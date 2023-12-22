package mqbasicrouter

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/golang-collections/collections/set"
	"github.com/signadot/routesapi/go-routesapi/models"
)

type pullMQRouter struct {
	*Config
	routeServerURL string
	init           chan struct{}
	mu             sync.RWMutex
	routingKeys    *set.Set
}

func NewPullMQRouter(ctx context.Context, cfg *Config) (*pullMQRouter, error) {
	// create an mq router
	mq := &pullMQRouter{
		Config:         cfg,
		routeServerURL: cfg.getRouteServerURL(),
		init:           make(chan struct{}),
	}
	// run the mq router
	go mq.run(ctx)
	return mq, nil
}

func (mq *pullMQRouter) run(ctx context.Context) {
	reloadTicker := time.NewTicker(mq.PullInterval)
	defer reloadTicker.Stop()
	for {
		mq.reload(ctx)
		select {
		case <-ctx.Done():
			mq.Log.Info("context cancelled, closing")
			return
		case <-reloadTicker.C:
		}
	}
}

func (mq *pullMQRouter) reload(ctx context.Context) {
	mq.Log.Debug("reloading routes", "baseline", mq.Baseline)

	// load routes from route server
	resp, err := mq.getRoutes()
	if err != nil {
		mq.Log.Error("couldn't get workload routes", "error", err)
		return
	}

	// collect received routing keys
	rkSet := set.New()
	rkList := make([]string, 0, len(resp.Routes))
	for _, route := range resp.Routes {
		rkSet.Insert(route.RoutingKey)
		rkList = append(rkList, route.RoutingKey)
	}
	mq.Log.Debug("routing keys received", "routingKeys", rkList)

	// update received routing keys
	mq.mu.Lock()
	mq.routingKeys = rkSet
	mq.mu.Unlock()

	// declare ourselves as initialized
	select {
	case <-mq.init:
	default:
		close(mq.init)
	}
}

func (mq *pullMQRouter) getRoutes() (*models.WorkloadRoutesResponse, error) {
	mq.Log.Debug("sending request to routeserver", "url", mq.routeServerURL)

	// send request to route server
	resp, err := http.Get(mq.routeServerURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse response
	routesResp := &models.WorkloadRoutesResponse{}
	err = json.NewDecoder(resp.Body).Decode(routesResp)
	if err != nil {
		return nil, err
	}
	return routesResp, nil
}

func (mq *pullMQRouter) ShouldProcess(ctx context.Context, routingKey string) bool {
	// wait until initialized or the context is done
	select {
	case <-mq.init:
	case <-ctx.Done():
		return false
	}

	// obtain a read lock
	mq.mu.RLock()
	defer mq.mu.RUnlock()

	if mq.SandboxName == "" {
		// we are a baseline workload, ignore received routing keys (they belong
		// to sandboxed workloads)
		return !mq.routingKeys.Has(routingKey)
	}
	// we are a sandboxed workload, only accept the received routing keys
	return mq.routingKeys.Has(routingKey)
}
