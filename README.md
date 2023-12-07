# Signadot Routes API

## Overview

The Signadot Routes API provides access to routing rules pertinent to Signadot
Sandboxes on a cluster with the Signadot Operator (>= v0.15.0) installed.
Effective use of the Routes API requires an understanding of [Sandbox
Routing](docs/sandbox-routing.md).

## Routeserver

The Signadot Operator packages a routeserver deployment and service in the
`signadot` namespace, running a GRPC service at
`routeserver.signadot.svc:7777`.  This repository provides client support for
using the routeserver.

## Contents

This repository hosts 

- The [GRPC service definition](routes.proto).
- A generated Go client.
- Libraries for destination workload routing.
- A command for querying the route server.
- Documentation









