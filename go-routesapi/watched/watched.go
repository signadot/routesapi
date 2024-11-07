package watched

import (
	"context"
	"sync"
	"time"

	"github.com/signadot/routesapi/go-routesapi"
	"github.com/signadot/routesapi/go-routesapi/internal/indices"
	"github.com/signadot/routesapi/go-routesapi/internal/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/proto"
)

// Watched provides an interface
type Watched interface {
	// Get returns a [routesapi.WorkloadRoutingRule] which indicates where to direct
	// requests originally destined to baseline workload baseline with routing
	// key rk.  Get returns nil, if no such rule exists.
	Get(baseline *routesapi.BaselineWorkload, rk string) *routesapi.WorkloadRoutingRule

	// Same as Get() but with context handling
	GetWithContext(ctx context.Context, baseline *routesapi.BaselineWorkload,
		rk string) (*routesapi.WorkloadRoutingRule, error)

	// RoutesTo indicates whether or not a request originally destined to
	// baseline workload with routing key rk should be delivered to the
	// corresponding sandboxed workload associated with a sandbox name (sbName).
	RoutesTo(baseline *routesapi.BaselineWorkload, rk, sbName string) bool

	// Same as RoutesTo() but with context handling
	RoutesToWithContext(ctx context.Context, baseline *routesapi.BaselineWorkload,
		rk, sbName string) (bool, error)
}

type watched struct {
	sync.RWMutex
	synced chan struct{}
	D      map[key]*routesapi.WorkloadRoutingRule
	I      indices.Index[key]
}

// NewWatched creates a Watched.  The set of the workload rules returned from
// the returned Watched corresponds to those specified in q.
func NewWatched(ctx context.Context, cfg *Config, q *routesapi.WorkloadRoutingRulesRequest) (Watched, error) {
	conn, err := grpc.Dial(cfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    5 * time.Second,
			Timeout: 15 * time.Second,
		}))
	if err != nil {
		return nil, err
	}
	tmpIn := proto.Clone(q).(*routesapi.WorkloadRoutingRulesRequest)
	grpcClient := routesapi.NewRoutesClient(conn)
	watcher := &watcher{
		Config:       cfg,
		grpcClient:   grpcClient,
		watchContext: ctx,
		watchOpts:    nil,
		watchArg:     tmpIn,
		watched:      newWatched(),
		pending:      queue.New[*routesapi.WorkloadRoutingRuleOp](0),
	}
	go func() {
		for {
			_, err := watcher.Recv()
			if err != nil {
				// TODO propagate non-restart errors if possible
				// for now, this doesn't, it just restarts on error
				cfg.Log.Error("error with retry watch receive", "error", err)
			}
			select {
			case <-ctx.Done():
				cfg.Log.Info("exiting watcher, context done")
				return
			default:
			}
		}
	}()
	return watcher.watched, nil
}

func newWatched() *watched {
	return &watched{
		synced: make(chan struct{}),
		D:      make(map[key]*routesapi.WorkloadRoutingRule),
		I:      make(indices.Index[key]),
	}
}

func (w *watched) Get(baseline *routesapi.BaselineWorkload, rk string) *routesapi.WorkloadRoutingRule {
	res, _ := w.GetWithContext(context.Background(), baseline, rk)
	return res
}

func (w *watched) GetWithContext(ctx context.Context, baseline *routesapi.BaselineWorkload,
	rk string) (*routesapi.WorkloadRoutingRule, error) {
	select {
	case <-w.synced:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	key := newKey(rk, baseline)
	w.RLock()
	defer w.RUnlock()
	return w.D[*key], nil
}

func (w *watched) RoutesTo(b *routesapi.BaselineWorkload, rk, sbName string) bool {
	res, _ := w.RoutesToWithContext(context.Background(), b, rk, sbName)
	return res
}

func (w *watched) RoutesToWithContext(ctx context.Context, b *routesapi.BaselineWorkload,
	rk, sbName string) (bool, error) {
	select {
	case <-w.synced:
	case <-ctx.Done():
		return false, ctx.Err()
	}
	key := newKey(rk, b)
	w.RLock()
	defer w.RUnlock()
	return w.I.Get(sbName)[*key], nil
}

func (w *watched) set(rr *routesapi.WorkloadRoutingRule) {
	k, v := kv(rr)
	w.Lock()
	defer w.Unlock()
	w.D[*k] = v
	w.I.Add(rr.DestinationSandbox.Name, *k)
}

func kv(rr *routesapi.WorkloadRoutingRule) (*key, *routesapi.WorkloadRoutingRule) {
	key := newKey(rr.RoutingKey, rr.Baseline)
	// deep copy the rule
	resRule := proto.Clone(rr).(*routesapi.WorkloadRoutingRule)
	return key, resRule
}

func (w *watched) remove(rr *routesapi.WorkloadRoutingRule) {
	key := newKey(rr.RoutingKey, rr.Baseline)
	w.Lock()
	defer w.Unlock()
	delete(w.D, *key)
	w.I.Remove(rr.DestinationSandbox.Name, *key)
}

func (w *watched) sync() {
	select {
	case <-w.synced:
	default:
		close(w.synced)
	}
}

func (w *watched) handleOp(op *routesapi.WorkloadRoutingRuleOp) {
	switch op.Op {
	case routesapi.WatchOp_ADD, routesapi.WatchOp_REPLACE:
		w.set(op.Route)
	case routesapi.WatchOp_REMOVE:
		w.remove(op.Route)
	case routesapi.WatchOp_SYNCED:
		w.sync()
	}
}
