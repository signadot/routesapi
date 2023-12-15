package routesmq

import (
	"context"
	"os"
	"testing"
	"time"

	"log/slog"

	"github.com/signadot/routesapi/go-routesapi"
	"github.com/signadot/routesapi/go-routesmq/pullrouter"
	"github.com/signadot/routesapi/go-routesmq/watchrouter"
)

func TestPullMQRouter(t *testing.T) {
	if os.Getenv("GOTEST_MANUAL") == "" {
		t.Skip()
		return
	}
	cfg := &pullrouter.Config{
		RouteServerAddr: os.Getenv("TEST_ROUTE_SERVER_ADDR"),
		PullInterval:    10 * time.Second,
		Log: slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
	}
	ctx := context.Background()

	sandboxID := os.Getenv("SIGNADOT_SANDBOX_ROUTING_KEY")
	baseline := &routesapi.BaselineWorkload{
		Kind:      "Deployment",
		Namespace: "hotrod",
		Name:      "route",
	}

	mq, err := NewPullMQRouter(ctx, cfg, baseline, sandboxID)
	if err != nil {
		t.Error(err)
		return
	}
	ok := mq.ShouldProcess(ctx, os.Getenv("TEST_ROUTING_KEY"))
	t.Logf("got should process:\n%v\n", ok)
}

func TestWatchMQRouter(t *testing.T) {
	if os.Getenv("GOTEST_MANUAL") == "" {
		t.Skip()
		return
	}
	cfg := &watchrouter.Config{
		RouteServerAddr: os.Getenv("TEST_ROUTE_SERVER_ADDR"),
		Log: slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
	}
	ctx := context.Background()

	sandboxID := os.Getenv("SIGNADOT_SANDBOX_ROUTING_KEY")
	baseline := &routesapi.BaselineWorkload{
		Kind:      "Deployment",
		Namespace: "hotrod",
		Name:      "route",
	}

	mq, err := NewWatchMQRouter(ctx, cfg, baseline, sandboxID)
	if err != nil {
		t.Error(err)
		return
	}
	ok := mq.ShouldProcess(ctx, os.Getenv("TEST_ROUTING_KEY"))
	t.Logf("got should process:\n%v\n", ok)
}
