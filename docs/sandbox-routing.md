# Sandbox Routing

## Overview

Signadot is a testing infrastructure product which permits creating test
variants of in-cluster workloads, called _sandboxed workloads_, inside a
sandbox.  In a sandbox, the associated sandboxed workloads interact with the
rest of the cluster, called _the baseline_, seamlessly. Requests can be
associated with sandboxes.  When a request is associated with a sandbox, it is
routed to the appropriate _sandboxed workloads_.  With context propagation,
requests associated with sandboxes can be sent to any service in the cluster.

The purpose of the Routes API is to provide access to the routing rules which
coordinate requests for sandboxes in a cluster.  As sandboxes are intended for
testing and development, they change frequently and so such routing rules are
typically much more dynamic and volatile than baseline routing configuraion.

## Origin and Destination Routing

Most service meshes implement routing rules at the _origin_ of a request.
Doing so allows customizing requests based on DNS and customizing requests to
entities outside the cluster.  But it also has the drawback of requiring
instrumentation of ingresses and gateways, which can be problematic.

An alternative [^1] is to process routing rules at the _destination_ of a request,
which has the the disadvantage of not working with DNS or entities outside a
cluster, as one can not know by what host or service a request was sent or may
not control the routing of an entity outside the cluster.


Destination based routing, however, has 2 advantages:

* It provides separation of concerns: any client, including ingresses and gateways, will route correctly without the need to take into account customizing the ingresses.
* It simplifies in-app routing a great deal, which can be useful for message queues.

The Signadot Operator can plug into service meshes which route at the origin of
requests, but it also supplies a DevMesh which implements destination based
routing.

The Signadot Routes API is used by the DevMesh sidecars, while it is not used
for service mesh integration.  

As a result, the Signadot Routes API only provides routing rules for consumption
at the destination of a request.

Future versions may provide routing rules for routing at the origin of requests.


## Notes

[^1] Of course, ideally, one would expect routing to occur neither at the source
nor at the destination but rather simply in-flight.  This option does not
exist on Kubernetes.
