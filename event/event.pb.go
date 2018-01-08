// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event/event.proto

/*
Package event is a generated protocol buffer package.

It is generated from these files:
	event/event.proto

It has these top-level messages:
	RawEventCreatedPayload
	UnsubscribeMessage
	RawEvents
	By
	RawEvent
	Subscription
	AllType
	ListEventsRequest
	UserEvent
*/
package event

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"
import account "bitbucket.org/subiz/header/account"
import conversation "bitbucket.org/subiz/header/conversation"
import content "bitbucket.org/subiz/header/content"
import user "bitbucket.org/subiz/header/user"

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
	Event_Sub Event = 0
	// Out_EventCreated = 0;
	Event_RawEventCreated  Event = 3
	Event_Subscribe        Event = 4
	Event_SubscribeReply   Event = 6
	Event_Unsubscribe      Event = 5
	Event_UnsubscribeReply Event = 7
)

var Event_name = map[int32]string{
	0: "Sub",
	3: "RawEventCreated",
	4: "Subscribe",
	6: "SubscribeReply",
	5: "Unsubscribe",
	7: "UnsubscribeReply",
}
var Event_value = map[string]int32{
	"Sub":              0,
	"RawEventCreated":  3,
	"Subscribe":        4,
	"SubscribeReply":   6,
	"Unsubscribe":      5,
	"UnsubscribeReply": 7,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Type int32

const (
	Type_c                          Type = 0
	Type_conversation_updated       Type = 9
	Type_message_sent               Type = 10
	Type_conversation_state_updated Type = 11
	// conversation_started = 1;
	Type_conversation_joined Type = 2
	// conversation_message = 3;
	Type_conversation_left Type = 4
	// conversation_closed = 5;
	Type_conversation_tagged   Type = 6
	Type_conversation_untagged Type = 7
	// conversation_waiting = 8;
	Type_channel_deintegrated       Type = 20
	Type_channel_integrated         Type = 21
	Type_message_seen               Type = 30
	Type_message_acked              Type = 31
	Type_message_received           Type = 32
	Type_conversation_member_typing Type = 33
	Type_conversation_postbacked    Type = 34
	Type_conversation_unassigned    Type = 35
	Type_conversation_assigned      Type = 36
)

var Type_name = map[int32]string{
	0:  "c",
	9:  "conversation_updated",
	10: "message_sent",
	11: "conversation_state_updated",
	2:  "conversation_joined",
	4:  "conversation_left",
	6:  "conversation_tagged",
	7:  "conversation_untagged",
	20: "channel_deintegrated",
	21: "channel_integrated",
	30: "message_seen",
	31: "message_acked",
	32: "message_received",
	33: "conversation_member_typing",
	34: "conversation_postbacked",
	35: "conversation_unassigned",
	36: "conversation_assigned",
}
var Type_value = map[string]int32{
	"c": 0,
	"conversation_updated":       9,
	"message_sent":               10,
	"conversation_state_updated": 11,
	"conversation_joined":        2,
	"conversation_left":          4,
	"conversation_tagged":        6,
	"conversation_untagged":      7,
	"channel_deintegrated":       20,
	"channel_integrated":         21,
	"message_seen":               30,
	"message_acked":              31,
	"message_received":           32,
	"conversation_member_typing": 33,
	"conversation_postbacked":    34,
	"conversation_unassigned":    35,
	"conversation_assigned":      36,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SubscriberType int32

const (
	SubscriberType_user    SubscriberType = 0
	SubscriberType_channel SubscriberType = 1
)

var SubscriberType_name = map[int32]string{
	0: "user",
	1: "channel",
}
var SubscriberType_value = map[string]int32{
	"user":    0,
	"channel": 1,
}

func (x SubscriberType) String() string {
	return proto.EnumName(SubscriberType_name, int32(x))
}
func (SubscriberType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SubPrefix int32

const (
	SubPrefix_Webhook   SubPrefix = 0
	SubPrefix_Websocket SubPrefix = 1
)

var SubPrefix_name = map[int32]string{
	0: "Webhook",
	1: "Websocket",
}
var SubPrefix_value = map[string]int32{
	"Webhook":   0,
	"Websocket": 1,
}

func (x SubPrefix) String() string {
	return proto.EnumName(SubPrefix_name, int32(x))
}
func (SubPrefix) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Object int32

const (
	Object_none         Object = 0
	Object_account      Object = 1
	Object_agent        Object = 2
	Object_message      Object = 3
	Object_conversation Object = 4
	Object_integration  Object = 30
	Object_postback     Object = 5
	Object_content      Object = 6
	Object_topic        Object = 7
)

var Object_name = map[int32]string{
	0:  "none",
	1:  "account",
	2:  "agent",
	3:  "message",
	4:  "conversation",
	30: "integration",
	5:  "postback",
	6:  "content",
	7:  "topic",
}
var Object_value = map[string]int32{
	"none":         0,
	"account":      1,
	"agent":        2,
	"message":      3,
	"conversation": 4,
	"integration":  30,
	"postback":     5,
	"content":      6,
	"topic":        7,
}

func (x Object) String() string {
	return proto.EnumName(Object_name, int32(x))
}
func (Object) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RawEventCreatedPayload struct {
	Ctx     *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	SubId   string          `protobuf:"bytes,3,opt,name=sub_id,json=subId" json:"sub_id,omitempty"`
	Topic   string          `protobuf:"bytes,4,opt,name=topic" json:"topic,omitempty"`
	Payload string          `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
}

func (m *RawEventCreatedPayload) Reset()                    { *m = RawEventCreatedPayload{} }
func (m *RawEventCreatedPayload) String() string            { return proto.CompactTextString(m) }
func (*RawEventCreatedPayload) ProtoMessage()               {}
func (*RawEventCreatedPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RawEventCreatedPayload) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *RawEventCreatedPayload) GetSubId() string {
	if m != nil {
		return m.SubId
	}
	return ""
}

func (m *RawEventCreatedPayload) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *RawEventCreatedPayload) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type UnsubscribeMessage struct {
	Topic string `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	SubId string `protobuf:"bytes,4,opt,name=sub_id,json=subId" json:"sub_id,omitempty"`
}

func (m *UnsubscribeMessage) Reset()                    { *m = UnsubscribeMessage{} }
func (m *UnsubscribeMessage) String() string            { return proto.CompactTextString(m) }
func (*UnsubscribeMessage) ProtoMessage()               {}
func (*UnsubscribeMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UnsubscribeMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *UnsubscribeMessage) GetSubId() string {
	if m != nil {
		return m.SubId
	}
	return ""
}

type RawEvents struct {
	Ctx    *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Events []*RawEvent     `protobuf:"bytes,4,rep,name=events" json:"events,omitempty"`
	Total  int64           `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	Anchor string          `protobuf:"bytes,3,opt,name=anchor" json:"anchor,omitempty"`
}

func (m *RawEvents) Reset()                    { *m = RawEvents{} }
func (m *RawEvents) String() string            { return proto.CompactTextString(m) }
func (*RawEvents) ProtoMessage()               {}
func (*RawEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RawEvents) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *RawEvents) GetEvents() []*RawEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *RawEvents) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RawEvents) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

type By struct {
	Device   *common.Device `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
	Id       string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	ClientId string         `protobuf:"bytes,4,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
}

func (m *By) Reset()                    { *m = By{} }
func (m *By) String() string            { return proto.CompactTextString(m) }
func (*By) ProtoMessage()               {}
func (*By) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *By) GetDevice() *common.Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *By) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *By) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type RawEvent struct {
	Ctx            *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Id             string          `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	AccountId      string          `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Created        int64           `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Type           string          `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	Topics         []string        `protobuf:"bytes,11,rep,name=topics" json:"topics,omitempty"`
	By             *By             `protobuf:"bytes,50,opt,name=by" json:"by,omitempty"`
	Object         string          `protobuf:"bytes,13,opt,name=object" json:"object,omitempty"`
	ConversationId string          `protobuf:"bytes,15,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	Data           *RawEvent_Data  `protobuf:"bytes,20,opt,name=data" json:"data,omitempty"`
}

func (m *RawEvent) Reset()                    { *m = RawEvent{} }
func (m *RawEvent) String() string            { return proto.CompactTextString(m) }
func (*RawEvent) ProtoMessage()               {}
func (*RawEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RawEvent) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *RawEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RawEvent) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *RawEvent) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RawEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RawEvent) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *RawEvent) GetBy() *By {
	if m != nil {
		return m.By
	}
	return nil
}

func (m *RawEvent) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *RawEvent) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *RawEvent) GetData() *RawEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type RawEvent_Data struct {
	Account      *account.Account           `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Agent        *account.Agent             `protobuf:"bytes,2,opt,name=agent" json:"agent,omitempty"`
	Message      *conversation.Message      `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Conversation *conversation.Conversation `protobuf:"bytes,4,opt,name=conversation" json:"conversation,omitempty"`
	Postback     *conversation.Postback     `protobuf:"bytes,5,opt,name=postback" json:"postback,omitempty"`
	Content      *content.Content           `protobuf:"bytes,6,opt,name=content" json:"content,omitempty"`
	Topic        *user.Topic                `protobuf:"bytes,7,opt,name=topic" json:"topic,omitempty"`
	Presence     *user.Presence             `protobuf:"bytes,8,opt,name=presence" json:"presence,omitempty"`
	Previewing   *user.Previewing           `protobuf:"bytes,9,opt,name=previewing" json:"previewing,omitempty"`
	User         *user.User                 `protobuf:"bytes,10,opt,name=user" json:"user,omitempty"`
}

func (m *RawEvent_Data) Reset()                    { *m = RawEvent_Data{} }
func (m *RawEvent_Data) String() string            { return proto.CompactTextString(m) }
func (*RawEvent_Data) ProtoMessage()               {}
func (*RawEvent_Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *RawEvent_Data) GetAccount() *account.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *RawEvent_Data) GetAgent() *account.Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *RawEvent_Data) GetMessage() *conversation.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RawEvent_Data) GetConversation() *conversation.Conversation {
	if m != nil {
		return m.Conversation
	}
	return nil
}

func (m *RawEvent_Data) GetPostback() *conversation.Postback {
	if m != nil {
		return m.Postback
	}
	return nil
}

func (m *RawEvent_Data) GetContent() *content.Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *RawEvent_Data) GetTopic() *user.Topic {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *RawEvent_Data) GetPresence() *user.Presence {
	if m != nil {
		return m.Presence
	}
	return nil
}

func (m *RawEvent_Data) GetPreviewing() *user.Previewing {
	if m != nil {
		return m.Previewing
	}
	return nil
}

func (m *RawEvent_Data) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

type Subscription struct {
	Ctx   *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Topic string          `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	// optional SubscribeScope scope = 5;
	// optional string account_id = 6;
	SubId string `protobuf:"bytes,7,opt,name=sub_id,json=subId" json:"sub_id,omitempty"`
	// optional SubscriberType subscriber_type = 8;
	// optional string subscriber_id = 9;
	TargetTopic string `protobuf:"bytes,10,opt,name=target_topic,json=targetTopic" json:"target_topic,omitempty"`
	TargetKey   string `protobuf:"bytes,11,opt,name=target_key,json=targetKey" json:"target_key,omitempty"`
	// optional string target_partition = 13;
	// time to life in seconds
	Ttls int64 `protobuf:"varint,12,opt,name=ttls" json:"ttls,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Subscription) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Subscription) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Subscription) GetSubId() string {
	if m != nil {
		return m.SubId
	}
	return ""
}

func (m *Subscription) GetTargetTopic() string {
	if m != nil {
		return m.TargetTopic
	}
	return ""
}

func (m *Subscription) GetTargetKey() string {
	if m != nil {
		return m.TargetKey
	}
	return ""
}

func (m *Subscription) GetTtls() int64 {
	if m != nil {
		return m.Ttls
	}
	return 0
}

type AllType struct {
	O         Object             `protobuf:"varint,2,opt,name=o,enum=event.Object" json:"o,omitempty"`
	E         *RawEvent          `protobuf:"bytes,3,opt,name=e" json:"e,omitempty"`
	B         *By                `protobuf:"bytes,4,opt,name=b" json:"b,omitempty"`
	Es        *RawEvents         `protobuf:"bytes,5,opt,name=es" json:"es,omitempty"`
	Ler       *ListEventsRequest `protobuf:"bytes,6,opt,name=ler" json:"ler,omitempty"`
	Userevent *UserEvent         `protobuf:"bytes,7,opt,name=userevent" json:"userevent,omitempty"`
}

func (m *AllType) Reset()                    { *m = AllType{} }
func (m *AllType) String() string            { return proto.CompactTextString(m) }
func (*AllType) ProtoMessage()               {}
func (*AllType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AllType) GetO() Object {
	if m != nil {
		return m.O
	}
	return Object_none
}

func (m *AllType) GetE() *RawEvent {
	if m != nil {
		return m.E
	}
	return nil
}

func (m *AllType) GetB() *By {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *AllType) GetEs() *RawEvents {
	if m != nil {
		return m.Es
	}
	return nil
}

func (m *AllType) GetLer() *ListEventsRequest {
	if m != nil {
		return m.Ler
	}
	return nil
}

func (m *AllType) GetUserevent() *UserEvent {
	if m != nil {
		return m.Userevent
	}
	return nil
}

type ListEventsRequest struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId    string          `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Query     string          `protobuf:"bytes,4,opt,name=query" json:"query,omitempty"`
	Anchor    string          `protobuf:"bytes,5,opt,name=anchor" json:"anchor,omitempty"`
	Limit     int32           `protobuf:"varint,6,opt,name=limit" json:"limit,omitempty"`
	Category  string          `protobuf:"bytes,9,opt,name=category" json:"category,omitempty"`
}

func (m *ListEventsRequest) Reset()                    { *m = ListEventsRequest{} }
func (m *ListEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEventsRequest) ProtoMessage()               {}
func (*ListEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListEventsRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ListEventsRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *ListEventsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListEventsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ListEventsRequest) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

func (m *ListEventsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListEventsRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type UserEvent struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId    string          `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Event     *RawEvent       `protobuf:"bytes,5,opt,name=event" json:"event,omitempty"`
}

func (m *UserEvent) Reset()                    { *m = UserEvent{} }
func (m *UserEvent) String() string            { return proto.CompactTextString(m) }
func (*UserEvent) ProtoMessage()               {}
func (*UserEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UserEvent) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *UserEvent) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *UserEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserEvent) GetEvent() *RawEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*RawEventCreatedPayload)(nil), "event.RawEventCreatedPayload")
	proto.RegisterType((*UnsubscribeMessage)(nil), "event.UnsubscribeMessage")
	proto.RegisterType((*RawEvents)(nil), "event.RawEvents")
	proto.RegisterType((*By)(nil), "event.By")
	proto.RegisterType((*RawEvent)(nil), "event.RawEvent")
	proto.RegisterType((*RawEvent_Data)(nil), "event.RawEvent.Data")
	proto.RegisterType((*Subscription)(nil), "event.Subscription")
	proto.RegisterType((*AllType)(nil), "event.AllType")
	proto.RegisterType((*ListEventsRequest)(nil), "event.ListEventsRequest")
	proto.RegisterType((*UserEvent)(nil), "event.UserEvent")
	proto.RegisterEnum("event.Event", Event_name, Event_value)
	proto.RegisterEnum("event.Type", Type_name, Type_value)
	proto.RegisterEnum("event.SubscriberType", SubscriberType_name, SubscriberType_value)
	proto.RegisterEnum("event.SubPrefix", SubPrefix_name, SubPrefix_value)
	proto.RegisterEnum("event.Object", Object_name, Object_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MyService service

type MyServiceClient interface {
	Do(ctx context.Context, in *AllType, opts ...grpc.CallOption) (*AllType, error)
}

type myServiceClient struct {
	cc *grpc.ClientConn
}

func NewMyServiceClient(cc *grpc.ClientConn) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) Do(ctx context.Context, in *AllType, opts ...grpc.CallOption) (*AllType, error) {
	out := new(AllType)
	err := grpc.Invoke(ctx, "/event.MyService/Do", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyService service

type MyServiceServer interface {
	Do(context.Context, *AllType) (*AllType, error)
}

func RegisterMyServiceServer(s *grpc.Server, srv MyServiceServer) {
	s.RegisterService(&_MyService_serviceDesc, srv)
}

func _MyService_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.MyService/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).Do(ctx, req.(*AllType))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _MyService_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event/event.proto",
}

func init() { proto.RegisterFile("event/event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xdb, 0x6e, 0x1b, 0x37,
	0x13, 0xf6, 0x4a, 0x94, 0x44, 0x8d, 0x7c, 0xa0, 0xe9, 0xd3, 0x5a, 0x8e, 0x1d, 0x59, 0x71, 0x6c,
	0x47, 0x3f, 0x20, 0x07, 0xca, 0xdd, 0x7f, 0xf1, 0x03, 0x89, 0xf3, 0x5f, 0x78, 0xdb, 0xa0, 0x06,
	0x9d, 0x20, 0xe8, 0x95, 0xb1, 0x07, 0x5a, 0xde, 0x78, 0xb5, 0xab, 0xec, 0x72, 0x9d, 0xa8, 0x57,
	0x45, 0x81, 0x5e, 0xf6, 0x15, 0xda, 0xf7, 0x29, 0x8a, 0x3e, 0x41, 0x1f, 0xa6, 0xe0, 0x61, 0xa5,
	0x95, 0x0d, 0xb8, 0xee, 0x8d, 0x96, 0x33, 0xf3, 0x71, 0x38, 0x33, 0x1c, 0x7e, 0xa4, 0x60, 0x95,
	0xdf, 0xf2, 0x58, 0x9c, 0xa8, 0xdf, 0xfe, 0x38, 0x4d, 0x44, 0x42, 0x6b, 0x4a, 0x68, 0xf7, 0xbd,
	0x50, 0x78, 0xb9, 0x7f, 0xc3, 0x45, 0x3f, 0x49, 0x87, 0x27, 0x59, 0xee, 0x85, 0x3f, 0x9c, 0x5c,
	0x73, 0x37, 0xe0, 0xe9, 0x89, 0x9f, 0x8c, 0x46, 0x49, 0x6c, 0x3e, 0x7a, 0x5a, 0xfb, 0xe5, 0x03,
	0x78, 0xd7, 0xf7, 0x93, 0x3c, 0x16, 0xc5, 0xd7, 0xcc, 0xf8, 0xef, 0x83, 0x2b, 0xc4, 0xb7, 0x3c,
	0xcd, 0x5c, 0x11, 0xaa, 0x75, 0x66, 0xc2, 0x23, 0x56, 0xf3, 0x93, 0x58, 0xc8, 0xa4, 0xcc, 0xd7,
	0xcc, 0xe8, 0x3d, 0x30, 0x23, 0xcf, 0xcc, 0x8f, 0xc6, 0x76, 0x7f, 0xb2, 0x60, 0x93, 0xb9, 0x5f,
	0xfe, 0x2f, 0x0b, 0x71, 0x9a, 0x72, 0x57, 0xf0, 0xe0, 0xdc, 0x9d, 0x44, 0x89, 0x1b, 0xd0, 0x7d,
	0xa8, 0xfa, 0xe2, 0xab, 0x6d, 0x75, 0xac, 0xe3, 0xd6, 0x60, 0xa5, 0x6f, 0x4a, 0x70, 0x2a, 0x97,
	0xfa, 0x2a, 0x98, 0xb4, 0xd1, 0x0d, 0xa8, 0x67, 0xb9, 0x77, 0x19, 0x06, 0x76, 0xb5, 0x63, 0x1d,
	0x37, 0x59, 0x2d, 0xcb, 0xbd, 0xb3, 0x80, 0xae, 0x43, 0x4d, 0x24, 0xe3, 0xd0, 0xb7, 0x91, 0xd6,
	0x2a, 0x81, 0xda, 0xd0, 0x18, 0x6b, 0xd7, 0x76, 0x4d, 0xe9, 0x0b, 0xb1, 0xfb, 0x1a, 0xe8, 0x87,
	0x38, 0xcb, 0xbd, 0xcc, 0x4f, 0x43, 0x8f, 0xbf, 0xe3, 0x59, 0xe6, 0x0e, 0xf9, 0xcc, 0x4b, 0xb5,
	0xec, 0x65, 0xb6, 0x24, 0x2a, 0x2d, 0xd9, 0xfd, 0xd9, 0x82, 0x66, 0x91, 0x47, 0xf6, 0x98, 0xd0,
	0x8f, 0xa0, 0xae, 0x76, 0x3f, 0xb3, 0x51, 0xa7, 0xaa, 0x50, 0xba, 0x33, 0x0a, 0x27, 0xcc, 0x98,
	0x75, 0x18, 0xc2, 0x8d, 0xec, 0x4a, 0xc7, 0x3a, 0xae, 0x32, 0x2d, 0xd0, 0x4d, 0xa8, 0xbb, 0xb1,
	0x7f, 0x9d, 0xa4, 0x26, 0x3a, 0x23, 0x75, 0xbf, 0x87, 0xca, 0x9b, 0x09, 0x3d, 0x84, 0x7a, 0xc0,
	0x6f, 0x43, 0x9f, 0xab, 0x49, 0xad, 0xc1, 0x72, 0x11, 0xc2, 0x5b, 0xa5, 0x65, 0xc6, 0x4a, 0x97,
	0xa1, 0x32, 0xad, 0x5d, 0x25, 0x0c, 0xe8, 0x0e, 0x34, 0xfd, 0x28, 0xe4, 0xb1, 0x98, 0xe5, 0x87,
	0xb5, 0xe2, 0x2c, 0xe8, 0xfe, 0xda, 0x04, 0x5c, 0x44, 0xf7, 0x98, 0x0c, 0xef, 0x3a, 0xdf, 0x05,
	0x30, 0x5d, 0x39, 0xf3, 0xde, 0x34, 0x9a, 0xb3, 0x40, 0x6e, 0x8f, 0xaf, 0x1b, 0xc0, 0xc6, 0x2a,
	0xd3, 0x42, 0xa4, 0x14, 0x90, 0x98, 0x8c, 0xb9, 0xdd, 0x54, 0x53, 0xd4, 0x58, 0xe6, 0xaf, 0xf6,
	0x23, 0xb3, 0x5b, 0x9d, 0xaa, 0xcc, 0x5f, 0x4b, 0x74, 0x1b, 0x2a, 0xde, 0xc4, 0x1e, 0xa8, 0xb0,
	0x9a, 0xa6, 0xa4, 0x6f, 0x26, 0xac, 0xe2, 0x4d, 0xe4, 0x94, 0xc4, 0xfb, 0xc4, 0x7d, 0x61, 0x2f,
	0xe9, 0x92, 0x69, 0x89, 0x1e, 0xc1, 0x4a, 0xb9, 0xed, 0x65, 0x70, 0x2b, 0x0a, 0xb0, 0x5c, 0x56,
	0x9f, 0x05, 0xf4, 0x18, 0x50, 0xe0, 0x0a, 0xd7, 0x5e, 0x57, 0xde, 0xd7, 0xef, 0x6c, 0x58, 0xff,
	0xad, 0x2b, 0x5c, 0xa6, 0x10, 0xed, 0xdf, 0xab, 0x80, 0xa4, 0x48, 0x7b, 0xd0, 0x30, 0x19, 0x9a,
	0x52, 0x91, 0x7e, 0x71, 0x32, 0x5f, 0xeb, 0x2f, 0x2b, 0x00, 0xf4, 0x00, 0x6a, 0xee, 0x90, 0xc7,
	0x62, 0xba, 0x67, 0x53, 0xa4, 0xd4, 0x32, 0x6d, 0xa4, 0x27, 0xd0, 0x18, 0xe9, 0x06, 0x55, 0xa5,
	0x6d, 0x0d, 0x36, 0xfa, 0x73, 0x87, 0xd6, 0x74, 0x2f, 0x2b, 0x50, 0xf4, 0x7f, 0xb0, 0x58, 0x06,
	0xa8, 0xc2, 0xb7, 0x06, 0xed, 0xf9, 0x59, 0xa7, 0x25, 0x81, 0xcd, 0xe1, 0xe9, 0x00, 0xf0, 0x38,
	0xc9, 0x84, 0xe7, 0xfa, 0x37, 0xea, 0xdc, 0xb4, 0x06, 0x9b, 0xf3, 0x73, 0xcf, 0x8d, 0x95, 0x4d,
	0x71, 0x32, 0x6d, 0x43, 0x09, 0x76, 0xdd, 0xa4, 0x5d, 0x50, 0xc4, 0xa9, 0xfe, 0xb2, 0x02, 0x40,
	0xf7, 0x8b, 0x63, 0xd6, 0x50, 0xc8, 0x56, 0x5f, 0xb1, 0xc3, 0x7b, 0xa9, 0x2a, 0xce, 0x5c, 0x0f,
	0xf0, 0x38, 0xe5, 0x19, 0x8f, 0x7d, 0xae, 0x7a, 0x43, 0x16, 0x47, 0xa1, 0xce, 0x8d, 0x96, 0x4d,
	0xed, 0xf4, 0x25, 0xc0, 0x38, 0xe5, 0xb7, 0x21, 0xff, 0x12, 0xc6, 0x43, 0xd5, 0x32, 0x72, 0xf5,
	0x02, 0x6d, 0xf4, 0xac, 0x84, 0xa1, 0x7b, 0x80, 0xa4, 0xd9, 0x06, 0x85, 0x05, 0x8d, 0xfd, 0x90,
	0xf1, 0x94, 0x29, 0xbd, 0x83, 0x30, 0x90, 0x96, 0x83, 0xf0, 0x22, 0x59, 0x72, 0x10, 0x26, 0x64,
	0xd5, 0x41, 0x78, 0x95, 0x50, 0x07, 0x61, 0x4a, 0xd6, 0x1c, 0x84, 0xd7, 0xc8, 0xba, 0x83, 0xf0,
	0x06, 0xd9, 0x74, 0x10, 0xde, 0x24, 0x5b, 0x0e, 0xc2, 0x5b, 0xc4, 0x76, 0x10, 0xb6, 0xc9, 0xb6,
	0x83, 0xf0, 0x36, 0x69, 0x3b, 0x08, 0xb7, 0xc9, 0x8e, 0x83, 0xf0, 0x0e, 0x79, 0xe2, 0x20, 0xfc,
	0x84, 0xec, 0x3a, 0x08, 0xef, 0x92, 0x3d, 0x07, 0xe1, 0x3d, 0xf2, 0xd4, 0x41, 0xf8, 0x29, 0xe9,
	0x38, 0x08, 0x77, 0xc8, 0xbe, 0x83, 0xf0, 0x3e, 0xe9, 0x3a, 0x08, 0x3f, 0x23, 0x07, 0x0e, 0xc2,
	0x07, 0xe4, 0xb9, 0x83, 0xf0, 0x73, 0x72, 0xe8, 0x20, 0x7c, 0x48, 0x8e, 0x1c, 0x84, 0x8f, 0xc8,
	0xb1, 0x83, 0xf0, 0x31, 0x79, 0xe1, 0x20, 0xfc, 0x82, 0xf4, 0x1c, 0x84, 0x7b, 0xe4, 0x3f, 0x0c,
	0xfc, 0x6b, 0x37, 0x8e, 0x79, 0x74, 0x19, 0x06, 0x0c, 0xbc, 0xc9, 0xa5, 0x0c, 0x5b, 0x8e, 0x6b,
	0x99, 0x70, 0x05, 0x67, 0x60, 0x5a, 0xe1, 0x52, 0x24, 0x0c, 0x8f, 0xe5, 0x20, 0x4f, 0x23, 0x06,
	0x6a, 0x24, 0x42, 0x11, 0x71, 0x46, 0xbc, 0x34, 0xf9, 0x22, 0x27, 0x45, 0x6e, 0x3c, 0xcc, 0x65,
	0xfb, 0xb4, 0x34, 0x43, 0x5c, 0xca, 0xc3, 0xc6, 0x1a, 0x85, 0xc3, 0xe6, 0xa7, 0x24, 0x8c, 0xf5,
	0xb0, 0x65, 0x86, 0x0a, 0xd0, 0x8c, 0xb8, 0x7b, 0x6b, 0xf4, 0xae, 0x10, 0xae, 0x7f, 0x3d, 0x92,
	0xdc, 0xc5, 0x90, 0x62, 0x83, 0xfa, 0x55, 0x92, 0x8e, 0x5c, 0x21, 0x7d, 0x46, 0xe1, 0x2d, 0x4f,
	0x43, 0x1e, 0x64, 0xac, 0x7e, 0x15, 0xf2, 0x28, 0xc8, 0x18, 0xf6, 0x93, 0xd1, 0x38, 0x17, 0x3c,
	0x98, 0x1e, 0x04, 0xd3, 0xe9, 0xd3, 0x0e, 0x9e, 0xef, 0xc7, 0x59, 0x97, 0x4d, 0x5b, 0xc8, 0xb4,
	0x49, 0xf7, 0x0f, 0x0b, 0x16, 0x2f, 0x34, 0x8b, 0x8f, 0x55, 0xeb, 0x3e, 0x82, 0xa4, 0xfe, 0x89,
	0xe4, 0x1b, 0xe5, 0x7b, 0x65, 0x1f, 0x16, 0x85, 0x9b, 0x0e, 0xb9, 0xb8, 0xd4, 0x73, 0x40, 0x19,
	0x5b, 0x5a, 0xa7, 0x3a, 0x56, 0x92, 0x9c, 0x81, 0xdc, 0xf0, 0x89, 0xdd, 0xd2, 0x24, 0xa7, 0x35,
	0xdf, 0xf0, 0x89, 0xa2, 0x32, 0x11, 0x65, 0xf6, 0xa2, 0x62, 0x38, 0x35, 0x76, 0x10, 0x5e, 0x22,
	0xcb, 0x8c, 0x98, 0x69, 0x63, 0x37, 0x15, 0xa1, 0x8c, 0xbe, 0xfb, 0x97, 0x05, 0x8d, 0xd7, 0x51,
	0xf4, 0x5e, 0xd2, 0xdd, 0x0e, 0x58, 0x89, 0xe2, 0x85, 0xe5, 0xc1, 0x92, 0xe1, 0x9d, 0xef, 0x14,
	0x7b, 0x31, 0x2b, 0xa1, 0xbb, 0x60, 0x15, 0x64, 0x70, 0xef, 0x16, 0xb1, 0x38, 0xdd, 0x02, 0xcb,
	0x33, 0xa7, 0xbe, 0xc4, 0x88, 0x96, 0x47, 0x3b, 0x50, 0xe1, 0x99, 0x39, 0xd3, 0xe4, 0xce, 0xc4,
	0x8c, 0x55, 0x78, 0x46, 0x7b, 0x50, 0x8d, 0x78, 0x6a, 0xce, 0xb0, 0x6d, 0x20, 0xdf, 0x86, 0x99,
	0x30, 0x18, 0xfe, 0x39, 0xe7, 0x99, 0x60, 0x12, 0x44, 0xfb, 0xd0, 0x94, 0x6d, 0xa2, 0x30, 0xe6,
	0x2c, 0x17, 0x4e, 0xe5, 0x61, 0xd2, 0xe1, 0xcc, 0x20, 0xdd, 0x3f, 0x2d, 0x58, 0xbd, 0xe7, 0xea,
	0x31, 0x5b, 0x36, 0x7f, 0x8f, 0x54, 0xee, 0xde, 0x23, 0x5b, 0x50, 0xb4, 0x6b, 0x71, 0x35, 0x4a,
	0x51, 0xbf, 0x0a, 0x3e, 0xe7, 0x3c, 0x9d, 0x14, 0x17, 0xb7, 0x12, 0x4a, 0x17, 0x69, 0xad, 0x7c,
	0x91, 0x4a, 0x74, 0x14, 0x8e, 0x42, 0x4d, 0x60, 0x35, 0xa6, 0x05, 0xda, 0x06, 0xec, 0xbb, 0x82,
	0x0f, 0x93, 0x74, 0x62, 0xae, 0xa3, 0xa9, 0xdc, 0xfd, 0xc5, 0x82, 0xe6, 0x34, 0xd3, 0x7f, 0x9f,
	0x48, 0xf5, 0x81, 0x44, 0xd0, 0x5c, 0x22, 0xcf, 0x41, 0x3f, 0x1c, 0xcd, 0xd6, 0xdd, 0xdb, 0x73,
	0x6d, 0xed, 0x25, 0x50, 0xd3, 0xa1, 0x34, 0xa0, 0x7a, 0x91, 0x7b, 0x64, 0x81, 0xae, 0xc1, 0xca,
	0x9d, 0xb7, 0x16, 0xa9, 0xd2, 0x25, 0x68, 0x5e, 0x14, 0x4f, 0x1f, 0x82, 0x28, 0x85, 0xe5, 0xa9,
	0xc8, 0xf8, 0x38, 0x9a, 0x90, 0x3a, 0x5d, 0x81, 0x56, 0xe9, 0x7d, 0x44, 0x6a, 0x74, 0x1d, 0x48,
	0x49, 0xa1, 0x61, 0x8d, 0xde, 0x6f, 0x55, 0x40, 0xaa, 0x5b, 0x6b, 0x60, 0xf9, 0x64, 0x81, 0xda,
	0xb0, 0x3e, 0x77, 0xb1, 0xe6, 0xe3, 0x40, 0xad, 0xd9, 0xa4, 0x04, 0x16, 0x0b, 0x4e, 0xca, 0x78,
	0x2c, 0x08, 0xd0, 0x3d, 0x68, 0xcf, 0x61, 0x15, 0x73, 0x4d, 0x67, 0xb4, 0xe8, 0x16, 0xac, 0xcd,
	0xd9, 0x15, 0xfb, 0x04, 0xa4, 0x42, 0x37, 0x60, 0x75, 0xce, 0x10, 0xf1, 0x2b, 0x41, 0xd0, 0x3d,
	0xbc, 0x70, 0x87, 0x43, 0x1e, 0x90, 0x3a, 0xdd, 0x86, 0x8d, 0xf9, 0xa0, 0x62, 0x63, 0x6a, 0xa8,
	0x78, 0x0d, 0x91, 0x06, 0x3c, 0x8c, 0x05, 0x1f, 0xa6, 0x6a, 0xf5, 0x75, 0xba, 0x09, 0x74, 0x4a,
	0xb1, 0x33, 0xfd, 0xc6, 0x7c, 0x1e, 0x3c, 0x26, 0x7b, 0x74, 0x15, 0x96, 0x0a, 0x8d, 0xeb, 0xdf,
	0xf0, 0x80, 0x3c, 0x95, 0xc5, 0x2a, 0x54, 0x29, 0xf7, 0x79, 0x78, 0xcb, 0x03, 0xd2, 0xb9, 0x97,
	0xf0, 0x88, 0x8f, 0x3c, 0x4d, 0xa7, 0x61, 0x3c, 0x24, 0xfb, 0x74, 0x07, 0xb6, 0xe6, 0xec, 0x05,
	0xeb, 0xf1, 0x80, 0x74, 0xef, 0x19, 0xf3, 0xd8, 0xcd, 0xb2, 0x70, 0x28, 0x2b, 0xf2, 0xec, 0x5e,
	0x86, 0x53, 0xd3, 0x41, 0xef, 0xa8, 0xb4, 0xb9, 0xa9, 0xda, 0x2a, 0xac, 0x2f, 0x3f, 0xb2, 0x40,
	0x5b, 0xd0, 0x30, 0x39, 0x12, 0xab, 0x77, 0xa4, 0x9a, 0xe2, 0x3c, 0xe5, 0x57, 0xe1, 0x57, 0x69,
	0xf9, 0xc8, 0xbd, 0xeb, 0x24, 0xb9, 0x21, 0x0b, 0xb2, 0x5d, 0x3e, 0x72, 0x2f, 0x4b, 0xe4, 0xf3,
	0x9e, 0x58, 0xbd, 0x1f, 0x2d, 0xa8, 0x6b, 0x26, 0x92, 0xae, 0xe2, 0x24, 0xe6, 0xda, 0x95, 0x69,
	0x63, 0x62, 0xd1, 0xa6, 0x79, 0xd6, 0x90, 0x8a, 0xd4, 0x9b, 0x4a, 0x90, 0xaa, 0xac, 0x5d, 0x39,
	0x4c, 0x82, 0x64, 0x9b, 0x15, 0xd5, 0x95, 0x8a, 0x3d, 0xba, 0x38, 0x7b, 0x7a, 0x90, 0x9a, 0x0a,
	0x50, 0x13, 0x3e, 0xa9, 0x4b, 0xaf, 0x8a, 0x83, 0x49, 0x63, 0xf0, 0x0a, 0x9a, 0xef, 0x26, 0x17,
	0x3c, 0x55, 0x2f, 0xda, 0x43, 0xa8, 0xbc, 0x4d, 0xe8, 0xb2, 0x39, 0x12, 0x86, 0x3e, 0xdb, 0x77,
	0xe4, 0xee, 0x82, 0x57, 0x57, 0x7f, 0x3f, 0x5e, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xf6,
	0x3b, 0x0a, 0x96, 0x0d, 0x00, 0x00,
}
