// Code generated by protoc-gen-go. DO NOT EDIT.
// source: partitioner/partitioner.proto

package partitioner

import (
	fmt "fmt"
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

type Configuration_Type int32

const (
	// coordinator is in middle of configuration change process
	Configuration_PENDING Configuration_Type = 0
	Configuration_NORMAL  Configuration_Type = 1
)

var Configuration_Type_name = map[int32]string{
	0: "PENDING",
	1: "NORMAL",
}

var Configuration_Type_value = map[string]int32{
	"PENDING": 0,
	"NORMAL":  1,
}

func (x Configuration_Type) String() string {
	return proto.EnumName(Configuration_Type_name, int32(x))
}

func (Configuration_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{0, 0}
}

type Configuration struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term    int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	// ID of the cluster
	Cluster string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// indicates current state of coordinator
	State                string                 `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Workers              []*WorkerConfiguration `protobuf:"bytes,5,rep,name=workers,proto3" json:"workers,omitempty"`
	TotalPartitions      int32                  `protobuf:"varint,6,opt,name=total_partitions,json=totalPartitions,proto3" json:"total_partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Configuration) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Configuration) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Configuration) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Configuration) GetWorkers() []*WorkerConfiguration {
	if m != nil {
		return m.Workers
	}
	return nil
}

func (m *Configuration) GetTotalPartitions() int32 {
	if m != nil {
		return m.TotalPartitions
	}
	return 0
}

type WorkerConfiguration struct {
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Host                 string   `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Partitions           []int32  `protobuf:"varint,7,rep,packed,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerConfiguration) Reset()         { *m = WorkerConfiguration{} }
func (m *WorkerConfiguration) String() string { return proto.CompactTextString(m) }
func (*WorkerConfiguration) ProtoMessage()    {}
func (*WorkerConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{1}
}

func (m *WorkerConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerConfiguration.Unmarshal(m, b)
}
func (m *WorkerConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerConfiguration.Marshal(b, m, deterministic)
}
func (m *WorkerConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerConfiguration.Merge(m, src)
}
func (m *WorkerConfiguration) XXX_Size() int {
	return xxx_messageInfo_WorkerConfiguration.Size(m)
}
func (m *WorkerConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerConfiguration proto.InternalMessageInfo

func (m *WorkerConfiguration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerConfiguration) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *WorkerConfiguration) GetPartitions() []int32 {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type JoinClusterRequest struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term    int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	Cluster string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// unique string for each workers, ex: web-1, web-2
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// used to contact worker, examples:  web-4:9090, 192.168.5.134:9090
	Host                 string   `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinClusterRequest) Reset()         { *m = JoinClusterRequest{} }
func (m *JoinClusterRequest) String() string { return proto.CompactTextString(m) }
func (*JoinClusterRequest) ProtoMessage()    {}
func (*JoinClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{2}
}

func (m *JoinClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinClusterRequest.Unmarshal(m, b)
}
func (m *JoinClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinClusterRequest.Marshal(b, m, deterministic)
}
func (m *JoinClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinClusterRequest.Merge(m, src)
}
func (m *JoinClusterRequest) XXX_Size() int {
	return xxx_messageInfo_JoinClusterRequest.Size(m)
}
func (m *JoinClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinClusterRequest proto.InternalMessageInfo

func (m *JoinClusterRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *JoinClusterRequest) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *JoinClusterRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *JoinClusterRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JoinClusterRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type Empty struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term    int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	// ID of the cluster
	Cluster              string   `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Empty) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Empty) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func init() {
	proto.RegisterEnum("partitioner.Configuration_Type", Configuration_Type_name, Configuration_Type_value)
	proto.RegisterType((*Configuration)(nil), "partitioner.Configuration")
	proto.RegisterType((*WorkerConfiguration)(nil), "partitioner.WorkerConfiguration")
	proto.RegisterType((*JoinClusterRequest)(nil), "partitioner.JoinClusterRequest")
	proto.RegisterType((*Empty)(nil), "partitioner.Empty")
}

func init() { proto.RegisterFile("partitioner/partitioner.proto", fileDescriptor_e94f51d1df456051) }

var fileDescriptor_e94f51d1df456051 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0xbd, 0xb6, 0xfe, 0xe0, 0x11, 0x6d, 0xcd, 0xb4, 0x87, 0xc5, 0xd0, 0x5a, 0xe8, 0xa4,
	0x5e, 0x5c, 0xaa, 0xde, 0x5a, 0x7c, 0x30, 0xaa, 0x31, 0x2d, 0xad, 0x6d, 0x44, 0xa1, 0xf4, 0x54,
	0xe4, 0x7a, 0x9b, 0x2c, 0xb1, 0xb4, 0xca, 0x6a, 0x9d, 0xe0, 0x5b, 0xee, 0xf9, 0xa8, 0x39, 0xe7,
	0x1e, 0xb4, 0xb2, 0x8d, 0x44, 0x44, 0x48, 0x88, 0x6f, 0x33, 0x4f, 0x33, 0x7a, 0x3f, 0xbd, 0x41,
	0xf0, 0x36, 0x8b, 0xa5, 0xe2, 0x8a, 0x8b, 0x94, 0xc9, 0x0f, 0x95, 0x7a, 0x98, 0x49, 0xa1, 0x04,
	0x3a, 0x15, 0xc9, 0xbb, 0x25, 0xf0, 0x22, 0x14, 0xe9, 0x7f, 0x7e, 0xb2, 0x91, 0x71, 0xa1, 0x21,
	0x05, 0xfb, 0x82, 0xc9, 0x9c, 0x8b, 0x94, 0x12, 0x97, 0xf8, 0xdd, 0x68, 0xdf, 0x22, 0x82, 0xa1,
	0x98, 0x4c, 0x68, 0xdb, 0x25, 0xbe, 0x19, 0xe9, 0xba, 0x98, 0xfe, 0xb7, 0xde, 0xe4, 0x8a, 0x49,
	0xda, 0x29, 0xa7, 0x77, 0x2d, 0xbe, 0x01, 0x33, 0x57, 0xb1, 0x62, 0xd4, 0xd0, 0x7a, 0xd9, 0xe0,
	0x67, 0xb0, 0x2f, 0x85, 0x3c, 0x63, 0x32, 0xa7, 0xa6, 0xdb, 0xf1, 0x9d, 0xc0, 0x1d, 0x56, 0x09,
	0x7f, 0xeb, 0x67, 0x35, 0xa0, 0x68, 0xbf, 0x80, 0xef, 0xa1, 0xa7, 0x84, 0x8a, 0xd7, 0x7f, 0x0f,
	0x1b, 0x39, 0xb5, 0x34, 0xcb, 0x2b, 0xad, 0x2f, 0x0e, 0xb2, 0x37, 0x00, 0xe3, 0xd7, 0x36, 0x63,
	0xe8, 0x80, 0xbd, 0x98, 0xcc, 0xbe, 0x7e, 0x9b, 0x4d, 0x7b, 0x2d, 0x04, 0xb0, 0x66, 0xf3, 0xe8,
	0xe7, 0xf8, 0x47, 0x8f, 0x78, 0x7f, 0xe0, 0x75, 0x83, 0x17, 0xbe, 0x84, 0x36, 0x5f, 0xed, 0x88,
	0xdb, 0x7c, 0x55, 0x7c, 0xf2, 0xa9, 0xc8, 0x15, 0x35, 0xb5, 0xa2, 0x6b, 0x7c, 0x07, 0x50, 0x01,
	0xb0, 0xdd, 0x8e, 0x6f, 0x46, 0x15, 0xc5, 0xbb, 0x22, 0x80, 0xdf, 0x05, 0x4f, 0xc3, 0x32, 0x88,
	0x88, 0x9d, 0x6f, 0x58, 0xae, 0x8e, 0x96, 0xeb, 0x23, 0x10, 0xbd, 0x39, 0x98, 0x93, 0x24, 0x53,
	0xdb, 0x63, 0x99, 0x06, 0xd7, 0x04, 0x9c, 0x50, 0x08, 0xb9, 0xe2, 0x69, 0xac, 0x84, 0xc4, 0x11,
	0x74, 0xa7, 0x4c, 0x95, 0xd9, 0x21, 0xd6, 0x4e, 0xa8, 0x8d, 0xfb, 0xfd, 0x9a, 0x56, 0x0b, 0xd9,
	0x6b, 0xe1, 0x08, 0x8c, 0x22, 0x21, 0x1c, 0xd4, 0xa6, 0xee, 0x87, 0xd6, 0x6f, 0x78, 0xb5, 0xd7,
	0x0a, 0x6e, 0x08, 0x58, 0xe5, 0xf5, 0x9e, 0x0b, 0xf2, 0x05, 0xec, 0xfd, 0x7d, 0x1e, 0x18, 0x6c,
	0xc6, 0xc0, 0x00, 0xac, 0x50, 0x24, 0x09, 0x57, 0x8d, 0xc6, 0xcd, 0x3b, 0x1f, 0xc1, 0x1c, 0x2f,
	0x85, 0x7c, 0xc2, 0xca, 0xd2, 0xd2, 0xbf, 0xed, 0xa7, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x9f, 0x24, 0xde, 0xd7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorClient interface {
	GetConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Configuration, error)
	Join(ctx context.Context, in *JoinClusterRequest, opts ...grpc.CallOption) (*Empty, error)
}

type coordinatorClient struct {
	cc *grpc.ClientConn
}

func NewCoordinatorClient(cc *grpc.ClientConn) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) GetConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/partitioner.Coordinator/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) Join(ctx context.Context, in *JoinClusterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/partitioner.Coordinator/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
type CoordinatorServer interface {
	GetConfig(context.Context, *Empty) (*Configuration, error)
	Join(context.Context, *JoinClusterRequest) (*Empty, error)
}

func RegisterCoordinatorServer(s *grpc.Server, srv CoordinatorServer) {
	s.RegisterService(&_Coordinator_serviceDesc, srv)
}

func _Coordinator_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Coordinator/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).GetConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Coordinator/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).Join(ctx, req.(*JoinClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "partitioner.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Coordinator_GetConfig_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Coordinator_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partitioner/partitioner.proto",
}

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	GetConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Configuration, error)
	Request(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Empty, error)
	Commit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Abort(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) GetConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/partitioner.Worker/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Request(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/partitioner.Worker/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Commit(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/partitioner.Worker/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Abort(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/partitioner.Worker/Abort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	GetConfig(context.Context, *Empty) (*Configuration, error)
	Request(context.Context, *Configuration) (*Empty, error)
	Commit(context.Context, *Empty) (*Empty, error)
	Abort(context.Context, *Empty) (*Empty, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Worker/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configuration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Worker/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Request(ctx, req.(*Configuration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Worker/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Commit(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Abort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Abort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Worker/Abort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Abort(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "partitioner.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Worker_GetConfig_Handler,
		},
		{
			MethodName: "Request",
			Handler:    _Worker_Request_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Worker_Commit_Handler,
		},
		{
			MethodName: "Abort",
			Handler:    _Worker_Abort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partitioner/partitioner.proto",
}
