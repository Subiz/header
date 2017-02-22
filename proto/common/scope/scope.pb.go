// Code generated by protoc-gen-go.
// source: scope.proto
// DO NOT EDIT!

/*
Package scope is a generated protocol buffer package.

It is generated from these files:
	scope.proto

It has these top-level messages:
	Scope
*/
package scope

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Scope struct {
	ChannelManage  bool `protobuf:"varint,1,opt,name=channel_manage,json=channelManage" json:"channel_manage,omitempty"`
	ClientRead     bool `protobuf:"varint,2,opt,name=client_read,json=clientRead" json:"client_read,omitempty"`
	ClientManage   bool `protobuf:"varint,3,opt,name=client_manage,json=clientManage" json:"client_manage,omitempty"`
	ClientValidate bool `protobuf:"varint,4,opt,name=client_validate,json=clientValidate" json:"client_validate,omitempty"`
	ClientDelete   bool `protobuf:"varint,5,opt,name=client_delete,json=clientDelete" json:"client_delete,omitempty"`
	ClientVerify   bool `protobuf:"varint,6,opt,name=client_verify,json=clientVerify" json:"client_verify,omitempty"`
}

func (m *Scope) Reset()                    { *m = Scope{} }
func (m *Scope) String() string            { return proto.CompactTextString(m) }
func (*Scope) ProtoMessage()               {}
func (*Scope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Scope) GetChannelManage() bool {
	if m != nil {
		return m.ChannelManage
	}
	return false
}

func (m *Scope) GetClientRead() bool {
	if m != nil {
		return m.ClientRead
	}
	return false
}

func (m *Scope) GetClientManage() bool {
	if m != nil {
		return m.ClientManage
	}
	return false
}

func (m *Scope) GetClientValidate() bool {
	if m != nil {
		return m.ClientValidate
	}
	return false
}

func (m *Scope) GetClientDelete() bool {
	if m != nil {
		return m.ClientDelete
	}
	return false
}

func (m *Scope) GetClientVerify() bool {
	if m != nil {
		return m.ClientVerify
	}
	return false
}

func init() {
	proto.RegisterType((*Scope)(nil), "scope.Scope")
}

func init() { proto.RegisterFile("scope.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xcf, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0xc6, 0x71, 0xaa, 0xb6, 0xc8, 0xd5, 0x56, 0xc8, 0x94, 0x4d, 0x51, 0x44, 0x27, 0x17, 0x5f,
	0xc1, 0xd5, 0xa5, 0x42, 0xd7, 0x72, 0x36, 0xa7, 0x16, 0x62, 0x5a, 0x6a, 0x28, 0xf8, 0xc4, 0xbe,
	0x86, 0x78, 0xbd, 0x21, 0x5b, 0xf2, 0xe3, 0xcf, 0x07, 0x07, 0xe9, 0xbb, 0x6e, 0x3b, 0x3a, 0x76,
	0x7d, 0xeb, 0x5b, 0x15, 0xf3, 0x67, 0xf3, 0x8d, 0x20, 0xbe, 0xfe, 0x5f, 0x6a, 0x07, 0x79, 0xfd,
	0x44, 0xe7, 0xc8, 0x56, 0x2f, 0x74, 0xf8, 0x20, 0x1d, 0xad, 0xa3, 0xc3, 0xbc, 0xc8, 0x44, 0x2f,
	0x8c, 0x6a, 0x05, 0x69, 0x6d, 0x1b, 0x72, 0xbe, 0xea, 0x09, 0x8d, 0x9e, 0x70, 0x03, 0x23, 0x15,
	0x84, 0x46, 0x6d, 0x21, 0x93, 0x40, 0x66, 0xa6, 0x9c, 0x2c, 0x46, 0x94, 0x95, 0x3d, 0x2c, 0x25,
	0x1a, 0xd0, 0x36, 0x06, 0x3d, 0xe9, 0x19, 0x67, 0xf9, 0xc8, 0xa5, 0x68, 0xb0, 0x66, 0xc8, 0x92,
	0x27, 0x1d, 0x87, 0x6b, 0x67, 0xb6, 0x20, 0x1a, 0xa8, 0x6f, 0xee, 0x1f, 0x9d, 0x84, 0x51, 0xc9,
	0x76, 0x4b, 0xf8, 0xee, 0xd3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xee, 0x31, 0x92, 0x98, 0x06, 0x01,
	0x00, 0x00,
}
