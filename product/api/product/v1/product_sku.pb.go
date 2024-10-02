// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: product/api/product/v1/product_sku.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/jace996/multiapp/pkg/query"
	v1 "github.com/jace996/multiapp/product/api/price/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProductSku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	MainPic *ProductMedia     `protobuf:"bytes,13,opt,name=main_pic,json=mainPic,proto3" json:"main_pic,omitempty"`
	Medias  []*ProductMedia   `protobuf:"bytes,14,rep,name=medias,proto3" json:"medias,omitempty"`
	Prices  []*v1.PriceParams `protobuf:"bytes,26,rep,name=prices,proto3" json:"prices,omitempty"`
	Barcode string            `protobuf:"bytes,19,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Stock   []*Stock          `protobuf:"bytes,62,rep,name=stock,proto3" json:"stock,omitempty"`
}

func (x *CreateProductSku) Reset() {
	*x = CreateProductSku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_product_sku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductSku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductSku) ProtoMessage() {}

func (x *CreateProductSku) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_product_sku_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductSku.ProtoReflect.Descriptor instead.
func (*CreateProductSku) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_product_sku_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductSku) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProductSku) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateProductSku) GetMainPic() *ProductMedia {
	if x != nil {
		return x.MainPic
	}
	return nil
}

func (x *CreateProductSku) GetMedias() []*ProductMedia {
	if x != nil {
		return x.Medias
	}
	return nil
}

func (x *CreateProductSku) GetPrices() []*v1.PriceParams {
	if x != nil {
		return x.Prices
	}
	return nil
}

func (x *CreateProductSku) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *CreateProductSku) GetStock() []*Stock {
	if x != nil {
		return x.Stock
	}
	return nil
}

type ProductSku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	TenantId     string                 `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Version      string                 `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	Title        string                 `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	MainPic      *ProductMedia          `protobuf:"bytes,13,opt,name=main_pic,json=mainPic,proto3" json:"main_pic,omitempty"`
	Medias       []*ProductMedia        `protobuf:"bytes,14,rep,name=medias,proto3" json:"medias,omitempty"`
	Barcode      string                 `protobuf:"bytes,19,opt,name=barcode,proto3" json:"barcode,omitempty"`
	SaleableFrom *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=saleable_from,json=saleableFrom,proto3" json:"saleable_from,omitempty"`
	SaleableTo   *timestamppb.Timestamp `protobuf:"bytes,24,opt,name=saleable_to,json=saleableTo,proto3" json:"saleable_to,omitempty"`
	Keywords     []*Keyword             `protobuf:"bytes,25,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Prices       []*v1.Price            `protobuf:"bytes,26,rep,name=prices,proto3" json:"prices,omitempty"`
	Stocks       []*Stock               `protobuf:"bytes,62,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *ProductSku) Reset() {
	*x = ProductSku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_product_sku_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSku) ProtoMessage() {}

func (x *ProductSku) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_product_sku_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSku.ProtoReflect.Descriptor instead.
func (*ProductSku) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_product_sku_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSku) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductSku) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProductSku) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProductSku) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ProductSku) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ProductSku) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductSku) GetMainPic() *ProductMedia {
	if x != nil {
		return x.MainPic
	}
	return nil
}

func (x *ProductSku) GetMedias() []*ProductMedia {
	if x != nil {
		return x.Medias
	}
	return nil
}

func (x *ProductSku) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *ProductSku) GetSaleableFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.SaleableFrom
	}
	return nil
}

func (x *ProductSku) GetSaleableTo() *timestamppb.Timestamp {
	if x != nil {
		return x.SaleableTo
	}
	return nil
}

func (x *ProductSku) GetKeywords() []*Keyword {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *ProductSku) GetPrices() []*v1.Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

func (x *ProductSku) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

var File_product_api_product_v1_product_sku_proto protoreflect.FileDescriptor

var file_product_api_product_v1_product_sku_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x69, 0x63,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x07, 0x6d, 0x61,
	0x69, 0x6e, 0x50, 0x69, 0x63, 0x12, 0x3c, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x06, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x1a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x3e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x9f, 0x05,
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x3f, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x07, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x69, 0x63,
	0x12, 0x3c, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x61, 0x6c, 0x65,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x73, 0x61, 0x6c,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x61, 0x6c,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x12, 0x3b, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x1a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x3e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_api_product_v1_product_sku_proto_rawDescOnce sync.Once
	file_product_api_product_v1_product_sku_proto_rawDescData = file_product_api_product_v1_product_sku_proto_rawDesc
)

func file_product_api_product_v1_product_sku_proto_rawDescGZIP() []byte {
	file_product_api_product_v1_product_sku_proto_rawDescOnce.Do(func() {
		file_product_api_product_v1_product_sku_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_api_product_v1_product_sku_proto_rawDescData)
	})
	return file_product_api_product_v1_product_sku_proto_rawDescData
}

var file_product_api_product_v1_product_sku_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_api_product_v1_product_sku_proto_goTypes = []interface{}{
	(*CreateProductSku)(nil),      // 0: product.api.product.v1.CreateProductSku
	(*ProductSku)(nil),            // 1: product.api.product.v1.ProductSku
	(*ProductMedia)(nil),          // 2: product.api.product.v1.ProductMedia
	(*v1.PriceParams)(nil),        // 3: product.api.price.v1.PriceParams
	(*Stock)(nil),                 // 4: product.api.product.v1.Stock
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Keyword)(nil),               // 6: product.api.product.v1.Keyword
	(*v1.Price)(nil),              // 7: product.api.price.v1.Price
}
var file_product_api_product_v1_product_sku_proto_depIdxs = []int32{
	2,  // 0: product.api.product.v1.CreateProductSku.main_pic:type_name -> product.api.product.v1.ProductMedia
	2,  // 1: product.api.product.v1.CreateProductSku.medias:type_name -> product.api.product.v1.ProductMedia
	3,  // 2: product.api.product.v1.CreateProductSku.prices:type_name -> product.api.price.v1.PriceParams
	4,  // 3: product.api.product.v1.CreateProductSku.stock:type_name -> product.api.product.v1.Stock
	5,  // 4: product.api.product.v1.ProductSku.created_at:type_name -> google.protobuf.Timestamp
	5,  // 5: product.api.product.v1.ProductSku.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 6: product.api.product.v1.ProductSku.main_pic:type_name -> product.api.product.v1.ProductMedia
	2,  // 7: product.api.product.v1.ProductSku.medias:type_name -> product.api.product.v1.ProductMedia
	5,  // 8: product.api.product.v1.ProductSku.saleable_from:type_name -> google.protobuf.Timestamp
	5,  // 9: product.api.product.v1.ProductSku.saleable_to:type_name -> google.protobuf.Timestamp
	6,  // 10: product.api.product.v1.ProductSku.keywords:type_name -> product.api.product.v1.Keyword
	7,  // 11: product.api.product.v1.ProductSku.prices:type_name -> product.api.price.v1.Price
	4,  // 12: product.api.product.v1.ProductSku.stocks:type_name -> product.api.product.v1.Stock
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_product_api_product_v1_product_sku_proto_init() }
func file_product_api_product_v1_product_sku_proto_init() {
	if File_product_api_product_v1_product_sku_proto != nil {
		return
	}
	file_product_api_product_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_api_product_v1_product_sku_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductSku); i {
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
		file_product_api_product_v1_product_sku_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSku); i {
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
			RawDescriptor: file_product_api_product_v1_product_sku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_api_product_v1_product_sku_proto_goTypes,
		DependencyIndexes: file_product_api_product_v1_product_sku_proto_depIdxs,
		MessageInfos:      file_product_api_product_v1_product_sku_proto_msgTypes,
	}.Build()
	File_product_api_product_v1_product_sku_proto = out.File
	file_product_api_product_v1_product_sku_proto_rawDesc = nil
	file_product_api_product_v1_product_sku_proto_goTypes = nil
	file_product_api_product_v1_product_sku_proto_depIdxs = nil
}
