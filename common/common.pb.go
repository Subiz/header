// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	auth "git.subiz.net/header/auth"
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

func (Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

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

func (Device_DeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{9, 0}
}

// A string
type String struct {
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// An int64 value
type Int64 struct {
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (m *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(m, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// A bool value
type Bool struct {
	Value                bool     `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (m *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(m, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type Empty struct {
	Ctx                  *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

type Id struct {
	Ctx                  *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

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
	Ctx                  *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Ids                  []string `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ids) Reset()         { *m = Ids{} }
func (m *Ids) String() string { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()    {}
func (*Ids) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{5}
}

func (m *Ids) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ids.Unmarshal(m, b)
}
func (m *Ids) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ids.Marshal(b, m, deterministic)
}
func (m *Ids) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ids.Merge(m, src)
}
func (m *Ids) XXX_Size() int {
	return xxx_messageInfo_Ids.Size(m)
}
func (m *Ids) XXX_DiscardUnknown() {
	xxx_messageInfo_Ids.DiscardUnknown(m)
}

var xxx_messageInfo_Ids proto.InternalMessageInfo

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

type ObjectPathRequest struct {
	Ctx                  *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectPathRequest) Reset()         { *m = ObjectPathRequest{} }
func (m *ObjectPathRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectPathRequest) ProtoMessage()    {}
func (*ObjectPathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{6}
}

func (m *ObjectPathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectPathRequest.Unmarshal(m, b)
}
func (m *ObjectPathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectPathRequest.Marshal(b, m, deterministic)
}
func (m *ObjectPathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectPathRequest.Merge(m, src)
}
func (m *ObjectPathRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectPathRequest.Size(m)
}
func (m *ObjectPathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectPathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectPathRequest proto.InternalMessageInfo

func (m *ObjectPathRequest) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ObjectPathRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectPathRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Context struct {
	EventId    string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	State      []byte `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Node       string `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	ReplyTopic string `protobuf:"bytes,4,opt,name=reply_topic,json=replyTopic,proto3" json:"reply_topic,omitempty"`
	// 	optional int32 reply_partition = 5;
	Credential *auth.Credential `protobuf:"bytes,6,opt,name=credential,proto3" json:"credential,omitempty"`
	Tracing    []byte           `protobuf:"bytes,7,opt,name=tracing,proto3" json:"tracing,omitempty"`
	ReplyKey   string           `protobuf:"bytes,8,opt,name=reply_key,json=replyKey,proto3" json:"reply_key,omitempty"`
	ByDevice   *Device          `protobuf:"bytes,10,opt,name=by_device,json=byDevice,proto3" json:"by_device,omitempty"`
	// for kafka
	Topic                string   `protobuf:"bytes,11,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition            int32    `protobuf:"varint,13,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset               int64    `protobuf:"varint,14,opt,name=offset,proto3" json:"offset,omitempty"`
	Term                 uint64   `protobuf:"varint,15,opt,name=term,proto3" json:"term,omitempty"`
	RouterTopic          string   `protobuf:"bytes,16,opt,name=router_topic,json=routerTopic,proto3" json:"router_topic,omitempty"`
	IdempotencyKey       string   `protobuf:"bytes,17,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{7}
}

func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

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

func (m *Context) GetIdempotencyKey() string {
	if m != nil {
		return m.IdempotencyKey
	}
	return ""
}

type Reply struct {
	Ctx            *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	EventId        string   `protobuf:"bytes,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	State          []byte   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Service        string   `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	ServiceId      string   `protobuf:"bytes,6,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Err            bool     `protobuf:"varint,10,opt,name=err,proto3" json:"err,omitempty"`
	ErrDescription string   `protobuf:"bytes,12,opt,name=err_description,json=errDescription,proto3" json:"err_description,omitempty"`
	// lang.T err_code = 13;
	ErrClass             int32    `protobuf:"varint,15,opt,name=err_class,json=errClass,proto3" json:"err_class,omitempty"`
	ErrHash              string   `protobuf:"bytes,16,opt,name=err_hash,json=errHash,proto3" json:"err_hash,omitempty"`
	Payload              []byte   `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
	Published            int64    `protobuf:"varint,20,opt,name=published,proto3" json:"published,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{8}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

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
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	UserAgent            string   `protobuf:"bytes,4,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	ScreenResolution     string   `protobuf:"bytes,5,opt,name=screen_resolution,json=screenResolution,proto3" json:"screen_resolution,omitempty"`
	Timezone             string   `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Language             string   `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	Referrer             string   `protobuf:"bytes,8,opt,name=referrer,proto3" json:"referrer,omitempty"`
	Type                 string   `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{9}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

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
	Ctx                  *Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Debug                string   `protobuf:"bytes,3,opt,name=debug,proto3" json:"debug,omitempty"`
	Hash                 string   `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Class                int32    `protobuf:"varint,6,opt,name=class,proto3" json:"class,omitempty"`
	Stack                string   `protobuf:"bytes,7,opt,name=stack,proto3" json:"stack,omitempty"`
	Created              int64    `protobuf:"varint,8,opt,name=created,proto3" json:"created,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{10}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

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
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{11}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Pong struct {
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{12}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.Event", Event_name, Event_value)
	proto.RegisterEnum("common.Device_DeviceType", Device_DeviceType_name, Device_DeviceType_value)
	proto.RegisterType((*String)(nil), "common.String")
	proto.RegisterType((*Int64)(nil), "common.Int64")
	proto.RegisterType((*Bool)(nil), "common.Bool")
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*Id)(nil), "common.Id")
	proto.RegisterType((*Ids)(nil), "common.Ids")
	proto.RegisterType((*ObjectPathRequest)(nil), "common.ObjectPathRequest")
	proto.RegisterType((*Context)(nil), "common.Context")
	proto.RegisterType((*Reply)(nil), "common.Reply")
	proto.RegisterType((*Device)(nil), "common.Device")
	proto.RegisterType((*Error)(nil), "common.Error")
	proto.RegisterType((*PingRequest)(nil), "common.PingRequest")
	proto.RegisterType((*Pong)(nil), "common.Pong")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 905 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0xdb, 0x36,
	0x14, 0xad, 0x2d, 0xcb, 0xb6, 0xae, 0xb3, 0x44, 0x61, 0x8b, 0x41, 0xcb, 0xda, 0xcd, 0x15, 0x06,
	0xd4, 0x6b, 0x01, 0x67, 0xc8, 0x86, 0x3d, 0x0e, 0xe8, 0xd2, 0x00, 0x33, 0xfa, 0xb0, 0x40, 0x2d,
	0x06, 0x6c, 0x2f, 0x06, 0x25, 0xde, 0xc8, 0x5c, 0x64, 0x52, 0x23, 0xa9, 0xac, 0xea, 0x37, 0xec,
	0x67, 0xf6, 0x05, 0xfb, 0x85, 0x7d, 0xd2, 0x40, 0x52, 0x72, 0x9c, 0x37, 0x3f, 0xec, 0x25, 0xb9,
	0xe7, 0xdc, 0x1b, 0xf2, 0xde, 0x73, 0x0f, 0x15, 0x78, 0x5c, 0xc8, 0xed, 0x56, 0x8a, 0x73, 0xff,
	0x6b, 0x59, 0x2b, 0x69, 0x24, 0x19, 0x7b, 0x74, 0xf6, 0x55, 0xc9, 0xcd, 0x52, 0x37, 0x39, 0xff,
	0xb8, 0x14, 0x68, 0xce, 0x37, 0x48, 0x19, 0xaa, 0x73, 0xda, 0x98, 0x8d, 0xfb, 0xe1, 0xab, 0xd3,
	0x2f, 0x60, 0xfc, 0xce, 0x28, 0x2e, 0x4a, 0xf2, 0x04, 0xc2, 0x3b, 0x5a, 0x35, 0x98, 0x0c, 0xe7,
	0x83, 0x45, 0x94, 0x79, 0x90, 0x3e, 0x83, 0x70, 0x25, 0xcc, 0xf7, 0xdf, 0x3d, 0x4c, 0x07, 0x7d,
	0xfa, 0x29, 0x8c, 0x7e, 0x94, 0xb2, 0x7a, 0x98, 0x9d, 0xf6, 0xd9, 0x97, 0x10, 0x5e, 0x6d, 0x6b,
	0xd3, 0x92, 0xe7, 0x10, 0x14, 0xe6, 0x43, 0x32, 0x98, 0x0f, 0x16, 0xb3, 0x8b, 0x93, 0x65, 0xd7,
	0xef, 0xa5, 0x14, 0x06, 0x3f, 0x98, 0xcc, 0xe6, 0xd2, 0x5f, 0x60, 0xb8, 0x62, 0x07, 0x14, 0x92,
	0x67, 0x00, 0xb4, 0x28, 0x64, 0x23, 0xcc, 0x9a, 0xb3, 0xae, 0xd9, 0xa8, 0x63, 0x56, 0x8c, 0x1c,
	0xc3, 0x90, 0xb3, 0x24, 0x70, 0xf4, 0x90, 0xb3, 0xf4, 0x57, 0x08, 0x56, 0x4c, 0xff, 0x0f, 0x07,
	0xc7, 0x10, 0x70, 0xa6, 0x93, 0x60, 0x1e, 0x2c, 0xa2, 0xcc, 0x86, 0xe9, 0x6f, 0x70, 0xfa, 0x73,
	0xfe, 0x3b, 0x16, 0xe6, 0x9a, 0x9a, 0x4d, 0x86, 0x7f, 0x34, 0xa8, 0xcd, 0x21, 0x17, 0xf9, 0x16,
	0x87, 0x7d, 0x8b, 0x84, 0xc0, 0xa8, 0xa6, 0x66, 0xd3, 0x35, 0xed, 0xe2, 0xf4, 0xef, 0x00, 0x26,
	0xdd, 0x1f, 0x91, 0xcf, 0x60, 0x8a, 0x77, 0xe8, 0xdb, 0x1a, 0xb8, 0x9a, 0x89, 0xc3, 0x2b, 0x66,
	0x75, 0xd7, 0x86, 0x1a, 0xaf, 0xfb, 0x51, 0xe6, 0x81, 0x3d, 0x50, 0x48, 0x86, 0xfd, 0x81, 0x36,
	0x26, 0x5f, 0xc2, 0x4c, 0x61, 0x5d, 0xb5, 0x6b, 0x23, 0x6b, 0x5e, 0x24, 0x23, 0x97, 0x02, 0x47,
	0xbd, 0xb7, 0x0c, 0xf9, 0x06, 0xa0, 0x50, 0xc8, 0x50, 0x18, 0x4e, 0xab, 0x64, 0xec, 0xfa, 0x8f,
	0x97, 0xce, 0x2a, 0x97, 0x3b, 0x3e, 0xdb, 0xab, 0x21, 0x09, 0x4c, 0x8c, 0xa2, 0x05, 0x17, 0x65,
	0x32, 0x71, 0xd7, 0xf7, 0x90, 0x7c, 0x0e, 0x91, 0xbf, 0xec, 0x16, 0xdb, 0x64, 0xea, 0xae, 0x9a,
	0x3a, 0xe2, 0x2d, 0xb6, 0xe4, 0x15, 0x44, 0x79, 0xbb, 0x66, 0x78, 0xc7, 0x0b, 0x4c, 0xc0, 0xdd,
	0x73, 0xdc, 0xeb, 0xf4, 0xc6, 0xb1, 0xd9, 0x34, 0x6f, 0x7d, 0x64, 0x07, 0xf4, 0x0d, 0xcf, 0xbc,
	0x2b, 0x1d, 0x20, 0x4f, 0x21, 0xaa, 0xa9, 0x32, 0xdc, 0x70, 0x29, 0x92, 0x4f, 0xe6, 0x83, 0x45,
	0x98, 0xdd, 0x13, 0xe4, 0x53, 0x18, 0xcb, 0x9b, 0x1b, 0x8d, 0x26, 0x39, 0x76, 0x5e, 0xed, 0x90,
	0x95, 0xc5, 0xa0, 0xda, 0x26, 0x27, 0xf3, 0xc1, 0x62, 0x94, 0xb9, 0x98, 0x3c, 0x87, 0x23, 0x25,
	0x1b, 0x83, 0xaa, 0xd3, 0x25, 0x76, 0xd7, 0xcc, 0x3c, 0xe7, 0x85, 0x79, 0x01, 0x27, 0x9c, 0xe1,
	0xb6, 0x96, 0x06, 0x45, 0xe1, 0x47, 0x3a, 0x75, 0x55, 0xc7, 0x7b, 0xf4, 0x5b, 0x6c, 0xd3, 0x7f,
	0x86, 0x10, 0x66, 0x76, 0xca, 0x43, 0x4c, 0xb0, 0xbf, 0xd4, 0xe0, 0x90, 0xa5, 0x26, 0x30, 0xd1,
	0xa8, 0x9c, 0x68, 0xa1, 0xaf, 0xef, 0xa0, 0x35, 0x6e, 0x17, 0xda, 0xc3, 0xc6, 0xde, 0xb8, 0x1d,
	0xe3, 0x8d, 0x8b, 0x4a, 0x39, 0xa5, 0xa7, 0x99, 0x0d, 0xed, 0x44, 0xa8, 0xd4, 0x9a, 0xa1, 0x2e,
	0x14, 0xaf, 0x9d, 0x88, 0x47, 0x7e, 0x22, 0x54, 0xea, 0xcd, 0x3d, 0x6b, 0xf7, 0x68, 0x0b, 0x8b,
	0x8a, 0x6a, 0xed, 0x64, 0x0b, 0xb3, 0x29, 0x2a, 0x75, 0x69, 0xb1, 0x9b, 0x40, 0xa9, 0xf5, 0x86,
	0xea, 0x4d, 0x27, 0xdb, 0x04, 0x95, 0xfa, 0x89, 0xea, 0x8d, 0xed, 0xb5, 0xa6, 0x6d, 0x25, 0x29,
	0xeb, 0x9d, 0xd1, 0x41, 0xb7, 0xb9, 0x26, 0xaf, 0xb8, 0xde, 0x20, 0x4b, 0x9e, 0xb8, 0xf5, 0xdc,
	0x13, 0xe9, 0x5f, 0x43, 0x18, 0x77, 0x8b, 0xb7, 0x8f, 0xa4, 0xde, 0xbd, 0xe3, 0xda, 0x0e, 0xd9,
	0x68, 0x54, 0x6b, 0x5a, 0xa2, 0x30, 0x9d, 0x7d, 0x23, 0xcb, 0xbc, 0xb6, 0x04, 0x79, 0x05, 0xa7,
	0xba, 0x50, 0x88, 0x62, 0xad, 0x50, 0xcb, 0xaa, 0x71, 0x43, 0x79, 0x9d, 0x62, 0x9f, 0xc8, 0x76,
	0x3c, 0x39, 0x83, 0xa9, 0xe1, 0x5b, 0xfc, 0x28, 0x05, 0x76, 0x72, 0xed, 0xb0, 0xcd, 0x55, 0x54,
	0x94, 0x0d, 0x2d, 0xd1, 0xf5, 0x1e, 0x65, 0x3b, 0x6c, 0x73, 0x0a, 0x6f, 0x50, 0x29, 0x54, 0xf7,
	0xae, 0xf6, 0xd8, 0x99, 0xab, 0xad, 0x31, 0x89, 0xfc, 0x9b, 0xb3, 0x71, 0xfa, 0x03, 0x80, 0x9f,
	0xe6, 0x7d, 0x5b, 0x23, 0x99, 0xc1, 0xa4, 0x11, 0xb7, 0x42, 0xfe, 0x29, 0xe2, 0x47, 0x04, 0x60,
	0xbc, 0x95, 0x39, 0xaf, 0x30, 0x1e, 0xd8, 0xd8, 0xd0, 0xbc, 0x42, 0x13, 0x0f, 0x6d, 0x11, 0x43,
	0x7d, 0x6b, 0x64, 0x1d, 0x07, 0xe9, 0xbf, 0x03, 0x08, 0xaf, 0x94, 0x92, 0xea, 0x10, 0x43, 0xcd,
	0x61, 0xb6, 0xbf, 0x50, 0xff, 0x79, 0xd9, 0xa7, 0xac, 0xaf, 0x18, 0xe6, 0x4d, 0xd9, 0xa9, 0xea,
	0x81, 0x6d, 0xdc, 0xad, 0xd0, 0x8b, 0xe5, 0x62, 0x5b, 0xe9, 0x77, 0x3e, 0x76, 0x3b, 0xf7, 0xa0,
	0xf3, 0x65, 0x71, 0xdb, 0xe9, 0xe2, 0x81, 0xdd, 0x75, 0xa1, 0x90, 0x1a, 0x64, 0x4e, 0x93, 0x20,
	0xeb, 0xa1, 0x3d, 0xb9, 0xb0, 0x9f, 0x21, 0xbf, 0x2c, 0x17, 0xa7, 0x2f, 0x60, 0x76, 0xcd, 0x45,
	0xd9, 0x7f, 0x2d, 0x13, 0x98, 0x6c, 0x51, 0x6b, 0x2b, 0xb6, 0x6f, 0xb8, 0x87, 0xe9, 0x1c, 0x46,
	0xd7, 0x52, 0x94, 0xfb, 0x15, 0xc1, 0x83, 0x8a, 0x97, 0x73, 0x08, 0xaf, 0xec, 0x8b, 0x21, 0x11,
	0x84, 0xef, 0x50, 0xb0, 0x75, 0xfc, 0x88, 0x1c, 0xc1, 0xf4, 0x75, 0xcd, 0xdd, 0x23, 0x8c, 0x87,
	0x17, 0x17, 0x10, 0x7a, 0x77, 0x7c, 0x0d, 0x23, 0x7b, 0x2b, 0x79, 0xdc, 0x2b, 0xb7, 0xd7, 0xc3,
	0xd9, 0xd1, 0x8e, 0x94, 0xa2, 0xcc, 0xc7, 0xee, 0xff, 0xe2, 0xb7, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x2c, 0x8d, 0xcd, 0xec, 0x5c, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	err := c.cc.Invoke(ctx, "/common.Agent/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
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
