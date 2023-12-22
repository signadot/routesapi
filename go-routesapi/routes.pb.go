// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: routes.proto

package routesapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WatchOp int32

const (
	WatchOp_ADD     WatchOp = 0
	WatchOp_REMOVE  WatchOp = 1
	WatchOp_REPLACE WatchOp = 2
	// In any WatchWorkloadRoutes rpc call, the returned stream will send at most
	// 1 SYNCED WatchOp, indicating the client has all the information about
	// Sandboxes and RouteGroups available in the cluster. Prior to sending a
	// SYNCED WatchOp, all WatchOps are ADDs.
	WatchOp_SYNCED WatchOp = 3
)

// Enum value maps for WatchOp.
var (
	WatchOp_name = map[int32]string{
		0: "ADD",
		1: "REMOVE",
		2: "REPLACE",
		3: "SYNCED",
	}
	WatchOp_value = map[string]int32{
		"ADD":     0,
		"REMOVE":  1,
		"REPLACE": 2,
		"SYNCED":  3,
	}
)

func (x WatchOp) Enum() *WatchOp {
	p := new(WatchOp)
	*p = x
	return p
}

func (x WatchOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatchOp) Descriptor() protoreflect.EnumDescriptor {
	return file_routes_proto_enumTypes[0].Descriptor()
}

func (WatchOp) Type() protoreflect.EnumType {
	return &file_routes_proto_enumTypes[0]
}

func (x WatchOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatchOp.Descriptor instead.
func (WatchOp) EnumDescriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{0}
}

// A WorkloadRoute defines for a given baseline and a routing key, a single
// `destinationSandbox` and `mappings`. The mappings map each port of the
// baseline workload with corresponding TCP addresses belonging to the
// `destinationSandbox` where traffic is routed instead.
type WorkloadRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutingKey string `protobuf:"bytes,1,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty"`
	// baseline indicates the corresponding baseline workload.
	Baseline *BaselineWorkload `protobuf:"bytes,2,opt,name=baseline,proto3" json:"baseline,omitempty"`
	// destination_sandbox indicates the sandbox associated with the destination
	// sandboxed workloads.
	DestinationSandbox *DestinationSandbox `protobuf:"bytes,3,opt,name=destination_sandbox,json=destinationSandbox,proto3" json:"destination_sandbox,omitempty"`
	// mappings represents a mapping from a port on the workload to a set of
	// destinations.
	Mappings []*WorkloadPortMapping `protobuf:"bytes,4,rep,name=mappings,proto3" json:"mappings,omitempty"`
}

func (x *WorkloadRoute) Reset() {
	*x = WorkloadRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRoute) ProtoMessage() {}

func (x *WorkloadRoute) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRoute.ProtoReflect.Descriptor instead.
func (*WorkloadRoute) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{0}
}

func (x *WorkloadRoute) GetRoutingKey() string {
	if x != nil {
		return x.RoutingKey
	}
	return ""
}

func (x *WorkloadRoute) GetBaseline() *BaselineWorkload {
	if x != nil {
		return x.Baseline
	}
	return nil
}

func (x *WorkloadRoute) GetDestinationSandbox() *DestinationSandbox {
	if x != nil {
		return x.DestinationSandbox
	}
	return nil
}

func (x *WorkloadRoute) GetMappings() []*WorkloadPortMapping {
	if x != nil {
		return x.Mappings
	}
	return nil
}

// A DestinationSandbox represents a sandbox that will receive traffic intended
// for a baseline workload in the presence of a routing key.
type DestinationSandbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sandbox name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DestinationSandbox) Reset() {
	*x = DestinationSandbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationSandbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationSandbox) ProtoMessage() {}

func (x *DestinationSandbox) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationSandbox.ProtoReflect.Descriptor instead.
func (*DestinationSandbox) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{1}
}

func (x *DestinationSandbox) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A BaselineWorkload identifies a given baseline workload. In the context of a
// WorkloadRoutesRequest, all the fields are optional. In the context of a
// response from the server, all the fields are filled in.
type BaselineWorkload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind      string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BaselineWorkload) Reset() {
	*x = BaselineWorkload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaselineWorkload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaselineWorkload) ProtoMessage() {}

func (x *BaselineWorkload) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaselineWorkload.ProtoReflect.Descriptor instead.
func (*BaselineWorkload) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{2}
}

func (x *BaselineWorkload) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *BaselineWorkload) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *BaselineWorkload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A WorkloadPortMapping provides a mapping from a port on the workload to a set
// of destinations. Each destination in the response corresponds to a sandbox
// service matching the sandboxed workload. As a result, any of the destinations
// can be used.
type WorkloadPortMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkloadPort uint32      `protobuf:"varint,1,opt,name=workload_port,json=workloadPort,proto3" json:"workload_port,omitempty"`
	Destinations []*Location `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *WorkloadPortMapping) Reset() {
	*x = WorkloadPortMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadPortMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadPortMapping) ProtoMessage() {}

func (x *WorkloadPortMapping) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadPortMapping.ProtoReflect.Descriptor instead.
func (*WorkloadPortMapping) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{3}
}

func (x *WorkloadPortMapping) GetWorkloadPort() uint32 {
	if x != nil {
		return x.WorkloadPort
	}
	return 0
}

func (x *WorkloadPortMapping) GetDestinations() []*Location {
	if x != nil {
		return x.Destinations
	}
	return nil
}

// Location gives a TCP address as a host, port pair.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{4}
}

func (x *Location) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Location) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// WorkloadRoutesRequest is a request for a set of WorkloadRoutes, which give
// information about how to route requests when they are intercepted at a given
// workload. Each field is optional and constrains the the set of WorkloadRoutes
// returned accordingly.
type WorkloadRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// baseline_workload specifies the kind, namespace, and name of the baseline
	// workload to which requests are directed. Each field is optional.
	BaselineWorkload *BaselineWorkload `protobuf:"bytes,1,opt,name=baseline_workload,json=baselineWorkload,proto3" json:"baseline_workload,omitempty"`
	// routing_key specifies the routing key associated with the request.
	RoutingKey string `protobuf:"bytes,2,opt,name=routing_key,json=routingKey,proto3" json:"routing_key,omitempty"`
	// destination_sandbox specifies the sandbox associated with the destination
	// sandboxed workloads.
	DestinationSandbox *DestinationSandbox `protobuf:"bytes,3,opt,name=destination_sandbox,json=destinationSandbox,proto3" json:"destination_sandbox,omitempty"`
}

func (x *WorkloadRoutesRequest) Reset() {
	*x = WorkloadRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRoutesRequest) ProtoMessage() {}

func (x *WorkloadRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRoutesRequest.ProtoReflect.Descriptor instead.
func (*WorkloadRoutesRequest) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{5}
}

func (x *WorkloadRoutesRequest) GetBaselineWorkload() *BaselineWorkload {
	if x != nil {
		return x.BaselineWorkload
	}
	return nil
}

func (x *WorkloadRoutesRequest) GetRoutingKey() string {
	if x != nil {
		return x.RoutingKey
	}
	return ""
}

func (x *WorkloadRoutesRequest) GetDestinationSandbox() *DestinationSandbox {
	if x != nil {
		return x.DestinationSandbox
	}
	return nil
}

// a GetWorkloadRoutesResponse gives the set of WorkloadRoutes which match a
// given WorkloadRoutesRequest.
type GetWorkloadRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*WorkloadRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *GetWorkloadRoutesResponse) Reset() {
	*x = GetWorkloadRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkloadRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkloadRoutesResponse) ProtoMessage() {}

func (x *GetWorkloadRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkloadRoutesResponse.ProtoReflect.Descriptor instead.
func (*GetWorkloadRoutesResponse) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{6}
}

func (x *GetWorkloadRoutesResponse) GetRoutes() []*WorkloadRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

// WorkloadRouteOp describes a diff operation against a set of workload routes:
// adding, removing, and replacing WorkloadRoutes are possible. Additionally,
// there is a SYNCED operation to indicate when the client has received all
// relevant WorkloadRoutes.
type WorkloadRouteOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    WatchOp        `protobuf:"varint,1,opt,name=op,proto3,enum=routes.WatchOp" json:"op,omitempty"`
	Route *WorkloadRoute `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *WorkloadRouteOp) Reset() {
	*x = WorkloadRouteOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRouteOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRouteOp) ProtoMessage() {}

func (x *WorkloadRouteOp) ProtoReflect() protoreflect.Message {
	mi := &file_routes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRouteOp.ProtoReflect.Descriptor instead.
func (*WorkloadRouteOp) Descriptor() ([]byte, []int) {
	return file_routes_proto_rawDescGZIP(), []int{7}
}

func (x *WorkloadRouteOp) GetOp() WatchOp {
	if x != nil {
		return x.Op
	}
	return WatchOp_ADD
}

func (x *WorkloadRouteOp) GetRoute() *WorkloadRoute {
	if x != nil {
		return x.Route
	}
	return nil
}

var File_routes_proto protoreflect.FileDescriptor

var file_routes_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xec, 0x01, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x4b, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x37, 0x0a, 0x08,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x58, 0x0a, 0x10, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x13, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x32, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0xcc, 0x01, 0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x62, 0x61, 0x73,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x10,
	0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65,
	0x79, 0x12, 0x4b, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x22, 0x4a,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x0f, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4f, 0x70, 0x12, 0x1f, 0x0a,
	0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x2b,
	0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2a, 0x37, 0x0a, 0x07, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x4e, 0x43,
	0x45, 0x44, 0x10, 0x03, 0x32, 0xb4, 0x01, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x57, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x13, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x4f, 0x70, 0x22, 0x00, 0x30, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x64,
	0x6f, 0x74, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routes_proto_rawDescOnce sync.Once
	file_routes_proto_rawDescData = file_routes_proto_rawDesc
)

func file_routes_proto_rawDescGZIP() []byte {
	file_routes_proto_rawDescOnce.Do(func() {
		file_routes_proto_rawDescData = protoimpl.X.CompressGZIP(file_routes_proto_rawDescData)
	})
	return file_routes_proto_rawDescData
}

var file_routes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_routes_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_routes_proto_goTypes = []interface{}{
	(WatchOp)(0),                      // 0: routes.WatchOp
	(*WorkloadRoute)(nil),             // 1: routes.WorkloadRoute
	(*DestinationSandbox)(nil),        // 2: routes.DestinationSandbox
	(*BaselineWorkload)(nil),          // 3: routes.BaselineWorkload
	(*WorkloadPortMapping)(nil),       // 4: routes.WorkloadPortMapping
	(*Location)(nil),                  // 5: routes.Location
	(*WorkloadRoutesRequest)(nil),     // 6: routes.WorkloadRoutesRequest
	(*GetWorkloadRoutesResponse)(nil), // 7: routes.GetWorkloadRoutesResponse
	(*WorkloadRouteOp)(nil),           // 8: routes.WorkloadRouteOp
}
var file_routes_proto_depIdxs = []int32{
	3,  // 0: routes.WorkloadRoute.baseline:type_name -> routes.BaselineWorkload
	2,  // 1: routes.WorkloadRoute.destination_sandbox:type_name -> routes.DestinationSandbox
	4,  // 2: routes.WorkloadRoute.mappings:type_name -> routes.WorkloadPortMapping
	5,  // 3: routes.WorkloadPortMapping.destinations:type_name -> routes.Location
	3,  // 4: routes.WorkloadRoutesRequest.baseline_workload:type_name -> routes.BaselineWorkload
	2,  // 5: routes.WorkloadRoutesRequest.destination_sandbox:type_name -> routes.DestinationSandbox
	1,  // 6: routes.GetWorkloadRoutesResponse.routes:type_name -> routes.WorkloadRoute
	0,  // 7: routes.WorkloadRouteOp.op:type_name -> routes.WatchOp
	1,  // 8: routes.WorkloadRouteOp.route:type_name -> routes.WorkloadRoute
	6,  // 9: routes.Routes.GetWorkloadRoutes:input_type -> routes.WorkloadRoutesRequest
	6,  // 10: routes.Routes.WatchWorkloadRoutes:input_type -> routes.WorkloadRoutesRequest
	7,  // 11: routes.Routes.GetWorkloadRoutes:output_type -> routes.GetWorkloadRoutesResponse
	8,  // 12: routes.Routes.WatchWorkloadRoutes:output_type -> routes.WorkloadRouteOp
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_routes_proto_init() }
func file_routes_proto_init() {
	if File_routes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRoute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationSandbox); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaselineWorkload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadPortMapping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRoutesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkloadRoutesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_routes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRouteOp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_routes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routes_proto_goTypes,
		DependencyIndexes: file_routes_proto_depIdxs,
		EnumInfos:         file_routes_proto_enumTypes,
		MessageInfos:      file_routes_proto_msgTypes,
	}.Build()
	File_routes_proto = out.File
	file_routes_proto_rawDesc = nil
	file_routes_proto_goTypes = nil
	file_routes_proto_depIdxs = nil
}
