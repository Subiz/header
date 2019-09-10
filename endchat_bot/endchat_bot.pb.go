// Code generated by protoc-gen-go. DO NOT EDIT.
// source: endchat_bot.proto

package endchat_bot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/subiz/header/common"
	math "math"
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

type ConnectorSetting struct {
	ConnectorId          string   `protobuf:"bytes,3,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	AtMidnight           bool     `protobuf:"varint,4,opt,name=at_midnight,json=atMidnight,proto3" json:"at_midnight,omitempty"`
	AfterAgentMessage    int64    `protobuf:"varint,6,opt,name=after_agent_message,json=afterAgentMessage,proto3" json:"after_agent_message,omitempty"`
	AfterAnyMessage      int64    `protobuf:"varint,7,opt,name=after_any_message,json=afterAnyMessage,proto3" json:"after_any_message,omitempty"`
	AfterUserMessage     int64    `protobuf:"varint,8,opt,name=after_user_message,json=afterUserMessage,proto3" json:"after_user_message,omitempty"`
	Age                  int64    `protobuf:"varint,10,opt,name=age,proto3" json:"age,omitempty"`
	Enabled              bool     `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectorSetting) Reset()         { *m = ConnectorSetting{} }
func (m *ConnectorSetting) String() string { return proto.CompactTextString(m) }
func (*ConnectorSetting) ProtoMessage()    {}
func (*ConnectorSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_b405950ab0fad1c7, []int{0}
}

func (m *ConnectorSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectorSetting.Unmarshal(m, b)
}
func (m *ConnectorSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectorSetting.Marshal(b, m, deterministic)
}
func (m *ConnectorSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectorSetting.Merge(m, src)
}
func (m *ConnectorSetting) XXX_Size() int {
	return xxx_messageInfo_ConnectorSetting.Size(m)
}
func (m *ConnectorSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectorSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectorSetting proto.InternalMessageInfo

func (m *ConnectorSetting) GetConnectorId() string {
	if m != nil {
		return m.ConnectorId
	}
	return ""
}

func (m *ConnectorSetting) GetAtMidnight() bool {
	if m != nil {
		return m.AtMidnight
	}
	return false
}

func (m *ConnectorSetting) GetAfterAgentMessage() int64 {
	if m != nil {
		return m.AfterAgentMessage
	}
	return 0
}

func (m *ConnectorSetting) GetAfterAnyMessage() int64 {
	if m != nil {
		return m.AfterAnyMessage
	}
	return 0
}

func (m *ConnectorSetting) GetAfterUserMessage() int64 {
	if m != nil {
		return m.AfterUserMessage
	}
	return 0
}

func (m *ConnectorSetting) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *ConnectorSetting) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type Setting struct {
	AccountId            string              `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ConnectorSettings    []*ConnectorSetting `protobuf:"bytes,3,rep,name=connector_settings,json=connectorSettings,proto3" json:"connector_settings,omitempty"`
	GlobalSetting        *ConnectorSetting   `protobuf:"bytes,4,opt,name=global_setting,json=globalSetting,proto3" json:"global_setting,omitempty"`
	Updated              int64               `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Setting) Reset()         { *m = Setting{} }
func (m *Setting) String() string { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()    {}
func (*Setting) Descriptor() ([]byte, []int) {
	return fileDescriptor_b405950ab0fad1c7, []int{1}
}

func (m *Setting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setting.Unmarshal(m, b)
}
func (m *Setting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setting.Marshal(b, m, deterministic)
}
func (m *Setting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setting.Merge(m, src)
}
func (m *Setting) XXX_Size() int {
	return xxx_messageInfo_Setting.Size(m)
}
func (m *Setting) XXX_DiscardUnknown() {
	xxx_messageInfo_Setting.DiscardUnknown(m)
}

var xxx_messageInfo_Setting proto.InternalMessageInfo

func (m *Setting) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Setting) GetConnectorSettings() []*ConnectorSetting {
	if m != nil {
		return m.ConnectorSettings
	}
	return nil
}

func (m *Setting) GetGlobalSetting() *ConnectorSetting {
	if m != nil {
		return m.GlobalSetting
	}
	return nil
}

func (m *Setting) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// (create_min, account_id, id)
// (account_id, id)
type Conversation struct {
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	LastMessageSent      int64    `protobuf:"varint,5,opt,name=last_message_sent,json=lastMessageSent,proto3" json:"last_message_sent,omitempty"`
	LastUserMessageSent  int64    `protobuf:"varint,6,opt,name=last_user_message_sent,json=lastUserMessageSent,proto3" json:"last_user_message_sent,omitempty"`
	LastAgentMessageSent int64    `protobuf:"varint,7,opt,name=last_agent_message_sent,json=lastAgentMessageSent,proto3" json:"last_agent_message_sent,omitempty"`
	Ended                int64    `protobuf:"varint,8,opt,name=ended,proto3" json:"ended,omitempty"`
	EndRequested         int64    `protobuf:"varint,9,opt,name=end_requested,json=endRequested,proto3" json:"end_requested,omitempty"`
	State                string   `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	IntegrationId        string   `protobuf:"bytes,11,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	ConnectorId          string   `protobuf:"bytes,12,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conversation) Reset()         { *m = Conversation{} }
func (m *Conversation) String() string { return proto.CompactTextString(m) }
func (*Conversation) ProtoMessage()    {}
func (*Conversation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b405950ab0fad1c7, []int{2}
}

func (m *Conversation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conversation.Unmarshal(m, b)
}
func (m *Conversation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conversation.Marshal(b, m, deterministic)
}
func (m *Conversation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conversation.Merge(m, src)
}
func (m *Conversation) XXX_Size() int {
	return xxx_messageInfo_Conversation.Size(m)
}
func (m *Conversation) XXX_DiscardUnknown() {
	xxx_messageInfo_Conversation.DiscardUnknown(m)
}

var xxx_messageInfo_Conversation proto.InternalMessageInfo

func (m *Conversation) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Conversation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Conversation) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Conversation) GetLastMessageSent() int64 {
	if m != nil {
		return m.LastMessageSent
	}
	return 0
}

func (m *Conversation) GetLastUserMessageSent() int64 {
	if m != nil {
		return m.LastUserMessageSent
	}
	return 0
}

func (m *Conversation) GetLastAgentMessageSent() int64 {
	if m != nil {
		return m.LastAgentMessageSent
	}
	return 0
}

func (m *Conversation) GetEnded() int64 {
	if m != nil {
		return m.Ended
	}
	return 0
}

func (m *Conversation) GetEndRequested() int64 {
	if m != nil {
		return m.EndRequested
	}
	return 0
}

func (m *Conversation) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Conversation) GetIntegrationId() string {
	if m != nil {
		return m.IntegrationId
	}
	return ""
}

func (m *Conversation) GetConnectorId() string {
	if m != nil {
		return m.ConnectorId
	}
	return ""
}

type EndCallback struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ConversationId       string          `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ConnectorId          string          `protobuf:"bytes,4,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EndCallback) Reset()         { *m = EndCallback{} }
func (m *EndCallback) String() string { return proto.CompactTextString(m) }
func (*EndCallback) ProtoMessage()    {}
func (*EndCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_b405950ab0fad1c7, []int{3}
}

func (m *EndCallback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndCallback.Unmarshal(m, b)
}
func (m *EndCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndCallback.Marshal(b, m, deterministic)
}
func (m *EndCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndCallback.Merge(m, src)
}
func (m *EndCallback) XXX_Size() int {
	return xxx_messageInfo_EndCallback.Size(m)
}
func (m *EndCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_EndCallback.DiscardUnknown(m)
}

var xxx_messageInfo_EndCallback proto.InternalMessageInfo

func (m *EndCallback) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *EndCallback) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *EndCallback) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *EndCallback) GetConnectorId() string {
	if m != nil {
		return m.ConnectorId
	}
	return ""
}

type MidnightCallback struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	MidnightSec          int64           `protobuf:"varint,3,opt,name=midnight_sec,json=midnightSec,proto3" json:"midnight_sec,omitempty"`
	ConnectorId          string          `protobuf:"bytes,4,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MidnightCallback) Reset()         { *m = MidnightCallback{} }
func (m *MidnightCallback) String() string { return proto.CompactTextString(m) }
func (*MidnightCallback) ProtoMessage()    {}
func (*MidnightCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_b405950ab0fad1c7, []int{4}
}

func (m *MidnightCallback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MidnightCallback.Unmarshal(m, b)
}
func (m *MidnightCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MidnightCallback.Marshal(b, m, deterministic)
}
func (m *MidnightCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MidnightCallback.Merge(m, src)
}
func (m *MidnightCallback) XXX_Size() int {
	return xxx_messageInfo_MidnightCallback.Size(m)
}
func (m *MidnightCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_MidnightCallback.DiscardUnknown(m)
}

var xxx_messageInfo_MidnightCallback proto.InternalMessageInfo

func (m *MidnightCallback) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *MidnightCallback) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *MidnightCallback) GetMidnightSec() int64 {
	if m != nil {
		return m.MidnightSec
	}
	return 0
}

func (m *MidnightCallback) GetConnectorId() string {
	if m != nil {
		return m.ConnectorId
	}
	return ""
}

func init() {
	proto.RegisterType((*ConnectorSetting)(nil), "endchat_bot.ConnectorSetting")
	proto.RegisterType((*Setting)(nil), "endchat_bot.Setting")
	proto.RegisterType((*Conversation)(nil), "endchat_bot.Conversation")
	proto.RegisterType((*EndCallback)(nil), "endchat_bot.EndCallback")
	proto.RegisterType((*MidnightCallback)(nil), "endchat_bot.MidnightCallback")
}

func init() { proto.RegisterFile("endchat_bot.proto", fileDescriptor_b405950ab0fad1c7) }

var fileDescriptor_b405950ab0fad1c7 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x95, 0xe3, 0xb6, 0x69, 0xc6, 0x6e, 0x7e, 0xb6, 0xd5, 0xf7, 0x59, 0x48, 0x15, 0x69, 0xaa,
	0x8a, 0x08, 0xa1, 0x54, 0xa2, 0xe2, 0x01, 0x20, 0x70, 0x51, 0x89, 0xde, 0xb8, 0xe2, 0x86, 0x1b,
	0x6b, 0xbd, 0x3b, 0x38, 0x16, 0xce, 0x6e, 0xf1, 0xae, 0x51, 0xcb, 0x23, 0x20, 0x9e, 0x00, 0xde,
	0x8b, 0xe7, 0x41, 0xbb, 0xeb, 0x4d, 0x4d, 0x8b, 0x54, 0x2e, 0xb8, 0xcb, 0x9c, 0x73, 0x66, 0x32,
	0x3b, 0x67, 0xc6, 0x30, 0x41, 0xc1, 0xd9, 0x8a, 0xea, 0x2c, 0x97, 0x7a, 0x71, 0x55, 0x4b, 0x2d,
	0x49, 0xd4, 0x81, 0x1e, 0xc5, 0x4c, 0xae, 0xd7, 0x52, 0x38, 0x6a, 0xf6, 0xad, 0x07, 0xe3, 0xa5,
	0x14, 0x02, 0x99, 0x96, 0xf5, 0x25, 0x6a, 0x5d, 0x8a, 0x82, 0x1c, 0x41, 0xcc, 0x3c, 0x96, 0x95,
	0x3c, 0x09, 0xa7, 0xc1, 0x7c, 0x90, 0x46, 0x1b, 0xec, 0x9c, 0x93, 0xc7, 0x10, 0x51, 0x9d, 0xad,
	0x4b, 0x2e, 0xca, 0x62, 0xa5, 0x93, 0xad, 0x69, 0x30, 0xdf, 0x4d, 0x81, 0xea, 0x8b, 0x16, 0x21,
	0x0b, 0xd8, 0xa7, 0x1f, 0x34, 0xd6, 0x19, 0x2d, 0x50, 0xe8, 0x6c, 0x8d, 0x4a, 0xd1, 0x02, 0x93,
	0x9d, 0x69, 0x30, 0x0f, 0xd3, 0x89, 0xa5, 0x5e, 0x1a, 0xe6, 0xc2, 0x11, 0xe4, 0x29, 0x4c, 0x5a,
	0xbd, 0xb8, 0xd9, 0xa8, 0xfb, 0x56, 0x3d, 0x72, 0x6a, 0x71, 0xe3, 0xb5, 0xcf, 0x80, 0x38, 0x6d,
	0xa3, 0xb0, 0xde, 0x88, 0x77, 0xad, 0x78, 0x6c, 0x99, 0x77, 0x0a, 0x6b, 0xaf, 0x1e, 0x43, 0x68,
	0x68, 0xb0, 0xb4, 0xf9, 0x49, 0x12, 0xe8, 0xa3, 0xa0, 0x79, 0x85, 0x3c, 0x89, 0x6c, 0xe3, 0x3e,
	0x9c, 0xfd, 0x0c, 0xa0, 0xef, 0xa7, 0x70, 0x08, 0x40, 0x19, 0x93, 0x8d, 0xd0, 0x66, 0x06, 0x3d,
	0x3b, 0x83, 0x41, 0x8b, 0x9c, 0x73, 0xf2, 0x16, 0xc8, 0xed, 0x90, 0x94, 0xcb, 0x51, 0x49, 0x38,
	0x0d, 0xe7, 0xd1, 0xf3, 0xc3, 0x45, 0xd7, 0x84, 0xbb, 0xf3, 0x4d, 0x27, 0xec, 0x0e, 0xa2, 0xc8,
	0x6b, 0x18, 0x16, 0x95, 0xcc, 0x69, 0xe5, 0x4b, 0xd9, 0x91, 0x3e, 0x58, 0x69, 0xcf, 0x25, 0xf9,
	0x96, 0x13, 0xe8, 0x37, 0x57, 0x9c, 0x6a, 0xe4, 0xc9, 0xb6, 0x7d, 0xae, 0x0f, 0x67, 0x5f, 0x43,
	0x88, 0x97, 0x52, 0x7c, 0xc6, 0x5a, 0x51, 0x5d, 0x4a, 0xf1, 0xd0, 0xeb, 0x86, 0xd0, 0xdb, 0x18,
	0xdf, 0x2b, 0xb9, 0xa9, 0xcc, 0x6a, 0xb4, 0x95, 0xb7, 0x5c, 0xe5, 0x36, 0x34, 0xc6, 0x55, 0x54,
	0x6d, 0x1c, 0xce, 0x14, 0x0a, 0xdd, 0xfe, 0xfb, 0xc8, 0x10, 0xad, 0x0d, 0x97, 0x28, 0x34, 0x39,
	0x83, 0xff, 0xac, 0xb6, 0xeb, 0x9b, 0x4b, 0x70, 0x7b, 0xb1, 0x6f, 0xd8, 0x8e, 0x77, 0x36, 0xe9,
	0x05, 0xfc, 0x6f, 0x93, 0x7e, 0x5b, 0x24, 0x97, 0xe5, 0xf6, 0xe3, 0xc0, 0xd0, 0xdd, 0x65, 0xb2,
	0x69, 0x07, 0xb0, 0x8d, 0x82, 0x23, 0x6f, 0xf7, 0xc2, 0x05, 0xe4, 0x18, 0xf6, 0x50, 0xf0, 0xac,
	0xc6, 0x4f, 0x0d, 0x2a, 0xf3, 0x9a, 0x81, 0x65, 0x63, 0x14, 0x3c, 0xf5, 0x98, 0x49, 0x55, 0x9a,
	0x6a, 0xb7, 0x33, 0x83, 0xd4, 0x05, 0xe4, 0x04, 0x86, 0xa5, 0xd0, 0x58, 0xd4, 0x76, 0x80, 0x66,
	0x6a, 0x91, 0xa5, 0xf7, 0x3a, 0xe8, 0x39, 0xbf, 0x77, 0x3c, 0xf1, 0xbd, 0xe3, 0x99, 0xfd, 0x08,
	0x20, 0x7a, 0x23, 0xf8, 0x92, 0x56, 0x55, 0x4e, 0xd9, 0x47, 0x72, 0x04, 0x21, 0xd3, 0xd7, 0x49,
	0x60, 0x1d, 0x1f, 0x2d, 0xda, 0x03, 0x5d, 0x4a, 0xa1, 0xf1, 0x5a, 0xa7, 0x86, 0x7b, 0xc8, 0xae,
	0x27, 0x30, 0x62, 0x1d, 0x77, 0x6f, 0x8f, 0x76, 0xd8, 0x85, 0xff, 0xd0, 0xdd, 0xd6, 0xfd, 0xee,
	0xbe, 0x07, 0x30, 0xf6, 0x67, 0xfc, 0x0f, 0x5b, 0x3c, 0x82, 0xd8, 0x7f, 0x2e, 0x32, 0x85, 0xcc,
	0xf6, 0x17, 0xa6, 0x91, 0xc7, 0x2e, 0x91, 0xfd, 0x45, 0x73, 0xaf, 0x4e, 0xde, 0x1f, 0x17, 0xa5,
	0x5e, 0x35, 0xb9, 0x69, 0xe1, 0x54, 0x35, 0x79, 0xf9, 0xe5, 0x74, 0x85, 0x94, 0x63, 0x7d, 0xda,
	0x39, 0x94, 0x7c, 0xc7, 0x7e, 0xdd, 0xce, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x83, 0x6c,
	0x0c, 0x0d, 0x05, 0x00, 0x00,
}
