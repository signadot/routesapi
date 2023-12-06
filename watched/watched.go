package watched

import (
	"context"
	"sync"

	"github.com/signadot/routesapi"
	"github.com/signadot/routesapi/internal/indices"
	"github.com/signadot/routesapi/internal/queue"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Watched provides an interface
type Watched interface {

	// Get returns a [routesapi.WorkloadRule] which indicates where to direct
	// requests originally destined to baseline workload baseline with
	// routing key rk.  Get returns nil, if no such rule exists.
	Get(baseline *routesapi.BaselineWorkload, rk string) *routesapi.WorkloadRule

	// RoutesTo indicates whether or not a request originally destined
	// to baseline workload with routing key rk should be delivered to the
	// corresponding sandboxed workload associated with sbID.
	RoutesTo(baseline *routesapi.BaselineWorkload, rk, sbID string) bool
}

type watched struct {
	sync.RWMutex
	synced chan struct{}
	D      map[key]*routesapi.WorkloadRule
	I      indices.Index[key]
}

// NewWatched creates a Watched.  The set of the workload rules
// returned from the returned Watched corresponds to those
// specified in q.
func NewWatched(ctx context.Context, cfg *Config, q *routesapi.WorkloadRoutesRequest) (Watched, error) {
	conn, err := grpc.Dial(cfg.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	tmpIn := proto.Clone(q).(*routesapi.WorkloadRoutesRequest)
	grpcClient := routesapi.NewRoutesClient(conn)
	watcher := &watcher{
		Config:       cfg,
		grpcClient:   grpcClient,
		watchContext: ctx,
		watchOpts:    nil,
		watchArg:     tmpIn,
		watched:      newWatched(),
		pending:      queue.New[*routesapi.WorkloadRuleOp](0),
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
		D:      make(map[key]*routesapi.WorkloadRule),
		I:      make(indices.Index[key]),
	}
}

func (w *watched) Get(baseline *routesapi.BaselineWorkload, rk string) *routesapi.WorkloadRule {
	key := newKey(rk, baseline)
	<-w.synced
	w.RLock()
	defer w.RUnlock()
	return w.D[*key]
}

func (w *watched) RoutesTo(b *routesapi.BaselineWorkload, rk, sbID string) bool {
	key := newKey(rk, b)
	<-w.synced
	w.RLock()
	defer w.RUnlock()
	return w.I.Get(sbID)[*key]
}

func (w *watched) set(rr *routesapi.WorkloadRule) {
	k, v := kv(rr)
	w.Lock()
	defer w.Unlock()
	w.D[*k] = v
	w.I.Add(rr.SandboxedWorkload.SandboxID, *k)
}

func kv(rr *routesapi.WorkloadRule) (*key, *routesapi.WorkloadRule) {
	key := newKey(rr.RoutingKey, rr.SandboxedWorkload.Baseline)
	// deep copy the rule
	resRule := proto.Clone(rr).(*routesapi.WorkloadRule)
	return key, resRule
}

func (w *watched) remove(rr *routesapi.WorkloadRule) {
	key := newKey(rr.RoutingKey, rr.SandboxedWorkload.Baseline)
	w.Lock()
	defer w.Unlock()
	delete(w.D, *key)
	w.I.Remove(rr.SandboxedWorkload.SandboxID, *key)
}

func (w *watched) sync() {
	select {
	case <-w.synced:
	default:
		close(w.synced)
	}
}

func (w *watched) handleOp(op *routesapi.WorkloadRuleOp) {
	switch op.Op {
	case routesapi.WatchOp_ADD, routesapi.WatchOp_REPLACE:
		w.set(op.Rule)
	case routesapi.WatchOp_REMOVE:
		w.remove(op.Rule)
	case routesapi.WatchOp_SYNCED:
		w.sync()
	}
}
