// Code generated by protoc-gen-go.
// source: scope/scopemgr.proto
// DO NOT EDIT!

/*
Package scope is a generated protocol buffer package.

It is generated from these files:
	scope/scopemgr.proto

It has these top-level messages:
	IdRequest
	ScopesResponse
	Empty
*/
package scope

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import auth "bitbucket.org/subiz/servicespec/proto/auth"

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

type ErrorCode int32

const (
	ErrorCode_NO_ERROR          ErrorCode = 0
	ErrorCode_RESOURCE_NOTFOUND ErrorCode = 1
	ErrorCode_INVALID_REQUEST   ErrorCode = 2
	ErrorCode_UNAUTHORIZED      ErrorCode = 3
	ErrorCode_INTERNAL_ERROR    ErrorCode = 4
)

var ErrorCode_name = map[int32]string{
	0: "NO_ERROR",
	1: "RESOURCE_NOTFOUND",
	2: "INVALID_REQUEST",
	3: "UNAUTHORIZED",
	4: "INTERNAL_ERROR",
}
var ErrorCode_value = map[string]int32{
	"NO_ERROR":          0,
	"RESOURCE_NOTFOUND": 1,
	"INVALID_REQUEST":   2,
	"UNAUTHORIZED":      3,
	"INTERNAL_ERROR":    4,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type IdRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *IdRequest) Reset()                    { *m = IdRequest{} }
func (m *IdRequest) String() string            { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()               {}
func (*IdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ScopesResponse struct {
	Scopes []*auth.Scope `protobuf:"bytes,1,rep,name=Scopes,json=scopes" json:"Scopes,omitempty"`
}

func (m *ScopesResponse) Reset()                    { *m = ScopesResponse{} }
func (m *ScopesResponse) String() string            { return proto.CompactTextString(m) }
func (*ScopesResponse) ProtoMessage()               {}
func (*ScopesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ScopesResponse) GetScopes() []*auth.Scope {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*IdRequest)(nil), "scope.IdRequest")
	proto.RegisterType((*ScopesResponse)(nil), "scope.ScopesResponse")
	proto.RegisterType((*Empty)(nil), "scope.Empty")
	proto.RegisterEnum("scope.ErrorCode", ErrorCode_name, ErrorCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ScopeMgt service

type ScopeMgtClient interface {
	Read(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*auth.Scope, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ScopesResponse, error)
}

type scopeMgtClient struct {
	cc *grpc.ClientConn
}

func NewScopeMgtClient(cc *grpc.ClientConn) ScopeMgtClient {
	return &scopeMgtClient{cc}
}

func (c *scopeMgtClient) Read(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*auth.Scope, error) {
	out := new(auth.Scope)
	err := grpc.Invoke(ctx, "/scope.ScopeMgt/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopeMgtClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ScopesResponse, error) {
	out := new(ScopesResponse)
	err := grpc.Invoke(ctx, "/scope.ScopeMgt/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScopeMgt service

type ScopeMgtServer interface {
	Read(context.Context, *IdRequest) (*auth.Scope, error)
	List(context.Context, *Empty) (*ScopesResponse, error)
}

func RegisterScopeMgtServer(s *grpc.Server, srv ScopeMgtServer) {
	s.RegisterService(&_ScopeMgt_serviceDesc, srv)
}

func _ScopeMgt_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeMgtServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.ScopeMgt/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeMgtServer).Read(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopeMgt_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeMgtServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.ScopeMgt/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeMgtServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScopeMgt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scope.ScopeMgt",
	HandlerType: (*ScopeMgtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _ScopeMgt_Read_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ScopeMgt_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scope/scopemgr.proto",
}

func init() { proto.RegisterFile("scope/scopemgr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x5f, 0x4b, 0x02, 0x41,
	0x14, 0xc5, 0x5d, 0xff, 0xa5, 0x57, 0xb1, 0xed, 0x96, 0x20, 0xf6, 0x22, 0xdb, 0x43, 0x12, 0xb4,
	0x0b, 0x86, 0x1f, 0x40, 0x74, 0xa2, 0x05, 0xdb, 0xa5, 0xeb, 0x6e, 0x0f, 0xbd, 0x88, 0xbb, 0x3b,
	0xd8, 0x22, 0x36, 0xdb, 0xcc, 0x6c, 0x50, 0x9f, 0x3e, 0x1a, 0x25, 0xea, 0xe5, 0xc2, 0x39, 0x9c,
	0x39, 0xfc, 0xce, 0xc0, 0x85, 0x4a, 0x45, 0xc1, 0x3d, 0x73, 0xf7, 0x5b, 0xe9, 0x16, 0x52, 0x68,
	0x81, 0x0d, 0xa3, 0x87, 0xd3, 0x24, 0xd7, 0x49, 0x99, 0xee, 0xb8, 0x76, 0x85, 0xdc, 0x7a, 0xaa,
	0x4c, 0xf2, 0x2f, 0x4f, 0x71, 0xf9, 0x91, 0xa7, 0x5c, 0x15, 0x3c, 0xf5, 0x4c, 0xda, 0xdb, 0x94,
	0xfa, 0xd5, 0x9c, 0xc3, 0x6b, 0xe7, 0x12, 0xda, 0x7e, 0x46, 0xfc, 0xbd, 0xe4, 0x4a, 0x63, 0x0f,
	0xaa, 0x7e, 0x36, 0xb0, 0x46, 0xd6, 0xb8, 0x41, 0xd5, 0x3c, 0x73, 0xa6, 0xd0, 0x5b, 0xfd, 0x94,
	0x2b, 0xe2, 0xaa, 0x10, 0x6f, 0x8a, 0xe3, 0x15, 0x34, 0x0f, 0xce, 0xc0, 0x1a, 0xd5, 0xc6, 0x9d,
	0x49, 0xc7, 0x35, 0x5d, 0xc6, 0xa3, 0xa6, 0x21, 0x51, 0xce, 0x09, 0x34, 0xd8, 0xbe, 0xd0, 0x9f,
	0x37, 0x3b, 0x68, 0x33, 0x29, 0x85, 0x9c, 0x8b, 0x8c, 0x63, 0x17, 0x5a, 0x41, 0xb8, 0x66, 0x44,
	0x21, 0xd9, 0x15, 0xec, 0xc3, 0x19, 0xb1, 0x55, 0x18, 0xd3, 0x9c, 0xad, 0x83, 0x30, 0xba, 0x0f,
	0xe3, 0x60, 0x61, 0x5b, 0x78, 0x0e, 0xa7, 0x7e, 0xf0, 0x3c, 0x5b, 0xfa, 0x8b, 0x35, 0xb1, 0xa7,
	0x98, 0xad, 0x22, 0xbb, 0x8a, 0x36, 0x74, 0xe3, 0x60, 0x16, 0x47, 0x0f, 0x21, 0xf9, 0x2f, 0x6c,
	0x61, 0xd7, 0x10, 0xa1, 0xe7, 0x07, 0x11, 0xa3, 0x60, 0xb6, 0x3c, 0x36, 0xd6, 0x27, 0x09, 0xb4,
	0x0c, 0xc6, 0xe3, 0x56, 0xe3, 0x35, 0xd4, 0x89, 0x6f, 0x32, 0xb4, 0x5d, 0x83, 0xe4, 0xfe, 0x4e,
	0x1c, 0xfe, 0x05, 0x76, 0x2a, 0x78, 0x0b, 0xf5, 0x65, 0xae, 0x34, 0x76, 0x8f, 0x41, 0xc3, 0x3d,
	0xec, 0x1f, 0xd5, 0xff, 0xf1, 0x4e, 0x25, 0x69, 0x9a, 0x4f, 0xbb, 0xfb, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x91, 0xf2, 0x8a, 0x4f, 0x8a, 0x01, 0x00, 0x00,
}
