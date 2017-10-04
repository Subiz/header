// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push/push.proto

/*
Package push is a generated protocol buffer package.

It is generated from these files:
	push/push.proto

It has these top-level messages:
	Token
	TokenRemoveRequest
	PushRequest
	Result
	Id
	PushConfiguration
*/
package push

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

type DeviceType int32

const (
	DeviceType_NONE    DeviceType = 0
	DeviceType_BROWSER DeviceType = 1
	DeviceType_MOBILE  DeviceType = 2
)

var DeviceType_name = map[int32]string{
	0: "NONE",
	1: "BROWSER",
	2: "MOBILE",
}
var DeviceType_value = map[string]int32{
	"NONE":    0,
	"BROWSER": 1,
	"MOBILE":  2,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event int32

const (
	Event_PushTokenRegisterRequested Event = 0
	Event_PushTokenRemoveRequested   Event = 2
	Event_PushRequested              Event = 3
	Event_PushConfigUpdateRequested  Event = 4
	Event_PushConfigReadRequested    Event = 5
)

var Event_name = map[int32]string{
	0: "PushTokenRegisterRequested",
	2: "PushTokenRemoveRequested",
	3: "PushRequested",
	4: "PushConfigUpdateRequested",
	5: "PushConfigReadRequested",
}
var Event_value = map[string]int32{
	"PushTokenRegisterRequested": 0,
	"PushTokenRemoveRequested":   2,
	"PushRequested":              3,
	"PushConfigUpdateRequested":  4,
	"PushConfigReadRequested":    5,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NotiType int32

const (
	NotiType_SystemMaintainance NotiType = 0
	NotiType_RuleUpdate         NotiType = 3
	NotiType_AccountUpdate      NotiType = 4
	NotiType_DirectMessage      NotiType = 5
	NotiType_News               NotiType = 6
)

var NotiType_name = map[int32]string{
	0: "SystemMaintainance",
	3: "RuleUpdate",
	4: "AccountUpdate",
	5: "DirectMessage",
	6: "News",
}
var NotiType_value = map[string]int32{
	"SystemMaintainance": 0,
	"RuleUpdate":         3,
	"AccountUpdate":      4,
	"DirectMessage":      5,
	"News":               6,
}

func (x NotiType) String() string {
	return proto.EnumName(NotiType_name, int32(x))
}
func (NotiType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Token struct {
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	// string ExpiredIn = 4;
	DeviceType DeviceType `protobuf:"varint,5,opt,name=device_type,json=deviceType,enum=push.DeviceType" json:"device_type,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Token) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_NONE
}

type TokenRemoveRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *TokenRemoveRequest) Reset()                    { *m = TokenRemoveRequest{} }
func (m *TokenRemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRemoveRequest) ProtoMessage()               {}
func (*TokenRemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TokenRemoveRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *TokenRemoveRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TokenRemoveRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PushRequest struct {
	AccountId   string     `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserIds     []string   `protobuf:"bytes,2,rep,name=user_ids,json=userIds" json:"user_ids,omitempty"`
	DeviceType  DeviceType `protobuf:"varint,3,opt,name=device_type,json=deviceType,enum=push.DeviceType" json:"device_type,omitempty"`
	Title       string     `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Body        string     `protobuf:"bytes,5,opt,name=body" json:"body,omitempty"`
	Icon        string     `protobuf:"bytes,6,opt,name=icon" json:"icon,omitempty"`
	ClickAction string     `protobuf:"bytes,7,opt,name=click_action,json=clickAction" json:"click_action,omitempty"`
}

func (m *PushRequest) Reset()                    { *m = PushRequest{} }
func (m *PushRequest) String() string            { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()               {}
func (*PushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *PushRequest) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *PushRequest) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_NONE
}

func (m *PushRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PushRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *PushRequest) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *PushRequest) GetClickAction() string {
	if m != nil {
		return m.ClickAction
	}
	return ""
}

type Result struct {
	AccountId string   `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserIds   []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds" json:"user_ids,omitempty"`
	Success   []bool   `protobuf:"varint,3,rep,packed,name=success" json:"success,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Result) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Result) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *Result) GetSuccess() []bool {
	if m != nil {
		return m.Success
	}
	return nil
}

type Id struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id        string          `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
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

func (m *Id) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PushConfiguration struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId    string          `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Email     string          `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	// optional bool NotificationSound = 1;
	EmailNotis   []NotiType `protobuf:"varint,6,rep,packed,name=email_notis,json=emailNotis,enum=push.NotiType" json:"email_notis,omitempty"`
	DesktopNotis []NotiType `protobuf:"varint,7,rep,packed,name=desktop_notis,json=desktopNotis,enum=push.NotiType" json:"desktop_notis,omitempty"`
	MobileNotis  []NotiType `protobuf:"varint,9,rep,packed,name=mobile_notis,json=mobileNotis,enum=push.NotiType" json:"mobile_notis,omitempty"`
	// MobilePushDelayTiming is the seconds subiz delay sending
	// notification to mobile after notifiy user in desktop device.
	// If user sees the notification during this time, the push will
	// be cancelled.
	MobilePushDelaySeconds int32 `protobuf:"varint,10,opt,name=mobile_push_delay_seconds,json=mobilePushDelaySeconds" json:"mobile_push_delay_seconds,omitempty"`
	// EmailThreshold is a duration (in minute) in which no more
	// than 1 email is being sent
	EmailThreshold int32 `protobuf:"varint,11,opt,name=email_threshold,json=emailThreshold" json:"email_threshold,omitempty"`
}

func (m *PushConfiguration) Reset()                    { *m = PushConfiguration{} }
func (m *PushConfiguration) String() string            { return proto.CompactTextString(m) }
func (*PushConfiguration) ProtoMessage()               {}
func (*PushConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PushConfiguration) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *PushConfiguration) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *PushConfiguration) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PushConfiguration) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PushConfiguration) GetEmailNotis() []NotiType {
	if m != nil {
		return m.EmailNotis
	}
	return nil
}

func (m *PushConfiguration) GetDesktopNotis() []NotiType {
	if m != nil {
		return m.DesktopNotis
	}
	return nil
}

func (m *PushConfiguration) GetMobileNotis() []NotiType {
	if m != nil {
		return m.MobileNotis
	}
	return nil
}

func (m *PushConfiguration) GetMobilePushDelaySeconds() int32 {
	if m != nil {
		return m.MobilePushDelaySeconds
	}
	return 0
}

func (m *PushConfiguration) GetEmailThreshold() int32 {
	if m != nil {
		return m.EmailThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*Token)(nil), "push.Token")
	proto.RegisterType((*TokenRemoveRequest)(nil), "push.TokenRemoveRequest")
	proto.RegisterType((*PushRequest)(nil), "push.PushRequest")
	proto.RegisterType((*Result)(nil), "push.Result")
	proto.RegisterType((*Id)(nil), "push.Id")
	proto.RegisterType((*PushConfiguration)(nil), "push.PushConfiguration")
	proto.RegisterEnum("push.DeviceType", DeviceType_name, DeviceType_value)
	proto.RegisterEnum("push.Event", Event_name, Event_value)
	proto.RegisterEnum("push.NotiType", NotiType_name, NotiType_value)
}

func init() { proto.RegisterFile("push/push.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x8e, 0x24, 0xff, 0x8e, 0x12, 0x47, 0x21, 0x16, 0x89, 0x92, 0xdd, 0x2c, 0x1c, 0x5f, 0xd6,
	0xc8, 0xc1, 0x46, 0x92, 0xd3, 0x1e, 0xf3, 0xe3, 0x83, 0x81, 0x26, 0x29, 0x94, 0xb4, 0xbd, 0x14,
	0x30, 0x24, 0x72, 0x6a, 0x13, 0x96, 0x45, 0x57, 0xa4, 0xd2, 0xb8, 0xb7, 0xde, 0xfa, 0x00, 0x7d,
	0xbb, 0xbe, 0x4c, 0x41, 0x52, 0xa9, 0x0d, 0x23, 0x40, 0x0b, 0x34, 0x17, 0x7b, 0xe6, 0xfb, 0xe6,
	0x1b, 0xce, 0x0c, 0x87, 0x82, 0xed, 0x79, 0x21, 0x27, 0x7d, 0xfd, 0xd3, 0x9b, 0xe7, 0x42, 0x09,
	0x52, 0xd1, 0xf6, 0x41, 0x2f, 0xe1, 0x2a, 0x29, 0xe8, 0x14, 0x55, 0x4f, 0xe4, 0xe3, 0xbe, 0x2c,
	0x12, 0xfe, 0xb9, 0x3f, 0xc1, 0x98, 0x61, 0xde, 0xa7, 0x62, 0x36, 0x13, 0x59, 0xf9, 0x67, 0x55,
	0x9d, 0xaf, 0x0e, 0x54, 0xef, 0xc5, 0x14, 0x33, 0x72, 0x08, 0x10, 0x53, 0x2a, 0x8a, 0x4c, 0x8d,
	0x38, 0x0b, 0x9d, 0xb6, 0xd3, 0x6d, 0x46, 0xcd, 0x12, 0x19, 0x32, 0xb2, 0x07, 0xf5, 0x42, 0x62,
	0xae, 0x39, 0xd7, 0x70, 0x35, 0xed, 0x0e, 0x19, 0xf9, 0x0b, 0xaa, 0x4a, 0x27, 0x08, 0x3d, 0x03,
	0x5b, 0x87, 0x9c, 0x80, 0xcf, 0xf0, 0x81, 0x53, 0x1c, 0xa9, 0xc5, 0x1c, 0xc3, 0x6a, 0xdb, 0xe9,
	0xb6, 0x4e, 0x83, 0x9e, 0xa9, 0xf7, 0xca, 0x10, 0xf7, 0x8b, 0x39, 0x46, 0xc0, 0x7e, 0xda, 0x9d,
	0x04, 0x88, 0xa9, 0x24, 0xc2, 0x99, 0x78, 0xc0, 0x08, 0x3f, 0x16, 0x28, 0xd5, 0xcb, 0x96, 0xd5,
	0xf9, 0xee, 0x80, 0xff, 0xba, 0x90, 0x93, 0xdf, 0xcc, 0xbe, 0x0f, 0x8d, 0x32, 0xbb, 0x0c, 0xdd,
	0xb6, 0xd7, 0x6d, 0x46, 0x75, 0x9b, 0x5e, 0xae, 0x37, 0xe8, 0xfd, 0xba, 0x41, 0x53, 0x12, 0x57,
	0x29, 0x86, 0x95, 0xb2, 0x24, 0xed, 0x10, 0x02, 0x95, 0x44, 0xb0, 0x85, 0x19, 0x51, 0x33, 0x32,
	0xb6, 0xc6, 0x38, 0x15, 0x59, 0x58, 0xb3, 0x98, 0xb6, 0xc9, 0x11, 0x6c, 0xd2, 0x94, 0xd3, 0xe9,
	0x28, 0xa6, 0x8a, 0x8b, 0x2c, 0xac, 0x1b, 0xce, 0x37, 0xd8, 0xb9, 0x81, 0x3a, 0xef, 0xa1, 0x16,
	0xa1, 0x2c, 0xd2, 0x3f, 0xe9, 0x2b, 0x84, 0xba, 0x2c, 0x28, 0x45, 0x29, 0x43, 0xaf, 0xed, 0x75,
	0x1b, 0xd1, 0x93, 0xdb, 0x79, 0x0b, 0xee, 0x90, 0x91, 0x23, 0xf0, 0xa8, 0x7a, 0x34, 0x29, 0xfd,
	0xd3, 0xed, 0x5e, 0xb9, 0x4c, 0x97, 0x22, 0x53, 0xf8, 0xa8, 0x22, 0xcd, 0xad, 0x1d, 0xee, 0xae,
	0x1f, 0xde, 0x02, 0x97, 0xb3, 0xf2, 0x5a, 0x5c, 0xce, 0x3a, 0x5f, 0x3c, 0xd8, 0xd1, 0x77, 0x72,
	0x29, 0xb2, 0x0f, 0x7c, 0x5c, 0xe4, 0xb1, 0xee, 0xe5, 0x05, 0xce, 0x59, 0x59, 0x0d, 0x6f, 0x7d,
	0x35, 0x70, 0x16, 0xf3, 0xb4, 0x1c, 0xb9, 0x75, 0x48, 0x1f, 0x7c, 0x63, 0x8c, 0x32, 0xa1, 0xb8,
	0x0c, 0x6b, 0x6d, 0xaf, 0xdb, 0x3a, 0x6d, 0xd9, 0x0b, 0xbd, 0x11, 0x8a, 0xdb, 0xeb, 0x34, 0x21,
	0xda, 0x95, 0xe4, 0x0c, 0xb6, 0x18, 0xca, 0xa9, 0x12, 0xf3, 0x52, 0x52, 0x7f, 0x56, 0xb2, 0x59,
	0x06, 0x59, 0xd1, 0x09, 0x6c, 0xce, 0x44, 0xc2, 0x53, 0x2c, 0x35, 0xcd, 0x67, 0x35, 0xbe, 0x8d,
	0xb1, 0x92, 0xff, 0x61, 0xbf, 0x94, 0xe8, 0xa0, 0x11, 0xc3, 0x34, 0x5e, 0x8c, 0x24, 0x52, 0x91,
	0x31, 0x19, 0x42, 0xdb, 0xe9, 0x56, 0xa3, 0x5d, 0x1b, 0xa0, 0xa7, 0x78, 0xa5, 0xe9, 0x3b, 0xcb,
	0x92, 0xff, 0x60, 0xdb, 0xf6, 0xa4, 0x26, 0x39, 0xca, 0x89, 0x48, 0x59, 0xe8, 0x1b, 0x41, 0xcb,
	0xc0, 0xf7, 0x4f, 0xe8, 0x71, 0x1f, 0x60, 0xb9, 0xb4, 0xa4, 0x01, 0x95, 0x9b, 0xdb, 0x9b, 0x41,
	0xb0, 0x41, 0x7c, 0xa8, 0x5f, 0x44, 0xb7, 0xef, 0xee, 0x06, 0x51, 0xe0, 0x10, 0x80, 0xda, 0xf5,
	0xed, 0xc5, 0xf0, 0xd5, 0x20, 0x70, 0x8f, 0xbf, 0x39, 0x50, 0x1d, 0x3c, 0x60, 0xa6, 0xc8, 0xbf,
	0x70, 0xa0, 0xcf, 0x2d, 0x9f, 0xee, 0x98, 0x4b, 0x85, 0x79, 0xf9, 0xbc, 0x90, 0x05, 0x1b, 0xe4,
	0x1f, 0x08, 0x57, 0xf8, 0x95, 0xa7, 0x8d, 0x2c, 0x70, 0xc9, 0x0e, 0x6c, 0xad, 0xbc, 0x47, 0x64,
	0x81, 0x47, 0x0e, 0x61, 0x7f, 0xb9, 0x0e, 0x6f, 0xe6, 0x2c, 0x56, 0x2b, 0x8a, 0x0a, 0xf9, 0x1b,
	0xf6, 0x96, 0x74, 0x84, 0x31, 0x5b, 0x92, 0xd5, 0xe3, 0x04, 0x1a, 0x4f, 0x43, 0x24, 0xbb, 0x40,
	0xee, 0x16, 0x52, 0xe1, 0xec, 0x3a, 0xe6, 0x99, 0x8a, 0x79, 0x16, 0x67, 0x14, 0x83, 0x0d, 0xd2,
	0x02, 0x88, 0x8a, 0x14, 0x6d, 0xe6, 0xc0, 0xd3, 0x25, 0x9c, 0xdb, 0xa5, 0x29, 0xa1, 0x8a, 0x86,
	0xae, 0x78, 0x8e, 0x54, 0x5d, 0xa3, 0x94, 0xf1, 0x18, 0x83, 0xaa, 0x99, 0x09, 0x7e, 0x92, 0x41,
	0x2d, 0xa9, 0x99, 0x2f, 0xe7, 0xd9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x6f, 0x83, 0x61,
	0x82, 0x05, 0x00, 0x00,
}