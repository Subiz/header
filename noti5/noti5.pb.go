// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noti5/noti5.proto

/*
Package noti5 is a generated protocol buffer package.

It is generated from these files:
	noti5/noti5.proto

It has these top-level messages:
	Setting
*/
package noti5

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

type Type int32

const (
	Type_new_message      Type = 0
	Type_new_conversation Type = 1
	Type_nothing          Type = 2
)

var Type_name = map[int32]string{
	0: "new_message",
	1: "new_conversation",
	2: "nothing",
}
var Type_value = map[string]int32{
	"new_message":      0,
	"new_conversation": 1,
	"nothing":          2,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Setting struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AgentId          *string         `protobuf:"bytes,3,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	WebType          *string         `protobuf:"bytes,4,opt,name=web_type,json=webType" json:"web_type,omitempty"`
	MobileType       *string         `protobuf:"bytes,5,opt,name=mobile_type,json=mobileType" json:"mobile_type,omitempty"`
	EmailType        *string         `protobuf:"bytes,6,opt,name=email_type,json=emailType" json:"email_type,omitempty"`
	EmailAfter       *int32          `protobuf:"varint,7,opt,name=email_after,json=emailAfter" json:"email_after,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Setting) Reset()                    { *m = Setting{} }
func (m *Setting) String() string            { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()               {}
func (*Setting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Setting) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Setting) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Setting) GetAgentId() string {
	if m != nil && m.AgentId != nil {
		return *m.AgentId
	}
	return ""
}

func (m *Setting) GetWebType() string {
	if m != nil && m.WebType != nil {
		return *m.WebType
	}
	return ""
}

func (m *Setting) GetMobileType() string {
	if m != nil && m.MobileType != nil {
		return *m.MobileType
	}
	return ""
}

func (m *Setting) GetEmailType() string {
	if m != nil && m.EmailType != nil {
		return *m.EmailType
	}
	return ""
}

func (m *Setting) GetEmailAfter() int32 {
	if m != nil && m.EmailAfter != nil {
		return *m.EmailAfter
	}
	return 0
}

func init() {
	proto.RegisterType((*Setting)(nil), "noti5.Setting")
	proto.RegisterEnum("noti5.Type", Type_name, Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Noti5Service service

type Noti5ServiceClient interface {
	ReadNotificationSetting(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*Setting, error)
	UpdateNotificationSetting(ctx context.Context, in *Setting, opts ...grpc.CallOption) (*common.Empty, error)
}

type noti5ServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoti5ServiceClient(cc *grpc.ClientConn) Noti5ServiceClient {
	return &noti5ServiceClient{cc}
}

func (c *noti5ServiceClient) ReadNotificationSetting(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*Setting, error) {
	out := new(Setting)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/ReadNotificationSetting", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noti5ServiceClient) UpdateNotificationSetting(ctx context.Context, in *Setting, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/UpdateNotificationSetting", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Noti5Service service

type Noti5ServiceServer interface {
	ReadNotificationSetting(context.Context, *common.Empty) (*Setting, error)
	UpdateNotificationSetting(context.Context, *Setting) (*common.Empty, error)
}

func RegisterNoti5ServiceServer(s *grpc.Server, srv Noti5ServiceServer) {
	s.RegisterService(&_Noti5Service_serviceDesc, srv)
}

func _Noti5Service_ReadNotificationSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).ReadNotificationSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/ReadNotificationSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).ReadNotificationSetting(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noti5Service_UpdateNotificationSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Setting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).UpdateNotificationSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/UpdateNotificationSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).UpdateNotificationSetting(ctx, req.(*Setting))
	}
	return interceptor(ctx, in, info, handler)
}

var _Noti5Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "noti5.Noti5Service",
	HandlerType: (*Noti5ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadNotificationSetting",
			Handler:    _Noti5Service_ReadNotificationSetting_Handler,
		},
		{
			MethodName: "UpdateNotificationSetting",
			Handler:    _Noti5Service_UpdateNotificationSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noti5/noti5.proto",
}

func init() { proto.RegisterFile("noti5/noti5.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0x86, 0xe3, 0x7c, 0xac, 0x37, 0xe3, 0xdd, 0x8d, 0x57, 0x2c, 0xac, 0x13, 0x08, 0x4d, 0x73,
	0x0a, 0x3d, 0xd8, 0x10, 0xe8, 0x25, 0xf4, 0x52, 0x4a, 0x0f, 0xbd, 0xf4, 0xe0, 0xb4, 0xe7, 0x20,
	0xcb, 0x13, 0x47, 0x34, 0x96, 0x8c, 0x3d, 0xf9, 0xea, 0x2f, 0xe8, 0xef, 0xec, 0x2f, 0x29, 0x96,
	0x9c, 0x43, 0x4b, 0x2f, 0x36, 0xf3, 0x3c, 0xef, 0x2b, 0xc1, 0x08, 0xfe, 0x2a, 0x4d, 0xf2, 0x3a,
	0x32, 0xdf, 0xb0, 0x28, 0x35, 0x69, 0xd6, 0x33, 0xc3, 0x28, 0x4c, 0x24, 0x25, 0x3b, 0xf1, 0x82,
	0x14, 0xea, 0x32, 0x8b, 0xaa, 0x5d, 0x22, 0x5f, 0xa3, 0x0d, 0xf2, 0x14, 0xcb, 0x48, 0xe8, 0x3c,
	0xd7, 0xaa, 0xf9, 0xd9, 0xda, 0xf4, 0xdd, 0x01, 0x77, 0x89, 0x44, 0x52, 0x65, 0xec, 0x12, 0x3a,
	0x82, 0x8e, 0x81, 0x33, 0x71, 0x66, 0xde, 0x7c, 0x10, 0x36, 0xb9, 0x3b, 0xad, 0x08, 0x8f, 0x14,
	0xd7, 0x8e, 0x8d, 0x01, 0xb8, 0x10, 0x7a, 0xa7, 0x68, 0x25, 0xd3, 0xa0, 0x3d, 0x71, 0x66, 0xfd,
	0xb8, 0xdf, 0x90, 0x87, 0x94, 0x0d, 0xe1, 0x27, 0xcf, 0xd0, 0xca, 0x8e, 0x91, 0xae, 0x99, 0xad,
	0x3a, 0x60, 0xb2, 0xa2, 0x53, 0x81, 0x41, 0xd7, 0xaa, 0x03, 0x26, 0x4f, 0xa7, 0x02, 0xd9, 0x05,
	0x78, 0xb9, 0x4e, 0xe4, 0x16, 0xad, 0xed, 0x19, 0x0b, 0x16, 0x99, 0xc0, 0x18, 0x00, 0x73, 0x2e,
	0xb7, 0xd6, 0xff, 0xb0, 0xb7, 0x1a, 0x72, 0xee, 0x5b, 0xcd, 0xd7, 0x84, 0x65, 0xe0, 0x4e, 0x9c,
	0x59, 0x2f, 0xb6, 0x8d, 0xdb, 0x9a, 0x5c, 0x2d, 0xa0, 0x6b, 0x82, 0x03, 0xf0, 0x14, 0x1e, 0x56,
	0x39, 0x56, 0x15, 0xcf, 0xd0, 0x6f, 0xb1, 0x7f, 0xe0, 0xd7, 0x40, 0x68, 0xb5, 0xc7, 0xb2, 0xe2,
	0x24, 0xb5, 0xf2, 0x1d, 0xe6, 0x81, 0xab, 0x34, 0x6d, 0xa4, 0xca, 0xfc, 0xf6, 0xfc, 0xcd, 0x81,
	0x5f, 0x8f, 0xf5, 0x6a, 0x97, 0x58, 0xee, 0xa5, 0x40, 0xb6, 0x80, 0xff, 0x31, 0xf2, 0xb4, 0x66,
	0x6b, 0x29, 0x4c, 0xe7, 0xbc, 0xc0, 0xdf, 0xe7, 0x9d, 0xdd, 0xe7, 0x05, 0x9d, 0x46, 0x7f, 0x42,
	0xfb, 0x40, 0x8d, 0x9e, 0xb6, 0xd8, 0x0d, 0x0c, 0x9f, 0x8b, 0x94, 0x13, 0x7e, 0xd7, 0xfe, 0x12,
	0x1f, 0x7d, 0x3e, 0x6d, 0xda, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x67, 0xef, 0x98, 0xf6,
	0x01, 0x00, 0x00,
}