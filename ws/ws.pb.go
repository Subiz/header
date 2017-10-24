// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ws/ws.proto

/*
Package ws is a generated protocol buffer package.

It is generated from these files:
	ws/ws.proto

It has these top-level messages:
	Subscribe
	ListRequest
	Payload
*/
package ws

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event int32

const (
	Event_WsSend Event = 0
)

var Event_name = map[int32]string{
	0: "WsSend",
}
var Event_value = map[string]int32{
	"WsSend": 0,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Subscribe struct {
	// common.Context ctx = 1;
	// optional string connection_id = 2;
	Events []string `protobuf:"bytes,3,rep,name=events" json:"events,omitempty"`
	// optional: only use this in widget; otherwise must provide credential through access token
	MaskId       string `protobuf:"bytes,4,opt,name=mask_id,json=maskId" json:"mask_id,omitempty"`
	ConnectionId string `protobuf:"bytes,5,opt,name=connection_id,json=connectionId" json:"connection_id,omitempty"`
}

func (m *Subscribe) Reset()                    { *m = Subscribe{} }
func (m *Subscribe) String() string            { return proto.CompactTextString(m) }
func (*Subscribe) ProtoMessage()               {}
func (*Subscribe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Subscribe) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Subscribe) GetMaskId() string {
	if m != nil {
		return m.MaskId
	}
	return ""
}

func (m *Subscribe) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type ListRequest struct {
	Ctx          *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	ConnectionId string          `protobuf:"bytes,2,opt,name=connection_id,json=connectionId" json:"connection_id,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ListRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type Payload struct {
	Id      int64  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Payload) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Payload) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Subscribe)(nil), "ws.Subscribe")
	proto.RegisterType((*ListRequest)(nil), "ws.ListRequest")
	proto.RegisterType((*Payload)(nil), "ws.Payload")
	proto.RegisterEnum("ws.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("ws/ws.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x62, 0x13, 0x3a, 0xf1, 0x8b, 0x15, 0x34, 0x78, 0x8a, 0xf5, 0x12, 0x3c, 0x24,
	0x60, 0x7f, 0x82, 0x78, 0x28, 0x78, 0x90, 0x14, 0xf1, 0x28, 0xfb, 0x31, 0xd4, 0xa5, 0x66, 0x57,
	0x33, 0x13, 0x53, 0xfd, 0xf5, 0x92, 0x1a, 0xf1, 0x60, 0x4f, 0xcb, 0xb3, 0x0f, 0xf3, 0xf2, 0xce,
	0x40, 0xda, 0x53, 0xd5, 0x53, 0xf9, 0xd6, 0x7a, 0xf6, 0x22, 0xec, 0xe9, 0xa2, 0x54, 0x96, 0x55,
	0xa7, 0xd7, 0xc8, 0xa5, 0x6f, 0x57, 0x15, 0x75, 0xca, 0x7e, 0x55, 0x2f, 0x28, 0x0d, 0xb6, 0x95,
	0xf6, 0x4d, 0xe3, 0xdd, 0xf8, 0xfc, 0xcc, 0xcc, 0x24, 0x4c, 0x97, 0x9d, 0x22, 0xdd, 0x5a, 0x85,
	0xe2, 0x0c, 0x62, 0xfc, 0x40, 0xc7, 0x94, 0x45, 0x79, 0x54, 0x4c, 0xeb, 0x91, 0xc4, 0x39, 0x24,
	0x8d, 0xa4, 0xf5, 0xb3, 0x35, 0xd9, 0x7e, 0x1e, 0x0c, 0x62, 0xc0, 0x85, 0x11, 0x57, 0x70, 0xa8,
	0xbd, 0x73, 0xa8, 0xd9, 0x7a, 0x37, 0xe8, 0xc9, 0x56, 0x1f, 0xfc, 0x7d, 0x2e, 0xcc, 0xec, 0x11,
	0xd2, 0x7b, 0x4b, 0x5c, 0xe3, 0x7b, 0x87, 0xc4, 0xe2, 0x12, 0x22, 0xcd, 0x9b, 0x2c, 0xc8, 0x83,
	0x22, 0xbd, 0x39, 0x2e, 0xc7, 0x36, 0xb7, 0xde, 0x31, 0x6e, 0xb8, 0x1e, 0xdc, 0xff, 0xd8, 0x70,
	0x47, 0xec, 0x1c, 0x92, 0x07, 0xf9, 0xf9, 0xea, 0xa5, 0x11, 0x47, 0x10, 0x5a, 0x93, 0x45, 0x79,
	0x50, 0x44, 0x75, 0x68, 0x8d, 0xc8, 0x20, 0x69, 0x90, 0x48, 0xae, 0x70, 0xec, 0xfb, 0x8b, 0xd7,
	0xa7, 0x30, 0xb9, 0x1b, 0x76, 0x12, 0x00, 0xf1, 0x13, 0x2d, 0xd1, 0x99, 0x93, 0x3d, 0x15, 0x6f,
	0x4f, 0x31, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x38, 0xef, 0x16, 0xcf, 0x4d, 0x01, 0x00, 0x00,
}
