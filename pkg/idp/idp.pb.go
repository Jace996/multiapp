// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: idp/idp.proto

package idp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
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

	Wechat *WeChat `protobuf:"bytes,10,opt,name=wechat,proto3" json:"wechat,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[0]
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
	return file_idp_idp_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetWechat() *WeChat {
	if x != nil {
		return x.Wechat
	}
	return nil
}

type WeChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenPlatform    map[string]*WeChat_OpenPlatform    `protobuf:"bytes,1,rep,name=open_platform,json=openPlatform,proto3" json:"open_platform,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OfficialAccount map[string]*WeChat_OfficialAccount `protobuf:"bytes,2,rep,name=official_account,json=officialAccount,proto3" json:"official_account,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MiniProgram     map[string]*WeChat_MiniProgram     `protobuf:"bytes,3,rep,name=mini_program,json=miniProgram,proto3" json:"mini_program,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MiniGame        map[string]*WeChat_MiniGame        `protobuf:"bytes,4,rep,name=mini_game,json=miniGame,proto3" json:"mini_game,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pay             map[string]*WeChat_Pay             `protobuf:"bytes,5,rep,name=pay,proto3" json:"pay,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Work            map[string]*WeChat_Work            `protobuf:"bytes,6,rep,name=work,proto3" json:"work,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WeChat) Reset() {
	*x = WeChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat) ProtoMessage() {}

func (x *WeChat) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat.ProtoReflect.Descriptor instead.
func (*WeChat) Descriptor() ([]byte, []int) {
	return file_idp_idp_proto_rawDescGZIP(), []int{1}
}

func (x *WeChat) GetOpenPlatform() map[string]*WeChat_OpenPlatform {
	if x != nil {
		return x.OpenPlatform
	}
	return nil
}

func (x *WeChat) GetOfficialAccount() map[string]*WeChat_OfficialAccount {
	if x != nil {
		return x.OfficialAccount
	}
	return nil
}

func (x *WeChat) GetMiniProgram() map[string]*WeChat_MiniProgram {
	if x != nil {
		return x.MiniProgram
	}
	return nil
}

func (x *WeChat) GetMiniGame() map[string]*WeChat_MiniGame {
	if x != nil {
		return x.MiniGame
	}
	return nil
}

func (x *WeChat) GetPay() map[string]*WeChat_Pay {
	if x != nil {
		return x.Pay
	}
	return nil
}

func (x *WeChat) GetWork() map[string]*WeChat_Work {
	if x != nil {
		return x.Work
	}
	return nil
}

type WeChat_OpenPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId          string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppSecret      string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
	Token          string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	EncodingAesKey string `protobuf:"bytes,4,opt,name=encoding_aes_key,json=encodingAesKey,proto3" json:"encoding_aes_key,omitempty"`
}

func (x *WeChat_OpenPlatform) Reset() {
	*x = WeChat_OpenPlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat_OpenPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat_OpenPlatform) ProtoMessage() {}

func (x *WeChat_OpenPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat_OpenPlatform.ProtoReflect.Descriptor instead.
func (*WeChat_OpenPlatform) Descriptor() ([]byte, []int) {
	return file_idp_idp_proto_rawDescGZIP(), []int{1, 0}
}

func (x *WeChat_OpenPlatform) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WeChat_OpenPlatform) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

func (x *WeChat_OpenPlatform) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *WeChat_OpenPlatform) GetEncodingAesKey() string {
	if x != nil {
		return x.EncodingAesKey
	}
	return ""
}

type WeChat_OfficialAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId          string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppSecret      string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
	Token          string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	EncodingAesKey string `protobuf:"bytes,4,opt,name=encoding_aes_key,json=encodingAesKey,proto3" json:"encoding_aes_key,omitempty"`
}

func (x *WeChat_OfficialAccount) Reset() {
	*x = WeChat_OfficialAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat_OfficialAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat_OfficialAccount) ProtoMessage() {}

func (x *WeChat_OfficialAccount) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat_OfficialAccount.ProtoReflect.Descriptor instead.
func (*WeChat_OfficialAccount) Descriptor() ([]byte, []int) {
	return file_idp_idp_proto_rawDescGZIP(), []int{1, 1}
}

func (x *WeChat_OfficialAccount) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WeChat_OfficialAccount) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

func (x *WeChat_OfficialAccount) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *WeChat_OfficialAccount) GetEncodingAesKey() string {
	if x != nil {
		return x.EncodingAesKey
	}
	return ""
}

type WeChat_MiniProgram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppSecret string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
}

func (x *WeChat_MiniProgram) Reset() {
	*x = WeChat_MiniProgram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat_MiniProgram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat_MiniProgram) ProtoMessage() {}

func (x *WeChat_MiniProgram) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat_MiniProgram.ProtoReflect.Descriptor instead.
func (*WeChat_MiniProgram) Descriptor() ([]byte, []int) {
	return file_idp_idp_proto_rawDescGZIP(), []int{1, 2}
}

func (x *WeChat_MiniProgram) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WeChat_MiniProgram) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

type WeChat_MiniGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppSecret string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
}

func (x *WeChat_MiniGame) Reset() {
	*x = WeChat_MiniGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat_MiniGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat_MiniGame) ProtoMessage() {}

func (x *WeChat_MiniGame) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat_MiniGame.ProtoReflect.Descriptor instead.
func (*WeChat_MiniGame) Descriptor() ([]byte, []int) {
	return file_idp_idp_proto_rawDescGZIP(), []int{1, 3}
}

func (x *WeChat_MiniGame) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WeChat_MiniGame) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

type WeChat_Work struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// corp_id
	CorpId string `protobuf:"bytes,1,opt,name=corp_id,json=corpId,proto3" json:"corp_id,omitempty"`
	// corp_secret,如果需要获取会话存档实例，当前参数请填写聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
	CorpSecret string `protobuf:"bytes,2,opt,name=corp_secret,json=corpSecret,proto3" json:"corp_secret,omitempty"`
	// agent_id
	AgentId       string `protobuf:"bytes,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	RasPrivateKey string `protobuf:"bytes,8,opt,name=ras_private_key,json=rasPrivateKey,proto3" json:"ras_private_key,omitempty"`
	// 微信客服回调配置，用于生成签名校验回调请求的合法性
	Token string `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty"`
	// 微信客服回调p配置，用于解密回调消息内容对应的密文
	EncodingAesKey string `protobuf:"bytes,10,opt,name=encoding_aes_key,json=encodingAesKey,proto3" json:"encoding_aes_key,omitempty"`
}

func (x *WeChat_Work) Reset() {
	*x = WeChat_Work{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat_Work) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat_Work) ProtoMessage() {}

func (x *WeChat_Work) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat_Work.ProtoReflect.Descriptor instead.
func (*WeChat_Work) Descriptor() ([]byte, []int) {
	return file_idp_idp_proto_rawDescGZIP(), []int{1, 4}
}

func (x *WeChat_Work) GetCorpId() string {
	if x != nil {
		return x.CorpId
	}
	return ""
}

func (x *WeChat_Work) GetCorpSecret() string {
	if x != nil {
		return x.CorpSecret
	}
	return ""
}

func (x *WeChat_Work) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *WeChat_Work) GetRasPrivateKey() string {
	if x != nil {
		return x.RasPrivateKey
	}
	return ""
}

func (x *WeChat_Work) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *WeChat_Work) GetEncodingAesKey() string {
	if x != nil {
		return x.EncodingAesKey
	}
	return ""
}

type WeChat_Pay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	MchId     string `protobuf:"bytes,3,opt,name=mch_id,json=mchId,proto3" json:"mch_id,omitempty"`
	NotifyUrl string `protobuf:"bytes,4,opt,name=notify_url,json=notifyUrl,proto3" json:"notify_url,omitempty"`
}

func (x *WeChat_Pay) Reset() {
	*x = WeChat_Pay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idp_idp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeChat_Pay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeChat_Pay) ProtoMessage() {}

func (x *WeChat_Pay) ProtoReflect() protoreflect.Message {
	mi := &file_idp_idp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeChat_Pay.ProtoReflect.Descriptor instead.
func (*WeChat_Pay) Descriptor() ([]byte, []int) {
	return file_idp_idp_proto_rawDescGZIP(), []int{1, 5}
}

func (x *WeChat_Pay) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WeChat_Pay) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WeChat_Pay) GetMchId() string {
	if x != nil {
		return x.MchId
	}
	return ""
}

func (x *WeChat_Pay) GetNotifyUrl() string {
	if x != nil {
		return x.NotifyUrl
	}
	return ""
}

var File_idp_idp_proto protoreflect.FileDescriptor

var file_idp_idp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x64, 0x70, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x69, 0x64, 0x70, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23,
	0x0a, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x06, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x22, 0xa5, 0x0c, 0x0a, 0x06, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x42,
	0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x4b, 0x0a, 0x10, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69,
	0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3f, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x12, 0x36, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x70, 0x61, 0x79, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x70, 0x61, 0x79,
	0x12, 0x29, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x84, 0x01, 0x0a, 0x0c,
	0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x65, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x65, 0x73, 0x4b,
	0x65, 0x79, 0x1a, 0x87, 0x01, 0x0a, 0x0f, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x65, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x1a, 0x43, 0x0a, 0x0b,
	0x4d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x1a, 0x40, 0x0a, 0x08, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x1a, 0xc3, 0x01, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x6f, 0x72, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x72, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x70, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x70,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x61, 0x73, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x28, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x65, 0x73, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x41, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x1a, 0x64, 0x0a, 0x03, 0x50, 0x61, 0x79,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x63, 0x68,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x1a,
	0x59, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5f, 0x0a, 0x14, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x57, 0x0a, 0x10, 0x4d,
	0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x69,
	0x6e, 0x69, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x47, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x49, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x69, 0x64, 0x70, 0x2e, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x63, 0x65, 0x39, 0x39,
	0x36, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69,
	0x64, 0x70, 0x3b, 0x69, 0x64, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idp_idp_proto_rawDescOnce sync.Once
	file_idp_idp_proto_rawDescData = file_idp_idp_proto_rawDesc
)

func file_idp_idp_proto_rawDescGZIP() []byte {
	file_idp_idp_proto_rawDescOnce.Do(func() {
		file_idp_idp_proto_rawDescData = protoimpl.X.CompressGZIP(file_idp_idp_proto_rawDescData)
	})
	return file_idp_idp_proto_rawDescData
}

var file_idp_idp_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_idp_idp_proto_goTypes = []any{
	(*Config)(nil),                 // 0: idp.Config
	(*WeChat)(nil),                 // 1: idp.WeChat
	(*WeChat_OpenPlatform)(nil),    // 2: idp.WeChat.OpenPlatform
	(*WeChat_OfficialAccount)(nil), // 3: idp.WeChat.OfficialAccount
	(*WeChat_MiniProgram)(nil),     // 4: idp.WeChat.MiniProgram
	(*WeChat_MiniGame)(nil),        // 5: idp.WeChat.MiniGame
	(*WeChat_Work)(nil),            // 6: idp.WeChat.Work
	(*WeChat_Pay)(nil),             // 7: idp.WeChat.Pay
	nil,                            // 8: idp.WeChat.OpenPlatformEntry
	nil,                            // 9: idp.WeChat.OfficialAccountEntry
	nil,                            // 10: idp.WeChat.MiniProgramEntry
	nil,                            // 11: idp.WeChat.MiniGameEntry
	nil,                            // 12: idp.WeChat.PayEntry
	nil,                            // 13: idp.WeChat.WorkEntry
}
var file_idp_idp_proto_depIdxs = []int32{
	1,  // 0: idp.Config.wechat:type_name -> idp.WeChat
	8,  // 1: idp.WeChat.open_platform:type_name -> idp.WeChat.OpenPlatformEntry
	9,  // 2: idp.WeChat.official_account:type_name -> idp.WeChat.OfficialAccountEntry
	10, // 3: idp.WeChat.mini_program:type_name -> idp.WeChat.MiniProgramEntry
	11, // 4: idp.WeChat.mini_game:type_name -> idp.WeChat.MiniGameEntry
	12, // 5: idp.WeChat.pay:type_name -> idp.WeChat.PayEntry
	13, // 6: idp.WeChat.work:type_name -> idp.WeChat.WorkEntry
	2,  // 7: idp.WeChat.OpenPlatformEntry.value:type_name -> idp.WeChat.OpenPlatform
	3,  // 8: idp.WeChat.OfficialAccountEntry.value:type_name -> idp.WeChat.OfficialAccount
	4,  // 9: idp.WeChat.MiniProgramEntry.value:type_name -> idp.WeChat.MiniProgram
	5,  // 10: idp.WeChat.MiniGameEntry.value:type_name -> idp.WeChat.MiniGame
	7,  // 11: idp.WeChat.PayEntry.value:type_name -> idp.WeChat.Pay
	6,  // 12: idp.WeChat.WorkEntry.value:type_name -> idp.WeChat.Work
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_idp_idp_proto_init() }
func file_idp_idp_proto_init() {
	if File_idp_idp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idp_idp_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_idp_idp_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WeChat); i {
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
		file_idp_idp_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WeChat_OpenPlatform); i {
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
		file_idp_idp_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WeChat_OfficialAccount); i {
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
		file_idp_idp_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WeChat_MiniProgram); i {
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
		file_idp_idp_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*WeChat_MiniGame); i {
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
		file_idp_idp_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*WeChat_Work); i {
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
		file_idp_idp_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*WeChat_Pay); i {
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
			RawDescriptor: file_idp_idp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_idp_idp_proto_goTypes,
		DependencyIndexes: file_idp_idp_proto_depIdxs,
		MessageInfos:      file_idp_idp_proto_msgTypes,
	}.Build()
	File_idp_idp_proto = out.File
	file_idp_idp_proto_rawDesc = nil
	file_idp_idp_proto_goTypes = nil
	file_idp_idp_proto_depIdxs = nil
}
