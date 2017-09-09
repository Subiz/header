// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	User
	Users
	Device
	Trace
	Id
	Empty
	Ids
	ListRequest
	MergeRequest
	GreetingRequest
*/
package user

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

type User struct {
	Id        string    `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Alias     []string  `protobuf:"bytes,2,rep,name=Alias" json:"Alias,omitempty"`
	AccountId string    `protobuf:"bytes,3,opt,name=AccountId" json:"AccountId,omitempty"`
	FirstName string    `protobuf:"bytes,4,opt,name=FirstName" json:"FirstName,omitempty"`
	LastName  string    `protobuf:"bytes,5,opt,name=LastName" json:"LastName,omitempty"`
	Phones    []string  `protobuf:"bytes,6,rep,name=Phones" json:"Phones,omitempty"`
	Emails    []string  `protobuf:"bytes,7,rep,name=Emails" json:"Emails,omitempty"`
	Traces    []*Trace  `protobuf:"bytes,8,rep,name=Traces" json:"Traces,omitempty"`
	Devices   []*Device `protobuf:"bytes,9,rep,name=Devices" json:"Devices,omitempty"`
	IsBan     bool      `protobuf:"varint,10,opt,name=IsBan" json:"IsBan,omitempty"`
	AvatarUrl string    `protobuf:"bytes,11,opt,name=AvatarUrl" json:"AvatarUrl,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *User) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetPhones() []string {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *User) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *User) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

func (m *User) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *User) GetIsBan() bool {
	if m != nil {
		return m.IsBan
	}
	return false
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

type Users struct {
	Users []*User `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *Users) Reset()                    { *m = Users{} }
func (m *Users) String() string            { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()               {}
func (*Users) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Device struct {
	Id              int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	UserAgentZipped int32  `protobuf:"varint,2,opt,name=UserAgentZipped" json:"UserAgentZipped,omitempty"`
	Resolution      string `protobuf:"bytes,3,opt,name=Resolution" json:"Resolution,omitempty"`
	LanguageZipped  int32  `protobuf:"varint,4,opt,name=LanguageZipped" json:"LanguageZipped,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Device) GetUserAgentZipped() int32 {
	if m != nil {
		return m.UserAgentZipped
	}
	return 0
}

func (m *Device) GetResolution() string {
	if m != nil {
		return m.Resolution
	}
	return ""
}

func (m *Device) GetLanguageZipped() int32 {
	if m != nil {
		return m.LanguageZipped
	}
	return 0
}

type Trace struct {
	Id             string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	IP             string `protobuf:"bytes,2,opt,name=IP" json:"IP,omitempty"`
	LocationZipped int32  `protobuf:"varint,3,opt,name=LocationZipped" json:"LocationZipped,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Trace) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trace) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Trace) GetLocationZipped() int32 {
	if m != nil {
		return m.LocationZipped
	}
	return 0
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

type Ids struct {
	Ids []string `protobuf:"bytes,1,rep,name=Ids" json:"Ids,omitempty"`
}

func (m *Ids) Reset()                    { *m = Ids{} }
func (m *Ids) String() string            { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()               {}
func (*Ids) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Ids) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId" json:"AccountId,omitempty"`
	StartId   string `protobuf:"bytes,2,opt,name=StartId" json:"StartId,omitempty"`
	Limit     int32  `protobuf:"varint,3,opt,name=Limit" json:"Limit,omitempty"`
	Keyword   string `protobuf:"bytes,4,opt,name=Keyword" json:"Keyword,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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

type MergeRequest struct {
	FormerUserId string `protobuf:"bytes,1,opt,name=FormerUserId" json:"FormerUserId,omitempty"`
	RecentUserId string `protobuf:"bytes,2,opt,name=RecentUserId" json:"RecentUserId,omitempty"`
}

func (m *MergeRequest) Reset()                    { *m = MergeRequest{} }
func (m *MergeRequest) String() string            { return proto.CompactTextString(m) }
func (*MergeRequest) ProtoMessage()               {}
func (*MergeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MergeRequest) GetFormerUserId() string {
	if m != nil {
		return m.FormerUserId
	}
	return ""
}

func (m *MergeRequest) GetRecentUserId() string {
	if m != nil {
		return m.RecentUserId
	}
	return ""
}

type GreetingRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId" json:"AccountId,omitempty"`
	UserAgent string `protobuf:"bytes,2,opt,name=UserAgent" json:"UserAgent,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=UserId" json:"UserId,omitempty"`
	UUID      string `protobuf:"bytes,4,opt,name=UUID" json:"UUID,omitempty"`
}

func (m *GreetingRequest) Reset()                    { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()               {}
func (*GreetingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GreetingRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *GreetingRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *GreetingRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GreetingRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Users)(nil), "user.Users")
	proto.RegisterType((*Device)(nil), "user.Device")
	proto.RegisterType((*Trace)(nil), "user.Trace")
	proto.RegisterType((*Id)(nil), "user.Id")
	proto.RegisterType((*Empty)(nil), "user.Empty")
	proto.RegisterType((*Ids)(nil), "user.Ids")
	proto.RegisterType((*ListRequest)(nil), "user.ListRequest")
	proto.RegisterType((*MergeRequest)(nil), "user.MergeRequest")
	proto.RegisterType((*GreetingRequest)(nil), "user.GreetingRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserMgr service

type UserMgrClient interface {
	Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*Id, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	Ban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	Unban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	Merge(ctx context.Context, in *MergeRequest, opts ...grpc.CallOption) (*Empty, error)
	ReadBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Users, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Users, error)
}

type userMgrClient struct {
	cc *grpc.ClientConn
}

func NewUserMgrClient(cc *grpc.ClientConn) UserMgrClient {
	return &userMgrClient{cc}
}

func (c *userMgrClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/user.UserMgr/Greeting", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Ban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Ban", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Unban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Unban", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Merge(ctx context.Context, in *MergeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Merge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) ReadBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := grpc.Invoke(ctx, "/user.UserMgr/ReadBulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := grpc.Invoke(ctx, "/user.UserMgr/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserMgr service

type UserMgrServer interface {
	Greeting(context.Context, *GreetingRequest) (*Id, error)
	Update(context.Context, *User) (*Empty, error)
	Ban(context.Context, *Id) (*Empty, error)
	Unban(context.Context, *Id) (*Empty, error)
	Merge(context.Context, *MergeRequest) (*Empty, error)
	ReadBulk(context.Context, *Ids) (*Users, error)
	List(context.Context, *ListRequest) (*Users, error)
}

func RegisterUserMgrServer(s *grpc.Server, srv UserMgrServer) {
	s.RegisterService(&_UserMgr_serviceDesc, srv)
}

func _UserMgr_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Greeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Ban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Ban(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Unban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Unban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Unban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Unban(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Merge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Merge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Merge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Merge(ctx, req.(*MergeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_ReadBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).ReadBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/ReadBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).ReadBulk(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserMgr",
	HandlerType: (*UserMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _UserMgr_Greeting_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserMgr_Update_Handler,
		},
		{
			MethodName: "Ban",
			Handler:    _UserMgr_Ban_Handler,
		},
		{
			MethodName: "Unban",
			Handler:    _UserMgr_Unban_Handler,
		},
		{
			MethodName: "Merge",
			Handler:    _UserMgr_Merge_Handler,
		},
		{
			MethodName: "ReadBulk",
			Handler:    _UserMgr_ReadBulk_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserMgr_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x3e,
	0x14, 0x5f, 0xf3, 0xd1, 0x26, 0xa7, 0xd3, 0xf6, 0xff, 0x5b, 0x03, 0xac, 0x69, 0x9a, 0x2a, 0x83,
	0xa6, 0x80, 0xc4, 0x90, 0xc6, 0x13, 0x6c, 0xda, 0x86, 0x22, 0x3a, 0x98, 0x02, 0xe5, 0x82, 0x3b,
	0xaf, 0xb1, 0x42, 0x44, 0x9b, 0x14, 0xdb, 0x19, 0xea, 0x2b, 0xf0, 0x0a, 0x3c, 0x07, 0xef, 0x87,
	0x8e, 0xed, 0x74, 0x69, 0x90, 0x10, 0x37, 0x9b, 0x7f, 0x1f, 0x39, 0xc7, 0xe7, 0xc3, 0x85, 0xfd,
	0x46, 0x09, 0xf9, 0x0a, 0xff, 0x9c, 0xae, 0x64, 0xad, 0x6b, 0x12, 0xe0, 0x99, 0xfd, 0xf2, 0x20,
	0x98, 0x29, 0x21, 0xc9, 0x1e, 0x78, 0x69, 0x4e, 0x07, 0x93, 0x41, 0x12, 0x67, 0x5e, 0x9a, 0x93,
	0x03, 0x08, 0xcf, 0x17, 0x25, 0x57, 0xd4, 0x9b, 0xf8, 0x49, 0x9c, 0x59, 0x40, 0x8e, 0x20, 0x3e,
	0x9f, 0xcf, 0xeb, 0xa6, 0xd2, 0x69, 0x4e, 0x7d, 0x63, 0x7e, 0x20, 0x50, 0xbd, 0x2e, 0xa5, 0xd2,
	0xef, 0xf8, 0x52, 0xd0, 0xc0, 0xaa, 0x1b, 0x82, 0x1c, 0x42, 0x34, 0xe5, 0x4e, 0x0c, 0x8d, 0xb8,
	0xc1, 0xe4, 0x31, 0x0c, 0x6f, 0xbf, 0xd4, 0x95, 0x50, 0x74, 0x68, 0xd2, 0x39, 0x84, 0xfc, 0xd5,
	0x92, 0x97, 0x0b, 0x45, 0x47, 0x96, 0xb7, 0x88, 0x3c, 0x85, 0xe1, 0x47, 0xc9, 0xe7, 0x42, 0xd1,
	0x68, 0xe2, 0x27, 0xe3, 0xb3, 0xf1, 0xa9, 0xa9, 0xcc, 0x70, 0x99, 0x93, 0xc8, 0x09, 0x8c, 0x2e,
	0xc5, 0x7d, 0x89, 0xae, 0xd8, 0xb8, 0x76, 0xad, 0xcb, 0x92, 0x59, 0x2b, 0x62, 0xa9, 0xa9, 0xba,
	0xe0, 0x15, 0x85, 0xc9, 0x20, 0x89, 0x32, 0x0b, 0x4c, 0xa9, 0xf7, 0x5c, 0x73, 0x39, 0x93, 0x0b,
	0x3a, 0x76, 0xa5, 0xb6, 0x04, 0x7b, 0x0e, 0x21, 0xb6, 0x4d, 0x91, 0x89, 0x3b, 0xd0, 0x81, 0x49,
	0x01, 0x36, 0x05, 0x52, 0x99, 0x15, 0xd8, 0x8f, 0x01, 0x0c, 0x6d, 0xaa, 0x4e, 0x93, 0x43, 0xd3,
	0xe4, 0x04, 0xf6, 0xd1, 0x73, 0x5e, 0x88, 0x4a, 0x7f, 0x2e, 0x57, 0x2b, 0x91, 0x53, 0xcf, 0x88,
	0x7d, 0x9a, 0x1c, 0x03, 0x64, 0x42, 0xd5, 0x8b, 0x46, 0x97, 0x75, 0xe5, 0x3a, 0xdf, 0x61, 0xc8,
	0x09, 0xec, 0x4d, 0x79, 0x55, 0x34, 0xbc, 0x10, 0x2e, 0x50, 0x60, 0x02, 0xf5, 0x58, 0xf6, 0x1e,
	0x42, 0xd3, 0x9d, 0x3f, 0xe6, 0x8d, 0xf8, 0xd6, 0x64, 0x47, 0x7c, 0x6b, 0x02, 0xd6, 0x73, 0x8e,
	0xc1, 0x5d, 0x40, 0xdf, 0x05, 0xdc, 0x62, 0xd9, 0x01, 0xb4, 0x5f, 0x77, 0xa2, 0xb1, 0x11, 0x84,
	0x57, 0xcb, 0x95, 0x5e, 0xb3, 0x27, 0xe0, 0xa7, 0xb9, 0x22, 0xff, 0x99, 0x7f, 0xa6, 0x47, 0x71,
	0x86, 0x47, 0xd6, 0xc0, 0x78, 0x5a, 0x2a, 0x9d, 0x89, 0x6f, 0x8d, 0x50, 0x7a, 0x7b, 0xb1, 0x06,
	0xfd, 0xc5, 0xa2, 0x30, 0xfa, 0xa0, 0xb9, 0x44, 0xcd, 0xde, 0xb0, 0x85, 0x38, 0xbb, 0x69, 0xb9,
	0x2c, 0xb5, 0xbb, 0x9d, 0x05, 0xe8, 0x7f, 0x2b, 0xd6, 0xdf, 0x6b, 0x99, 0xbb, 0x35, 0x6c, 0x21,
	0xfb, 0x04, 0xbb, 0x37, 0x42, 0x16, 0xa2, 0xcd, 0xcb, 0x60, 0xf7, 0xba, 0x96, 0x4b, 0x21, 0xb1,
	0xe1, 0x9b, 0xd4, 0x5b, 0x1c, 0x7a, 0x32, 0x31, 0x17, 0x95, 0x76, 0x1e, 0x7b, 0x85, 0x2d, 0x8e,
	0xad, 0x61, 0xff, 0x8d, 0x14, 0x42, 0x97, 0x55, 0xf1, 0x6f, 0x25, 0x1d, 0x41, 0xbc, 0x99, 0xb1,
	0x8b, 0xf8, 0x40, 0xe0, 0xde, 0xbb, 0x64, 0x76, 0xd4, 0x0e, 0x11, 0x02, 0xc1, 0x6c, 0x96, 0x5e,
	0xba, 0xaa, 0xcc, 0xf9, 0xec, 0xa7, 0x07, 0x23, 0x94, 0x6f, 0x0a, 0x49, 0x5e, 0x42, 0xd4, 0x5e,
	0x83, 0x3c, 0xb2, 0xab, 0xd8, 0xbb, 0xd6, 0x61, 0x64, 0xe9, 0x34, 0x67, 0x3b, 0xf8, 0x8c, 0x66,
	0xab, 0x9c, 0x6b, 0x41, 0x3a, 0x7b, 0x7b, 0xe8, 0x1e, 0x93, 0x1d, 0xe0, 0x0e, 0x39, 0x06, 0x1f,
	0xdf, 0xc3, 0xe6, 0xbb, 0xbe, 0x8e, 0x2f, 0xa0, 0xba, 0xfb, 0x9b, 0xe3, 0x05, 0x84, 0xa6, 0xe9,
	0x84, 0x58, 0xbe, 0x3b, 0x81, 0xbe, 0xf7, 0x19, 0x44, 0x99, 0xe0, 0xf9, 0x45, 0xb3, 0xf8, 0x4a,
	0xe2, 0x36, 0xa0, 0x6a, 0x5d, 0xf6, 0x45, 0xed, 0x90, 0x04, 0x02, 0xdc, 0x1e, 0xf2, 0xbf, 0xa5,
	0x3b, 0x9b, 0xd4, 0x73, 0xde, 0x0d, 0xcd, 0xaf, 0xdd, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0x23, 0xce, 0x63, 0x00, 0x05, 0x00, 0x00,
}
