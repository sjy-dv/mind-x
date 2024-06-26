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

// NodesManagerClient is the client API for NodesManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodesManagerClient interface {
	ListNodes(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (NodesManager_ListNodesClient, error)
	AddNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (NodesManager_AddNodeClient, error)
	RemoveNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyMessage, error)
	LoadInfo(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NodeLoadInfo, error)
}

type nodesManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewNodesManagerClient(cc grpc.ClientConnInterface) NodesManagerClient {
	return &nodesManagerClient{cc}
}

func (c *nodesManagerClient) ListNodes(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (NodesManager_ListNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodesManager_ServiceDesc.Streams[0], "/mindxv.v0.NodesManager/ListNodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodesManagerListNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodesManager_ListNodesClient interface {
	Recv() (*Node, error)
	grpc.ClientStream
}

type nodesManagerListNodesClient struct {
	grpc.ClientStream
}

func (x *nodesManagerListNodesClient) Recv() (*Node, error) {
	m := new(Node)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodesManagerClient) AddNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (NodesManager_AddNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodesManager_ServiceDesc.Streams[1], "/mindxv.v0.NodesManager/AddNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodesManagerAddNodeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodesManager_AddNodeClient interface {
	Recv() (*Node, error)
	grpc.ClientStream
}

type nodesManagerAddNodeClient struct {
	grpc.ClientStream
}

func (x *nodesManagerAddNodeClient) Recv() (*Node, error) {
	m := new(Node)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodesManagerClient) RemoveNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/mindxv.v0.NodesManager/RemoveNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesManagerClient) LoadInfo(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NodeLoadInfo, error) {
	out := new(NodeLoadInfo)
	err := c.cc.Invoke(ctx, "/mindxv.v0.NodesManager/LoadInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodesManagerServer is the server API for NodesManager service.
// All implementations should embed UnimplementedNodesManagerServer
// for forward compatibility
type NodesManagerServer interface {
	ListNodes(*EmptyMessage, NodesManager_ListNodesServer) error
	AddNode(*Node, NodesManager_AddNodeServer) error
	RemoveNode(context.Context, *Node) (*EmptyMessage, error)
	LoadInfo(context.Context, *EmptyMessage) (*NodeLoadInfo, error)
}

// UnimplementedNodesManagerServer should be embedded to have forward compatible implementations.
type UnimplementedNodesManagerServer struct {
}

func (UnimplementedNodesManagerServer) ListNodes(*EmptyMessage, NodesManager_ListNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedNodesManagerServer) AddNode(*Node, NodesManager_AddNodeServer) error {
	return status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedNodesManagerServer) RemoveNode(context.Context, *Node) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedNodesManagerServer) LoadInfo(context.Context, *EmptyMessage) (*NodeLoadInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadInfo not implemented")
}

// UnsafeNodesManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodesManagerServer will
// result in compilation errors.
type UnsafeNodesManagerServer interface {
	mustEmbedUnimplementedNodesManagerServer()
}

func RegisterNodesManagerServer(s grpc.ServiceRegistrar, srv NodesManagerServer) {
	s.RegisterService(&NodesManager_ServiceDesc, srv)
}

func _NodesManager_ListNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodesManagerServer).ListNodes(m, &nodesManagerListNodesServer{stream})
}

type NodesManager_ListNodesServer interface {
	Send(*Node) error
	grpc.ServerStream
}

type nodesManagerListNodesServer struct {
	grpc.ServerStream
}

func (x *nodesManagerListNodesServer) Send(m *Node) error {
	return x.ServerStream.SendMsg(m)
}

func _NodesManager_AddNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Node)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodesManagerServer).AddNode(m, &nodesManagerAddNodeServer{stream})
}

type NodesManager_AddNodeServer interface {
	Send(*Node) error
	grpc.ServerStream
}

type nodesManagerAddNodeServer struct {
	grpc.ServerStream
}

func (x *nodesManagerAddNodeServer) Send(m *Node) error {
	return x.ServerStream.SendMsg(m)
}

func _NodesManager_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesManagerServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindxv.v0.NodesManager/RemoveNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesManagerServer).RemoveNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodesManager_LoadInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesManagerServer).LoadInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindxv.v0.NodesManager/LoadInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesManagerServer).LoadInfo(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// NodesManager_ServiceDesc is the grpc.ServiceDesc for NodesManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodesManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mindxv.v0.NodesManager",
	HandlerType: (*NodesManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveNode",
			Handler:    _NodesManager_RemoveNode_Handler,
		},
		{
			MethodName: "LoadInfo",
			Handler:    _NodesManager_LoadInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListNodes",
			Handler:       _NodesManager_ListNodes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddNode",
			Handler:       _NodesManager_AddNode_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v0/cluster.proto",
}
