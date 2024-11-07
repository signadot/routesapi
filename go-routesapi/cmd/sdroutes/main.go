package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/signadot/routesapi/go-routesapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	addr        string
	namespace   string
	name        string
	kind        string
	routingKey  string
	sandboxName string
	watch       bool
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
	flag.StringVar(&sandboxName, "sandbox-name", "", "routes routing to sandbox with given name")
	flag.BoolVar(&watch, "watch", false, "whether to watch")
	flag.Parse()
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    5 * time.Second,
			Timeout: 15 * time.Second,
		}))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to %s\n", addr)
	routesClient := routesapi.NewRoutesClient(conn)
	req := &routesapi.WorkloadRoutingRulesRequest{
		BaselineWorkload: &routesapi.BaselineWorkload{
			Namespace: namespace,
			Name:      name,
			Kind:      kind,
		},
		RoutingKey: routingKey,
		DestinationSandbox: &routesapi.DestinationSandbox{
			Name: sandboxName,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if !watch {
		resp, err := routesClient.GetWorkloadRoutingRules(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		s := pbjOpts.Format(resp)
		fmt.Println(s)
		return
	}
	w, err := routesClient.WatchWorkloadRoutingRules(context.Background(), req)
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
