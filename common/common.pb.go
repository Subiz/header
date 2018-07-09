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
	Error
	Device
	PingRequest
	Pong
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import lang "bitbucket.org/subiz/header/lang"
import auth "bitbucket.org/subiz/header/auth"

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
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (Device_DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

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

type Error struct {
	Error       string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	RequestId   string `protobuf:"bytes,3,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Hash        string `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	Debug       string `protobuf:"bytes,6,opt,name=debug" json:"debug,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Error) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Error) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Error) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Error) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
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
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
	proto.RegisterType((*Error)(nil), "common.Error")
	proto.RegisterType((*Device)(nil), "common.Device")
	proto.RegisterType((*PingRequest)(nil), "common.PingRequest")
	proto.RegisterType((*Pong)(nil), "common.Pong")
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
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xc1, 0x6e, 0xdb, 0x46,
	0x10, 0x8d, 0x44, 0x53, 0x12, 0x47, 0xaa, 0xcc, 0x6e, 0x82, 0x82, 0x75, 0x5a, 0x54, 0xe1, 0x25,
	0xaa, 0x03, 0xc8, 0x85, 0x7b, 0x2f, 0x10, 0x38, 0x06, 0x2a, 0xf4, 0x12, 0xb0, 0x46, 0x81, 0x9e,
	0x04, 0x92, 0x3b, 0x96, 0x16, 0x96, 0x76, 0xd9, 0xdd, 0x65, 0x1a, 0xe5, 0x0b, 0x72, 0xe8, 0x5f,
	0xf6, 0x47, 0x8a, 0x99, 0xa5, 0x64, 0xf5, 0x52, 0xf8, 0xd0, 0x0b, 0x35, 0xef, 0xed, 0xec, 0xce,
	0xcc, 0x9b, 0x19, 0x08, 0x9e, 0xd7, 0x66, 0xb7, 0x33, 0xfa, 0x2a, 0xfc, 0x2c, 0x1a, 0x6b, 0xbc,
	0x11, 0x83, 0x80, 0x2e, 0x2e, 0x2b, 0xe5, 0xab, 0xb6, 0x7e, 0x40, 0xbf, 0x30, 0x76, 0x7d, 0xe5,
	0xda, 0x4a, 0x7d, 0xba, 0xda, 0x60, 0x29, 0xd1, 0x5e, 0x6d, 0x4b, 0xbd, 0xe6, 0x4f, 0xb8, 0xf3,
	0x9f, 0xbe, 0x65, 0xeb, 0x37, 0xfc, 0x09, 0xbe, 0xf9, 0x25, 0xc4, 0xb7, 0xbb, 0xc6, 0xef, 0xc5,
	0x2b, 0x88, 0x6a, 0xff, 0x31, 0xeb, 0xcd, 0x7a, 0xf3, 0xf1, 0xf5, 0xf9, 0xa2, 0x4b, 0xe2, 0xc6,
	0x68, 0x8f, 0x1f, 0x7d, 0x41, 0x67, 0xf9, 0x6f, 0xd0, 0x5f, 0xca, 0x27, 0x38, 0x8a, 0x6f, 0x01,
	0xca, 0xba, 0x36, 0xad, 0xf6, 0x2b, 0x25, 0xb3, 0xfe, 0xac, 0x37, 0x4f, 0x8a, 0xa4, 0x63, 0x96,
	0x52, 0x4c, 0xa1, 0xaf, 0x64, 0x16, 0x31, 0xdd, 0x57, 0x32, 0xff, 0x1d, 0xa2, 0xa5, 0x74, 0xff,
	0xc3, 0xc3, 0x29, 0x44, 0x4a, 0xba, 0x2c, 0x9a, 0x45, 0xf3, 0xa4, 0x20, 0x33, 0xff, 0x1c, 0xc1,
	0xb0, 0x7b, 0x41, 0x7c, 0x0d, 0x23, 0xfc, 0x80, 0xe1, 0x6a, 0x8f, 0xaf, 0x0e, 0x19, 0x2f, 0xa5,
	0x78, 0x01, 0xb1, 0xf3, 0xa5, 0x47, 0x7e, 0x72, 0x52, 0x04, 0x20, 0x04, 0x9c, 0x69, 0x23, 0xb1,
	0xcb, 0x94, 0x6d, 0xf1, 0x1d, 0x8c, 0x2d, 0x36, 0xdb, 0xfd, 0xca, 0x9b, 0x46, 0xd5, 0xd9, 0x19,
	0x1f, 0x01, 0x53, 0x77, 0xc4, 0x88, 0x1f, 0x00, 0x6a, 0x8b, 0x12, 0xb5, 0x57, 0xe5, 0x36, 0x1b,
	0x70, 0x31, 0xe9, 0x82, 0x15, 0xbf, 0x39, 0xf2, 0xc5, 0x89, 0x8f, 0xc8, 0x60, 0xe8, 0x6d, 0x59,
	0x2b, 0xbd, 0xce, 0x86, 0x1c, 0xfe, 0x00, 0xc5, 0x4b, 0x48, 0x42, 0xb0, 0x07, 0xdc, 0x67, 0x23,
	0x0e, 0x35, 0x62, 0xe2, 0x17, 0xdc, 0x8b, 0x37, 0x90, 0x54, 0xfb, 0x95, 0xc4, 0x0f, 0xaa, 0xc6,
	0x0c, 0x38, 0xce, 0xf4, 0x20, 0xda, 0x3b, 0x66, 0x8b, 0x51, 0xb5, 0x0f, 0x16, 0x15, 0x18, 0x12,
	0x1e, 0xf3, 0x2b, 0x01, 0x88, 0x6f, 0x20, 0x69, 0x4a, 0xeb, 0x95, 0x57, 0x46, 0x67, 0x5f, 0xcc,
	0x7a, 0xf3, 0xb8, 0x78, 0x24, 0xc4, 0x57, 0x30, 0x30, 0xf7, 0xf7, 0x0e, 0x7d, 0x36, 0x9d, 0xf5,
	0xe6, 0x51, 0xd1, 0x21, 0x92, 0xc5, 0xa3, 0xdd, 0x65, 0xe7, 0xb3, 0xde, 0xfc, 0xac, 0x60, 0x5b,
	0xbc, 0x82, 0x89, 0x35, 0xad, 0x47, 0xdb, 0xe9, 0x92, 0x72, 0x98, 0x71, 0xe0, 0x58, 0x98, 0xfc,
	0xef, 0x3e, 0xc4, 0x05, 0x25, 0xff, 0x94, 0x46, 0x9f, 0xf6, 0x2a, 0x7a, 0x4a, 0xaf, 0x32, 0x18,
	0x3a, 0xb4, 0xac, 0x45, 0x1c, 0xfc, 0x3b, 0x48, 0x33, 0xd3, 0x99, 0xf4, 0xd8, 0x20, 0xcc, 0x4c,
	0xc7, 0x84, 0x99, 0x41, 0x6b, 0x59, 0xc0, 0x51, 0x41, 0xa6, 0x78, 0x0d, 0xe7, 0x68, 0xed, 0x4a,
	0xa2, 0xab, 0xad, 0x6a, 0x58, 0x9b, 0x09, 0xdf, 0x9a, 0xa2, 0xb5, 0xef, 0x1e, 0x59, 0x91, 0xc3,
	0x88, 0x1c, 0x6b, 0x9a, 0x11, 0x52, 0x6f, 0x7a, 0x3d, 0x5c, 0xf0, 0x1a, 0xde, 0x15, 0x43, 0xb4,
	0xf6, 0x86, 0xe6, 0xe5, 0x25, 0x24, 0xec, 0xb3, 0x2d, 0x9d, 0x63, 0xc5, 0xe2, 0x82, 0x2e, 0xdd,
	0x10, 0xe6, 0x2a, 0xad, 0x5d, 0x6d, 0x4a, 0xb7, 0xe9, 0x14, 0xa3, 0x7b, 0x3f, 0x97, 0x6e, 0x43,
	0xf5, 0x34, 0xe5, 0x7e, 0x6b, 0x4a, 0x79, 0x18, 0x8a, 0x0e, 0x72, 0xd3, 0xda, 0x6a, 0xab, 0xdc,
	0x06, 0x65, 0xf6, 0x82, 0x3b, 0xf3, 0x48, 0xe4, 0x9f, 0x7b, 0x10, 0xdf, 0x5a, 0x6b, 0x2c, 0xe9,
	0x84, 0x64, 0x74, 0x6b, 0x12, 0x00, 0xa9, 0x61, 0xf1, 0x8f, 0x16, 0xdd, 0x89, 0xb4, 0x49, 0xc7,
	0x2c, 0xa5, 0x98, 0xc1, 0xf8, 0xb4, 0xee, 0x30, 0xde, 0xa7, 0x14, 0x75, 0x9f, 0xf3, 0x0d, 0x2a,
	0xb3, 0x4d, 0xa1, 0x24, 0x56, 0xed, 0xba, 0x53, 0x37, 0x80, 0xfc, 0xaf, 0x3e, 0x0c, 0xba, 0xf1,
	0xa3, 0x8d, 0x6f, 0x8e, 0x1b, 0xdf, 0x50, 0x16, 0xad, 0x43, 0xbb, 0x2a, 0xd7, 0xa8, 0x7d, 0x17,
	0x25, 0x21, 0xe6, 0x2d, 0x11, 0xe2, 0x0d, 0x7c, 0xe9, 0x6a, 0x8b, 0xa8, 0x57, 0x16, 0x9d, 0xd9,
	0xb6, 0x9c, 0x4b, 0x08, 0x98, 0x86, 0x83, 0xe2, 0xc8, 0x8b, 0x0b, 0x18, 0x79, 0xb5, 0xc3, 0x4f,
	0x46, 0x63, 0x17, 0xff, 0x88, 0xe9, 0x8c, 0x1a, 0xd2, 0x96, 0x6b, 0x64, 0x19, 0x93, 0xe2, 0x88,
	0xe9, 0xcc, 0xe2, 0x3d, 0x5a, 0x8b, 0xf6, 0x71, 0xb7, 0x02, 0xe6, 0x11, 0xdf, 0x37, 0x98, 0x25,
	0xa1, 0x48, 0xb2, 0xf3, 0x9f, 0x00, 0x42, 0x35, 0x77, 0xfb, 0x06, 0xc5, 0x18, 0x86, 0xad, 0x7e,
	0xd0, 0xe6, 0x4f, 0x9d, 0x3e, 0x13, 0x00, 0x83, 0x9d, 0xa9, 0xd4, 0x16, 0xd3, 0x1e, 0xd9, 0xbe,
	0xac, 0xb6, 0xe8, 0xd3, 0x3e, 0x39, 0x49, 0x74, 0x0f, 0xde, 0x34, 0x69, 0x94, 0xbf, 0x86, 0xf1,
	0x7b, 0xa5, 0xd7, 0x45, 0xd0, 0x9a, 0x1a, 0xbc, 0x43, 0xe7, 0x28, 0xb3, 0xd0, 0xa0, 0x03, 0xcc,
	0x67, 0x70, 0xf6, 0xde, 0xe8, 0xf5, 0xa9, 0x47, 0xf4, 0x2f, 0x8f, 0xcb, 0x19, 0xc4, 0xb7, 0xb4,
	0x0d, 0x22, 0x81, 0xf8, 0x57, 0xd4, 0x72, 0x95, 0x3e, 0x13, 0x13, 0x18, 0xbd, 0x6d, 0x14, 0x2f,
	0x58, 0xda, 0xbf, 0xbe, 0x86, 0x38, 0x48, 0xf9, 0x3d, 0x9c, 0x51, 0x54, 0xf1, 0xfc, 0xb0, 0x66,
	0x27, 0x39, 0x5c, 0x4c, 0x8e, 0xa4, 0xd1, 0xeb, 0x6a, 0xc0, 0xff, 0x08, 0x3f, 0xfe, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x1b, 0x3b, 0x5c, 0x88, 0x06, 0x00, 0x00,
}
