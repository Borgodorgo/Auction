// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: ReplicationService.proto

package Replication

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReplicationService_ReplicateBid_FullMethodName      = "/replicationservice.ReplicationService/ReplicateBid"
	ReplicationService_ConfirmLeader_FullMethodName     = "/replicationservice.ReplicationService/ConfirmLeader"
	ReplicationService_PropagateToLeader_FullMethodName = "/replicationservice.ReplicationService/PropagateToLeader"
)

// ReplicationServiceClient is the client API for ReplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// P2PNetwork
type ReplicationServiceClient interface {
	ReplicateBid(ctx context.Context, in *NewBid, opts ...grpc.CallOption) (*Response, error)
	ConfirmLeader(ctx context.Context, in *NewLeader, opts ...grpc.CallOption) (*Response, error)
	PropagateToLeader(ctx context.Context, in *NewBid, opts ...grpc.CallOption) (*Response, error)
}

type replicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationServiceClient(cc grpc.ClientConnInterface) ReplicationServiceClient {
	return &replicationServiceClient{cc}
}

func (c *replicationServiceClient) ReplicateBid(ctx context.Context, in *NewBid, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, ReplicationService_ReplicateBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) ConfirmLeader(ctx context.Context, in *NewLeader, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, ReplicationService_ConfirmLeader_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) PropagateToLeader(ctx context.Context, in *NewBid, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, ReplicationService_PropagateToLeader_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServiceServer is the server API for ReplicationService service.
// All implementations must embed UnimplementedReplicationServiceServer
// for forward compatibility.
//
// P2PNetwork
type ReplicationServiceServer interface {
	ReplicateBid(context.Context, *NewBid) (*Response, error)
	ConfirmLeader(context.Context, *NewLeader) (*Response, error)
	PropagateToLeader(context.Context, *NewBid) (*Response, error)
	mustEmbedUnimplementedReplicationServiceServer()
}

// UnimplementedReplicationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReplicationServiceServer struct{}

func (UnimplementedReplicationServiceServer) ReplicateBid(context.Context, *NewBid) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateBid not implemented")
}
func (UnimplementedReplicationServiceServer) ConfirmLeader(context.Context, *NewLeader) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmLeader not implemented")
}
func (UnimplementedReplicationServiceServer) PropagateToLeader(context.Context, *NewBid) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropagateToLeader not implemented")
}
func (UnimplementedReplicationServiceServer) mustEmbedUnimplementedReplicationServiceServer() {}
func (UnimplementedReplicationServiceServer) testEmbeddedByValue()                            {}

// UnsafeReplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServiceServer will
// result in compilation errors.
type UnsafeReplicationServiceServer interface {
	mustEmbedUnimplementedReplicationServiceServer()
}

func RegisterReplicationServiceServer(s grpc.ServiceRegistrar, srv ReplicationServiceServer) {
	// If the following call pancis, it indicates UnimplementedReplicationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReplicationService_ServiceDesc, srv)
}

func _ReplicationService_ReplicateBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).ReplicateBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_ReplicateBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).ReplicateBid(ctx, req.(*NewBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_ConfirmLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewLeader)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).ConfirmLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_ConfirmLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).ConfirmLeader(ctx, req.(*NewLeader))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_PropagateToLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).PropagateToLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationService_PropagateToLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).PropagateToLeader(ctx, req.(*NewBid))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplicationService_ServiceDesc is the grpc.ServiceDesc for ReplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replicationservice.ReplicationService",
	HandlerType: (*ReplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplicateBid",
			Handler:    _ReplicationService_ReplicateBid_Handler,
		},
		{
			MethodName: "ConfirmLeader",
			Handler:    _ReplicationService_ConfirmLeader_Handler,
		},
		{
			MethodName: "PropagateToLeader",
			Handler:    _ReplicationService_PropagateToLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ReplicationService.proto",
}
