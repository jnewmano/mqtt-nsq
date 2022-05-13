// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: message.proto

package nsqexporter

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

type MQTTMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message data
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Topic     string                 `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Payload   []byte                 `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	// meta data
	SourceAddress string `protobuf:"bytes,5,opt,name=SourceAddress,proto3" json:"SourceAddress,omitempty"`
	PacketID      uint32 `protobuf:"varint,6,opt,name=PacketID,proto3" json:"PacketID,omitempty"`
}

func (x *MQTTMessage) Reset() {
	*x = MQTTMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQTTMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQTTMessage) ProtoMessage() {}

func (x *MQTTMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQTTMessage.ProtoReflect.Descriptor instead.
func (*MQTTMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *MQTTMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *MQTTMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *MQTTMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MQTTMessage) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

func (x *MQTTMessage) GetPacketID() uint32 {
	if x != nil {
		return x.PacketID
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6e, 0x73, 0x71, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01,
	0x0a, 0x0b, 0x4d, 0x51, 0x54, 0x54, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6e, 0x65, 0x77, 0x6d, 0x61, 0x6e, 0x6f,
	0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2d, 0x6e, 0x73, 0x71, 0x2f, 0x6e, 0x73, 0x71, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x3b, 0x6e, 0x73, 0x71, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_proto_goTypes = []interface{}{
	(*MQTTMessage)(nil),           // 0: nsqexporter.MQTTMessage
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_message_proto_depIdxs = []int32{
	1, // 0: nsqexporter.MQTTMessage.Timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQTTMessage); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
