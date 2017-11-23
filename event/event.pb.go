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
	Event_ZZZ Event = 0
	// Out_EventCreated = 0;
	Event_RawEventCreated  Event = 3
	Event_Subscribe        Event = 4
	Event_SubscribeReply   Event = 6
	Event_Unsubscribe      Event = 5
	Event_UnsubscribeReply Event = 7
)

var Event_name = map[int32]string{
	0: "ZZZ",
	3: "RawEventCreated",
	4: "Subscribe",
	6: "SubscribeReply",
	5: "Unsubscribe",
	7: "UnsubscribeReply",
}
var Event_value = map[string]int32{
	"ZZZ":              0,
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

func init() {
	proto.RegisterType((*RawEventCreatedPayload)(nil), "event.RawEventCreatedPayload")
	proto.RegisterType((*UnsubscribeMessage)(nil), "event.UnsubscribeMessage")
	proto.RegisterType((*RawEvents)(nil), "event.RawEvents")
	proto.RegisterType((*By)(nil), "event.By")
	proto.RegisterType((*RawEvent)(nil), "event.RawEvent")
	proto.RegisterType((*RawEvent_Data)(nil), "event.RawEvent.Data")
	proto.RegisterType((*Subscription)(nil), "event.Subscription")
	proto.RegisterEnum("event.Event", Event_name, Event_value)
	proto.RegisterEnum("event.Type", Type_name, Type_value)
	proto.RegisterEnum("event.SubscriberType", SubscriberType_name, SubscriberType_value)
	proto.RegisterEnum("event.SubPrefix", SubPrefix_name, SubPrefix_value)
	proto.RegisterEnum("event.Object", Object_name, Object_value)
}

func init() { proto.RegisterFile("event/event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x6e, 0x1b, 0xb7,
	0x13, 0xb7, 0x24, 0x4a, 0xa2, 0x66, 0x65, 0x7b, 0x4c, 0xcb, 0xf6, 0x5a, 0x4e, 0x1c, 0x59, 0xff,
	0x24, 0x56, 0xf4, 0x07, 0xe4, 0xc2, 0xbd, 0xf5, 0x50, 0x20, 0x71, 0x7a, 0x30, 0x8b, 0xa2, 0x01,
	0x93, 0x22, 0x68, 0x2e, 0xc6, 0x7e, 0xd0, 0xf2, 0xc6, 0xd2, 0xae, 0xb0, 0x4b, 0x39, 0x51, 0x4f,
	0x45, 0x9f, 0xa4, 0x0f, 0x94, 0xf7, 0xe8, 0x6b, 0x14, 0xe4, 0x72, 0x25, 0x6d, 0x02, 0xb8, 0xbe,
	0x88, 0x9c, 0x99, 0xdf, 0x0c, 0xe7, 0x7b, 0x05, 0x3b, 0xf2, 0x4e, 0xc6, 0xea, 0xcc, 0xfc, 0x8e,
	0x66, 0x69, 0xa2, 0x12, 0x56, 0x37, 0x44, 0x77, 0xe4, 0x47, 0xca, 0x9f, 0x07, 0xb7, 0x52, 0x8d,
	0x92, 0x74, 0x7c, 0x96, 0xcd, 0xfd, 0xe8, 0x8f, 0xb3, 0x1b, 0xe9, 0x85, 0x32, 0x3d, 0x0b, 0x92,
	0xe9, 0x34, 0x89, 0xed, 0x91, 0xab, 0x75, 0xbf, 0xbb, 0x07, 0xef, 0x05, 0x41, 0x32, 0x8f, 0x55,
	0x71, 0x5a, 0x8d, 0x1f, 0xee, 0x7d, 0x21, 0xbe, 0x93, 0x69, 0xe6, 0xa9, 0xc8, 0xbc, 0xb3, 0x22,
	0x1e, 0xf0, 0x5a, 0x90, 0xc4, 0x4a, 0x07, 0x65, 0x4f, 0xab, 0x31, 0xbc, 0x47, 0x63, 0x9e, 0xd9,
	0x9f, 0x1c, 0xdb, 0xff, 0xab, 0x02, 0xfb, 0xc2, 0xfb, 0xf4, 0x93, 0x4e, 0xc4, 0x45, 0x2a, 0x3d,
	0x25, 0xc3, 0x37, 0xde, 0x62, 0x92, 0x78, 0x21, 0x3b, 0x81, 0x5a, 0xa0, 0x3e, 0xbb, 0x95, 0x5e,
	0x65, 0xe0, 0x9c, 0x6f, 0x8f, 0x6c, 0x0a, 0x2e, 0xf4, 0x53, 0x9f, 0x95, 0xd0, 0x32, 0xb6, 0x07,
	0x8d, 0x6c, 0xee, 0x5f, 0x45, 0xa1, 0x5b, 0xeb, 0x55, 0x06, 0x2d, 0x51, 0xcf, 0xe6, 0xfe, 0x65,
	0xc8, 0x3a, 0x50, 0x57, 0xc9, 0x2c, 0x0a, 0x5c, 0x92, 0x73, 0x0d, 0xc1, 0x5c, 0x68, 0xce, 0x72,
	0xd3, 0x6e, 0xdd, 0xf0, 0x0b, 0xb2, 0xff, 0x12, 0xd8, 0x6f, 0x71, 0x36, 0xf7, 0xb3, 0x20, 0x8d,
	0x7c, 0xf9, 0x8b, 0xcc, 0x32, 0x6f, 0x2c, 0x57, 0x56, 0x6a, 0xeb, 0x56, 0x56, 0x4f, 0x92, 0xb5,
	0x27, 0xfb, 0xef, 0xa1, 0x55, 0x84, 0x91, 0x3d, 0xc4, 0xf3, 0x53, 0x68, 0x98, 0xe2, 0x67, 0x2e,
	0xe9, 0xd5, 0x0c, 0x2a, 0x6f, 0x8c, 0xc2, 0x88, 0xb0, 0xe2, 0xfe, 0xef, 0x50, 0x7d, 0xb5, 0x60,
	0xcf, 0xa1, 0x11, 0xca, 0xbb, 0x28, 0x90, 0x6e, 0xd5, 0x18, 0xdd, 0x2a, 0x8c, 0xbe, 0x36, 0x5c,
	0x61, 0xa5, 0x6c, 0x0b, 0xaa, 0xcb, 0x64, 0x54, 0xa3, 0x90, 0x1d, 0x41, 0x2b, 0x98, 0x44, 0x32,
	0x56, 0x2b, 0x87, 0x69, 0xce, 0xb8, 0x0c, 0xfb, 0x7f, 0x53, 0xa0, 0xc5, 0x7b, 0x0f, 0xf1, 0xf9,
	0x6b, 0xe3, 0x8f, 0x01, 0x6c, 0x9b, 0xad, 0xac, 0xb7, 0x2c, 0xe7, 0x32, 0xd4, 0xf9, 0x0e, 0xf2,
	0x8a, 0xba, 0xb4, 0x57, 0x19, 0xd4, 0x44, 0x41, 0x32, 0x06, 0x44, 0x2d, 0x66, 0xd2, 0x6d, 0x19,
	0x15, 0x73, 0x67, 0xfb, 0xd0, 0x30, 0x09, 0xce, 0x5c, 0xa7, 0x57, 0x1b, 0xb4, 0x84, 0xa5, 0xd8,
	0x21, 0x54, 0xfd, 0x85, 0x7b, 0x6e, 0xdc, 0x6a, 0xd9, 0x24, 0xbd, 0x5a, 0x88, 0xaa, 0xbf, 0xd0,
	0x2a, 0x89, 0xff, 0x51, 0x06, 0xca, 0xdd, 0x34, 0x86, 0x2c, 0xc5, 0x4e, 0x61, 0x7b, 0xbd, 0x8f,
	0xb5, 0x73, 0xdb, 0x06, 0xb0, 0xb5, 0xce, 0xbe, 0x0c, 0xd9, 0x00, 0x48, 0xe8, 0x29, 0xcf, 0xed,
	0x18, 0xeb, 0x9d, 0xaf, 0x4a, 0x30, 0x7a, 0xed, 0x29, 0x4f, 0x18, 0x44, 0xf7, 0x4b, 0x15, 0x88,
	0x26, 0xd9, 0x10, 0x9a, 0x36, 0x42, 0x9b, 0x2a, 0x1c, 0x15, 0xa3, 0xf6, 0x32, 0x3f, 0x45, 0x01,
	0x60, 0x4f, 0xa1, 0xee, 0x8d, 0x65, 0xac, 0x96, 0x35, 0x5b, 0x22, 0x35, 0x57, 0xe4, 0x42, 0x76,
	0x06, 0xcd, 0x69, 0xde, 0x71, 0x26, 0xb5, 0xce, 0xf9, 0xde, 0xa8, 0x34, 0x85, 0xb6, 0x1d, 0x45,
	0x81, 0x62, 0x3f, 0x42, 0x7b, 0x1d, 0x60, 0x12, 0xef, 0x9c, 0x77, 0xcb, 0x5a, 0x17, 0x6b, 0x84,
	0x28, 0xe1, 0xd9, 0x39, 0xd0, 0x59, 0x92, 0x29, 0xdf, 0x0b, 0x6e, 0xcd, 0x20, 0x38, 0xe7, 0xfb,
	0x65, 0xdd, 0x37, 0x56, 0x2a, 0x96, 0x38, 0x1d, 0xb6, 0x9d, 0x71, 0xb7, 0x61, 0xc3, 0x2e, 0x66,
	0xfe, 0x22, 0x3f, 0x45, 0x01, 0x60, 0x27, 0xc5, 0xdc, 0x34, 0x0d, 0xd2, 0x19, 0x99, 0x71, 0x7f,
	0xa7, 0x59, 0x76, 0x88, 0x38, 0xa1, 0x80, 0x0e, 0x27, 0xb4, 0x8d, 0x9b, 0x9c, 0x50, 0xc4, 0x1d,
	0x4e, 0xe8, 0x0e, 0x32, 0x4e, 0x28, 0xc3, 0x5d, 0x4e, 0xe8, 0x2e, 0x76, 0x38, 0xa1, 0x7b, 0xb8,
	0xcf, 0x09, 0xdd, 0xc7, 0x03, 0x4e, 0xe8, 0x01, 0xba, 0x9c, 0x50, 0x17, 0x0f, 0x39, 0xa1, 0x87,
	0xd8, 0xe5, 0x84, 0x76, 0xf1, 0x88, 0x13, 0x7a, 0x84, 0x8f, 0x38, 0xa1, 0x8f, 0xf0, 0x31, 0x27,
	0xf4, 0x31, 0x1e, 0x73, 0x42, 0x8f, 0xf1, 0x09, 0x27, 0xf4, 0x09, 0xf6, 0x38, 0xa1, 0x3d, 0x3c,
	0xe1, 0x84, 0x9e, 0x60, 0x9f, 0x13, 0xfa, 0x3f, 0x7c, 0xca, 0x09, 0x7d, 0x8a, 0xcf, 0x38, 0xa1,
	0xcf, 0xf0, 0x39, 0x27, 0xf4, 0x39, 0x9e, 0x72, 0x42, 0x4f, 0x71, 0xc0, 0x09, 0x1d, 0xe0, 0x0b,
	0x4e, 0xe8, 0x0b, 0x1c, 0x72, 0x42, 0x87, 0xf8, 0x7f, 0x01, 0xc1, 0x8d, 0x17, 0xc7, 0x72, 0x72,
	0x15, 0x85, 0x02, 0xfc, 0xc5, 0x95, 0x8e, 0x41, 0xdf, 0xeb, 0x99, 0xf2, 0x94, 0x14, 0x60, 0x8b,
	0x71, 0xa5, 0x12, 0x41, 0x67, 0xfa, 0x32, 0x4f, 0x27, 0x02, 0xcc, 0x4d, 0x45, 0x6a, 0x22, 0x05,
	0xfa, 0x69, 0xf2, 0x49, 0x2b, 0x4d, 0xbc, 0x78, 0x3c, 0xd7, 0x05, 0x74, 0xf2, 0x19, 0xbd, 0xd2,
	0xed, 0x2e, 0x9a, 0x85, 0xc1, 0xd6, 0xc7, 0x24, 0x8a, 0xf3, 0xab, 0x63, 0xaf, 0x06, 0xd0, 0x9a,
	0x48, 0xef, 0xce, 0xf2, 0x3d, 0xa5, 0xbc, 0xe0, 0x66, 0xaa, 0xf7, 0x81, 0x20, 0x66, 0x1e, 0x1b,
	0xd7, 0x49, 0x3a, 0xf5, 0x94, 0xb6, 0x39, 0x89, 0xee, 0x64, 0x1a, 0xc9, 0x30, 0x13, 0x8d, 0xeb,
	0x48, 0x4e, 0xc2, 0x4c, 0xd0, 0x20, 0x99, 0xce, 0xe6, 0x4a, 0x86, 0xcb, 0x56, 0xb4, 0xbd, 0xb6,
	0xec, 0xa1, 0x72, 0x47, 0xac, 0xea, 0xbc, 0x2c, 0xa2, 0x2d, 0x54, 0xff, 0x4b, 0x05, 0xda, 0x6f,
	0xf3, 0xc5, 0x38, 0x33, 0xcd, 0xf3, 0x80, 0x35, 0xf1, 0x5f, 0x7b, 0xb3, 0xb9, 0xbe, 0xaa, 0x4f,
	0xa0, 0xad, 0xbc, 0x74, 0x2c, 0xd5, 0x55, 0xae, 0x03, 0x46, 0xe8, 0xe4, 0x3c, 0xd3, 0x33, 0x7a,
	0xcd, 0x58, 0xc8, 0xad, 0x5c, 0xb8, 0x4e, 0xbe, 0x66, 0x72, 0xce, 0xcf, 0x72, 0x61, 0x96, 0x89,
	0x9a, 0x64, 0x6e, 0xdb, 0xec, 0x18, 0x73, 0xe7, 0x84, 0x6e, 0xe2, 0x96, 0x40, 0xab, 0x36, 0xf3,
	0x52, 0x15, 0x69, 0xef, 0x87, 0x09, 0xd4, 0xf3, 0x6d, 0xd7, 0x84, 0xda, 0x87, 0x0f, 0x1f, 0x70,
	0x83, 0xed, 0xc2, 0xf6, 0x57, 0x9f, 0x1f, 0xac, 0xb1, 0x4d, 0x68, 0xbd, 0x2d, 0xbe, 0x06, 0x48,
	0x18, 0x83, 0xad, 0x25, 0x29, 0xe4, 0x6c, 0xb2, 0xc0, 0x06, 0xdb, 0x06, 0x67, 0xed, 0x93, 0x81,
	0x75, 0xd6, 0x01, 0x5c, 0x63, 0xe4, 0xb0, 0xe6, 0xf0, 0x9f, 0x2a, 0x90, 0x77, 0x7a, 0xbd, 0xd5,
	0xa1, 0x12, 0xe0, 0x06, 0x73, 0xa1, 0x53, 0x5a, 0x4d, 0xf3, 0x59, 0x68, 0xde, 0x6c, 0x31, 0x84,
	0x76, 0xd1, 0x53, 0x99, 0x8c, 0x15, 0x02, 0x3b, 0x86, 0x6e, 0x09, 0x6b, 0x3a, 0x6f, 0xa9, 0xe1,
	0xb0, 0x03, 0xd8, 0x2d, 0xc9, 0x4d, 0xf7, 0x84, 0x58, 0x65, 0x7b, 0xb0, 0x53, 0x12, 0x4c, 0xe4,
	0xb5, 0x42, 0xf2, 0x0d, 0x5e, 0x79, 0xe3, 0xb1, 0x0c, 0xb1, 0xc1, 0x0e, 0x61, 0xaf, 0xec, 0x54,
	0x6c, 0x45, 0x4d, 0xe3, 0xaf, 0x1d, 0x84, 0x50, 0x46, 0xb1, 0x92, 0xe3, 0xd4, 0xbc, 0xde, 0x61,
	0xfb, 0xc0, 0x96, 0x23, 0xb2, 0xe2, 0xef, 0x95, 0xe3, 0x90, 0x31, 0x1e, 0xb3, 0x1d, 0xd8, 0x2c,
	0x38, 0x5e, 0x70, 0x2b, 0x43, 0x7c, 0xa2, 0x93, 0x55, 0xb0, 0x52, 0x19, 0xc8, 0xe8, 0x4e, 0x86,
	0xd8, 0xfb, 0x26, 0xe0, 0xa9, 0x9c, 0xfa, 0xf9, 0x38, 0x44, 0xf1, 0x18, 0x4f, 0xd8, 0x11, 0x1c,
	0x94, 0xe4, 0x45, 0xd7, 0xca, 0x10, 0xfb, 0xc3, 0xd3, 0xb5, 0x22, 0xa5, 0x26, 0xe5, 0x14, 0x88,
	0x1e, 0x32, 0xdc, 0x60, 0x0e, 0x34, 0xad, 0xaf, 0x58, 0x19, 0x9e, 0x9a, 0xe2, 0xbe, 0x49, 0xe5,
	0x75, 0xf4, 0x59, 0x4b, 0xde, 0x4b, 0xff, 0x26, 0x49, 0x6e, 0x71, 0x43, 0x97, 0xfd, 0xbd, 0xf4,
	0xb3, 0x44, 0xff, 0x73, 0xc1, 0xca, 0xf0, 0xcf, 0x0a, 0x34, 0x7e, 0xcd, 0xbf, 0x28, 0x14, 0x48,
	0x9c, 0xc4, 0x32, 0x37, 0x65, 0x67, 0x0a, 0x2b, 0xac, 0x65, 0x17, 0x3c, 0x56, 0x35, 0xdf, 0x46,
	0x84, 0x35, 0x9d, 0x83, 0x75, 0x47, 0x91, 0xe8, 0x76, 0x29, 0xb2, 0xa4, 0x19, 0xc7, 0xac, 0xbd,
	0x5a, 0xc2, 0x58, 0x37, 0x0e, 0xe6, 0x83, 0x87, 0x0d, 0x6d, 0xd5, 0xcc, 0x02, 0x36, 0xfd, 0x86,
	0xf9, 0x93, 0xf4, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x13, 0x97, 0x0e, 0x3c, 0x0a,
	0x00, 0x00,
}
