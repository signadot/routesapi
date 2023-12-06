package watched

import (
	"context"
	"time"

	"github.com/signadot/routesapi"
	"github.com/signadot/routesapi/internal/queue"
	"github.com/signadot/routesapi/internal/retrypolicy"
	"google.golang.org/grpc"
)

// watcher watches (with retries on error) a set of WorkloadRules.
type watcher struct {
	*Config

	grpcClient   routesapi.RoutesClient
	watchContext context.Context                  // watch context
	watchArg     *routesapi.WorkloadRoutesRequest // watch request
	watchOpts    []grpc.CallOption                // options
	watched      *watched                         // output watched

	// underlying grpc watch (without retries)
	underWatch routesapi.Routes_WatchWorkloadRoutesClient

	// pending requeusts
	pending       *queue.Queue[*routesapi.WorkloadRuleOp]
	pendingSynced bool
}

func (w *watcher) Recv() (*routesapi.WorkloadRuleOp, error) {
	var (
		op  *routesapi.WorkloadRuleOp
		err error
	)

	for {
		if w.pending.Len() > 0 {
			//w.Log.Debug("queue")
			op = w.pending.Pop()
			w.watched.handleOp(op)
			return op, nil
		}
		if w.underWatch != nil {
			//w.Log.Debug("underwatch")
			op, err = w.underWatch.Recv()
			if err == nil {
				w.watched.handleOp(op)
				return op, nil
			}
		}
		w.retryRefetch(err)
		err = nil
	}
}

func (w *watcher) retryRefetch(err error) {
	if err != nil {
		w.Log.Error("restart because of error receiving", "error", err)
	} else {
		w.Log.Info("starting watch")
	}
	luby := retrypolicy.DefaultLubyNexter()
	ticker := time.NewTicker(luby.Next())
	defer ticker.Stop()
	for {
		err := w.slurp()
		if err == nil {
			return
		}
		w.Log.Error("slurp", "error", err)
		<-ticker.C
		ticker.Reset(luby.Next())
	}
}

func (w *watcher) slurp() error {
	uw, err := w.grpcClient.WatchWorkloadRoutes(w.watchContext, w.watchArg, w.watchOpts...)
	if err != nil {
		return err
	}
	d := map[key]*routesapi.WorkloadRule{}
	for {
		op, err := uw.Recv()
		if err != nil {
			return err
		}
		switch op.Op {
		case routesapi.WatchOp_ADD:
			k, v := kv(op.Rule)
			d[*k] = v
		case routesapi.WatchOp_SYNCED:
			w.Log.Debug("synced")
			w.update(d, uw)
			return nil
		default:
			panic("impossible")
		}
	}
}

func (w *watcher) update(d map[key]*routesapi.WorkloadRule, uw routesapi.Routes_WatchWorkloadRoutesClient) {
	for k, cur := range d {
		_, ok := w.watched.D[k]
		pendingOp := routesapi.WatchOp_ADD
		if ok {
			pendingOp = routesapi.WatchOp_REPLACE
		}
		w.pending.Push(&routesapi.WorkloadRuleOp{
			Op:   pendingOp,
			Rule: cur,
		})
	}
	for k, old := range w.watched.D {
		_, present := d[k]
		if present {
			continue
		}
		w.pending.Push(&routesapi.WorkloadRuleOp{
			Op:   routesapi.WatchOp_REMOVE,
			Rule: old,
		})
	}
	if !w.pendingSynced {
		w.Log.Debug("adding pending sync")
		w.pending.Push(&routesapi.WorkloadRuleOp{
			Op: routesapi.WatchOp_SYNCED,
		})
		w.pendingSynced = true
	}
	w.underWatch = uw
}
