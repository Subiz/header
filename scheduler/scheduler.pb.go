// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler/scheduler.proto

/*
Package scheduler is a generated protocol buffer package.

schedule help publish a kafka event in future

It is generated from these files:
	scheduler/scheduler.proto

It has these top-level messages:
	Task
	Id
*/
package scheduler

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
	Event_SchedulerRequested Event = 0
	Event_SchedulerCancelled Event = 1
)

var Event_name = map[int32]string{
	0: "SchedulerRequested",
	1: "SchedulerCancelled",
}
var Event_value = map[string]int32{
	"SchedulerRequested": 0,
	"SchedulerCancelled": 1,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Task struct {
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// string ScheduleTime = 2;
	// callback_time is *nano second*, golang: time.Now().Unix(), node: new Dat() * 1000000
	CallbackTime int64  `protobuf:"varint,3,opt,name=callback_time,json=callbackTime" json:"callback_time,omitempty"`
	Topic        string `protobuf:"bytes,4,opt,name=topic" json:"topic,omitempty"`
	// Data is value of kafka event in bytes
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Key  string `protobuf:"bytes,6,opt,name=key" json:"key,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetCallbackTime() int64 {
	if m != nil {
		return m.CallbackTime
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

type Id struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CallbackTime int64  `protobuf:"varint,2,opt,name=callback_time,json=callbackTime" json:"callback_time,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Id) GetCallbackTime() int64 {
	if m != nil {
		return m.CallbackTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "scheduler.Task")
	proto.RegisterType((*Id)(nil), "scheduler.Id")
	proto.RegisterEnum("scheduler.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("scheduler/scheduler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xbb, 0x4b, 0xc6, 0x30,
	0x14, 0x47, 0x4d, 0xfa, 0x80, 0x5e, 0xaa, 0x94, 0x8b, 0x48, 0xdc, 0x4a, 0x5d, 0x8a, 0x83, 0x0e,
	0x0e, 0xe2, 0x2c, 0x0e, 0xae, 0xb1, 0xbb, 0xa4, 0xc9, 0x05, 0x43, 0xd3, 0x87, 0x6d, 0xaa, 0xf8,
	0xdf, 0x8b, 0xd1, 0x16, 0xe4, 0xe3, 0xdb, 0x4e, 0x0e, 0xb9, 0x1c, 0x7e, 0x70, 0xb9, 0xe8, 0x37,
	0x32, 0xab, 0xa3, 0xf9, 0x76, 0xa7, 0x9b, 0x69, 0x1e, 0xfd, 0x88, 0xd9, 0x2e, 0xaa, 0x4f, 0x88,
	0x1b, 0xb5, 0x74, 0x78, 0x06, 0xdc, 0x1a, 0xc1, 0x4b, 0x56, 0x67, 0x92, 0x5b, 0x83, 0x57, 0x70,
	0xaa, 0x95, 0x73, 0xad, 0xd2, 0xdd, 0xab, 0xb7, 0x3d, 0x89, 0xa8, 0x64, 0x75, 0x24, 0xf3, 0x4d,
	0x36, 0xb6, 0x27, 0x3c, 0x87, 0xc4, 0x8f, 0x93, 0xd5, 0x22, 0x0e, 0x77, 0xbf, 0x0f, 0x44, 0x88,
	0x8d, 0xf2, 0x4a, 0x24, 0x25, 0xab, 0x73, 0x19, 0x18, 0x0b, 0x88, 0x3a, 0xfa, 0x12, 0x69, 0xf8,
	0xf7, 0x83, 0xd5, 0x03, 0xf0, 0x67, 0xf3, 0x97, 0x65, 0xc7, 0xb3, 0xfc, 0x30, 0x7b, 0x7d, 0x0f,
	0xc9, 0xd3, 0x07, 0x0d, 0x1e, 0x2f, 0x00, 0x5f, 0xb6, 0x25, 0x92, 0xde, 0x57, 0x5a, 0x3c, 0x99,
	0xe2, 0xe4, 0x9f, 0x7f, 0x54, 0x83, 0x26, 0xe7, 0xc8, 0x14, 0xac, 0x4d, 0xc3, 0xfc, 0xbb, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x55, 0x20, 0x43, 0x1b, 0x01, 0x00, 0x00,
}