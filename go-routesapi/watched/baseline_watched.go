package watched

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/signadot/routesapi/go-routesapi"
	"google.golang.org/protobuf/proto"
)

// BaselineWatched wraps [watched.Watched] with a [routesapi.BaselineWorkload].
type BaselineWatched interface {
	// Get returns a workload routing rule associated with
	// the routing key rk for the associated baseline workload.
	Get(rk string) *routesapi.WorkloadRoutingRule

	// RoutesTo returns true if traffic destined to the associated
	// baseline workload with routing key rk should be directed to
	// the sandbox named sbName.
	RoutesTo(rk string, sbName string) bool
}

type baselineWatched struct {
	*watched
	baseline *routesapi.BaselineWorkload
}

// NewBaselineWatched creates a BaselineWatched instance using
// ctx for the underlying grpc watch rpc, using the routeserver
// specified in [cfg.Addr], and associated with the baseline
// specified in b.
func NewBaselineWatched(ctx context.Context, cfg *Config, b *routesapi.BaselineWorkload) (BaselineWatched, error) {
	bb := proto.Clone(b).(*routesapi.BaselineWorkload)
	w, err := NewWatched(ctx, cfg, &routesapi.WorkloadRoutingRulesRequest{
		BaselineWorkload: bb,
	})
	if err != nil {
		return nil, err
	}
	return &baselineWatched{
		watched:  w.(*watched),
		baseline: bb,
	}, nil
}

// BaselineWatchedFromEnv attempts to construct a BaselineWatched instance
// using configuration from the environment, by using
// [BaselineFromEnv] and [GetRouteServerAddr] and calling
// [NewBaselineWatched].  The context associated with the watch
// is the result of calling [context.Background] and the logger
// is the result of calling [slog.Default].  For more control,
// please use [NewBaselineWatched] directly.
func BaselineWatchedFromEnv() (BaselineWatched, error) {
	baseline, err := BaselineFromEnv()
	if err != nil {
		return nil, fmt.Errorf("could not get baseline from env: %w", err)
	}
	cfg := &Config{
		Addr: GetRouteServerAddr(),
		Log:  slog.Default(),
	}
	apiBaseline := &routesapi.BaselineWorkload{
		Kind:      baseline.Kind,
		Namespace: baseline.Namespace,
		Name:      baseline.Name,
	}
	return NewBaselineWatched(context.Background(), cfg, apiBaseline)
}

func (bw *baselineWatched) Get(rk string) *routesapi.WorkloadRoutingRule {
	return bw.watched.Get(bw.baseline, rk)
}

func (bw *baselineWatched) RoutesTo(rk string, sbName string) bool {
	return bw.watched.RoutesTo(bw.baseline, rk, sbName)
}
