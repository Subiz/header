// Code generated by protoc-gen-go. DO NOT EDIT.
// source: partitioner/partitioner.proto

package partitioner

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

type ConfigurationResponse_Type int32

const (
	// coordinator is in middle of configuration change process
	ConfigurationResponse_PENDING ConfigurationResponse_Type = 0
	ConfigurationResponse_NORMAL  ConfigurationResponse_Type = 1
)

var ConfigurationResponse_Type_name = map[int32]string{
	0: "PENDING",
	1: "NORMAL",
}

var ConfigurationResponse_Type_value = map[string]int32{
	"PENDING": 0,
	"NORMAL":  1,
}

func (x ConfigurationResponse_Type) String() string {
	return proto.EnumName(ConfigurationResponse_Type_name, int32(x))
}

func (ConfigurationResponse_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{1, 0}
}

type Event_Type int32

const (
	Event_ABORT Event_Type = 0
	Event_APPLY Event_Type = 1
	Event_BLOCK Event_Type = 2
)

var Event_Type_name = map[int32]string{
	0: "ABORT",
	1: "APPLY",
	2: "BLOCK",
}

var Event_Type_value = map[string]int32{
	"ABORT": 0,
	"APPLY": 1,
	"BLOCK": 2,
}

func (x Event_Type) String() string {
	return proto.EnumName(Event_Type_name, int32(x))
}

func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{3, 0}
}

// A string
type Configuration struct {
	Version              string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term                 int32                  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Workers              []*WorkerConfiguration `protobuf:"bytes,3,rep,name=workers,proto3" json:"workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Configuration) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Configuration) GetWorkers() []*WorkerConfiguration {
	if m != nil {
		return m.Workers
	}
	return nil
}

type ConfigurationResponse struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term    int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	// indicates current state of coordinator
	State                string                 `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Workers              []*WorkerConfiguration `protobuf:"bytes,3,rep,name=workers,proto3" json:"workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConfigurationResponse) Reset()         { *m = ConfigurationResponse{} }
func (m *ConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigurationResponse) ProtoMessage()    {}
func (*ConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{1}
}

func (m *ConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigurationResponse.Unmarshal(m, b)
}
func (m *ConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigurationResponse.Marshal(b, m, deterministic)
}
func (m *ConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationResponse.Merge(m, src)
}
func (m *ConfigurationResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigurationResponse.Size(m)
}
func (m *ConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationResponse proto.InternalMessageInfo

func (m *ConfigurationResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ConfigurationResponse) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *ConfigurationResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ConfigurationResponse) GetWorkers() []*WorkerConfiguration {
	if m != nil {
		return m.Workers
	}
	return nil
}

type WorkerConfiguration struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term    int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	// unique string for each workers, ex: web-1, web-2
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// list of job
	Partitions           []int32  `protobuf:"varint,4,rep,packed,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerConfiguration) Reset()         { *m = WorkerConfiguration{} }
func (m *WorkerConfiguration) String() string { return proto.CompactTextString(m) }
func (*WorkerConfiguration) ProtoMessage()    {}
func (*WorkerConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{2}
}

func (m *WorkerConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerConfiguration.Unmarshal(m, b)
}
func (m *WorkerConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerConfiguration.Marshal(b, m, deterministic)
}
func (m *WorkerConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerConfiguration.Merge(m, src)
}
func (m *WorkerConfiguration) XXX_Size() int {
	return xxx_messageInfo_WorkerConfiguration.Size(m)
}
func (m *WorkerConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerConfiguration proto.InternalMessageInfo

func (m *WorkerConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *WorkerConfiguration) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *WorkerConfiguration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerConfiguration) GetPartitions() []int32 {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Event used by coordinator to notify workers about configuration change process
type Event struct {
	Version              string         `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term                 int32          `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Type                 string         `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Configuration        *Configuration `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Created              int64          `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{3}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Event) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetConfiguration() *Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Event) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func init() {
	proto.RegisterEnum("partitioner.ConfigurationResponse_Type", ConfigurationResponse_Type_name, ConfigurationResponse_Type_value)
	proto.RegisterEnum("partitioner.Event_Type", Event_Type_name, Event_Type_value)
	proto.RegisterType((*Configuration)(nil), "partitioner.Configuration")
	proto.RegisterType((*ConfigurationResponse)(nil), "partitioner.ConfigurationResponse")
	proto.RegisterType((*WorkerConfiguration)(nil), "partitioner.WorkerConfiguration")
	proto.RegisterType((*Event)(nil), "partitioner.Event")
}

func init() { proto.RegisterFile("partitioner/partitioner.proto", fileDescriptor_e94f51d1df456051) }

var fileDescriptor_e94f51d1df456051 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4b, 0x03, 0x31,
	0x14, 0x6c, 0xf6, 0xa3, 0xa5, 0x6f, 0xa9, 0x2c, 0x51, 0x21, 0x08, 0x6a, 0xd8, 0x8b, 0x7b, 0xaa,
	0x50, 0x6f, 0x9e, 0x6c, 0x6b, 0x11, 0xb1, 0xb6, 0x25, 0x14, 0xc4, 0xe3, 0xda, 0x3e, 0x25, 0x88,
	0x9b, 0x25, 0x89, 0x95, 0xfe, 0x33, 0x7f, 0x85, 0xbf, 0x49, 0x36, 0xb6, 0xb2, 0x0b, 0x5e, 0x8a,
	0xb7, 0x99, 0x79, 0xef, 0x65, 0x86, 0x21, 0x70, 0x5c, 0x64, 0xda, 0x4a, 0x2b, 0x55, 0x8e, 0xfa,
	0xbc, 0x82, 0xbb, 0x85, 0x56, 0x56, 0xd1, 0xa8, 0x22, 0x25, 0x6b, 0xe8, 0x0c, 0x55, 0xfe, 0x2c,
	0x5f, 0xde, 0x75, 0x56, 0x4a, 0x94, 0x41, 0x6b, 0x85, 0xda, 0x48, 0x95, 0x33, 0xc2, 0x49, 0xda,
	0x16, 0x5b, 0x4a, 0x29, 0x04, 0x16, 0xf5, 0x1b, 0xf3, 0x38, 0x49, 0x43, 0xe1, 0x30, 0xbd, 0x84,
	0xd6, 0x87, 0xd2, 0xaf, 0xa8, 0x0d, 0xf3, 0xb9, 0x9f, 0x46, 0x3d, 0xde, 0xad, 0x1a, 0x3e, 0xb8,
	0x59, 0xcd, 0x40, 0x6c, 0x0f, 0x92, 0x4f, 0x02, 0x87, 0xf5, 0x11, 0x9a, 0x42, 0xe5, 0x06, 0x77,
	0xcc, 0x70, 0x00, 0xa1, 0xb1, 0x99, 0x45, 0x16, 0xb8, 0xdd, 0x1f, 0xf2, 0xaf, 0x64, 0xa7, 0x10,
	0xcc, 0xd7, 0x05, 0xd2, 0x08, 0x5a, 0xb3, 0xd1, 0xe4, 0xfa, 0x76, 0x72, 0x13, 0x37, 0x28, 0x40,
	0x73, 0x32, 0x15, 0xf7, 0xfd, 0x71, 0x4c, 0x12, 0x03, 0xfb, 0x7f, 0x3c, 0xb0, 0x63, 0xee, 0x3d,
	0xf0, 0xe4, 0x92, 0xf9, 0x6e, 0xd1, 0x93, 0x4b, 0x7a, 0x02, 0xf0, 0x9b, 0xd0, 0xb0, 0x80, 0xfb,
	0x69, 0x28, 0x2a, 0x4a, 0xf2, 0x45, 0x20, 0x1c, 0xad, 0x30, 0xb7, 0x3b, 0xfa, 0x94, 0xda, 0xba,
	0xc0, 0x8d, 0x93, 0xc3, 0xf4, 0x0a, 0x3a, 0x8b, 0x6a, 0x74, 0xd7, 0x5d, 0xd4, 0x3b, 0xaa, 0x75,
	0x54, 0x6f, 0xa7, 0x7e, 0x50, 0x66, 0x58, 0x68, 0xcc, 0x2c, 0x2e, 0x59, 0xc8, 0x49, 0xea, 0x8b,
	0x2d, 0x4d, 0xce, 0x36, 0xed, 0xb5, 0x21, 0xec, 0x0f, 0xa6, 0x62, 0x1e, 0x37, 0x1c, 0x9c, 0xcd,
	0xc6, 0x8f, 0x31, 0x29, 0xe1, 0x60, 0x3c, 0x1d, 0xde, 0xc5, 0xde, 0x53, 0xd3, 0xfd, 0xc7, 0x8b,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xfb, 0xc3, 0x69, 0xb0, 0x02, 0x00, 0x00,
}
