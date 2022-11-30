// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: internal/microservices/post/protobuf/posts.proto

package protobuf

import (
	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID          uint64                 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ContentTemplate string                 `protobuf:"bytes,3,opt,name=ContentTemplate,proto3" json:"ContentTemplate,omitempty"`
	Content         string                 `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Tier            uint64                 `protobuf:"varint,5,opt,name=Tier,proto3" json:"Tier,omitempty"`
	IsAllowed       bool                   `protobuf:"varint,6,opt,name=IsAllowed,proto3" json:"IsAllowed,omitempty"`
	DateCreated     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=DateCreated,proto3" json:"DateCreated,omitempty"`
	Tags            []string               `protobuf:"bytes,8,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Author          *protobuf.LessUser     `protobuf:"bytes,9,opt,name=Author,proto3" json:"Author,omitempty"`
	LikesNum        uint64                 `protobuf:"varint,10,opt,name=LikesNum,proto3" json:"LikesNum,omitempty"`
	IsLiked         bool                   `protobuf:"varint,11,opt,name=IsLiked,proto3" json:"IsLiked,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Post) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Post) GetContentTemplate() string {
	if x != nil {
		return x.ContentTemplate
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetTier() uint64 {
	if x != nil {
		return x.Tier
	}
	return 0
}

func (x *Post) GetIsAllowed() bool {
	if x != nil {
		return x.IsAllowed
	}
	return false
}

func (x *Post) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

func (x *Post) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Post) GetAuthor() *protobuf.LessUser {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Post) GetLikesNum() uint64 {
	if x != nil {
		return x.LikesNum
	}
	return 0
}

func (x *Post) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

type PostArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PostArray) Reset() {
	*x = PostArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostArray) ProtoMessage() {}

func (x *PostArray) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostArray.ProtoReflect.Descriptor instead.
func (*PostArray) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{1}
}

func (x *PostArray) GetPosts() []*Post {
	if x != nil {
		return x.Posts
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
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostID) ProtoMessage() {}

func (x *PostID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[2]
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
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{2}
}

func (x *PostID) GetPostID() uint64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type PostUserIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID uint64 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
	UserID uint64 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *PostUserIDs) Reset() {
	*x = PostUserIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostUserIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostUserIDs) ProtoMessage() {}

func (x *PostUserIDs) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostUserIDs.ProtoReflect.Descriptor instead.
func (*PostUserIDs) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{3}
}

func (x *PostUserIDs) GetPostID() uint64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

func (x *PostUserIDs) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	PostID uint64 `protobuf:"varint,2,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *Like) Reset() {
	*x = Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Like.ProtoReflect.Descriptor instead.
func (*Like) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{4}
}

func (x *Like) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Like) GetPostID() uint64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type Likes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Likes []*Like `protobuf:"bytes,1,rep,name=likes,proto3" json:"likes,omitempty"`
}

func (x *Likes) Reset() {
	*x = Likes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Likes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Likes) ProtoMessage() {}

func (x *Likes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Likes.ProtoReflect.Descriptor instead.
func (*Likes) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{5}
}

func (x *Likes) GetLikes() []*Like {
	if x != nil {
		return x.Likes
	}
	return nil
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TagName string `protobuf:"bytes,2,opt,name=tagName,proto3" json:"tagName,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{6}
}

func (x *Tag) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

type TagDep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID uint64 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
	TagID  uint64 `protobuf:"varint,2,opt,name=tagID,proto3" json:"tagID,omitempty"`
}

func (x *TagDep) Reset() {
	*x = TagDep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagDep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagDep) ProtoMessage() {}

func (x *TagDep) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagDep.ProtoReflect.Descriptor instead.
func (*TagDep) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{7}
}

func (x *TagDep) GetPostID() uint64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

func (x *TagDep) GetTagID() uint64 {
	if x != nil {
		return x.TagID
	}
	return 0
}

type TagDeps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagDeps []*TagDep `protobuf:"bytes,1,rep,name=tagDeps,proto3" json:"tagDeps,omitempty"`
}

func (x *TagDeps) Reset() {
	*x = TagDeps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagDeps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagDeps) ProtoMessage() {}

func (x *TagDeps) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagDeps.ProtoReflect.Descriptor instead.
func (*TagDeps) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{8}
}

func (x *TagDeps) GetTagDeps() []*TagDep {
	if x != nil {
		return x.TagDeps
	}
	return nil
}

type TagName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagName string `protobuf:"bytes,1,opt,name=tagName,proto3" json:"tagName,omitempty"`
}

func (x *TagName) Reset() {
	*x = TagName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagName) ProtoMessage() {}

func (x *TagName) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagName.ProtoReflect.Descriptor instead.
func (*TagName) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{9}
}

func (x *TagName) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

type TagID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagID uint64 `protobuf:"varint,1,opt,name=tagID,proto3" json:"tagID,omitempty"`
}

func (x *TagID) Reset() {
	*x = TagID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagID) ProtoMessage() {}

func (x *TagID) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_post_protobuf_posts_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagID.ProtoReflect.Descriptor instead.
func (*TagID) Descriptor() ([]byte, []int) {
	return file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP(), []int{10}
}

func (x *TagID) GetTagID() uint64 {
	if x != nil {
		return x.TagID
	}
	return 0
}

var File_internal_microservices_post_protobuf_posts_proto protoreflect.FileDescriptor

var file_internal_microservices_post_protobuf_posts_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x04, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x54,
	0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x61, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x22, 0x2e, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x21,
	0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x22, 0x20, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x44, 0x22, 0x3d, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x36, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x2a, 0x0a, 0x05, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52,
	0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x06, 0x54, 0x61, 0x67, 0x44, 0x65,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x44, 0x22,
	0x32, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x44, 0x65, 0x70, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x61,
	0x67, 0x44, 0x65, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x44, 0x65, 0x70, 0x52, 0x07, 0x74, 0x61, 0x67, 0x44,
	0x65, 0x70, 0x73, 0x22, 0x23, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x61, 0x67, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x44, 0x32, 0xb2, 0x06, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x30, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x12, 0x29, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x44, 0x1a, 0x0b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x12, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x12, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x32,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x42, 0x79, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65,
	0x12, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e,
	0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e,
	0x54, 0x61, 0x67, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x49,
	0x44, 0x1a, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x12, 0x2a, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0a, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x12, 0x35, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x54, 0x61, 0x67, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x54, 0x61, 0x67, 0x44, 0x65, 0x70, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x33, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x44, 0x65, 0x70, 0x73, 0x42, 0x79,
	0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x61,
	0x67, 0x44, 0x65, 0x70, 0x73, 0x12, 0x35, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x70, 0x54, 0x61, 0x67, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x54, 0x61,
	0x67, 0x44, 0x65, 0x70, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x50, 0x5a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61,
	0x72, 0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f, 0x32, 0x30, 0x32, 0x32, 0x5f,
	0x32, 0x5f, 0x56, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_microservices_post_protobuf_posts_proto_rawDescOnce sync.Once
	file_internal_microservices_post_protobuf_posts_proto_rawDescData = file_internal_microservices_post_protobuf_posts_proto_rawDesc
)

func file_internal_microservices_post_protobuf_posts_proto_rawDescGZIP() []byte {
	file_internal_microservices_post_protobuf_posts_proto_rawDescOnce.Do(func() {
		file_internal_microservices_post_protobuf_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_microservices_post_protobuf_posts_proto_rawDescData)
	})
	return file_internal_microservices_post_protobuf_posts_proto_rawDescData
}

var file_internal_microservices_post_protobuf_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_microservices_post_protobuf_posts_proto_goTypes = []interface{}{
	(*Post)(nil),                  // 0: posts.Post
	(*PostArray)(nil),             // 1: posts.PostArray
	(*PostID)(nil),                // 2: posts.PostID
	(*PostUserIDs)(nil),           // 3: posts.PostUserIDs
	(*Like)(nil),                  // 4: posts.Like
	(*Likes)(nil),                 // 5: posts.Likes
	(*Tag)(nil),                   // 6: posts.Tag
	(*TagDep)(nil),                // 7: posts.TagDep
	(*TagDeps)(nil),               // 8: posts.TagDeps
	(*TagName)(nil),               // 9: posts.TagName
	(*TagID)(nil),                 // 10: posts.TagID
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*protobuf.LessUser)(nil),     // 12: user.LessUser
	(*protobuf.UserID)(nil),       // 13: user.UserID
	(*emptypb.Empty)(nil),         // 14: google.protobuf.Empty
}
var file_internal_microservices_post_protobuf_posts_proto_depIdxs = []int32{
	11, // 0: posts.Post.DateCreated:type_name -> google.protobuf.Timestamp
	12, // 1: posts.Post.Author:type_name -> user.LessUser
	0,  // 2: posts.PostArray.posts:type_name -> posts.Post
	4,  // 3: posts.Likes.likes:type_name -> posts.Like
	7,  // 4: posts.TagDeps.tagDeps:type_name -> posts.TagDep
	13, // 5: posts.Posts.GetAllByUserID:input_type -> user.UserID
	2,  // 6: posts.Posts.GetPostByID:input_type -> posts.PostID
	0,  // 7: posts.Posts.Create:input_type -> posts.Post
	0,  // 8: posts.Posts.Update:input_type -> posts.Post
	2,  // 9: posts.Posts.DeleteByID:input_type -> posts.PostID
	13, // 10: posts.Posts.GetPostsBySubscriptions:input_type -> user.UserID
	3,  // 11: posts.Posts.GetLikeByUserAndPostID:input_type -> posts.PostUserIDs
	2,  // 12: posts.Posts.GetAllLikesByPostID:input_type -> posts.PostID
	3,  // 13: posts.Posts.CreateLike:input_type -> posts.PostUserIDs
	3,  // 14: posts.Posts.DeleteLikeByID:input_type -> posts.PostUserIDs
	9,  // 15: posts.Posts.CreateTag:input_type -> posts.TagName
	10, // 16: posts.Posts.GetTagById:input_type -> posts.TagID
	9,  // 17: posts.Posts.GetTagByName:input_type -> posts.TagName
	7,  // 18: posts.Posts.CreateDepTag:input_type -> posts.TagDep
	2,  // 19: posts.Posts.GetTagDepsByPostId:input_type -> posts.PostID
	7,  // 20: posts.Posts.DeleteDepTag:input_type -> posts.TagDep
	1,  // 21: posts.Posts.GetAllByUserID:output_type -> posts.PostArray
	0,  // 22: posts.Posts.GetPostByID:output_type -> posts.Post
	2,  // 23: posts.Posts.Create:output_type -> posts.PostID
	14, // 24: posts.Posts.Update:output_type -> google.protobuf.Empty
	14, // 25: posts.Posts.DeleteByID:output_type -> google.protobuf.Empty
	1,  // 26: posts.Posts.GetPostsBySubscriptions:output_type -> posts.PostArray
	4,  // 27: posts.Posts.GetLikeByUserAndPostID:output_type -> posts.Like
	5,  // 28: posts.Posts.GetAllLikesByPostID:output_type -> posts.Likes
	14, // 29: posts.Posts.CreateLike:output_type -> google.protobuf.Empty
	14, // 30: posts.Posts.DeleteLikeByID:output_type -> google.protobuf.Empty
	10, // 31: posts.Posts.CreateTag:output_type -> posts.TagID
	6,  // 32: posts.Posts.GetTagById:output_type -> posts.Tag
	6,  // 33: posts.Posts.GetTagByName:output_type -> posts.Tag
	14, // 34: posts.Posts.CreateDepTag:output_type -> google.protobuf.Empty
	8,  // 35: posts.Posts.GetTagDepsByPostId:output_type -> posts.TagDeps
	14, // 36: posts.Posts.DeleteDepTag:output_type -> google.protobuf.Empty
	21, // [21:37] is the sub-list for method output_type
	5,  // [5:21] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_internal_microservices_post_protobuf_posts_proto_init() }
func file_internal_microservices_post_protobuf_posts_proto_init() {
	if File_internal_microservices_post_protobuf_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostArray); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostUserIDs); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Like); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Likes); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagDep); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagDeps); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagName); i {
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
		file_internal_microservices_post_protobuf_posts_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagID); i {
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
			RawDescriptor: file_internal_microservices_post_protobuf_posts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_microservices_post_protobuf_posts_proto_goTypes,
		DependencyIndexes: file_internal_microservices_post_protobuf_posts_proto_depIdxs,
		MessageInfos:      file_internal_microservices_post_protobuf_posts_proto_msgTypes,
	}.Build()
	File_internal_microservices_post_protobuf_posts_proto = out.File
	file_internal_microservices_post_protobuf_posts_proto_rawDesc = nil
	file_internal_microservices_post_protobuf_posts_proto_goTypes = nil
	file_internal_microservices_post_protobuf_posts_proto_depIdxs = nil
}
