// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

package client

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

type Client_Type int32

const (
	Client_app       Client_Type = 0
	Client_connector Client_Type = 1
	Client_bot       Client_Type = 3
)

var Client_Type_name = map[int32]string{
	0: "app",
	1: "connector",
	3: "bot",
}

var Client_Type_value = map[string]int32{
	"app":       0,
	"connector": 1,
	"bot":       3,
}

func (x Client_Type) Enum() *Client_Type {
	p := new(Client_Type)
	*p = x
	return p
}

func (x Client_Type) String() string {
	return proto.EnumName(Client_Type_name, int32(x))
}

func (x *Client_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Client_Type_value, data, "Client_Type")
	if err != nil {
		return err
	}
	*x = Client_Type(value)
	return nil
}

func (Client_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1, 0}
}

type Client_ChannelType int32

const (
	Client_subiz    Client_ChannelType = 0
	Client_email    Client_ChannelType = 1
	Client_facebook Client_ChannelType = 2
	Client_viber    Client_ChannelType = 3
)

var Client_ChannelType_name = map[int32]string{
	0: "subiz",
	1: "email",
	2: "facebook",
	3: "viber",
}

var Client_ChannelType_value = map[string]int32{
	"subiz":    0,
	"email":    1,
	"facebook": 2,
	"viber":    3,
}

func (x Client_ChannelType) Enum() *Client_ChannelType {
	p := new(Client_ChannelType)
	*p = x
	return p
}

func (x Client_ChannelType) String() string {
	return proto.EnumName(Client_ChannelType_name, int32(x))
}

func (x *Client_ChannelType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Client_ChannelType_value, data, "Client_ChannelType")
	if err != nil {
		return err
	}
	*x = Client_ChannelType(value)
	return nil
}

func (Client_ChannelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1, 1}
}

type Webhook_State int32

const (
	Webhook_active        Webhook_State = 0
	Webhook_backoff_sleep Webhook_State = 1
	Webhook_dead          Webhook_State = 2
)

var Webhook_State_name = map[int32]string{
	0: "active",
	1: "backoff_sleep",
	2: "dead",
}

var Webhook_State_value = map[string]int32{
	"active":        0,
	"backoff_sleep": 1,
	"dead":          2,
}

func (x Webhook_State) Enum() *Webhook_State {
	p := new(Webhook_State)
	*p = x
	return p
}

func (x Webhook_State) String() string {
	return proto.EnumName(Webhook_State_name, int32(x))
}

func (x *Webhook_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Webhook_State_value, data, "Webhook_State")
	if err != nil {
		return err
	}
	*x = Webhook_State(value)
	return nil
}

func (Webhook_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{4, 0}
}

type Clients struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Clients              []*Client       `protobuf:"bytes,2,rep,name=clients" json:"clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Clients) Reset()         { *m = Clients{} }
func (m *Clients) String() string { return proto.CompactTextString(m) }
func (*Clients) ProtoMessage()    {}
func (*Clients) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{0}
}

func (m *Clients) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clients.Unmarshal(m, b)
}
func (m *Clients) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clients.Marshal(b, m, deterministic)
}
func (m *Clients) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clients.Merge(m, src)
}
func (m *Clients) XXX_Size() int {
	return xxx_messageInfo_Clients.Size(m)
}
func (m *Clients) XXX_DiscardUnknown() {
	xxx_messageInfo_Clients.DiscardUnknown(m)
}

var xxx_messageInfo_Clients proto.InternalMessageInfo

func (m *Clients) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Clients) GetClients() []*Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

type Client struct {
	Ctx *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Id  *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	// secre used to authorize client with oauth2 server
	Secret *string `protobuf:"bytes,4,opt,name=secret" json:"secret,omitempty"`
	// LogoUrl is url to logo of the client, should be 256x256 and lessthan 256KB
	LogoUrl   *string `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	AccountId *string `protobuf:"bytes,6,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// IsVerified will be true if client is verified by the authority
	IsVerified *bool  `protobuf:"varint,8,opt,name=is_verified,json=isVerified" json:"is_verified,omitempty"`
	Verified   *int64 `protobuf:"varint,9,opt,name=verified" json:"verified,omitempty"`
	// List of URLs which client must register for oauth redirection
	RedirectUri *string `protobuf:"bytes,10,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
	Type        *string `protobuf:"bytes,11,opt,name=type" json:"type,omitempty"`
	Name        *string `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
	// Version number of the client.
	Version    *string  `protobuf:"bytes,14,opt,name=version" json:"version,omitempty"`
	IsEnabled  *bool    `protobuf:"varint,15,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	Created    *int64   `protobuf:"varint,17,opt,name=created" json:"created,omitempty"`
	Modified   *int64   `protobuf:"varint,18,opt,name=modified" json:"modified,omitempty"`
	WebhookUri *string  `protobuf:"bytes,20,opt,name=webhook_uri,json=webhookUri" json:"webhook_uri,omitempty"`
	Events     []string `protobuf:"bytes,19,rep,name=events" json:"events,omitempty"`
	// for connector only
	ChannelType          *string  `protobuf:"bytes,21,opt,name=channel_type,json=channelType" json:"channel_type,omitempty"`
	AvailabilityUri      *string  `protobuf:"bytes,22,opt,name=availability_uri,json=availabilityUri" json:"availability_uri,omitempty"`
	PingUri              *string  `protobuf:"bytes,23,opt,name=ping_uri,json=pingUri" json:"ping_uri,omitempty"`
	IsInternal           *bool    `protobuf:"varint,24,opt,name=is_internal,json=isInternal" json:"is_internal,omitempty"`
	UnsubscribeUri       *string  `protobuf:"bytes,25,opt,name=unsubscribe_uri,json=unsubscribeUri" json:"unsubscribe_uri,omitempty"`
	Scopes               []string `protobuf:"bytes,26,rep,name=scopes" json:"scopes,omitempty"`
	BotDefaultJobTitle   *string  `protobuf:"bytes,28,opt,name=bot_default_job_title,json=botDefaultJobTitle" json:"bot_default_job_title,omitempty"`
	BotDefaultFullname   *string  `protobuf:"bytes,29,opt,name=bot_default_fullname,json=botDefaultFullname" json:"bot_default_fullname,omitempty"`
	BotIsSupervisor      *bool    `protobuf:"varint,30,opt,name=bot_is_supervisor,json=botIsSupervisor" json:"bot_is_supervisor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Client) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Client) GetSecret() string {
	if m != nil && m.Secret != nil {
		return *m.Secret
	}
	return ""
}

func (m *Client) GetLogoUrl() string {
	if m != nil && m.LogoUrl != nil {
		return *m.LogoUrl
	}
	return ""
}

func (m *Client) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Client) GetIsVerified() bool {
	if m != nil && m.IsVerified != nil {
		return *m.IsVerified
	}
	return false
}

func (m *Client) GetVerified() int64 {
	if m != nil && m.Verified != nil {
		return *m.Verified
	}
	return 0
}

func (m *Client) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
	}
	return ""
}

func (m *Client) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Client) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Client) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *Client) GetIsEnabled() bool {
	if m != nil && m.IsEnabled != nil {
		return *m.IsEnabled
	}
	return false
}

func (m *Client) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Client) GetModified() int64 {
	if m != nil && m.Modified != nil {
		return *m.Modified
	}
	return 0
}

func (m *Client) GetWebhookUri() string {
	if m != nil && m.WebhookUri != nil {
		return *m.WebhookUri
	}
	return ""
}

func (m *Client) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Client) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return ""
}

func (m *Client) GetAvailabilityUri() string {
	if m != nil && m.AvailabilityUri != nil {
		return *m.AvailabilityUri
	}
	return ""
}

func (m *Client) GetPingUri() string {
	if m != nil && m.PingUri != nil {
		return *m.PingUri
	}
	return ""
}

func (m *Client) GetIsInternal() bool {
	if m != nil && m.IsInternal != nil {
		return *m.IsInternal
	}
	return false
}

func (m *Client) GetUnsubscribeUri() string {
	if m != nil && m.UnsubscribeUri != nil {
		return *m.UnsubscribeUri
	}
	return ""
}

func (m *Client) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Client) GetBotDefaultJobTitle() string {
	if m != nil && m.BotDefaultJobTitle != nil {
		return *m.BotDefaultJobTitle
	}
	return ""
}

func (m *Client) GetBotDefaultFullname() string {
	if m != nil && m.BotDefaultFullname != nil {
		return *m.BotDefaultFullname
	}
	return ""
}

func (m *Client) GetBotIsSupervisor() bool {
	if m != nil && m.BotIsSupervisor != nil {
		return *m.BotIsSupervisor
	}
	return false
}

type AuthorizedClient struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Client               *Client         `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
	IssueAccountId       *string         `protobuf:"bytes,3,opt,name=issue_account_id,json=issueAccountId" json:"issue_account_id,omitempty"`
	Issuer               *string         `protobuf:"bytes,4,opt,name=issuer" json:"issuer,omitempty"`
	Scopes               []string        `protobuf:"bytes,5,rep,name=scopes" json:"scopes,omitempty"`
	Kind                 *string         `protobuf:"bytes,6,opt,name=kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthorizedClient) Reset()         { *m = AuthorizedClient{} }
func (m *AuthorizedClient) String() string { return proto.CompactTextString(m) }
func (*AuthorizedClient) ProtoMessage()    {}
func (*AuthorizedClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{2}
}

func (m *AuthorizedClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizedClient.Unmarshal(m, b)
}
func (m *AuthorizedClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizedClient.Marshal(b, m, deterministic)
}
func (m *AuthorizedClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizedClient.Merge(m, src)
}
func (m *AuthorizedClient) XXX_Size() int {
	return xxx_messageInfo_AuthorizedClient.Size(m)
}
func (m *AuthorizedClient) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizedClient.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizedClient proto.InternalMessageInfo

func (m *AuthorizedClient) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *AuthorizedClient) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *AuthorizedClient) GetIssueAccountId() string {
	if m != nil && m.IssueAccountId != nil {
		return *m.IssueAccountId
	}
	return ""
}

func (m *AuthorizedClient) GetIssuer() string {
	if m != nil && m.Issuer != nil {
		return *m.Issuer
	}
	return ""
}

func (m *AuthorizedClient) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *AuthorizedClient) GetKind() string {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return ""
}

type Webhooks struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Webhooks             []*Webhook      `protobuf:"bytes,2,rep,name=webhooks" json:"webhooks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Webhooks) Reset()         { *m = Webhooks{} }
func (m *Webhooks) String() string { return proto.CompactTextString(m) }
func (*Webhooks) ProtoMessage()    {}
func (*Webhooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{3}
}

func (m *Webhooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Webhooks.Unmarshal(m, b)
}
func (m *Webhooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Webhooks.Marshal(b, m, deterministic)
}
func (m *Webhooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Webhooks.Merge(m, src)
}
func (m *Webhooks) XXX_Size() int {
	return xxx_messageInfo_Webhooks.Size(m)
}
func (m *Webhooks) XXX_DiscardUnknown() {
	xxx_messageInfo_Webhooks.DiscardUnknown(m)
}

var xxx_messageInfo_Webhooks proto.InternalMessageInfo

func (m *Webhooks) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Webhooks) GetWebhooks() []*Webhook {
	if m != nil {
		return m.Webhooks
	}
	return nil
}

type Webhook struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,12,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ClientId             *string         `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Url                  *string         `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Secret               *string         `protobuf:"bytes,4,opt,name=secret" json:"secret,omitempty"`
	Events               []string        `protobuf:"bytes,6,rep,name=events" json:"events,omitempty"`
	State                *string         `protobuf:"bytes,7,opt,name=state" json:"state,omitempty"`
	LastHookAt           *int64          `protobuf:"varint,9,opt,name=last_hook_at,json=lastHookAt" json:"last_hook_at,omitempty"`
	EventsCount          *int64          `protobuf:"varint,13,opt,name=events_count,json=eventsCount" json:"events_count,omitempty"`
	LastHookResponse     *string         `protobuf:"bytes,14,opt,name=last_hook_response,json=lastHookResponse" json:"last_hook_response,omitempty"`
	LastHookStatus       *int32          `protobuf:"varint,15,opt,name=last_hook_status,json=lastHookStatus" json:"last_hook_status,omitempty"`
	LastHookPayload      *string         `protobuf:"bytes,16,opt,name=last_hook_payload,json=lastHookPayload" json:"last_hook_payload,omitempty"`
	Modified             *int64          `protobuf:"varint,19,opt,name=modified" json:"modified,omitempty"`
	Created              *int64          `protobuf:"varint,20,opt,name=created" json:"created,omitempty"`
	Creator              *string         `protobuf:"bytes,21,opt,name=creator" json:"creator,omitempty"`
	IsEnabled            *bool           `protobuf:"varint,22,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	IsProtected          *bool           `protobuf:"varint,23,opt,name=is_protected,json=isProtected" json:"is_protected,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Webhook) Reset()         { *m = Webhook{} }
func (m *Webhook) String() string { return proto.CompactTextString(m) }
func (*Webhook) ProtoMessage()    {}
func (*Webhook) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{4}
}

func (m *Webhook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Webhook.Unmarshal(m, b)
}
func (m *Webhook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Webhook.Marshal(b, m, deterministic)
}
func (m *Webhook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Webhook.Merge(m, src)
}
func (m *Webhook) XXX_Size() int {
	return xxx_messageInfo_Webhook.Size(m)
}
func (m *Webhook) XXX_DiscardUnknown() {
	xxx_messageInfo_Webhook.DiscardUnknown(m)
}

var xxx_messageInfo_Webhook proto.InternalMessageInfo

func (m *Webhook) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Webhook) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Webhook) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *Webhook) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Webhook) GetSecret() string {
	if m != nil && m.Secret != nil {
		return *m.Secret
	}
	return ""
}

func (m *Webhook) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Webhook) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *Webhook) GetLastHookAt() int64 {
	if m != nil && m.LastHookAt != nil {
		return *m.LastHookAt
	}
	return 0
}

func (m *Webhook) GetEventsCount() int64 {
	if m != nil && m.EventsCount != nil {
		return *m.EventsCount
	}
	return 0
}

func (m *Webhook) GetLastHookResponse() string {
	if m != nil && m.LastHookResponse != nil {
		return *m.LastHookResponse
	}
	return ""
}

func (m *Webhook) GetLastHookStatus() int32 {
	if m != nil && m.LastHookStatus != nil {
		return *m.LastHookStatus
	}
	return 0
}

func (m *Webhook) GetLastHookPayload() string {
	if m != nil && m.LastHookPayload != nil {
		return *m.LastHookPayload
	}
	return ""
}

func (m *Webhook) GetModified() int64 {
	if m != nil && m.Modified != nil {
		return *m.Modified
	}
	return 0
}

func (m *Webhook) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Webhook) GetCreator() string {
	if m != nil && m.Creator != nil {
		return *m.Creator
	}
	return ""
}

func (m *Webhook) GetIsEnabled() bool {
	if m != nil && m.IsEnabled != nil {
		return *m.IsEnabled
	}
	return false
}

func (m *Webhook) GetIsProtected() bool {
	if m != nil && m.IsProtected != nil {
		return *m.IsProtected
	}
	return false
}

type WebhookPayload struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	HubId                *string         `protobuf:"bytes,3,opt,name=hub_id,json=hubId" json:"hub_id,omitempty"`
	Payload              []byte          `protobuf:"bytes,4,opt,name=payload" json:"payload,omitempty"`
	Event                *string         `protobuf:"bytes,5,opt,name=event" json:"event,omitempty"`
	WebhookId            *string         `protobuf:"bytes,6,opt,name=webhook_id,json=webhookId" json:"webhook_id,omitempty"`
	QueueOffset          *int64          `protobuf:"varint,8,opt,name=queue_offset,json=queueOffset" json:"queue_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WebhookPayload) Reset()         { *m = WebhookPayload{} }
func (m *WebhookPayload) String() string { return proto.CompactTextString(m) }
func (*WebhookPayload) ProtoMessage()    {}
func (*WebhookPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{5}
}

func (m *WebhookPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookPayload.Unmarshal(m, b)
}
func (m *WebhookPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookPayload.Marshal(b, m, deterministic)
}
func (m *WebhookPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookPayload.Merge(m, src)
}
func (m *WebhookPayload) XXX_Size() int {
	return xxx_messageInfo_WebhookPayload.Size(m)
}
func (m *WebhookPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookPayload.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookPayload proto.InternalMessageInfo

func (m *WebhookPayload) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WebhookPayload) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *WebhookPayload) GetHubId() string {
	if m != nil && m.HubId != nil {
		return *m.HubId
	}
	return ""
}

func (m *WebhookPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *WebhookPayload) GetEvent() string {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return ""
}

func (m *WebhookPayload) GetWebhookId() string {
	if m != nil && m.WebhookId != nil {
		return *m.WebhookId
	}
	return ""
}

func (m *WebhookPayload) GetQueueOffset() int64 {
	if m != nil && m.QueueOffset != nil {
		return *m.QueueOffset
	}
	return 0
}

type WebhookBackoff struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	WorkerId             *string         `protobuf:"bytes,3,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	NumBackOff           *int32          `protobuf:"varint,5,opt,name=num_back_off,json=numBackOff" json:"num_back_off,omitempty"`
	BackoffId            *string         `protobuf:"bytes,6,opt,name=backoff_id,json=backoffId" json:"backoff_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WebhookBackoff) Reset()         { *m = WebhookBackoff{} }
func (m *WebhookBackoff) String() string { return proto.CompactTextString(m) }
func (*WebhookBackoff) ProtoMessage()    {}
func (*WebhookBackoff) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{6}
}

func (m *WebhookBackoff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookBackoff.Unmarshal(m, b)
}
func (m *WebhookBackoff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookBackoff.Marshal(b, m, deterministic)
}
func (m *WebhookBackoff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookBackoff.Merge(m, src)
}
func (m *WebhookBackoff) XXX_Size() int {
	return xxx_messageInfo_WebhookBackoff.Size(m)
}
func (m *WebhookBackoff) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookBackoff.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookBackoff proto.InternalMessageInfo

func (m *WebhookBackoff) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WebhookBackoff) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *WebhookBackoff) GetWorkerId() string {
	if m != nil && m.WorkerId != nil {
		return *m.WorkerId
	}
	return ""
}

func (m *WebhookBackoff) GetNumBackOff() int32 {
	if m != nil && m.NumBackOff != nil {
		return *m.NumBackOff
	}
	return 0
}

func (m *WebhookBackoff) GetBackoffId() string {
	if m != nil && m.BackoffId != nil {
		return *m.BackoffId
	}
	return ""
}

type WebhookTestResult struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ClientId             *string         `protobuf:"bytes,3,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ResponseBody         *string         `protobuf:"bytes,14,opt,name=response_body,json=responseBody" json:"response_body,omitempty"`
	Status               *int32          `protobuf:"varint,15,opt,name=status" json:"status,omitempty"`
	Payload              *string         `protobuf:"bytes,16,opt,name=payload" json:"payload,omitempty"`
	Latency              *int64          `protobuf:"varint,17,opt,name=latency" json:"latency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WebhookTestResult) Reset()         { *m = WebhookTestResult{} }
func (m *WebhookTestResult) String() string { return proto.CompactTextString(m) }
func (*WebhookTestResult) ProtoMessage()    {}
func (*WebhookTestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{7}
}

func (m *WebhookTestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookTestResult.Unmarshal(m, b)
}
func (m *WebhookTestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookTestResult.Marshal(b, m, deterministic)
}
func (m *WebhookTestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookTestResult.Merge(m, src)
}
func (m *WebhookTestResult) XXX_Size() int {
	return xxx_messageInfo_WebhookTestResult.Size(m)
}
func (m *WebhookTestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookTestResult.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookTestResult proto.InternalMessageInfo

func (m *WebhookTestResult) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WebhookTestResult) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *WebhookTestResult) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *WebhookTestResult) GetResponseBody() string {
	if m != nil && m.ResponseBody != nil {
		return *m.ResponseBody
	}
	return ""
}

func (m *WebhookTestResult) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *WebhookTestResult) GetPayload() string {
	if m != nil && m.Payload != nil {
		return *m.Payload
	}
	return ""
}

func (m *WebhookTestResult) GetLatency() int64 {
	if m != nil && m.Latency != nil {
		return *m.Latency
	}
	return 0
}

func init() {
	proto.RegisterEnum("client.Client_Type", Client_Type_name, Client_Type_value)
	proto.RegisterEnum("client.Client_ChannelType", Client_ChannelType_name, Client_ChannelType_value)
	proto.RegisterEnum("client.Webhook_State", Webhook_State_name, Webhook_State_value)
	proto.RegisterType((*Clients)(nil), "client.Clients")
	proto.RegisterType((*Client)(nil), "client.Client")
	proto.RegisterType((*AuthorizedClient)(nil), "client.AuthorizedClient")
	proto.RegisterType((*Webhooks)(nil), "client.Webhooks")
	proto.RegisterType((*Webhook)(nil), "client.Webhook")
	proto.RegisterType((*WebhookPayload)(nil), "client.WebhookPayload")
	proto.RegisterType((*WebhookBackoff)(nil), "client.WebhookBackoff")
	proto.RegisterType((*WebhookTestResult)(nil), "client.WebhookTestResult")
}

func init() { proto.RegisterFile("client/client.proto", fileDescriptor_4d3551c163a1d198) }

var fileDescriptor_4d3551c163a1d198 = []byte{
	// 1123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x5b, 0x45,
	0x10, 0xae, 0xff, 0xed, 0xb1, 0x63, 0x9f, 0x6c, 0x93, 0x74, 0x93, 0x52, 0x70, 0x8c, 0x44, 0x0d,
	0x45, 0x49, 0xe9, 0x35, 0x37, 0x69, 0x00, 0x61, 0x6e, 0x5a, 0x9d, 0xa4, 0x45, 0xe2, 0xe6, 0xe8,
	0xfc, 0x8c, 0xe3, 0xc5, 0xc7, 0x67, 0xcd, 0xee, 0x1e, 0xb7, 0xee, 0x15, 0x8f, 0xc4, 0x7b, 0x20,
	0x9e, 0x80, 0x4b, 0x5e, 0x04, 0xed, 0x5f, 0xec, 0x58, 0x42, 0x0a, 0x52, 0xaf, 0x7c, 0xe6, 0x9b,
	0xd9, 0xd9, 0xf9, 0x76, 0xf6, 0x9b, 0x35, 0x3c, 0x4c, 0x73, 0x86, 0x85, 0x3a, 0xb7, 0x3f, 0x67,
	0x4b, 0xc1, 0x15, 0x27, 0x4d, 0x6b, 0x9d, 0x3c, 0xbb, 0x61, 0x6a, 0x56, 0x26, 0x67, 0x29, 0x5f,
	0x9c, 0xcb, 0x32, 0x61, 0x1f, 0xce, 0x67, 0x18, 0x67, 0x28, 0xce, 0x53, 0xbe, 0x58, 0xf0, 0xc2,
	0xfd, 0xd8, 0x45, 0xa3, 0xb7, 0xd0, 0xba, 0x34, 0xcb, 0x24, 0x39, 0x85, 0x5a, 0xaa, 0xde, 0xd3,
	0xca, 0xb0, 0x32, 0xee, 0xbe, 0x18, 0x9c, 0xb9, 0xb0, 0x4b, 0x5e, 0x28, 0x7c, 0xaf, 0x42, 0xed,
	0x23, 0x63, 0x68, 0xd9, 0x4d, 0x24, 0xad, 0x0e, 0x6b, 0xe3, 0xee, 0x8b, 0xfe, 0x99, 0x2b, 0xc1,
	0x26, 0x09, 0xbd, 0x7b, 0xf4, 0x7b, 0x0b, 0x9a, 0x16, 0xbb, 0x4f, 0xde, 0x3e, 0x54, 0x59, 0x46,
	0x6b, 0xc3, 0xca, 0xb8, 0x13, 0x56, 0x59, 0x46, 0x8e, 0xa0, 0x29, 0x31, 0x15, 0xa8, 0x68, 0xdd,
	0x60, 0xce, 0x22, 0xc7, 0xd0, 0xce, 0xf9, 0x0d, 0x8f, 0x4a, 0x91, 0xd3, 0x86, 0xf1, 0xb4, 0xb4,
	0xfd, 0x46, 0xe4, 0xe4, 0x09, 0x40, 0x9c, 0xa6, 0xbc, 0x2c, 0x54, 0xc4, 0x32, 0xda, 0x34, 0xce,
	0x8e, 0x43, 0x26, 0x19, 0xf9, 0x0c, 0xba, 0x4c, 0x46, 0x2b, 0x14, 0x6c, 0xca, 0x30, 0xa3, 0xed,
	0x61, 0x65, 0xdc, 0x0e, 0x81, 0xc9, 0xb7, 0x0e, 0x21, 0x27, 0xd0, 0xbe, 0xf5, 0x76, 0x86, 0x95,
	0x71, 0x2d, 0xbc, 0xb5, 0xc9, 0x29, 0xf4, 0x04, 0x66, 0x4c, 0x60, 0xaa, 0xa2, 0x52, 0x30, 0x0a,
	0x26, 0x7b, 0xd7, 0x63, 0x6f, 0x04, 0x23, 0x04, 0xea, 0x6a, 0xbd, 0x44, 0xda, 0x35, 0x2e, 0xf3,
	0xad, 0xb1, 0x22, 0x5e, 0x20, 0xed, 0x59, 0x4c, 0x7f, 0x13, 0x0a, 0xad, 0x15, 0x0a, 0xc9, 0x78,
	0x41, 0xfb, 0x96, 0x80, 0x33, 0x35, 0x01, 0x26, 0x23, 0x2c, 0xe2, 0x24, 0xc7, 0x8c, 0x0e, 0x4c,
	0x81, 0x1d, 0x26, 0xbf, 0xb7, 0x80, 0x5e, 0x98, 0x0a, 0x8c, 0x15, 0x66, 0x74, 0xdf, 0x94, 0xe7,
	0x4d, 0x5d, 0xf9, 0x82, 0x67, 0xb6, 0x72, 0x62, 0x2b, 0xf7, 0xb6, 0xa6, 0xfd, 0x0e, 0x93, 0x19,
	0xe7, 0x73, 0x53, 0xf8, 0x81, 0xd9, 0x12, 0x1c, 0xa4, 0xeb, 0x3e, 0x82, 0x26, 0xae, 0x4c, 0x43,
	0x1f, 0x0e, 0x6b, 0xfa, 0xa4, 0xad, 0xa5, 0x29, 0xa7, 0xb3, 0xb8, 0x28, 0x30, 0x8f, 0x0c, 0xaf,
	0x43, 0x4b, 0xd9, 0x61, 0xd7, 0x9a, 0xde, 0x97, 0x10, 0xc4, 0xab, 0x98, 0xe5, 0x71, 0xc2, 0x72,
	0xa6, 0xd6, 0x66, 0x83, 0x23, 0x13, 0x36, 0xd8, 0xc6, 0xf5, 0x2e, 0xc7, 0xd0, 0x5e, 0xb2, 0xe2,
	0xc6, 0x84, 0x3c, 0xb2, 0xb4, 0xb5, 0xad, 0x5d, 0xb6, 0x31, 0xac, 0x50, 0x28, 0x8a, 0x38, 0xa7,
	0xd4, 0x37, 0x66, 0xe2, 0x10, 0xf2, 0x14, 0x06, 0x65, 0x21, 0xcb, 0x44, 0xa6, 0x82, 0x25, 0x68,
	0x52, 0x1c, 0x9b, 0x14, 0xfd, 0x2d, 0xd8, 0x51, 0x91, 0x29, 0x5f, 0xa2, 0xa4, 0x27, 0x96, 0x8a,
	0xb5, 0xc8, 0x37, 0x70, 0x98, 0x70, 0x15, 0x65, 0x38, 0x8d, 0xcb, 0x5c, 0x45, 0xbf, 0xf2, 0x24,
	0x52, 0x4c, 0xe5, 0x48, 0x3f, 0x31, 0x69, 0x48, 0xc2, 0xd5, 0x77, 0xd6, 0xf7, 0x13, 0x4f, 0xae,
	0xb5, 0x87, 0x3c, 0x87, 0x83, 0xed, 0x25, 0xd3, 0x32, 0xcf, 0x4d, 0x27, 0x9f, 0xec, 0xae, 0xf8,
	0xc1, 0x79, 0xc8, 0x57, 0xb0, 0xaf, 0x57, 0x30, 0x19, 0xc9, 0x72, 0x89, 0x62, 0xc5, 0x24, 0x17,
	0xf4, 0x53, 0x43, 0x66, 0x90, 0x70, 0x35, 0x91, 0x57, 0xb7, 0xf0, 0xe8, 0x29, 0xd4, 0xcd, 0x01,
	0xb6, 0xa0, 0x16, 0x2f, 0x97, 0xc1, 0x03, 0xb2, 0x07, 0x9d, 0x94, 0x17, 0x05, 0xa6, 0x8a, 0x8b,
	0xa0, 0xa2, 0xf1, 0x84, 0xab, 0xa0, 0x36, 0xfa, 0x16, 0xba, 0x97, 0x5b, 0x07, 0xde, 0x81, 0x86,
	0xd1, 0x73, 0xf0, 0x40, 0x7f, 0xe2, 0x22, 0x66, 0x79, 0x50, 0x21, 0x3d, 0x68, 0x4f, 0xe3, 0x14,
	0x13, 0xce, 0xe7, 0x41, 0x55, 0x3b, 0x56, 0x2c, 0x41, 0x11, 0xd4, 0x46, 0x7f, 0x56, 0x20, 0xb8,
	0x28, 0xd5, 0x8c, 0x0b, 0xf6, 0x01, 0xb3, 0xfb, 0x8b, 0xf1, 0x0b, 0x70, 0x93, 0x84, 0x56, 0x4d,
	0xd4, 0xae, 0xc6, 0x9d, 0x97, 0x8c, 0x21, 0x60, 0x52, 0x96, 0x18, 0x6d, 0xe9, 0xce, 0x4a, 0xb8,
	0x6f, 0xf0, 0x8b, 0x5b, 0xf1, 0x1d, 0x41, 0xd3, 0x20, 0xc2, 0xcb, 0xd9, 0x5a, 0x5b, 0x1d, 0x6b,
	0xdc, 0xe9, 0x18, 0x81, 0xfa, 0x9c, 0x15, 0x5e, 0xc5, 0xe6, 0x7b, 0xf4, 0x0b, 0xb4, 0x7f, 0xb6,
	0xd7, 0xf6, 0x5e, 0x93, 0xea, 0x19, 0xb4, 0xdd, 0x2d, 0xf7, 0xa3, 0x6a, 0xe0, 0x69, 0xb8, 0x34,
	0xe1, 0x6d, 0xc0, 0xe8, 0xaf, 0x3a, 0xb4, 0x1c, 0x7a, 0x9f, 0xdc, 0x77, 0x47, 0x4d, 0x6f, 0x77,
	0xd4, 0x3c, 0x86, 0x8e, 0xdd, 0x49, 0x7b, 0xab, 0xc6, 0xdb, 0xb6, 0xc0, 0x24, 0x23, 0x01, 0xd4,
	0xf4, 0xf0, 0xb2, 0xe7, 0xa4, 0x3f, 0xff, 0x73, 0xd6, 0x6d, 0x94, 0xd9, 0xbc, 0xa3, 0xcc, 0x03,
	0x68, 0x48, 0x15, 0x2b, 0xa4, 0x2d, 0x13, 0x6e, 0x0d, 0x32, 0x84, 0x5e, 0x1e, 0x4b, 0x15, 0x19,
	0xa9, 0xc7, 0xca, 0x8d, 0x30, 0xd0, 0xd8, 0x8f, 0x9c, 0xcf, 0x2f, 0x74, 0xe7, 0x7b, 0x36, 0x43,
	0x64, 0x0a, 0xa5, 0x7b, 0x26, 0xa2, 0x6b, 0xb1, 0x4b, 0x0d, 0x91, 0xaf, 0x81, 0x6c, 0x92, 0x08,
	0x94, 0x4b, 0x5e, 0x48, 0x74, 0x73, 0x2a, 0xf0, 0xa9, 0x42, 0x87, 0xeb, 0xfe, 0x6f, 0xa2, 0x75,
	0x15, 0xa5, 0x34, 0x63, 0xab, 0x11, 0xf6, 0x7d, 0xec, 0x95, 0x41, 0xb5, 0x38, 0x36, 0x91, 0xcb,
	0x78, 0x9d, 0xf3, 0x38, 0xa3, 0x81, 0x1d, 0x15, 0x3e, 0xf4, 0xb5, 0x85, 0xef, 0x4c, 0xb3, 0x87,
	0x3b, 0xd3, 0x6c, 0x6b, 0x06, 0x1e, 0xdc, 0x9d, 0x81, 0xde, 0xc3, 0x85, 0x9b, 0x54, 0xde, 0xdc,
	0x19, 0xab, 0x47, 0xbb, 0x63, 0xf5, 0x14, 0x7a, 0x4c, 0x46, 0xfa, 0x2d, 0xc4, 0x54, 0xe7, 0x7d,
	0x64, 0x02, 0xba, 0x4c, 0xbe, 0xf6, 0xd0, 0xe8, 0x39, 0x34, 0xae, 0xcc, 0x19, 0x03, 0x34, 0xe3,
	0x54, 0xb1, 0x15, 0x06, 0x0f, 0xc8, 0x3e, 0xec, 0x25, 0x71, 0x3a, 0xe7, 0xd3, 0x69, 0x24, 0x73,
	0xc4, 0x65, 0x50, 0x21, 0x6d, 0xa8, 0x67, 0x18, 0x67, 0x41, 0x75, 0xf4, 0x77, 0x05, 0xfa, 0xee,
	0x3e, 0x79, 0x5a, 0xff, 0xfb, 0x5a, 0x55, 0x77, 0xaf, 0xd5, 0x21, 0x34, 0x67, 0x65, 0xb2, 0x11,
	0x59, 0x63, 0x56, 0x26, 0x13, 0xc3, 0xdc, 0x9f, 0xa8, 0xbe, 0x3f, 0xbd, 0xd0, 0x9b, 0xfa, 0xa2,
	0x98, 0xe6, 0xba, 0x97, 0xd2, 0x1a, 0x7a, 0x17, 0xff, 0x22, 0x6c, 0xde, 0x49, 0x87, 0x4c, 0xcc,
	0x79, 0xfc, 0x56, 0x62, 0x89, 0x11, 0x9f, 0x4e, 0x25, 0x2a, 0xf3, 0x50, 0xd6, 0xc2, 0xae, 0xc1,
	0x5e, 0x19, 0x68, 0xf4, 0xc7, 0x86, 0xdd, 0x4b, 0x7b, 0x04, 0x1f, 0x81, 0xdd, 0x63, 0xe8, 0xbc,
	0xe3, 0x62, 0x8e, 0x62, 0x43, 0xb0, 0x6d, 0x81, 0x49, 0xa6, 0x2f, 0x77, 0x51, 0x2e, 0x22, 0x7d,
	0xe0, 0xba, 0x2e, 0x43, 0xa8, 0x11, 0x42, 0x51, 0x2e, 0x74, 0x01, 0xaf, 0xa6, 0x53, 0x9d, 0xdd,
	0xb7, 0x63, 0xc3, 0xca, 0x21, 0x93, 0x6c, 0xf4, 0x4f, 0x05, 0xf6, 0x5d, 0xc9, 0xd7, 0x28, 0x55,
	0x88, 0xb2, 0xcc, 0xd5, 0xc7, 0xa9, 0x7a, 0x23, 0xf5, 0xda, 0x8e, 0xd4, 0x3f, 0x87, 0x3d, 0xaf,
	0xa1, 0x28, 0xe1, 0xd9, 0xda, 0x09, 0xa9, 0xe7, 0xc1, 0x97, 0x3c, 0x5b, 0x1b, 0xf5, 0x6f, 0x4b,
	0xc7, 0x59, 0xdb, 0x6d, 0x0d, 0xdc, 0x83, 0xe9, 0xda, 0x4a, 0xa1, 0x95, 0xc7, 0x0a, 0x8b, 0x74,
	0xed, 0xff, 0x08, 0x38, 0xf3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe9, 0x44, 0xb0, 0x16,
	0x0a, 0x00, 0x00,
}
