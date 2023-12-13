# Workload Rules

This document describes how to use workload rules as found in the [routes
api](../routes.proto).

## Routing Context 

A WorkloadRule is intended to be applied inside of a Kubernetes workload, such
as a Deployment or Argo Rollout, in order to determine where to route requests
destined for the workload.  

Workloads run Pods, and all containers in a Pod share network context, such as
IP addresses and ports, so one container can intercept requests destined to another
and decide to send the request on the the other container, or send it elsewhere.
Typically, this is done with a sidecar, and WorkloadRules are applied in sidecars
in the Signadot DevMesh.

WorkloadRules can also be applied inside an application.  This can be useful
for [message queues](message-queues.md).

## Rule Contents

A WorkloadRule contains

1. A routing key
1. A baseline workload `kind`, `namespace`, and `name`
1. A mapping from workload ports to a list destinations TCP addresses.
1. A sandbox identifier associated with the baseline workload, which identifies to which sandboxed
workload the destination TCP addresses 

Given a routing key `rule.routing-key`, a baseline workload `rule.baseliine.{kind,namespace,name}`, a sandbox identifier
`rule.sandbox-id`, and a mapping `rule.portMap` of workload ports to destination TCP addresses, one can determine
how to route a request.

Given a request `req` with routing key `req.routing-key` arriving on port `req.port` of baseline workload 
`req.baseline`, a workload rule states

```
If `req.baseline` == `rule.baseline` and 
    `req.routing-key` == `rule.routing-key` and 
    some TCP addresses `addrs` are in `rule.portMap[req.port]` 
then route to any one address `addrs`, otherwise pass the request to the baseline.
```


### Multiple Destinations

In short, one can route to any one destination address in a WorkloadRule.  Here, we detail 
how the destination addresses are defined.

For every workload port `workloadPort` and every destination address `host:port` in `rule.portMap[workloadPort]` , there exists a baseline Kubernetes Service 
whose selector selects `rule.baseline` which contains a ServicePort whose port equals `port` and whose targetPort equals `workloadPort`.  Moreover, `host` corresponds to the sandbox version of the baseline Kubernetes service.


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

One may also have a workload rule in which the sandbox ID is not the same as the routing key.
This occurs when the routing key comes from a RouteGroup which selects the sandbox identified
by the sandbox ID.


