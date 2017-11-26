// Code generated by protoc-gen-go.
// source: detector/detector.proto
// DO NOT EDIT!

/*
Package detector is a generated protocol buffer package.

It is generated from these files:
	detector/detector.proto

It has these top-level messages:
	UserAgent
	String
	Zipped
	Location
*/
package detector

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

type UserAgent struct {
	Id                   int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	BrowserVersionFull   string `protobuf:"bytes,2,opt,name=BrowserVersionFull" json:"BrowserVersionFull,omitempty"`
	LayoutEngine         string `protobuf:"bytes,3,opt,name=LayoutEngine" json:"LayoutEngine,omitempty"`
	LayoutEngineVersion  string `protobuf:"bytes,4,opt,name=LayoutEngineVersion" json:"LayoutEngineVersion,omitempty"`
	HardwareArchitecture string `protobuf:"bytes,5,opt,name=HardwareArchitecture" json:"HardwareArchitecture,omitempty"`
	Os                   string `protobuf:"bytes,6,opt,name=Os" json:"Os,omitempty"`
	BrowserVersion       string `protobuf:"bytes,7,opt,name=BrowserVersion" json:"BrowserVersion,omitempty"`
	DeviceType           string `protobuf:"bytes,8,opt,name=DeviceType" json:"DeviceType,omitempty"`
	UserAgent            string `protobuf:"bytes,9,opt,name=UserAgent" json:"UserAgent,omitempty"`
}

func (m *UserAgent) Reset()                    { *m = UserAgent{} }
func (m *UserAgent) String() string            { return proto.CompactTextString(m) }
func (*UserAgent) ProtoMessage()               {}
func (*UserAgent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserAgent) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserAgent) GetBrowserVersionFull() string {
	if m != nil {
		return m.BrowserVersionFull
	}
	return ""
}

func (m *UserAgent) GetLayoutEngine() string {
	if m != nil {
		return m.LayoutEngine
	}
	return ""
}

func (m *UserAgent) GetLayoutEngineVersion() string {
	if m != nil {
		return m.LayoutEngineVersion
	}
	return ""
}

func (m *UserAgent) GetHardwareArchitecture() string {
	if m != nil {
		return m.HardwareArchitecture
	}
	return ""
}

func (m *UserAgent) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *UserAgent) GetBrowserVersion() string {
	if m != nil {
		return m.BrowserVersion
	}
	return ""
}

func (m *UserAgent) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *UserAgent) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

type String struct {
	String_ string `protobuf:"bytes,1,opt,name=String" json:"String,omitempty"`
}

func (m *String) Reset()                    { *m = String{} }
func (m *String) String() string            { return proto.CompactTextString(m) }
func (*String) ProtoMessage()               {}
func (*String) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *String) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

type Zipped struct {
	Id int32 `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
}

func (m *Zipped) Reset()                    { *m = Zipped{} }
func (m *Zipped) String() string            { return proto.CompactTextString(m) }
func (*Zipped) ProtoMessage()               {}
func (*Zipped) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Zipped) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Location struct {
	Id       int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Country  string `protobuf:"bytes,2,opt,name=Country" json:"Country,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=City" json:"City,omitempty"`
	TimeZone string `protobuf:"bytes,4,opt,name=TimeZone" json:"TimeZone,omitempty"`
	Language string `protobuf:"bytes,5,opt,name=Language" json:"Language,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Location) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Location) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Location) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Location) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *Location) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func init() {
	proto.RegisterType((*UserAgent)(nil), "UserAgent")
	proto.RegisterType((*String)(nil), "String")
	proto.RegisterType((*Zipped)(nil), "Zipped")
	proto.RegisterType((*Location)(nil), "Location")
}

func init() { proto.RegisterFile("detector/detector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x49, 0x6c, 0xd3, 0x66, 0x90, 0x1e, 0x46, 0xd1, 0x45, 0x44, 0x4a, 0x0e, 0xd2, 0x53,
	0x15, 0x7d, 0x82, 0x5a, 0x15, 0x0b, 0x85, 0x42, 0xad, 0x1e, 0x7a, 0x5b, 0x93, 0x21, 0x2e, 0xd4,
	0xdd, 0x30, 0xd9, 0x58, 0x72, 0xf3, 0xfd, 0x7c, 0x29, 0xc9, 0x92, 0xd4, 0xb6, 0xf4, 0x36, 0xff,
	0x8f, 0x85, 0x99, 0x5f, 0x02, 0xe7, 0x09, 0x59, 0x8a, 0xad, 0xe1, 0x9b, 0x66, 0x18, 0x66, 0x6c,
	0xac, 0x89, 0x7e, 0x7d, 0x08, 0xdf, 0x72, 0xe2, 0x51, 0x4a, 0xda, 0x62, 0x0f, 0xfc, 0x49, 0x22,
	0xbc, 0xbe, 0x37, 0x68, 0xcf, 0xfd, 0x49, 0x82, 0x43, 0xc0, 0x07, 0x36, 0xeb, 0x9c, 0xf8, 0x9d,
	0x38, 0x57, 0x46, 0x3f, 0x17, 0xab, 0x95, 0xf0, 0xfb, 0xde, 0x20, 0x9c, 0x1f, 0x48, 0x30, 0x82,
	0xe3, 0xa9, 0x2c, 0x4d, 0x61, 0x9f, 0x74, 0xaa, 0x34, 0x89, 0x23, 0xd7, 0xdc, 0xf1, 0xf0, 0x16,
	0x4e, 0xb6, 0x75, 0xfd, 0x5c, 0xb4, 0x5c, 0xf5, 0x50, 0x84, 0x77, 0x70, 0xfa, 0x22, 0x39, 0x59,
	0x4b, 0xa6, 0x11, 0xc7, 0x9f, 0xaa, 0x3a, 0xa0, 0x60, 0x12, 0x6d, 0xf7, 0xe4, 0x60, 0x56, 0x5d,
	0x32, 0xcb, 0x45, 0xe0, 0x1a, 0xfe, 0x2c, 0xc7, 0x6b, 0xe8, 0xed, 0xee, 0x2b, 0x3a, 0x2e, 0xdb,
	0x73, 0xf1, 0x0a, 0xe0, 0x91, 0xbe, 0x55, 0x4c, 0x8b, 0x32, 0x23, 0xd1, 0x75, 0x9d, 0x2d, 0x07,
	0x2f, 0xb7, 0x70, 0x89, 0xd0, 0xc5, 0xff, 0x46, 0xd4, 0x87, 0xe0, 0xd5, 0xb2, 0xd2, 0x29, 0x9e,
	0x35, 0x93, 0xa3, 0x19, 0xce, 0x6b, 0x15, 0x09, 0x08, 0x96, 0x2a, 0xcb, 0x28, 0xd9, 0x67, 0x1d,
	0xfd, 0x78, 0xd0, 0x9d, 0x9a, 0x58, 0xda, 0x6a, 0x8d, 0xfd, 0x0f, 0x21, 0xa0, 0x33, 0x36, 0x85,
	0xb6, 0x5c, 0xd6, 0xf4, 0x1b, 0x89, 0x08, 0xad, 0xb1, 0xb2, 0x65, 0x8d, 0xda, 0xcd, 0x78, 0x01,
	0xdd, 0x85, 0xfa, 0xa2, 0xa5, 0xd1, 0x54, 0x73, 0xdd, 0xe8, 0x2a, 0x9b, 0x4a, 0x9d, 0x16, 0x32,
	0x6d, 0x00, 0x6e, 0xf4, 0x47, 0xe0, 0xfe, 0x89, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43,
	0x5d, 0xec, 0x29, 0x2e, 0x02, 0x00, 0x00,
}
