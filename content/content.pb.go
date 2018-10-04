// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content/content.proto

package content

import (
	fmt "fmt"
	common "git.subiz.net/header/common"
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

type SearchContentRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Query                *string         `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	Limit                *int32          `protobuf:"varint,6,opt,name=limit" json:"limit,omitempty"`
	Plan                 *string         `protobuf:"bytes,7,opt,name=plan" json:"plan,omitempty"`
	OrderBy              *string         `protobuf:"bytes,10,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SearchContentRequest) Reset()         { *m = SearchContentRequest{} }
func (m *SearchContentRequest) String() string { return proto.CompactTextString(m) }
func (*SearchContentRequest) ProtoMessage()    {}
func (*SearchContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{0}
}

func (m *SearchContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchContentRequest.Unmarshal(m, b)
}
func (m *SearchContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchContentRequest.Marshal(b, m, deterministic)
}
func (m *SearchContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchContentRequest.Merge(m, src)
}
func (m *SearchContentRequest) XXX_Size() int {
	return xxx_messageInfo_SearchContentRequest.Size(m)
}
func (m *SearchContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchContentRequest proto.InternalMessageInfo

func (m *SearchContentRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *SearchContentRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *SearchContentRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *SearchContentRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *SearchContentRequest) GetPlan() string {
	if m != nil && m.Plan != nil {
		return *m.Plan
	}
	return ""
}

func (m *SearchContentRequest) GetOrderBy() string {
	if m != nil && m.OrderBy != nil {
		return *m.OrderBy
	}
	return ""
}

type KeyValue struct {
	Key                  *string  `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{1}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Content struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Sbid                 *string         `protobuf:"bytes,2,opt,name=sbid" json:"sbid,omitempty"`
	Id                   *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	AccountId            *string         `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Description          *string         `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Title                *string         `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Url                  *string         `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
	Labels               []string        `protobuf:"bytes,8,rep,name=labels" json:"labels,omitempty"`
	Availability         *string         `protobuf:"bytes,9,opt,name=availability" json:"availability,omitempty"`
	Price                *float32        `protobuf:"fixed32,10,opt,name=price" json:"price,omitempty"`
	Currency             *string         `protobuf:"bytes,11,opt,name=currency" json:"currency,omitempty"`
	SalePrice            *float32        `protobuf:"fixed32,12,opt,name=sale_price,json=salePrice" json:"sale_price,omitempty"`
	Fields               []*KeyValue     `protobuf:"bytes,13,rep,name=fields" json:"fields,omitempty"`
	Categories           []string        `protobuf:"bytes,14,rep,name=categories" json:"categories,omitempty"`
	Relates              []string        `protobuf:"bytes,15,rep,name=relates" json:"relates,omitempty"`
	AttachmentUrls       []string        `protobuf:"bytes,16,rep,name=attachment_urls,json=attachmentUrls" json:"attachment_urls,omitempty"`
	Type                 *string         `protobuf:"bytes,17,opt,name=type" json:"type,omitempty"`
	Created              *int64          `protobuf:"varint,18,opt,name=created" json:"created,omitempty"`
	Updated              *int64          `protobuf:"varint,19,opt,name=updated" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{2}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Content) GetSbid() string {
	if m != nil && m.Sbid != nil {
		return *m.Sbid
	}
	return ""
}

func (m *Content) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Content) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Content) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Content) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Content) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Content) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Content) GetAvailability() string {
	if m != nil && m.Availability != nil {
		return *m.Availability
	}
	return ""
}

func (m *Content) GetPrice() float32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Content) GetCurrency() string {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return ""
}

func (m *Content) GetSalePrice() float32 {
	if m != nil && m.SalePrice != nil {
		return *m.SalePrice
	}
	return 0
}

func (m *Content) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Content) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Content) GetRelates() []string {
	if m != nil {
		return m.Relates
	}
	return nil
}

func (m *Content) GetAttachmentUrls() []string {
	if m != nil {
		return m.AttachmentUrls
	}
	return nil
}

func (m *Content) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Content) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Content) GetUpdated() int64 {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return 0
}

type Contents struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Contents             []*Content      `protobuf:"bytes,2,rep,name=contents" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Contents) Reset()         { *m = Contents{} }
func (m *Contents) String() string { return proto.CompactTextString(m) }
func (*Contents) ProtoMessage()    {}
func (*Contents) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{3}
}

func (m *Contents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contents.Unmarshal(m, b)
}
func (m *Contents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contents.Marshal(b, m, deterministic)
}
func (m *Contents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contents.Merge(m, src)
}
func (m *Contents) XXX_Size() int {
	return xxx_messageInfo_Contents.Size(m)
}
func (m *Contents) XXX_DiscardUnknown() {
	xxx_messageInfo_Contents.DiscardUnknown(m)
}

var xxx_messageInfo_Contents proto.InternalMessageInfo

func (m *Contents) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Contents) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ListRequest struct {
	AccountId            *string  `protobuf:"bytes,6,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Anchor               *string  `protobuf:"bytes,2,opt,name=anchor" json:"anchor,omitempty"`
	Category             *string  `protobuf:"bytes,3,opt,name=category" json:"category,omitempty"`
	Limit                *int32   `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	Label                *string  `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	Query                *string  `protobuf:"bytes,7,opt,name=query" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ListRequest) GetAnchor() string {
	if m != nil && m.Anchor != nil {
		return *m.Anchor
	}
	return ""
}

func (m *ListRequest) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *ListRequest) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *ListRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchContentRequest)(nil), "content.SearchContentRequest")
	proto.RegisterType((*KeyValue)(nil), "content.KeyValue")
	proto.RegisterType((*Content)(nil), "content.Content")
	proto.RegisterType((*Contents)(nil), "content.Contents")
	proto.RegisterType((*ListRequest)(nil), "content.ListRequest")
}

func init() { proto.RegisterFile("content/content.proto", fileDescriptor_d1ed60e503df5706) }

var fileDescriptor_d1ed60e503df5706 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6b, 0x1b, 0x3d,
	0x10, 0xc6, 0x1f, 0xf1, 0xc7, 0x38, 0x71, 0x62, 0xe5, 0x03, 0xbd, 0x86, 0xbc, 0xb8, 0xbe, 0xd4,
	0x29, 0xc1, 0x01, 0x17, 0x7a, 0x6f, 0xda, 0x52, 0x42, 0x5b, 0x08, 0x1b, 0xd2, 0x4b, 0x0f, 0x46,
	0xd6, 0x4e, 0x6d, 0x51, 0x79, 0x77, 0x23, 0x69, 0x43, 0xb6, 0xbf, 0xa5, 0xa7, 0xfe, 0x85, 0xfe,
	0xbe, 0x42, 0x91, 0x56, 0xbb, 0x76, 0xec, 0x1c, 0x72, 0x5a, 0x3d, 0xcf, 0xcc, 0xb3, 0xd2, 0x3c,
	0x9a, 0x11, 0x1c, 0xf3, 0x38, 0x32, 0x18, 0x99, 0x0b, 0xff, 0x1d, 0x27, 0x2a, 0x36, 0x31, 0x69,
	0x7a, 0xd8, 0x1f, 0xcd, 0x85, 0x19, 0xeb, 0x74, 0x26, 0x7e, 0x8e, 0x23, 0x34, 0x17, 0x0b, 0x64,
	0x21, 0xaa, 0x0b, 0x1e, 0x2f, 0x97, 0x71, 0xe4, 0x3f, 0xb9, 0x64, 0xf8, 0xa7, 0x02, 0x47, 0x37,
	0xc8, 0x14, 0x5f, 0xbc, 0xcb, 0xb5, 0x01, 0xde, 0xa5, 0xa8, 0x0d, 0x79, 0x01, 0x35, 0x6e, 0x1e,
	0x68, 0x65, 0x50, 0x19, 0x75, 0x26, 0xfb, 0x63, 0x2f, 0x72, 0x49, 0x0f, 0x26, 0xb0, 0x31, 0x72,
	0x0a, 0xc0, 0x38, 0x8f, 0xd3, 0xc8, 0x4c, 0x45, 0x48, 0xab, 0x83, 0xca, 0xa8, 0x1d, 0xb4, 0x3d,
	0x73, 0x15, 0x92, 0x23, 0xd8, 0xb9, 0x4b, 0x51, 0x65, 0xb4, 0xe6, 0x22, 0x39, 0xb0, 0xac, 0x14,
	0x4b, 0x61, 0x68, 0x63, 0x50, 0x19, 0xed, 0x04, 0x39, 0x20, 0x04, 0xea, 0x89, 0x64, 0x11, 0x6d,
	0xba, 0x54, 0xb7, 0x26, 0xff, 0x41, 0x2b, 0x56, 0x21, 0xaa, 0xe9, 0x2c, 0xa3, 0xe0, 0xf8, 0xa6,
	0xc3, 0x97, 0xd9, 0x70, 0x02, 0xad, 0x4f, 0x98, 0x7d, 0x65, 0x32, 0x45, 0x72, 0x00, 0xb5, 0x1f,
	0x58, 0x6c, 0x62, 0x97, 0x76, 0x8b, 0x7b, 0x1b, 0xa2, 0xf5, 0x7c, 0x63, 0x07, 0x86, 0xbf, 0xea,
	0xd0, 0xf4, 0x35, 0x3e, 0xa7, 0x38, 0x02, 0x75, 0x3d, 0x2b, 0xcb, 0x72, 0x6b, 0xd2, 0x85, 0xaa,
	0x08, 0xfd, 0x4e, 0x55, 0x11, 0x6e, 0x18, 0x50, 0xdf, 0x34, 0x60, 0x00, 0x9d, 0x10, 0x35, 0x57,
	0x22, 0x31, 0x22, 0x8e, 0xe8, 0x8e, 0x8b, 0xaf, 0x53, 0xf6, 0xa4, 0x46, 0x18, 0x89, 0xce, 0x8c,
	0x76, 0x90, 0x03, 0x5b, 0x51, 0xaa, 0xa4, 0xf7, 0xc2, 0x2e, 0xc9, 0x09, 0x34, 0x24, 0x9b, 0xa1,
	0xd4, 0xb4, 0x35, 0xa8, 0x8d, 0xda, 0x81, 0x47, 0x64, 0x08, 0xbb, 0xec, 0x9e, 0x09, 0xc9, 0x66,
	0x42, 0x0a, 0x93, 0xd1, 0xb6, 0x93, 0x3c, 0xe2, 0xec, 0x1e, 0x89, 0x12, 0x1c, 0x9d, 0x87, 0xd5,
	0x20, 0x07, 0xa4, 0x0f, 0x2d, 0x9e, 0x2a, 0x85, 0x11, 0xcf, 0x68, 0xc7, 0xa9, 0x4a, 0x6c, 0xcb,
	0xd2, 0x4c, 0xe2, 0x34, 0x97, 0xed, 0x3a, 0x59, 0xdb, 0x32, 0xd7, 0x4e, 0x7a, 0x06, 0x8d, 0xef,
	0x02, 0x65, 0xa8, 0xe9, 0xde, 0xa0, 0x36, 0xea, 0x4c, 0x7a, 0xe3, 0xa2, 0x0b, 0x8b, 0x3b, 0x09,
	0x7c, 0x02, 0xf9, 0x1f, 0x80, 0x33, 0x83, 0xf3, 0x58, 0x09, 0xd4, 0xb4, 0xeb, 0xce, 0xbe, 0xc6,
	0x10, 0x0a, 0x4d, 0x85, 0x92, 0x19, 0xd4, 0x74, 0xdf, 0x05, 0x0b, 0x48, 0x5e, 0xc2, 0x3e, 0x33,
	0x86, 0xf1, 0xc5, 0x12, 0x23, 0x33, 0x4d, 0x95, 0xd4, 0xf4, 0xc0, 0x65, 0x74, 0x57, 0xf4, 0xad,
	0x92, 0xda, 0xde, 0x93, 0xc9, 0x12, 0xa4, 0xbd, 0xfc, 0x9e, 0xec, 0xda, 0xfe, 0x96, 0x2b, 0x64,
	0x06, 0x43, 0x4a, 0x06, 0x95, 0x51, 0x2d, 0x28, 0xa0, 0x8d, 0xa4, 0x49, 0xe8, 0x22, 0x87, 0x79,
	0xc4, 0xc3, 0xe1, 0x37, 0x68, 0xf9, 0xee, 0xd0, 0xcf, 0x69, 0x8f, 0x73, 0x68, 0xf9, 0xaa, 0x35,
	0xad, 0x3a, 0x1b, 0x0e, 0x4a, 0x1b, 0x8a, 0x49, 0x2a, 0x33, 0x86, 0xbf, 0x2b, 0xd0, 0xf9, 0x2c,
	0x74, 0x39, 0x5c, 0x8f, 0x1b, 0xa7, 0xb1, 0xd9, 0x38, 0x27, 0xd0, 0x60, 0x11, 0x5f, 0xc4, 0xca,
	0x77, 0x9f, 0x47, 0xee, 0xd2, 0x72, 0xf3, 0x8a, 0x7e, 0x2f, 0xf1, 0x6a, 0xae, 0xea, 0xeb, 0x73,
	0x65, 0x59, 0xdb, 0x2a, 0xbe, 0xf9, 0x72, 0xb0, 0x9a, 0xcc, 0xe6, 0xda, 0x64, 0x4e, 0xfe, 0x56,
	0xe1, 0xf8, 0x06, 0xd5, 0xbd, 0xe0, 0xe8, 0x2b, 0xf8, 0xc2, 0x22, 0x36, 0x47, 0x45, 0xce, 0xa0,
	0x7e, 0x2d, 0xa2, 0x39, 0x39, 0x2c, 0xac, 0xb0, 0xc8, 0xd7, 0xd2, 0xdf, 0x2d, 0xc9, 0x38, 0x9a,
	0x93, 0x37, 0xd0, 0xbd, 0x5a, 0x26, 0xb1, 0x32, 0xa5, 0x99, 0xbd, 0x4d, 0x5f, 0x74, 0x7f, 0x9b,
	0x22, 0xaf, 0x61, 0xef, 0x36, 0xd1, 0x58, 0xea, 0xc8, 0x96, 0x9d, 0xfd, 0x2d, 0x86, 0x9c, 0x43,
	0xf7, 0x3d, 0x4a, 0x34, 0x58, 0xfe, 0xa6, 0x53, 0x1c, 0xe6, 0x2a, 0xd4, 0xfd, 0xbd, 0x02, 0x7c,
	0x58, 0x26, 0x26, 0x23, 0x63, 0xe8, 0xe5, 0xd9, 0x6f, 0xa5, 0x2c, 0x05, 0xb0, 0x12, 0x6c, 0xe6,
	0xbf, 0x02, 0xf8, 0x88, 0xe5, 0x79, 0xd6, 0x13, 0xb7, 0x4f, 0x72, 0x09, 0xdd, 0x47, 0xaf, 0xa8,
	0x26, 0xa7, 0x65, 0xce, 0x53, 0xcf, 0xeb, 0x13, 0x16, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xd7,
	0x4e, 0xf1, 0x02, 0xd5, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceContentManagerClient is the client API for ServiceContentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceContentManagerClient interface {
	Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error)
	ImportContents(ctx context.Context, in *Contents, opts ...grpc.CallOption) (*Contents, error)
	UpsertContent(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error)
	DeleteContents(ctx context.Context, in *common.Ids, opts ...grpc.CallOption) (*common.Empty, error)
	DeleteAllContents(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*common.Empty, error)
	GetContent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*Content, error)
	SearchContents(ctx context.Context, in *SearchContentRequest, opts ...grpc.CallOption) (*Contents, error)
}

type serviceContentManagerClient struct {
	cc *grpc.ClientConn
}

func NewServiceContentManagerClient(cc *grpc.ClientConn) ServiceContentManagerClient {
	return &serviceContentManagerClient{cc}
}

func (c *serviceContentManagerClient) Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error) {
	out := new(common.Pong)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) ImportContents(ctx context.Context, in *Contents, opts ...grpc.CallOption) (*Contents, error) {
	out := new(Contents)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/ImportContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) UpsertContent(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/UpsertContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) DeleteContents(ctx context.Context, in *common.Ids, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/DeleteContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) DeleteAllContents(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/DeleteAllContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) GetContent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/GetContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) SearchContents(ctx context.Context, in *SearchContentRequest, opts ...grpc.CallOption) (*Contents, error) {
	out := new(Contents)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/SearchContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceContentManagerServer is the server API for ServiceContentManager service.
type ServiceContentManagerServer interface {
	Ping(context.Context, *common.PingRequest) (*common.Pong, error)
	ImportContents(context.Context, *Contents) (*Contents, error)
	UpsertContent(context.Context, *Content) (*Content, error)
	DeleteContents(context.Context, *common.Ids) (*common.Empty, error)
	DeleteAllContents(context.Context, *common.Id) (*common.Empty, error)
	GetContent(context.Context, *common.Id) (*Content, error)
	SearchContents(context.Context, *SearchContentRequest) (*Contents, error)
}

func RegisterServiceContentManagerServer(s *grpc.Server, srv ServiceContentManagerServer) {
	s.RegisterService(&_ServiceContentManager_serviceDesc, srv)
}

func _ServiceContentManager_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).Ping(ctx, req.(*common.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_ImportContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).ImportContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/ImportContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).ImportContents(ctx, req.(*Contents))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_UpsertContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).UpsertContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/UpsertContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).UpsertContent(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_DeleteContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).DeleteContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/DeleteContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).DeleteContents(ctx, req.(*common.Ids))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_DeleteAllContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).DeleteAllContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/DeleteAllContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).DeleteAllContents(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).GetContent(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_SearchContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).SearchContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/SearchContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).SearchContents(ctx, req.(*SearchContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceContentManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "content.ServiceContentManager",
	HandlerType: (*ServiceContentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ServiceContentManager_Ping_Handler,
		},
		{
			MethodName: "ImportContents",
			Handler:    _ServiceContentManager_ImportContents_Handler,
		},
		{
			MethodName: "UpsertContent",
			Handler:    _ServiceContentManager_UpsertContent_Handler,
		},
		{
			MethodName: "DeleteContents",
			Handler:    _ServiceContentManager_DeleteContents_Handler,
		},
		{
			MethodName: "DeleteAllContents",
			Handler:    _ServiceContentManager_DeleteAllContents_Handler,
		},
		{
			MethodName: "GetContent",
			Handler:    _ServiceContentManager_GetContent_Handler,
		},
		{
			MethodName: "SearchContents",
			Handler:    _ServiceContentManager_SearchContents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/content.proto",
}
