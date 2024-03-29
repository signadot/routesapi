# Workload Routing Rules

This document describes how to use workload rules as found in the [routes
api](../routes.proto).

## Routing Context 

A WorkloadRoutingRule is intended to be applied inside of a Kubernetes workload, such
as a Deployment or Argo Rollout, in order to determine where to route requests
destined for the workload.  

Workloads run Pods, and all containers in a Pod share network context, such as
IP addresses and ports, so one container can intercept requests destined to another
and decide to send the request on the the other container, or send it elsewhere.
Typically, this is done with a sidecar, and WorkloadRoutingRules are applied in sidecars
in the Signadot DevMesh.

WorkloadRoutingRules can also be applied inside an application.  This can be useful
for [message queues](message-queues.md).

## Rule Contents

A WorkloadRoutingRule contains

1. A baseline workload `kind`, `namespace`, and `name`.
1. A routing key.
1. A sandbox identifier associated with the baseline workload, which identifies to which sandboxed
workload the destination TCP addresses belong.
1. A mapping from workload ports to a list destinations TCP addresses.

Given a routing key `rule.routing-key`, a baseline workload `rule.baseliine.{kind,namespace,name}`, a sandbox identifier
`rule.sandbox-id`, and a mapping `rule.portMap` of workload ports to destination TCP addresses, one can determine
how to route traffic destined to the baseline workload based on the routing key.

At a high level, a rule states

```
If the traffic is destined for the rule's baseline and 
  has a routing key which is the same as the rule routing key
then
  route to the destination Sandbox in the rule.
```

More specifically, at the TCP level, a rule states
```
If the traffic is destined for the rule's baseline and 
  has a routing key which is the same as the rule routing key and
  is destined on a port which has a destination address in the rule's mappings
then
  route to a destination address in the mapping with the same port
```

### Multiple Destinations

In short, one can route to any one destination address in a WorkloadRoutingRule.  Here, we detail 
how the destination addresses are defined.

For every workload port `workloadPort` and every destination address
`host:port` in `rule.portMap[workloadPort]` , there exists a baseline
Kubernetes Service whose selector selects `rule.baseline` which contains a
ServicePort whose port equals `port` and whose targetPort equals
`workloadPort`.  Moreover, `host` corresponds to the sandbox version of the
baseline Kubernetes service.


For example, given the following manifests

```yaml
apiVersion: v1
kind: Service
metadata:
  namespace: ns
  name: service-1
spec:
  type ClusterIP
  selector:
    app: example-app
  ports:
  - port: 80
    targetPort: 8080
  - port: 8000
    targetPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  namespace: ns
  name: service-2
spec:
  type ClusterIP
  selector:
    app: example-app
  ports:
  - port: 1080
    targetPort: 8080
---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: ns
  name: example-deploy
spec:
  template:
    metadata:
      labels:
        app: example-app
    spec:
      # ...
```

A workload rule may look as follows

```json
{
  "routing_key": "ltj9s8scupb86",
  "sandboxed_workload": {
    "sandbox_id": "ltj9s8scupb86",
    "baseline": {
      "kind": "Deployment",
      "namespace": "namespace",
      "name": "example-deploy"
    }
  },
  "port_rules": [
    {
      "workload_port": 8080,
      "destinations": [
        {
          "host": "sb1-service-1-ffc3f8f0.hotrod.svc",
          "port": 80
        },
        {
          "host": "sb1-service-1-ffc3f8f0.hotrod.svc",
          "port": 8000
        },
        {
          "host": "sb1-service-2-fad553c1e.hotrod.svc",
          "port": 1080
        }
      ]
    }
  ]
}
```

### Empty Destinations

A WorkloadRoutingRule may have an empty set of PortRules.  This occurs when there is no service
associated with the baseline workload.   Since there is no service, it does not participate
in sandbox routing.

### RouteGroups

One may also have a workload rule in which the sandbox ID is not the same as the routing key.
This occurs when the routing key comes from a RouteGroup which selects the sandbox identified
by the sandbox ID.


