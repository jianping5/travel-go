// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: user.proto

package pb

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

// 用户登录
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=AccessExpire,proto3" json:"AccessExpire,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

// 用户注册
type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisterResp) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *RegisterResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// 用户信息
type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfoReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Account    string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	Email      string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Signature  string `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Avatar     string `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	CreateTime string `protobuf:"bytes,6,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime string `protobuf:"bytes,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *UserInfoResp) Reset() {
	*x = UserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResp) ProtoMessage() {}

func (x *UserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResp.ProtoReflect.Descriptor instead.
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResp) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserInfoResp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoResp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UserInfoResp) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfoResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *UserInfoResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type GenerateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GenerateTokenReq) Reset() {
	*x = GenerateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenReq) ProtoMessage() {}

func (x *GenerateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenReq.ProtoReflect.Descriptor instead.
func (*GenerateTokenReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *GenerateTokenReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GenerateTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
}

func (x *GenerateTokenResp) Reset() {
	*x = GenerateTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResp) ProtoMessage() {}

func (x *GenerateTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResp.ProtoReflect.Descriptor instead.
func (*GenerateTokenResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateTokenResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GenerateTokenResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

type SearchUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortType int32  `protobuf:"varint,1,opt,name=SortType,proto3" json:"SortType,omitempty"`
	Offset   int32  `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Keyword  string `protobuf:"bytes,4,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
}

func (x *SearchUserReq) Reset() {
	*x = SearchUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserReq) ProtoMessage() {}

func (x *SearchUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserReq.ProtoReflect.Descriptor instead.
func (*SearchUserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *SearchUserReq) GetSortType() int32 {
	if x != nil {
		return x.SortType
	}
	return 0
}

func (x *SearchUserReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SearchUserReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchUserReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type UserInfoView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Account    string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	Avatar     string `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Email      string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	CreateTime string `protobuf:"bytes,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *UserInfoView) Reset() {
	*x = UserInfoView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoView) ProtoMessage() {}

func (x *UserInfoView) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoView.ProtoReflect.Descriptor instead.
func (*UserInfoView) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *UserInfoView) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoView) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserInfoView) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfoView) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoView) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type SearchUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32           `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Users []*UserInfoView `protobuf:"bytes,2,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *SearchUserResp) Reset() {
	*x = SearchUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResp) ProtoMessage() {}

func (x *SearchUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResp.ProtoReflect.Descriptor instead.
func (*SearchUserResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *SearchUserResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchUserResp) GetUsers() []*UserInfoView {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserIdsReq) Reset() {
	*x = GetUserIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIdsReq) ProtoMessage() {}

func (x *GetUserIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIdsReq.ProtoReflect.Descriptor instead.
func (*GetUserIdsReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

type GetUserIdsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []int64 `protobuf:"varint,1,rep,packed,name=UserIds,proto3" json:"UserIds,omitempty"`
}

func (x *GetUserIdsResp) Reset() {
	*x = GetUserIdsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIdsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIdsResp) ProtoMessage() {}

func (x *GetUserIdsResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIdsResp.ProtoReflect.Descriptor instead.
func (*GetUserIdsResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *GetUserIdsResp) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x40, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x51, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x59, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x4e, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22,
	0xc4, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x59, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x79, 0x0a,
	0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x4e, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x69, 0x65, 0x77, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x22, 0x2a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x32, 0xb2,
	0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0d, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x0a, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_user_proto_goTypes = []interface{}{
	(*LoginReq)(nil),          // 0: pb.LoginReq
	(*LoginResp)(nil),         // 1: pb.LoginResp
	(*RegisterReq)(nil),       // 2: pb.RegisterReq
	(*RegisterResp)(nil),      // 3: pb.RegisterResp
	(*UserInfoReq)(nil),       // 4: pb.UserInfoReq
	(*UserInfoResp)(nil),      // 5: pb.UserInfoResp
	(*GenerateTokenReq)(nil),  // 6: pb.GenerateTokenReq
	(*GenerateTokenResp)(nil), // 7: pb.GenerateTokenResp
	(*SearchUserReq)(nil),     // 8: pb.SearchUserReq
	(*UserInfoView)(nil),      // 9: pb.UserInfoView
	(*SearchUserResp)(nil),    // 10: pb.SearchUserResp
	(*GetUserIdsReq)(nil),     // 11: pb.GetUserIdsReq
	(*GetUserIdsResp)(nil),    // 12: pb.GetUserIdsResp
}
var file_user_proto_depIdxs = []int32{
	9,  // 0: pb.SearchUserResp.Users:type_name -> pb.UserInfoView
	0,  // 1: pb.User.Login:input_type -> pb.LoginReq
	2,  // 2: pb.User.Register:input_type -> pb.RegisterReq
	4,  // 3: pb.User.UserInfo:input_type -> pb.UserInfoReq
	6,  // 4: pb.User.generateToken:input_type -> pb.GenerateTokenReq
	8,  // 5: pb.User.SearchUser:input_type -> pb.SearchUserReq
	11, // 6: pb.User.GetUserIds:input_type -> pb.GetUserIdsReq
	1,  // 7: pb.User.Login:output_type -> pb.LoginResp
	3,  // 8: pb.User.Register:output_type -> pb.RegisterResp
	5,  // 9: pb.User.UserInfo:output_type -> pb.UserInfoResp
	7,  // 10: pb.User.generateToken:output_type -> pb.GenerateTokenResp
	10, // 11: pb.User.SearchUser:output_type -> pb.SearchUserResp
	12, // 12: pb.User.GetUserIds:output_type -> pb.GetUserIdsResp
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResp); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenReq); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenResp); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserReq); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoView); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserResp); i {
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
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserIdsReq); i {
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
		file_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserIdsResp); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
