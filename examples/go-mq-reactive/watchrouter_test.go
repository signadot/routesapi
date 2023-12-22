package mqreactiverouter

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/signadot/routesapi/go-routesapi"
)

func TestWatchMQRouter(t *testing.T) {
	if os.Getenv("GOTEST_MANUAL") == "" {
		t.Skip()
		return
	}
	cfg := &Config{
		RouteServerAddr: os.Getenv("TEST_ROUTE_SERVER_ADDR"),
		Log: slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
	}
	ctx := context.Background()

	sandboxName := os.Getenv("SIGNADOT_SANDBOX_NAME")
	baseline := &routesapi.BaselineWorkload{
		Kind:      "Deployment",
		Namespace: "hotrod",
		Name:      "route",
	}

	mq, err := NewWatchMQRouter(ctx, cfg, baseline, sandboxName)
	if err != nil {
		t.Error(err)
		return
	}
	ok := mq.ShouldProcess(ctx, os.Getenv("TEST_ROUTING_KEY"))
	t.Logf("got should process:\n%v\n", ok)
}
