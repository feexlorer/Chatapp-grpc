// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/chatapp.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatappServiceClient is the client API for ChatappService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatappServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*empty.Empty, error)
	ReceiveMessage(ctx context.Context, opts ...grpc.CallOption) (ChatappService_ReceiveMessageClient, error)
}

type chatappServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatappServiceClient(cc grpc.ClientConnInterface) ChatappServiceClient {
	return &chatappServiceClient{cc}
}

func (c *chatappServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chatapp.ChatappService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatappServiceClient) ReceiveMessage(ctx context.Context, opts ...grpc.CallOption) (ChatappService_ReceiveMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatappService_ServiceDesc.Streams[0], "/chatapp.ChatappService/ReceiveMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatappServiceReceiveMessageClient{stream}
	return x, nil
}

type ChatappService_ReceiveMessageClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatappServiceReceiveMessageClient struct {
	grpc.ClientStream
}

func (x *chatappServiceReceiveMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatappServiceReceiveMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatappServiceServer is the server API for ChatappService service.
// All implementations must embed UnimplementedChatappServiceServer
// for forward compatibility
type ChatappServiceServer interface {
	SendMessage(context.Context, *Message) (*empty.Empty, error)
	ReceiveMessage(ChatappService_ReceiveMessageServer) error
	mustEmbedUnimplementedChatappServiceServer()
}

// UnimplementedChatappServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatappServiceServer struct {
}

func (UnimplementedChatappServiceServer) SendMessage(context.Context, *Message) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatappServiceServer) ReceiveMessage(ChatappService_ReceiveMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
}
func (UnimplementedChatappServiceServer) mustEmbedUnimplementedChatappServiceServer() {}

// UnsafeChatappServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatappServiceServer will
// result in compilation errors.
type UnsafeChatappServiceServer interface {
	mustEmbedUnimplementedChatappServiceServer()
}

func RegisterChatappServiceServer(s grpc.ServiceRegistrar, srv ChatappServiceServer) {
	s.RegisterService(&ChatappService_ServiceDesc, srv)
}

func _ChatappService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatappServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatapp.ChatappService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatappServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatappService_ReceiveMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatappServiceServer).ReceiveMessage(&chatappServiceReceiveMessageServer{stream})
}

type ChatappService_ReceiveMessageServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatappServiceReceiveMessageServer struct {
	grpc.ServerStream
}

func (x *chatappServiceReceiveMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatappServiceReceiveMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatappService_ServiceDesc is the grpc.ServiceDesc for ChatappService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatappService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatapp.ChatappService",
	HandlerType: (*ChatappServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatappService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMessage",
			Handler:       _ChatappService_ReceiveMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chatapp.proto",
}