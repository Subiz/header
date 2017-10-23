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
	Ids
	ListRequest
	MergeRequest
	GreetingRequest
	CreateRequest
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event int32

const (
	Event_UserReadRequested     Event = 0
	Event_UserUpdateRequested   Event = 2
	Event_UserCreateRequested   Event = 3
	Event_UserReadBulkRequested Event = 4
)

var Event_name = map[int32]string{
	0: "UserReadRequested",
	2: "UserUpdateRequested",
	3: "UserCreateRequested",
	4: "UserReadBulkRequested",
}
var Event_value = map[string]int32{
	"UserReadRequested":     0,
	"UserUpdateRequested":   2,
	"UserCreateRequested":   3,
	"UserReadBulkRequested": 4,
}

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}
func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (x *Event) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_value, data, "Event")
	if err != nil {
		return err
	}
	*x = Event(value)
	return nil
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Id               *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	AccountId        *string         `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Fullname         *string         `protobuf:"bytes,5,opt,name=fullname" json:"fullname,omitempty"`
	Phones           []string        `protobuf:"bytes,7,rep,name=phones" json:"phones,omitempty"`
	Emails           []string        `protobuf:"bytes,10,rep,name=emails" json:"emails,omitempty"`
	Traces           []*Trace        `protobuf:"bytes,11,rep,name=traces" json:"traces,omitempty"`
	Alias            []string        `protobuf:"bytes,12,rep,name=alias" json:"alias,omitempty"`
	Devices          []*Device       `protobuf:"bytes,13,rep,name=devices" json:"devices,omitempty"`
	IsBan            *bool           `protobuf:"varint,14,opt,name=is_ban,json=isBan" json:"is_ban,omitempty"`
	AvatarUrl        *string         `protobuf:"bytes,15,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
	Segments         []string        `protobuf:"bytes,19,rep,name=segments" json:"segments,omitempty"`
	Labels           []string        `protobuf:"bytes,20,rep,name=labels" json:"labels,omitempty"`
	Unsubscribed     *bool           `protobuf:"varint,21,opt,name=unsubscribed" json:"unsubscribed,omitempty"`
	MarkedSpam       *bool           `protobuf:"varint,22,opt,name=marked_spam,json=markedSpam" json:"marked_spam,omitempty"`
	HardBounced      *bool           `protobuf:"varint,23,opt,name=hard_bounced,json=hardBounced" json:"hard_bounced,omitempty"`
	TotalSessions    *int32          `protobuf:"varint,24,opt,name=total_sessions,json=totalSessions" json:"total_sessions,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *User) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *User) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *User) GetFullname() string {
	if m != nil && m.Fullname != nil {
		return *m.Fullname
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

func (m *User) GetAlias() []string {
	if m != nil {
		return m.Alias
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
	if m != nil && m.IsBan != nil {
		return *m.IsBan
	}
	return false
}

func (m *User) GetAvatarUrl() string {
	if m != nil && m.AvatarUrl != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *User) GetSegments() []string {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *User) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *User) GetUnsubscribed() bool {
	if m != nil && m.Unsubscribed != nil {
		return *m.Unsubscribed
	}
	return false
}

func (m *User) GetMarkedSpam() bool {
	if m != nil && m.MarkedSpam != nil {
		return *m.MarkedSpam
	}
	return false
}

func (m *User) GetHardBounced() bool {
	if m != nil && m.HardBounced != nil {
		return *m.HardBounced
	}
	return false
}

func (m *User) GetTotalSessions() int32 {
	if m != nil && m.TotalSessions != nil {
		return *m.TotalSessions
	}
	return 0
}

type Users struct {
	Users            []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
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
	Id               *int32  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	UseragentId      *int32  `protobuf:"varint,4,opt,name=useragent_id,json=useragentId" json:"useragent_id,omitempty"`
	Useragent        *string `protobuf:"bytes,5,opt,name=useragent" json:"useragent,omitempty"`
	Resolution       *string `protobuf:"bytes,6,opt,name=resolution" json:"resolution,omitempty"`
	LanguageId       *int32  `protobuf:"varint,7,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	Language         *string `protobuf:"bytes,8,opt,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Device) GetUseragentId() int32 {
	if m != nil && m.UseragentId != nil {
		return *m.UseragentId
	}
	return 0
}

func (m *Device) GetUseragent() string {
	if m != nil && m.Useragent != nil {
		return *m.Useragent
	}
	return ""
}

func (m *Device) GetResolution() string {
	if m != nil && m.Resolution != nil {
		return *m.Resolution
	}
	return ""
}

func (m *Device) GetLanguageId() int32 {
	if m != nil && m.LanguageId != nil {
		return *m.LanguageId
	}
	return 0
}

func (m *Device) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

type Trace struct {
	Id               *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	PP               *string `protobuf:"bytes,4,opt,name=pP" json:"pP,omitempty"`
	LocationId       *int32  `protobuf:"varint,5,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	Location         *string `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Trace) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Trace) GetPP() string {
	if m != nil && m.PP != nil {
		return *m.PP
	}
	return ""
}

func (m *Trace) GetLocationId() int32 {
	if m != nil && m.LocationId != nil {
		return *m.LocationId
	}
	return 0
}

func (m *Trace) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

type Id struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Id               *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	AccountId        *string         `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Id) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Id) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Id) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

//
// message Empty {};
//
// service UserMgr {
// rpc Greeting(GreetingRequest) returns (Id) {}
//
// rpc Merge(MergeRequest) returns (Empty) {}
// rpc ReadBulk(Ids) returns (Users) {}
// rpc List(ListRequest) returns (Users) {}
// }
type Ids struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Ids              []string        `protobuf:"bytes,4,rep,name=ids" json:"ids,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Ids) Reset()                    { *m = Ids{} }
func (m *Ids) String() string            { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()               {}
func (*Ids) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Ids) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Ids) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Ids) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,5,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	StartId          *string         `protobuf:"bytes,6,opt,name=start_id,json=startId" json:"start_id,omitempty"`
	Limit            *int32          `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Keyword          *string         `protobuf:"bytes,4,opt,name=keyword" json:"keyword,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ListRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ListRequest) GetStartId() string {
	if m != nil && m.StartId != nil {
		return *m.StartId
	}
	return ""
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *ListRequest) GetKeyword() string {
	if m != nil && m.Keyword != nil {
		return *m.Keyword
	}
	return ""
}

type MergeRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id               *string         `protobuf:"bytes,5,opt,name=id" json:"id,omitempty"`
	RecentId         *string         `protobuf:"bytes,4,opt,name=recent_id,json=recentId" json:"recent_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MergeRequest) Reset()                    { *m = MergeRequest{} }
func (m *MergeRequest) String() string            { return proto.CompactTextString(m) }
func (*MergeRequest) ProtoMessage()               {}
func (*MergeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MergeRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *MergeRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *MergeRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *MergeRequest) GetRecentId() string {
	if m != nil && m.RecentId != nil {
		return *m.RecentId
	}
	return ""
}

type GreetingRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,5,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Useragent        *string         `protobuf:"bytes,6,opt,name=useragent" json:"useragent,omitempty"`
	UserId           *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Uuid             *string         `protobuf:"bytes,4,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GreetingRequest) Reset()                    { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()               {}
func (*GreetingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GreetingRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *GreetingRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *GreetingRequest) GetUseragent() string {
	if m != nil && m.Useragent != nil {
		return *m.Useragent
	}
	return ""
}

func (m *GreetingRequest) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *GreetingRequest) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

type CreateRequest struct {
	AccountId        *string `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ChallengeId      *string `protobuf:"bytes,3,opt,name=challenge_id,json=challengeId" json:"challenge_id,omitempty"`
	Answer           *string `protobuf:"bytes,4,opt,name=answer" json:"answer,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CreateRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *CreateRequest) GetChallengeId() string {
	if m != nil && m.ChallengeId != nil {
		return *m.ChallengeId
	}
	return ""
}

func (m *CreateRequest) GetAnswer() string {
	if m != nil && m.Answer != nil {
		return *m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Users)(nil), "user.Users")
	proto.RegisterType((*Device)(nil), "user.Device")
	proto.RegisterType((*Trace)(nil), "user.Trace")
	proto.RegisterType((*Id)(nil), "user.Id")
	proto.RegisterType((*Ids)(nil), "user.Ids")
	proto.RegisterType((*ListRequest)(nil), "user.ListRequest")
	proto.RegisterType((*MergeRequest)(nil), "user.MergeRequest")
	proto.RegisterType((*GreetingRequest)(nil), "user.GreetingRequest")
	proto.RegisterType((*CreateRequest)(nil), "user.CreateRequest")
	proto.RegisterEnum("user.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5d, 0x8f, 0xe3, 0x34,
	0x14, 0x25, 0x6d, 0xd3, 0x8f, 0x9b, 0xce, 0x07, 0xde, 0x9d, 0x1d, 0xef, 0xf2, 0xd5, 0x09, 0x02,
	0x15, 0x1e, 0x3a, 0xd2, 0xfe, 0x84, 0x59, 0x10, 0x8a, 0x04, 0xd2, 0x2a, 0xcb, 0x20, 0xf1, 0x54,
	0x39, 0xf1, 0xa5, 0x63, 0x35, 0x71, 0x8a, 0xed, 0xcc, 0x2e, 0x48, 0x48, 0xfc, 0x0e, 0x1e, 0xf8,
	0x19, 0xfc, 0x3e, 0xe4, 0x8f, 0xa4, 0xd3, 0xee, 0x0b, 0xac, 0xf6, 0xa5, 0xf5, 0x39, 0xd7, 0xf6,
	0xb9, 0xf7, 0xe8, 0x38, 0x70, 0xd6, 0x6a, 0x54, 0xd7, 0xf6, 0x67, 0xb5, 0x53, 0x8d, 0x69, 0xc8,
	0xc8, 0xae, 0x9f, 0xad, 0x0a, 0x61, 0x8a, 0xb6, 0xdc, 0xa2, 0x59, 0x35, 0x6a, 0x73, 0xad, 0xdb,
	0x42, 0xfc, 0x7e, 0x7d, 0x87, 0x8c, 0xa3, 0xba, 0x2e, 0x9b, 0xba, 0x6e, 0x64, 0xf8, 0xf3, 0xa7,
	0xd2, 0x3f, 0x47, 0x30, 0xba, 0xd5, 0xa8, 0xc8, 0x15, 0x0c, 0x4b, 0xf3, 0x86, 0x46, 0x8b, 0x68,
	0x99, 0x3c, 0x3f, 0x5b, 0x85, 0x4d, 0x2f, 0x1a, 0x69, 0xf0, 0x8d, 0xc9, 0x6d, 0x8d, 0x9c, 0xc2,
	0x40, 0x70, 0x3a, 0x5c, 0x44, 0xcb, 0x59, 0x3e, 0x10, 0x9c, 0x7c, 0x02, 0xc0, 0xca, 0xb2, 0x69,
	0xa5, 0x59, 0x0b, 0x4e, 0x47, 0x8e, 0x9f, 0x05, 0x26, 0xe3, 0xe4, 0x19, 0x4c, 0x7f, 0x69, 0xab,
	0x4a, 0xb2, 0x1a, 0x69, 0xec, 0x8a, 0x3d, 0x26, 0x4f, 0x60, 0xbc, 0xbb, 0x6b, 0x24, 0x6a, 0x3a,
	0x59, 0x0c, 0x97, 0xb3, 0x3c, 0x20, 0xcb, 0x63, 0xcd, 0x44, 0xa5, 0x29, 0x78, 0xde, 0x23, 0xf2,
	0x39, 0x8c, 0x8d, 0x62, 0x25, 0x6a, 0x9a, 0x2c, 0x86, 0xcb, 0xe4, 0x79, 0xb2, 0x72, 0x93, 0xff,
	0x68, 0xb9, 0x3c, 0x94, 0xc8, 0x63, 0x88, 0x59, 0x25, 0x98, 0xa6, 0x73, 0x77, 0xd6, 0x03, 0xf2,
	0x25, 0x4c, 0x38, 0xde, 0x0b, 0x7b, 0xf6, 0xc4, 0x9d, 0x9d, 0xfb, 0xb3, 0xdf, 0x38, 0x32, 0xef,
	0x8a, 0xe4, 0x02, 0xc6, 0x42, 0xaf, 0x0b, 0x26, 0xe9, 0xe9, 0x22, 0x5a, 0x4e, 0xf3, 0x58, 0xe8,
	0x1b, 0x26, 0xdd, 0x90, 0xf7, 0xcc, 0x30, 0xb5, 0x6e, 0x55, 0x45, 0xcf, 0xc2, 0x90, 0x8e, 0xb9,
	0x55, 0x95, 0x1d, 0x52, 0xe3, 0xa6, 0x46, 0x69, 0x34, 0x7d, 0xe4, 0x64, 0x7b, 0x6c, 0x87, 0xa9,
	0x58, 0x81, 0x95, 0xa6, 0x8f, 0xfd, 0x30, 0x1e, 0x91, 0x14, 0xe6, 0xad, 0xd4, 0x6d, 0xa1, 0x4b,
	0x25, 0x0a, 0xe4, 0xf4, 0xc2, 0xe9, 0x1d, 0x70, 0xe4, 0x33, 0x48, 0x6a, 0xa6, 0xb6, 0xc8, 0xd7,
	0x7a, 0xc7, 0x6a, 0xfa, 0xc4, 0x6d, 0x01, 0x4f, 0xbd, 0xda, 0xb1, 0x9a, 0x5c, 0xc1, 0xfc, 0x8e,
	0x29, 0xbe, 0x2e, 0x9a, 0x56, 0x96, 0xc8, 0xe9, 0xa5, 0xdb, 0x91, 0x58, 0xee, 0xc6, 0x53, 0xe4,
	0x0b, 0x38, 0x35, 0x8d, 0x61, 0xd5, 0x5a, 0xa3, 0xd6, 0xa2, 0x91, 0x9a, 0xd2, 0x45, 0xb4, 0x8c,
	0xf3, 0x13, 0xc7, 0xbe, 0x0a, 0x64, 0xfa, 0x15, 0xc4, 0x36, 0x01, 0x9a, 0x2c, 0x20, 0xb6, 0xce,
	0x68, 0x1a, 0x39, 0x9f, 0xc0, 0xfb, 0x64, 0x6b, 0xb9, 0x2f, 0xa4, 0xff, 0x44, 0x30, 0xf6, 0xbe,
	0x3d, 0x08, 0x43, 0xec, 0xc2, 0x70, 0x05, 0x73, 0xbb, 0x87, 0x6d, 0x70, 0x1f, 0x87, 0x38, 0x4f,
	0x7a, 0x2e, 0xe3, 0xe4, 0x63, 0x98, 0xf5, 0x30, 0x24, 0x62, 0x4f, 0x90, 0x4f, 0x01, 0x14, 0xea,
	0xa6, 0x6a, 0x8d, 0x68, 0x24, 0x1d, 0xbb, 0xf2, 0x03, 0xc6, 0x3a, 0x52, 0x31, 0xb9, 0x69, 0xd9,
	0x06, 0xed, 0xfd, 0x13, 0x77, 0x3f, 0x74, 0x94, 0xcf, 0x5b, 0x87, 0xe8, 0xd4, 0xe7, 0xad, 0xc3,
	0x29, 0x87, 0xd8, 0x65, 0xe5, 0xad, 0x0c, 0x9f, 0xc2, 0x60, 0xf7, 0x32, 0x64, 0x77, 0xb0, 0x7b,
	0xe9, 0x54, 0x9a, 0x92, 0x59, 0x45, 0xab, 0x12, 0x07, 0x95, 0x40, 0x05, 0x95, 0x80, 0x42, 0x93,
	0x3d, 0x4e, 0x7f, 0x82, 0x41, 0xc6, 0xdf, 0xff, 0x4b, 0x4a, 0x7f, 0x86, 0x61, 0xc6, 0xf5, 0x7f,
	0xb9, 0xf8, 0xf0, 0xa2, 0xe1, 0xf1, 0x93, 0x3c, 0x87, 0xa1, 0xe0, 0x9a, 0x8e, 0x5c, 0x1c, 0xed,
	0x32, 0xfd, 0x2b, 0x82, 0xe4, 0x7b, 0xa1, 0x4d, 0x8e, 0xbf, 0xb6, 0xa8, 0xcd, 0xff, 0xd7, 0x88,
	0x8f, 0x35, 0x9e, 0xc2, 0x54, 0x1b, 0xa6, 0x5c, 0xd1, 0x1b, 0x34, 0x71, 0x38, 0xe3, 0xf6, 0x81,
	0x56, 0xa2, 0x16, 0x26, 0xc4, 0xc6, 0x03, 0x42, 0x61, 0xb2, 0xc5, 0xdf, 0x5e, 0x37, 0xaa, 0x9b,
	0xbc, 0x83, 0xe9, 0x1f, 0x30, 0xff, 0x01, 0xd5, 0x06, 0xdf, 0xb9, 0xb9, 0xb7, 0x0c, 0xf0, 0xc6,
	0xc7, 0xbd, 0xf1, 0x1f, 0xc1, 0x4c, 0x61, 0x89, 0x0f, 0x7d, 0x9f, 0x7a, 0x22, 0xe3, 0xe9, 0xdf,
	0x11, 0x9c, 0x7d, 0xa7, 0x10, 0x8d, 0x90, 0x9b, 0xf7, 0xe7, 0xcf, 0xc1, 0x2b, 0x18, 0x1f, 0xbf,
	0x82, 0x4b, 0x98, 0x58, 0xb0, 0x6f, 0x7e, 0x6c, 0x61, 0xc6, 0x09, 0x81, 0x51, 0xdb, 0xf6, 0x4d,
	0xba, 0x75, 0x2a, 0xe0, 0xe4, 0x85, 0x42, 0x66, 0x7a, 0x83, 0x0e, 0xa5, 0x07, 0xc7, 0xd2, 0x57,
	0x30, 0x2f, 0xef, 0x58, 0x55, 0xa1, 0xf4, 0x6f, 0xc8, 0x2b, 0x24, 0x3d, 0x97, 0x71, 0xfb, 0xcd,
	0x62, 0x52, 0xbf, 0x46, 0x15, 0x84, 0x02, 0xfa, 0x7a, 0x0b, 0xf1, 0xb7, 0xf7, 0xb6, 0xc1, 0x0b,
	0xf8, 0xd0, 0x7d, 0x11, 0x90, 0xf1, 0xa0, 0x8a, 0xfc, 0xfc, 0x03, 0x72, 0x09, 0x8f, 0x2c, 0x7d,
	0xbb, 0xe3, 0xfb, 0x76, 0x90, 0x9f, 0x0f, 0xba, 0xc2, 0x41, 0x9f, 0xc8, 0xcf, 0x87, 0xe4, 0x29,
	0x5c, 0x74, 0x17, 0xdd, 0xb4, 0xd5, 0x76, 0x5f, 0x1a, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x35,
	0x3e, 0xe0, 0xb0, 0xdc, 0x06, 0x00, 0x00,
}
