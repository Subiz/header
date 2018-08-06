// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noti5/noti5.proto

/*
Package noti5 is a generated protocol buffer package.

It is generated from these files:
	noti5/noti5.proto

It has these top-level messages:
	Setting
	Token
	PushNoti
*/
package noti5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "git.subiz.net/header/common"

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

type Event int32

const (
	Event_Noti5Requested     Event = 0
	Event_Noti5PushRequested Event = 1
)

var Event_name = map[int32]string{
	0: "Noti5Requested",
	1: "Noti5PushRequested",
}
var Event_value = map[string]int32{
	"Noti5Requested":     0,
	"Noti5PushRequested": 1,
}

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}
func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (x *Event) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_value, data, "Event")
	if err != nil {
		return err
	}
	*x = Event(value)
	return nil
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

type Token struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserType         *string         `protobuf:"bytes,4,opt,name=user_type,json=userType" json:"user_type,omitempty"`
	Token            *string         `protobuf:"bytes,5,opt,name=token" json:"token,omitempty"`
	DeviceId         *string         `protobuf:"bytes,6,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Platform         *string         `protobuf:"bytes,7,opt,name=platform" json:"platform,omitempty"`
	Created          *int64          `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Topics           []string        `protobuf:"bytes,9,rep,name=topics" json:"topics,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Token) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Token) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Token) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *Token) GetUserType() string {
	if m != nil && m.UserType != nil {
		return *m.UserType
	}
	return ""
}

func (m *Token) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Token) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *Token) GetPlatform() string {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return ""
}

func (m *Token) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Token) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type PushNoti struct {
	AccountId        *string `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Platform         *string `protobuf:"bytes,5,opt,name=platform" json:"platform,omitempty"`
	Title            *string `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Body             *string `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PushNoti) Reset()                    { *m = PushNoti{} }
func (m *PushNoti) String() string            { return proto.CompactTextString(m) }
func (*PushNoti) ProtoMessage()               {}
func (*PushNoti) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushNoti) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *PushNoti) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *PushNoti) GetPlatform() string {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return ""
}

func (m *PushNoti) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *PushNoti) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Setting)(nil), "noti5.Setting")
	proto.RegisterType((*Token)(nil), "noti5.Token")
	proto.RegisterType((*PushNoti)(nil), "noti5.PushNoti")
	proto.RegisterEnum("noti5.Type", Type_name, Type_value)
	proto.RegisterEnum("noti5.Event", Event_name, Event_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Noti5Service service

type Noti5ServiceClient interface {
	// rpc ReadNotificationSetting(common.Empty) returns (Setting);
	// rpc UpdateNotificationSetting(Setting) returns (common.Empty);
	AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error)
	RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error)
	Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error)
}

type noti5ServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoti5ServiceClient(cc *grpc.ClientConn) Noti5ServiceClient {
	return &noti5ServiceClient{cc}
}

func (c *noti5ServiceClient) AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/AddToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noti5ServiceClient) RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/RemoveToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noti5ServiceClient) Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error) {
	out := new(common.Pong)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Noti5Service service

type Noti5ServiceServer interface {
	// rpc ReadNotificationSetting(common.Empty) returns (Setting);
	// rpc UpdateNotificationSetting(Setting) returns (common.Empty);
	AddToken(context.Context, *Token) (*common.Empty, error)
	RemoveToken(context.Context, *Token) (*common.Empty, error)
	Ping(context.Context, *common.PingRequest) (*common.Pong, error)
}

func RegisterNoti5ServiceServer(s *grpc.Server, srv Noti5ServiceServer) {
	s.RegisterService(&_Noti5Service_serviceDesc, srv)
}

func _Noti5Service_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).AddToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noti5Service_RemoveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).RemoveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/RemoveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).RemoveToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noti5Service_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).Ping(ctx, req.(*common.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Noti5Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "noti5.Noti5Service",
	HandlerType: (*Noti5ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToken",
			Handler:    _Noti5Service_AddToken_Handler,
		},
		{
			MethodName: "RemoveToken",
			Handler:    _Noti5Service_RemoveToken_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Noti5Service_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noti5/noti5.proto",
}

func init() { proto.RegisterFile("noti5/noti5.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xb1, 0x6f, 0xd3, 0x40,
	0x14, 0xc6, 0x7b, 0x49, 0x1c, 0x3b, 0x2f, 0x81, 0x86, 0xa3, 0x2a, 0x26, 0xa8, 0x22, 0x64, 0x21,
	0x64, 0x48, 0xa4, 0xa2, 0x2e, 0x6c, 0x15, 0xea, 0x90, 0x05, 0x55, 0x6e, 0xf7, 0xc8, 0xf1, 0xbd,
	0x3a, 0x27, 0xe2, 0x3b, 0x63, 0xbf, 0x24, 0x0d, 0x13, 0x23, 0x13, 0x7f, 0x24, 0x7f, 0x09, 0xba,
	0x3b, 0xbb, 0x2d, 0x1d, 0x10, 0x43, 0x17, 0xdb, 0xdf, 0xf7, 0xbd, 0x67, 0xbf, 0xdf, 0xf9, 0x0e,
	0x5e, 0x28, 0x4d, 0xf2, 0x6c, 0x66, 0xaf, 0xd3, 0xbc, 0xd0, 0xa4, 0xb9, 0x67, 0xc5, 0x60, 0x9c,
	0x4a, 0x9a, 0x96, 0x9b, 0xa5, 0xfc, 0x3e, 0x55, 0x48, 0xb3, 0x15, 0xc6, 0x02, 0x8b, 0x59, 0xa2,
	0xb3, 0x4c, 0xab, 0xea, 0xe6, 0x1a, 0x46, 0xbf, 0x19, 0xf8, 0x57, 0x48, 0x24, 0x55, 0xca, 0xdf,
	0x41, 0x33, 0xa1, 0xdb, 0x90, 0x0d, 0xd9, 0xb8, 0x7b, 0x7a, 0x38, 0xad, 0xea, 0x3e, 0x6b, 0x45,
	0x78, 0x4b, 0x91, 0xc9, 0xf8, 0x09, 0x40, 0x9c, 0x24, 0x7a, 0xa3, 0x68, 0x21, 0x45, 0xd8, 0x18,
	0xb2, 0x71, 0x27, 0xea, 0x54, 0xce, 0x5c, 0xf0, 0xd7, 0x10, 0xc4, 0x29, 0xba, 0xb0, 0x69, 0x43,
	0xdf, 0x6a, 0x17, 0xed, 0x70, 0xb9, 0xa0, 0x7d, 0x8e, 0x61, 0xcb, 0x45, 0x3b, 0x5c, 0x5e, 0xef,
	0x73, 0xe4, 0x6f, 0xa1, 0x9b, 0xe9, 0xa5, 0x5c, 0xa3, 0x4b, 0x3d, 0x9b, 0x82, 0xb3, 0x6c, 0xc1,
	0x09, 0x00, 0x66, 0xb1, 0x5c, 0xbb, 0xbc, 0xed, 0xbe, 0x6a, 0x9d, 0xba, 0xdf, 0xc5, 0xf1, 0x0d,
	0x61, 0x11, 0xfa, 0x43, 0x36, 0xf6, 0x22, 0xd7, 0x71, 0x6e, 0x9c, 0xd1, 0x8f, 0x06, 0x78, 0xd7,
	0xfa, 0x2b, 0xaa, 0x27, 0x40, 0x7c, 0x05, 0xfe, 0xa6, 0xc4, 0xe2, 0x9e, 0xb0, 0x6d, 0xe4, 0x5c,
	0xf0, 0x37, 0xd0, 0xb1, 0xc1, 0x03, 0xc2, 0xc0, 0x18, 0x76, 0xc4, 0x23, 0xf0, 0xc8, 0x0c, 0x50,
	0xc1, 0x39, 0x61, 0x5a, 0x04, 0x6e, 0x65, 0x82, 0xe6, 0x6d, 0x0e, 0x2b, 0x70, 0xc6, 0x5c, 0xf0,
	0x01, 0x04, 0xf9, 0x3a, 0xa6, 0x1b, 0x5d, 0x64, 0x16, 0xa9, 0x13, 0xdd, 0x69, 0x1e, 0x82, 0x9f,
	0x14, 0x18, 0x13, 0x8a, 0x30, 0x18, 0xb2, 0x71, 0x33, 0xaa, 0x25, 0x3f, 0x86, 0x36, 0xe9, 0x5c,
	0x26, 0x65, 0xd8, 0x19, 0x36, 0xcd, 0x74, 0x4e, 0x8d, 0x7e, 0x32, 0x08, 0x2e, 0x37, 0xe5, 0xea,
	0x8b, 0x26, 0xf9, 0x08, 0xb1, 0xf9, 0x0f, 0xc4, 0xd6, 0x5f, 0x88, 0x0f, 0x47, 0xf2, 0x1e, 0x8d,
	0x64, 0x08, 0x25, 0xad, 0xeb, 0xdf, 0xe3, 0x04, 0xe7, 0xd0, 0x5a, 0x6a, 0xb1, 0xaf, 0x00, 0xec,
	0xf3, 0xe4, 0x13, 0xb4, 0xec, 0x9a, 0x1c, 0x42, 0x57, 0xe1, 0x6e, 0x91, 0x61, 0x59, 0xc6, 0x29,
	0xf6, 0x0f, 0xf8, 0x11, 0xf4, 0x8d, 0x91, 0x68, 0xb5, 0xc5, 0xa2, 0x8c, 0x49, 0x6a, 0xd5, 0x67,
	0xbc, 0x0b, 0xbe, 0xd2, 0xb4, 0x92, 0x2a, 0xed, 0x37, 0x26, 0x1f, 0xc1, 0xbb, 0xd8, 0xa2, 0x22,
	0xce, 0xe1, 0xb9, 0x41, 0x39, 0x8b, 0xf0, 0xdb, 0x06, 0x4b, 0x42, 0xd1, 0x3f, 0xe0, 0xc7, 0xc0,
	0xad, 0x67, 0x38, 0xef, 0x7d, 0x76, 0xfa, 0x8b, 0x41, 0xcf, 0x06, 0x57, 0x58, 0x98, 0xc5, 0xe5,
	0xef, 0x21, 0x38, 0x17, 0xc2, 0xed, 0x88, 0xde, 0xd4, 0x9d, 0x1f, 0xab, 0x06, 0xcf, 0xea, 0x2d,
	0x71, 0x91, 0xe5, 0xb4, 0xe7, 0x13, 0xe8, 0x46, 0x98, 0xe9, 0x2d, 0xfe, 0x47, 0xed, 0x07, 0x68,
	0x5d, 0x9a, 0x53, 0xf4, 0xb2, 0xb6, 0x8d, 0xaa, 0xc6, 0x18, 0xf4, 0xee, 0x4c, 0xad, 0xd2, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x5e, 0xd0, 0xb4, 0xb9, 0x03, 0x00, 0x00,
}
