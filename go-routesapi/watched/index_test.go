package watched

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/signadot/routesapi/go-routesapi"
	"github.com/signadot/routesapi/go-routesapi/internal/queue"
	"google.golang.org/grpc"
)

func rule(rk, sb string) *routesapi.WorkloadRoutingRule {
	return &routesapi.WorkloadRoutingRule{
		RoutingKey: rk,
		Baseline: &routesapi.BaselineWorkload{
			Kind:      "Deployment",
			Namespace: "ns",
			Name:      "svc",
		},
		DestinationSandbox: &routesapi.DestinationSandbox{Name: sb},
	}
}

func TestReplaceChangesSandbox(t *testing.T) {
	w := newWatched()
	w.handleOp(&routesapi.WorkloadRoutingRuleOp{Op: routesapi.WatchOp_SYNCED})
	b := rule("rk", "").Baseline

	w.handleOp(&routesapi.WorkloadRoutingRuleOp{Op: routesapi.WatchOp_ADD, Route: rule("rk", "A")})
	w.handleOp(&routesapi.WorkloadRoutingRuleOp{Op: routesapi.WatchOp_REPLACE, Route: rule("rk", "B")})

	if got := w.Get(b, "rk").DestinationSandbox.Name; got != "B" {
		t.Errorf("Get: got sandbox %q, want B", got)
	}
	if w.RoutesTo(b, "rk", "A") {
		t.Error("RoutesTo(A) true after replace to B")
	}
	if !w.RoutesTo(b, "rk", "B") {
		t.Error("RoutesTo(B) false after replace to B")
	}

	// remove with the op naming the stale sandbox still clears the index
	w.handleOp(&routesapi.WorkloadRoutingRuleOp{Op: routesapi.WatchOp_REMOVE, Route: rule("rk", "A")})
	if w.Get(b, "rk") != nil {
		t.Error("Get non-nil after remove")
	}
	if w.RoutesTo(b, "rk", "A") || w.RoutesTo(b, "rk", "B") {
		t.Error("RoutesTo true after remove")
	}
	if len(w.I) != 0 {
		t.Errorf("index not empty after remove: %v", w.I)
	}
}

// failingClient fails every watch attempt.
type failingClient struct {
	routesapi.RoutesClient
}

func (failingClient) WatchWorkloadRoutingRules(context.Context, *routesapi.WorkloadRoutingRulesRequest,
	...grpc.CallOption) (routesapi.Routes_WatchWorkloadRoutingRulesClient, error) {
	return nil, errors.New("unavailable")
}

func TestRecvStopsOnContextDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	w := &watcher{
		Config:       &Config{Log: slog.New(slog.NewTextHandler(io.Discard, nil))},
		grpcClient:   failingClient{},
		watchContext: ctx,
		watched:      newWatched(),
		pending:      queue.New[*routesapi.WorkloadRoutingRuleOp](0),
	}
	done := make(chan error, 1)
	go func() {
		_, err := w.Recv()
		done <- err
	}()
	time.Sleep(50 * time.Millisecond)
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Errorf("got %v, want context.Canceled", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Recv did not return after context cancel")
	}
}
