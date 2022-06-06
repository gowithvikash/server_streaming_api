// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: greet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet_Many_Times(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_Greet_Many_TimesClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet_Many_Times(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_Greet_Many_TimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet.GreetService/Greet_Many_Times", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreet_Many_TimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_Greet_Many_TimesClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreet_Many_TimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreet_Many_TimesClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Greet_Many_Times(*GreetRequest, GreetService_Greet_Many_TimesServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Greet_Many_Times(*GreetRequest, GreetService_Greet_Many_TimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Greet_Many_Times not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Greet_Many_Times_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).Greet_Many_Times(m, &greetServiceGreet_Many_TimesServer{stream})
}

type GreetService_Greet_Many_TimesServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type greetServiceGreet_Many_TimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreet_Many_TimesServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet_Many_Times",
			Handler:       _GreetService_Greet_Many_Times_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greet.proto",
}