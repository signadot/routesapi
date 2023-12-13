# Sidecar Routing with the Routes API Go SDK

The Signadot DevMesh sidecar implements routing using
the Routes API Go SDK.

Istio provides a Go plugin and supports [WASM plugins](https://istio.io/latest/docs/reference/config/proxy_extensions/wasm-plugin/).

Some infrastructure teams build their own service mesh sidecars.  While capable
service meshes exist, this option can be useful if there are legacy
considerations which don't work within a classic service mesh or for
implementing various custom logic.

This document describes at a high level how the Routes API can be used
inside a sidecar, the details of adaptation to each of the above contexts 
is not addressed in this document.

## Sidecar Routing

### Context

Typically, a sidecar is running in an environment in which it listens to 
all incoming TCP traffic on a single port, and recovers the original
destination address of each incoming TCP address for routing.  This is 
usually accomplished using `iptables` and an init container in the Pod.

### Environment

| Name                        | Description                                 |
|-----------------------------|---------------------------------------------|
| SIGNADOT_BASELINE_KIND      | the kind of a the baseline workload         |
| SIGNADOT_BASELINE_NAMESPACE | the namespace of the baseline workload      |
| SIGNADOT_BASELINE_NAME      | the name of the baseline workload           |
| SIGNADOT_ROUTESERVER        | the TCP address of the signadot routeserver |

### Go Client Example


Given an environment containing the above variables, one 
can initialize the routing as follows:

```go
import "github.com/signadot/routesapi/go-routesapi/workload"

var routing workload.Router
func init() {
    var (
         cfg *workload.Config
         err error
    )
    cfg, err = workload.EnvConfig()
    if err != nil {
        panic(err)
    }
    routing, err = workload.NewRouter(cfg)
    if err != nil {
        panic(err)
    }
}

```

Then, when routing a request with routing key `routingKey` received on a give port, one can
retrieve any destination sandbox as follows

```
routing.GetTarget(port, routingKey)
```

This will return an empty string when the routing key does not indicate any sandboxed 
workload.  

## Considerations

Sandbox routing is more volatile than typical cluster routing, and HTTP proxies
typically keep idle connections open for some time.  Sometimes, destinations
can change whilst an idle connection is still in service to a stale
destination.  Depending on the intended deployment of your sidecar routing, you
may want to consider addressing this actively.  Disabling keep alive in the
proxy or otherwise limiting the keep alive or providing 1-1 mapping from client
HTTP connections to upstream HTTP connections can help.


