// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: http/app/app_http.proto

package app

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreatAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName         string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	AppDesc         string `protobuf:"bytes,2,opt,name=appDesc,proto3" json:"appDesc,omitempty"`
	AppIcon         string `protobuf:"bytes,3,opt,name=appIcon,proto3" json:"appIcon,omitempty"`
	AppType         int32  `protobuf:"varint,4,opt,name=appType,proto3" json:"appType,omitempty"`
	ScoringStrategy int32  `protobuf:"varint,5,opt,name=scoringStrategy,proto3" json:"scoringStrategy,omitempty"`
}

func (x *CreatAppRequest) Reset() {
	*x = CreatAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatAppRequest) ProtoMessage() {}

func (x *CreatAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatAppRequest.ProtoReflect.Descriptor instead.
func (*CreatAppRequest) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{0}
}

func (x *CreatAppRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *CreatAppRequest) GetAppDesc() string {
	if x != nil {
		return x.AppDesc
	}
	return ""
}

func (x *CreatAppRequest) GetAppIcon() string {
	if x != nil {
		return x.AppIcon
	}
	return ""
}

func (x *CreatAppRequest) GetAppType() int32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *CreatAppRequest) GetScoringStrategy() int32 {
	if x != nil {
		return x.ScoringStrategy
	}
	return 0
}

type CreatAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatAppReply) Reset() {
	*x = CreatAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatAppReply) ProtoMessage() {}

func (x *CreatAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatAppReply.ProtoReflect.Descriptor instead.
func (*CreatAppReply) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{1}
}

type ModifyAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      int64       `protobuf:"varint,1,opt,name=appId,proto3" json:"appId,omitempty"`
	AppModInfo *AppModInfo `protobuf:"bytes,2,opt,name=appModInfo,proto3" json:"appModInfo,omitempty"`
}

func (x *ModifyAppRequest) Reset() {
	*x = ModifyAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyAppRequest) ProtoMessage() {}

func (x *ModifyAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyAppRequest.ProtoReflect.Descriptor instead.
func (*ModifyAppRequest) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{2}
}

func (x *ModifyAppRequest) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ModifyAppRequest) GetAppModInfo() *AppModInfo {
	if x != nil {
		return x.AppModInfo
	}
	return nil
}

type ModifyAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModifyAppReply) Reset() {
	*x = ModifyAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyAppReply) ProtoMessage() {}

func (x *ModifyAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyAppReply.ProtoReflect.Descriptor instead.
func (*ModifyAppReply) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{3}
}

type AppModInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName         string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	AppDesc         string `protobuf:"bytes,2,opt,name=appDesc,proto3" json:"appDesc,omitempty"`
	AppIcon         string `protobuf:"bytes,3,opt,name=appIcon,proto3" json:"appIcon,omitempty"`
	AppType         int32  `protobuf:"varint,4,opt,name=appType,proto3" json:"appType,omitempty"`
	ScoringStrategy int32  `protobuf:"varint,5,opt,name=scoringStrategy,proto3" json:"scoringStrategy,omitempty"`
}

func (x *AppModInfo) Reset() {
	*x = AppModInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppModInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppModInfo) ProtoMessage() {}

func (x *AppModInfo) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppModInfo.ProtoReflect.Descriptor instead.
func (*AppModInfo) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{4}
}

func (x *AppModInfo) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AppModInfo) GetAppDesc() string {
	if x != nil {
		return x.AppDesc
	}
	return ""
}

func (x *AppModInfo) GetAppIcon() string {
	if x != nil {
		return x.AppIcon
	}
	return ""
}

func (x *AppModInfo) GetAppType() int32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *AppModInfo) GetScoringStrategy() int32 {
	if x != nil {
		return x.ScoringStrategy
	}
	return 0
}

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName         string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	AppDesc         string `protobuf:"bytes,2,opt,name=appDesc,proto3" json:"appDesc,omitempty"`
	AppIcon         string `protobuf:"bytes,3,opt,name=appIcon,proto3" json:"appIcon,omitempty"`
	AppType         int32  `protobuf:"varint,4,opt,name=appType,proto3" json:"appType,omitempty"`
	ScoringStrategy int32  `protobuf:"varint,5,opt,name=scoringStrategy,proto3" json:"scoringStrategy,omitempty"`
	ReviewStatus    int32  `protobuf:"varint,6,opt,name=reviewStatus,proto3" json:"reviewStatus,omitempty"`
	ReviewMessage   string `protobuf:"bytes,7,opt,name=reviewMessage,proto3" json:"reviewMessage,omitempty"`
	ReviewId        int64  `protobuf:"varint,8,opt,name=reviewId,proto3" json:"reviewId,omitempty"`
	UserId          int64  `protobuf:"varint,9,opt,name=userId,proto3" json:"userId,omitempty"`
	AppId           int64  `protobuf:"varint,10,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{5}
}

func (x *AppInfo) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AppInfo) GetAppDesc() string {
	if x != nil {
		return x.AppDesc
	}
	return ""
}

func (x *AppInfo) GetAppIcon() string {
	if x != nil {
		return x.AppIcon
	}
	return ""
}

func (x *AppInfo) GetAppType() int32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *AppInfo) GetScoringStrategy() int32 {
	if x != nil {
		return x.ScoringStrategy
	}
	return 0
}

func (x *AppInfo) GetReviewStatus() int32 {
	if x != nil {
		return x.ReviewStatus
	}
	return 0
}

func (x *AppInfo) GetReviewMessage() string {
	if x != nil {
		return x.ReviewMessage
	}
	return ""
}

func (x *AppInfo) GetReviewId() int64 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *AppInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AppInfo) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Account   string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Username  string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	AvatarUrl string `protobuf:"bytes,5,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	Gender    int32  `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	UserInfo  string `protobuf:"bytes,7,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	UserRole  int32  `protobuf:"varint,8,opt,name=userRole,proto3" json:"userRole,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfo) GetUserInfo() string {
	if x != nil {
		return x.UserInfo
	}
	return ""
}

func (x *UserInfo) GetUserRole() int32 {
	if x != nil {
		return x.UserRole
	}
	return 0
}

type GetAppByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId int64 `protobuf:"varint,1,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *GetAppByIdRequest) Reset() {
	*x = GetAppByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppByIdRequest) ProtoMessage() {}

func (x *GetAppByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppByIdRequest.ProtoReflect.Descriptor instead.
func (*GetAppByIdRequest) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppByIdRequest) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type GetAppByIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App  *AppInfo  `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	User *UserInfo `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetAppByIdReply) Reset() {
	*x = GetAppByIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppByIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppByIdReply) ProtoMessage() {}

func (x *GetAppByIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppByIdReply.ProtoReflect.Descriptor instead.
func (*GetAppByIdReply) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{8}
}

func (x *GetAppByIdReply) GetApp() *AppInfo {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *GetAppByIdReply) GetUser() *UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type ListAppPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int32 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListAppPageRequest) Reset() {
	*x = ListAppPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppPageRequest) ProtoMessage() {}

func (x *ListAppPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppPageRequest.ProtoReflect.Descriptor instead.
func (*ListAppPageRequest) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{9}
}

func (x *ListAppPageRequest) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ListAppPageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListAppPageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppList []*AppInfo `protobuf:"bytes,1,rep,name=appList,proto3" json:"appList,omitempty"`
	Total   int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListAppPageReply) Reset() {
	*x = ListAppPageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_app_app_http_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppPageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppPageReply) ProtoMessage() {}

func (x *ListAppPageReply) ProtoReflect() protoreflect.Message {
	mi := &file_http_app_app_http_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppPageReply.ProtoReflect.Descriptor instead.
func (*ListAppPageReply) Descriptor() ([]byte, []int) {
	return file_http_app_app_http_proto_rawDescGZIP(), []int{10}
}

func (x *ListAppPageReply) GetAppList() []*AppInfo {
	if x != nil {
		return x.AppList
	}
	return nil
}

func (x *ListAppPageReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_http_app_app_http_proto protoreflect.FileDescriptor

var file_http_app_app_http_proto_rawDesc = []byte{
	0x0a, 0x17, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x61, 0x70, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70,
	0x49, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x49,
	0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5e, 0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x4d, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x4d, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x41,
	0x70, 0x70, 0x4d, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x70, 0x70, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x63, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0xaf, 0x02, 0x0a, 0x07,
	0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x70, 0x70, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x63, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0xda, 0x01,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x26, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x55, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x91, 0x03, 0x0a, 0x03, 0x41, 0x70, 0x70,
	0x12, 0x5b, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x41, 0x70, 0x70, 0x12, 0x19, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x61, 0x70, 0x70, 0x12, 0x5f, 0x0a,
	0x09, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x70, 0x70, 0x12, 0x1a, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x61, 0x70, 0x70, 0x12, 0x63,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22,
	0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x70, 0x70, 0x62,
	0x79, 0x69, 0x64, 0x12, 0x67, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x70, 0x70, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x70, 0x70, 0x70, 0x61, 0x67, 0x65, 0x42, 0x1c, 0x5a, 0x1a,
	0x77, 0x75, 0x7a, 0x69, 0x67, 0x6f, 0x77, 0x65, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x3b, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_http_app_app_http_proto_rawDescOnce sync.Once
	file_http_app_app_http_proto_rawDescData = file_http_app_app_http_proto_rawDesc
)

func file_http_app_app_http_proto_rawDescGZIP() []byte {
	file_http_app_app_http_proto_rawDescOnce.Do(func() {
		file_http_app_app_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_app_app_http_proto_rawDescData)
	})
	return file_http_app_app_http_proto_rawDescData
}

var file_http_app_app_http_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_http_app_app_http_proto_goTypes = []interface{}{
	(*CreatAppRequest)(nil),    // 0: http.app.CreatAppRequest
	(*CreatAppReply)(nil),      // 1: http.app.CreatAppReply
	(*ModifyAppRequest)(nil),   // 2: http.app.ModifyAppRequest
	(*ModifyAppReply)(nil),     // 3: http.app.ModifyAppReply
	(*AppModInfo)(nil),         // 4: http.app.AppModInfo
	(*AppInfo)(nil),            // 5: http.app.AppInfo
	(*UserInfo)(nil),           // 6: http.app.UserInfo
	(*GetAppByIdRequest)(nil),  // 7: http.app.GetAppByIdRequest
	(*GetAppByIdReply)(nil),    // 8: http.app.GetAppByIdReply
	(*ListAppPageRequest)(nil), // 9: http.app.ListAppPageRequest
	(*ListAppPageReply)(nil),   // 10: http.app.ListAppPageReply
}
var file_http_app_app_http_proto_depIdxs = []int32{
	4,  // 0: http.app.ModifyAppRequest.appModInfo:type_name -> http.app.AppModInfo
	5,  // 1: http.app.GetAppByIdReply.app:type_name -> http.app.AppInfo
	6,  // 2: http.app.GetAppByIdReply.user:type_name -> http.app.UserInfo
	5,  // 3: http.app.ListAppPageReply.appList:type_name -> http.app.AppInfo
	0,  // 4: http.app.App.CreatApp:input_type -> http.app.CreatAppRequest
	2,  // 5: http.app.App.ModifyApp:input_type -> http.app.ModifyAppRequest
	7,  // 6: http.app.App.GetAppById:input_type -> http.app.GetAppByIdRequest
	9,  // 7: http.app.App.ListAppPage:input_type -> http.app.ListAppPageRequest
	1,  // 8: http.app.App.CreatApp:output_type -> http.app.CreatAppReply
	3,  // 9: http.app.App.ModifyApp:output_type -> http.app.ModifyAppReply
	8,  // 10: http.app.App.GetAppById:output_type -> http.app.GetAppByIdReply
	10, // 11: http.app.App.ListAppPage:output_type -> http.app.ListAppPageReply
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_http_app_app_http_proto_init() }
func file_http_app_app_http_proto_init() {
	if File_http_app_app_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_app_app_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatAppRequest); i {
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
		file_http_app_app_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatAppReply); i {
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
		file_http_app_app_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyAppRequest); i {
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
		file_http_app_app_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyAppReply); i {
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
		file_http_app_app_http_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppModInfo); i {
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
		file_http_app_app_http_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
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
		file_http_app_app_http_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_http_app_app_http_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppByIdRequest); i {
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
		file_http_app_app_http_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppByIdReply); i {
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
		file_http_app_app_http_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppPageRequest); i {
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
		file_http_app_app_http_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppPageReply); i {
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
			RawDescriptor: file_http_app_app_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_http_app_app_http_proto_goTypes,
		DependencyIndexes: file_http_app_app_http_proto_depIdxs,
		MessageInfos:      file_http_app_app_http_proto_msgTypes,
	}.Build()
	File_http_app_app_http_proto = out.File
	file_http_app_app_http_proto_rawDesc = nil
	file_http_app_app_http_proto_goTypes = nil
	file_http_app_app_http_proto_depIdxs = nil
}