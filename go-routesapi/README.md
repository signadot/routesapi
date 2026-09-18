# Signadot Routes API Go Client

## Overview

This directory contains a Go client of the [Signadot Routes API](../README.md).

## Requirements

Go 1.23 or newer. The module's `go` directive is kept at that floor on purpose,
so importing this client does not force your project onto the latest Go release.

## Contents

- A generated Go client.
- Libraries for destination workload routing.
- A command for querying the route server.
- Docs
  * [Sidecar Routing](docs/sidecar-routing.md)
  * [Message queues](docs/message-queues.md)
