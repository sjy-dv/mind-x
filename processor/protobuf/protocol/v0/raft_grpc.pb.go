// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

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

// RaftTransportClient is the client API for RaftTransport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftTransportClient interface {
	Receive(ctx context.Context, in *RaftMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type raftTransportClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftTransportClient(cc grpc.ClientConnInterface) RaftTransportClient {
	return &raftTransportClient{cc}
}

func (c *raftTransportClient) Receive(ctx context.Context, in *RaftMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/mindxv.v0.RaftTransport/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftTransportServer is the server API for RaftTransport service.
// All implementations should embed UnimplementedRaftTransportServer
// for forward compatibility
type RaftTransportServer interface {
	Receive(context.Context, *RaftMessage) (*EmptyMessage, error)
}

// UnimplementedRaftTransportServer should be embedded to have forward compatible implementations.
type UnimplementedRaftTransportServer struct {
}

func (UnimplementedRaftTransportServer) Receive(context.Context, *RaftMessage) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}

// UnsafeRaftTransportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftTransportServer will
// result in compilation errors.
type UnsafeRaftTransportServer interface {
	mustEmbedUnimplementedRaftTransportServer()
}

func RegisterRaftTransportServer(s grpc.ServiceRegistrar, srv RaftTransportServer) {
	s.RegisterService(&RaftTransport_ServiceDesc, srv)
}

func _RaftTransport_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftTransportServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindxv.v0.RaftTransport/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftTransportServer).Receive(ctx, req.(*RaftMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftTransport_ServiceDesc is the grpc.ServiceDesc for RaftTransport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftTransport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mindxv.v0.RaftTransport",
	HandlerType: (*RaftTransportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _RaftTransport_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v0/raft.proto",
}