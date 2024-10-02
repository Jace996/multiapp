// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: data/data.proto

package data

import (
	query "github.com/jace996/multiapp/pkg/query"
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

type DynamicValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*DynamicValue_IntValue
	//	*DynamicValue_LongValue
	//	*DynamicValue_FloatValue
	//	*DynamicValue_DoubleValue
	//	*DynamicValue_StringValue
	//	*DynamicValue_BoolValue
	//	*DynamicValue_NullValue
	//	*DynamicValue_JsonValue
	Value isDynamicValue_Value `protobuf_oneof:"value"`
}

func (x *DynamicValue) Reset() {
	*x = DynamicValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicValue) ProtoMessage() {}

func (x *DynamicValue) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicValue.ProtoReflect.Descriptor instead.
func (*DynamicValue) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0}
}

func (m *DynamicValue) GetValue() isDynamicValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *DynamicValue) GetIntValue() int32 {
	if x, ok := x.GetValue().(*DynamicValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *DynamicValue) GetLongValue() int64 {
	if x, ok := x.GetValue().(*DynamicValue_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *DynamicValue) GetFloatValue() float32 {
	if x, ok := x.GetValue().(*DynamicValue_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *DynamicValue) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*DynamicValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *DynamicValue) GetStringValue() string {
	if x, ok := x.GetValue().(*DynamicValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *DynamicValue) GetBoolValue() bool {
	if x, ok := x.GetValue().(*DynamicValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *DynamicValue) GetNullValue() structpb.NullValue {
	if x, ok := x.GetValue().(*DynamicValue_NullValue); ok {
		return x.NullValue
	}
	return structpb.NullValue(0)
}

func (x *DynamicValue) GetJsonValue() *structpb.Struct {
	if x, ok := x.GetValue().(*DynamicValue_JsonValue); ok {
		return x.JsonValue
	}
	return nil
}

type isDynamicValue_Value interface {
	isDynamicValue_Value()
}

type DynamicValue_IntValue struct {
	IntValue int32 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type DynamicValue_LongValue struct {
	LongValue int64 `protobuf:"varint,3,opt,name=long_value,json=longValue,proto3,oneof"`
}

type DynamicValue_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,4,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type DynamicValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type DynamicValue_StringValue struct {
	StringValue string `protobuf:"bytes,6,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type DynamicValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,7,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type DynamicValue_NullValue struct {
	NullValue structpb.NullValue `protobuf:"varint,8,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type DynamicValue_JsonValue struct {
	JsonValue *structpb.Struct `protobuf:"bytes,9,opt,name=json_value,json=jsonValue,proto3,oneof"`
}

func (*DynamicValue_IntValue) isDynamicValue_Value() {}

func (*DynamicValue_LongValue) isDynamicValue_Value() {}

func (*DynamicValue_FloatValue) isDynamicValue_Value() {}

func (*DynamicValue_DoubleValue) isDynamicValue_Value() {}

func (*DynamicValue_StringValue) isDynamicValue_Value() {}

func (*DynamicValue_BoolValue) isDynamicValue_Value() {}

func (*DynamicValue_NullValue) isDynamicValue_Value() {}

func (*DynamicValue_JsonValue) isDynamicValue_Value() {}

type DynamicValueFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Filter:
	//
	//	*DynamicValueFilter_IntValue
	//	*DynamicValueFilter_LongValue
	//	*DynamicValueFilter_FloatValue
	//	*DynamicValueFilter_DoubleValue
	//	*DynamicValueFilter_StringValue
	//	*DynamicValueFilter_BoolValue
	//	*DynamicValueFilter_NullValue
	Filter isDynamicValueFilter_Filter `protobuf_oneof:"filter"`
}

func (x *DynamicValueFilter) Reset() {
	*x = DynamicValueFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicValueFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicValueFilter) ProtoMessage() {}

func (x *DynamicValueFilter) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicValueFilter.ProtoReflect.Descriptor instead.
func (*DynamicValueFilter) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{1}
}

func (m *DynamicValueFilter) GetFilter() isDynamicValueFilter_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (x *DynamicValueFilter) GetIntValue() *query.Int32FilterOperators {
	if x, ok := x.GetFilter().(*DynamicValueFilter_IntValue); ok {
		return x.IntValue
	}
	return nil
}

func (x *DynamicValueFilter) GetLongValue() *query.Int64FilterOperators {
	if x, ok := x.GetFilter().(*DynamicValueFilter_LongValue); ok {
		return x.LongValue
	}
	return nil
}

func (x *DynamicValueFilter) GetFloatValue() *query.FloatFilterOperators {
	if x, ok := x.GetFilter().(*DynamicValueFilter_FloatValue); ok {
		return x.FloatValue
	}
	return nil
}

func (x *DynamicValueFilter) GetDoubleValue() *query.DoubleFilterOperators {
	if x, ok := x.GetFilter().(*DynamicValueFilter_DoubleValue); ok {
		return x.DoubleValue
	}
	return nil
}

func (x *DynamicValueFilter) GetStringValue() *query.StringFilterOperation {
	if x, ok := x.GetFilter().(*DynamicValueFilter_StringValue); ok {
		return x.StringValue
	}
	return nil
}

func (x *DynamicValueFilter) GetBoolValue() *query.BooleanFilterOperators {
	if x, ok := x.GetFilter().(*DynamicValueFilter_BoolValue); ok {
		return x.BoolValue
	}
	return nil
}

func (x *DynamicValueFilter) GetNullValue() *query.NullFilterOperators {
	if x, ok := x.GetFilter().(*DynamicValueFilter_NullValue); ok {
		return x.NullValue
	}
	return nil
}

type isDynamicValueFilter_Filter interface {
	isDynamicValueFilter_Filter()
}

type DynamicValueFilter_IntValue struct {
	IntValue *query.Int32FilterOperators `protobuf:"bytes,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type DynamicValueFilter_LongValue struct {
	LongValue *query.Int64FilterOperators `protobuf:"bytes,3,opt,name=long_value,json=longValue,proto3,oneof"`
}

type DynamicValueFilter_FloatValue struct {
	FloatValue *query.FloatFilterOperators `protobuf:"bytes,4,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type DynamicValueFilter_DoubleValue struct {
	DoubleValue *query.DoubleFilterOperators `protobuf:"bytes,5,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type DynamicValueFilter_StringValue struct {
	StringValue *query.StringFilterOperation `protobuf:"bytes,6,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type DynamicValueFilter_BoolValue struct {
	BoolValue *query.BooleanFilterOperators `protobuf:"bytes,7,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type DynamicValueFilter_NullValue struct {
	NullValue *query.NullFilterOperators `protobuf:"bytes,8,opt,name=null_value,json=nullValue,proto3,oneof"`
}

func (*DynamicValueFilter_IntValue) isDynamicValueFilter_Filter() {}

func (*DynamicValueFilter_LongValue) isDynamicValueFilter_Filter() {}

func (*DynamicValueFilter_FloatValue) isDynamicValueFilter_Filter() {}

func (*DynamicValueFilter_DoubleValue) isDynamicValueFilter_Filter() {}

func (*DynamicValueFilter_StringValue) isDynamicValueFilter_Filter() {}

func (*DynamicValueFilter_BoolValue) isDynamicValueFilter_Filter() {}

func (*DynamicValueFilter_NullValue) isDynamicValueFilter_Filter() {}

var File_data_data_proto protoreflect.FileDescriptor

var file_data_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a,
	0x0c, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a,
	0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a,
	0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f,
	0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x6e,
	0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6e,
	0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa1, 0x04, 0x0a, 0x12,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x48, 0x00, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x48, 0x00, 0x52, 0x0a,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x45,
	0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_data_proto_rawDescOnce sync.Once
	file_data_data_proto_rawDescData = file_data_data_proto_rawDesc
)

func file_data_data_proto_rawDescGZIP() []byte {
	file_data_data_proto_rawDescOnce.Do(func() {
		file_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_data_proto_rawDescData)
	})
	return file_data_data_proto_rawDescData
}

var file_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_data_data_proto_goTypes = []interface{}{
	(*DynamicValue)(nil),                 // 0: data.DynamicValue
	(*DynamicValueFilter)(nil),           // 1: data.DynamicValueFilter
	(structpb.NullValue)(0),              // 2: google.protobuf.NullValue
	(*structpb.Struct)(nil),              // 3: google.protobuf.Struct
	(*query.Int32FilterOperators)(nil),   // 4: query.operation.Int32FilterOperators
	(*query.Int64FilterOperators)(nil),   // 5: query.operation.Int64FilterOperators
	(*query.FloatFilterOperators)(nil),   // 6: query.operation.FloatFilterOperators
	(*query.DoubleFilterOperators)(nil),  // 7: query.operation.DoubleFilterOperators
	(*query.StringFilterOperation)(nil),  // 8: query.operation.StringFilterOperation
	(*query.BooleanFilterOperators)(nil), // 9: query.operation.BooleanFilterOperators
	(*query.NullFilterOperators)(nil),    // 10: query.operation.NullFilterOperators
}
var file_data_data_proto_depIdxs = []int32{
	2,  // 0: data.DynamicValue.null_value:type_name -> google.protobuf.NullValue
	3,  // 1: data.DynamicValue.json_value:type_name -> google.protobuf.Struct
	4,  // 2: data.DynamicValueFilter.int_value:type_name -> query.operation.Int32FilterOperators
	5,  // 3: data.DynamicValueFilter.long_value:type_name -> query.operation.Int64FilterOperators
	6,  // 4: data.DynamicValueFilter.float_value:type_name -> query.operation.FloatFilterOperators
	7,  // 5: data.DynamicValueFilter.double_value:type_name -> query.operation.DoubleFilterOperators
	8,  // 6: data.DynamicValueFilter.string_value:type_name -> query.operation.StringFilterOperation
	9,  // 7: data.DynamicValueFilter.bool_value:type_name -> query.operation.BooleanFilterOperators
	10, // 8: data.DynamicValueFilter.null_value:type_name -> query.operation.NullFilterOperators
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_data_data_proto_init() }
func file_data_data_proto_init() {
	if File_data_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicValue); i {
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
		file_data_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicValueFilter); i {
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
	file_data_data_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DynamicValue_IntValue)(nil),
		(*DynamicValue_LongValue)(nil),
		(*DynamicValue_FloatValue)(nil),
		(*DynamicValue_DoubleValue)(nil),
		(*DynamicValue_StringValue)(nil),
		(*DynamicValue_BoolValue)(nil),
		(*DynamicValue_NullValue)(nil),
		(*DynamicValue_JsonValue)(nil),
	}
	file_data_data_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DynamicValueFilter_IntValue)(nil),
		(*DynamicValueFilter_LongValue)(nil),
		(*DynamicValueFilter_FloatValue)(nil),
		(*DynamicValueFilter_DoubleValue)(nil),
		(*DynamicValueFilter_StringValue)(nil),
		(*DynamicValueFilter_BoolValue)(nil),
		(*DynamicValueFilter_NullValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_data_proto_goTypes,
		DependencyIndexes: file_data_data_proto_depIdxs,
		MessageInfos:      file_data_data_proto_msgTypes,
	}.Build()
	File_data_data_proto = out.File
	file_data_data_proto_rawDesc = nil
	file_data_data_proto_goTypes = nil
	file_data_data_proto_depIdxs = nil
}
