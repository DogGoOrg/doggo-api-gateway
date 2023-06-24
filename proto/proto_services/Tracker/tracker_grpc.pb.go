// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: doggo-proto/tracker.proto

package Tracker

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

const (
	Tracker_AddPoint_FullMethodName = "/Tracker.Tracker/AddPoint"
)

// TrackerClient is the client API for Tracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackerClient interface {
	AddPoint(ctx context.Context, in *AddPointReq, opts ...grpc.CallOption) (*AddPointRes, error)
}

type trackerClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackerClient(cc grpc.ClientConnInterface) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) AddPoint(ctx context.Context, in *AddPointReq, opts ...grpc.CallOption) (*AddPointRes, error) {
	out := new(AddPointRes)
	err := c.cc.Invoke(ctx, Tracker_AddPoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackerServer is the server API for Tracker service.
// All implementations must embed UnimplementedTrackerServer
// for forward compatibility
type TrackerServer interface {
	AddPoint(context.Context, *AddPointReq) (*AddPointRes, error)
	mustEmbedUnimplementedTrackerServer()
}

// UnimplementedTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedTrackerServer struct {
}

func (UnimplementedTrackerServer) AddPoint(context.Context, *AddPointReq) (*AddPointRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPoint not implemented")
}
func (UnimplementedTrackerServer) mustEmbedUnimplementedTrackerServer() {}

// UnsafeTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackerServer will
// result in compilation errors.
type UnsafeTrackerServer interface {
	mustEmbedUnimplementedTrackerServer()
}

func RegisterTrackerServer(s grpc.ServiceRegistrar, srv TrackerServer) {
	s.RegisterService(&Tracker_ServiceDesc, srv)
}

func _Tracker_AddPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPointReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).AddPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tracker_AddPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).AddPoint(ctx, req.(*AddPointReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Tracker_ServiceDesc is the grpc.ServiceDesc for Tracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Tracker.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPoint",
			Handler:    _Tracker_AddPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doggo-proto/tracker.proto",
}
