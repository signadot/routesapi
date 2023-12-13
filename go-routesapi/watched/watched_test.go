package watched

import (
	"context"
	"os"
	"testing"

	"log/slog"

	"github.com/signadot/routesapi/go-routesapi"
)

func TestWatched(t *testing.T) {
	if os.Getenv("GOTEST_MANUAL") == "" {
		t.Skip()
		return
	}
	cfg := &Config{
		Addr: ":7777",
		Log: slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
	}
	ctx := context.Background()
	w, err := NewWatched(ctx, cfg, &routesapi.WorkloadRoutesRequest{
		BaselineWorkload: &routesapi.BaselineWorkload{}})
	if err != nil {
		t.Error(err)
		return
	}
	rv := w.Get(&routesapi.BaselineWorkload{
		Kind:      "Deployment",
		Namespace: "hotrod",
		Name:      "route",
	}, "uhhny1x6s9q9d")
	t.Logf("got rv:\n%v\n", rv)
}
