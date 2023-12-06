package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/signadot/routesapi"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	addr       string
	namespace  string
	name       string
	kind       string
	routingKey string
	sandboxID  string
	watch      bool
)

var (
	pbjOpts = protojson.MarshalOptions{
		Multiline:       true,
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
)

func main() {
	flag.StringVar(&addr, "addr", "localhost:7777", "routeserver address to dial")
	flag.StringVar(&namespace, "namespace", "", "routes namespaces")
	flag.StringVar(&name, "name", "", "routes baseline workload name")
	flag.StringVar(&kind, "kind", "", "routes baseline workload kind")
	flag.StringVar(&routingKey, "routing-key", "", "routes routing key")
	flag.StringVar(&sandboxID, "sandbox-id", "", "routes routing to sandbox with given id")
	flag.BoolVar(&watch, "watch", false, "whether to watch")
	flag.Parse()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to %s\n", addr)
	routesClient := routesapi.NewRoutesClient(conn)
	req := &routesapi.WorkloadRoutesRequest{
		BaselineWorkload: &routesapi.BaselineWorkload{
			Namespace: namespace,
			Name:      name,
			Kind:      kind,
		},
		RoutingKey: routingKey,
		SandboxID:  sandboxID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if !watch {
		resp, err := routesClient.GetWorkloadRoutes(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		s := pbjOpts.Format(resp)
		fmt.Println(s)
		return
	}
	w, err := routesClient.WatchWorkloadRoutes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := w.Recv()
		if err != nil {
			log.Fatal(err)
		}
		s := pbjOpts.Format(resp)
		fmt.Println(s)
	}
}
