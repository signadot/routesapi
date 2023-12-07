package workload

import (
	"context"
	"fmt"

	"github.com/signadot/routesapi"
	"github.com/signadot/routesapi/watched"
)

type Router interface {
	GetTarget(containerPort int32, rks ...string) string
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
	for _, rk := range rks {
		rr := r.watched.Get(rk)
		if rr == nil {
			continue
		}
		for i := range rr.PortRules {
			pr := rr.PortRules[i]
			if int32(pr.WorkloadPort) != containerPort {
				continue
			}
			if len(pr.Destinations) == 0 {
				continue
			}
			dest := pr.Destinations[0]
			return fmt.Sprintf("%s:%d", dest.Host, dest.Port)
		}
	}
	return ""
}
