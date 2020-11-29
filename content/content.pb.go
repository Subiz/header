// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: content.proto

package content

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/subiz/header/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Content_Availability int32

const (
	Content_in_stock     Content_Availability = 1
	Content_out_of_stock Content_Availability = 2
	Content_preorder     Content_Availability = 3
	Content_discontinued Content_Availability = 4
)

// Enum value maps for Content_Availability.
var (
	Content_Availability_name = map[int32]string{
		1: "in_stock",
		2: "out_of_stock",
		3: "preorder",
		4: "discontinued",
	}
	Content_Availability_value = map[string]int32{
		"in_stock":     1,
		"out_of_stock": 2,
		"preorder":     3,
		"discontinued": 4,
	}
)

func (x Content_Availability) Enum() *Content_Availability {
	p := new(Content_Availability)
	*p = x
	return p
}

func (x Content_Availability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Content_Availability) Descriptor() protoreflect.EnumDescriptor {
	return file_content_proto_enumTypes[0].Descriptor()
}

func (Content_Availability) Type() protoreflect.EnumType {
	return &file_content_proto_enumTypes[0]
}

func (x Content_Availability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Content_Availability) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Content_Availability(num)
	return nil
}

// Deprecated: Use Content_Availability.Descriptor instead.
func (Content_Availability) EnumDescriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{4, 0}
}

type LookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Url       *string         `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

func (x *LookupRequest) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *LookupRequest) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *LookupRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type LinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Url       *string         `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Ids       []string        `protobuf:"bytes,4,rep,name=ids" json:"ids,omitempty"`
}

func (x *LinkRequest) Reset() {
	*x = LinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkRequest) ProtoMessage() {}

func (x *LinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkRequest.ProtoReflect.Descriptor instead.
func (*LinkRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{1}
}

func (x *LinkRequest) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *LinkRequest) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *LinkRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *LinkRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type SearchContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx         *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId   *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id          *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Title       *string         `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Url         *string         `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Labels      *string         `protobuf:"bytes,6,opt,name=labels" json:"labels,omitempty"`
	Categories  *string         `protobuf:"bytes,10,opt,name=categories" json:"categories,omitempty"`
	Relates     *string         `protobuf:"bytes,11,opt,name=relates" json:"relates,omitempty"`
	Fieldkeys   *string         `protobuf:"bytes,12,opt,name=fieldkeys" json:"fieldkeys,omitempty"`
	Query       *string         `protobuf:"bytes,13,opt,name=query" json:"query,omitempty"`
	Limit       *int32          `protobuf:"varint,14,opt,name=limit" json:"limit,omitempty"`
	Anchor      *string         `protobuf:"bytes,15,opt,name=anchor" json:"anchor,omitempty"`
	Stringify   *string         `protobuf:"bytes,16,opt,name=stringify" json:"stringify,omitempty"`
	Fieldvalues *string         `protobuf:"bytes,17,opt,name=fieldvalues" json:"fieldvalues,omitempty"`
	Created     *int64          `protobuf:"varint,18,opt,name=created" json:"created,omitempty"`
}

func (x *SearchContentRequest) Reset() {
	*x = SearchContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchContentRequest) ProtoMessage() {}

func (x *SearchContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchContentRequest.ProtoReflect.Descriptor instead.
func (*SearchContentRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2}
}

func (x *SearchContentRequest) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *SearchContentRequest) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *SearchContentRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *SearchContentRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *SearchContentRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *SearchContentRequest) GetLabels() string {
	if x != nil && x.Labels != nil {
		return *x.Labels
	}
	return ""
}

func (x *SearchContentRequest) GetCategories() string {
	if x != nil && x.Categories != nil {
		return *x.Categories
	}
	return ""
}

func (x *SearchContentRequest) GetRelates() string {
	if x != nil && x.Relates != nil {
		return *x.Relates
	}
	return ""
}

func (x *SearchContentRequest) GetFieldkeys() string {
	if x != nil && x.Fieldkeys != nil {
		return *x.Fieldkeys
	}
	return ""
}

func (x *SearchContentRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *SearchContentRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *SearchContentRequest) GetAnchor() string {
	if x != nil && x.Anchor != nil {
		return *x.Anchor
	}
	return ""
}

func (x *SearchContentRequest) GetStringify() string {
	if x != nil && x.Stringify != nil {
		return *x.Stringify
	}
	return ""
}

func (x *SearchContentRequest) GetFieldvalues() string {
	if x != nil && x.Fieldvalues != nil {
		return *x.Fieldvalues
	}
	return ""
}

func (x *SearchContentRequest) GetCreated() int64 {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return 0
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value *string `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3}
}

func (x *KeyValue) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	// optional string sbid = 2;
	Id             *string     `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"` // user input id
	AccountId      *string     `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Description    *string     `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Title          *string     `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Url            *string     `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
	Labels         []string    `protobuf:"bytes,8,rep,name=labels" json:"labels,omitempty"`
	Availability   *string     `protobuf:"bytes,9,opt,name=availability" json:"availability,omitempty"`
	Price          *float32    `protobuf:"fixed32,10,opt,name=price" json:"price,omitempty"`
	Currency       *string     `protobuf:"bytes,11,opt,name=currency" json:"currency,omitempty"`
	SalePrice      *float32    `protobuf:"fixed32,12,opt,name=sale_price,json=salePrice" json:"sale_price,omitempty"`
	Fields         []*KeyValue `protobuf:"bytes,13,rep,name=fields" json:"fields,omitempty"`
	Categories     []string    `protobuf:"bytes,14,rep,name=categories" json:"categories,omitempty"`
	Relates        []string    `protobuf:"bytes,15,rep,name=relates" json:"relates,omitempty"` //  TODO: delete
	Created        *int64      `protobuf:"varint,17,opt,name=created" json:"created,omitempty"`
	Updated        *int64      `protobuf:"varint,18,opt,name=updated" json:"updated,omitempty"`
	AttachmentUrls []string    `protobuf:"bytes,20,rep,name=attachment_urls,json=attachmentUrls" json:"attachment_urls,omitempty"`
	RelatedIds     []string    `protobuf:"bytes,21,rep,name=related_ids,json=relatedIds" json:"related_ids,omitempty"` // releated content ids
	ImageUrl       *string     `protobuf:"bytes,22,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{4}
}

func (x *Content) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *Content) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Content) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Content) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Content) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *Content) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Content) GetAvailability() string {
	if x != nil && x.Availability != nil {
		return *x.Availability
	}
	return ""
}

func (x *Content) GetPrice() float32 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *Content) GetCurrency() string {
	if x != nil && x.Currency != nil {
		return *x.Currency
	}
	return ""
}

func (x *Content) GetSalePrice() float32 {
	if x != nil && x.SalePrice != nil {
		return *x.SalePrice
	}
	return 0
}

func (x *Content) GetFields() []*KeyValue {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Content) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Content) GetRelates() []string {
	if x != nil {
		return x.Relates
	}
	return nil
}

func (x *Content) GetCreated() int64 {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return 0
}

func (x *Content) GetUpdated() int64 {
	if x != nil && x.Updated != nil {
		return *x.Updated
	}
	return 0
}

func (x *Content) GetAttachmentUrls() []string {
	if x != nil {
		return x.AttachmentUrls
	}
	return nil
}

func (x *Content) GetRelatedIds() []string {
	if x != nil {
		return x.RelatedIds
	}
	return nil
}

func (x *Content) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

type Contents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx      *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Contents []*Content      `protobuf:"bytes,2,rep,name=contents" json:"contents,omitempty"`
	Anchor   *string         `protobuf:"bytes,3,opt,name=anchor" json:"anchor,omitempty"`
	Total    *int64          `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
}

func (x *Contents) Reset() {
	*x = Contents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contents) ProtoMessage() {}

func (x *Contents) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contents.ProtoReflect.Descriptor instead.
func (*Contents) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{5}
}

func (x *Contents) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *Contents) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *Contents) GetAnchor() string {
	if x != nil && x.Anchor != nil {
		return *x.Anchor
	}
	return ""
}

func (x *Contents) GetTotal() int64 {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return 0
}

type ContentUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId *string `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Par       *int32  `protobuf:"varint,3,opt,name=par" json:"par,omitempty"`
	Url       *string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Err       *string `protobuf:"bytes,5,opt,name=err" json:"err,omitempty"`
}

func (x *ContentUrl) Reset() {
	*x = ContentUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentUrl) ProtoMessage() {}

func (x *ContentUrl) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentUrl.ProtoReflect.Descriptor instead.
func (*ContentUrl) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{6}
}

func (x *ContentUrl) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *ContentUrl) GetPar() int32 {
	if x != nil && x.Par != nil {
		return *x.Par
	}
	return 0
}

func (x *ContentUrl) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *ContentUrl) GetErr() string {
	if x != nil && x.Err != nil {
		return *x.Err
	}
	return ""
}

type ContentsByUrlTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId *string  `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Url       *string  `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Ids       []string `protobuf:"bytes,3,rep,name=ids" json:"ids,omitempty"`
}

func (x *ContentsByUrlTable) Reset() {
	*x = ContentsByUrlTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentsByUrlTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentsByUrlTable) ProtoMessage() {}

func (x *ContentsByUrlTable) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentsByUrlTable.ProtoReflect.Descriptor instead.
func (*ContentsByUrlTable) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{7}
}

func (x *ContentsByUrlTable) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *ContentsByUrlTable) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *ContentsByUrlTable) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_content_proto protoreflect.FileDescriptor

var file_content_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x73, 0x0a, 0x0b, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x9e, 0x03, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6b,
	0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x69, 0x66, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x82, 0x05, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x03, 0x63, 0x74, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x4e, 0x0a, 0x0c, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0c, 0x0a, 0x08, 0x69, 0x6e,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x5f,
	0x6f, 0x66, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x70, 0x72,
	0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x64, 0x10, 0x04, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x63, 0x68,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x61, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x70, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x57, 0x0a, 0x12, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x72, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x75, 0x62, 0x69, 0x7a, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74,
}

var (
	file_content_proto_rawDescOnce sync.Once
	file_content_proto_rawDescData = file_content_proto_rawDesc
)

func file_content_proto_rawDescGZIP() []byte {
	file_content_proto_rawDescOnce.Do(func() {
		file_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_proto_rawDescData)
	})
	return file_content_proto_rawDescData
}

var file_content_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_content_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_content_proto_goTypes = []interface{}{
	(Content_Availability)(0),    // 0: content.Content.Availability
	(*LookupRequest)(nil),        // 1: content.LookupRequest
	(*LinkRequest)(nil),          // 2: content.LinkRequest
	(*SearchContentRequest)(nil), // 3: content.SearchContentRequest
	(*KeyValue)(nil),             // 4: content.KeyValue
	(*Content)(nil),              // 5: content.Content
	(*Contents)(nil),             // 6: content.Contents
	(*ContentUrl)(nil),           // 7: content.ContentUrl
	(*ContentsByUrlTable)(nil),   // 8: content.ContentsByUrlTable
	(*common.Context)(nil),       // 9: common.Context
}
var file_content_proto_depIdxs = []int32{
	9, // 0: content.LookupRequest.ctx:type_name -> common.Context
	9, // 1: content.LinkRequest.ctx:type_name -> common.Context
	9, // 2: content.SearchContentRequest.ctx:type_name -> common.Context
	9, // 3: content.Content.ctx:type_name -> common.Context
	4, // 4: content.Content.fields:type_name -> content.KeyValue
	9, // 5: content.Contents.ctx:type_name -> common.Context
	5, // 6: content.Contents.contents:type_name -> content.Content
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_content_proto_init() }
func file_content_proto_init() {
	if File_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchContentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contents); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentUrl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentsByUrlTable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_content_proto_goTypes,
		DependencyIndexes: file_content_proto_depIdxs,
		EnumInfos:         file_content_proto_enumTypes,
		MessageInfos:      file_content_proto_msgTypes,
	}.Build()
	File_content_proto = out.File
	file_content_proto_rawDesc = nil
	file_content_proto_goTypes = nil
	file_content_proto_depIdxs = nil
}
