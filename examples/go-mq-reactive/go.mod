module github.com/signadot/routesapi/examples/go-mq-reactive

go 1.22
toolchain go1.23.6

replace github.com/signadot/routesapi/go-routesapi => ../../go-routesapi

require (
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
	github.com/signadot/routesapi/go-routesapi v0.0.0-20241107102336-873361c4ca81
	google.golang.org/grpc v1.60.0
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
