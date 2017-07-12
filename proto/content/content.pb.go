// Code generated by protoc-gen-go.
// source: content/content.proto
// DO NOT EDIT!

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
	Key   string `protobuf:"bytes,1,opt,name=Key,json=key" json:"Key"`
	Value string `protobuf:"bytes,2,opt,name=Value,json=value" json:"Value"`
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
	Id             string      `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id"`
	AccountId      string      `protobuf:"bytes,14,opt,name=AccountId,json=accountId" json:"AccountId"`
	Description    string      `protobuf:"bytes,2,opt,name=Description,json=description" json:"Description"`
	Title          string      `protobuf:"bytes,3,opt,name=Title,json=title" json:"Title"`
	Link           string      `protobuf:"bytes,4,opt,name=Link,json=link" json:"Link"`
	Categories     []string    `protobuf:"bytes,6,rep,name=Categories,json=categories" json:"Categories"`
	AttachmentLink string      `protobuf:"bytes,7,opt,name=AttachmentLink,json=attachmentLink" json:"AttachmentLink"`
	Labels         []string    `protobuf:"bytes,8,rep,name=Labels,json=labels" json:"Labels"`
	Availability   bool        `protobuf:"varint,9,opt,name=Availability,json=availability" json:"Availability"`
	Price          string      `protobuf:"bytes,10,opt,name=Price,json=price" json:"Price"`
	Currency       string      `protobuf:"bytes,11,opt,name=Currency,json=currency" json:"Currency"`
	SalePrice      string      `protobuf:"bytes,12,opt,name=SalePrice,json=salePrice" json:"SalePrice"`
	Fields         []*KeyValue `protobuf:"bytes,13,rep,name=Fields,json=fields" json:"Fields"`
	RelatedIds     []string    `protobuf:"bytes,15,rep,name=RelatedIds,json=relatedIds" json:"RelatedIds"`
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
	Contents []*Content `protobuf:"bytes,1,rep,name=Contents,json=contents" json:"Contents"`
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
	Ids []string `protobuf:"bytes,1,rep,name=Ids,json=ids" json:"Ids"`
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
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id"`
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
	AccountId string `protobuf:"bytes,1,opt,name=AccountId,json=accountId" json:"AccountId"`
	StartId   string `protobuf:"bytes,2,opt,name=StartId,json=startId" json:"StartId"`
	Category  string `protobuf:"bytes,3,opt,name=Category,json=category" json:"Category"`
	Limit     int32  `protobuf:"varint,4,opt,name=Limit,json=limit" json:"Limit"`
	Keyword   string `protobuf:"bytes,5,opt,name=Keyword,json=keyword" json:"Keyword"`
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
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0xe3, 0xf8, 0x27, 0xe3, 0x7c, 0xf9, 0xda, 0x55, 0x81, 0x55, 0x85, 0x50, 0xe4, 0x0b,
	0x08, 0x52, 0x55, 0x44, 0xb9, 0xe1, 0xb6, 0xb4, 0x20, 0x59, 0x0d, 0x12, 0x72, 0x11, 0xf7, 0x1b,
	0xef, 0x50, 0x56, 0xd9, 0xd8, 0xc1, 0xbb, 0x29, 0xf2, 0x63, 0xf0, 0x06, 0x88, 0x27, 0x45, 0xbb,
	0xeb, 0xb8, 0xf9, 0xbb, 0x4a, 0xce, 0x99, 0x99, 0xb3, 0x33, 0x73, 0x46, 0x86, 0x27, 0x45, 0x55,
	0x6a, 0x2c, 0xf5, 0x9b, 0xf6, 0xf7, 0x62, 0x59, 0x57, 0xba, 0x22, 0x51, 0x0b, 0xd3, 0x4b, 0x88,
	0x6f, 0xb1, 0xf9, 0xc6, 0xe4, 0x0a, 0xc9, 0x31, 0xf8, 0xb7, 0xd8, 0x50, 0x6f, 0xec, 0x4d, 0x06,
	0xb9, 0x3f, 0xc7, 0x86, 0x9c, 0x42, 0x60, 0x43, 0xb4, 0x67, 0xb9, 0xe0, 0xc1, 0x80, 0xf4, 0xaf,
	0x0f, 0xd1, 0xb5, 0xab, 0x27, 0x23, 0xe8, 0x65, 0xbc, 0x2d, 0xe9, 0x09, 0x4e, 0x9e, 0xc3, 0xe0,
	0xaa, 0x28, 0xaa, 0x55, 0xa9, 0x33, 0x4e, 0x47, 0x96, 0x1e, 0xb0, 0x35, 0x41, 0xc6, 0x90, 0xdc,
	0xa0, 0x2a, 0x6a, 0xb1, 0xd4, 0xa2, 0x2a, 0x5b, 0xd5, 0x84, 0x3f, 0x52, 0xe6, 0xc5, 0xaf, 0x42,
	0x4b, 0xa4, 0xbe, 0x7b, 0x51, 0x1b, 0x40, 0x08, 0xf4, 0xa7, 0xa2, 0x9c, 0xd3, 0xbe, 0x25, 0xfb,
	0x52, 0x94, 0x73, 0xf2, 0x02, 0xe0, 0x9a, 0x69, 0xbc, 0xaf, 0x6a, 0x81, 0x8a, 0x86, 0x63, 0x7f,
	0x32, 0xc8, 0xa1, 0xe8, 0x18, 0xf2, 0x12, 0x46, 0x57, 0x5a, 0xb3, 0xe2, 0xc7, 0x02, 0x4b, 0x6d,
	0xab, 0x23, 0x5b, 0x3d, 0x62, 0x5b, 0x2c, 0x79, 0x0a, 0xe1, 0x94, 0xcd, 0x50, 0x2a, 0x1a, 0x5b,
	0x8d, 0x50, 0x5a, 0x44, 0x52, 0x18, 0x5e, 0x3d, 0x30, 0x21, 0xd9, 0x4c, 0x48, 0xa1, 0x1b, 0x3a,
	0x18, 0x7b, 0x93, 0x38, 0x1f, 0xb2, 0x0d, 0xce, 0x74, 0xfb, 0xa5, 0x16, 0x05, 0x52, 0x70, 0xdd,
	0x2e, 0x0d, 0x20, 0x67, 0x10, 0x5f, 0xaf, 0xea, 0x1a, 0xcb, 0xa2, 0xa1, 0x89, 0x0d, 0xc4, 0x45,
	0x8b, 0xcd, 0x7e, 0xee, 0x98, 0x44, 0x57, 0x35, 0x74, 0xfb, 0x51, 0x6b, 0x82, 0xbc, 0x86, 0xf0,
	0x93, 0x40, 0xc9, 0x15, 0xfd, 0x6f, 0xec, 0x4f, 0x92, 0xcb, 0x93, 0x8b, 0xb5, 0x6d, 0x6b, 0x93,
	0xf2, 0xf0, 0xbb, 0x4d, 0x30, 0xe3, 0xe7, 0x28, 0x99, 0x46, 0x9e, 0x71, 0x45, 0xff, 0x77, 0xe3,
	0xd7, 0x1d, 0x93, 0xbe, 0x87, 0xb8, 0xf5, 0x48, 0x91, 0xf3, 0xc7, 0xff, 0xd4, 0xb3, 0xc2, 0xc7,
	0x9d, 0x70, 0x1b, 0xc8, 0xe3, 0x96, 0x50, 0xe9, 0x33, 0xf0, 0x33, 0xae, 0xcc, 0x35, 0x18, 0x65,
	0xcf, 0x2a, 0xfb, 0x82, 0xab, 0xf4, 0xd4, 0x78, 0xbd, 0xeb, 0x78, 0x1a, 0x41, 0xf0, 0x71, 0xb1,
	0xd4, 0x4d, 0xfa, 0xdb, 0x83, 0x64, 0x2a, 0x94, 0xce, 0xf1, 0xe7, 0x0a, 0x95, 0xde, 0x3e, 0x05,
	0x6f, 0xf7, 0x14, 0x28, 0x44, 0x77, 0x9a, 0xd5, 0x26, 0xe6, 0xce, 0x20, 0x52, 0x0e, 0xda, 0xf5,
	0x39, 0x1b, 0x9b, 0xf6, 0x0a, 0xe2, 0xd6, 0x56, 0xbb, 0xf0, 0xa9, 0x58, 0x08, 0x6d, 0x2f, 0x21,
	0xc8, 0x03, 0x69, 0x80, 0xd1, 0xba, 0xc5, 0xe6, 0x57, 0x55, 0x73, 0x1a, 0x38, 0xad, 0xb9, 0x83,
	0x97, 0x7f, 0x7a, 0x00, 0xed, 0x84, 0x9f, 0xef, 0x6b, 0x72, 0x0e, 0x61, 0x56, 0x2a, 0xac, 0x35,
	0xd9, 0x5b, 0xc0, 0xd9, 0xa8, 0x63, 0xdc, 0x38, 0x47, 0xe4, 0x2d, 0x80, 0xcb, 0xfe, 0xb0, 0x92,
	0x73, 0x72, 0xb2, 0x5b, 0xa1, 0x0e, 0x94, 0xbc, 0x82, 0xf0, 0x06, 0x25, 0x6a, 0x24, 0x49, 0x17,
	0xcb, 0xf8, 0x81, 0xc4, 0x73, 0x00, 0x97, 0x68, 0xb5, 0x87, 0x1b, 0xc9, 0x87, 0x65, 0xfb, 0x39,
	0x32, 0xbe, 0x2d, 0xba, 0x37, 0x82, 0x6d, 0xb9, 0x6f, 0x2c, 0x20, 0xa7, 0x5d, 0x6c, 0xc3, 0x91,
	0xb3, 0xfd, 0x11, 0xd2, 0xa3, 0x59, 0x68, 0xbf, 0x08, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0x68, 0x51, 0x41, 0x2a, 0x04, 0x00, 0x00,
}
