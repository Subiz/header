// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email/email.proto

/*
Package email is a generated protocol buffer package.

It is generated from these files:
	email/email.proto

It has these top-level messages:
	Email
	Attachment
*/
package email

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "git.subiz.net/header/common"

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
	Event_Email_SendRequest Event = 0
)

var Event_name = map[int32]string{
	0: "Email_SendRequest",
}
var Event_value = map[string]int32{
	"Email_SendRequest": 0,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Email struct {
	Ctx  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	From string          `protobuf:"bytes,3,opt,name=from" json:"from,omitempty"`
	// string to = 4;
	Subject     string            `protobuf:"bytes,5,opt,name=subject" json:"subject,omitempty"`
	Body        string            `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
	Header      map[string]string `protobuf:"bytes,7,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Attachments []*Attachment     `protobuf:"bytes,8,rep,name=attachments" json:"attachments,omitempty"`
	To          []string          `protobuf:"bytes,9,rep,name=to" json:"to,omitempty"`
	Cc          []string          `protobuf:"bytes,10,rep,name=cc" json:"cc,omitempty"`
	Bcc         []string          `protobuf:"bytes,11,rep,name=bcc" json:"bcc,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Email) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Email) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Email) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Email) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Email) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Email) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Email) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Email) GetCc() []string {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *Email) GetBcc() []string {
	if m != nil {
		return m.Bcc
	}
	return nil
}

type Attachment struct {
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Mimetype string `protobuf:"bytes,4,opt,name=mimetype" json:"mimetype,omitempty"`
}

func (m *Attachment) Reset()                    { *m = Attachment{} }
func (m *Attachment) String() string            { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()               {}
func (*Attachment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Attachment) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Attachment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Attachment) GetMimetype() string {
	if m != nil {
		return m.Mimetype
	}
	return ""
}

func init() {
	proto.RegisterType((*Email)(nil), "email.Email")
	proto.RegisterType((*Attachment)(nil), "email.Attachment")
	proto.RegisterEnum("email.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("email/email.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x14, 0xfc, 0xda, 0x52, 0xa0, 0xaf, 0xc9, 0x27, 0x6c, 0x34, 0xd9, 0x70, 0x30, 0x95, 0x53, 0xe3,
	0xa1, 0x18, 0xb8, 0xa8, 0x37, 0x63, 0x48, 0x3c, 0x79, 0xa8, 0x3f, 0xc0, 0xb4, 0xcb, 0x53, 0xaa,
	0xec, 0x2e, 0xb6, 0xaf, 0x84, 0xfa, 0xeb, 0xfc, 0x69, 0x66, 0xb7, 0x05, 0xb9, 0xb4, 0x33, 0xef,
	0xcd, 0xec, 0x4e, 0x66, 0x61, 0x8c, 0x32, 0x2b, 0x36, 0x33, 0xfb, 0x4d, 0xb6, 0xa5, 0x26, 0xcd,
	0x7c, 0x4b, 0x26, 0xf1, 0x7b, 0x41, 0x49, 0x55, 0xe7, 0xc5, 0x77, 0xa2, 0x90, 0x66, 0x6b, 0xcc,
	0x56, 0x58, 0xce, 0x84, 0x96, 0x52, 0xab, 0xee, 0xd7, 0x1a, 0xa6, 0x3f, 0x2e, 0xf8, 0x4b, 0xe3,
	0x61, 0x57, 0xe0, 0x09, 0xda, 0x73, 0x27, 0x72, 0xe2, 0x70, 0x7e, 0x96, 0x74, 0xaa, 0x47, 0xad,
	0x08, 0xf7, 0x94, 0x9a, 0x1d, 0x63, 0xd0, 0x7b, 0x2b, 0xb5, 0xe4, 0x5e, 0xe4, 0xc4, 0x41, 0x6a,
	0x31, 0xe3, 0x30, 0xa8, 0xea, 0xfc, 0x03, 0x05, 0x71, 0xdf, 0x8e, 0x0f, 0xd4, 0xa8, 0x73, 0xbd,
	0x6a, 0x78, 0xbf, 0x55, 0x1b, 0xcc, 0x6e, 0xa0, 0xdf, 0x86, 0xe1, 0x83, 0xc8, 0x8b, 0xc3, 0x39,
	0x4f, 0xda, 0xf4, 0x36, 0x42, 0xf2, 0x64, 0x57, 0x4b, 0x45, 0x65, 0x93, 0x76, 0x3a, 0xb6, 0x80,
	0x30, 0x23, 0xca, 0xc4, 0x5a, 0xa2, 0xa2, 0x8a, 0x0f, 0xad, 0x6d, 0xdc, 0xd9, 0x1e, 0x8e, 0x9b,
	0xf4, 0x54, 0xc5, 0xfe, 0x83, 0x4b, 0x9a, 0x07, 0x91, 0x17, 0x07, 0xa9, 0x4b, 0xda, 0x70, 0x21,
	0x38, 0xb4, 0x5c, 0x08, 0x36, 0x02, 0x2f, 0x17, 0x82, 0x87, 0x76, 0x60, 0xe0, 0xe4, 0x0e, 0xc2,
	0x93, 0xdb, 0x8d, 0xe0, 0x13, 0x1b, 0x5b, 0x46, 0x90, 0x1a, 0xc8, 0xce, 0xc1, 0xdf, 0x65, 0x9b,
	0x1a, 0xb9, 0x6b, 0x67, 0x2d, 0xb9, 0x77, 0x6f, 0x9d, 0xe9, 0x33, 0xc0, 0x5f, 0x0e, 0xe3, 0xac,
	0xcb, 0x4d, 0xa7, 0x32, 0xd0, 0xf4, 0xa0, 0x32, 0x89, 0x87, 0xd6, 0x0c, 0x66, 0x13, 0x18, 0xca,
	0x42, 0x22, 0x35, 0x5b, 0xe4, 0x3d, 0x3b, 0x3f, 0xf2, 0xeb, 0x4b, 0xf0, 0x97, 0x3b, 0x73, 0xd4,
	0x05, 0x8c, 0x6d, 0x2f, 0xaf, 0x2f, 0xa8, 0x56, 0x29, 0x7e, 0xd5, 0x58, 0xd1, 0xe8, 0x5f, 0xde,
	0xb7, 0x2f, 0xb7, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x74, 0x47, 0x35, 0xff, 0x01, 0x00,
	0x00,
}
