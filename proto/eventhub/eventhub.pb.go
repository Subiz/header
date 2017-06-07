// Code generated by protoc-gen-go.
// source: eventhub/eventhub.proto
// DO NOT EDIT!

/*
Package eventhub is a generated protocol buffer package.

It is generated from these files:
	eventhub/eventhub.proto

It has these top-level messages:
	Empty
	UserRecallRequest
	ChannelRecallRequest
	Registration
	Event
	Events
	Cluster
*/
package eventhub

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UserRecallRequest struct {
	AccountId   string `protobuf:"bytes,1,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
	LastEventId string `protobuf:"bytes,2,opt,name=LastEventId,json=lastEventId" json:"LastEventId,omitempty"`
	UserId      string `protobuf:"bytes,3,opt,name=UserId,json=userId" json:"UserId,omitempty"`
}

func (m *UserRecallRequest) Reset()                    { *m = UserRecallRequest{} }
func (m *UserRecallRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRecallRequest) ProtoMessage()               {}
func (*UserRecallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserRecallRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *UserRecallRequest) GetLastEventId() string {
	if m != nil {
		return m.LastEventId
	}
	return ""
}

func (m *UserRecallRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ChannelRecallRequest struct {
	AccountId   string `protobuf:"bytes,1,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
	LastEventId string `protobuf:"bytes,2,opt,name=LastEventId,json=lastEventId" json:"LastEventId,omitempty"`
	ChannelId   string `protobuf:"bytes,3,opt,name=ChannelId,json=channelId" json:"ChannelId,omitempty"`
}

func (m *ChannelRecallRequest) Reset()                    { *m = ChannelRecallRequest{} }
func (m *ChannelRecallRequest) String() string            { return proto.CompactTextString(m) }
func (*ChannelRecallRequest) ProtoMessage()               {}
func (*ChannelRecallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChannelRecallRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *ChannelRecallRequest) GetLastEventId() string {
	if m != nil {
		return m.LastEventId
	}
	return ""
}

func (m *ChannelRecallRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

type Registration struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
	AgentId   string `protobuf:"bytes,2,opt,name=AgentId,json=agentId" json:"AgentId,omitempty"`
	Event     string `protobuf:"bytes,3,opt,name=Event,json=event" json:"Event,omitempty"`
}

func (m *Registration) Reset()                    { *m = Registration{} }
func (m *Registration) String() string            { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()               {}
func (*Registration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Registration) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Registration) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *Registration) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

type Event struct {
	Id      string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=Payload,json=payload" json:"Payload,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type Events struct {
	Events []*Event `protobuf:"bytes,1,rep,name=Events,json=events" json:"Events,omitempty"`
}

func (m *Events) Reset()                    { *m = Events{} }
func (m *Events) String() string            { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()               {}
func (*Events) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Events) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Cluster struct {
	Service string   `protobuf:"bytes,1,opt,name=Service,json=service" json:"Service,omitempty"`
	Hosts   []string `protobuf:"bytes,2,rep,name=Hosts,json=hosts" json:"Hosts,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Cluster) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Cluster) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "eventhub.Empty")
	proto.RegisterType((*UserRecallRequest)(nil), "eventhub.UserRecallRequest")
	proto.RegisterType((*ChannelRecallRequest)(nil), "eventhub.ChannelRecallRequest")
	proto.RegisterType((*Registration)(nil), "eventhub.Registration")
	proto.RegisterType((*Event)(nil), "eventhub.Event")
	proto.RegisterType((*Events)(nil), "eventhub.Events")
	proto.RegisterType((*Cluster)(nil), "eventhub.Cluster")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventHub service

type EventHubClient interface {
	Subscribe(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error)
	Unsubscribe(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error)
	RecallFromChannel(ctx context.Context, in *ChannelRecallRequest, opts ...grpc.CallOption) (*Events, error)
	RecallFromUser(ctx context.Context, in *UserRecallRequest, opts ...grpc.CallOption) (*Events, error)
	UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Empty, error)
}

type eventHubClient struct {
	cc *grpc.ClientConn
}

func NewEventHubClient(cc *grpc.ClientConn) EventHubClient {
	return &eventHubClient{cc}
}

func (c *eventHubClient) Subscribe(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/eventhub.EventHub/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventHubClient) Unsubscribe(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/eventhub.EventHub/Unsubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventHubClient) RecallFromChannel(ctx context.Context, in *ChannelRecallRequest, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := grpc.Invoke(ctx, "/eventhub.EventHub/RecallFromChannel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventHubClient) RecallFromUser(ctx context.Context, in *UserRecallRequest, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := grpc.Invoke(ctx, "/eventhub.EventHub/RecallFromUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventHubClient) UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/eventhub.EventHub/UpdateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventHub service

type EventHubServer interface {
	Subscribe(context.Context, *Registration) (*Empty, error)
	Unsubscribe(context.Context, *Registration) (*Empty, error)
	RecallFromChannel(context.Context, *ChannelRecallRequest) (*Events, error)
	RecallFromUser(context.Context, *UserRecallRequest) (*Events, error)
	UpdateCluster(context.Context, *Cluster) (*Empty, error)
}

func RegisterEventHubServer(s *grpc.Server, srv EventHubServer) {
	s.RegisterService(&_EventHub_serviceDesc, srv)
}

func _EventHub_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventHubServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventhub.EventHub/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventHubServer).Subscribe(ctx, req.(*Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventHub_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventHubServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventhub.EventHub/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventHubServer).Unsubscribe(ctx, req.(*Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventHub_RecallFromChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelRecallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventHubServer).RecallFromChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventhub.EventHub/RecallFromChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventHubServer).RecallFromChannel(ctx, req.(*ChannelRecallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventHub_RecallFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRecallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventHubServer).RecallFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventhub.EventHub/RecallFromUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventHubServer).RecallFromUser(ctx, req.(*UserRecallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventHub_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventHubServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventhub.EventHub/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventHubServer).UpdateCluster(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventHub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventhub.EventHub",
	HandlerType: (*EventHubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _EventHub_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _EventHub_Unsubscribe_Handler,
		},
		{
			MethodName: "RecallFromChannel",
			Handler:    _EventHub_RecallFromChannel_Handler,
		},
		{
			MethodName: "RecallFromUser",
			Handler:    _EventHub_RecallFromUser_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _EventHub_UpdateCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventhub/eventhub.proto",
}

func init() { proto.RegisterFile("eventhub/eventhub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0x61, 0x84, 0x8d, 0xbd, 0x53, 0x94, 0x86, 0xe0, 0x82, 0xc4, 0x2c, 0xbd, 0xc8, 0x09,
	0x03, 0x46, 0xa3, 0x47, 0x42, 0x50, 0x48, 0x3c, 0x98, 0x11, 0x6e, 0x5e, 0xba, 0xad, 0x81, 0xc5,
	0xb1, 0xcd, 0xb5, 0x23, 0xe1, 0xd3, 0xf9, 0xd5, 0xcc, 0xba, 0xbf, 0x51, 0x12, 0x13, 0xe3, 0xad,
	0xcf, 0xd3, 0xf6, 0xf7, 0xbc, 0x7b, 0xfb, 0x0e, 0x2e, 0xe8, 0x9e, 0xfa, 0x7c, 0x1b, 0x5b, 0x37,
	0xf9, 0x62, 0x14, 0x46, 0x01, 0x0f, 0x50, 0x2b, 0xd7, 0x58, 0x81, 0xe6, 0x7c, 0x17, 0xf2, 0x03,
	0x7e, 0x87, 0xce, 0x9a, 0xd1, 0xc8, 0xa4, 0x36, 0xf1, 0x3c, 0x93, 0x7e, 0xc4, 0x94, 0x71, 0x34,
	0x00, 0x75, 0x6a, 0xdb, 0x41, 0xec, 0xf3, 0xa5, 0xa3, 0xd7, 0x8d, 0xfa, 0x50, 0x35, 0x55, 0x92,
	0x1b, 0xc8, 0x00, 0xed, 0x85, 0x30, 0x3e, 0x4f, 0x58, 0x4b, 0x47, 0x97, 0xc4, 0xbe, 0xe6, 0x95,
	0x16, 0xea, 0x81, 0x9c, 0x40, 0x97, 0x8e, 0xde, 0x10, 0x9b, 0x72, 0x2c, 0x14, 0xe6, 0xd0, 0x9d,
	0x6d, 0x89, 0xef, 0x53, 0xef, 0x7f, 0xf3, 0x06, 0xa0, 0x66, 0xdc, 0x22, 0x52, 0xb5, 0x73, 0x03,
	0xbf, 0xc1, 0x89, 0x49, 0x37, 0x2e, 0xe3, 0x11, 0xe1, 0x6e, 0xe0, 0xff, 0x92, 0xa6, 0x83, 0x32,
	0xdd, 0x54, 0x93, 0x14, 0x92, 0x4a, 0xd4, 0x85, 0xa6, 0x08, 0xcc, 0x12, 0x9a, 0xa2, 0x99, 0x78,
	0x9c, 0xb9, 0xa8, 0x0d, 0x52, 0xc1, 0x93, 0x5c, 0x01, 0x7a, 0x25, 0x07, 0x2f, 0x20, 0x05, 0x28,
	0x4c, 0x25, 0x1e, 0x83, 0x2c, 0xae, 0x30, 0x74, 0x9d, 0xaf, 0xf4, 0xba, 0xd1, 0x18, 0x6a, 0x93,
	0xb3, 0x51, 0xf1, 0x62, 0xc2, 0x37, 0x65, 0xa1, 0x19, 0x7e, 0x04, 0x65, 0xe6, 0xc5, 0x8c, 0xd3,
	0x28, 0xe1, 0xae, 0x68, 0xb4, 0x77, 0x6d, 0x9a, 0x85, 0x29, 0x2c, 0x95, 0x49, 0x81, 0x8b, 0x80,
	0x71, 0xa6, 0x4b, 0x46, 0x23, 0x29, 0x70, 0x9b, 0x88, 0xc9, 0xa7, 0x04, 0x2d, 0x01, 0x5b, 0xc4,
	0x16, 0xba, 0x07, 0x75, 0x15, 0x5b, 0xcc, 0x8e, 0x5c, 0x8b, 0xa2, 0x5e, 0x99, 0x56, 0x6d, 0x50,
	0xbf, 0x5a, 0x85, 0x18, 0x92, 0x1a, 0x7a, 0x00, 0x6d, 0xed, 0xb3, 0xbf, 0xdc, 0x7c, 0x86, 0x4e,
	0xfa, 0xd8, 0x4f, 0x51, 0xb0, 0xcb, 0x5e, 0x09, 0x5d, 0x95, 0xe7, 0x8e, 0x0d, 0x44, 0xff, 0xfc,
	0x5b, 0x1f, 0x18, 0xae, 0xa1, 0x29, 0xb4, 0x4b, 0x50, 0x32, 0x5e, 0xe8, 0xb2, 0x3c, 0xf5, 0x63,
	0x86, 0x8f, 0x22, 0xee, 0xe0, 0x74, 0x1d, 0x3a, 0x84, 0xd3, 0xbc, 0x97, 0x9d, 0x4a, 0x1d, 0xa9,
	0x75, 0xe4, 0x13, 0x2c, 0x59, 0xfc, 0x3d, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0xab,
	0xbb, 0x1e, 0x58, 0x03, 0x00, 0x00,
}