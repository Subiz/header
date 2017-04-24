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

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type UserAgent struct {
	Id                   int32  `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
	BrowserVersionFull   string `protobuf:"bytes,2,opt,name=BrowserVersionFull,json=browserVersionFull" json:"BrowserVersionFull,omitempty"`
	LayoutEngine         string `protobuf:"bytes,3,opt,name=LayoutEngine,json=layoutEngine" json:"LayoutEngine,omitempty"`
	LayoutEngineVersion  string `protobuf:"bytes,4,opt,name=LayoutEngineVersion,json=layoutEngineVersion" json:"LayoutEngineVersion,omitempty"`
	HardwareArchitecture string `protobuf:"bytes,5,opt,name=HardwareArchitecture,json=hardwareArchitecture" json:"HardwareArchitecture,omitempty"`
	Os                   string `protobuf:"bytes,6,opt,name=Os,json=os" json:"Os,omitempty"`
	BrowserVersion       string `protobuf:"bytes,7,opt,name=BrowserVersion,json=browserVersion" json:"BrowserVersion,omitempty"`
	DeviceType           string `protobuf:"bytes,8,opt,name=DeviceType,json=deviceType" json:"DeviceType,omitempty"`
	UserAgent            string `protobuf:"bytes,9,opt,name=UserAgent,json=userAgent" json:"UserAgent,omitempty"`
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
	String_ string `protobuf:"bytes,1,opt,name=String,json=string" json:"String,omitempty"`
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
	Id int32 `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
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
	Id       int32  `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
	Country  string `protobuf:"bytes,2,opt,name=Country,json=country" json:"Country,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=City,json=city" json:"City,omitempty"`
	TimeZone string `protobuf:"bytes,4,opt,name=TimeZone,json=timeZone" json:"TimeZone,omitempty"`
	Language string `protobuf:"bytes,5,opt,name=Language,json=language" json:"Language,omitempty"`
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Detector service

type DetectorClient interface {
	DetectUserAgent(ctx context.Context, in *String, opts ...grpc.CallOption) (*UserAgent, error)
	ZipUserAgent(ctx context.Context, in *String, opts ...grpc.CallOption) (*Zipped, error)
	UnzipUserAgent(ctx context.Context, in *Zipped, opts ...grpc.CallOption) (*UserAgent, error)
	DetectLocation(ctx context.Context, in *String, opts ...grpc.CallOption) (*Location, error)
	ZipLanguage(ctx context.Context, in *String, opts ...grpc.CallOption) (*Zipped, error)
	UnzipLanguage(ctx context.Context, in *Zipped, opts ...grpc.CallOption) (*String, error)
}

type detectorClient struct {
	cc *grpc.ClientConn
}

func NewDetectorClient(cc *grpc.ClientConn) DetectorClient {
	return &detectorClient{cc}
}

func (c *detectorClient) DetectUserAgent(ctx context.Context, in *String, opts ...grpc.CallOption) (*UserAgent, error) {
	out := new(UserAgent)
	err := grpc.Invoke(ctx, "/Detector/DetectUserAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detectorClient) ZipUserAgent(ctx context.Context, in *String, opts ...grpc.CallOption) (*Zipped, error) {
	out := new(Zipped)
	err := grpc.Invoke(ctx, "/Detector/ZipUserAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detectorClient) UnzipUserAgent(ctx context.Context, in *Zipped, opts ...grpc.CallOption) (*UserAgent, error) {
	out := new(UserAgent)
	err := grpc.Invoke(ctx, "/Detector/UnzipUserAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detectorClient) DetectLocation(ctx context.Context, in *String, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := grpc.Invoke(ctx, "/Detector/DetectLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detectorClient) ZipLanguage(ctx context.Context, in *String, opts ...grpc.CallOption) (*Zipped, error) {
	out := new(Zipped)
	err := grpc.Invoke(ctx, "/Detector/ZipLanguage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *detectorClient) UnzipLanguage(ctx context.Context, in *Zipped, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := grpc.Invoke(ctx, "/Detector/UnzipLanguage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Detector service

type DetectorServer interface {
	DetectUserAgent(context.Context, *String) (*UserAgent, error)
	ZipUserAgent(context.Context, *String) (*Zipped, error)
	UnzipUserAgent(context.Context, *Zipped) (*UserAgent, error)
	DetectLocation(context.Context, *String) (*Location, error)
	ZipLanguage(context.Context, *String) (*Zipped, error)
	UnzipLanguage(context.Context, *Zipped) (*String, error)
}

func RegisterDetectorServer(s *grpc.Server, srv DetectorServer) {
	s.RegisterService(&_Detector_serviceDesc, srv)
}

func _Detector_DetectUserAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectorServer).DetectUserAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Detector/DetectUserAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectorServer).DetectUserAgent(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Detector_ZipUserAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectorServer).ZipUserAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Detector/ZipUserAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectorServer).ZipUserAgent(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Detector_UnzipUserAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Zipped)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectorServer).UnzipUserAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Detector/UnzipUserAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectorServer).UnzipUserAgent(ctx, req.(*Zipped))
	}
	return interceptor(ctx, in, info, handler)
}

func _Detector_DetectLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectorServer).DetectLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Detector/DetectLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectorServer).DetectLocation(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Detector_ZipLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectorServer).ZipLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Detector/ZipLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectorServer).ZipLanguage(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Detector_UnzipLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Zipped)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectorServer).UnzipLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Detector/UnzipLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectorServer).UnzipLanguage(ctx, req.(*Zipped))
	}
	return interceptor(ctx, in, info, handler)
}

var _Detector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Detector",
	HandlerType: (*DetectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetectUserAgent",
			Handler:    _Detector_DetectUserAgent_Handler,
		},
		{
			MethodName: "ZipUserAgent",
			Handler:    _Detector_ZipUserAgent_Handler,
		},
		{
			MethodName: "UnzipUserAgent",
			Handler:    _Detector_UnzipUserAgent_Handler,
		},
		{
			MethodName: "DetectLocation",
			Handler:    _Detector_DetectLocation_Handler,
		},
		{
			MethodName: "ZipLanguage",
			Handler:    _Detector_ZipLanguage_Handler,
		},
		{
			MethodName: "UnzipLanguage",
			Handler:    _Detector_UnzipLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "detector/detector.proto",
}

func init() { proto.RegisterFile("detector/detector.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x53, 0xdb, 0x6e, 0xd4, 0x30,
	0x10, 0x4d, 0xc2, 0x36, 0x97, 0xa1, 0x04, 0x69, 0x5a, 0x81, 0x55, 0x21, 0xb4, 0x18, 0xa9, 0xec,
	0x53, 0x40, 0xe5, 0x0b, 0x4a, 0x0b, 0x02, 0x69, 0x25, 0xa4, 0xa5, 0xe5, 0x21, 0x6f, 0x69, 0x62,
	0xa5, 0x96, 0x82, 0x1d, 0x39, 0x0e, 0x55, 0x78, 0xe2, 0xff, 0xf8, 0x1a, 0xfe, 0x00, 0xc5, 0x71,
	0x42, 0xf6, 0xf2, 0x36, 0x73, 0xce, 0xb1, 0x27, 0xe7, 0x78, 0x02, 0xcf, 0x0b, 0xa6, 0x59, 0xae,
	0xa5, 0x7a, 0x3b, 0x16, 0x49, 0xad, 0xa4, 0x96, 0xf4, 0x8f, 0x07, 0xd1, 0x6d, 0xc3, 0xd4, 0x65,
	0xc9, 0x84, 0xc6, 0x18, 0xbc, 0x2f, 0x05, 0x71, 0x97, 0xee, 0xea, 0x68, 0xe3, 0xf1, 0x02, 0x13,
	0xc0, 0x0f, 0x4a, 0x3e, 0x34, 0x4c, 0x7d, 0x67, 0xaa, 0xe1, 0x52, 0x7c, 0x6a, 0xab, 0x8a, 0x78,
	0x4b, 0x77, 0x15, 0x6d, 0xf0, 0x6e, 0x8f, 0x41, 0x0a, 0xc7, 0xeb, 0xac, 0x93, 0xad, 0xfe, 0x28,
	0x4a, 0x2e, 0x18, 0x79, 0x64, 0x94, 0xc7, 0xd5, 0x0c, 0xc3, 0x77, 0x70, 0x32, 0xd7, 0xd8, 0xe3,
	0x64, 0x61, 0xa4, 0x27, 0xd5, 0x3e, 0x85, 0x17, 0x70, 0xfa, 0x39, 0x53, 0xc5, 0x43, 0xa6, 0xd8,
	0xa5, 0xca, 0xef, 0x79, 0x6f, 0xa0, 0x55, 0x8c, 0x1c, 0x99, 0x23, 0xa7, 0xf7, 0x07, 0xb8, 0xde,
	0xc9, 0xd7, 0x86, 0xf8, 0x46, 0xe1, 0xc9, 0x06, 0xcf, 0x21, 0xde, 0x76, 0x42, 0x02, 0xc3, 0xc5,
	0xdb, 0x2e, 0xf0, 0x25, 0xc0, 0x35, 0xfb, 0xc9, 0x73, 0x76, 0xd3, 0xd5, 0x8c, 0x84, 0x46, 0x03,
	0xc5, 0x84, 0xe0, 0x8b, 0x59, 0x5c, 0x24, 0x32, 0x74, 0xd4, 0x8e, 0x00, 0x5d, 0x82, 0xff, 0x4d,
	0x2b, 0x2e, 0x4a, 0x7c, 0x36, 0x56, 0x26, 0xcd, 0x68, 0xe3, 0x37, 0xa6, 0xa3, 0x04, 0xfc, 0x94,
	0xd7, 0x35, 0x2b, 0x76, 0xb3, 0xa6, 0xbf, 0x5d, 0x08, 0xd7, 0x32, 0xcf, 0x74, 0xff, 0x19, 0xbb,
	0x0f, 0x41, 0x20, 0xb8, 0x92, 0xad, 0xd0, 0xaa, 0xb3, 0xe9, 0x07, 0xf9, 0xd0, 0x22, 0xc2, 0xe2,
	0x8a, 0xeb, 0xce, 0x46, 0xbd, 0xc8, 0xb9, 0xee, 0xf0, 0x0c, 0xc2, 0x1b, 0xfe, 0x83, 0xa5, 0x52,
	0x30, 0x9b, 0x6b, 0xa8, 0x6d, 0xdf, 0x73, 0xeb, 0x4c, 0x94, 0x6d, 0x56, 0x8e, 0x01, 0x86, 0x95,
	0xed, 0x2f, 0xfe, 0xba, 0x10, 0x5e, 0xdb, 0xfd, 0xc0, 0x15, 0x3c, 0x1d, 0xea, 0xff, 0xeb, 0x11,
	0x24, 0x83, 0xa7, 0x33, 0x48, 0x26, 0x90, 0x3a, 0xfd, 0xab, 0xa7, 0xbc, 0x3e, 0x20, 0x0b, 0x92,
	0xc1, 0x2b, 0x75, 0xf0, 0x0d, 0xc4, 0xb7, 0xe2, 0xd7, 0xb6, 0x6a, 0x20, 0x77, 0x2e, 0x3b, 0x87,
	0x78, 0x18, 0x3b, 0x65, 0x31, 0x5d, 0x17, 0x25, 0x23, 0x46, 0x1d, 0x7c, 0x05, 0x8f, 0x53, 0x5e,
	0x8f, 0x56, 0x0e, 0xce, 0x7c, 0x0d, 0x4f, 0xcc, 0xcc, 0x99, 0xc8, 0x8e, 0x1c, 0xd5, 0xd4, 0xb9,
	0xf3, 0xcd, 0x7f, 0xf0, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x85, 0xed, 0x87, 0x22,
	0x03, 0x00, 0x00,
}
