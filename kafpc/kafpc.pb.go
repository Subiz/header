// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kafpc/kafpc.proto

package kafpc

import (
	fmt "fmt"
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

type Request struct {
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ResponseHost         string   `protobuf:"bytes,3,opt,name=response_host,json=responseHost,proto3" json:"response_host,omitempty"`
	Body                 []byte   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Created              int64    `protobuf:"varint,7,opt,name=created,proto3" json:"created,omitempty"`
	Forget               bool     `protobuf:"varint,8,opt,name=forget,proto3" json:"forget,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_48733a375d5689d0, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetResponseHost() string {
	if m != nil {
		return m.ResponseHost
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Request) GetForget() bool {
	if m != nil {
		return m.Forget
	}
	return false
}

type Response struct {
	RequestId            string   `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Code                 int32    `protobuf:"varint,5,opt,name=code,proto3" json:"code,omitempty"`
	Error                []byte   `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Created              int64    `protobuf:"varint,8,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_48733a375d5689d0, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetError() []byte {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_48733a375d5689d0, []int{2}
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

func init() {
	proto.RegisterType((*Request)(nil), "kafpc.Request")
	proto.RegisterType((*Response)(nil), "kafpc.Response")
	proto.RegisterType((*Empty)(nil), "kafpc.Empty")
}

func init() { proto.RegisterFile("kafpc/kafpc.proto", fileDescriptor_48733a375d5689d0) }

var fileDescriptor_48733a375d5689d0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xdd, 0x4a, 0xc4, 0x30,
	0x10, 0x85, 0xc9, 0xee, 0xa6, 0xed, 0x0e, 0x55, 0x71, 0x10, 0x09, 0x82, 0x50, 0x2a, 0x48, 0xaf,
	0x76, 0x41, 0x9f, 0x41, 0x50, 0xbc, 0xcb, 0x0b, 0x2c, 0xdd, 0x66, 0xd6, 0x2d, 0xfe, 0x24, 0x26,
	0xf1, 0xa2, 0x97, 0xbe, 0x85, 0x8f, 0x2b, 0x3b, 0x4d, 0xc5, 0xbd, 0x09, 0xe7, 0x9c, 0xcc, 0x30,
	0x1f, 0x07, 0xce, 0x5f, 0xdb, 0x9d, 0xeb, 0xd6, 0xfc, 0xae, 0x9c, 0xb7, 0xd1, 0xa2, 0x64, 0x53,
	0xff, 0x08, 0xc8, 0x35, 0x7d, 0x7e, 0x51, 0x88, 0x78, 0x0a, 0xb3, 0xde, 0xa8, 0x59, 0x25, 0x9a,
	0xa5, 0x9e, 0xf5, 0x06, 0x6f, 0xe0, 0xc4, 0x53, 0x70, 0xf6, 0x23, 0xd0, 0x66, 0x6f, 0x43, 0x54,
	0x73, 0xfe, 0x2a, 0xa7, 0xf0, 0xd1, 0x86, 0x88, 0x08, 0x8b, 0xad, 0x35, 0x83, 0x92, 0x95, 0x68,
	0x4a, 0xcd, 0xfa, 0x90, 0xb9, 0x36, 0xee, 0x55, 0xc6, 0xf3, 0xac, 0x51, 0x41, 0xde, 0x79, 0x6a,
	0x23, 0x19, 0x95, 0x57, 0xa2, 0x99, 0xeb, 0xc9, 0xe2, 0x25, 0x64, 0x3b, 0xeb, 0x5f, 0x28, 0xaa,
	0xa2, 0x12, 0x4d, 0xa1, 0x93, 0xab, 0xbf, 0x05, 0x14, 0x3a, 0x9d, 0xc2, 0x6b, 0x00, 0x3f, 0x62,
	0x6e, 0x7a, 0x93, 0x40, 0x96, 0x29, 0x79, 0x32, 0x7f, 0x14, 0x8b, 0x63, 0x8a, 0xce, 0x1a, 0x62,
	0x32, 0xa9, 0x59, 0xe3, 0x05, 0x48, 0xf2, 0xde, 0x7a, 0x46, 0x2b, 0xf5, 0x68, 0xfe, 0xb3, 0x15,
	0x47, 0x6c, 0x75, 0x0e, 0xf2, 0xe1, 0xdd, 0xc5, 0xe1, 0x6e, 0x0d, 0xf2, 0xf9, 0x50, 0x18, 0xde,
	0x82, 0xd4, 0xe4, 0xde, 0x06, 0x3c, 0x5b, 0x8d, 0x75, 0x4e, 0x88, 0x57, 0x65, 0x0a, 0x78, 0x61,
	0x9b, 0x71, 0xcd, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc0, 0x26, 0x59, 0x7b, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KafpcClient is the client API for Kafpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KafpcClient interface {
	Reply(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Empty, error)
}

type kafpcClient struct {
	cc *grpc.ClientConn
}

func NewKafpcClient(cc *grpc.ClientConn) KafpcClient {
	return &kafpcClient{cc}
}

func (c *kafpcClient) Reply(ctx context.Context, in *Response, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kafpc.Kafpc/Reply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafpcServer is the server API for Kafpc service.
type KafpcServer interface {
	Reply(context.Context, *Response) (*Empty, error)
}

func RegisterKafpcServer(s *grpc.Server, srv KafpcServer) {
	s.RegisterService(&_Kafpc_serviceDesc, srv)
}

func _Kafpc_Reply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafpcServer).Reply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kafpc.Kafpc/Reply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafpcServer).Reply(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kafpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kafpc.Kafpc",
	HandlerType: (*KafpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reply",
			Handler:    _Kafpc_Reply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kafpc/kafpc.proto",
}
