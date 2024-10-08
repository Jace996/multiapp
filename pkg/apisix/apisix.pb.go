// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: apisix/apisix.proto

package apisix

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host   string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port   uint64 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Weight int64  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apisix_apisix_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_apisix_apisix_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_apisix_apisix_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Node) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Node) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type Upstream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes  []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Type   string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Scheme string  `protobuf:"bytes,3,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Name   string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Upstream) Reset() {
	*x = Upstream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apisix_apisix_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upstream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upstream) ProtoMessage() {}

func (x *Upstream) ProtoReflect() protoreflect.Message {
	mi := &file_apisix_apisix_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upstream.ProtoReflect.Descriptor instead.
func (*Upstream) Descriptor() ([]byte, []int) {
	return file_apisix_apisix_proto_rawDescGZIP(), []int{1}
}

func (x *Upstream) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Upstream) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Upstream) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Upstream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string    `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ApiKey   string    `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	Modules  []*Module `protobuf:"bytes,3,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apisix_apisix_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_apisix_apisix_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_apisix_apisix_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Config) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *Config) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// https://apisix.apache.org/docs/apisix/admin-api/#route
	Routes map[string]*structpb.Struct `protobuf:"bytes,10,rep,name=routes,proto3" json:"routes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// https://apisix.apache.org/docs/apisix/admin-api/#global-rule
	GlobalRules map[string]*structpb.Struct `protobuf:"bytes,11,rep,name=global_rules,json=globalRules,proto3" json:"global_rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// https://apisix.apache.org/docs/apisix/admin-api/#upstream
	Upstreams map[string]*structpb.Struct `protobuf:"bytes,12,rep,name=upstreams,proto3" json:"upstreams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// https://apisix.apache.org/docs/apisix/admin-api/#stream-route
	StreamRoutes map[string]*structpb.Struct `protobuf:"bytes,13,rep,name=stream_routes,json=streamRoutes,proto3" json:"stream_routes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apisix_apisix_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_apisix_apisix_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_apisix_apisix_proto_rawDescGZIP(), []int{3}
}

func (x *Module) GetRoutes() map[string]*structpb.Struct {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *Module) GetGlobalRules() map[string]*structpb.Struct {
	if x != nil {
		return x.GlobalRules
	}
	return nil
}

func (x *Module) GetUpstreams() map[string]*structpb.Struct {
	if x != nil {
		return x.Upstreams
	}
	return nil
}

func (x *Module) GetStreamRoutes() map[string]*structpb.Struct {
	if x != nil {
		return x.StreamRoutes
	}
	return nil
}

var File_apisix_apisix_proto protoreflect.FileDescriptor

var file_apisix_apisix_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x04, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x6e, 0x0a, 0x08, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x22, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xe2, 0x04, 0x0a,
	0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x3b, 0x0a, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2e, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x45, 0x0a, 0x0d,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x1a, 0x52, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x57, 0x0a, 0x10, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x55, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x58, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x61, 0x63, 0x65, 0x39, 0x39, 0x36, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x70, 0x70,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x3b, 0x61, 0x70, 0x69, 0x73,
	0x69, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apisix_apisix_proto_rawDescOnce sync.Once
	file_apisix_apisix_proto_rawDescData = file_apisix_apisix_proto_rawDesc
)

func file_apisix_apisix_proto_rawDescGZIP() []byte {
	file_apisix_apisix_proto_rawDescOnce.Do(func() {
		file_apisix_apisix_proto_rawDescData = protoimpl.X.CompressGZIP(file_apisix_apisix_proto_rawDescData)
	})
	return file_apisix_apisix_proto_rawDescData
}

var file_apisix_apisix_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apisix_apisix_proto_goTypes = []any{
	(*Node)(nil),            // 0: apisix.Node
	(*Upstream)(nil),        // 1: apisix.Upstream
	(*Config)(nil),          // 2: apisix.Config
	(*Module)(nil),          // 3: apisix.Module
	nil,                     // 4: apisix.Module.RoutesEntry
	nil,                     // 5: apisix.Module.GlobalRulesEntry
	nil,                     // 6: apisix.Module.UpstreamsEntry
	nil,                     // 7: apisix.Module.StreamRoutesEntry
	(*structpb.Struct)(nil), // 8: google.protobuf.Struct
}
var file_apisix_apisix_proto_depIdxs = []int32{
	0,  // 0: apisix.Upstream.nodes:type_name -> apisix.Node
	3,  // 1: apisix.Config.modules:type_name -> apisix.Module
	4,  // 2: apisix.Module.routes:type_name -> apisix.Module.RoutesEntry
	5,  // 3: apisix.Module.global_rules:type_name -> apisix.Module.GlobalRulesEntry
	6,  // 4: apisix.Module.upstreams:type_name -> apisix.Module.UpstreamsEntry
	7,  // 5: apisix.Module.stream_routes:type_name -> apisix.Module.StreamRoutesEntry
	8,  // 6: apisix.Module.RoutesEntry.value:type_name -> google.protobuf.Struct
	8,  // 7: apisix.Module.GlobalRulesEntry.value:type_name -> google.protobuf.Struct
	8,  // 8: apisix.Module.UpstreamsEntry.value:type_name -> google.protobuf.Struct
	8,  // 9: apisix.Module.StreamRoutesEntry.value:type_name -> google.protobuf.Struct
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_apisix_apisix_proto_init() }
func file_apisix_apisix_proto_init() {
	if File_apisix_apisix_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apisix_apisix_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Node); i {
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
		file_apisix_apisix_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Upstream); i {
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
		file_apisix_apisix_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
		file_apisix_apisix_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Module); i {
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
			RawDescriptor: file_apisix_apisix_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apisix_apisix_proto_goTypes,
		DependencyIndexes: file_apisix_apisix_proto_depIdxs,
		MessageInfos:      file_apisix_apisix_proto_msgTypes,
	}.Build()
	File_apisix_apisix_proto = out.File
	file_apisix_apisix_proto_rawDesc = nil
	file_apisix_apisix_proto_goTypes = nil
	file_apisix_apisix_proto_depIdxs = nil
}
