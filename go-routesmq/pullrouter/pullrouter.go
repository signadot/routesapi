package pullrouter

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/signadot/routesapi/go-routesapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Config provides the information to create a MQRouter
type Config struct {
	RouteServerAddr string // address of the routeserver
	PullInterval    time.Duration
	Log             *slog.Logger
}

type pullMQRouter struct {
	*Config
	sandboxID  string
	baseline   *routesapi.BaselineWorkload
	grpcClient routesapi.RoutesClient
	init       chan struct{}
	mu         sync.RWMutex
	routingMap map[string]string // this is a map from routing key to sandbox ID
}

func NewPullMQRouter(ctx context.Context, cfg *Config, b *routesapi.BaselineWorkload, sbID string) (*pullMQRouter, error) {
	// connect the route server
	conn, err := grpc.Dial(cfg.RouteServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	// create a route api client
	grpcClient := routesapi.NewRoutesClient(conn)
	// create an mq router
	mq := &pullMQRouter{
		Config:     cfg,
		sandboxID:  sbID,
		baseline:   b,
		grpcClient: grpcClient,
		init:       make(chan struct{}),
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
	mq.Log.Debug("reloading routing rules", "baseline", mq.baseline)
	// load routes from route server
	resp, err := mq.grpcClient.GetWorkloadRoutes(ctx, &routesapi.WorkloadRoutesRequest{
		BaselineWorkload: mq.baseline,
	})
	if err != nil {
		mq.Log.Error("couldn't get workload routes", "error", err)
		return
	}

	// recompute the routing map
	routingMap := make(map[string]string, len(resp.Rules))
	for _, rule := range resp.Rules {
		routingMap[rule.RoutingKey] = rule.SandboxedWorkload.SandboxID
	}
	mq.Log.Debug("new routing map", "routingMap", routingMap)

	// update the routing map
	mq.mu.Lock()
	mq.routingMap = routingMap
	mq.mu.Unlock()

	// declare ourselves as initialized
	select {
	case <-mq.init:
	default:
		close(mq.init)
	}

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

	// there are 2 possible cases here:
	//
	// 1. we are a baseline workload (mq.sandboxID == ""), in which case we will
	// only process the message if there is no sandboxed workload for the given
	// routing key (in other words: mq.routingMap[routingKey] == "")
	//
	// 2. we are a sandboxed workload, in which case we will only process the
	// message if the routing key points to our sandbox ID
	return mq.routingMap[routingKey] == mq.sandboxID
}
