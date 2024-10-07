package workload

import (
	"context"
	"fmt"

	"github.com/signadot/routesapi/go-routesapi"
	"github.com/signadot/routesapi/go-routesapi/watched"
)

type Router interface {
	GetTarget(containerPort int32, rks ...string) string
	GetTargetWithContext(ctx context.Context, containerPort int32, rks ...string) (string, error)
}

type router struct {
	watched watched.BaselineWatched
}

func NewRouter(cfg *Config) (Router, error) {
	w, err := watched.NewBaselineWatched(context.Background(),
		&watched.Config{
			Addr: cfg.RouteServerAddr,
			Log:  cfg.Log,
		},
		&routesapi.BaselineWorkload{
			Kind:      cfg.Baseline.Kind,
			Namespace: cfg.Baseline.Namespace,
			Name:      cfg.Baseline.Name,
		})
	if err != nil {
		return nil, err
	}
	return &router{w}, nil
}

func (r *router) GetTarget(containerPort int32, rks ...string) string {
	res, _ := r.GetTargetWithContext(context.Background(), containerPort, rks...)
	return res
}

func (r *router) GetTargetWithContext(ctx context.Context, containerPort int32, rks ...string) (string, error) {
	for _, rk := range rks {
		rr, err := r.watched.GetWithContext(ctx, rk)
		if err != nil {
			return "", err
		}
		if rr == nil {
			continue
		}
		for i := range rr.Mappings {
			pr := rr.Mappings[i]
			if int32(pr.WorkloadPort) != containerPort {
				continue
			}
			if len(pr.Destinations) == 0 {
				continue
			}
			dest := pr.Destinations[0]
			return fmt.Sprintf("%s:%d", dest.Host, dest.Port), nil
		}
	}
	return "", nil
}
