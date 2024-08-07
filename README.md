# Signadot Routes API

## Overview

The Signadot Routes API provides access to routing rules pertinent to Signadot
Sandboxes on a cluster with the Signadot Operator (>= v0.15.0) installed.
Effective use of the Routes API requires an understanding of [Sandbox
Routing](docs/sandbox-routing.md).

## Routeserver

The Signadot Operator packages a routeserver deployment and service in the
`signadot` namespace, running a GRPC service at
`routeserver.signadot.svc:7777` and a corresponding HTTP service at 
`routeserver.signadot.svc:7778`. 

## Contents

This repository hosts 

- The [GRPC service definition](routes.proto).
- The [HTTP openapi spec](routes-openapi.yaml) and corresponding [docs](https://signadot.github.io/routesapi/http-api).
- A [Go client](go-routesapi/README.md) with supporting libraries and a command for querying the gRPC
  and HTTP services.
- [Example applications](examples/) with message queues.
- Documentation
  * [Sandbox Routing](docs/sandbox-routing.md)
  * [Workload Rules](docs/workload-rules.md)









