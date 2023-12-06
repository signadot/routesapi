package watched

import "github.com/signadot/routesapi"

// key represents a way to identify a [routes.WorkloadRule].
type key struct {
	RoutingKey string `json:"routingKey"`
	Baseline
}

func newKey(rk string, b *routesapi.BaselineWorkload) *key {
	return &key{
		RoutingKey: rk,
		Baseline: Baseline{
			Namespace: b.Namespace,
			Name:      b.Name,
			Kind:      b.Kind,
		},
	}
}
