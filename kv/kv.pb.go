// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv/kv.proto

/*
Package kv is a generated protocol buffer package.

It is generated from these files:
	kv/kv.proto

It has these top-level messages:
	Key
	Value
	Bool
*/
package kv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"

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

type Key struct {
	Partition string `protobuf:"bytes,2,opt,name=partition" json:"partition,omitempty"`
	Key       string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Key) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Value struct {
	Partition string `protobuf:"bytes,2,opt,name=partition" json:"partition,omitempty"`
	Key       string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Bytes     []byte `protobuf:"bytes,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Value     string `protobuf:"bytes,5,opt,name=value" json:"value,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Value) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *Value) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Value) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Bool struct {
	Value bool `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*Key)(nil), "kv.Key")
	proto.RegisterType((*Value)(nil), "kv.Value")
	proto.RegisterType((*Bool)(nil), "kv.Bool")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KV service

type KVClient interface {
	Set(ctx context.Context, in *Value, opts ...grpc.CallOption) (*Value, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Has(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Bool, error)
	Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error)
}

type kVClient struct {
	cc *grpc.ClientConn
}

func NewKVClient(cc *grpc.ClientConn) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Set(ctx context.Context, in *Value, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := grpc.Invoke(ctx, "/kv.KV/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := grpc.Invoke(ctx, "/kv.KV/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Has(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := grpc.Invoke(ctx, "/kv.KV/Has", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error) {
	out := new(common.Pong)
	err := grpc.Invoke(ctx, "/kv.KV/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KV service

type KVServer interface {
	Set(context.Context, *Value) (*Value, error)
	Get(context.Context, *Key) (*Value, error)
	Has(context.Context, *Key) (*Bool, error)
	Ping(context.Context, *common.PingRequest) (*common.Pong, error)
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Set(ctx, req.(*Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Has_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Has(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Has",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Has(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Ping(ctx, req.(*common.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kv.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _KV_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KV_Get_Handler,
		},
		{
			MethodName: "Has",
			Handler:    _KV_Has_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _KV_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv/kv.proto",
}

func init() { proto.RegisterFile("kv/kv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4f, 0x84, 0x30,
	0x14, 0xc6, 0x03, 0x05, 0x3d, 0x9e, 0x37, 0x98, 0xea, 0x50, 0xf1, 0x86, 0x0b, 0xd3, 0xb9, 0x94,
	0x44, 0xe3, 0x3f, 0xe0, 0xa2, 0x09, 0x8b, 0xc1, 0xe4, 0x76, 0x8a, 0x2f, 0xd8, 0xf4, 0x8e, 0x9e,
	0x50, 0x9a, 0xe0, 0xe6, 0x7f, 0x6e, 0x5a, 0x88, 0xb8, 0x3a, 0xf5, 0xfb, 0x7e, 0x5f, 0xdf, 0x6b,
	0xf3, 0xc1, 0x85, 0xb2, 0xb9, 0xb2, 0xfc, 0xd4, 0x69, 0xa3, 0x69, 0xa8, 0x6c, 0xca, 0x85, 0x34,
	0x62, 0xa8, 0x15, 0x1a, 0xae, 0xbb, 0x26, 0xef, 0x07, 0x21, 0xbf, 0xf2, 0x0f, 0xac, 0xde, 0xb1,
	0xcb, 0x6b, 0x7d, 0x3c, 0xea, 0x76, 0x3e, 0xa6, 0x99, 0xec, 0x11, 0x48, 0x81, 0x23, 0xdd, 0x40,
	0x72, 0xaa, 0x3a, 0x23, 0x8d, 0xd4, 0x2d, 0x0b, 0xb7, 0xc1, 0x2e, 0x29, 0x17, 0x40, 0x2f, 0x81,
	0x28, 0x1c, 0x19, 0xf1, 0xdc, 0xc9, 0xac, 0x86, 0x78, 0x5f, 0x1d, 0x06, 0xfc, 0xef, 0x20, 0xbd,
	0x86, 0x58, 0x8c, 0x06, 0x7b, 0x16, 0x6d, 0x83, 0xdd, 0xba, 0x9c, 0x8c, 0xa3, 0xd6, 0xad, 0x63,
	0xb1, 0xbf, 0x39, 0x99, 0x6c, 0x03, 0xd1, 0x93, 0xd6, 0x87, 0x25, 0x75, 0xfb, 0x57, 0x73, 0x7a,
	0xff, 0x1d, 0x40, 0x58, 0xec, 0xe9, 0x2d, 0x90, 0x37, 0x34, 0x34, 0xe1, 0xca, 0x72, 0xff, 0xa5,
	0x74, 0x91, 0xf4, 0x06, 0xc8, 0x33, 0x1a, 0x7a, 0xee, 0x48, 0x81, 0xe3, 0xdf, 0x88, 0x01, 0x79,
	0xa9, 0xfa, 0x25, 0x5a, 0x39, 0xe1, 0x9f, 0xbb, 0x83, 0xe8, 0x55, 0xb6, 0x0d, 0xbd, 0xe2, 0x73,
	0x53, 0xce, 0x95, 0xf8, 0x39, 0x60, 0x6f, 0xd2, 0xf5, 0x2f, 0xd4, 0x6d, 0x23, 0xce, 0x7c, 0x89,
	0x0f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x76, 0x4c, 0x9e, 0x87, 0x01, 0x00, 0x00,
}
