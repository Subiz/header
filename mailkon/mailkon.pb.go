// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mailkon/mailkon.proto

/*
Package mailkon is a generated protocol buffer package.

It is generated from these files:
	mailkon/mailkon.proto

It has these top-level messages:
	Address
	Domain
	Message
	User
	SendgridEvent
	UserAvail
	HttpChunk
	Job
*/
package mailkon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	Event_MaikonRequested               Event = 1000
	Event_MailkonSynced                 Event = 1001
	Event_MailkonJob                    Event = 2
	Event_MailkonCreateAccountRequested Event = 0
)

var Event_name = map[int32]string{
	1000: "MaikonRequested",
	1001: "MailkonSynced",
	2:    "MailkonJob",
	0:    "MailkonCreateAccountRequested",
}
var Event_value = map[string]int32{
	"MaikonRequested":               1000,
	"MailkonSynced":                 1001,
	"MailkonJob":                    2,
	"MailkonCreateAccountRequested": 0,
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
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type JobType int32

const (
	JobType_sendgrid_inbound JobType = 0
	JobType_sendgrid_event   JobType = 1
	JobType_subiz_event      JobType = 3
)

var JobType_name = map[int32]string{
	0: "sendgrid_inbound",
	1: "sendgrid_event",
	3: "subiz_event",
}
var JobType_value = map[string]int32{
	"sendgrid_inbound": 0,
	"sendgrid_event":   1,
	"subiz_event":      3,
}

func (x JobType) Enum() *JobType {
	p := new(JobType)
	*p = x
	return p
}
func (x JobType) String() string {
	return proto.EnumName(JobType_name, int32(x))
}
func (x *JobType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(JobType_value, data, "JobType")
	if err != nil {
		return err
	}
	*x = JobType(value)
	return nil
}
func (JobType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Address struct {
	AccountId        *string `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Updated          *int64  `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Address) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Address) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *Address) GetUpdated() int64 {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return 0
}

type Domain struct {
	AccountId        *string `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Domain           *string `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	Updated          *int64  `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
	Dnstype          *string `protobuf:"bytes,4,opt,name=dnstype" json:"dnstype,omitempty"`
	Data             *string `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Domain) Reset()                    { *m = Domain{} }
func (m *Domain) String() string            { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()               {}
func (*Domain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Domain) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Domain) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *Domain) GetUpdated() int64 {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return 0
}

func (m *Domain) GetDnstype() string {
	if m != nil && m.Dnstype != nil {
		return *m.Dnstype
	}
	return ""
}

func (m *Domain) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

type Message struct {
	MessageId        *string  `protobuf:"bytes,2,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	AccountId        *string  `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ConversationId   *string  `protobuf:"bytes,4,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	FromAddr         *string  `protobuf:"bytes,5,opt,name=from_addr,json=fromAddr" json:"from_addr,omitempty"`
	ToAddr           *string  `protobuf:"bytes,6,opt,name=to_addr,json=toAddr" json:"to_addr,omitempty"`
	Subject          *string  `protobuf:"bytes,7,opt,name=subject" json:"subject,omitempty"`
	Body             *string  `protobuf:"bytes,8,opt,name=body" json:"body,omitempty"`
	Header           *string  `protobuf:"bytes,9,opt,name=header" json:"header,omitempty"`
	Attachments      []string `protobuf:"bytes,10,rep,name=attachments" json:"attachments,omitempty"`
	Created          *int64   `protobuf:"varint,11,opt,name=created" json:"created,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Message) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Message) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Message) GetConversationId() string {
	if m != nil && m.ConversationId != nil {
		return *m.ConversationId
	}
	return ""
}

func (m *Message) GetFromAddr() string {
	if m != nil && m.FromAddr != nil {
		return *m.FromAddr
	}
	return ""
}

func (m *Message) GetToAddr() string {
	if m != nil && m.ToAddr != nil {
		return *m.ToAddr
	}
	return ""
}

func (m *Message) GetSubject() string {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *Message) GetHeader() string {
	if m != nil && m.Header != nil {
		return *m.Header
	}
	return ""
}

func (m *Message) GetAttachments() []string {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Message) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

type User struct {
	AccountId        *string `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	EmailAddress     *string `protobuf:"bytes,4,opt,name=email_address,json=emailAddress" json:"email_address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *User) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *User) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *User) GetEmailAddress() string {
	if m != nil && m.EmailAddress != nil {
		return *m.EmailAddress
	}
	return ""
}

type SendgridEvent struct {
	EventId          *string  `protobuf:"bytes,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	AccountId        *string  `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ConversationId   *string  `protobuf:"bytes,6,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	MessageId        *string  `protobuf:"bytes,7,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	Data             []string `protobuf:"bytes,8,rep,name=data" json:"data,omitempty"`
	FullMessageId    *string  `protobuf:"bytes,9,opt,name=full_message_id,json=fullMessageId" json:"full_message_id,omitempty"`
	Subject          *string  `protobuf:"bytes,10,opt,name=subject" json:"subject,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SendgridEvent) Reset()                    { *m = SendgridEvent{} }
func (m *SendgridEvent) String() string            { return proto.CompactTextString(m) }
func (*SendgridEvent) ProtoMessage()               {}
func (*SendgridEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SendgridEvent) GetEventId() string {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return ""
}

func (m *SendgridEvent) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *SendgridEvent) GetConversationId() string {
	if m != nil && m.ConversationId != nil {
		return *m.ConversationId
	}
	return ""
}

func (m *SendgridEvent) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *SendgridEvent) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SendgridEvent) GetFullMessageId() string {
	if m != nil && m.FullMessageId != nil {
		return *m.FullMessageId
	}
	return ""
}

func (m *SendgridEvent) GetSubject() string {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return ""
}

type UserAvail struct {
	AccountId        *string `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	EmailAddress     *string `protobuf:"bytes,5,opt,name=email_address,json=emailAddress" json:"email_address,omitempty"`
	Availability     *bool   `protobuf:"varint,6,opt,name=availability" json:"availability,omitempty"`
	Updated          *int64  `protobuf:"varint,7,opt,name=updated" json:"updated,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserAvail) Reset()                    { *m = UserAvail{} }
func (m *UserAvail) String() string            { return proto.CompactTextString(m) }
func (*UserAvail) ProtoMessage()               {}
func (*UserAvail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserAvail) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *UserAvail) GetEmailAddress() string {
	if m != nil && m.EmailAddress != nil {
		return *m.EmailAddress
	}
	return ""
}

func (m *UserAvail) GetAvailability() bool {
	if m != nil && m.Availability != nil {
		return *m.Availability
	}
	return false
}

func (m *UserAvail) GetUpdated() int64 {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return 0
}

type HttpChunk struct {
	Id               *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	ChunkId          *int32  `protobuf:"varint,4,opt,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	Data             []byte  `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
	ChunkSize        *int32  `protobuf:"varint,6,opt,name=chunk_size,json=chunkSize" json:"chunk_size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HttpChunk) Reset()                    { *m = HttpChunk{} }
func (m *HttpChunk) String() string            { return proto.CompactTextString(m) }
func (*HttpChunk) ProtoMessage()               {}
func (*HttpChunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HttpChunk) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *HttpChunk) GetChunkId() int32 {
	if m != nil && m.ChunkId != nil {
		return *m.ChunkId
	}
	return 0
}

func (m *HttpChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *HttpChunk) GetChunkSize() int32 {
	if m != nil && m.ChunkSize != nil {
		return *m.ChunkSize
	}
	return 0
}

type Job struct {
	Topic            *string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Partition        *int32  `protobuf:"varint,3,opt,name=partition" json:"partition,omitempty"`
	Offset           *int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	ContentType      *string `protobuf:"bytes,5,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Type             *string `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	RequestId        *string `protobuf:"bytes,7,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Job) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *Job) GetPartition() int32 {
	if m != nil && m.Partition != nil {
		return *m.Partition
	}
	return 0
}

func (m *Job) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *Job) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *Job) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Job) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func init() {
	proto.RegisterType((*Address)(nil), "mailkon.Address")
	proto.RegisterType((*Domain)(nil), "mailkon.Domain")
	proto.RegisterType((*Message)(nil), "mailkon.Message")
	proto.RegisterType((*User)(nil), "mailkon.User")
	proto.RegisterType((*SendgridEvent)(nil), "mailkon.SendgridEvent")
	proto.RegisterType((*UserAvail)(nil), "mailkon.UserAvail")
	proto.RegisterType((*HttpChunk)(nil), "mailkon.HttpChunk")
	proto.RegisterType((*Job)(nil), "mailkon.Job")
	proto.RegisterEnum("mailkon.Event", Event_name, Event_value)
	proto.RegisterEnum("mailkon.JobType", JobType_name, JobType_value)
}

func init() { proto.RegisterFile("mailkon/mailkon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x5d, 0x9a, 0xa6, 0x6e, 0xbe, 0xfe, 0xca, 0x1a, 0x2c, 0x08, 0x26, 0x75, 0x41, 0x82, 0x69,
	0x17, 0xf0, 0x0c, 0xd3, 0x86, 0x44, 0x27, 0xf5, 0x26, 0x83, 0x3b, 0xa4, 0xca, 0x89, 0xdd, 0xcd,
	0xac, 0xb5, 0x4b, 0xec, 0x4c, 0xea, 0x5e, 0x80, 0x2b, 0x9e, 0x03, 0xde, 0x0c, 0x78, 0x0b, 0xe4,
	0x9f, 0xae, 0x69, 0x07, 0x1a, 0x57, 0xf5, 0x39, 0x9f, 0xf3, 0xf9, 0xf8, 0x7c, 0xc7, 0x85, 0x27,
	0x0b, 0xc2, 0xe7, 0x37, 0x52, 0xbc, 0xf5, 0xbf, 0x6f, 0x96, 0xa5, 0xd4, 0x12, 0x23, 0x0f, 0xd3,
	0x4f, 0x80, 0x4e, 0x29, 0x2d, 0x99, 0x52, 0xf8, 0x10, 0x80, 0x14, 0x85, 0xac, 0x84, 0x9e, 0x72,
	0x9a, 0x04, 0xa3, 0xe0, 0x38, 0xce, 0x62, 0xcf, 0x8c, 0x29, 0x4e, 0x00, 0x11, 0xb7, 0x33, 0x69,
	0xd8, 0xda, 0x1a, 0x9a, 0x4a, 0xb5, 0xa4, 0x44, 0x33, 0x9a, 0x84, 0xa3, 0xe0, 0x38, 0xcc, 0xd6,
	0x30, 0xfd, 0x1a, 0x40, 0xeb, 0x5c, 0x2e, 0x08, 0x17, 0x8f, 0x75, 0x7f, 0x0a, 0x2d, 0x6a, 0x37,
	0xfa, 0xe6, 0x1e, 0xfd, 0xbb, 0xb7, 0xa9, 0x50, 0xa1, 0xf4, 0x6a, 0xc9, 0x92, 0xa6, 0xd3, 0xe3,
	0x21, 0xc6, 0xd0, 0xa4, 0x44, 0x93, 0x24, 0xb2, 0xb4, 0x5d, 0xa7, 0x3f, 0x1a, 0x80, 0x26, 0x4c,
	0x29, 0x72, 0xc5, 0x8c, 0x94, 0x85, 0x5b, 0x1a, 0x29, 0xee, 0xbc, 0xd8, 0x33, 0x63, 0xba, 0xa3,
	0x34, 0xdc, 0x55, 0xfa, 0x1a, 0x06, 0x85, 0x14, 0xb7, 0xac, 0x54, 0x44, 0x73, 0x29, 0xcc, 0x1e,
	0x77, 0x7e, 0xbf, 0x4e, 0x8f, 0x29, 0x7e, 0x0e, 0xf1, 0xac, 0x94, 0x8b, 0xa9, 0xb1, 0xc9, 0x6b,
	0x69, 0x1b, 0xc2, 0xf8, 0x8d, 0x0f, 0x00, 0x69, 0xe9, 0x4a, 0x2d, 0x77, 0x61, 0x2d, 0x6d, 0x21,
	0x01, 0xa4, 0xaa, 0xfc, 0x33, 0x2b, 0x74, 0x82, 0xdc, 0xb5, 0x3c, 0x34, 0xd7, 0xca, 0x25, 0x5d,
	0x25, 0x6d, 0x77, 0x2d, 0xb3, 0x36, 0xb6, 0x5d, 0x33, 0x42, 0x59, 0x99, 0xc4, 0xae, 0x8b, 0x43,
	0x78, 0x04, 0x1d, 0xa2, 0x35, 0x29, 0xae, 0x17, 0x4c, 0x68, 0x95, 0xc0, 0x28, 0x3c, 0x8e, 0xb3,
	0x3a, 0x65, 0xce, 0x29, 0x4a, 0x66, 0x8d, 0xed, 0x38, 0x63, 0x3d, 0x4c, 0x0b, 0x68, 0x7e, 0x54,
	0xac, 0xdc, 0xf1, 0xa1, 0xb1, 0xeb, 0xc3, 0x01, 0xa0, 0x4a, 0xb1, 0x72, 0xe3, 0x51, 0xcb, 0xc0,
	0x31, 0xc5, 0x2f, 0xa1, 0xc7, 0x4c, 0xbc, 0xa6, 0xeb, 0xb8, 0x38, 0x7b, 0xba, 0x96, 0xf4, 0x61,
	0x4b, 0x7f, 0x06, 0xd0, 0xbb, 0x64, 0x82, 0x5e, 0x95, 0x9c, 0xbe, 0xbb, 0x65, 0x42, 0xe3, 0x67,
	0xd0, 0x66, 0x66, 0xb1, 0x69, 0x88, 0x2c, 0x7e, 0x30, 0x91, 0xe6, 0x7f, 0x4c, 0xa4, 0xf5, 0xd7,
	0x89, 0x6c, 0x0f, 0x1e, 0xed, 0x0e, 0x7e, 0x9d, 0x9b, 0xb6, 0x75, 0xcb, 0xae, 0xf1, 0x2b, 0x18,
	0xcc, 0xaa, 0xf9, 0x7c, 0x5a, 0xfb, 0xce, 0x39, 0xdd, 0x33, 0xf4, 0xe4, 0xfe, 0xdb, 0xda, 0xd8,
	0x60, 0x6b, 0x6c, 0xe9, 0xb7, 0x00, 0x62, 0xe3, 0xe7, 0xe9, 0x2d, 0xe1, 0xf3, 0xc7, 0xc2, 0xf5,
	0xc0, 0xbb, 0xe8, 0xa1, 0x77, 0x38, 0x85, 0x2e, 0x31, 0xcd, 0x48, 0xce, 0xe7, 0x5c, 0xaf, 0xec,
	0x65, 0xdb, 0xd9, 0x16, 0x57, 0x7f, 0x37, 0x68, 0xfb, 0x4d, 0x72, 0x88, 0xdf, 0x6b, 0xbd, 0x3c,
	0xbb, 0xae, 0xc4, 0x0d, 0xee, 0x43, 0xe3, 0x5e, 0x46, 0x83, 0x53, 0x33, 0x84, 0xc2, 0x14, 0xd6,
	0x3e, 0x47, 0x19, 0xb2, 0xb8, 0xe6, 0x8e, 0x51, 0xd4, 0xf5, 0xee, 0x1c, 0x02, 0xb8, 0xed, 0x8a,
	0xdf, 0x31, 0xab, 0x23, 0xca, 0x62, 0xcb, 0x5c, 0xf2, 0x3b, 0x96, 0x7e, 0x0f, 0x20, 0xbc, 0x90,
	0x39, 0xde, 0x87, 0x48, 0xcb, 0x25, 0x2f, 0x7c, 0x88, 0x1c, 0xc0, 0x2f, 0x20, 0x5e, 0x92, 0x52,
	0x73, 0x33, 0x1c, 0x2b, 0x21, 0xca, 0x36, 0x84, 0x49, 0xb6, 0x9c, 0xcd, 0x14, 0xd3, 0x56, 0x47,
	0x98, 0x79, 0x84, 0x8f, 0xa0, 0x5b, 0x48, 0xa1, 0x4d, 0x50, 0xec, 0xdb, 0x77, 0x06, 0x75, 0x3c,
	0xf7, 0xc1, 0xbf, 0x7f, 0x5b, 0x72, 0x21, 0xb0, 0x6b, 0xa3, 0xb4, 0x64, 0x5f, 0x2a, 0xa6, 0x74,
	0x6d, 0xf4, 0x9e, 0x19, 0xd3, 0x13, 0x0a, 0x91, 0x4b, 0xe1, 0x3e, 0x0c, 0x26, 0x84, 0xdf, 0x48,
	0x91, 0xb9, 0x1a, 0xa3, 0xc3, 0x5f, 0x08, 0x63, 0xe8, 0x4d, 0xdc, 0x1f, 0xe6, 0xe5, 0x4a, 0x14,
	0x8c, 0x0e, 0x7f, 0x23, 0xdc, 0x07, 0xf0, 0xdc, 0x85, 0xcc, 0x87, 0x0d, 0x7c, 0x04, 0x87, 0x1e,
	0x9f, 0xd9, 0x87, 0x74, 0xea, 0x86, 0xba, 0xe9, 0xb3, 0x77, 0x72, 0x0e, 0xe8, 0x42, 0xe6, 0x56,
	0xe3, 0x3e, 0x0c, 0x95, 0x8f, 0xff, 0x94, 0x8b, 0x5c, 0x56, 0x82, 0x0e, 0xf7, 0x30, 0x86, 0xfe,
	0x3d, 0x6b, 0xc3, 0x3f, 0x0c, 0xf0, 0x00, 0x3a, 0xaa, 0xca, 0xf9, 0x9d, 0x27, 0xc2, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x74, 0x2a, 0xa7, 0x94, 0xd3, 0x05, 0x00, 0x00,
}
