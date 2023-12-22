package mqreactiverouter

import (
	"log/slog"
)

// Config provides the information to create a MQRouter
type Config struct {
	// Logger interface
	Log *slog.Logger
	// Routeserver address (just host and port)
	RouteServerAddr string
	// Baseline workload reference
	Baseline *BaselineWorkload
	// Sandbox reference, nil if not running within a sandbox
	Sandbox *Sandbox
}

func (c *Config) isSandboxedWorkload() bool {
	if c.Sandbox == nil {
		return false
	}
	return c.Sandbox.Name != ""
}

// A BaselineWorkload identifies the baseline workload.
type BaselineWorkload struct {
	Kind      string
	Namespace string
	Name      string
}

// A Sandbox information
type Sandbox struct {
	Name string
}
