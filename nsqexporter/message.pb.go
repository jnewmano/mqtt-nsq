// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package nsqexporter is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	MQTTMessage
*/
package nsqexporter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MQTTMessage struct {
	// message data
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
	Topic     string                     `protobuf:"bytes,2,opt,name=Topic,json=topic" json:"Topic,omitempty"`
	Payload   []byte                     `protobuf:"bytes,3,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
	// meta data
	SourceAddress string `protobuf:"bytes,5,opt,name=SourceAddress,json=sourceAddress" json:"SourceAddress,omitempty"`
	PacketID      uint32 `protobuf:"varint,6,opt,name=PacketID,json=packetID" json:"PacketID,omitempty"`
}

func (m *MQTTMessage) Reset()                    { *m = MQTTMessage{} }
func (m *MQTTMessage) String() string            { return proto.CompactTextString(m) }
func (*MQTTMessage) ProtoMessage()               {}
func (*MQTTMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MQTTMessage) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *MQTTMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *MQTTMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *MQTTMessage) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MQTTMessage) GetPacketID() uint32 {
	if m != nil {
		return m.PacketID
	}
	return 0
}

func init() {
	proto.RegisterType((*MQTTMessage)(nil), "nsqexporter.MQTTMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x40, 0xb3, 0x1a, 0xee, 0x8e, 0x45, 0x9a, 0x8d, 0xc5, 0x86, 0x8a, 0x18, 0x0b, 0x2a, 0x48,
	0xb4, 0x31, 0xb1, 0x32, 0xb1, 0xb1, 0xb8, 0xe4, 0x44, 0x7e, 0x60, 0x81, 0xb9, 0x95, 0x08, 0x37,
	0xeb, 0xce, 0x90, 0x78, 0x9f, 0xe6, 0xdf, 0x19, 0xe1, 0xc4, 0xeb, 0xf6, 0x6d, 0xde, 0xbc, 0x19,
	0x19, 0x0f, 0x40, 0x64, 0x2c, 0xe4, 0xce, 0x23, 0xa3, 0x8a, 0x0e, 0xf4, 0x09, 0x5f, 0x0e, 0x3d,
	0x83, 0x4f, 0x1e, 0x6d, 0xc7, 0xef, 0x63, 0x9d, 0x37, 0x38, 0x14, 0x16, 0x7b, 0x73, 0xb0, 0xc5,
	0x64, 0xd5, 0xe3, 0xbe, 0x70, 0x7c, 0x74, 0x40, 0x05, 0x77, 0x03, 0x10, 0x9b, 0xc1, 0xfd, 0xbf,
	0xe6, 0xd2, 0xcd, 0xb7, 0x90, 0xd1, 0xf6, 0xb5, 0xaa, 0xb6, 0x73, 0x5f, 0x3d, 0xc8, 0xb0, 0xfa,
	0x53, 0xb4, 0x48, 0x45, 0x16, 0xdd, 0x25, 0xb9, 0x45, 0xb4, 0xfd, 0x69, 0x77, 0x3d, 0xee, 0xf3,
	0xc5, 0x28, 0xc3, 0xa5, 0xa7, 0xae, 0x65, 0x50, 0xa1, 0xeb, 0x1a, 0x7d, 0x91, 0x8a, 0x2c, 0x2c,
	0x03, 0xfe, 0x05, 0xa5, 0xe5, 0x7a, 0x67, 0x8e, 0x3d, 0x9a, 0x56, 0x5f, 0xa6, 0x22, 0xbb, 0x2a,
	0xd7, 0x6e, 0x46, 0x75, 0x2b, 0xe3, 0x37, 0x1c, 0x7d, 0x03, 0x4f, 0x6d, 0xeb, 0x81, 0x48, 0x07,
	0xd3, 0x5c, 0x4c, 0xe7, 0x9f, 0x2a, 0x91, 0x9b, 0x9d, 0x69, 0x3e, 0x80, 0x5f, 0x9e, 0xf5, 0x2a,
	0x15, 0x59, 0x5c, 0x6e, 0xdc, 0x89, 0xeb, 0xd5, 0x74, 0xd0, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x1d, 0x13, 0x00, 0x1d, 0x01, 0x00, 0x00,
}
