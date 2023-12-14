package routesmq

import "log/slog"

// Config provides the information to create a MQRouter
type Config struct {
	Addr string // address of the routeserver
	Log  *slog.Logger
}

type MQRouter interface {
	ShouldProcess(routingKey string) bool
}
