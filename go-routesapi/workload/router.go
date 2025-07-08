package workload

import (
	"context"
	"fmt"

	"github.com/signadot/routesapi/go-routesapi"
	"github.com/signadot/routesapi/go-routesapi/watched"
)

type router struct {
	watched watched.BaselineWatched
}

func NewRouter(cfg *Config) (Router, error) {
	var cliInfo *routesapi.ClientInfo
	if cfg.ClientInfo != nil {
		cliInfo = cfg.ClientInfo
	} else {
		// by default support virtual workloads
		cliInfo = &routesapi.ClientInfo{
			EnableVirtualWorkloads: true,
		}
	}
	w, err := watched.NewBaselineWatched(context.Background(),
		&watched.Config{
			Addr: cfg.RouteServerAddr,
			Log:  cfg.Log,
		},
		&routesapi.BaselineWorkload{
			Kind:      cfg.Baseline.Kind,
			Namespace: cfg.Baseline.Namespace,
			Name:      cfg.Baseline.Name,
		},
		cliInfo)
	if err != nil {
		return nil, err
	}
	return &router{w}, nil
}

func (r *router) GetTarget(containerPort int32, rks ...string) *Target {
	res, _ := r.GetTargetWithContext(context.Background(), containerPort, rks...)
	return res
}

func (r *router) GetTargetWithContext(ctx context.Context, containerPort int32, rks ...string) (*Target, error) {
	for _, rk := range rks {
		rr, err := r.watched.GetWithContext(ctx, rk)
		if err != nil {
			return nil, err
		}
		if rr == nil {
			continue
		}
		for i := range rr.Mappings {
			pr := rr.Mappings[i]
			if int32(pr.WorkloadPort) != containerPort {
				continue
			}
			if len(pr.Destinations) == 0 && (pr.TrafficManager == nil ||
				len(pr.TrafficManager.NextDestinations) == 0) {
				continue
			}

			var target Target
			if len(pr.Destinations) > 0 {
				dest := pr.Destinations[0]
				target.Destination = fmt.Sprintf("%s:%d", dest.Host, dest.Port)
			}
			if pr.TrafficManager != nil && len(pr.TrafficManager.NextDestinations) > 0 {
				dest := pr.TrafficManager.NextDestinations[0]
				target.TrafficManager = &TrafficManager{
					NextDestination: fmt.Sprintf("%s:%d", dest.Host, dest.Port),
					AllTraffic:      pr.TrafficManager.AllTraffic,
				}
			}
			return &target, nil
		}
	}
	return nil, nil
}
