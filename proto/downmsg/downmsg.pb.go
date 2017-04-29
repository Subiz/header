// Code generated by protoc-gen-go.
// source: downmsg/downmsg.proto
// DO NOT EDIT!

/*
Package downmsg is a generated protocol buffer package.

It is generated from these files:
	downmsg/downmsg.proto

It has these top-level messages:
	SendRequest
	Id
	Empty
	Registration
*/
package downmsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SendRequest struct {
	UserId  string `protobuf:"bytes,1,opt,name=UserId,json=userId" json:"UserId,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content,json=content" json:"Content,omitempty"`
}

func (m *SendRequest) Reset()                    { *m = SendRequest{} }
func (m *SendRequest) String() string            { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()               {}
func (*SendRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SendRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Registration struct {
	UserId   string `protobuf:"bytes,1,opt,name=UserId,json=userId" json:"UserId,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=ClientId,json=clientId" json:"ClientId,omitempty"`
}

func (m *Registration) Reset()                    { *m = Registration{} }
func (m *Registration) String() string            { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()               {}
func (*Registration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Registration) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Registration) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "downmsg.SendRequest")
	proto.RegisterType((*Id)(nil), "downmsg.Id")
	proto.RegisterType((*Empty)(nil), "downmsg.Empty")
	proto.RegisterType((*Registration)(nil), "downmsg.Registration")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DownstreamMsg service

type DownstreamMsgClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*Empty, error)
	// must call when client webhook is updated
	OnClientUpdated(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	// must be call when registration change
	OnRevoke(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error)
	OnRegister(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error)
}

type downstreamMsgClient struct {
	cc *grpc.ClientConn
}

func NewDownstreamMsgClient(cc *grpc.ClientConn) DownstreamMsgClient {
	return &downstreamMsgClient{cc}
}

func (c *downstreamMsgClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/downmsg.DownstreamMsg/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downstreamMsgClient) OnClientUpdated(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/downmsg.DownstreamMsg/OnClientUpdated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downstreamMsgClient) OnRevoke(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/downmsg.DownstreamMsg/OnRevoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downstreamMsgClient) OnRegister(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/downmsg.DownstreamMsg/OnRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DownstreamMsg service

type DownstreamMsgServer interface {
	Send(context.Context, *SendRequest) (*Empty, error)
	// must call when client webhook is updated
	OnClientUpdated(context.Context, *Id) (*Empty, error)
	// must be call when registration change
	OnRevoke(context.Context, *Registration) (*Empty, error)
	OnRegister(context.Context, *Registration) (*Empty, error)
}

func RegisterDownstreamMsgServer(s *grpc.Server, srv DownstreamMsgServer) {
	s.RegisterService(&_DownstreamMsg_serviceDesc, srv)
}

func _DownstreamMsg_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownstreamMsgServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/downmsg.DownstreamMsg/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownstreamMsgServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownstreamMsg_OnClientUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownstreamMsgServer).OnClientUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/downmsg.DownstreamMsg/OnClientUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownstreamMsgServer).OnClientUpdated(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownstreamMsg_OnRevoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownstreamMsgServer).OnRevoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/downmsg.DownstreamMsg/OnRevoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownstreamMsgServer).OnRevoke(ctx, req.(*Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownstreamMsg_OnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownstreamMsgServer).OnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/downmsg.DownstreamMsg/OnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownstreamMsgServer).OnRegister(ctx, req.(*Registration))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownstreamMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "downmsg.DownstreamMsg",
	HandlerType: (*DownstreamMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _DownstreamMsg_Send_Handler,
		},
		{
			MethodName: "OnClientUpdated",
			Handler:    _DownstreamMsg_OnClientUpdated_Handler,
		},
		{
			MethodName: "OnRevoke",
			Handler:    _DownstreamMsg_OnRevoke_Handler,
		},
		{
			MethodName: "OnRegister",
			Handler:    _DownstreamMsg_OnRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "downmsg/downmsg.proto",
}

func init() { proto.RegisterFile("downmsg/downmsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0xd0, 0x24, 0x4e, 0xb5, 0xc2, 0xd2, 0x4a, 0xe8, 0x49, 0x72, 0xf2, 0x14, 0xc5,
	0xe2, 0x59, 0xb0, 0x7a, 0xc8, 0x41, 0x0a, 0x91, 0x3e, 0x40, 0xec, 0x0c, 0x21, 0x68, 0x66, 0xe3,
	0xee, 0xd4, 0xe2, 0xc3, 0xfa, 0x2e, 0x92, 0x64, 0x0d, 0x3d, 0x54, 0xe8, 0x65, 0x97, 0x7f, 0x67,
	0xbe, 0x9d, 0xff, 0x67, 0x60, 0x8a, 0x7a, 0xcb, 0xb5, 0x2d, 0x6f, 0xdc, 0x9d, 0x36, 0x46, 0x8b,
	0x56, 0xa1, 0x93, 0xc9, 0x03, 0x8c, 0x5e, 0x89, 0x31, 0xa7, 0xcf, 0x0d, 0x59, 0x51, 0x97, 0x10,
	0xac, 0x2c, 0x99, 0x0c, 0x63, 0xef, 0xca, 0xbb, 0x3e, 0xcd, 0x83, 0x4d, 0xa7, 0x54, 0x0c, 0xe1,
	0x42, 0xb3, 0x10, 0x4b, 0xec, 0x77, 0x85, 0x70, 0xdd, 0xcb, 0x64, 0x02, 0x7e, 0x86, 0x6a, 0xdc,
	0x9e, 0x8e, 0xf1, 0x2b, 0x4c, 0x42, 0x38, 0x79, 0xae, 0x1b, 0xf9, 0x4e, 0x1e, 0xe1, 0x2c, 0xa7,
	0xb2, 0xb2, 0x62, 0x0a, 0xa9, 0x34, 0xff, 0x3b, 0x60, 0x06, 0xd1, 0xe2, 0xa3, 0x22, 0x96, 0x0c,
	0xdd, 0x84, 0x68, 0xed, 0xf4, 0xdd, 0x8f, 0x07, 0xe7, 0x4f, 0x7a, 0xcb, 0x56, 0x0c, 0x15, 0xf5,
	0x8b, 0x2d, 0x55, 0x0a, 0xc7, 0xad, 0x6b, 0x35, 0x49, 0xff, 0x62, 0xed, 0x84, 0x98, 0x8d, 0x87,
	0xd7, 0xde, 0xc3, 0x91, 0xba, 0x85, 0x8b, 0x25, 0xf7, 0xff, 0xaf, 0x1a, 0x2c, 0x84, 0x50, 0x8d,
	0x86, 0xa6, 0x0c, 0xf7, 0x10, 0x73, 0x88, 0x96, 0x9c, 0xd3, 0x97, 0x7e, 0x27, 0x35, 0x1d, 0xaa,
	0xbb, 0x51, 0xf6, 0x40, 0xf7, 0x00, 0x2d, 0xd4, 0xf6, 0x90, 0x39, 0x18, 0x7b, 0x0b, 0xba, 0x9d,
	0xcc, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x32, 0xd8, 0xb6, 0x20, 0xac, 0x01, 0x00, 0x00,
}