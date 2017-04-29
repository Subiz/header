// Code generated by protoc-gen-go.
// source: presence/presence.proto
// DO NOT EDIT!

/*
Package presence is a generated protocol buffer package.

It is generated from these files:
	presence/presence.proto

It has these top-level messages:
	Id
	Empty
	LastFocusTime
*/
package presence

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

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LastFocusTime struct {
	TimeSpan string `protobuf:"bytes,1,opt,name=TimeSpan,json=timeSpan" json:"TimeSpan,omitempty"`
}

func (m *LastFocusTime) Reset()                    { *m = LastFocusTime{} }
func (m *LastFocusTime) String() string            { return proto.CompactTextString(m) }
func (*LastFocusTime) ProtoMessage()               {}
func (*LastFocusTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LastFocusTime) GetTimeSpan() string {
	if m != nil {
		return m.TimeSpan
	}
	return ""
}

func init() {
	proto.RegisterType((*Id)(nil), "presence.Id")
	proto.RegisterType((*Empty)(nil), "presence.Empty")
	proto.RegisterType((*LastFocusTime)(nil), "presence.LastFocusTime")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PresenceMgr service

type PresenceMgrClient interface {
	Ping(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	Bye(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	LastFocusingTime(ctx context.Context, in *Id, opts ...grpc.CallOption) (*LastFocusTime, error)
}

type presenceMgrClient struct {
	cc *grpc.ClientConn
}

func NewPresenceMgrClient(cc *grpc.ClientConn) PresenceMgrClient {
	return &presenceMgrClient{cc}
}

func (c *presenceMgrClient) Ping(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/presence.PresenceMgr/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceMgrClient) Bye(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/presence.PresenceMgr/Bye", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenceMgrClient) LastFocusingTime(ctx context.Context, in *Id, opts ...grpc.CallOption) (*LastFocusTime, error) {
	out := new(LastFocusTime)
	err := grpc.Invoke(ctx, "/presence.PresenceMgr/LastFocusingTime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PresenceMgr service

type PresenceMgrServer interface {
	Ping(context.Context, *Id) (*Empty, error)
	Bye(context.Context, *Id) (*Empty, error)
	LastFocusingTime(context.Context, *Id) (*LastFocusTime, error)
}

func RegisterPresenceMgrServer(s *grpc.Server, srv PresenceMgrServer) {
	s.RegisterService(&_PresenceMgr_serviceDesc, srv)
}

func _PresenceMgr_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceMgrServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/presence.PresenceMgr/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceMgrServer).Ping(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenceMgr_Bye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceMgrServer).Bye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/presence.PresenceMgr/Bye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceMgrServer).Bye(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenceMgr_LastFocusingTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceMgrServer).LastFocusingTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/presence.PresenceMgr/LastFocusingTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceMgrServer).LastFocusingTime(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _PresenceMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "presence.PresenceMgr",
	HandlerType: (*PresenceMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PresenceMgr_Ping_Handler,
		},
		{
			MethodName: "Bye",
			Handler:    _PresenceMgr_Bye_Handler,
		},
		{
			MethodName: "LastFocusingTime",
			Handler:    _PresenceMgr_LastFocusingTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "presence/presence.proto",
}

func init() { proto.RegisterFile("presence/presence.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0x4a, 0x2d,
	0x4e, 0xcd, 0x4b, 0x4e, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x25, 0x11, 0x2e, 0x26, 0xcf, 0x14, 0x21, 0x3e, 0x10, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0xc4, 0x94, 0x99, 0xa2, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52, 0xa9, 0xa4, 0xcd,
	0xc5, 0xeb, 0x93, 0x58, 0x5c, 0xe2, 0x96, 0x9f, 0x5c, 0x5a, 0x1c, 0x92, 0x99, 0x9b, 0x2a, 0x24,
	0xc5, 0xc5, 0x01, 0xa2, 0x83, 0x0b, 0x12, 0xf3, 0xa0, 0xea, 0x39, 0x4a, 0xa0, 0x7c, 0xa3, 0xd9,
	0x8c, 0x5c, 0xdc, 0x01, 0x50, 0x83, 0x7d, 0xd3, 0x8b, 0x84, 0xd4, 0xb9, 0x58, 0x02, 0x32, 0xf3,
	0xd2, 0x85, 0x78, 0xf4, 0xe0, 0xd6, 0x7b, 0xa6, 0x48, 0xf1, 0x23, 0x78, 0x10, 0x3b, 0x18, 0x84,
	0xd4, 0xb8, 0x98, 0x9d, 0x2a, 0x53, 0x09, 0xab, 0xb3, 0xe6, 0x12, 0x80, 0xbb, 0x26, 0x33, 0x2f,
	0x1d, 0xec, 0x20, 0x54, 0x4d, 0xe2, 0x08, 0x1e, 0x8a, 0xbb, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x5e,
	0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x93, 0xc1, 0xd6, 0x15, 0x01, 0x00, 0x00,
}