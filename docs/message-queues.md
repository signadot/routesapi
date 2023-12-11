# Message Queue Consumer Routing

One interesting way of placing message queues in sandboxes is to configure
multiple consumers of the message queue to accept or reject messages based on
routing keys, which are to be injected into message queue messages as metadata
or headers.

This configuration allows one to work on consumers of message queues in a
sandbox without needing to include the corresponding producers in the the
sandbox.

Other mechanisms for sandboxing message queues exist, such as configuring an
application to use sandbox-specific topics or spinning up a new message queue
broker per-sandbox.  However, these are heavier weight and not as flexible.

This document describes how to use the Routes API to configure message queue
consumers to accept or reject message queue messages.

## Message Queue Consumers and Sandboxes

Many message queues, such as Kafka, support having multiple consumers of
messages.  In this "multi-fanout" scenario, each consumer consumes all
messages.  Additionally, we assume that the relevant messages are decorated
with a routing key on the producer side.

When working with sandboxes involving message queues, one can dispatch messages
conditionally to one consumer or another by:

1. Sending all messages to all consumers, such as Kafka consumer groups.
1. Configuring each consumer to process or not each message as a function
of the routing key of the message.


## Configuring Consumers

As each consumer will behave different depending on whether it is running in a
baseline or sandboxed workload, it is essential that each consumer be configured
appropriately.

Signadot sandboxes provide a [downward api](https://www.signadot.com/docs/reference/sandboxes/sandbox-downward-api) via environment variables which can be used by a workload to determine if it is running in
the baseline or in a sandbox.

Additionally, the Go Routes API works with the
following environment variables

| Name                        | Description                                 |
|-----------------------------|---------------------------------------------|
| SIGNADOT_BASELINE_KIND      | the kind of a the baseline workload         |
| SIGNADOT_BASELINE_NAMESPACE | the namespace of the baseline workload      |
| SIGNADOT_BASELINE_NAME      | the name of the baseline workload           |
| SIGNADOT_ROUTESERVER        | the TCP address of the signadot routeserver |
| ----                        | ----                                        |

The Go Routes API SDK provides a simple function to get access to the routing
rules for a baseline workload from the environment:

```go
import "github.com/signadot/routesapi/watched"
routing, err := watched.BaselineFromEnv()
```

## Using `routesapi` Go SDK

In the baseline:
```go
import "github.com/signadot/routesapi/watched"

func ShouldProcess(routing watched.BaselineWatched, routingKey string) bool {
    return routing.Get(routingKey) == nil
}
```

In a sandboxed workload:

```go
import "github.com/signadot/routesapi/watched"

func ShouldProcess(routing watched.BaselineWatched, routingKey string) bool {
    return routing.RoutesTo(routingKey, os.Getenv("SIGNADOT_SANDBOX_ROUTING_KEY"))
}
```

It can be useful to combine the code so that the code can be deployed
unconditionally.  This can be accomplished as follows:

```go
func ShouldProcess(routing watched.BaselineWatched, routingKey string) bool {
    if sbID := os.Getenv("SIGNADOT_SANDBOX_ROUTING_KEY"); sbID != "" {
       return routing.RoutesTo(routingKey, sbID)
    }
    return routing.Get(routingKey) == nil
}
```
