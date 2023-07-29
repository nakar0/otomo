// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: message_with_otomo.proto

package grpcgen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Role int32

const (
	Role_User  Role = 0
	Role_Otomo Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "User",
		1: "Otomo",
	}
	Role_value = map[string]int32{
		"User":  0,
		"Otomo": 1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_message_with_otomo_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_message_with_otomo_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_message_with_otomo_proto_rawDescGZIP(), []int{0}
}

type MessageWithOtomo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Sender   Role                   `protobuf:"varint,3,opt,name=sender,proto3,enum=Role" json:"sender,omitempty"`
	Reciever Role                   `protobuf:"varint,4,opt,name=reciever,proto3,enum=Role" json:"reciever,omitempty"`
	Text     string                 `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	SentAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *MessageWithOtomo) Reset() {
	*x = MessageWithOtomo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_with_otomo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOtomo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOtomo) ProtoMessage() {}

func (x *MessageWithOtomo) ProtoReflect() protoreflect.Message {
	mi := &file_message_with_otomo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOtomo.ProtoReflect.Descriptor instead.
func (*MessageWithOtomo) Descriptor() ([]byte, []int) {
	return file_message_with_otomo_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithOtomo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageWithOtomo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MessageWithOtomo) GetSender() Role {
	if x != nil {
		return x.Sender
	}
	return Role_User
}

func (x *MessageWithOtomo) GetReciever() Role {
	if x != nil {
		return x.Reciever
	}
	return Role_User
}

func (x *MessageWithOtomo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MessageWithOtomo) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

var File_message_with_otomo_proto protoreflect.FileDescriptor

var file_message_with_otomo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6f,
	0x74, 0x6f, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x10,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x74, 0x6f, 0x6d, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x2a, 0x1b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x74, 0x6f, 0x6d, 0x6f, 0x10,
	0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_with_otomo_proto_rawDescOnce sync.Once
	file_message_with_otomo_proto_rawDescData = file_message_with_otomo_proto_rawDesc
)

func file_message_with_otomo_proto_rawDescGZIP() []byte {
	file_message_with_otomo_proto_rawDescOnce.Do(func() {
		file_message_with_otomo_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_with_otomo_proto_rawDescData)
	})
	return file_message_with_otomo_proto_rawDescData
}

var file_message_with_otomo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_with_otomo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_with_otomo_proto_goTypes = []interface{}{
	(Role)(0),                     // 0: Role
	(*MessageWithOtomo)(nil),      // 1: MessageWithOtomo
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_message_with_otomo_proto_depIdxs = []int32{
	0, // 0: MessageWithOtomo.sender:type_name -> Role
	0, // 1: MessageWithOtomo.reciever:type_name -> Role
	2, // 2: MessageWithOtomo.sent_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_with_otomo_proto_init() }
func file_message_with_otomo_proto_init() {
	if File_message_with_otomo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_with_otomo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOtomo); i {
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
			RawDescriptor: file_message_with_otomo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_with_otomo_proto_goTypes,
		DependencyIndexes: file_message_with_otomo_proto_depIdxs,
		EnumInfos:         file_message_with_otomo_proto_enumTypes,
		MessageInfos:      file_message_with_otomo_proto_msgTypes,
	}.Build()
	File_message_with_otomo_proto = out.File
	file_message_with_otomo_proto_rawDesc = nil
	file_message_with_otomo_proto_goTypes = nil
	file_message_with_otomo_proto_depIdxs = nil
}