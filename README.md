# Signadot Routes API

## Overview

The Signadot Routes API provides access to routing rules pertinent to Signadot
Sandboxes on a cluster with the Signadot Operator (>= v0.14.2) installed.
Effective use of the Routes API requires an understanding of [Sandbox
Routing](docs/sandbox-routing.md).

NOTE: this is currently in preview mode, there is no Signadot Operator >= v0.14.2
at the time of making this repository public.  It will be coming soon!

## Routeserver

The Signadot Operator packages a routeserver deployment and service in the
`signadot` namespace, running a GRPC service at
`routeserver.signadot.svc:7777`.  This repository provides the service
definition, examples, and [a Go client](go-routesapi/README.md).

## Contents

This repository hosts 

- The [GRPC service definition](routes.proto).
- A generated Go client with supporting libraries and a command for querying the routeserver.
- Documentation
  * [Sandbox Routing](docs/sandbox-routing.md)
  * [Workload Rules](docs/workload-rules.md)
  * [Sidecar Routing](docs/sidecar-routing.md)
  * [Message queues](docs/message-queues.md)









