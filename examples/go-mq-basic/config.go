package mqbasicrouter

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// Config provides the information to create a MQRouter
type Config struct {
	Log             *slog.Logger
	RouteServerAddr string // address of the routeserver
	PullInterval    time.Duration
	Baseline        *BaselineWorkload
	SandboxName     string
}

func (c *Config) getRouteServerURL() string {
	// construct base request
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/v1/workloads/routes", c.RouteServerAddr), nil)
	if err != nil {
		panic(fmt.Sprintf("Error building request, %v", err))
	}

	// add query params
	q := req.URL.Query()
	q.Add("baselineKind", c.Baseline.Kind)
	q.Add("baselineNamespace", c.Baseline.Namespace)
	q.Add("baselineName", c.Baseline.Name)
	q.Add("destinationSandboxName", c.SandboxName) // if SandboxName is empty it will return all sandboxes
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

// A BaselineWorkload identifies a given baseline workload.
type BaselineWorkload struct {
	Kind      string
	Namespace string
	Name      string
}
