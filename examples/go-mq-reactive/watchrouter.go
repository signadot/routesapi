package mqreactiverouter

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/signadot/routesapi/go-routesapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// Config provides the information to create a MQRouter
type Config struct {
	RouteServerAddr string // address of the routeserver
	Log             *slog.Logger
}

type watchMQRouter struct {
	*Config
	sandboxID   string
	baseline    *routesapi.BaselineWorkload
	grpcClient  routesapi.RoutesClient
	starting    bool
	startingMap map[string]string
	init        chan struct{}
	mu          sync.RWMutex
	routingMap  map[string]string // this is a map from routing key to sandbox ID
}

func NewWatchMQRouter(ctx context.Context, cfg *Config, b *routesapi.BaselineWorkload, sbID string) (*watchMQRouter, error) {
	// connect the route server
	conn, err := grpc.Dial(cfg.RouteServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	// create a route api client
	grpcClient := routesapi.NewRoutesClient(conn)
	// create an mq router
	mq := &watchMQRouter{
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

func (mq *watchMQRouter) run(ctx context.Context) {
	// watch loop
	for {
		watchClient, err := mq.grpcClient.WatchWorkloadRoutes(ctx, &routesapi.WorkloadRoutesRequest{
			BaselineWorkload: mq.baseline,
		})
		if err != nil {
			// don't retry if the context has been cancelled
			select {
			case <-ctx.Done():
				mq.Log.Info("context cancelled, closing")
				return
			default:
			}

			mq.Log.Error("couldn't watch workload routes", "error", err)
			<-time.After(3 * time.Second)
			continue
		}
		mq.Log.Debug("successfully got routes watch client")
		mq.readStream(ctx, watchClient)
	}
}

func (mq *watchMQRouter) readStream(ctx context.Context,
	watchClient routesapi.Routes_WatchWorkloadRoutesClient) {
	// put us in starting state
	mq.starting = true
	mq.startingMap = map[string]string{}

	// read route operations from the watch routes stream
	for {
		op, err := watchClient.Recv()
		if err != nil {
			// just return if the context has been cancelled
			select {
			case <-ctx.Done():
				return
			default:
			}

			// extract the grpc status
			grpcStatus, ok := status.FromError(err)
			if !ok {
				mq.Log.Error("watch routes stream error: no status", "error", err)
				break
			}
			switch grpcStatus.Code() {
			case codes.OK:
				mq.Log.Debug("watch routes error code is ok")
				goto PROCESS
			default:
				mq.Log.Error("watch routes error", "error", err)
				<-time.After(3 * time.Second)
			}
			break
		}

	PROCESS:
		// here we can be in two different states: constructing the initial
		// state (SYNCED op hasn't arrived yet) or receiving deltas
		if mq.starting {
			mq.processStartingOp(op)
		} else {
			mq.processDeltaOp(op)
		}
	}
}

func (mq *watchMQRouter) processStartingOp(op *routesapi.WorkloadRuleOp) {
	// no need to lock here, only one goroutine is acting on the starting fields
	switch op.Op {
	case routesapi.WatchOp_ADD:
		mq.startingMap[op.Rule.RoutingKey] = op.Rule.SandboxedWorkload.SandboxID
	case routesapi.WatchOp_SYNCED:
		mq.Log.Debug("synced")

		// update the routing map
		mq.mu.Lock()
		mq.routingMap = mq.startingMap
		mq.mu.Unlock()
		mq.Log.Debug("routing map", "routingMap", mq.routingMap)

		// move out of starting state
		mq.starting = false
		mq.startingMap = nil

		// declare ourselves as initialized
		select {
		case <-mq.init:
		default:
			close(mq.init)
		}
	default:
		mq.Log.Error("received unexpected watch op while starting", "op", op.Op.String())
	}
}

func (mq *watchMQRouter) processDeltaOp(op *routesapi.WorkloadRuleOp) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	switch op.Op {
	case routesapi.WatchOp_ADD:
		mq.routingMap[op.Rule.RoutingKey] = op.Rule.SandboxedWorkload.SandboxID
	case routesapi.WatchOp_REPLACE:
		mq.routingMap[op.Rule.RoutingKey] = op.Rule.SandboxedWorkload.SandboxID
	case routesapi.WatchOp_REMOVE:
		delete(mq.routingMap, op.Rule.RoutingKey)
	default:
		mq.Log.Error("received unexpected watch op while receiving deltas", "op", op.Op.String())
		return
	}
	mq.Log.Debug("new routing map", "routingMap", mq.routingMap)
}

func (mq *watchMQRouter) ShouldProcess(ctx context.Context, routingKey string) bool {
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
