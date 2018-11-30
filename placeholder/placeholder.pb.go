// Code generated by protoc-gen-go. DO NOT EDIT.
// source: placeholder/placeholder.proto

package placeholder

import (
	fmt "fmt"
	common "git.subiz.net/header/common"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Event int32

const (
	Event_PlaceholderSynced Event = 0
	Event_ResolveRequested  Event = 1
)

var Event_name = map[int32]string{
	0: "PlaceholderSynced",
	1: "ResolveRequested",
}

var Event_value = map[string]int32{
	"PlaceholderSynced": 0,
	"ResolveRequested":  1,
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

func (Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_005c33552af3c3f0, []int{0}
}

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
	return fileDescriptor_005c33552af3c3f0, []int{0}
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
	return fileDescriptor_005c33552af3c3f0, []int{1}
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
	proto.RegisterEnum("placeholder.Event", Event_name, Event_value)
	proto.RegisterType((*ResolveRequest)(nil), "placeholder.ResolveRequest")
	proto.RegisterType((*Resolved)(nil), "placeholder.Resolved")
}

func init() { proto.RegisterFile("placeholder/placeholder.proto", fileDescriptor_005c33552af3c3f0) }

var fileDescriptor_005c33552af3c3f0 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0x13, 0x31,
	0x10, 0x6e, 0x08, 0x4d, 0xdb, 0x49, 0xf3, 0xa8, 0xd5, 0x96, 0xed, 0x0b, 0xd2, 0x20, 0x68, 0xd4,
	0x43, 0x2a, 0x55, 0xdc, 0x11, 0x42, 0x48, 0xac, 0x10, 0xa8, 0x0a, 0x05, 0x89, 0x53, 0xb4, 0xb1,
	0xa7, 0xa9, 0xd5, 0x5d, 0x3b, 0x78, 0xbd, 0x51, 0xc3, 0x7f, 0xe6, 0x3f, 0x20, 0xcf, 0xb8, 0x21,
	0x95, 0x38, 0x71, 0x4a, 0xbe, 0xc7, 0x8c, 0xc7, 0x9f, 0xed, 0x85, 0x93, 0x59, 0x9e, 0x49, 0xbc,
	0xb5, 0xb9, 0x42, 0x77, 0xb1, 0xf2, 0x7f, 0x38, 0x73, 0xd6, 0x5b, 0xd1, 0x5c, 0xa1, 0x0e, 0x07,
	0x53, 0xed, 0x87, 0x65, 0x35, 0xd1, 0xbf, 0x86, 0x06, 0xfd, 0xc5, 0x2d, 0x66, 0xa1, 0x48, 0xda,
	0xa2, 0xb0, 0x26, 0xfe, 0x70, 0x59, 0xff, 0x77, 0x03, 0xda, 0x23, 0x2c, 0x6d, 0x3e, 0xc7, 0x11,
	0xfe, 0xac, 0xb0, 0xf4, 0xe2, 0x14, 0xea, 0xd2, 0xdf, 0x27, 0xb5, 0x5e, 0x6d, 0xd0, 0xbc, 0xec,
	0x0c, 0xa3, 0xfd, 0xbd, 0x35, 0x1e, 0xef, 0xfd, 0x28, 0x68, 0x42, 0xc0, 0xd3, 0x00, 0x92, 0x27,
	0xbd, 0xda, 0x60, 0x6b, 0x44, 0xff, 0xc5, 0x4b, 0x68, 0x29, 0xbc, 0xc9, 0xaa, 0xdc, 0x8f, 0xe7,
	0x59, 0x5e, 0x61, 0x52, 0x27, 0x71, 0x3b, 0x92, 0xdf, 0x03, 0x27, 0x8e, 0x60, 0xcb, 0x56, 0x7e,
	0xec, 0xed, 0x4c, 0xcb, 0x64, 0x9d, 0x0c, 0x9b, 0xb6, 0xf2, 0xd7, 0x01, 0x8b, 0x67, 0xb0, 0x11,
	0xc4, 0x3b, 0x5c, 0x24, 0x0d, 0x92, 0x1a, 0xb6, 0xf2, 0x9f, 0x70, 0x21, 0x4e, 0x00, 0x32, 0x29,
	0x6d, 0x65, 0xfc, 0x58, 0xab, 0xa4, 0x49, 0xda, 0x56, 0x64, 0x52, 0x25, 0x0e, 0x60, 0x33, 0x9b,
	0x22, 0x8b, 0xdb, 0x24, 0x6e, 0x10, 0x4e, 0x55, 0xa8, 0x2c, 0xb0, 0x2c, 0xb3, 0x29, 0x06, 0xb1,
	0xc5, 0x95, 0x91, 0x49, 0x95, 0x38, 0x83, 0x8e, 0xb4, 0x66, 0x8e, 0xae, 0xcc, 0xbc, 0xb6, 0x26,
	0x78, 0xda, 0xe4, 0x69, 0xaf, 0xd2, 0xa9, 0x12, 0x2f, 0xa0, 0x39, 0xb3, 0xa5, 0x9f, 0x64, 0xf2,
	0x2e, 0x98, 0x3a, 0x64, 0x82, 0x07, 0x8a, 0x17, 0x92, 0x21, 0x21, 0x9e, 0xa2, 0xcb, 0x0b, 0x45,
	0x86, 0x47, 0xa4, 0x3d, 0x07, 0x71, 0x87, 0x47, 0x24, 0x1c, 0x5b, 0x3b, 0x2c, 0xd1, 0x48, 0x9a,
	0x51, 0xc4, 0xd6, 0x91, 0x4a, 0x55, 0x88, 0xa5, 0x2a, 0xd1, 0x05, 0x71, 0x97, 0x63, 0x09, 0x30,
	0x55, 0xe2, 0x35, 0x74, 0x2a, 0xe3, 0x30, 0x53, 0xe3, 0x65, 0xef, 0x3d, 0x32, 0xb4, 0x98, 0xbe,
	0x8e, 0x2b, 0x1c, 0x03, 0x14, 0x8b, 0xf1, 0x43, 0x8f, 0x7d, 0x4e, 0xbd, 0x58, 0x7c, 0xe3, 0x2e,
	0x67, 0xd0, 0x31, 0xd6, 0xeb, 0x1b, 0x2d, 0x97, 0x19, 0x24, 0x9c, 0xc1, 0x2a, 0xcd, 0x5b, 0x0c,
	0xcc, 0xc4, 0xde, 0x07, 0xcf, 0x01, 0x6f, 0x31, 0x32, 0xa9, 0x12, 0x7d, 0x68, 0xf1, 0x29, 0xcc,
	0xd0, 0x15, 0xc1, 0x71, 0x48, 0x8e, 0x26, 0x91, 0x57, 0xe8, 0x0a, 0x9e, 0x78, 0xea, 0x6c, 0x35,
	0x1b, 0x17, 0x58, 0x4c, 0x78, 0x9c, 0x23, 0x9e, 0x98, 0xe8, 0xcf, 0xc4, 0x72, 0x5c, 0xec, 0xd3,
	0x2a, 0x39, 0xe6, 0xb8, 0x08, 0xa7, 0x4a, 0xbc, 0x82, 0xb6, 0x36, 0x1e, 0xa7, 0x6e, 0x39, 0xed,
	0x80, 0x3b, 0xac, 0xb0, 0xdc, 0x21, 0xd7, 0x85, 0xa6, 0xd3, 0x38, 0xe1, 0x0e, 0x84, 0x53, 0x25,
	0xce, 0x61, 0x87, 0xb2, 0xc8, 0xbc, 0x77, 0x7a, 0x52, 0x79, 0x8a, 0xfd, 0x39, 0x79, 0x3a, 0x41,
	0x78, 0xf7, 0xc0, 0xc7, 0xab, 0x95, 0xeb, 0xac, 0x0c, 0x96, 0x5e, 0xbc, 0x5a, 0x01, 0xf3, 0xb1,
	0x18, 0xcb, 0xc5, 0xa7, 0x7c, 0x2c, 0x01, 0xa6, 0x2a, 0xdc, 0x71, 0xaf, 0xe5, 0x1d, 0xd2, 0xda,
	0x7d, 0x4e, 0x9b, 0x89, 0x54, 0xf5, 0x7f, 0xc0, 0x66, 0x7c, 0x6e, 0xea, 0x7f, 0x1f, 0xda, 0x3e,
	0x34, 0xd0, 0x39, 0xeb, 0xca, 0xa4, 0xde, 0xab, 0x87, 0x75, 0x19, 0x9d, 0xbf, 0x81, 0xf5, 0x0f,
	0x73, 0x34, 0x5e, 0xec, 0xc1, 0xce, 0xd5, 0xdf, 0x8f, 0xc1, 0xd7, 0x85, 0x91, 0xa8, 0xba, 0x6b,
	0x62, 0x17, 0xba, 0x8f, 0x5f, 0x3a, 0xaa, 0x6e, 0xed, 0xf2, 0x0b, 0x34, 0xc9, 0xfc, 0x91, 0xcc,
	0xe2, 0x2d, 0x6c, 0x44, 0x93, 0x38, 0x1a, 0xae, 0x7e, 0x65, 0x1e, 0x97, 0x1e, 0xee, 0xfd, 0x4b,
	0x54, 0xfd, 0xb5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x38, 0x97, 0x2b, 0xa7, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlaceHolderClient is the client API for PlaceHolder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaceHolderClient interface {
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*Resolved, error)
}

type placeHolderClient struct {
	cc *grpc.ClientConn
}

func NewPlaceHolderClient(cc *grpc.ClientConn) PlaceHolderClient {
	return &placeHolderClient{cc}
}

func (c *placeHolderClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*Resolved, error) {
	out := new(Resolved)
	err := c.cc.Invoke(ctx, "/placeholder.PlaceHolder/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaceHolderServer is the server API for PlaceHolder service.
type PlaceHolderServer interface {
	Resolve(context.Context, *ResolveRequest) (*Resolved, error)
}

func RegisterPlaceHolderServer(s *grpc.Server, srv PlaceHolderServer) {
	s.RegisterService(&_PlaceHolder_serviceDesc, srv)
}

func _PlaceHolder_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaceHolderServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/placeholder.PlaceHolder/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaceHolderServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlaceHolder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "placeholder.PlaceHolder",
	HandlerType: (*PlaceHolderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resolve",
			Handler:    _PlaceHolder_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "placeholder/placeholder.proto",
}