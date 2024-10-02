// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sys/private/conf/conf.proto

package conf

import (
	apisix "github.com/jace996/multiapp/pkg/apisix"
	_ "github.com/jace996/multiapp/pkg/blob"
	conf "github.com/jace996/multiapp/pkg/conf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     *conf.Data      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Security *conf.Security  `protobuf:"bytes,3,opt,name=security,proto3" json:"security,omitempty"`
	Services *conf.Services  `protobuf:"bytes,4,opt,name=services,proto3" json:"services,omitempty"`
	Logging  *conf.Logging   `protobuf:"bytes,6,opt,name=logging,proto3" json:"logging,omitempty"`
	Tracing  *conf.Tracers   `protobuf:"bytes,7,opt,name=tracing,proto3" json:"tracing,omitempty"`
	App      *conf.AppConfig `protobuf:"bytes,8,opt,name=app,proto3" json:"app,omitempty"`
	Dev      *conf.Dev       `protobuf:"bytes,9,opt,name=dev,proto3" json:"dev,omitempty"`
	Sys      *SysConf        `protobuf:"bytes,10,opt,name=sys,proto3" json:"sys,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_private_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_sys_private_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_sys_private_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetData() *conf.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetSecurity() *conf.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Bootstrap) GetServices() *conf.Services {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Bootstrap) GetLogging() *conf.Logging {
	if x != nil {
		return x.Logging
	}
	return nil
}

func (x *Bootstrap) GetTracing() *conf.Tracers {
	if x != nil {
		return x.Tracing
	}
	return nil
}

func (x *Bootstrap) GetApp() *conf.AppConfig {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetDev() *conf.Dev {
	if x != nil {
		return x.Dev
	}
	return nil
}

func (x *Bootstrap) GetSys() *SysConf {
	if x != nil {
		return x.Sys
	}
	return nil
}

type SysConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apisix *apisix.Config `protobuf:"bytes,1,opt,name=apisix,proto3" json:"apisix,omitempty"`
}

func (x *SysConf) Reset() {
	*x = SysConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_private_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysConf) ProtoMessage() {}

func (x *SysConf) ProtoReflect() protoreflect.Message {
	mi := &file_sys_private_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysConf.ProtoReflect.Descriptor instead.
func (*SysConf) Descriptor() ([]byte, []int) {
	return file_sys_private_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *SysConf) GetApisix() *apisix.Config {
	if x != nil {
		return x.Apisix
	}
	return nil
}

var File_sys_private_conf_conf_proto protoreflect.FileDescriptor

var file_sys_private_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x79, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73,
	0x79, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x0f, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62, 0x6c,
	0x6f, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61,
	0x70, 0x69, 0x73, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbe, 0x02, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x1e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a,
	0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12,
	0x27, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x70, 0x70,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x1b, 0x0a, 0x03, 0x64,
	0x65, 0x76, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x44, 0x65, 0x76, 0x52, 0x03, 0x64, 0x65, 0x76, 0x12, 0x27, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x03, 0x73, 0x79,
	0x73, 0x22, 0x31, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x26, 0x0a, 0x06,
	0x61, 0x70, 0x69, 0x73, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x69, 0x78, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x73, 0x69, 0x78, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x73,
	0x79, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b,
	0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_private_conf_conf_proto_rawDescOnce sync.Once
	file_sys_private_conf_conf_proto_rawDescData = file_sys_private_conf_conf_proto_rawDesc
)

func file_sys_private_conf_conf_proto_rawDescGZIP() []byte {
	file_sys_private_conf_conf_proto_rawDescOnce.Do(func() {
		file_sys_private_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_private_conf_conf_proto_rawDescData)
	})
	return file_sys_private_conf_conf_proto_rawDescData
}

var file_sys_private_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sys_private_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),      // 0: sys.internal.Bootstrap
	(*SysConf)(nil),        // 1: sys.internal.SysConf
	(*conf.Data)(nil),      // 2: conf.Data
	(*conf.Security)(nil),  // 3: conf.Security
	(*conf.Services)(nil),  // 4: conf.Services
	(*conf.Logging)(nil),   // 5: conf.Logging
	(*conf.Tracers)(nil),   // 6: conf.Tracers
	(*conf.AppConfig)(nil), // 7: conf.AppConfig
	(*conf.Dev)(nil),       // 8: conf.Dev
	(*apisix.Config)(nil),  // 9: apisix.Config
}
var file_sys_private_conf_conf_proto_depIdxs = []int32{
	2, // 0: sys.internal.Bootstrap.data:type_name -> conf.Data
	3, // 1: sys.internal.Bootstrap.security:type_name -> conf.Security
	4, // 2: sys.internal.Bootstrap.services:type_name -> conf.Services
	5, // 3: sys.internal.Bootstrap.logging:type_name -> conf.Logging
	6, // 4: sys.internal.Bootstrap.tracing:type_name -> conf.Tracers
	7, // 5: sys.internal.Bootstrap.app:type_name -> conf.AppConfig
	8, // 6: sys.internal.Bootstrap.dev:type_name -> conf.Dev
	1, // 7: sys.internal.Bootstrap.sys:type_name -> sys.internal.SysConf
	9, // 8: sys.internal.SysConf.apisix:type_name -> apisix.Config
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_sys_private_conf_conf_proto_init() }
func file_sys_private_conf_conf_proto_init() {
	if File_sys_private_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_private_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_sys_private_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysConf); i {
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
			RawDescriptor: file_sys_private_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sys_private_conf_conf_proto_goTypes,
		DependencyIndexes: file_sys_private_conf_conf_proto_depIdxs,
		MessageInfos:      file_sys_private_conf_conf_proto_msgTypes,
	}.Build()
	File_sys_private_conf_conf_proto = out.File
	file_sys_private_conf_conf_proto_rawDesc = nil
	file_sys_private_conf_conf_proto_goTypes = nil
	file_sys_private_conf_conf_proto_depIdxs = nil
}
