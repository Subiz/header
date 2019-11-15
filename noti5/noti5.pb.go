// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noti5.proto

package noti5

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

func (Token_Platform) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{3, 0}
}

type Subscription struct {
	NewMessage             *bool    `protobuf:"varint,3,opt,name=new_message,json=newMessage" json:"new_message,omitempty"`
	UnassignedConversation *bool    `protobuf:"varint,4,opt,name=unassigned_conversation,json=unassignedConversation" json:"unassigned_conversation,omitempty"`
	Delay                  *int32   `protobuf:"varint,10,opt,name=delay" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{0}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetNewMessage() bool {
	if m != nil && m.NewMessage != nil {
		return *m.NewMessage
	}
	return false
}

func (m *Subscription) GetUnassignedConversation() bool {
	if m != nil && m.UnassignedConversation != nil {
		return *m.UnassignedConversation
	}
	return false
}

func (m *Subscription) GetDelay() int32 {
	if m != nil && m.Delay != nil {
		return *m.Delay
	}
	return 0
}

type Setting struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AgentId              *string         `protobuf:"bytes,3,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Updated              *int64          `protobuf:"varint,8,opt,name=updated" json:"updated,omitempty"`
	Web                  *Subscription   `protobuf:"bytes,9,opt,name=web" json:"web,omitempty"`
	Mobile               *Subscription   `protobuf:"bytes,10,opt,name=mobile" json:"mobile,omitempty"`
	Email                *Subscription   `protobuf:"bytes,11,opt,name=email" json:"email,omitempty"`
	DoNotDisturb         *DoNotDisturb   `protobuf:"bytes,12,opt,name=do_not_disturb,json=doNotDisturb" json:"do_not_disturb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Setting) Reset()         { *m = Setting{} }
func (m *Setting) String() string { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()    {}
func (*Setting) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{1}
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

func (m *Setting) GetUpdated() int64 {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return 0
}

func (m *Setting) GetWeb() *Subscription {
	if m != nil {
		return m.Web
	}
	return nil
}

func (m *Setting) GetMobile() *Subscription {
	if m != nil {
		return m.Mobile
	}
	return nil
}

func (m *Setting) GetEmail() *Subscription {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *Setting) GetDoNotDisturb() *DoNotDisturb {
	if m != nil {
		return m.DoNotDisturb
	}
	return nil
}

type DoNotDisturb struct {
	DailyFrom            *int64   `protobuf:"varint,2,opt,name=daily_from,json=dailyFrom" json:"daily_from,omitempty"`
	DailyTo              *int64   `protobuf:"varint,3,opt,name=daily_to,json=dailyTo" json:"daily_to,omitempty"`
	EnableDaily          *bool    `protobuf:"varint,5,opt,name=enable_daily,json=enableDaily" json:"enable_daily,omitempty"`
	Until                *int64   `protobuf:"varint,4,opt,name=until" json:"until,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoNotDisturb) Reset()         { *m = DoNotDisturb{} }
func (m *DoNotDisturb) String() string { return proto.CompactTextString(m) }
func (*DoNotDisturb) ProtoMessage()    {}
func (*DoNotDisturb) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{2}
}

func (m *DoNotDisturb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoNotDisturb.Unmarshal(m, b)
}
func (m *DoNotDisturb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoNotDisturb.Marshal(b, m, deterministic)
}
func (m *DoNotDisturb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoNotDisturb.Merge(m, src)
}
func (m *DoNotDisturb) XXX_Size() int {
	return xxx_messageInfo_DoNotDisturb.Size(m)
}
func (m *DoNotDisturb) XXX_DiscardUnknown() {
	xxx_messageInfo_DoNotDisturb.DiscardUnknown(m)
}

var xxx_messageInfo_DoNotDisturb proto.InternalMessageInfo

func (m *DoNotDisturb) GetDailyFrom() int64 {
	if m != nil && m.DailyFrom != nil {
		return *m.DailyFrom
	}
	return 0
}

func (m *DoNotDisturb) GetDailyTo() int64 {
	if m != nil && m.DailyTo != nil {
		return *m.DailyTo
	}
	return 0
}

func (m *DoNotDisturb) GetEnableDaily() bool {
	if m != nil && m.EnableDaily != nil {
		return *m.EnableDaily
	}
	return false
}

func (m *DoNotDisturb) GetUntil() int64 {
	if m != nil && m.Until != nil {
		return *m.Until
	}
	return 0
}

type Token struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId               *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserType             *string         `protobuf:"bytes,4,opt,name=user_type,json=userType" json:"user_type,omitempty"`
	FcmToken             *string         `protobuf:"bytes,5,opt,name=fcm_token,json=fcmToken" json:"fcm_token,omitempty"`
	DeviceId             *string         `protobuf:"bytes,6,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Platform             *string         `protobuf:"bytes,7,opt,name=platform" json:"platform,omitempty"`
	Created              *int64          `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Topics               []string        `protobuf:"bytes,9,rep,name=topics" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

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
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId               *string         `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Platform             *string         `protobuf:"bytes,5,opt,name=platform" json:"platform,omitempty"`
	Title                *string         `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Body                 *string         `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	ConversationId       *string         `protobuf:"bytes,8,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	SenderId             *string         `protobuf:"bytes,9,opt,name=sender_id,json=senderId" json:"sender_id,omitempty"`
	SenderType           *string         `protobuf:"bytes,10,opt,name=sender_type,json=senderType" json:"sender_type,omitempty"`
	IconUrl              *string         `protobuf:"bytes,11,opt,name=icon_url,json=iconUrl" json:"icon_url,omitempty"`
	LastPageViewUrl      *string         `protobuf:"bytes,12,opt,name=last_page_view_url,json=lastPageViewUrl" json:"last_page_view_url,omitempty"`
	Type                 *string         `protobuf:"bytes,13,opt,name=type" json:"type,omitempty"`
	Payload              []byte          `protobuf:"bytes,14,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PushNoti) Reset()         { *m = PushNoti{} }
func (m *PushNoti) String() string { return proto.CompactTextString(m) }
func (*PushNoti) ProtoMessage()    {}
func (*PushNoti) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{4}
}

func (m *PushNoti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNoti.Unmarshal(m, b)
}
func (m *PushNoti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNoti.Marshal(b, m, deterministic)
}
func (m *PushNoti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNoti.Merge(m, src)
}
func (m *PushNoti) XXX_Size() int {
	return xxx_messageInfo_PushNoti.Size(m)
}
func (m *PushNoti) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNoti.DiscardUnknown(m)
}

var xxx_messageInfo_PushNoti proto.InternalMessageInfo

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

func (m *PushNoti) GetConversationId() string {
	if m != nil && m.ConversationId != nil {
		return *m.ConversationId
	}
	return ""
}

func (m *PushNoti) GetSenderId() string {
	if m != nil && m.SenderId != nil {
		return *m.SenderId
	}
	return ""
}

func (m *PushNoti) GetSenderType() string {
	if m != nil && m.SenderType != nil {
		return *m.SenderType
	}
	return ""
}

func (m *PushNoti) GetIconUrl() string {
	if m != nil && m.IconUrl != nil {
		return *m.IconUrl
	}
	return ""
}

func (m *PushNoti) GetLastPageViewUrl() string {
	if m != nil && m.LastPageViewUrl != nil {
		return *m.LastPageViewUrl
	}
	return ""
}

func (m *PushNoti) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *PushNoti) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Message struct {
	FcmToken             *string  `protobuf:"bytes,2,opt,name=fcm_token,json=fcmToken" json:"fcm_token,omitempty"`
	Title                *string  `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Body                 *string  `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{5}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

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

type SubscribeStatus struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId               *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Status               *bool           `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SubscribeStatus) Reset()         { *m = SubscribeStatus{} }
func (m *SubscribeStatus) String() string { return proto.CompactTextString(m) }
func (*SubscribeStatus) ProtoMessage()    {}
func (*SubscribeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29df08143386a8, []int{6}
}

func (m *SubscribeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeStatus.Unmarshal(m, b)
}
func (m *SubscribeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeStatus.Marshal(b, m, deterministic)
}
func (m *SubscribeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeStatus.Merge(m, src)
}
func (m *SubscribeStatus) XXX_Size() int {
	return xxx_messageInfo_SubscribeStatus.Size(m)
}
func (m *SubscribeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeStatus proto.InternalMessageInfo

func (m *SubscribeStatus) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *SubscribeStatus) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *SubscribeStatus) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *SubscribeStatus) GetStatus() bool {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return false
}

func init() {
	proto.RegisterEnum("noti5.Token_Platform", Token_Platform_name, Token_Platform_value)
	proto.RegisterType((*Subscription)(nil), "noti5.Subscription")
	proto.RegisterType((*Setting)(nil), "noti5.Setting")
	proto.RegisterType((*DoNotDisturb)(nil), "noti5.DoNotDisturb")
	proto.RegisterType((*Token)(nil), "noti5.Token")
	proto.RegisterType((*PushNoti)(nil), "noti5.PushNoti")
	proto.RegisterType((*Message)(nil), "noti5.Message")
	proto.RegisterType((*SubscribeStatus)(nil), "noti5.SubscribeStatus")
}

func init() { proto.RegisterFile("noti5.proto", fileDescriptor_3a29df08143386a8) }

var fileDescriptor_3a29df08143386a8 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x8a, 0xdc, 0x36,
	0x14, 0xad, 0xc7, 0xe3, 0xb1, 0x7d, 0xc7, 0xdd, 0x09, 0x6a, 0xd8, 0x38, 0x09, 0x21, 0x13, 0x97,
	0xd2, 0x29, 0x81, 0x5d, 0x28, 0x94, 0xd2, 0xd7, 0x66, 0x29, 0xcc, 0x43, 0xc3, 0xe0, 0xdd, 0xf6,
	0xa1, 0x2f, 0x46, 0xb6, 0xee, 0xcc, 0x8a, 0xd8, 0x92, 0xb1, 0xe4, 0x9d, 0x4c, 0x1f, 0x0a, 0x85,
	0x3e, 0xf4, 0x63, 0xfa, 0x05, 0xfd, 0xba, 0x22, 0xc9, 0xd3, 0xf5, 0x2e, 0xdd, 0x42, 0xa1, 0x79,
	0xf3, 0x39, 0xe7, 0xca, 0xf7, 0xea, 0x1c, 0x49, 0x30, 0x17, 0x52, 0xf3, 0xaf, 0xce, 0xda, 0x4e,
	0x6a, 0x49, 0x02, 0x0b, 0x9e, 0x25, 0x95, 0x6c, 0x1a, 0x29, 0x1c, 0x99, 0xfd, 0x02, 0xc9, 0x65,
	0x5f, 0xaa, 0xaa, 0xe3, 0xad, 0xe6, 0x52, 0x90, 0x97, 0x30, 0x17, 0xb8, 0x2f, 0x1a, 0x54, 0x8a,
	0xee, 0x30, 0xf5, 0x97, 0xde, 0x2a, 0xca, 0x41, 0xe0, 0xfe, 0x7b, 0xc7, 0x90, 0xaf, 0xe1, 0x49,
	0x2f, 0xa8, 0x52, 0x7c, 0x27, 0x90, 0x15, 0x95, 0x14, 0x37, 0xd8, 0x29, 0x6a, 0xd6, 0xa6, 0x53,
	0x5b, 0x7c, 0x7a, 0x2b, 0xbf, 0x19, 0xa9, 0xe4, 0x31, 0x04, 0x0c, 0x6b, 0x7a, 0x48, 0x61, 0xe9,
	0xad, 0x82, 0xdc, 0x81, 0xec, 0xcf, 0x09, 0x84, 0x97, 0xa8, 0x35, 0x17, 0x3b, 0xf2, 0x0a, 0xfc,
	0x4a, 0xbf, 0x4f, 0xbd, 0xa5, 0xb7, 0x9a, 0x7f, 0xb9, 0x38, 0x1b, 0xe6, 0x7c, 0x23, 0x85, 0xc6,
	0xf7, 0x3a, 0x37, 0x1a, 0x79, 0x01, 0x40, 0xab, 0x4a, 0xf6, 0x42, 0x17, 0x9c, 0xa5, 0x93, 0xa5,
	0xb7, 0x8a, 0xf3, 0x78, 0x60, 0xd6, 0x8c, 0x3c, 0x85, 0x88, 0xee, 0xd0, 0x89, 0xbe, 0x15, 0x43,
	0x8b, 0xd7, 0x8c, 0xa4, 0x10, 0xf6, 0x2d, 0xa3, 0x1a, 0x59, 0x1a, 0x2d, 0xbd, 0x95, 0x9f, 0x1f,
	0x21, 0xf9, 0x0c, 0xfc, 0x3d, 0x96, 0x69, 0x6c, 0xdb, 0x7e, 0x72, 0xe6, 0x2c, 0x1b, 0x9b, 0x92,
	0x1b, 0x9d, 0xbc, 0x86, 0x59, 0x23, 0x4b, 0x5e, 0xa3, 0xdd, 0xc0, 0x03, 0x95, 0x43, 0x09, 0xf9,
	0x02, 0x02, 0x6c, 0x28, 0xaf, 0xd3, 0xf9, 0xc3, 0xb5, 0xae, 0x82, 0x7c, 0x03, 0x27, 0x4c, 0x16,
	0x42, 0xea, 0x82, 0x71, 0xa5, 0xfb, 0xae, 0x4c, 0x93, 0x3b, 0x6b, 0x2e, 0xe4, 0x5b, 0xa9, 0x2f,
	0x9c, 0x94, 0x27, 0x6c, 0x84, 0xb2, 0x5f, 0x3d, 0x48, 0xc6, 0xb2, 0xb1, 0x87, 0x51, 0x5e, 0x1f,
	0x8a, 0x6d, 0x27, 0x1b, 0x6b, 0x8f, 0x9f, 0xc7, 0x96, 0xf9, 0xae, 0x93, 0x8d, 0xb1, 0xc7, 0xc9,
	0x5a, 0x5a, 0x7b, 0xfc, 0x3c, 0xb4, 0xf8, 0x4a, 0x92, 0x57, 0x90, 0xa0, 0xa0, 0x65, 0x8d, 0x85,
	0x65, 0xd2, 0xc0, 0x66, 0x39, 0x77, 0xdc, 0x85, 0xa1, 0x4c, 0x80, 0xbd, 0xd0, 0xbc, 0xb6, 0x39,
	0xfb, 0xb9, 0x03, 0xd9, 0x1f, 0x13, 0x08, 0xae, 0xe4, 0x3b, 0x14, 0xff, 0x43, 0x7c, 0x4f, 0x20,
	0xec, 0x15, 0x76, 0xb7, 0xe9, 0xcd, 0x0c, 0x5c, 0x33, 0xf2, 0x1c, 0x62, 0x2b, 0xe8, 0x43, 0x8b,
	0xb6, 0x7d, 0x9c, 0x47, 0x86, 0xb8, 0x3a, 0xb4, 0x68, 0xc4, 0x6d, 0xd5, 0x14, 0xda, 0x0c, 0x61,
	0xe7, 0x8e, 0xf3, 0x68, 0x5b, 0x35, 0x6e, 0xa8, 0xe7, 0x10, 0x33, 0xbc, 0xe1, 0x15, 0x9a, 0x9f,
	0xce, 0x9c, 0xe8, 0x88, 0x35, 0x23, 0xcf, 0x20, 0x6a, 0x6b, 0xaa, 0xb7, 0xb2, 0x6b, 0xd2, 0xd0,
	0x69, 0x47, 0x6c, 0xce, 0x4b, 0xd5, 0xe1, 0xf8, 0xbc, 0x0c, 0x90, 0x9c, 0xc2, 0x4c, 0xcb, 0x96,
	0x57, 0x2a, 0x8d, 0x97, 0xbe, 0x19, 0xd2, 0xa1, 0xec, 0x53, 0x88, 0x36, 0xc7, 0xd5, 0x73, 0x08,
	0x19, 0xaa, 0x77, 0x5a, 0xb6, 0x8f, 0x3e, 0x22, 0x70, 0x3c, 0x39, 0x8f, 0xbc, 0xec, 0x77, 0x1f,
	0xa2, 0x4d, 0xaf, 0xae, 0xdf, 0x4a, 0xcd, 0xff, 0xbb, 0x63, 0xfe, 0xbf, 0x38, 0x36, 0xbd, 0xe3,
	0xd8, 0x78, 0x6b, 0xc1, 0xbd, 0xad, 0x3d, 0x86, 0x40, 0x73, 0x5d, 0xe3, 0xe0, 0x87, 0x03, 0x84,
	0xc0, 0xb4, 0x94, 0xec, 0x30, 0x18, 0x61, 0xbf, 0xc9, 0xe7, 0xb0, 0x18, 0xdf, 0x70, 0xd3, 0x26,
	0xb2, 0xf2, 0xc9, 0x98, 0x76, 0x01, 0x29, 0x14, 0xcc, 0x4d, 0x12, 0xbb, 0x7e, 0x8e, 0x58, 0x33,
	0xf3, 0xa6, 0x0c, 0xa2, 0xcd, 0x0f, 0xac, 0x0c, 0x8e, 0xb2, 0x09, 0x3e, 0x85, 0x88, 0x57, 0x52,
	0x14, 0x7d, 0xe7, 0x2e, 0x4c, 0x9c, 0x87, 0x06, 0xff, 0xd0, 0xd5, 0xe4, 0x35, 0x90, 0x9a, 0x2a,
	0x5d, 0xb4, 0x74, 0x87, 0xc5, 0x0d, 0xc7, 0xbd, 0x2d, 0x4a, 0x6c, 0xd1, 0xc2, 0x28, 0x1b, 0xba,
	0xc3, 0x1f, 0x39, 0xee, 0x4d, 0x31, 0x81, 0xa9, 0xed, 0xf0, 0xb1, 0xdb, 0x82, 0xf9, 0x36, 0x39,
	0xb6, 0xf4, 0x50, 0x4b, 0xca, 0xd2, 0x93, 0xa5, 0xb7, 0x4a, 0xf2, 0x23, 0xcc, 0x36, 0x10, 0x1e,
	0x1f, 0xb5, 0x3b, 0x47, 0x68, 0x72, 0xef, 0x08, 0xfd, 0x6d, 0x97, 0xff, 0x4f, 0x76, 0x4d, 0x6f,
	0xed, 0xca, 0x7e, 0xf3, 0x60, 0x31, 0x5c, 0xf1, 0x12, 0x2f, 0x35, 0xd5, 0xbd, 0xfa, 0x90, 0xb7,
	0xe2, 0x14, 0x66, 0xca, 0x36, 0x19, 0x5e, 0xde, 0x01, 0x7d, 0xfb, 0xf2, 0xa7, 0x17, 0x3b, 0xae,
	0xaf, 0xfb, 0xd2, 0x74, 0x3b, 0x57, 0x7d, 0xc9, 0x7f, 0x3e, 0xbf, 0x46, 0xca, 0xb0, 0x3b, 0xb7,
	0x4f, 0xca, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x55, 0xd0, 0x6c, 0x94, 0x17, 0x06, 0x00, 0x00,
}
