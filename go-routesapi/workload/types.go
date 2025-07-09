package workload

import (
	"context"
)

type Router interface {
	GetTarget(containerPort int32, rks ...string) *Target
	GetTargetWithContext(ctx context.Context, containerPort int32, rks ...string) (*Target, error)
}

type Target struct {
	// Destination where traffic should be forwarded when not using the
	// traffic-manager. An empty string represents the baseline.
	Destination string `json:"destination,omitempty"`

	// TrafficManager specification, used to specify how to interact with
	// traffic-manager.  If absent, no traffic goes to TM.  If present, some
	// traffic may go to TM.
	TrafficManager *TrafficManager `json:"trafficManager,omitempty"`
}

type TrafficManager struct {
	// nextDestination specifies the signadot-next-host header content to be set
	// when forwarding traffic to traffic-manager
	NextDestination string `json:"nextDestination"`
	// if allTraffic is specified and true, indicates that any traffic goes to
	// traffic-manager, otherwise, only traffic with sd-traffic member key goes
	// to TM.
	AllTraffic bool `json:"allTraffic,omitempty"`
}
