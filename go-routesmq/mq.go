package routesmq

import (
	"context"

	"github.com/signadot/routesapi/go-routesapi"
	"github.com/signadot/routesapi/go-routesmq/pullrouter"
	"github.com/signadot/routesapi/go-routesmq/watchrouter"
)

type MQRouter interface {
	ShouldProcess(ctx context.Context, routingKey string) bool
}

func NewPullMQRouter(ctx context.Context, cfg *pullrouter.Config,
	b *routesapi.BaselineWorkload, sbID string) (MQRouter, error) {
	return pullrouter.NewPullMQRouter(ctx, cfg, b, sbID)
}

func NewWatchMQRouter(ctx context.Context, cfg *watchrouter.Config,
	b *routesapi.BaselineWorkload, sbID string) (MQRouter, error) {
	return watchrouter.NewWatchMQRouter(ctx, cfg, b, sbID)
}
