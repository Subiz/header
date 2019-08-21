// Code generated by protoc-gen-go. DO NOT EDIT.
// source: placeholder.proto

package placeholder

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ResolveRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Text                 *string         `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	DefaultValue         *string         `protobuf:"bytes,3,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	OutTopic             *string         `protobuf:"bytes,5,opt,name=out_topic,json=outTopic" json:"out_topic,omitempty"`
	OutKey               *string         `protobuf:"bytes,6,opt,name=out_key,json=outKey" json:"out_key,omitempty"`
	AccountId            *string         `protobuf:"bytes,11,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AgentId              *string         `protobuf:"bytes,12,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MessageId            *string         `protobuf:"bytes,13,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	ConversationId       *string         `protobuf:"bytes,14,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	PostbackId           *string         `protobuf:"bytes,15,opt,name=postback_id,json=postbackId" json:"postback_id,omitempty"`
	ContentId            *string         `protobuf:"bytes,16,opt,name=content_id,json=contentId" json:"content_id,omitempty"`
	TopicId              *string         `protobuf:"bytes,17,opt,name=topic_id,json=topicId" json:"topic_id,omitempty"`
	PresenceId           *string         `protobuf:"bytes,18,opt,name=presence_id,json=presenceId" json:"presence_id,omitempty"`
	UserId               *string         `protobuf:"bytes,20,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UnreadTopicId        *string         `protobuf:"bytes,21,opt,name=unread_topic_id,json=unreadTopicId" json:"unread_topic_id,omitempty"`
	MyUserId             *string         `protobuf:"bytes,22,opt,name=my_user_id,json=myUserId" json:"my_user_id,omitempty"`
	NotificationId       *string         `protobuf:"bytes,24,opt,name=notification_id,json=notificationId" json:"notification_id,omitempty"`
	NotiboxId            *string         `protobuf:"bytes,25,opt,name=notibox_id,json=notiboxId" json:"notibox_id,omitempty"`
	AgentPermId          *string         `protobuf:"bytes,26,opt,name=agent_perm_id,json=agentPermId" json:"agent_perm_id,omitempty"`
	GroupMemberId        *string         `protobuf:"bytes,27,opt,name=group_member_id,json=groupMemberId" json:"group_member_id,omitempty"`
	GroupId              *string         `protobuf:"bytes,28,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	IntegrationId        *string         `protobuf:"bytes,40,opt,name=integration_id,json=integrationId" json:"integration_id,omitempty"`
	LimitId              *string         `protobuf:"bytes,29,opt,name=limit_id,json=limitId" json:"limit_id,omitempty"`
	UserAttributeId      *string         `protobuf:"bytes,30,opt,name=user_attribute_id,json=userAttributeId" json:"user_attribute_id,omitempty"`
	AliasId              *string         `protobuf:"bytes,32,opt,name=alias_id,json=aliasId" json:"alias_id,omitempty"`
	NoteId               *string         `protobuf:"bytes,33,opt,name=note_id,json=noteId" json:"note_id,omitempty"`
	TicketId             *string         `protobuf:"bytes,34,opt,name=ticket_id,json=ticketId" json:"ticket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResolveRequest) Reset()         { *m = ResolveRequest{} }
func (m *ResolveRequest) String() string { return proto.CompactTextString(m) }
func (*ResolveRequest) ProtoMessage()    {}
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81af5d00480424c6, []int{0}
}

func (m *ResolveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveRequest.Unmarshal(m, b)
}
func (m *ResolveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveRequest.Marshal(b, m, deterministic)
}
func (m *ResolveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveRequest.Merge(m, src)
}
func (m *ResolveRequest) XXX_Size() int {
	return xxx_messageInfo_ResolveRequest.Size(m)
}
func (m *ResolveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveRequest proto.InternalMessageInfo

func (m *ResolveRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ResolveRequest) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *ResolveRequest) GetDefaultValue() string {
	if m != nil && m.DefaultValue != nil {
		return *m.DefaultValue
	}
	return ""
}

func (m *ResolveRequest) GetOutTopic() string {
	if m != nil && m.OutTopic != nil {
		return *m.OutTopic
	}
	return ""
}

func (m *ResolveRequest) GetOutKey() string {
	if m != nil && m.OutKey != nil {
		return *m.OutKey
	}
	return ""
}

func (m *ResolveRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ResolveRequest) GetAgentId() string {
	if m != nil && m.AgentId != nil {
		return *m.AgentId
	}
	return ""
}

func (m *ResolveRequest) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *ResolveRequest) GetConversationId() string {
	if m != nil && m.ConversationId != nil {
		return *m.ConversationId
	}
	return ""
}

func (m *ResolveRequest) GetPostbackId() string {
	if m != nil && m.PostbackId != nil {
		return *m.PostbackId
	}
	return ""
}

func (m *ResolveRequest) GetContentId() string {
	if m != nil && m.ContentId != nil {
		return *m.ContentId
	}
	return ""
}

func (m *ResolveRequest) GetTopicId() string {
	if m != nil && m.TopicId != nil {
		return *m.TopicId
	}
	return ""
}

func (m *ResolveRequest) GetPresenceId() string {
	if m != nil && m.PresenceId != nil {
		return *m.PresenceId
	}
	return ""
}

func (m *ResolveRequest) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *ResolveRequest) GetUnreadTopicId() string {
	if m != nil && m.UnreadTopicId != nil {
		return *m.UnreadTopicId
	}
	return ""
}

func (m *ResolveRequest) GetMyUserId() string {
	if m != nil && m.MyUserId != nil {
		return *m.MyUserId
	}
	return ""
}

func (m *ResolveRequest) GetNotificationId() string {
	if m != nil && m.NotificationId != nil {
		return *m.NotificationId
	}
	return ""
}

func (m *ResolveRequest) GetNotiboxId() string {
	if m != nil && m.NotiboxId != nil {
		return *m.NotiboxId
	}
	return ""
}

func (m *ResolveRequest) GetAgentPermId() string {
	if m != nil && m.AgentPermId != nil {
		return *m.AgentPermId
	}
	return ""
}

func (m *ResolveRequest) GetGroupMemberId() string {
	if m != nil && m.GroupMemberId != nil {
		return *m.GroupMemberId
	}
	return ""
}

func (m *ResolveRequest) GetGroupId() string {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return ""
}

func (m *ResolveRequest) GetIntegrationId() string {
	if m != nil && m.IntegrationId != nil {
		return *m.IntegrationId
	}
	return ""
}

func (m *ResolveRequest) GetLimitId() string {
	if m != nil && m.LimitId != nil {
		return *m.LimitId
	}
	return ""
}

func (m *ResolveRequest) GetUserAttributeId() string {
	if m != nil && m.UserAttributeId != nil {
		return *m.UserAttributeId
	}
	return ""
}

func (m *ResolveRequest) GetAliasId() string {
	if m != nil && m.AliasId != nil {
		return *m.AliasId
	}
	return ""
}

func (m *ResolveRequest) GetNoteId() string {
	if m != nil && m.NoteId != nil {
		return *m.NoteId
	}
	return ""
}

func (m *ResolveRequest) GetTicketId() string {
	if m != nil && m.TicketId != nil {
		return *m.TicketId
	}
	return ""
}

type Resolved struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Text                 *string         `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Errors               []string        `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Resolved) Reset()         { *m = Resolved{} }
func (m *Resolved) String() string { return proto.CompactTextString(m) }
func (*Resolved) ProtoMessage()    {}
func (*Resolved) Descriptor() ([]byte, []int) {
	return fileDescriptor_81af5d00480424c6, []int{1}
}

func (m *Resolved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resolved.Unmarshal(m, b)
}
func (m *Resolved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resolved.Marshal(b, m, deterministic)
}
func (m *Resolved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resolved.Merge(m, src)
}
func (m *Resolved) XXX_Size() int {
	return xxx_messageInfo_Resolved.Size(m)
}
func (m *Resolved) XXX_DiscardUnknown() {
	xxx_messageInfo_Resolved.DiscardUnknown(m)
}

var xxx_messageInfo_Resolved proto.InternalMessageInfo

func (m *Resolved) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Resolved) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *Resolved) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*ResolveRequest)(nil), "placeholder.ResolveRequest")
	proto.RegisterType((*Resolved)(nil), "placeholder.Resolved")
}

func init() { proto.RegisterFile("placeholder.proto", fileDescriptor_81af5d00480424c6) }

var fileDescriptor_81af5d00480424c6 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x6e, 0x13, 0x3d,
	0x10, 0xc5, 0xd5, 0x2f, 0x1f, 0x69, 0xeb, 0x34, 0x59, 0x6a, 0x41, 0xd9, 0xfe, 0x83, 0x34, 0x55,
	0x69, 0xc5, 0x45, 0x2b, 0xf1, 0x06, 0xc0, 0xd5, 0x0a, 0x21, 0xa1, 0xa8, 0x20, 0xc1, 0xcd, 0xca,
	0x6b, 0x4f, 0x53, 0xab, 0xbb, 0xeb, 0xc5, 0x6b, 0x57, 0x0d, 0xef, 0xcc, 0x3b, 0xa0, 0x99, 0x71,
	0xa2, 0x5c, 0x73, 0xb7, 0xf3, 0x3b, 0x67, 0xc7, 0xe3, 0x63, 0x5b, 0xec, 0x77, 0xb5, 0xd2, 0x70,
	0xef, 0x6a, 0x03, 0xfe, 0xba, 0xf3, 0x2e, 0x38, 0x39, 0xda, 0x40, 0x47, 0x7b, 0xda, 0x35, 0x8d,
	0x6b, 0x59, 0x9a, 0xfd, 0x19, 0x8a, 0xc9, 0x1c, 0x7a, 0x57, 0x3f, 0xc2, 0x1c, 0x7e, 0x45, 0xe8,
	0x83, 0x3c, 0x13, 0x03, 0x1d, 0x9e, 0xf2, 0xad, 0xe9, 0xd6, 0xd5, 0xe8, 0x7d, 0x76, 0x9d, 0xec,
	0x9f, 0x5c, 0x1b, 0xe0, 0x29, 0xcc, 0x51, 0x93, 0x52, 0xfc, 0x8f, 0x45, 0xfe, 0xdf, 0x74, 0xeb,
	0x6a, 0x77, 0x4e, 0xdf, 0xf2, 0x5c, 0x8c, 0x0d, 0xdc, 0xa9, 0x58, 0x87, 0xf2, 0x51, 0xd5, 0x11,
	0xf2, 0x01, 0x89, 0x7b, 0x09, 0x7e, 0x47, 0x26, 0x8f, 0xc5, 0xae, 0x8b, 0xa1, 0x0c, 0xae, 0xb3,
	0x3a, 0x7f, 0x46, 0x86, 0x1d, 0x17, 0xc3, 0x2d, 0xd6, 0xf2, 0x95, 0xd8, 0x46, 0xf1, 0x01, 0x96,
	0xf9, 0x90, 0xa4, 0xa1, 0x8b, 0xe1, 0x33, 0x2c, 0xe5, 0xa9, 0x10, 0x4a, 0x6b, 0x17, 0xdb, 0x50,
	0x5a, 0x93, 0x8f, 0x48, 0xdb, 0x4d, 0xa4, 0x30, 0xf2, 0x50, 0xec, 0xa8, 0x05, 0xb0, 0xb8, 0x47,
	0xe2, 0x36, 0xd5, 0x85, 0xc1, 0x3f, 0x1b, 0xe8, 0x7b, 0xb5, 0x00, 0x14, 0xc7, 0xfc, 0x67, 0x22,
	0x85, 0x91, 0x97, 0x22, 0xd3, 0xae, 0x7d, 0x04, 0xdf, 0xab, 0x60, 0x5d, 0x8b, 0x9e, 0x09, 0x79,
	0x26, 0x9b, 0xb8, 0x30, 0xf2, 0x8d, 0x18, 0x75, 0xae, 0x0f, 0x95, 0xd2, 0x0f, 0x68, 0xca, 0xc8,
	0x24, 0x56, 0x88, 0x17, 0xd2, 0x98, 0x10, 0x4f, 0xf1, 0x9c, 0x17, 0x4a, 0x84, 0x47, 0xa4, 0x3d,
	0xa3, 0xb8, 0xcf, 0x23, 0x52, 0x9d, 0x5a, 0x7b, 0xe8, 0xa1, 0xd5, 0x34, 0xa3, 0x4c, 0xad, 0x13,
	0x2a, 0x0c, 0xc6, 0x12, 0x7b, 0xf0, 0x28, 0xbe, 0xe0, 0x58, 0xb0, 0x2c, 0x8c, 0x7c, 0x2b, 0xb2,
	0xd8, 0x7a, 0x50, 0xa6, 0x5c, 0xf7, 0x7e, 0x49, 0x86, 0x31, 0xe3, 0xdb, 0xb4, 0xc2, 0x89, 0x10,
	0xcd, 0xb2, 0x5c, 0xf5, 0x38, 0xe0, 0xd4, 0x9b, 0xe5, 0x37, 0xee, 0x72, 0x29, 0xb2, 0xd6, 0x05,
	0x7b, 0x67, 0xf5, 0x3a, 0x83, 0x9c, 0x33, 0xd8, 0xc4, 0xbc, 0x45, 0x24, 0x95, 0x7b, 0x42, 0xcf,
	0x21, 0x6f, 0x31, 0x91, 0xc2, 0xc8, 0x99, 0x18, 0xf3, 0x29, 0x74, 0xe0, 0x1b, 0x74, 0x1c, 0x91,
	0x63, 0x44, 0xf0, 0x2b, 0xf8, 0x86, 0x27, 0x5e, 0x78, 0x17, 0xbb, 0xb2, 0x81, 0xa6, 0xe2, 0x71,
	0x8e, 0x79, 0x62, 0xc2, 0x5f, 0x88, 0x72, 0x5c, 0xec, 0xb3, 0x26, 0x3f, 0xe1, 0xb8, 0xa8, 0x2e,
	0x8c, 0xbc, 0x10, 0x13, 0xdb, 0x06, 0x58, 0xf8, 0xf5, 0xb4, 0x57, 0xdc, 0x61, 0x83, 0x72, 0x87,
	0xda, 0x36, 0x96, 0x4e, 0xe3, 0x94, 0x3b, 0x50, 0x5d, 0x18, 0xf9, 0x4e, 0xec, 0x53, 0x16, 0x2a,
	0x04, 0x6f, 0xab, 0x18, 0x28, 0xf6, 0xd7, 0xe4, 0xc9, 0x50, 0xf8, 0xb0, 0xe2, 0xe9, 0x6a, 0xd5,
	0x56, 0xf5, 0x68, 0x99, 0xa6, 0xab, 0x85, 0x35, 0x1f, 0x4b, 0xeb, 0xf8, 0xe7, 0x33, 0x3e, 0x16,
	0x2c, 0x0b, 0x83, 0x77, 0x3c, 0x58, 0xfd, 0x00, 0xb4, 0xf6, 0x8c, 0xd3, 0x66, 0x50, 0x98, 0xd9,
	0x0f, 0xb1, 0x93, 0x9e, 0x9b, 0xf9, 0xd7, 0x87, 0x76, 0x20, 0x86, 0xe0, 0xbd, 0xf3, 0x7d, 0x3e,
	0x98, 0x0e, 0x70, 0x5d, 0xae, 0x3e, 0x5e, 0xfc, 0x3c, 0x5f, 0xd8, 0x70, 0x1f, 0x2b, 0xec, 0x74,
	0xd3, 0xc7, 0xca, 0xfe, 0xbe, 0xb9, 0x07, 0x65, 0xc0, 0xdf, 0x6c, 0xbc, 0xff, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x94, 0x02, 0x96, 0xf9, 0x20, 0x04, 0x00, 0x00,
}
