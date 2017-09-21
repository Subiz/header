// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto

It has these top-level messages:
	Empty
	Error
	Context
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import auth "bitbucket.org/subiz/servicespec/proto/auth"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Error struct {
	Ctx              *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Error) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

type Context struct {
	EventId          *string          `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	State            []byte           `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Node             *string          `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	ReplyTopic       *string          `protobuf:"bytes,4,opt,name=reply_topic,json=replyTopic" json:"reply_topic,omitempty"`
	ReplyPartition   *int32           `protobuf:"varint,5,opt,name=reply_partition,json=replyPartition" json:"reply_partition,omitempty"`
	Credential       *auth.Credential `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Tracing          []byte           `protobuf:"bytes,7,opt,name=tracing" json:"tracing,omitempty"`
	ReplyKey         *string          `protobuf:"bytes,8,opt,name=reply_key,json=replyKey" json:"reply_key,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Context) Reset()                    { *m = Context{} }
func (m *Context) String() string            { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()               {}
func (*Context) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Context) GetEventId() string {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return ""
}

func (m *Context) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Context) GetNode() string {
	if m != nil && m.Node != nil {
		return *m.Node
	}
	return ""
}

func (m *Context) GetReplyTopic() string {
	if m != nil && m.ReplyTopic != nil {
		return *m.ReplyTopic
	}
	return ""
}

func (m *Context) GetReplyPartition() int32 {
	if m != nil && m.ReplyPartition != nil {
		return *m.ReplyPartition
	}
	return 0
}

func (m *Context) GetCredential() *auth.Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *Context) GetTracing() []byte {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *Context) GetReplyKey() string {
	if m != nil && m.ReplyKey != nil {
		return *m.ReplyKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*Error)(nil), "common.Error")
	proto.RegisterType((*Context)(nil), "common.Context")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0x53, 0xa0, 0x14, 0x86, 0x7f, 0xfe, 0x98, 0xd5, 0xc3, 0xaa, 0x07, 0x91, 0x8b, 0xc4,
	0x43, 0x6b, 0x4c, 0xfc, 0x04, 0x84, 0x83, 0xf1, 0x62, 0x1a, 0xef, 0xa4, 0xdd, 0x4e, 0x70, 0x03,
	0xdd, 0xdd, 0x6c, 0xa7, 0x84, 0x7a, 0xf3, 0x9b, 0x1b, 0x66, 0xc1, 0x78, 0xd9, 0x9d, 0xf7, 0x9b,
	0x37, 0x99, 0x79, 0x70, 0xa9, 0x6c, 0x5d, 0x5b, 0x93, 0x85, 0x2f, 0x75, 0xde, 0x92, 0x15, 0xc3,
	0xa0, 0x6e, 0x5e, 0x4a, 0x4d, 0x65, 0xab, 0xb6, 0x48, 0xa9, 0xf5, 0x9b, 0xac, 0x69, 0x4b, 0xfd,
	0x95, 0x35, 0xe8, 0xf7, 0x5a, 0x61, 0xe3, 0x50, 0x65, 0x6c, 0xcf, 0x8a, 0x96, 0x3e, 0xf9, 0x09,
	0xe3, 0xf3, 0x04, 0xe2, 0x55, 0xed, 0xa8, 0x9b, 0x3f, 0x42, 0xbc, 0xf2, 0xde, 0x7a, 0x71, 0x0f,
	0x7d, 0x45, 0x07, 0x19, 0xcd, 0xa2, 0xc5, 0xe4, 0x79, 0x9a, 0x9e, 0x96, 0x2d, 0xad, 0x21, 0x3c,
	0x50, 0x7e, 0xec, 0xcd, 0xbf, 0x7b, 0x90, 0x9c, 0x80, 0xb8, 0x86, 0x11, 0xee, 0xd1, 0xd0, 0x5a,
	0x57, 0x3c, 0x33, 0xce, 0x13, 0xd6, 0xaf, 0x95, 0xb8, 0x82, 0xb8, 0xa1, 0x82, 0x50, 0xf6, 0x66,
	0xd1, 0xe2, 0x5f, 0x1e, 0x84, 0x10, 0x30, 0x30, 0xb6, 0x42, 0xd9, 0x67, 0x33, 0xd7, 0xe2, 0x0e,
	0x26, 0x1e, 0xdd, 0xae, 0x5b, 0x93, 0x75, 0x5a, 0xc9, 0x01, 0xb7, 0x80, 0xd1, 0xc7, 0x91, 0x88,
	0x07, 0x98, 0x06, 0x83, 0x2b, 0x3c, 0x69, 0xd2, 0xd6, 0xc8, 0x78, 0x16, 0x2d, 0xe2, 0xfc, 0x3f,
	0xe3, 0xf7, 0x33, 0x15, 0x4f, 0x00, 0xca, 0x63, 0x85, 0x86, 0x74, 0xb1, 0x93, 0x43, 0x0e, 0x71,
	0x91, 0x72, 0xe0, 0xe5, 0x2f, 0xcf, 0xff, 0x78, 0x84, 0x84, 0x84, 0x7c, 0xa1, 0xb4, 0xd9, 0xc8,
	0x84, 0xef, 0x3c, 0x4b, 0x71, 0x0b, 0xe3, 0xb0, 0x74, 0x8b, 0x9d, 0x1c, 0xf1, 0x4d, 0x23, 0x06,
	0x6f, 0xd8, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xda, 0x07, 0xf4, 0x2d, 0x8d, 0x01, 0x00, 0x00,
}
