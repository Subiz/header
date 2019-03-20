// Code generated by protoc-gen-go. DO NOT EDIT.
// source: callback/callback.proto

// schedule help publish a kafka event in future

package callback

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Task struct {
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// callback_sec is second*, golang: time.Now().Unix(),
	// node: new Date() * 1000000
	CallbackSec int64  `protobuf:"varint,3,opt,name=callback_sec,json=callbackSec,proto3" json:"callback_sec,omitempty"`
	Topic       string `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	// Data is value of kafka event in bytes
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Key                  string   `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Called               int64    `protobuf:"varint,8,opt,name=called,proto3" json:"called,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0902222c118453, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetCallbackSec() int64 {
	if m != nil {
		return m.CallbackSec
	}
	return 0
}

func (m *Task) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Task) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Task) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Task) GetCalled() int64 {
	if m != nil {
		return m.Called
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "callback.Task")
}

func init() { proto.RegisterFile("callback/callback.proto", fileDescriptor_6a0902222c118453) }

var fileDescriptor_6a0902222c118453 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4e, 0xcc, 0xc9,
	0x49, 0x4a, 0x4c, 0xce, 0xd6, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0xa5, 0x5e, 0x46, 0x2e, 0x96, 0x90, 0xc4, 0xe2, 0x6c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x45, 0x2e, 0x1e, 0x98, 0xa2,
	0xf8, 0xe2, 0xd4, 0x64, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x6e, 0x98, 0x58, 0x70, 0x6a,
	0xb2, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x7e, 0x41, 0x66, 0xb2, 0x04, 0x0b, 0x58, 0x17, 0x84, 0x23,
	0x24, 0xc4, 0xc5, 0x92, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x66,
	0x0b, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0xb0, 0x81, 0xd5, 0x81, 0x98, 0x42, 0x62, 0x5c,
	0x6c, 0x20, 0xa3, 0x52, 0x53, 0x24, 0x38, 0xc0, 0x06, 0x43, 0x79, 0x49, 0x6c, 0x60, 0x07, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x19, 0xe8, 0x9a, 0x52, 0xbb, 0x00, 0x00, 0x00,
}