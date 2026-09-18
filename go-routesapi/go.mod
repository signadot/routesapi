module github.com/signadot/routesapi/go-routesapi

// go-routesapi is imported by third parties, so the `go` directive below is the
// minimum Go version every importer is forced onto. Keep it at the floor the
// code and its dependencies actually need (currently set by golang.org/x/*),
// not at the latest Go release. Use the `toolchain` directive to choose the Go
// version used to develop and test this repo; importers ignore that line.
go 1.23.0

toolchain go1.27.1

require (
	google.golang.org/grpc v1.60.0
	google.golang.org/protobuf v1.34.2
	k8s.io/apimachinery v0.31.1
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8 // indirect
)
