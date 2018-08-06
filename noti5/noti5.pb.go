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
	Message
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

type Token_Platform int32

const (
	Token_desktop Token_Platform = 0
	Token_mobile  Token_Platform = 1
)

var Token_Platform_name = map[int32]string{
	0: "desktop",
	1: "mobile",
}
var Token_Platform_value = map[string]int32{
	"desktop": 0,
	"mobile":  1,
}

func (x Token_Platform) Enum() *Token_Platform {
	p := new(Token_Platform)
	*p = x
	return p
}
func (x Token_Platform) String() string {
	return proto.EnumName(Token_Platform_name, int32(x))
}
func (x *Token_Platform) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Token_Platform_value, data, "Token_Platform")
	if err != nil {
		return err
	}
	*x = Token_Platform(value)
	return nil
}
func (Token_Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

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
	FcmToken         *string         `protobuf:"bytes,5,opt,name=fcm_token,json=fcmToken" json:"fcm_token,omitempty"`
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

func (m *Token) GetFcmToken() string {
	if m != nil && m.FcmToken != nil {
		return *m.FcmToken
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
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string         `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Platform         *string         `protobuf:"bytes,5,opt,name=platform" json:"platform,omitempty"`
	Title            *string         `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Body             *string         `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PushNoti) Reset()                    { *m = PushNoti{} }
func (m *PushNoti) String() string            { return proto.CompactTextString(m) }
func (*PushNoti) ProtoMessage()               {}
func (*PushNoti) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushNoti) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

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

type Message struct {
	FcmToken         *string `protobuf:"bytes,2,opt,name=fcm_token,json=fcmToken" json:"fcm_token,omitempty"`
	Title            *string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Body             *string `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Message) GetFcmToken() string {
	if m != nil && m.FcmToken != nil {
		return *m.FcmToken
	}
	return ""
}

func (m *Message) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Setting)(nil), "noti5.Setting")
	proto.RegisterType((*Token)(nil), "noti5.Token")
	proto.RegisterType((*PushNoti)(nil), "noti5.PushNoti")
	proto.RegisterType((*Message)(nil), "noti5.Message")
	proto.RegisterEnum("noti5.Type", Type_name, Type_value)
	proto.RegisterEnum("noti5.Event", Event_name, Event_value)
	proto.RegisterEnum("noti5.Token_Platform", Token_Platform_name, Token_Platform_value)
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

// Client API for Noti5TokenService service

type Noti5TokenServiceClient interface {
	AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error)
	RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error)
	Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error)
}

type noti5TokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoti5TokenServiceClient(cc *grpc.ClientConn) Noti5TokenServiceClient {
	return &noti5TokenServiceClient{cc}
}

func (c *noti5TokenServiceClient) AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/noti5.Noti5TokenService/AddToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noti5TokenServiceClient) RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/noti5.Noti5TokenService/RemoveToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noti5TokenServiceClient) Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error) {
	out := new(common.Pong)
	err := grpc.Invoke(ctx, "/noti5.Noti5TokenService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Noti5TokenService service

type Noti5TokenServiceServer interface {
	AddToken(context.Context, *Token) (*common.Empty, error)
	RemoveToken(context.Context, *Token) (*common.Empty, error)
	Ping(context.Context, *common.PingRequest) (*common.Pong, error)
}

func RegisterNoti5TokenServiceServer(s *grpc.Server, srv Noti5TokenServiceServer) {
	s.RegisterService(&_Noti5TokenService_serviceDesc, srv)
}

func _Noti5TokenService_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5TokenServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5TokenService/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5TokenServiceServer).AddToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noti5TokenService_RemoveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5TokenServiceServer).RemoveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5TokenService/RemoveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5TokenServiceServer).RemoveToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noti5TokenService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5TokenServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5TokenService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5TokenServiceServer).Ping(ctx, req.(*common.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Noti5TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "noti5.Noti5TokenService",
	HandlerType: (*Noti5TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToken",
			Handler:    _Noti5TokenService_AddToken_Handler,
		},
		{
			MethodName: "RemoveToken",
			Handler:    _Noti5TokenService_RemoveToken_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Noti5TokenService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noti5/noti5.proto",
}

func init() { proto.RegisterFile("noti5/noti5.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0xc5, 0xf9, 0x67, 0x7b, 0xc2, 0x0f, 0xcc, 0xfe, 0x10, 0x98, 0x20, 0xd4, 0x34, 0x3d, 0x34,
	0xe5, 0x10, 0x24, 0x2a, 0xa4, 0x8a, 0x1b, 0xaa, 0x38, 0x70, 0x68, 0x15, 0x19, 0x7a, 0x8e, 0x1c,
	0xef, 0x24, 0xac, 0x88, 0x77, 0x5d, 0x7b, 0x12, 0x48, 0xcf, 0xfd, 0x06, 0xfd, 0x18, 0xed, 0xa7,
	0xea, 0x27, 0xa9, 0x76, 0xd7, 0xe6, 0x4f, 0x84, 0xaa, 0x56, 0xea, 0x25, 0xf1, 0x7b, 0x6f, 0x66,
	0x3c, 0xef, 0x8d, 0x0c, 0x5b, 0x52, 0x91, 0x38, 0x39, 0x32, 0xbf, 0x83, 0x2c, 0x57, 0xa4, 0x58,
	0xd3, 0x80, 0x4e, 0x7f, 0x2a, 0x68, 0x50, 0xcc, 0xc7, 0xe2, 0xcb, 0x40, 0x22, 0x1d, 0x5d, 0x63,
	0xcc, 0x31, 0x3f, 0x4a, 0x54, 0x9a, 0x2a, 0x59, 0xfe, 0xd9, 0x86, 0xde, 0x4f, 0x07, 0xdc, 0x4b,
	0x24, 0x12, 0x72, 0xca, 0x5e, 0x42, 0x3d, 0xa1, 0xbb, 0xd0, 0xe9, 0x3a, 0xfd, 0xf6, 0xf1, 0xe6,
	0xa0, 0xac, 0x7b, 0xaf, 0x24, 0xe1, 0x1d, 0x45, 0x5a, 0x63, 0x07, 0x00, 0x71, 0x92, 0xa8, 0xb9,
	0xa4, 0x91, 0xe0, 0x61, 0xad, 0xeb, 0xf4, 0xfd, 0xc8, 0x2f, 0x99, 0x0b, 0xce, 0xf6, 0xc0, 0x8b,
	0xa7, 0x68, 0xc5, 0xba, 0x11, 0x5d, 0x83, 0xad, 0x74, 0x8b, 0xe3, 0x11, 0x2d, 0x33, 0x0c, 0x1b,
	0x56, 0xba, 0xc5, 0xf1, 0xd5, 0x32, 0x43, 0xf6, 0x02, 0xda, 0xa9, 0x1a, 0x8b, 0x19, 0x5a, 0xb5,
	0x69, 0x54, 0xb0, 0x94, 0x29, 0x38, 0x00, 0xc0, 0x34, 0x16, 0x33, 0xab, 0xb7, 0xec, 0x5b, 0x0d,
	0x53, 0xf5, 0x5b, 0x39, 0x9e, 0x10, 0xe6, 0xa1, 0xdb, 0x75, 0xfa, 0xcd, 0xc8, 0x76, 0x9c, 0x69,
	0xa6, 0xf7, 0xa3, 0x06, 0xcd, 0x2b, 0x75, 0x83, 0xf2, 0x1f, 0x58, 0xdc, 0x05, 0x77, 0x5e, 0x60,
	0xfe, 0xe0, 0xb0, 0xa5, 0xe1, 0x05, 0x67, 0xfb, 0xe0, 0x1b, 0xe1, 0x91, 0x43, 0x4f, 0x13, 0x66,
	0xc5, 0x7d, 0xf0, 0x27, 0x49, 0x3a, 0x22, 0xbd, 0x44, 0x69, 0xd0, 0x9b, 0x24, 0xa9, 0x5d, 0x6a,
	0x1f, 0x7c, 0x8e, 0x0b, 0x91, 0xa0, 0x1e, 0x6a, 0xdd, 0x79, 0x96, 0xb8, 0xe0, 0xac, 0x03, 0x5e,
	0x36, 0x8b, 0x69, 0xa2, 0xf2, 0xd4, 0x38, 0xf3, 0xa3, 0x7b, 0xcc, 0x42, 0x70, 0x93, 0x1c, 0x63,
	0x42, 0x1e, 0x7a, 0x5d, 0xa7, 0x5f, 0x8f, 0x2a, 0xc8, 0x76, 0xa0, 0x45, 0x2a, 0x13, 0x49, 0x11,
	0xfa, 0xdd, 0xba, 0x5e, 0xd2, 0xa2, 0xde, 0x2b, 0xf0, 0x86, 0x55, 0x77, 0x1b, 0x5c, 0x8e, 0xc5,
	0x0d, 0xa9, 0x2c, 0x58, 0x63, 0x00, 0x2d, 0x1b, 0x78, 0xe0, 0xf4, 0xbe, 0x3b, 0xe0, 0x0d, 0xe7,
	0xc5, 0xf5, 0x47, 0x45, 0xe2, 0xef, 0x13, 0xab, 0xff, 0x26, 0xb1, 0xc6, 0x93, 0xc4, 0x1e, 0x5b,
	0x6b, 0xae, 0x58, 0xdb, 0x86, 0x26, 0x09, 0x9a, 0x55, 0xd7, 0xb6, 0x80, 0x31, 0x68, 0x8c, 0x15,
	0x5f, 0x96, 0x41, 0x98, 0xe7, 0xde, 0x10, 0xdc, 0x0f, 0x58, 0x14, 0xf1, 0x74, 0x25, 0xe5, 0xda,
	0x4a, 0xca, 0xf7, 0x13, 0xeb, 0xcf, 0x4d, 0x6c, 0x3c, 0x4c, 0x3c, 0x3c, 0x85, 0x86, 0x39, 0xda,
	0x26, 0xb4, 0x25, 0xde, 0x8e, 0x52, 0x3b, 0x3d, 0x58, 0x63, 0xdb, 0x10, 0x68, 0x22, 0x51, 0x72,
	0x81, 0x79, 0x11, 0x93, 0x50, 0x32, 0x70, 0x74, 0x8e, 0x52, 0xd1, 0xb5, 0x90, 0xd3, 0xa0, 0x76,
	0xf8, 0x16, 0x9a, 0xe7, 0x0b, 0x94, 0xc4, 0x18, 0x6c, 0xe8, 0xfc, 0x4e, 0x22, 0xfc, 0x3c, 0xc7,
	0x82, 0x90, 0x07, 0x6b, 0x6c, 0x07, 0x98, 0xe1, 0x74, 0xb8, 0x0f, 0xbc, 0x73, 0xfc, 0xd5, 0x81,
	0x75, 0x23, 0x5c, 0x62, 0xae, 0xcf, 0xce, 0xde, 0xc1, 0x6e, 0x84, 0x31, 0xd7, 0xdc, 0x44, 0x24,
	0xe6, 0x45, 0xd5, 0x47, 0xfa, 0x5f, 0x75, 0x82, 0xf3, 0x34, 0xa3, 0x65, 0x67, 0x63, 0x60, 0x3f,
	0xff, 0x4a, 0x3e, 0x85, 0xbd, 0x4f, 0x19, 0x8f, 0x09, 0x9f, 0xeb, 0x5d, 0x29, 0xee, 0x3c, 0x9d,
	0x75, 0xfc, 0xcd, 0x81, 0x2d, 0xb3, 0x86, 0x09, 0xac, 0xda, 0xe5, 0x35, 0x78, 0x67, 0x9c, 0xdb,
	0x0c, 0xd7, 0xcb, 0x01, 0x06, 0xad, 0xb4, 0xb3, 0x43, 0x68, 0x47, 0x98, 0xaa, 0x05, 0xfe, 0x41,
	0xed, 0x1b, 0x68, 0x0c, 0xf5, 0x46, 0xff, 0x57, 0xb4, 0x46, 0x65, 0x24, 0x9d, 0xf5, 0x7b, 0x52,
	0xc9, 0xe9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x83, 0x70, 0xa2, 0xe6, 0x04, 0x00, 0x00,
}
