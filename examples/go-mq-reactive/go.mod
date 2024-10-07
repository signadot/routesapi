module github.com/signadot/routesapi/examples/go-mq-reactive

go 1.22

toolchain go1.22.5

replace github.com/signadot/routesapi/go-routesapi => ../../go-routesapi

require (
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
	github.com/signadot/routesapi/go-routesapi v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.60.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
