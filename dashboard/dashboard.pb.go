// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dashboard.proto

package dashboard

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	account "github.com/subiz/header/account"
	common "github.com/subiz/header/common"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AccessLog_Action int32

const (
	AccessLog_login_success         AccessLog_Action = 0
	AccessLog_refresh_token_success AccessLog_Action = 1
	AccessLog_logout                AccessLog_Action = 2
	AccessLog_login_failed          AccessLog_Action = 3
	AccessLog_refresh_token_failed  AccessLog_Action = 4
)

var AccessLog_Action_name = map[int32]string{
	0: "login_success",
	1: "refresh_token_success",
	2: "logout",
	3: "login_failed",
	4: "refresh_token_failed",
}

var AccessLog_Action_value = map[string]int32{
	"login_success":         0,
	"refresh_token_success": 1,
	"logout":                2,
	"login_failed":          3,
	"refresh_token_failed":  4,
}

func (x AccessLog_Action) String() string {
	return proto.EnumName(AccessLog_Action_name, int32(x))
}

func (AccessLog_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{2, 0}
}

type DashboardAgent struct {
	AccountId                           string           `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id                                  string           `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Account                             *account.Account `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	Agent                               *account.Agent   `protobuf:"bytes,5,opt,name=agent,proto3" json:"agent,omitempty"`
	TourGuideWebShowed                  int64            `protobuf:"varint,6,opt,name=tour_guide_web_showed,json=tourGuideWebShowed,proto3" json:"tour_guide_web_showed,omitempty"`
	RefreshToken                        string           `protobuf:"bytes,7,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	AccessToken                         string           `protobuf:"bytes,8,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	LanguageUrl                         string           `protobuf:"bytes,9,opt,name=language_url,json=languageUrl,proto3" json:"language_url,omitempty"`
	TourGuideGettingStartedWebShowed    int64            `protobuf:"varint,10,opt,name=tour_guide_getting_started_web_showed,json=tourGuideGettingStartedWebShowed,proto3" json:"tour_guide_getting_started_web_showed,omitempty"`
	TourGuideChannelInstallingWebShowed int64            `protobuf:"varint,11,opt,name=tour_guide_channel_installing_web_showed,json=tourGuideChannelInstallingWebShowed,proto3" json:"tour_guide_channel_installing_web_showed,omitempty"`
	XXX_NoUnkeyedLiteral                struct{}         `json:"-"`
	XXX_unrecognized                    []byte           `json:"-"`
	XXX_sizecache                       int32            `json:"-"`
}

func (m *DashboardAgent) Reset()         { *m = DashboardAgent{} }
func (m *DashboardAgent) String() string { return proto.CompactTextString(m) }
func (*DashboardAgent) ProtoMessage()    {}
func (*DashboardAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{0}
}

func (m *DashboardAgent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardAgent.Unmarshal(m, b)
}
func (m *DashboardAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardAgent.Marshal(b, m, deterministic)
}
func (m *DashboardAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardAgent.Merge(m, src)
}
func (m *DashboardAgent) XXX_Size() int {
	return xxx_messageInfo_DashboardAgent.Size(m)
}
func (m *DashboardAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardAgent.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardAgent proto.InternalMessageInfo

func (m *DashboardAgent) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *DashboardAgent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DashboardAgent) GetAccount() *account.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *DashboardAgent) GetAgent() *account.Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *DashboardAgent) GetTourGuideWebShowed() int64 {
	if m != nil {
		return m.TourGuideWebShowed
	}
	return 0
}

func (m *DashboardAgent) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *DashboardAgent) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *DashboardAgent) GetLanguageUrl() string {
	if m != nil {
		return m.LanguageUrl
	}
	return ""
}

func (m *DashboardAgent) GetTourGuideGettingStartedWebShowed() int64 {
	if m != nil {
		return m.TourGuideGettingStartedWebShowed
	}
	return 0
}

func (m *DashboardAgent) GetTourGuideChannelInstallingWebShowed() int64 {
	if m != nil {
		return m.TourGuideChannelInstallingWebShowed
	}
	return 0
}

type DashboardAccount struct {
	Id                   string           `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Account              *account.Account `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DashboardAccount) Reset()         { *m = DashboardAccount{} }
func (m *DashboardAccount) String() string { return proto.CompactTextString(m) }
func (*DashboardAccount) ProtoMessage()    {}
func (*DashboardAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{1}
}

func (m *DashboardAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardAccount.Unmarshal(m, b)
}
func (m *DashboardAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardAccount.Marshal(b, m, deterministic)
}
func (m *DashboardAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardAccount.Merge(m, src)
}
func (m *DashboardAccount) XXX_Size() int {
	return xxx_messageInfo_DashboardAccount.Size(m)
}
func (m *DashboardAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardAccount.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardAccount proto.InternalMessageInfo

func (m *DashboardAccount) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DashboardAccount) GetAccount() *account.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AccessLog struct {
	Id            string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	AccountId     string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AgentId       string `protobuf:"bytes,4,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	AcToken       string `protobuf:"bytes,5,opt,name=ac_token,json=acToken,proto3" json:"ac_token,omitempty"`
	RfToken       string `protobuf:"bytes,6,opt,name=rf_token,json=rfToken,proto3" json:"rf_token,omitempty"`
	Ip            string `protobuf:"bytes,7,opt,name=ip,proto3" json:"ip,omitempty"`
	LocationId    int32  `protobuf:"varint,8,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	CityName      string `protobuf:"bytes,9,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`
	CountryName   string `protobuf:"bytes,10,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
	CountryCode   string `protobuf:"bytes,11,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	ContinentCode string `protobuf:"bytes,12,opt,name=continent_code,json=continentCode,proto3" json:"continent_code,omitempty"`
	ContinentName string `protobuf:"bytes,13,opt,name=continent_name,json=continentName,proto3" json:"continent_name,omitempty"`
	//optional string coutry_code = 14;
	Latitude             float32  `protobuf:"fixed32,15,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,16,opt,name=longitude,proto3" json:"longitude,omitempty"`
	PostalCode           string   `protobuf:"bytes,17,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Timezone             string   `protobuf:"bytes,18,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Isp                  string   `protobuf:"bytes,19,opt,name=isp,proto3" json:"isp,omitempty"`
	UserAgent            string   `protobuf:"bytes,20,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Created              int64    `protobuf:"varint,22,opt,name=created,proto3" json:"created,omitempty"`
	Day                  int32    `protobuf:"varint,23,opt,name=day,proto3" json:"day,omitempty"`
	Action               string   `protobuf:"bytes,24,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessLog) Reset()         { *m = AccessLog{} }
func (m *AccessLog) String() string { return proto.CompactTextString(m) }
func (*AccessLog) ProtoMessage()    {}
func (*AccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{2}
}

func (m *AccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLog.Unmarshal(m, b)
}
func (m *AccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLog.Marshal(b, m, deterministic)
}
func (m *AccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLog.Merge(m, src)
}
func (m *AccessLog) XXX_Size() int {
	return xxx_messageInfo_AccessLog.Size(m)
}
func (m *AccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLog proto.InternalMessageInfo

func (m *AccessLog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccessLog) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AccessLog) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *AccessLog) GetAcToken() string {
	if m != nil {
		return m.AcToken
	}
	return ""
}

func (m *AccessLog) GetRfToken() string {
	if m != nil {
		return m.RfToken
	}
	return ""
}

func (m *AccessLog) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *AccessLog) GetLocationId() int32 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func (m *AccessLog) GetCityName() string {
	if m != nil {
		return m.CityName
	}
	return ""
}

func (m *AccessLog) GetCountryName() string {
	if m != nil {
		return m.CountryName
	}
	return ""
}

func (m *AccessLog) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *AccessLog) GetContinentCode() string {
	if m != nil {
		return m.ContinentCode
	}
	return ""
}

func (m *AccessLog) GetContinentName() string {
	if m != nil {
		return m.ContinentName
	}
	return ""
}

func (m *AccessLog) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *AccessLog) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *AccessLog) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *AccessLog) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *AccessLog) GetIsp() string {
	if m != nil {
		return m.Isp
	}
	return ""
}

func (m *AccessLog) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *AccessLog) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *AccessLog) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *AccessLog) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type SessionCookie struct {
	RefreshToken         string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiredAt            int64    `protobuf:"varint,4,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	IssuedAt             int64    `protobuf:"varint,5,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	RememberMe           bool     `protobuf:"varint,8,opt,name=remember_me,json=rememberMe,proto3" json:"remember_me,omitempty"`
	AccountId            string   `protobuf:"bytes,9,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AgentId              string   `protobuf:"bytes,10,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionCookie) Reset()         { *m = SessionCookie{} }
func (m *SessionCookie) String() string { return proto.CompactTextString(m) }
func (*SessionCookie) ProtoMessage()    {}
func (*SessionCookie) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{3}
}

func (m *SessionCookie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionCookie.Unmarshal(m, b)
}
func (m *SessionCookie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionCookie.Marshal(b, m, deterministic)
}
func (m *SessionCookie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionCookie.Merge(m, src)
}
func (m *SessionCookie) XXX_Size() int {
	return xxx_messageInfo_SessionCookie.Size(m)
}
func (m *SessionCookie) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionCookie.DiscardUnknown(m)
}

var xxx_messageInfo_SessionCookie proto.InternalMessageInfo

func (m *SessionCookie) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *SessionCookie) GetExpiredAt() int64 {
	if m != nil {
		return m.ExpiredAt
	}
	return 0
}

func (m *SessionCookie) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *SessionCookie) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SessionCookie) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SessionCookie) GetRememberMe() bool {
	if m != nil {
		return m.RememberMe
	}
	return false
}

func (m *SessionCookie) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *SessionCookie) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

// global setting. eg: dashboard_4/4234098234
type Global struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data                 string          `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Pk                   string          `protobuf:"bytes,4,opt,name=pk,proto3" json:"pk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{4}
}

func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (m *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(m, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Global) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Global) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Global) GetPk() string {
	if m != nil {
		return m.Pk
	}
	return ""
}

func init() {
	proto.RegisterEnum("dashboard.AccessLog_Action", AccessLog_Action_name, AccessLog_Action_value)
	proto.RegisterType((*DashboardAgent)(nil), "dashboard.DashboardAgent")
	proto.RegisterType((*DashboardAccount)(nil), "dashboard.DashboardAccount")
	proto.RegisterType((*AccessLog)(nil), "dashboard.AccessLog")
	proto.RegisterType((*SessionCookie)(nil), "dashboard.SessionCookie")
	proto.RegisterType((*Global)(nil), "dashboard.Global")
}

func init() { proto.RegisterFile("dashboard.proto", fileDescriptor_9b97678da3a35dfb) }

var fileDescriptor_9b97678da3a35dfb = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdb, 0x6e, 0xe3, 0x36,
	0x10, 0xad, 0xad, 0xd8, 0xb1, 0x26, 0x76, 0xe2, 0x65, 0xb3, 0x5b, 0x25, 0xed, 0xa2, 0x8e, 0xd3,
	0x05, 0x8c, 0x3e, 0x24, 0x68, 0xfb, 0x05, 0x6e, 0x0a, 0x04, 0x01, 0xda, 0x2d, 0xa0, 0x74, 0x51,
	0xa0, 0x2f, 0x02, 0x2d, 0x4e, 0x64, 0x22, 0x12, 0x29, 0x50, 0x14, 0x36, 0xd9, 0x3f, 0xec, 0x17,
	0xf4, 0x4b, 0xfa, 0x5e, 0x70, 0x48, 0xcb, 0xb9, 0xa1, 0xfb, 0x64, 0xce, 0x39, 0x47, 0xc3, 0x31,
	0xcf, 0x0c, 0x09, 0x07, 0x82, 0x37, 0xeb, 0x95, 0xe6, 0x46, 0x9c, 0xd5, 0x46, 0x5b, 0xcd, 0xe2,
	0x0e, 0x38, 0x1e, 0xe7, 0xba, 0xaa, 0xb4, 0xf2, 0xc4, 0xf1, 0x84, 0xe7, 0xb9, 0x6e, 0x95, 0xf5,
	0xe1, 0xfc, 0x9f, 0x08, 0xf6, 0x7f, 0xd9, 0x48, 0x97, 0x05, 0x2a, 0xcb, 0xde, 0x02, 0x04, 0x4d,
	0x26, 0x45, 0xd2, 0x9f, 0xf5, 0x16, 0x71, 0x1a, 0x07, 0xe4, 0x4a, 0xb0, 0x7d, 0xe8, 0x4b, 0x91,
	0x44, 0x04, 0xf7, 0xa5, 0x60, 0xdf, 0xc3, 0x6e, 0x20, 0x93, 0x9d, 0x59, 0x6f, 0xb1, 0xf7, 0xe3,
	0xf4, 0x6c, 0xb3, 0xc5, 0xd2, 0xff, 0xa6, 0x1b, 0x01, 0xfb, 0x0e, 0x06, 0xdc, 0xed, 0x91, 0x0c,
	0x48, 0xb9, 0xbf, 0x55, 0x3a, 0x34, 0xf5, 0x24, 0xfb, 0x01, 0x5e, 0x5b, 0xdd, 0x9a, 0xac, 0x68,
	0xa5, 0xc0, 0xec, 0x23, 0xae, 0xb2, 0x66, 0xad, 0x3f, 0xa2, 0x48, 0x86, 0xb3, 0xde, 0x22, 0x4a,
	0x99, 0x23, 0x2f, 0x1d, 0xf7, 0x27, 0xae, 0xae, 0x89, 0x61, 0xa7, 0x30, 0x31, 0x78, 0x63, 0xb0,
	0x59, 0x67, 0x56, 0xdf, 0xa2, 0x4a, 0x76, 0xa9, 0xbe, 0x71, 0x00, 0xff, 0x70, 0x18, 0x3b, 0x81,
	0x31, 0xcf, 0x73, 0x6c, 0x9a, 0xa0, 0x19, 0x91, 0x66, 0xcf, 0x63, 0x9d, 0xa4, 0xe4, 0xaa, 0x68,
	0x79, 0x81, 0x59, 0x6b, 0xca, 0x24, 0xf6, 0x92, 0x0d, 0xf6, 0xc1, 0x94, 0xec, 0x77, 0x78, 0xf7,
	0xa0, 0xba, 0x02, 0xad, 0x95, 0xaa, 0xc8, 0x1a, 0xcb, 0x8d, 0x45, 0xf1, 0xb0, 0x5a, 0xa0, 0x6a,
	0x67, 0x5d, 0xb5, 0x97, 0x5e, 0x7a, 0xed, 0x95, 0xdb, 0xda, 0x3f, 0xc0, 0xe2, 0x41, 0xc2, 0x7c,
	0xcd, 0x95, 0xc2, 0x32, 0x93, 0xaa, 0xb1, 0xbc, 0x2c, 0x5d, 0xee, 0x07, 0x39, 0xf7, 0x28, 0xe7,
	0x69, 0x97, 0xf3, 0xc2, 0xab, 0xaf, 0x3a, 0x71, 0x97, 0x76, 0xfe, 0x1e, 0xa6, 0x5b, 0x63, 0xc3,
	0xf9, 0x7b, 0xef, 0xfa, 0x2f, 0x79, 0x17, 0x7d, 0xc6, 0xbb, 0xf9, 0xdf, 0x03, 0x88, 0x97, 0x74,
	0x54, 0xbf, 0xea, 0xe2, 0x59, 0xa6, 0xc7, 0x4d, 0x13, 0x3d, 0x6d, 0x9a, 0x23, 0x18, 0x91, 0xb7,
	0x8e, 0xdc, 0x21, 0x72, 0x97, 0xe2, 0x40, 0xe5, 0xc1, 0x91, 0x41, 0xa0, 0x72, 0xef, 0xc6, 0x11,
	0x8c, 0xcc, 0x4d, 0xa0, 0x86, 0x9e, 0x32, 0x37, 0x9e, 0x72, 0xfb, 0xd7, 0xc1, 0xe5, 0xbe, 0xac,
	0xd9, 0xb7, 0xb0, 0x57, 0xea, 0x9c, 0x5b, 0xa9, 0x95, 0xdb, 0xc3, 0x59, 0x3b, 0x48, 0x61, 0x03,
	0x5d, 0x09, 0xf6, 0x35, 0xc4, 0xb9, 0xb4, 0xf7, 0x99, 0xe2, 0x15, 0x06, 0x5b, 0x47, 0x0e, 0x78,
	0xcf, 0x2b, 0x74, 0xb6, 0x53, 0xa5, 0x26, 0xf0, 0xe0, 0x6d, 0x0f, 0xd8, 0x53, 0x49, 0xae, 0x05,
	0x92, 0x13, 0x5b, 0xc9, 0x85, 0x16, 0xc8, 0xde, 0xc1, 0x7e, 0xae, 0x95, 0x95, 0xca, 0xfd, 0x51,
	0x12, 0x8d, 0x49, 0x34, 0xe9, 0xd0, 0xe7, 0x32, 0xda, 0x6e, 0xf2, 0x44, 0x46, 0x1b, 0x1e, 0xc3,
	0xa8, 0xe4, 0x56, 0xda, 0x56, 0x60, 0x72, 0x30, 0xeb, 0x2d, 0xfa, 0x69, 0x17, 0xb3, 0x6f, 0x20,
	0x2e, 0xb5, 0x2a, 0x3c, 0x39, 0x25, 0x72, 0x0b, 0xb8, 0xb3, 0xa8, 0xb5, 0x6b, 0x08, 0x5f, 0xc4,
	0x2b, 0xca, 0x0e, 0x1e, 0xa2, 0x0a, 0x8e, 0x61, 0x64, 0x65, 0x85, 0x9f, 0xb4, 0xc2, 0x84, 0xf9,
	0xa3, 0xd8, 0xc4, 0x6c, 0x0a, 0x91, 0x6c, 0xea, 0xe4, 0x4b, 0x82, 0xdd, 0xd2, 0x59, 0xdb, 0x36,
	0x68, 0x32, 0x3f, 0xb9, 0x87, 0xde, 0x5a, 0x87, 0xf8, 0xeb, 0x22, 0x81, 0xdd, 0xdc, 0x20, 0xb7,
	0x28, 0x92, 0x37, 0xd4, 0x9d, 0x9b, 0xd0, 0xa5, 0x12, 0xfc, 0x3e, 0xf9, 0x8a, 0xbc, 0x70, 0x4b,
	0xf6, 0x06, 0x86, 0x3c, 0x77, 0x86, 0x24, 0x09, 0xa5, 0x09, 0xd1, 0x5c, 0xc1, 0x70, 0x49, 0x2b,
	0xf6, 0x0a, 0x26, 0xa5, 0x2e, 0xa4, 0xca, 0x9a, 0x96, 0x7a, 0x6d, 0xfa, 0x05, 0x3b, 0x82, 0xd7,
	0x8f, 0x66, 0xbb, 0xa3, 0x7a, 0x0c, 0x60, 0x58, 0xea, 0x42, 0xb7, 0x76, 0xda, 0x67, 0x53, 0x18,
	0xfb, 0x2f, 0x6f, 0xb8, 0x2c, 0x51, 0x4c, 0x23, 0x96, 0xc0, 0xe1, 0xe3, 0x0f, 0x03, 0xb3, 0x33,
	0xff, 0xb7, 0x07, 0x93, 0x6b, 0x6c, 0x1a, 0xa9, 0xd5, 0x85, 0xd6, 0xb7, 0x12, 0x9f, 0x5f, 0x20,
	0xd1, 0x0b, 0x17, 0xc8, 0x5b, 0x00, 0xbc, 0xab, 0xa5, 0x41, 0x91, 0x71, 0x7f, 0xdb, 0x45, 0x69,
	0x1c, 0x90, 0xa5, 0x75, 0x2d, 0x26, 0x9b, 0xa6, 0xf5, 0xec, 0x80, 0xd8, 0x91, 0x07, 0x96, 0x96,
	0x31, 0xd8, 0xb1, 0xf7, 0x35, 0x86, 0x3e, 0xa6, 0x35, 0x3b, 0x84, 0x01, 0x56, 0x5c, 0x96, 0xa1,
	0x8f, 0x7d, 0xe0, 0xec, 0x33, 0x58, 0x61, 0xb5, 0x42, 0x93, 0x55, 0x48, 0xad, 0x3c, 0x4a, 0x61,
	0x03, 0xfd, 0x86, 0x4f, 0x66, 0x2d, 0xfe, 0xbf, 0x59, 0x83, 0x47, 0xb3, 0x36, 0xcf, 0x61, 0x78,
	0x59, 0xea, 0x15, 0x2f, 0xd9, 0x09, 0x44, 0xb9, 0xbd, 0x4b, 0x7a, 0x34, 0xf5, 0x07, 0x67, 0xe1,
	0x89, 0xb8, 0xd0, 0xca, 0xe2, 0x9d, 0x4d, 0x1d, 0xe7, 0x2a, 0xa6, 0xee, 0xf4, 0x43, 0x4e, 0x6b,
	0x87, 0x09, 0x6e, 0x79, 0x38, 0x1d, 0x5a, 0xbb, 0x51, 0xac, 0x6f, 0xc3, 0x54, 0xf7, 0xeb, 0xdb,
	0x9f, 0x4f, 0xff, 0x3a, 0x29, 0xa4, 0x5d, 0xb7, 0x2b, 0x97, 0xf5, 0xbc, 0x69, 0x57, 0xf2, 0xd3,
	0xf9, 0x1a, 0xb9, 0x40, 0x73, 0xde, 0x3d, 0x4a, 0xab, 0x21, 0x3d, 0x3f, 0x3f, 0xfd, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x00, 0xa1, 0xa2, 0x26, 0xb9, 0x06, 0x00, 0x00,
}
