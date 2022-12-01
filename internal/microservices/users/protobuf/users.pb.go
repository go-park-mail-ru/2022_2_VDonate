// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: internal/microservices/users/protobuf/users.proto

package protobuf

import (
	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/auth/protobuf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username           string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email              string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password           string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Avatar             string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsAuthor           bool   `protobuf:"varint,6,opt,name=is_author,json=isAuthor,proto3" json:"is_author,omitempty"`
	About              string `protobuf:"bytes,7,opt,name=about,proto3" json:"about,omitempty"`
	CountSubscriptions uint64 `protobuf:"varint,8,opt,name=count_subscriptions,json=countSubscriptions,proto3" json:"count_subscriptions,omitempty"`
	CountSubscribers   uint64 `protobuf:"varint,9,opt,name=count_subscribers,json=countSubscribers,proto3" json:"count_subscribers,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetIsAuthor() bool {
	if x != nil {
		return x.IsAuthor
	}
	return false
}

func (x *User) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

func (x *User) GetCountSubscriptions() uint64 {
	if x != nil {
		return x.CountSubscriptions
	}
	return 0
}

func (x *User) GetCountSubscribers() uint64 {
	if x != nil {
		return x.CountSubscribers
	}
	return 0
}

type LessUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *LessUser) Reset() {
	*x = LessUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LessUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LessUser) ProtoMessage() {}

func (x *LessUser) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LessUser.ProtoReflect.Descriptor instead.
func (*LessUser) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{1}
}

func (x *LessUser) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LessUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LessUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UsersArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersArray) Reset() {
	*x = UsersArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersArray) ProtoMessage() {}

func (x *UsersArray) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersArray.ProtoReflect.Descriptor instead.
func (*UsersArray) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{2}
}

func (x *UsersArray) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{3}
}

func (x *UserID) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserAuthorPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthorId uint64 `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *UserAuthorPair) Reset() {
	*x = UserAuthorPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuthorPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuthorPair) ProtoMessage() {}

func (x *UserAuthorPair) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuthorPair.ProtoReflect.Descriptor instead.
func (*UserAuthorPair) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{4}
}

func (x *UserAuthorPair) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserAuthorPair) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type Username struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Username) Reset() {
	*x = Username{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Username) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Username) ProtoMessage() {}

func (x *Username) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Username.ProtoReflect.Descriptor instead.
func (*Username) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{5}
}

func (x *Username) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Keyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *Keyword) Reset() {
	*x = Keyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyword) ProtoMessage() {}

func (x *Keyword) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyword.ProtoReflect.Descriptor instead.
func (*Keyword) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{6}
}

func (x *Keyword) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{7}
}

func (x *Email) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{8}
}

func (x *File) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type PostID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID uint64 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *PostID) Reset() {
	*x = PostID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostID) ProtoMessage() {}

func (x *PostID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_users_protobuf_users_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostID.ProtoReflect.Descriptor instead.
func (*PostID) Descriptor() ([]byte, []int) {
	return file_internal_microservices_users_protobuf_users_proto_rawDescGZIP(), []int{9}
}

func (x *PostID) GetPostID() uint64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

var File_internal_microservices_users_protobuf_users_proto protoreflect.FileDescriptor

var file_internal_microservices_users_protobuf_users_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x22, 0x4e, 0x0a, 0x08, 0x4c, 0x65, 0x73, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x2e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x21, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x0e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x49, 0x64, 0x22, 0x26, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x1d, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x1a,
	0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x32, 0xa1, 0x03, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x10,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f, 0x32,
	0x30, 0x32, 0x32, 0x5f, 0x32, 0x5f, 0x56, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_microservices_users_protobuf_users_proto_rawDescOnce sync.Once
	file_internal_microservices_users_protobuf_users_proto_rawDescData = file_internal_microservices_users_protobuf_users_proto_rawDesc
)

func file_internal_microservices_users_protobuf_users_proto_rawDescGZIP() []byte {
	file_internal_microservices_users_protobuf_users_proto_rawDescOnce.Do(func() {
		file_internal_microservices_users_protobuf_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_microservices_users_protobuf_users_proto_rawDescData)
	})
	return file_internal_microservices_users_protobuf_users_proto_rawDescData
}

var file_internal_microservices_users_protobuf_users_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_microservices_users_protobuf_users_proto_goTypes = []interface{}{
	(*User)(nil),               // 0: user.User
	(*LessUser)(nil),           // 1: user.LessUser
	(*UsersArray)(nil),         // 2: user.UsersArray
	(*UserID)(nil),             // 3: user.UserID
	(*UserAuthorPair)(nil),     // 4: user.UserAuthorPair
	(*Username)(nil),           // 5: user.Username
	(*Keyword)(nil),            // 6: user.Keyword
	(*Email)(nil),              // 7: user.Email
	(*File)(nil),               // 8: user.File
	(*PostID)(nil),             // 9: user.PostID
	(*protobuf.SessionID)(nil), // 10: auth.SessionID
	(*emptypb.Empty)(nil),      // 11: google.protobuf.Empty
}
var file_internal_microservices_users_protobuf_users_proto_depIdxs = []int32{
	0,  // 0: user.UsersArray.users:type_name -> user.User
	0,  // 1: user.Users.Update:input_type -> user.User
	0,  // 2: user.Users.Create:input_type -> user.User
	6,  // 3: user.Users.GetAuthorByUsername:input_type -> user.Keyword
	3,  // 4: user.Users.GetByID:input_type -> user.UserID
	10, // 5: user.Users.GetBySessionID:input_type -> auth.SessionID
	7,  // 6: user.Users.GetByEmail:input_type -> user.Email
	5,  // 7: user.Users.GetByUsername:input_type -> user.Username
	9,  // 8: user.Users.GetUserByPostID:input_type -> user.PostID
	11, // 9: user.Users.GetAllAuthors:input_type -> google.protobuf.Empty
	11, // 10: user.Users.Update:output_type -> google.protobuf.Empty
	3,  // 11: user.Users.Create:output_type -> user.UserID
	2,  // 12: user.Users.GetAuthorByUsername:output_type -> user.UsersArray
	0,  // 13: user.Users.GetByID:output_type -> user.User
	0,  // 14: user.Users.GetBySessionID:output_type -> user.User
	0,  // 15: user.Users.GetByEmail:output_type -> user.User
	0,  // 16: user.Users.GetByUsername:output_type -> user.User
	0,  // 17: user.Users.GetUserByPostID:output_type -> user.User
	2,  // 18: user.Users.GetAllAuthors:output_type -> user.UsersArray
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_internal_microservices_users_protobuf_users_proto_init() }
func file_internal_microservices_users_protobuf_users_proto_init() {
	if File_internal_microservices_users_protobuf_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_microservices_users_protobuf_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LessUser); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersArray); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuthorPair); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Username); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyword); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_internal_microservices_users_protobuf_users_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostID); i {
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
			RawDescriptor: file_internal_microservices_users_protobuf_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_microservices_users_protobuf_users_proto_goTypes,
		DependencyIndexes: file_internal_microservices_users_protobuf_users_proto_depIdxs,
		MessageInfos:      file_internal_microservices_users_protobuf_users_proto_msgTypes,
	}.Build()
	File_internal_microservices_users_protobuf_users_proto = out.File
	file_internal_microservices_users_protobuf_users_proto_rawDesc = nil
	file_internal_microservices_users_protobuf_users_proto_goTypes = nil
	file_internal_microservices_users_protobuf_users_proto_depIdxs = nil
}
