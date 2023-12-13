package watched

import (
	"context"

	"github.com/signadot/routesapi/go-routesapi"
	"google.golang.org/protobuf/proto"
)

// BaselineWatched wraps [watched.Watched] with a [routesapi.BaselineWorkload].
type BaselineWatched interface {
	Get(rk string) *routesapi.WorkloadRule
	RoutesTo(rk string, sbID string) bool
}

type baselineWatched struct {
	*watched
	baseline *routesapi.BaselineWorkload
}

func NewBaselineWatched(ctx context.Context, cfg *Config, b *routesapi.BaselineWorkload) (BaselineWatched, error) {
	bb := proto.Clone(b).(*routesapi.BaselineWorkload)
	w, err := NewWatched(ctx, cfg, &routesapi.WorkloadRoutesRequest{
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

func (bw *baselineWatched) Get(rk string) *routesapi.WorkloadRule {
	return bw.watched.Get(bw.baseline, rk)
}

func (bw *baselineWatched) RoutesTo(rk string, sbID string) bool {
	return bw.watched.RoutesTo(bw.baseline, rk, sbID)
}
