module github.com/signadot/routesapi/examples/go-mq-reactive

go 1.27.1

replace github.com/signadot/routesapi/go-routesapi => ../../go-routesapi

require (
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
	github.com/signadot/routesapi/go-routesapi v0.0.0-20250417101624-9ce649411ae0
	google.golang.org/grpc v1.83.2
)

require (
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
