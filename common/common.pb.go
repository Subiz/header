// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto

It has these top-level messages:
	Empty
	Id
	Ids
	Context
	Reply
	Device
	Error
	PingRequest
	Pong
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import lang "git.subiz.net/header/lang"
import auth "git.subiz.net/header/auth"

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

type E int32

const (
	E_no_error            E = 0
	E_unknown             E = 1
	E_missing_account_id  E = 10
	E_missing_stripe_info E = 212
)

var E_name = map[int32]string{
	0:   "no_error",
	1:   "unknown",
	10:  "missing_account_id",
	212: "missing_stripe_info",
}
var E_value = map[string]int32{
	"no_error":            0,
	"unknown":             1,
	"missing_account_id":  10,
	"missing_stripe_info": 212,
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}
func (E) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event int32

const (
	Event_Send_    Event = 0
	Event_ApiReply Event = 2
)

var Event_name = map[int32]string{
	0: "Send_",
	2: "ApiReply",
}
var Event_value = map[string]int32{
	"Send_":    0,
	"ApiReply": 2,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Device_DeviceType int32

const (
	Device_unknown Device_DeviceType = 0
	Device_mobile  Device_DeviceType = 1
	Device_tablet  Device_DeviceType = 2
	Device_desktop Device_DeviceType = 3
)

var Device_DeviceType_name = map[int32]string{
	0: "unknown",
	1: "mobile",
	2: "tablet",
	3: "desktop",
}
var Device_DeviceType_value = map[string]int32{
	"unknown": 0,
	"mobile":  1,
	"tablet":  2,
	"desktop": 3,
}

func (x Device_DeviceType) String() string {
	return proto.EnumName(Device_DeviceType_name, int32(x))
}
func (Device_DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Empty struct {
	Ctx *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Empty) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

type Id struct {
	Ctx       *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string   `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id        string   `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Id) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Ids struct {
	Ctx       *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string   `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Ids       []string `protobuf:"bytes,3,rep,name=ids" json:"ids,omitempty"`
}

func (m *Ids) Reset()                    { *m = Ids{} }
func (m *Ids) String() string            { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()               {}
func (*Ids) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ids) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Ids) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Ids) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Context struct {
	EventId    string `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	State      []byte `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Node       string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	ReplyTopic string `protobuf:"bytes,4,opt,name=reply_topic,json=replyTopic" json:"reply_topic,omitempty"`
	// 	optional int32 reply_partition = 5;
	Credential *auth.Credential `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Tracing    []byte           `protobuf:"bytes,7,opt,name=tracing,proto3" json:"tracing,omitempty"`
	ReplyKey   string           `protobuf:"bytes,8,opt,name=reply_key,json=replyKey" json:"reply_key,omitempty"`
	ByDevice   *Device          `protobuf:"bytes,10,opt,name=by_device,json=byDevice" json:"by_device,omitempty"`
	// for kafka
	Topic       string `protobuf:"bytes,11,opt,name=topic" json:"topic,omitempty"`
	Partition   int32  `protobuf:"varint,13,opt,name=partition" json:"partition,omitempty"`
	Offset      int64  `protobuf:"varint,14,opt,name=offset" json:"offset,omitempty"`
	Term        uint64 `protobuf:"varint,15,opt,name=term" json:"term,omitempty"`
	RouterTopic string `protobuf:"bytes,16,opt,name=router_topic,json=routerTopic" json:"router_topic,omitempty"`
}

func (m *Context) Reset()                    { *m = Context{} }
func (m *Context) String() string            { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()               {}
func (*Context) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Context) GetEventId() string {
	if m != nil {
		return m.EventId
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
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Context) GetReplyTopic() string {
	if m != nil {
		return m.ReplyTopic
	}
	return ""
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
	if m != nil {
		return m.ReplyKey
	}
	return ""
}

func (m *Context) GetByDevice() *Device {
	if m != nil {
		return m.ByDevice
	}
	return nil
}

func (m *Context) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Context) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *Context) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Context) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Context) GetRouterTopic() string {
	if m != nil {
		return m.RouterTopic
	}
	return ""
}

type Reply struct {
	Ctx            *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	EventId        string   `protobuf:"bytes,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	State          []byte   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Service        string   `protobuf:"bytes,5,opt,name=service" json:"service,omitempty"`
	ServiceId      string   `protobuf:"bytes,6,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	Err            bool     `protobuf:"varint,10,opt,name=err" json:"err,omitempty"`
	ErrDescription string   `protobuf:"bytes,12,opt,name=err_description,json=errDescription" json:"err_description,omitempty"`
	ErrCode        lang.T   `protobuf:"varint,13,opt,name=err_code,json=errCode,enum=lang.T" json:"err_code,omitempty"`
	ErrClass       int32    `protobuf:"varint,15,opt,name=err_class,json=errClass" json:"err_class,omitempty"`
	ErrHash        string   `protobuf:"bytes,16,opt,name=err_hash,json=errHash" json:"err_hash,omitempty"`
	Payload        []byte   `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
	Published      int64    `protobuf:"varint,20,opt,name=published" json:"published,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Reply) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Reply) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Reply) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Reply) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Reply) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *Reply) GetErr() bool {
	if m != nil {
		return m.Err
	}
	return false
}

func (m *Reply) GetErrDescription() string {
	if m != nil {
		return m.ErrDescription
	}
	return ""
}

func (m *Reply) GetErrCode() lang.T {
	if m != nil {
		return m.ErrCode
	}
	return lang.T_undefined
}

func (m *Reply) GetErrClass() int32 {
	if m != nil {
		return m.ErrClass
	}
	return 0
}

func (m *Reply) GetErrHash() string {
	if m != nil {
		return m.ErrHash
	}
	return ""
}

func (m *Reply) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Reply) GetPublished() int64 {
	if m != nil {
		return m.Published
	}
	return 0
}

type Device struct {
	Ip               string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	UserAgent        string `protobuf:"bytes,4,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	ScreenResolution string `protobuf:"bytes,5,opt,name=screen_resolution,json=screenResolution" json:"screen_resolution,omitempty"`
	Timezone         string `protobuf:"bytes,6,opt,name=timezone" json:"timezone,omitempty"`
	Language         string `protobuf:"bytes,7,opt,name=language" json:"language,omitempty"`
	Referrer         string `protobuf:"bytes,8,opt,name=referrer" json:"referrer,omitempty"`
	Type             string `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Device) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Device) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *Device) GetScreenResolution() string {
	if m != nil {
		return m.ScreenResolution
	}
	return ""
}

func (m *Device) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Device) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Device) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *Device) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Error struct {
	Ctx         *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Debug       string   `protobuf:"bytes,3,opt,name=debug" json:"debug,omitempty"`
	Hash        string   `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	Class       int32    `protobuf:"varint,6,opt,name=class" json:"class,omitempty"`
	Stack       string   `protobuf:"bytes,7,opt,name=stack" json:"stack,omitempty"`
	Created     int64    `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Code        string   `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Error) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Error) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
}

func (m *Error) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Error) GetClass() int32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *Error) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

func (m *Error) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PingRequest struct {
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Pong struct {
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Pong) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*Id)(nil), "common.Id")
	proto.RegisterType((*Ids)(nil), "common.Ids")
	proto.RegisterType((*Context)(nil), "common.Context")
	proto.RegisterType((*Reply)(nil), "common.Reply")
	proto.RegisterType((*Device)(nil), "common.Device")
	proto.RegisterType((*Error)(nil), "common.Error")
	proto.RegisterType((*PingRequest)(nil), "common.PingRequest")
	proto.RegisterType((*Pong)(nil), "common.Pong")
	proto.RegisterEnum("common.E", E_name, E_value)
	proto.RegisterEnum("common.Event", Event_name, Event_value)
	proto.RegisterEnum("common.Device_DeviceType", Device_DeviceType_name, Device_DeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agent service

type AgentClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Pong, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/common.Agent/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agent service

type AgentServer interface {
	Ping(context.Context, *PingRequest) (*Pong, error)
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Agent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Agent_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/common.proto",
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 889 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0x8e, 0x2c, 0xcb, 0xb6, 0xc6, 0xae, 0xa3, 0x32, 0x8b, 0x85, 0x9a, 0x6d, 0x51, 0xaf, 0x50,
	0x60, 0xdd, 0x2c, 0xe0, 0x14, 0xe9, 0xbd, 0xc0, 0x22, 0x1b, 0xa0, 0x41, 0x2f, 0x81, 0x1a, 0x14,
	0xe8, 0x49, 0x90, 0xc5, 0x89, 0x4c, 0xc4, 0x26, 0x55, 0x92, 0xda, 0xae, 0xf3, 0x04, 0x3d, 0xf4,
	0xe1, 0x7a, 0xe8, 0x1b, 0xf4, 0x45, 0x0a, 0x0e, 0xe5, 0x9f, 0x1e, 0x0a, 0xe4, 0xb0, 0x17, 0x67,
	0xbe, 0x8f, 0x43, 0x0e, 0xe7, 0x9b, 0x8f, 0x0a, 0x9c, 0x55, 0x6a, 0xb3, 0x51, 0xf2, 0xd2, 0xff,
	0x59, 0x34, 0x5a, 0x59, 0xc5, 0x06, 0x1e, 0x9d, 0x7f, 0x53, 0x0b, 0xbb, 0x30, 0xed, 0x52, 0x3c,
	0x2d, 0x24, 0xda, 0xcb, 0x15, 0x96, 0x1c, 0xf5, 0xe5, 0xba, 0x94, 0x35, 0xfd, 0xf8, 0xec, 0xff,
	0xc9, 0x2a, 0x5b, 0xbb, 0xa2, 0x1f, 0x9f, 0x95, 0x5d, 0x40, 0x74, 0xb3, 0x69, 0xec, 0x96, 0xbd,
	0x86, 0xb0, 0xb2, 0x1f, 0xd3, 0x60, 0x16, 0xcc, 0xc7, 0x57, 0xa7, 0x8b, 0xae, 0xf0, 0xb5, 0x92,
	0x16, 0x3f, 0xda, 0xdc, 0xad, 0x65, 0xbf, 0x40, 0xef, 0x96, 0x3f, 0x23, 0x91, 0x7d, 0x05, 0x50,
	0x56, 0x95, 0x6a, 0xa5, 0x2d, 0x04, 0x4f, 0x7b, 0xb3, 0x60, 0x1e, 0xe7, 0x71, 0xc7, 0xdc, 0x72,
	0x36, 0x85, 0x9e, 0xe0, 0x69, 0x48, 0x74, 0x4f, 0xf0, 0xec, 0x57, 0x08, 0x6f, 0xb9, 0xf9, 0x04,
	0x07, 0x27, 0x10, 0x0a, 0x6e, 0xd2, 0x70, 0x16, 0xce, 0xe3, 0xdc, 0x85, 0xd9, 0x1f, 0x21, 0x0c,
	0xbb, 0x13, 0xd8, 0x17, 0x30, 0xc2, 0x0f, 0xe8, 0xb7, 0x06, 0xb4, 0x75, 0x48, 0xf8, 0x96, 0xb3,
	0x17, 0x10, 0x19, 0x5b, 0x5a, 0xa4, 0x23, 0x27, 0xb9, 0x07, 0x8c, 0x41, 0x5f, 0x2a, 0x8e, 0xdd,
	0x4d, 0x29, 0x66, 0x5f, 0xc3, 0x58, 0x63, 0xb3, 0xde, 0x16, 0x56, 0x35, 0xa2, 0x4a, 0xfb, 0xb4,
	0x04, 0x44, 0xdd, 0x3b, 0x86, 0x7d, 0x07, 0x50, 0x69, 0xe4, 0x28, 0xad, 0x28, 0xd7, 0xe9, 0x80,
	0x9a, 0x49, 0x16, 0xa4, 0xf8, 0xf5, 0x9e, 0xcf, 0x8f, 0x72, 0x58, 0x0a, 0x43, 0xab, 0xcb, 0x4a,
	0xc8, 0x3a, 0x1d, 0x52, 0xf9, 0x1d, 0x64, 0xaf, 0x20, 0xf6, 0xc5, 0x1e, 0x71, 0x9b, 0x8e, 0xa8,
	0xd4, 0x88, 0x88, 0x9f, 0x70, 0xcb, 0xde, 0x42, 0xbc, 0xdc, 0x16, 0x1c, 0x3f, 0x88, 0x0a, 0x53,
	0xa0, 0x3a, 0xd3, 0x9d, 0x68, 0xef, 0x89, 0xcd, 0x47, 0xcb, 0xad, 0x8f, 0x5c, 0x83, 0xfe, 0xc2,
	0x63, 0x3a, 0xc5, 0x03, 0xf6, 0x25, 0xc4, 0x4d, 0xa9, 0xad, 0xb0, 0x42, 0xc9, 0xf4, 0xb3, 0x59,
	0x30, 0x8f, 0xf2, 0x03, 0xc1, 0x5e, 0xc2, 0x40, 0x3d, 0x3c, 0x18, 0xb4, 0xe9, 0x74, 0x16, 0xcc,
	0xc3, 0xbc, 0x43, 0x4e, 0x16, 0x8b, 0x7a, 0x93, 0x9e, 0xce, 0x82, 0x79, 0x3f, 0xa7, 0x98, 0xbd,
	0x86, 0x89, 0x56, 0xad, 0x45, 0xdd, 0xe9, 0x92, 0x50, 0x99, 0xb1, 0xe7, 0x48, 0x98, 0xec, 0x9f,
	0x1e, 0x44, 0xb9, 0xbb, 0xfc, 0x73, 0x06, 0x7d, 0x3c, 0xab, 0xf0, 0x39, 0xb3, 0x4a, 0x61, 0x68,
	0x50, 0x93, 0x16, 0x91, 0xcf, 0xef, 0xa0, 0xf3, 0x4c, 0x17, 0xba, 0xc3, 0x06, 0xde, 0x33, 0x1d,
	0xe3, 0x3d, 0x83, 0x5a, 0x93, 0x80, 0xa3, 0xdc, 0x85, 0xec, 0x0d, 0x9c, 0xa2, 0xd6, 0x05, 0x47,
	0x53, 0x69, 0xd1, 0x90, 0x36, 0x13, 0xda, 0x35, 0x45, 0xad, 0xdf, 0x1f, 0x58, 0x96, 0xc1, 0xc8,
	0x25, 0x56, 0xce, 0x23, 0x4e, 0xbd, 0xe9, 0xd5, 0x70, 0x41, 0x0f, 0xf0, 0x3e, 0x1f, 0xa2, 0xd6,
	0xd7, 0xce, 0x2f, 0xaf, 0x20, 0xa6, 0x9c, 0x75, 0x69, 0x0c, 0x29, 0x16, 0xe5, 0x6e, 0xd3, 0xb5,
	0xc3, 0xd4, 0xa5, 0xd6, 0xc5, 0xaa, 0x34, 0xab, 0x4e, 0x31, 0xb7, 0xef, 0xc7, 0xd2, 0xac, 0x5c,
	0x3f, 0x4d, 0xb9, 0x5d, 0xab, 0x92, 0xef, 0x4c, 0xd1, 0x41, 0x1a, 0x5a, 0xbb, 0x5c, 0x0b, 0xb3,
	0x42, 0x9e, 0xbe, 0xa0, 0xc9, 0x1c, 0x88, 0xec, 0xcf, 0x1e, 0x0c, 0xba, 0x99, 0xbb, 0x67, 0xd6,
	0xec, 0x9f, 0x59, 0xe3, 0x84, 0x68, 0x0d, 0xea, 0xa2, 0xac, 0x51, 0xda, 0xce, 0xb9, 0xb1, 0x63,
	0xde, 0x39, 0x82, 0xbd, 0x85, 0xcf, 0x4d, 0xa5, 0x11, 0x65, 0xa1, 0xd1, 0xa8, 0x75, 0x4b, 0x8d,
	0x7b, 0x2d, 0x13, 0xbf, 0x90, 0xef, 0x79, 0x76, 0x0e, 0x23, 0x2b, 0x36, 0xf8, 0xa4, 0x24, 0x76,
	0x92, 0xee, 0xb1, 0x5b, 0x73, 0x2a, 0xb4, 0x65, 0x8d, 0x74, 0xf7, 0x38, 0xdf, 0x63, 0xb7, 0xa6,
	0xf1, 0x01, 0xb5, 0x46, 0x7d, 0x30, 0xb4, 0xc7, 0xe4, 0xab, 0x6d, 0x83, 0x69, 0xec, 0x9f, 0x9b,
	0x8b, 0xb3, 0x1f, 0x00, 0x7c, 0x37, 0xf7, 0xdb, 0x06, 0xd9, 0x18, 0x86, 0xad, 0x7c, 0x94, 0xea,
	0x77, 0x99, 0x9c, 0x30, 0x80, 0xc1, 0x46, 0x2d, 0xc5, 0x1a, 0x93, 0xc0, 0xc5, 0xb6, 0x5c, 0xae,
	0xd1, 0x26, 0x3d, 0x97, 0xc4, 0xd1, 0x3c, 0x5a, 0xd5, 0x24, 0x61, 0xf6, 0x57, 0x00, 0xd1, 0x8d,
	0xd6, 0x4a, 0x3f, 0xc7, 0x74, 0x33, 0x18, 0x1f, 0x0f, 0xdd, 0x7f, 0x5e, 0x8e, 0x29, 0xe7, 0x3d,
	0x8e, 0xcb, 0xb6, 0xee, 0x54, 0xf5, 0xc0, 0x5d, 0x9c, 0x46, 0xe8, 0xc5, 0xa2, 0xd8, 0x65, 0xfa,
	0x99, 0x0f, 0x68, 0xe6, 0x1e, 0x74, 0xde, 0xad, 0x1e, 0x3b, 0x5d, 0x3c, 0x70, 0xb3, 0xae, 0x34,
	0x96, 0x16, 0x39, 0x69, 0x12, 0xe6, 0x3b, 0xe8, 0x4e, 0x26, 0x77, 0xf9, 0x61, 0x51, 0x9c, 0xbd,
	0x81, 0xf1, 0x9d, 0x90, 0x75, 0x8e, 0xbf, 0xb5, 0x68, 0xac, 0xdb, 0xbc, 0x41, 0x63, 0x9c, 0xd8,
	0xfe, 0xc2, 0x3b, 0x98, 0xcd, 0xa0, 0x7f, 0xa7, 0x64, 0x7d, 0x9c, 0x11, 0xfe, 0x27, 0xe3, 0xe2,
	0x0e, 0x82, 0x1b, 0x36, 0x81, 0x91, 0x54, 0x05, 0x3a, 0x91, 0x92, 0x93, 0x63, 0x89, 0x03, 0xf6,
	0x12, 0xd8, 0x46, 0x18, 0x23, 0x64, 0x5d, 0x1c, 0x3e, 0xbb, 0x09, 0xb0, 0x14, 0xce, 0x76, 0xbc,
	0xb1, 0x5a, 0x34, 0x58, 0x08, 0xf9, 0xa0, 0x92, 0xbf, 0x83, 0x8b, 0x19, 0x44, 0x37, 0xee, 0x9d,
	0xb2, 0x18, 0xa2, 0x9f, 0x51, 0xf2, 0x22, 0x39, 0x71, 0x05, 0xde, 0x35, 0x82, 0x9e, 0x7e, 0xd2,
	0xbb, 0xba, 0x82, 0xc8, 0xfb, 0xed, 0x5b, 0xe8, 0xbb, 0x3e, 0xd8, 0xd9, 0x6e, 0x16, 0x47, 0x5d,
	0x9d, 0x4f, 0xf6, 0xa4, 0x92, 0xf5, 0x72, 0x40, 0xff, 0xab, 0xbe, 0xff, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x3f, 0x57, 0xbd, 0x16, 0x07, 0x00, 0x00,
}
