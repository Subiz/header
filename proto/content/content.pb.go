// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content/content.proto

/*
Package content is a generated protocol buffer package.

It is generated from these files:
	content/content.proto

It has these top-level messages:
	KeyValue
	Content
	Contents
	Ids
	Id
	Empty
	ListRequest
*/
package content

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

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Content struct {
	Id             string      `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AccountId      string      `protobuf:"bytes,14,opt,name=AccountId" json:"AccountId,omitempty"`
	Description    string      `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
	Title          string      `protobuf:"bytes,3,opt,name=Title" json:"Title,omitempty"`
	Link           string      `protobuf:"bytes,4,opt,name=Link" json:"Link,omitempty"`
	Categories     []string    `protobuf:"bytes,6,rep,name=Categories" json:"Categories,omitempty"`
	AttachmentLink string      `protobuf:"bytes,7,opt,name=AttachmentLink" json:"AttachmentLink,omitempty"`
	Labels         []string    `protobuf:"bytes,8,rep,name=Labels" json:"Labels,omitempty"`
	Availability   bool        `protobuf:"varint,9,opt,name=Availability" json:"Availability,omitempty"`
	Price          string      `protobuf:"bytes,10,opt,name=Price" json:"Price,omitempty"`
	Currency       string      `protobuf:"bytes,11,opt,name=Currency" json:"Currency,omitempty"`
	SalePrice      string      `protobuf:"bytes,12,opt,name=SalePrice" json:"SalePrice,omitempty"`
	Fields         []*KeyValue `protobuf:"bytes,13,rep,name=Fields" json:"Fields,omitempty"`
	RelatedIds     []string    `protobuf:"bytes,15,rep,name=RelatedIds" json:"RelatedIds,omitempty"`
}

func (m *Content) Reset()                    { *m = Content{} }
func (m *Content) String() string            { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()               {}
func (*Content) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Content) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Content) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Content) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Content) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Content) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Content) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Content) GetAttachmentLink() string {
	if m != nil {
		return m.AttachmentLink
	}
	return ""
}

func (m *Content) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Content) GetAvailability() bool {
	if m != nil {
		return m.Availability
	}
	return false
}

func (m *Content) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Content) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Content) GetSalePrice() string {
	if m != nil {
		return m.SalePrice
	}
	return ""
}

func (m *Content) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Content) GetRelatedIds() []string {
	if m != nil {
		return m.RelatedIds
	}
	return nil
}

type Contents struct {
	Contents []*Content `protobuf:"bytes,1,rep,name=Contents" json:"Contents,omitempty"`
}

func (m *Contents) Reset()                    { *m = Contents{} }
func (m *Contents) String() string            { return proto.CompactTextString(m) }
func (*Contents) ProtoMessage()               {}
func (*Contents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Contents) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Ids struct {
	Ids []string `protobuf:"bytes,1,rep,name=Ids" json:"Ids,omitempty"`
}

func (m *Ids) Reset()                    { *m = Ids{} }
func (m *Ids) String() string            { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()               {}
func (*Ids) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ids) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId" json:"AccountId,omitempty"`
	StartId   string `protobuf:"bytes,2,opt,name=StartId" json:"StartId,omitempty"`
	Category  string `protobuf:"bytes,3,opt,name=Category" json:"Category,omitempty"`
	Limit     int32  `protobuf:"varint,4,opt,name=Limit" json:"Limit,omitempty"`
	Keyword   string `protobuf:"bytes,5,opt,name=Keyword" json:"Keyword,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *ListRequest) GetStartId() string {
	if m != nil {
		return m.StartId
	}
	return ""
}

func (m *ListRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyValue)(nil), "content.KeyValue")
	proto.RegisterType((*Content)(nil), "content.Content")
	proto.RegisterType((*Contents)(nil), "content.Contents")
	proto.RegisterType((*Ids)(nil), "content.Ids")
	proto.RegisterType((*Id)(nil), "content.Id")
	proto.RegisterType((*Empty)(nil), "content.Empty")
	proto.RegisterType((*ListRequest)(nil), "content.ListRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContentMgr service

type ContentMgrClient interface {
	Insert(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Empty, error)
	InsertBulk(ctx context.Context, in *Contents, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	DeleteBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error)
	Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Content, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Contents, error)
}

type contentMgrClient struct {
	cc *grpc.ClientConn
}

func NewContentMgrClient(cc *grpc.ClientConn) ContentMgrClient {
	return &contentMgrClient{cc}
}

func (c *contentMgrClient) Insert(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/Insert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) InsertBulk(ctx context.Context, in *Contents, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/InsertBulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) DeleteBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/DeleteBulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := grpc.Invoke(ctx, "/content.ContentMgr/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Contents, error) {
	out := new(Contents)
	err := grpc.Invoke(ctx, "/content.ContentMgr/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContentMgr service

type ContentMgrServer interface {
	Insert(context.Context, *Content) (*Empty, error)
	InsertBulk(context.Context, *Contents) (*Empty, error)
	Delete(context.Context, *Id) (*Empty, error)
	DeleteBulk(context.Context, *Ids) (*Empty, error)
	Read(context.Context, *Id) (*Content, error)
	List(context.Context, *ListRequest) (*Contents, error)
}

func RegisterContentMgrServer(s *grpc.Server, srv ContentMgrServer) {
	s.RegisterService(&_ContentMgr_serviceDesc, srv)
}

func _ContentMgr_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).Insert(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_InsertBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).InsertBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/InsertBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).InsertBulk(ctx, req.(*Contents))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_DeleteBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).DeleteBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/DeleteBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).DeleteBulk(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).Read(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "content.ContentMgr",
	HandlerType: (*ContentMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _ContentMgr_Insert_Handler,
		},
		{
			MethodName: "InsertBulk",
			Handler:    _ContentMgr_InsertBulk_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ContentMgr_Delete_Handler,
		},
		{
			MethodName: "DeleteBulk",
			Handler:    _ContentMgr_DeleteBulk_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ContentMgr_Read_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ContentMgr_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/content.proto",
}

func init() { proto.RegisterFile("content/content.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x5f, 0x8f, 0xd2, 0x5e,
	0x10, 0xdd, 0x52, 0x28, 0x30, 0xf0, 0xe3, 0xb7, 0x3b, 0x41, 0xbd, 0x21, 0xc6, 0x90, 0x3e, 0x28,
	0x26, 0x64, 0x8d, 0xf8, 0xe2, 0x2b, 0xb2, 0x9a, 0x34, 0x60, 0x62, 0xba, 0xc6, 0xf7, 0xd2, 0x4e,
	0xd6, 0x9b, 0x2d, 0x2d, 0xf6, 0xde, 0x6a, 0xfa, 0x31, 0xfc, 0x06, 0xc6, 0x4f, 0x6a, 0xee, 0x1f,
	0xca, 0xdf, 0x27, 0xe6, 0x9c, 0x99, 0x39, 0xbd, 0x33, 0x67, 0x02, 0x3c, 0x89, 0xf3, 0x4c, 0x52,
	0x26, 0xdf, 0xd8, 0xdf, 0xdb, 0x6d, 0x91, 0xcb, 0x1c, 0xdb, 0x16, 0xfa, 0x33, 0xe8, 0x2c, 0xa9,
	0xfa, 0x16, 0xa5, 0x25, 0xe1, 0x35, 0xb8, 0x4b, 0xaa, 0x98, 0x33, 0x76, 0x26, 0xdd, 0x50, 0x85,
	0x38, 0x84, 0x96, 0x4e, 0xb1, 0x86, 0xe6, 0x0c, 0xf0, 0xff, 0xba, 0xd0, 0x5e, 0x98, 0x7e, 0x1c,
	0x40, 0x23, 0x48, 0x6c, 0x4b, 0x23, 0x48, 0xf0, 0x39, 0x74, 0xe7, 0x71, 0x9c, 0x97, 0x99, 0x0c,
	0x12, 0x36, 0xd0, 0xf4, 0x9e, 0xc0, 0x31, 0xf4, 0xee, 0x48, 0xc4, 0x05, 0xdf, 0x4a, 0x9e, 0x67,
	0x56, 0xf5, 0x90, 0x52, 0x5f, 0xfc, 0xca, 0x65, 0x4a, 0xcc, 0x35, 0x5f, 0xd4, 0x00, 0x11, 0x9a,
	0x2b, 0x9e, 0x3d, 0xb2, 0xa6, 0x26, 0x75, 0x8c, 0x2f, 0x00, 0x16, 0x91, 0xa4, 0x87, 0xbc, 0xe0,
	0x24, 0x98, 0x37, 0x76, 0x27, 0xdd, 0xf0, 0x80, 0xc1, 0x97, 0x30, 0x98, 0x4b, 0x19, 0xc5, 0xdf,
	0x37, 0x94, 0x49, 0xdd, 0xdd, 0xd6, 0xdd, 0x27, 0x2c, 0x3e, 0x05, 0x6f, 0x15, 0xad, 0x29, 0x15,
	0xac, 0xa3, 0x35, 0x2c, 0x42, 0x1f, 0xfa, 0xf3, 0x9f, 0x11, 0x4f, 0xa3, 0x35, 0x4f, 0xb9, 0xac,
	0x58, 0x77, 0xec, 0x4c, 0x3a, 0xe1, 0x11, 0xa7, 0x5e, 0xfb, 0xa5, 0xe0, 0x31, 0x31, 0x30, 0xaf,
	0xd5, 0x00, 0x47, 0xd0, 0x59, 0x94, 0x45, 0x41, 0x59, 0x5c, 0xb1, 0x9e, 0x4e, 0xd4, 0x58, 0xed,
	0xe7, 0x3e, 0x4a, 0xc9, 0x74, 0xf5, 0xcd, 0x7e, 0x6a, 0x02, 0x5f, 0x83, 0xf7, 0x89, 0x53, 0x9a,
	0x08, 0xf6, 0xdf, 0xd8, 0x9d, 0xf4, 0x66, 0x37, 0xb7, 0x3b, 0xdb, 0x76, 0x26, 0x85, 0xb6, 0x40,
	0x8d, 0x1f, 0x52, 0x1a, 0x49, 0x4a, 0x82, 0x44, 0xb0, 0xff, 0xcd, 0xf8, 0x7b, 0xc6, 0x7f, 0x0f,
	0x1d, 0xeb, 0x91, 0xc0, 0xe9, 0x3e, 0x66, 0x8e, 0x16, 0xbe, 0xae, 0x85, 0x6d, 0x22, 0xac, 0x2b,
	0xfc, 0x67, 0xe0, 0x06, 0x89, 0x50, 0xd7, 0xa0, 0x94, 0x1d, 0xad, 0xac, 0x42, 0x7f, 0xa8, 0xbc,
	0x3e, 0x75, 0xdc, 0x6f, 0x43, 0xeb, 0xe3, 0x66, 0x2b, 0x2b, 0xff, 0xb7, 0x03, 0xbd, 0x15, 0x17,
	0x32, 0xa4, 0x1f, 0x25, 0x09, 0x79, 0x7c, 0x0a, 0xce, 0xe9, 0x29, 0x30, 0x68, 0xdf, 0xcb, 0xa8,
	0x50, 0x39, 0x73, 0x06, 0x3b, 0xa8, 0xd7, 0x67, 0x6c, 0xac, 0xec, 0x15, 0xd4, 0x58, 0x2d, 0x7c,
	0xc5, 0x37, 0x5c, 0xea, 0x4b, 0x68, 0x85, 0x06, 0x28, 0xad, 0x25, 0x55, 0xbf, 0xf2, 0x22, 0x61,
	0x2d, 0xa3, 0x65, 0xe1, 0xec, 0x4f, 0x03, 0xc0, 0x0e, 0xf6, 0xf9, 0xa1, 0xc0, 0x29, 0x78, 0x41,
	0x26, 0xa8, 0x90, 0x78, 0xb6, 0x80, 0xd1, 0xa0, 0x66, 0xcc, 0x38, 0x57, 0xf8, 0x16, 0xc0, 0x54,
	0x7f, 0x28, 0xd3, 0x47, 0xbc, 0x39, 0xed, 0x10, 0x17, 0x5a, 0x5e, 0x81, 0x77, 0x47, 0x29, 0x49,
	0xc2, 0x5e, 0x9d, 0x0b, 0x92, 0x0b, 0x85, 0x53, 0x00, 0x53, 0xa8, 0xb5, 0xfb, 0x07, 0xc5, 0x97,
	0x65, 0x9b, 0x21, 0x45, 0xc9, 0xb1, 0xe8, 0xd9, 0x08, 0xfa, 0xc9, 0x4d, 0x65, 0x01, 0x0e, 0xeb,
	0xdc, 0x81, 0x23, 0xa3, 0xf3, 0x11, 0xfc, 0xab, 0xb5, 0xa7, 0xff, 0x11, 0xde, 0xfd, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xfb, 0x82, 0xec, 0x0c, 0x2a, 0x04, 0x00, 0x00,
}
