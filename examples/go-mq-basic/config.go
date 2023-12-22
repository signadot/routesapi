package mqbasicrouter

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// Config provides the information to create a MQRouter
type Config struct {
	// Logger interface
	Log *slog.Logger
	// Routeserver address (just host and port)
	RouteServerAddr string
	// Interval for refreshing local routes
	PullInterval time.Duration
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

func (c *Config) getRouteServerURL() (string, error) {
	// make sure we have a baseline
	if c.Baseline == nil {
		return "", fmt.Errorf("empty baseline")
	}

	// construct base request
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/v1/workloads/routes", c.RouteServerAddr), nil)
	if err != nil {
		return "", fmt.Errorf("error building request, %w", err)
	}

	// add query params
	q := req.URL.Query()
	q.Add("baselineKind", c.Baseline.Kind)
	q.Add("baselineNamespace", c.Baseline.Namespace)
	q.Add("baselineName", c.Baseline.Name)
	if c.isSandboxedWorkload() {
		q.Add("destinationSandboxName", c.Sandbox.Name)
	}
	req.URL.RawQuery = q.Encode()

	// return the resulting URL
	return req.URL.String(), nil
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
