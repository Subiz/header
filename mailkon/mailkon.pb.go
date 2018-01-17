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
	Event_MailkonCreateAccountRequested Event = 0
)

var Event_name = map[int32]string{
	1000: "MaikonRequested",
	1001: "MailkonSynced",
	0:    "MailkonCreateAccountRequested",
}
var Event_value = map[string]int32{
	"MaikonRequested":               1000,
	"MailkonSynced":                 1001,
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

func init() {
	proto.RegisterType((*Address)(nil), "mailkon.Address")
	proto.RegisterType((*Domain)(nil), "mailkon.Domain")
	proto.RegisterType((*Message)(nil), "mailkon.Message")
	proto.RegisterType((*User)(nil), "mailkon.User")
	proto.RegisterType((*SendgridEvent)(nil), "mailkon.SendgridEvent")
	proto.RegisterType((*UserAvail)(nil), "mailkon.UserAvail")
	proto.RegisterEnum("mailkon.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("mailkon/mailkon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x66, 0x9b, 0x64, 0xbd, 0x3b, 0x6d, 0x68, 0x65, 0x01, 0x35, 0x42, 0x95, 0xc2, 0x72, 0x20,
	0xe2, 0x00, 0xcf, 0x10, 0x01, 0x87, 0x1c, 0x72, 0xd9, 0xc2, 0x0d, 0x29, 0x72, 0xd6, 0x43, 0xbb,
	0x34, 0x6b, 0x07, 0xdb, 0x1b, 0x29, 0x4f, 0xc0, 0x89, 0xc7, 0x40, 0xe2, 0xd1, 0xe0, 0x2d, 0x90,
	0x7f, 0xa2, 0x24, 0x1b, 0x50, 0x39, 0xc5, 0xdf, 0x37, 0xce, 0xf8, 0x9b, 0x6f, 0xbe, 0x85, 0xc7,
	0x0d, 0xaf, 0x97, 0x77, 0x4a, 0xbe, 0x89, 0xbf, 0xaf, 0x57, 0x5a, 0x59, 0x45, 0x49, 0x84, 0xc5,
	0x27, 0x20, 0x13, 0x21, 0x34, 0x1a, 0x43, 0xaf, 0x00, 0x78, 0x55, 0xa9, 0x56, 0xda, 0x79, 0x2d,
	0x58, 0x32, 0x4a, 0xc6, 0x79, 0x99, 0x47, 0x66, 0x2a, 0x28, 0x03, 0xc2, 0xc3, 0x4d, 0x76, 0xe2,
	0x6b, 0x5b, 0xe8, 0x2a, 0xed, 0x4a, 0x70, 0x8b, 0x82, 0xf5, 0x46, 0xc9, 0xb8, 0x57, 0x6e, 0x61,
	0xf1, 0x2d, 0x81, 0xf4, 0x9d, 0x6a, 0x78, 0x2d, 0xef, 0xeb, 0xfe, 0x04, 0x52, 0xe1, 0x2f, 0xc6,
	0xe6, 0x11, 0xfd, 0xbb, 0xb7, 0xab, 0x08, 0x69, 0xec, 0x66, 0x85, 0xac, 0x1f, 0xf4, 0x44, 0x48,
	0x29, 0xf4, 0x05, 0xb7, 0x9c, 0x0d, 0x3c, 0xed, 0xcf, 0xc5, 0xcf, 0x13, 0x20, 0x33, 0x34, 0x86,
	0xdf, 0xa0, 0x93, 0xd2, 0x84, 0xa3, 0x93, 0x12, 0xde, 0xcb, 0x23, 0x33, 0x15, 0x1d, 0xa5, 0xbd,
	0xae, 0xd2, 0x97, 0x70, 0x5e, 0x29, 0xb9, 0x46, 0x6d, 0xb8, 0xad, 0x95, 0x74, 0x77, 0xc2, 0xfb,
	0x0f, 0xf7, 0xe9, 0xa9, 0xa0, 0xcf, 0x20, 0xff, 0xac, 0x55, 0x33, 0x77, 0x36, 0x45, 0x2d, 0x99,
	0x23, 0x9c, 0xdf, 0xf4, 0x12, 0x88, 0x55, 0xa1, 0x94, 0x86, 0x81, 0xad, 0xf2, 0x05, 0x06, 0xc4,
	0xb4, 0x8b, 0x2f, 0x58, 0x59, 0x46, 0xc2, 0x58, 0x11, 0xba, 0xb1, 0x16, 0x4a, 0x6c, 0x58, 0x16,
	0xc6, 0x72, 0x67, 0x67, 0xdb, 0x2d, 0x72, 0x81, 0x9a, 0xe5, 0xa1, 0x4b, 0x40, 0x74, 0x04, 0xa7,
	0xdc, 0x5a, 0x5e, 0xdd, 0x36, 0x28, 0xad, 0x61, 0x30, 0xea, 0x8d, 0xf3, 0x72, 0x9f, 0x72, 0xef,
	0x54, 0x1a, 0xbd, 0xb1, 0xa7, 0xc1, 0xd8, 0x08, 0x8b, 0x0a, 0xfa, 0x1f, 0x0d, 0xea, 0x8e, 0x0f,
	0x27, 0x5d, 0x1f, 0x2e, 0x81, 0xb4, 0x06, 0xf5, 0xce, 0xa3, 0xd4, 0xc1, 0xa9, 0xa0, 0x2f, 0x60,
	0x88, 0x2e, 0x5e, 0xf3, 0x6d, 0x5c, 0x82, 0x3d, 0x67, 0x9e, 0x8c, 0x61, 0x2b, 0x7e, 0x24, 0x30,
	0xbc, 0x46, 0x29, 0x6e, 0x74, 0x2d, 0xde, 0xaf, 0x51, 0x5a, 0xfa, 0x14, 0x32, 0x74, 0x87, 0x5d,
	0x43, 0xe2, 0xf1, 0xd1, 0x46, 0xfa, 0xff, 0xb1, 0x91, 0xf4, 0xaf, 0x1b, 0x39, 0x5c, 0x3c, 0xe9,
	0x2e, 0x7e, 0x9b, 0x9b, 0xcc, 0xbb, 0x15, 0x72, 0xf3, 0x3d, 0x81, 0xdc, 0xb9, 0x31, 0x59, 0xf3,
	0x7a, 0x79, 0x5f, 0x34, 0x8e, 0x26, 0x1f, 0x1c, 0x4f, 0x4e, 0x0b, 0x38, 0xe3, 0xae, 0x19, 0x5f,
	0xd4, 0xcb, 0xda, 0x6e, 0xbc, 0xd4, 0xac, 0x3c, 0xe0, 0xf6, 0x53, 0x4f, 0x0e, 0x52, 0xff, 0xea,
	0x03, 0x0c, 0x82, 0x5d, 0x8f, 0xe0, 0x7c, 0xc6, 0xeb, 0x3b, 0x25, 0x4b, 0xfc, 0xda, 0xa2, 0xb1,
	0x28, 0x2e, 0x7e, 0x11, 0x4a, 0x61, 0x38, 0x0b, 0x5f, 0xf6, 0xf5, 0x46, 0x56, 0x28, 0x2e, 0x7e,
	0x13, 0xfa, 0x1c, 0xae, 0x22, 0xf7, 0xd6, 0x6f, 0x78, 0x12, 0xf4, 0xee, 0xfe, 0xf7, 0xe0, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x0b, 0xa3, 0x92, 0x26, 0x04, 0x00, 0x00,
}
