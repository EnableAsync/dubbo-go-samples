// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

import (
	"dubbo.apache.org/dubbo-go/v3/protocol"
	dgrpc "dubbo.apache.org/dubbo-go/v3/protocol/grpc"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "main.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "main.HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d,
	0xcc, 0xcc, 0x53, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0xc9, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x4a, 0x6a, 0x5c, 0x5c, 0x50, 0x35, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9,
	0xc5, 0xc5, 0x89, 0xe9, 0x30, 0x45, 0x30, 0xae, 0xd1, 0x65, 0x46, 0x2e, 0x76, 0xf7, 0xa2, 0xd4,
	0xd4, 0x92, 0xd4, 0x22, 0x21, 0x17, 0x2e, 0xb1, 0xe0, 0xc4, 0x4a, 0xb0, 0xb6, 0x90, 0xf2, 0xfc,
	0xe0, 0xcc, 0x94, 0xd4, 0xe2, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0x21, 0x21, 0x3d, 0x90, 0xc5,
	0x7a, 0xc8, 0xb6, 0x4a, 0x09, 0xa0, 0x88, 0x15, 0xe4, 0x54, 0x2a, 0x31, 0x68, 0x30, 0x1a, 0x30,
	0x0a, 0x39, 0x70, 0x89, 0xc0, 0x4c, 0x71, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x21, 0xd5, 0x0c, 0x64,
	0x13, 0x82, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x48, 0x33, 0xc1, 0x80, 0xd1, 0xc9, 0x8c, 0x4b, 0x3a,
	0x33, 0x5f, 0x2f, 0xbd, 0xa8, 0x20, 0x59, 0x2f, 0xb5, 0x22, 0x31, 0xb7, 0x20, 0x27, 0xb5, 0x58,
	0x0f, 0x11, 0x98, 0x4e, 0xfc, 0x60, 0xe5, 0xe1, 0x20, 0x76, 0x00, 0x28, 0x5c, 0x03, 0x18, 0x17,
	0x31, 0x31, 0x7b, 0xf8, 0x84, 0x27, 0xb1, 0x81, 0x83, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x0c, 0xb8, 0x0a, 0x7a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHelloTwoSidesStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloTwoSidesStreamClient, error)
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error)
	SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHelloTwoSidesStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloTwoSidesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/main.Greeter/SayHelloTwoSidesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloTwoSidesStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloTwoSidesStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloTwoSidesStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloTwoSidesStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloTwoSidesStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[1], "/main.Greeter/SayHelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloClientStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[2], "/main.Greeter/SayHelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloServerStreamClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloServerStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHelloTwoSidesStream(Greeter_SayHelloTwoSidesStreamServer) error
	SayHelloClientStream(Greeter_SayHelloClientStreamServer) error
	SayHelloServerStream(*HelloRequest, Greeter_SayHelloServerStreamServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHelloTwoSidesStream(srv Greeter_SayHelloTwoSidesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloTwoSidesStream not implemented")
}
func (*UnimplementedGreeterServer) SayHelloClientStream(srv Greeter_SayHelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (*UnimplementedGreeterServer) SayHelloServerStream(req *HelloRequest, srv Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHelloTwoSidesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloTwoSidesStream(&greeterSayHelloTwoSidesStreamServer{stream})
}

type Greeter_SayHelloTwoSidesStreamServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloTwoSidesStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloTwoSidesStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloTwoSidesStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloClientStream(&greeterSayHelloClientStreamServer{stream})
}

type Greeter_SayHelloClientStreamServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloClientStreamServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloServerStream(m, &greeterSayHelloServerStreamServer{stream})
}

type Greeter_SayHelloServerStreamServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloServerStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloTwoSidesStream",
			Handler:       _Greeter_SayHelloTwoSidesStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _Greeter_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _Greeter_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

// GreeterClientImpl is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClientImpl struct {
	// Sends a greeting
	SayHelloTwoSidesStream func(ctx context.Context) (Greeter_SayHelloTwoSidesStreamClient, error)
	SayHelloClientStream   func(ctx context.Context) (Greeter_SayHelloClientStreamClient, error)
	SayHelloServerStream   func(ctx context.Context, in *HelloRequest) (Greeter_SayHelloServerStreamClient, error)
}

func (c *GreeterClientImpl) Reference() string {
	return "greeterImpl"
}

func (c *GreeterClientImpl) GetDubboStub(cc *grpc.ClientConn) GreeterClient {
	return NewGreeterClient(cc)
}

type GreeterProviderBase struct {
	proxyImpl protocol.Invoker
}

func (s *GreeterProviderBase) SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *GreeterProviderBase) GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func _DUBBO_Greeter_SayHelloTwoSidesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	_, ok := srv.(dgrpc.DubboGrpcService)
	invo := invocation.NewRPCInvocation("SayHelloTwoSidesStream", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	return srv.(GreeterServer).SayHelloTwoSidesStream(&greeterSayHelloTwoSidesStreamServer{stream})
}

func _DUBBO_Greeter_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	_, ok := srv.(dgrpc.DubboGrpcService)
	invo := invocation.NewRPCInvocation("SayHelloClientStream", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	return srv.(GreeterServer).SayHelloClientStream(&greeterSayHelloClientStreamServer{stream})
}

func _DUBBO_Greeter_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	_, ok := srv.(dgrpc.DubboGrpcService)
	invo := invocation.NewRPCInvocation("SayHelloServerStream", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloServerStream(m, &greeterSayHelloServerStreamServer{stream})
}

func (s *GreeterProviderBase) ServiceDesc() *grpc.ServiceDesc {
	return &grpc.ServiceDesc{
		ServiceName: "main.Greeter",
		HandlerType: (*GreeterServer)(nil),
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "SayHelloTwoSidesStream",
				Handler:       _DUBBO_Greeter_SayHelloTwoSidesStream_Handler,
				ServerStreams: true,
				ClientStreams: true,
			},
			{
				StreamName:    "SayHelloClientStream",
				Handler:       _DUBBO_Greeter_SayHelloClientStream_Handler,
				ClientStreams: true,
			},
			{
				StreamName:    "SayHelloServerStream",
				Handler:       _DUBBO_Greeter_SayHelloServerStream_Handler,
				ServerStreams: true,
			},
		},
		Metadata: "helloworld.proto",
	}
}
