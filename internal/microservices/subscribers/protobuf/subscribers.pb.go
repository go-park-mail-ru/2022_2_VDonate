// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: internal/microservices/subscribers/protobuf/subscribers.proto

package protobuf

import (
	protobuf "github.com/go-park-mail-ru/2022_2_VDonate/internal/microservices/users/protobuf"
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

type Subscriber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorID             uint64 `protobuf:"varint,1,opt,name=AuthorID,proto3" json:"AuthorID,omitempty"`
	SubscriberID         uint64 `protobuf:"varint,2,opt,name=SubscriberID,proto3" json:"SubscriberID,omitempty"`
	AuthorSubscriptionID uint64 `protobuf:"varint,3,opt,name=AuthorSubscriptionID,proto3" json:"AuthorSubscriptionID,omitempty"`
}

func (x *Subscriber) Reset() {
	*x = Subscriber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_subscribers_protobuf_subscribers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscriber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscriber) ProtoMessage() {}

func (x *Subscriber) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_subscribers_protobuf_subscribers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscriber.ProtoReflect.Descriptor instead.
func (*Subscriber) Descriptor() ([]byte, []int) {
	return file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescGZIP(), []int{0}
}

func (x *Subscriber) GetAuthorID() uint64 {
	if x != nil {
		return x.AuthorID
	}
	return 0
}

func (x *Subscriber) GetSubscriberID() uint64 {
	if x != nil {
		return x.SubscriberID
	}
	return 0
}

func (x *Subscriber) GetAuthorSubscriptionID() uint64 {
	if x != nil {
		return x.AuthorSubscriptionID
	}
	return 0
}

var File_internal_microservices_subscribers_protobuf_subscribers_proto protoreflect.FileDescriptor

var file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x31, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a,
	0x0a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x32,
	0xb7, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x2d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a,
	0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12, 0x3c,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0b,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x50, 0x61, 0x69,
	0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d,
	0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f, 0x32, 0x30, 0x32, 0x32, 0x5f, 0x32, 0x5f, 0x56,
	0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescOnce sync.Once
	file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescData = file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDesc
)

func file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescGZIP() []byte {
	file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescOnce.Do(func() {
		file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescData)
	})
	return file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDescData
}

var file_internal_microservices_subscribers_protobuf_subscribers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_microservices_subscribers_protobuf_subscribers_proto_goTypes = []interface{}{
	(*Subscriber)(nil),              // 0: subscribers.Subscriber
	(*protobuf.UserID)(nil),         // 1: user.UserID
	(*protobuf.UserAuthorPair)(nil), // 2: user.UserAuthorPair
	(*protobuf.UserIDs)(nil),        // 3: user.UserIDs
	(*emptypb.Empty)(nil),           // 4: google.protobuf.Empty
}
var file_internal_microservices_subscribers_protobuf_subscribers_proto_depIdxs = []int32{
	1, // 0: subscribers.Subscribers.GetSubscribers:input_type -> user.UserID
	0, // 1: subscribers.Subscribers.Subscribe:input_type -> subscribers.Subscriber
	2, // 2: subscribers.Subscribers.Unsubscribe:input_type -> user.UserAuthorPair
	3, // 3: subscribers.Subscribers.GetSubscribers:output_type -> user.UserIDs
	4, // 4: subscribers.Subscribers.Subscribe:output_type -> google.protobuf.Empty
	4, // 5: subscribers.Subscribers.Unsubscribe:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_microservices_subscribers_protobuf_subscribers_proto_init() }
func file_internal_microservices_subscribers_protobuf_subscribers_proto_init() {
	if File_internal_microservices_subscribers_protobuf_subscribers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_microservices_subscribers_protobuf_subscribers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscriber); i {
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
			RawDescriptor: file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_microservices_subscribers_protobuf_subscribers_proto_goTypes,
		DependencyIndexes: file_internal_microservices_subscribers_protobuf_subscribers_proto_depIdxs,
		MessageInfos:      file_internal_microservices_subscribers_protobuf_subscribers_proto_msgTypes,
	}.Build()
	File_internal_microservices_subscribers_protobuf_subscribers_proto = out.File
	file_internal_microservices_subscribers_protobuf_subscribers_proto_rawDesc = nil
	file_internal_microservices_subscribers_protobuf_subscribers_proto_goTypes = nil
	file_internal_microservices_subscribers_protobuf_subscribers_proto_depIdxs = nil
}
