// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: redis/redis.proto

package redis

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addrs           []string                `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	ReadTimeout     *durationpb.Duration    `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3,oneof" json:"read_timeout,omitempty"`
	WriteTimeout    *durationpb.Duration    `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3,oneof" json:"write_timeout,omitempty"`
	Username        *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password        *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Db              *wrapperspb.Int32Value  `protobuf:"bytes,7,opt,name=db,proto3,oneof" json:"db,omitempty"`
	MaxRetries      *wrapperspb.Int32Value  `protobuf:"bytes,8,opt,name=max_retries,json=maxRetries,proto3,oneof" json:"max_retries,omitempty"`
	MinRetryBackoff *durationpb.Duration    `protobuf:"bytes,9,opt,name=min_retry_backoff,json=minRetryBackoff,proto3,oneof" json:"min_retry_backoff,omitempty"`
	MaxRetryBackoff *durationpb.Duration    `protobuf:"bytes,10,opt,name=max_retry_backoff,json=maxRetryBackoff,proto3,oneof" json:"max_retry_backoff,omitempty"`
	DialTimeout     *durationpb.Duration    `protobuf:"bytes,11,opt,name=dial_timeout,json=dialTimeout,proto3,oneof" json:"dial_timeout,omitempty"`
	MasterName      *string                 `protobuf:"bytes,12,opt,name=master_name,json=masterName,proto3,oneof" json:"master_name,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_redis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_redis_redis_proto_msgTypes[0]
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
	return file_redis_redis_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

func (x *Config) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Config) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Config) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *Config) GetPassword() *wrapperspb.StringValue {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *Config) GetDb() *wrapperspb.Int32Value {
	if x != nil {
		return x.Db
	}
	return nil
}

func (x *Config) GetMaxRetries() *wrapperspb.Int32Value {
	if x != nil {
		return x.MaxRetries
	}
	return nil
}

func (x *Config) GetMinRetryBackoff() *durationpb.Duration {
	if x != nil {
		return x.MinRetryBackoff
	}
	return nil
}

func (x *Config) GetMaxRetryBackoff() *durationpb.Duration {
	if x != nil {
		return x.MaxRetryBackoff
	}
	return nil
}

func (x *Config) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Config) GetMasterName() string {
	if x != nil && x.MasterName != nil {
		return *x.MasterName
	}
	return ""
}

var File_redis_redis_proto protoreflect.FileDescriptor

var file_redis_redis_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x06, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b,
	0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x43,
	0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x01, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x02, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x03, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x30, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x04, 0x52, 0x02, 0x64, 0x62,
	0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x11, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x06, 0x52, 0x0f,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x88,
	0x01, 0x01, 0x12, 0x4a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x07, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x88, 0x01, 0x01, 0x12, 0x41,
	0x0a, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x08, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x64, 0x62, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66,
	0x66, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x61, 0x6c,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b,
	0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x3b, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redis_redis_proto_rawDescOnce sync.Once
	file_redis_redis_proto_rawDescData = file_redis_redis_proto_rawDesc
)

func file_redis_redis_proto_rawDescGZIP() []byte {
	file_redis_redis_proto_rawDescOnce.Do(func() {
		file_redis_redis_proto_rawDescData = protoimpl.X.CompressGZIP(file_redis_redis_proto_rawDescData)
	})
	return file_redis_redis_proto_rawDescData
}

var file_redis_redis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_redis_redis_proto_goTypes = []interface{}{
	(*Config)(nil),                 // 0: redis.Config
	(*durationpb.Duration)(nil),    // 1: google.protobuf.Duration
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 3: google.protobuf.Int32Value
}
var file_redis_redis_proto_depIdxs = []int32{
	1, // 0: redis.Config.read_timeout:type_name -> google.protobuf.Duration
	1, // 1: redis.Config.write_timeout:type_name -> google.protobuf.Duration
	2, // 2: redis.Config.username:type_name -> google.protobuf.StringValue
	2, // 3: redis.Config.password:type_name -> google.protobuf.StringValue
	3, // 4: redis.Config.db:type_name -> google.protobuf.Int32Value
	3, // 5: redis.Config.max_retries:type_name -> google.protobuf.Int32Value
	1, // 6: redis.Config.min_retry_backoff:type_name -> google.protobuf.Duration
	1, // 7: redis.Config.max_retry_backoff:type_name -> google.protobuf.Duration
	1, // 8: redis.Config.dial_timeout:type_name -> google.protobuf.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_redis_redis_proto_init() }
func file_redis_redis_proto_init() {
	if File_redis_redis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redis_redis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_redis_redis_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redis_redis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redis_redis_proto_goTypes,
		DependencyIndexes: file_redis_redis_proto_depIdxs,
		MessageInfos:      file_redis_redis_proto_msgTypes,
	}.Build()
	File_redis_redis_proto = out.File
	file_redis_redis_proto_rawDesc = nil
	file_redis_redis_proto_goTypes = nil
	file_redis_redis_proto_depIdxs = nil
}
