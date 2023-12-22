package mqreactiverouter

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/golang-collections/collections/set"
	"github.com/signadot/routesapi/go-routesapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type watchMQRouter struct {
	*Config
	grpcClient  routesapi.RoutesClient
	init        chan struct{}
	mu          sync.RWMutex
	startingSet *set.Set
	routingKeys *set.Set
}

func NewWatchMQRouter(ctx context.Context, cfg *Config) (*watchMQRouter, error) {
	// make sure we have a baseline
	if cfg.Baseline == nil {
		return nil, fmt.Errorf("empty baseline")
	}
	// connect the route server
	conn, err := grpc.Dial(cfg.RouteServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	// create a route api client
	grpcClient := routesapi.NewRoutesClient(conn)
	// create an mq router
	mq := &watchMQRouter{
		Config:     cfg,
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
		// create the gRPC request
		req := &routesapi.WorkloadRoutesRequest{
			BaselineWorkload: &routesapi.BaselineWorkload{
				Kind:      mq.Baseline.Kind,
				Namespace: mq.Baseline.Namespace,
				Name:      mq.Baseline.Name,
			},
		}
		if mq.isSandboxedWorkload() {
			req.DestinationSandbox = &routesapi.DestinationSandbox{
				Name: mq.Sandbox.Name,
			}
		}

		// start watching the stream
		watchClient, err := mq.grpcClient.WatchWorkloadRoutes(ctx, req)
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
	mq.startingSet = set.New()

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
		if mq.startingSet != nil {
			mq.processStartingOp(op)
		} else {
			mq.processDeltaOp(op)
		}
	}
}

func (mq *watchMQRouter) processStartingOp(op *routesapi.WorkloadRouteOp) {
	// no need to lock here, only one goroutine is acting on the starting fields
	switch op.Op {
	case routesapi.WatchOp_ADD:
		mq.startingSet.Insert(op.Route.RoutingKey)
	case routesapi.WatchOp_SYNCED:
		mq.Log.Debug("synced")

		// update the routing map
		mq.mu.Lock()
		mq.routingKeys = mq.startingSet
		mq.mu.Unlock()
		mq.Log.Debug("initial routing keys", "routingKeys", set2String(mq.routingKeys))

		// move out of starting state
		mq.startingSet = nil

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

func (mq *watchMQRouter) processDeltaOp(op *routesapi.WorkloadRouteOp) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	switch op.Op {
	case routesapi.WatchOp_ADD:
		mq.routingKeys.Insert(op.Route.RoutingKey)
	case routesapi.WatchOp_REPLACE:
		// do nothing, we only care about routing keys
	case routesapi.WatchOp_REMOVE:
		mq.routingKeys.Remove(op.Route.RoutingKey)
	default:
		mq.Log.Error("received unexpected watch op while receiving deltas", "op", op.Op.String())
		return
	}
	mq.Log.Debug("new routing keys", "routingKeys", set2String(mq.routingKeys))
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

	if mq.isSandboxedWorkload() {
		// we are a sandboxed workload, only accept the received routing keys
		return mq.routingKeys.Has(routingKey)
	}
	// we are a baseline workload, ignore received routing keys (they belong
	// to sandboxed workloads)
	return !mq.routingKeys.Has(routingKey)
}

func set2String(s *set.Set) string {
	values := ""
	s.Do(func(val any) {
		if values == "" {
			values = fmt.Sprintf("%v", val)
		} else {
			values += fmt.Sprintf(", %v", val)
		}
	})
	return fmt.Sprintf("[%s]", values)
}
