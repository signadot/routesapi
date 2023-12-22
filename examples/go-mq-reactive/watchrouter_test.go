package mqreactiverouter

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

func TestWatchMQRouter(t *testing.T) {
	if os.Getenv("GOTEST_MANUAL") == "" {
		t.Skip()
		return
	}

	routeServerAddr := os.Getenv("TEST_ROUTE_SERVER_ADDR")
	if routeServerAddr == "" {
		// use default location
		routeServerAddr = "routeserver.signadot.svc:7777"
	}

	cfg := &Config{
		Log: slog.New(
			slog.NewTextHandler(os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelDebug,
				}),
		),
		RouteServerAddr: routeServerAddr,
		Baseline: &BaselineWorkload{
			Kind:      "Deployment",
			Namespace: "hotrod",
			Name:      "route",
		},
	}

	sandboxName := os.Getenv("SIGNADOT_SANDBOX_NAME")
	if sandboxName != "" {
		// we are running within a sandbox, set the sandbox reference
		cfg.Sandbox = &Sandbox{
			Name: sandboxName,
		}
	}

	ctx := context.Background()
	mq, err := NewWatchMQRouter(ctx, cfg)
	if err != nil {
		t.Error(err)
		return
	}
	ok := mq.ShouldProcess(ctx, os.Getenv("TEST_ROUTING_KEY"))
	t.Logf("got should process:\n%v\n", ok)
}
