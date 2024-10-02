// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: product/private/biz/job.proto

package biz

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

type ProductUpdatedJobParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId      string                             `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductVersion string                             `protobuf:"bytes,2,opt,name=product_version,json=productVersion,proto3" json:"product_version,omitempty"`
	TenantId       string                             `protobuf:"bytes,3,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	IsDelete       bool                               `protobuf:"varint,4,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	SyncLinks      []*ProductUpdatedJobParam_SyncLink `protobuf:"bytes,5,rep,name=sync_links,json=syncLinks,proto3" json:"sync_links,omitempty"`
}

func (x *ProductUpdatedJobParam) Reset() {
	*x = ProductUpdatedJobParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_private_biz_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdatedJobParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdatedJobParam) ProtoMessage() {}

func (x *ProductUpdatedJobParam) ProtoReflect() protoreflect.Message {
	mi := &file_product_private_biz_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdatedJobParam.ProtoReflect.Descriptor instead.
func (*ProductUpdatedJobParam) Descriptor() ([]byte, []int) {
	return file_product_private_biz_job_proto_rawDescGZIP(), []int{0}
}

func (x *ProductUpdatedJobParam) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductUpdatedJobParam) GetProductVersion() string {
	if x != nil {
		return x.ProductVersion
	}
	return ""
}

func (x *ProductUpdatedJobParam) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ProductUpdatedJobParam) GetIsDelete() bool {
	if x != nil {
		return x.IsDelete
	}
	return false
}

func (x *ProductUpdatedJobParam) GetSyncLinks() []*ProductUpdatedJobParam_SyncLink {
	if x != nil {
		return x.SyncLinks
	}
	return nil
}

type ProductUpdatedJobParam_SyncLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderName string `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	ProviderId   string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
}

func (x *ProductUpdatedJobParam_SyncLink) Reset() {
	*x = ProductUpdatedJobParam_SyncLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_private_biz_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductUpdatedJobParam_SyncLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdatedJobParam_SyncLink) ProtoMessage() {}

func (x *ProductUpdatedJobParam_SyncLink) ProtoReflect() protoreflect.Message {
	mi := &file_product_private_biz_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdatedJobParam_SyncLink.ProtoReflect.Descriptor instead.
func (*ProductUpdatedJobParam_SyncLink) Descriptor() ([]byte, []int) {
	return file_product_private_biz_job_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProductUpdatedJobParam_SyncLink) GetProviderName() string {
	if x != nil {
		return x.ProviderName
	}
	return ""
}

func (x *ProductUpdatedJobParam_SyncLink) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

var File_product_private_biz_job_proto protoreflect.FileDescriptor

var file_product_private_biz_job_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x62, 0x69, 0x7a, 0x22, 0xc1, 0x02, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x53, 0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x09, 0x73, 0x79, 0x6e,
	0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x1a, 0x50, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b,
	0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2f, 0x62, 0x69, 0x7a, 0x3b, 0x62, 0x69, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_product_private_biz_job_proto_rawDescOnce sync.Once
	file_product_private_biz_job_proto_rawDescData = file_product_private_biz_job_proto_rawDesc
)

func file_product_private_biz_job_proto_rawDescGZIP() []byte {
	file_product_private_biz_job_proto_rawDescOnce.Do(func() {
		file_product_private_biz_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_private_biz_job_proto_rawDescData)
	})
	return file_product_private_biz_job_proto_rawDescData
}

var file_product_private_biz_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_private_biz_job_proto_goTypes = []interface{}{
	(*ProductUpdatedJobParam)(nil),          // 0: product.private.biz.ProductUpdatedJobParam
	(*ProductUpdatedJobParam_SyncLink)(nil), // 1: product.private.biz.ProductUpdatedJobParam.SyncLink
}
var file_product_private_biz_job_proto_depIdxs = []int32{
	1, // 0: product.private.biz.ProductUpdatedJobParam.sync_links:type_name -> product.private.biz.ProductUpdatedJobParam.SyncLink
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_private_biz_job_proto_init() }
func file_product_private_biz_job_proto_init() {
	if File_product_private_biz_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_private_biz_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductUpdatedJobParam); i {
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
		file_product_private_biz_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductUpdatedJobParam_SyncLink); i {
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
			RawDescriptor: file_product_private_biz_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_private_biz_job_proto_goTypes,
		DependencyIndexes: file_product_private_biz_job_proto_depIdxs,
		MessageInfos:      file_product_private_biz_job_proto_msgTypes,
	}.Build()
	File_product_private_biz_job_proto = out.File
	file_product_private_biz_job_proto_rawDesc = nil
	file_product_private_biz_job_proto_goTypes = nil
	file_product_private_biz_job_proto_depIdxs = nil
}
