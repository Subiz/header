// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: file.proto

package file

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

type FileHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx         *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId   string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name        string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type        string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Size        int64           `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Md5         string          `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	Description string          `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FileHeader) Reset() {
	*x = FileHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHeader) ProtoMessage() {}

func (x *FileHeader) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHeader.ProtoReflect.Descriptor instead.
func (*FileHeader) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileHeader) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *FileHeader) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *FileHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileHeader) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FileHeader) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileHeader) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *FileHeader) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PresignResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Url       string          `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Id        string          `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	SignedUrl string          `protobuf:"bytes,5,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
}

func (x *PresignResult) Reset() {
	*x = PresignResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresignResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresignResult) ProtoMessage() {}

func (x *PresignResult) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresignResult.ProtoReflect.Descriptor instead.
func (*PresignResult) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *PresignResult) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *PresignResult) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *PresignResult) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PresignResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PresignResult) GetSignedUrl() string {
	if x != nil {
		return x.SignedUrl
	}
	return ""
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id        string          `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *FileRequest) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *FileRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *FileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx         *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId   string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name        string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type        string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Size        int64           `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Md5         string          `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	Description string          `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	Created     int64           `protobuf:"varint,7,opt,name=created,proto3" json:"created,omitempty"`
	Url         string          `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Creator     string          `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
	Id          string          `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *File) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *File) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *File) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *File) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *File) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id         string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Uploaded   int64  `protobuf:"varint,4,opt,name=uploaded,proto3" json:"uploaded,omitempty"`
	Name       string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Size       int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	FullUrl    string `protobuf:"bytes,7,opt,name=full_url,json=fullUrl,proto3" json:"full_url,omitempty"`
	MimeType   string `protobuf:"bytes,8,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Downloaded int64  `protobuf:"varint,10,opt,name=downloaded,proto3" json:"downloaded,omitempty"`
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{4}
}

func (x *FileMeta) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *FileMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileMeta) GetUploaded() int64 {
	if x != nil {
		return x.Uploaded
	}
	return 0
}

func (x *FileMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileMeta) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileMeta) GetFullUrl() string {
	if x != nil {
		return x.FullUrl
	}
	return ""
}

func (x *FileMeta) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *FileMeta) GetDownloaded() int64 {
	if x != nil {
		return x.Downloaded
	}
	return 0
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbe, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63,
	0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x92, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x5f, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8e, 0x02, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03,
	0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd5, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6c, 0x6c,
	0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x75, 0x62, 0x69, 0x7a, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_file_proto_goTypes = []interface{}{
	(*FileHeader)(nil),     // 0: file.FileHeader
	(*PresignResult)(nil),  // 1: file.PresignResult
	(*FileRequest)(nil),    // 2: file.FileRequest
	(*File)(nil),           // 3: file.File
	(*FileMeta)(nil),       // 4: file.FileMeta
	(*common.Context)(nil), // 5: common.Context
}
var file_file_proto_depIdxs = []int32{
	5, // 0: file.FileHeader.ctx:type_name -> common.Context
	5, // 1: file.PresignResult.ctx:type_name -> common.Context
	5, // 2: file.FileRequest.ctx:type_name -> common.Context
	5, // 3: file.File.ctx:type_name -> common.Context
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHeader); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresignResult); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMeta); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
