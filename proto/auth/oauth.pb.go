// Code generated by protoc-gen-go.
// source: oauth.proto
// DO NOT EDIT!

package auth

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

type ErrorCode int32

const (
	ErrorCode_no_error                  ErrorCode = 0
	ErrorCode_resource_notfoud          ErrorCode = 1
	ErrorCode_invalid_request           ErrorCode = 2
	ErrorCode_unauthorized              ErrorCode = 3
	ErrorCode_invalide_client           ErrorCode = 4
	ErrorCode_invalid_grant             ErrorCode = 5
	ErrorCode_unauthorized_client       ErrorCode = 6
	ErrorCode_unsupported_grant_type    ErrorCode = 7
	ErrorCode_invalid_scope             ErrorCode = 8
	ErrorCode_access_denied             ErrorCode = 9
	ErrorCode_unsupported_response_type ErrorCode = 10
	ErrorCode_server_error              ErrorCode = 11
	ErrorCode_temporarily_unavailable   ErrorCode = 12
)

var ErrorCode_name = map[int32]string{
	0:  "no_error",
	1:  "resource_notfoud",
	2:  "invalid_request",
	3:  "unauthorized",
	4:  "invalide_client",
	5:  "invalid_grant",
	6:  "unauthorized_client",
	7:  "unsupported_grant_type",
	8:  "invalid_scope",
	9:  "access_denied",
	10: "unsupported_response_type",
	11: "server_error",
	12: "temporarily_unavailable",
}
var ErrorCode_value = map[string]int32{
	"no_error":                  0,
	"resource_notfoud":          1,
	"invalid_request":           2,
	"unauthorized":              3,
	"invalide_client":           4,
	"invalid_grant":             5,
	"unauthorized_client":       6,
	"unsupported_grant_type":    7,
	"invalid_scope":             8,
	"access_denied":             9,
	"unsupported_response_type": 10,
	"server_error":              11,
	"temporarily_unavailable":   12,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type TokenRequest_GrantType int32

const (
	TokenRequest_Refresh_token      TokenRequest_GrantType = 0
	TokenRequest_Authorization_code TokenRequest_GrantType = 1
	TokenRequest_Password           TokenRequest_GrantType = 2
	TokenRequest_Client_credentials TokenRequest_GrantType = 3
)

var TokenRequest_GrantType_name = map[int32]string{
	0: "Refresh_token",
	1: "Authorization_code",
	2: "Password",
	3: "Client_credentials",
}
var TokenRequest_GrantType_value = map[string]int32{
	"Refresh_token":      0,
	"Authorization_code": 1,
	"Password":           2,
	"Client_credentials": 3,
}

func (x TokenRequest_GrantType) String() string {
	return proto.EnumName(TokenRequest_GrantType_name, int32(x))
}
func (TokenRequest_GrantType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 0} }

type AccessTokenResponse_TokenType int32

const (
	AccessTokenResponse_bearer AccessTokenResponse_TokenType = 0
	AccessTokenResponse_mac    AccessTokenResponse_TokenType = 1
)

var AccessTokenResponse_TokenType_name = map[int32]string{
	0: "bearer",
	1: "mac",
}
var AccessTokenResponse_TokenType_value = map[string]int32{
	"bearer": 0,
	"mac":    1,
}

func (x AccessTokenResponse_TokenType) String() string {
	return proto.EnumName(AccessTokenResponse_TokenType_name, int32(x))
}
func (AccessTokenResponse_TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{6, 0}
}

type AuthorizeRequest_AuthorizeType int32

const (
	AuthorizeRequest_code  AuthorizeRequest_AuthorizeType = 0
	AuthorizeRequest_token AuthorizeRequest_AuthorizeType = 1
)

var AuthorizeRequest_AuthorizeType_name = map[int32]string{
	0: "code",
	1: "token",
}
var AuthorizeRequest_AuthorizeType_value = map[string]int32{
	"code":  0,
	"token": 1,
}

func (x AuthorizeRequest_AuthorizeType) String() string {
	return proto.EnumName(AuthorizeRequest_AuthorizeType_name, int32(x))
}
func (AuthorizeRequest_AuthorizeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{7, 0}
}

type UserRequest struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	FacebookId string `protobuf:"bytes,3,opt,name=facebook_id,json=facebookId" json:"facebook_id,omitempty"`
	GoogleId   string `protobuf:"bytes,4,opt,name=google_id,json=googleId" json:"google_id,omitempty"`
	YahooId    string `protobuf:"bytes,5,opt,name=yahoo_id,json=yahooId" json:"yahoo_id,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *UserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRequest) GetFacebookId() string {
	if m != nil {
		return m.FacebookId
	}
	return ""
}

func (m *UserRequest) GetGoogleId() string {
	if m != nil {
		return m.GoogleId
	}
	return ""
}

func (m *UserRequest) GetYahooId() string {
	if m != nil {
		return m.YahooId
	}
	return ""
}

type ClientRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	// any client internal = true are grant unlimited access. Only set this field to internal client (service)
	Internal bool `protobuf:"varint,3,opt,name=internal" json:"internal,omitempty"`
}

func (m *ClientRequest) Reset()                    { *m = ClientRequest{} }
func (m *ClientRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()               {}
func (*ClientRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ClientRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ClientRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *ClientRequest) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

type Response struct {
	Error            ErrorCode `protobuf:"varint,1,opt,name=error,enum=auth.ErrorCode" json:"error,omitempty"`
	ErrorDescription string    `protobuf:"bytes,2,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Response) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_no_error
}

func (m *Response) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type IdRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *IdRequest) Reset()                    { *m = IdRequest{} }
func (m *IdRequest) String() string            { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()               {}
func (*IdRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *IdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RenewTokenRequest struct {
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *RenewTokenRequest) Reset()                    { *m = RenewTokenRequest{} }
func (m *RenewTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*RenewTokenRequest) ProtoMessage()               {}
func (*RenewTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *RenewTokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type TokenRequest struct {
	GrantType    TokenRequest_GrantType `protobuf:"varint,1,opt,name=grant_type,json=grantType,enum=auth.TokenRequest_GrantType" json:"grant_type,omitempty"`
	Username     string                 `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password     string                 `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Code         string                 `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	RedirectUri  string                 `protobuf:"bytes,5,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
	Client       string                 `protobuf:"bytes,6,opt,name=client" json:"client,omitempty"`
	Scope        []string               `protobuf:"bytes,7,rep,name=scope" json:"scope,omitempty"`
	RefreshToken string                 `protobuf:"bytes,8,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *TokenRequest) GetGrantType() TokenRequest_GrantType {
	if m != nil {
		return m.GrantType
	}
	return TokenRequest_Refresh_token
}

func (m *TokenRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *TokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *TokenRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *TokenRequest) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *TokenRequest) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *TokenRequest) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *TokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type AccessTokenResponse struct {
	Error            ErrorCode                     `protobuf:"varint,1,opt,name=error,enum=auth.ErrorCode" json:"error,omitempty"`
	ErrorDescription string                        `protobuf:"bytes,2,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
	AccessToken      string                        `protobuf:"bytes,3,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	TokenType        AccessTokenResponse_TokenType `protobuf:"varint,4,opt,name=token_type,json=tokenType,enum=auth.AccessTokenResponse_TokenType" json:"token_type,omitempty"`
	ExpiresIn        int64                         `protobuf:"varint,5,opt,name=expires_in,json=expiresIn" json:"expires_in,omitempty"`
	RefreshToken     string                        `protobuf:"bytes,6,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	Scope            []string                      `protobuf:"bytes,7,rep,name=scope" json:"scope,omitempty"`
	// authorize_code response
	Code string `protobuf:"bytes,8,opt,name=code" json:"code,omitempty"`
}

func (m *AccessTokenResponse) Reset()                    { *m = AccessTokenResponse{} }
func (m *AccessTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*AccessTokenResponse) ProtoMessage()               {}
func (*AccessTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *AccessTokenResponse) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_no_error
}

func (m *AccessTokenResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

func (m *AccessTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AccessTokenResponse) GetTokenType() AccessTokenResponse_TokenType {
	if m != nil {
		return m.TokenType
	}
	return AccessTokenResponse_bearer
}

func (m *AccessTokenResponse) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *AccessTokenResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *AccessTokenResponse) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *AccessTokenResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type AuthorizeRequest struct {
	ResponseType AuthorizeRequest_AuthorizeType `protobuf:"varint,1,opt,name=response_type,json=responseType,enum=auth.AuthorizeRequest_AuthorizeType" json:"response_type,omitempty"`
	ClientId     string                         `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	RedirectUri  string                         `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
	Scope        []string                       `protobuf:"bytes,4,rep,name=scope" json:"scope,omitempty"`
}

func (m *AuthorizeRequest) Reset()                    { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()               {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *AuthorizeRequest) GetResponseType() AuthorizeRequest_AuthorizeType {
	if m != nil {
		return m.ResponseType
	}
	return AuthorizeRequest_code
}

func (m *AuthorizeRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AuthorizeRequest) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *AuthorizeRequest) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

type AuthorizeResponse struct {
	Error            ErrorCode `protobuf:"varint,1,opt,name=error,enum=auth.ErrorCode" json:"error,omitempty"`
	ErrorDescription string    `protobuf:"bytes,2,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
}

func (m *AuthorizeResponse) Reset()                    { *m = AuthorizeResponse{} }
func (m *AuthorizeResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthorizeResponse) ProtoMessage()               {}
func (*AuthorizeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *AuthorizeResponse) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_no_error
}

func (m *AuthorizeResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type UnauthorizeRequest struct {
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *UnauthorizeRequest) Reset()                    { *m = UnauthorizeRequest{} }
func (m *UnauthorizeRequest) String() string            { return proto.CompactTextString(m) }
func (*UnauthorizeRequest) ProtoMessage()               {}
func (*UnauthorizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *UnauthorizeRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *UnauthorizeRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "auth.UserRequest")
	proto.RegisterType((*ClientRequest)(nil), "auth.ClientRequest")
	proto.RegisterType((*Response)(nil), "auth.Response")
	proto.RegisterType((*IdRequest)(nil), "auth.IdRequest")
	proto.RegisterType((*RenewTokenRequest)(nil), "auth.RenewTokenRequest")
	proto.RegisterType((*TokenRequest)(nil), "auth.TokenRequest")
	proto.RegisterType((*AccessTokenResponse)(nil), "auth.AccessTokenResponse")
	proto.RegisterType((*AuthorizeRequest)(nil), "auth.AuthorizeRequest")
	proto.RegisterType((*AuthorizeResponse)(nil), "auth.AuthorizeResponse")
	proto.RegisterType((*UnauthorizeRequest)(nil), "auth.UnauthorizeRequest")
	proto.RegisterEnum("auth.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("auth.TokenRequest_GrantType", TokenRequest_GrantType_name, TokenRequest_GrantType_value)
	proto.RegisterEnum("auth.AccessTokenResponse_TokenType", AccessTokenResponse_TokenType_name, AccessTokenResponse_TokenType_value)
	proto.RegisterEnum("auth.AuthorizeRequest_AuthorizeType", AuthorizeRequest_AuthorizeType_name, AuthorizeRequest_AuthorizeType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Oauth2 service

type Oauth2Client interface {
	// Create a new user
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error)
	// Delete user
	DeleteUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error)
	// Update user info
	UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error)
	CreateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error)
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*AccessTokenResponse, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	Unauthorize(ctx context.Context, in *UnauthorizeRequest, opts ...grpc.CallOption) (*Response, error)
}

type oauth2Client struct {
	cc *grpc.ClientConn
}

func NewOauth2Client(cc *grpc.ClientConn) Oauth2Client {
	return &oauth2Client{cc}
}

func (c *oauth2Client) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Oauth2/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) DeleteUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Oauth2/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) UpdateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Oauth2/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) CreateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Oauth2/CreateClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) UpdateClient(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Oauth2/UpdateClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) DeleteClient(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Oauth2/DeleteClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*AccessTokenResponse, error) {
	out := new(AccessTokenResponse)
	err := grpc.Invoke(ctx, "/auth.Oauth2/Token", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := grpc.Invoke(ctx, "/auth.Oauth2/Authorize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) Unauthorize(ctx context.Context, in *UnauthorizeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Oauth2/Unauthorize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Oauth2 service

type Oauth2Server interface {
	// Create a new user
	CreateUser(context.Context, *UserRequest) (*Response, error)
	// Delete user
	DeleteUser(context.Context, *IdRequest) (*Response, error)
	// Update user info
	UpdateUser(context.Context, *UserRequest) (*Response, error)
	CreateClient(context.Context, *ClientRequest) (*Response, error)
	UpdateClient(context.Context, *ClientRequest) (*Response, error)
	DeleteClient(context.Context, *IdRequest) (*Response, error)
	Token(context.Context, *TokenRequest) (*AccessTokenResponse, error)
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	Unauthorize(context.Context, *UnauthorizeRequest) (*Response, error)
}

func RegisterOauth2Server(s *grpc.Server, srv Oauth2Server) {
	s.RegisterService(&_Oauth2_serviceDesc, srv)
}

func _Oauth2_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).DeleteUser(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).UpdateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).CreateClient(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).UpdateClient(ctx, req.(*ClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).DeleteClient(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_Unauthorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnauthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).Unauthorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Oauth2/Unauthorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).Unauthorize(ctx, req.(*UnauthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Oauth2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Oauth2",
	HandlerType: (*Oauth2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Oauth2_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Oauth2_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Oauth2_UpdateUser_Handler,
		},
		{
			MethodName: "CreateClient",
			Handler:    _Oauth2_CreateClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _Oauth2_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Oauth2_DeleteClient_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Oauth2_Token_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _Oauth2_Authorize_Handler,
		},
		{
			MethodName: "Unauthorize",
			Handler:    _Oauth2_Unauthorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}

func init() { proto.RegisterFile("oauth.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 934 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xf6, 0xbf, 0x3d, 0x65, 0x27, 0xdb, 0xae, 0xac, 0x12, 0xc7, 0x61, 0x45, 0x98, 0x5d, 0xa4,
	0x15, 0x48, 0x41, 0x09, 0x42, 0x5a, 0x84, 0x84, 0xb4, 0x64, 0x11, 0x32, 0x17, 0xd0, 0xb0, 0xb9,
	0x21, 0x46, 0x9d, 0xe9, 0x8a, 0xd3, 0x5a, 0x67, 0x7a, 0xe8, 0x99, 0xc9, 0x12, 0x1e, 0x83, 0x97,
	0xe0, 0x99, 0xb8, 0x71, 0xe7, 0xcc, 0x1d, 0xf5, 0xcf, 0xd8, 0xe3, 0xd8, 0x91, 0xd8, 0xc3, 0xde,
	0xba, 0xaa, 0xab, 0xa6, 0xbe, 0xaf, 0xbe, 0xea, 0xee, 0x81, 0xa1, 0xe2, 0x65, 0x71, 0x7d, 0x92,
	0x69, 0x55, 0x28, 0xec, 0x98, 0x75, 0xf8, 0x47, 0x13, 0x86, 0x17, 0x39, 0xe9, 0x88, 0x7e, 0x2d,
	0x29, 0x2f, 0x70, 0x17, 0x5a, 0x52, 0x4c, 0x9a, 0xc7, 0xcd, 0xe7, 0x41, 0xd4, 0x92, 0x02, 0xa7,
	0x30, 0xc8, 0x78, 0x9e, 0xbf, 0x55, 0x5a, 0x4c, 0x5a, 0xd6, 0xbb, 0xb4, 0xf1, 0x43, 0x18, 0x5e,
	0xf1, 0x84, 0x2e, 0x95, 0x7a, 0x13, 0x4b, 0x31, 0x69, 0xdb, 0x6d, 0xa8, 0x5c, 0x33, 0x81, 0x47,
	0x10, 0xcc, 0x95, 0x9a, 0x2f, 0xc8, 0x6c, 0x77, 0x5c, 0xb6, 0x73, 0xcc, 0x04, 0x1e, 0xc2, 0xe0,
	0x8e, 0x5f, 0x2b, 0x65, 0xf6, 0xba, 0x76, 0xaf, 0x6f, 0xed, 0x99, 0x08, 0x7f, 0x82, 0x9d, 0xf3,
	0x85, 0xa4, 0xb4, 0x78, 0x08, 0xd5, 0x3e, 0xf4, 0x72, 0x4a, 0x34, 0x15, 0x1e, 0x93, 0xb7, 0x0c,
	0x5a, 0x99, 0x16, 0xa4, 0x53, 0xbe, 0xb0, 0x70, 0x06, 0xd1, 0xd2, 0x0e, 0x7f, 0x81, 0x41, 0x44,
	0x79, 0xa6, 0xd2, 0x9c, 0xf0, 0x63, 0xe8, 0x92, 0xd6, 0x4a, 0xdb, 0x4f, 0xee, 0x9e, 0x3d, 0x3a,
	0xb1, 0x7d, 0xf9, 0xd6, 0xb8, 0xce, 0x95, 0xa0, 0xc8, 0xed, 0xe2, 0xa7, 0x30, 0xb6, 0x8b, 0x58,
	0x50, 0x9e, 0x68, 0x99, 0x15, 0x52, 0xa5, 0xbe, 0x22, 0xb3, 0x1b, 0xaf, 0x56, 0xfe, 0xf0, 0x08,
	0x82, 0x99, 0x78, 0x00, 0x70, 0xf8, 0x02, 0xc6, 0x11, 0xa5, 0xf4, 0xf6, 0xb5, 0x7a, 0x43, 0x69,
	0x15, 0xf4, 0x14, 0x76, 0x34, 0x5d, 0x69, 0xca, 0xaf, 0xe3, 0xc2, 0xf8, 0x7d, 0xfc, 0xc8, 0x3b,
	0x6d, 0x6c, 0xf8, 0x4f, 0x0b, 0x46, 0x6b, 0x59, 0x5f, 0x01, 0xcc, 0x35, 0x4f, 0x8b, 0xb8, 0xb8,
	0xcb, 0xc8, 0x13, 0xf8, 0xc0, 0x11, 0xa8, 0xc7, 0x9d, 0x7c, 0x67, 0x82, 0x5e, 0xdf, 0x65, 0x14,
	0x05, 0xf3, 0x6a, 0x69, 0x1a, 0x54, 0xe6, 0xa6, 0x1f, 0x37, 0x54, 0xc9, 0x59, 0xd9, 0x6b, 0x52,
	0xb7, 0xef, 0x49, 0x8d, 0xd0, 0x49, 0x94, 0x20, 0x2f, 0xa2, 0x5d, 0xe3, 0x47, 0x30, 0xd2, 0x24,
	0xa4, 0xa6, 0xa4, 0x88, 0x4b, 0x2d, 0xbd, 0x88, 0xc3, 0xca, 0x77, 0xa1, 0xa5, 0xd1, 0x29, 0xb1,
	0x42, 0x4e, 0x7a, 0x4e, 0x27, 0x67, 0xe1, 0x63, 0xe8, 0xe6, 0x89, 0xca, 0x68, 0xd2, 0x3f, 0x6e,
	0x3f, 0x0f, 0x22, 0x67, 0x6c, 0xf6, 0x63, 0xb0, 0xa5, 0x1f, 0x3f, 0x43, 0xb0, 0x64, 0x86, 0x63,
	0xd8, 0x89, 0xea, 0x19, 0xac, 0x81, 0xfb, 0x80, 0x2f, 0xcb, 0xe2, 0x5a, 0x69, 0xf9, 0x3b, 0x37,
	0xba, 0xc4, 0x06, 0x2b, 0x6b, 0xe2, 0x08, 0x06, 0x3f, 0x7a, 0x36, 0xac, 0x65, 0xa2, 0xdc, 0x84,
	0xc5, 0x89, 0x26, 0x41, 0x69, 0x21, 0xf9, 0x22, 0x67, 0xed, 0xf0, 0xef, 0x16, 0xec, 0xbd, 0x4c,
	0x12, 0xca, 0x73, 0xdf, 0xcb, 0xf7, 0x37, 0x30, 0xa6, 0x7f, 0xdc, 0x96, 0xf2, 0x6c, 0x5d, 0xcf,
	0x87, 0x7c, 0x55, 0x1e, 0xbf, 0x01, 0xb0, 0x7b, 0x4e, 0xeb, 0x8e, 0xad, 0xfd, 0xd4, 0xd5, 0xde,
	0x82, 0xd2, 0xe9, 0xef, 0x24, 0x2f, 0xaa, 0x25, 0x3e, 0x01, 0xa0, 0xdf, 0x32, 0xa9, 0x29, 0x8f,
	0x65, 0x6a, 0x45, 0x6a, 0x47, 0x81, 0xf7, 0xcc, 0xd2, 0xcd, 0xa6, 0xf7, 0x36, 0x9b, 0xfe, 0x80,
	0x5e, 0xd5, 0x50, 0x0c, 0x56, 0x43, 0x11, 0x1e, 0x43, 0xb0, 0x44, 0x81, 0x00, 0xbd, 0x4b, 0xe2,
	0x9a, 0x34, 0x6b, 0x60, 0x1f, 0xda, 0x37, 0x3c, 0x61, 0xcd, 0xf0, 0xaf, 0x26, 0xb0, 0x4a, 0x21,
	0xaa, 0x86, 0x7a, 0x66, 0x50, 0x38, 0x16, 0xf5, 0xb9, 0x7e, 0xe6, 0xb9, 0xde, 0x0b, 0x5f, 0x39,
	0x2c, 0xd9, 0x51, 0x95, 0x6a, 0x8b, 0x1e, 0x41, 0xe0, 0xa6, 0xcc, 0x5c, 0x2c, 0x7e, 0xc6, 0x9d,
	0x63, 0x26, 0x36, 0x66, 0xb6, 0xbd, 0x39, 0xb3, 0x4b, 0xae, 0x9d, 0x1a, 0xd7, 0xf0, 0x19, 0xec,
	0xac, 0x15, 0xc5, 0x81, 0x23, 0xcf, 0x1a, 0x18, 0x40, 0xd7, 0x0d, 0x5f, 0x33, 0x9c, 0xc3, 0xb8,
	0x86, 0xf5, 0x3d, 0x5e, 0x36, 0xdf, 0x03, 0x5e, 0xa4, 0xfc, 0x7e, 0x17, 0xd7, 0xa8, 0x37, 0xef,
	0x51, 0x3f, 0x80, 0xbe, 0x39, 0xea, 0xab, 0xae, 0xf4, 0x8c, 0x39, 0x13, 0x9f, 0xfc, 0xd9, 0x82,
	0x60, 0x89, 0xc6, 0x9c, 0x93, 0x54, 0xc5, 0xb6, 0x20, 0x6b, 0xe0, 0x63, 0x60, 0x9a, 0x72, 0x55,
	0xea, 0x84, 0xe2, 0x54, 0x15, 0x57, 0xaa, 0x14, 0xac, 0x89, 0x7b, 0xf0, 0x48, 0xa6, 0xb7, 0x7c,
	0x21, 0x45, 0xac, 0x5d, 0x69, 0xd6, 0x42, 0x06, 0xa3, 0x72, 0x05, 0x49, 0xb0, 0x76, 0x2d, 0x8c,
	0x62, 0x07, 0x83, 0x75, 0xcc, 0x91, 0xad, 0x72, 0xed, 0xb5, 0xc4, 0xba, 0x78, 0x00, 0x7b, 0xf5,
	0xcc, 0x2a, 0xb6, 0x87, 0x53, 0xd8, 0x2f, 0xd3, 0xbc, 0xcc, 0x32, 0xa5, 0x0b, 0xf2, 0xf1, 0x76,
	0x3c, 0x58, 0xbf, 0xfe, 0x1d, 0xab, 0x10, 0x1b, 0x18, 0x97, 0x3f, 0x50, 0x82, 0x52, 0x49, 0x82,
	0x05, 0xf8, 0x04, 0x0e, 0xeb, 0x5f, 0x58, 0x9b, 0x31, 0x06, 0x06, 0x73, 0x4e, 0xfa, 0x96, 0xb4,
	0x27, 0x3c, 0xc4, 0x23, 0x38, 0x28, 0xe8, 0x26, 0x53, 0x9a, 0x6b, 0xb9, 0xb8, 0x8b, 0xcb, 0x94,
	0xdf, 0x72, 0xb9, 0xe0, 0x97, 0x0b, 0x62, 0xa3, 0xb3, 0x7f, 0xdb, 0xd0, 0xfb, 0xc1, 0x00, 0x3d,
	0xc3, 0x53, 0x80, 0x73, 0x4d, 0xbc, 0x20, 0xf3, 0x78, 0xe2, 0xd8, 0x69, 0x5a, 0x7b, 0x48, 0xa7,
	0xbb, 0xce, 0x55, 0x4d, 0x41, 0xd8, 0xc0, 0xcf, 0x00, 0x5e, 0xd1, 0x82, 0x7c, 0x8a, 0x1f, 0x83,
	0xe5, 0x93, 0xb1, 0x25, 0xe1, 0x14, 0xe0, 0x22, 0x13, 0xef, 0x54, 0xe3, 0x0b, 0x18, 0x39, 0x58,
	0xee, 0x76, 0xc3, 0x3d, 0x17, 0xb1, 0xf6, 0x9a, 0x6e, 0x4f, 0x73, 0x95, 0xde, 0x2d, 0xed, 0x14,
	0x46, 0x8e, 0x91, 0x4f, 0xfb, 0x1f, 0x9c, 0x5e, 0x40, 0xd7, 0x5d, 0x29, 0xb8, 0xf9, 0x64, 0x4d,
	0x0f, 0x1f, 0xbc, 0xda, 0xc2, 0x06, 0x7e, 0x0d, 0xc1, 0xf2, 0x6c, 0xe1, 0xfe, 0xf6, 0x8b, 0x61,
	0x7a, 0xb0, 0xe1, 0x5f, 0xe6, 0x7f, 0x09, 0xc3, 0xda, 0x91, 0xc1, 0x89, 0x6f, 0xe7, 0xc6, 0x29,
	0xda, 0x04, 0x7d, 0xd9, 0xb3, 0x7f, 0x4c, 0x9f, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x35,
	0xcf, 0xc3, 0x40, 0x09, 0x00, 0x00,
}
