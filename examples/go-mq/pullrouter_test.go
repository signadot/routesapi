package mqrouter

import (
	"context"
	"os"
	"testing"
	"time"

	"log/slog"

	"github.com/signadot/routesapi/go-routesapi"
)

func TestPullMQRouter(t *testing.T) {
	if os.Getenv("GOTEST_MANUAL") == "" {
		t.Skip()
		return
	}
	cfg := &Config{
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
