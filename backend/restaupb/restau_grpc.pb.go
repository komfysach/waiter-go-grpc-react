// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package restaupb

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TJLServiceClient is the client API for TJLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TJLServiceClient interface {
}

type tJLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTJLServiceClient(cc grpc.ClientConnInterface) TJLServiceClient {
	return &tJLServiceClient{cc}
}

// TJLServiceServer is the server API for TJLService service.
// All implementations must embed UnimplementedTJLServiceServer
// for forward compatibility
type TJLServiceServer interface {
	mustEmbedUnimplementedTJLServiceServer()
}

// UnimplementedTJLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTJLServiceServer struct {
}

func (UnimplementedTJLServiceServer) mustEmbedUnimplementedTJLServiceServer() {}

// UnsafeTJLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TJLServiceServer will
// result in compilation errors.
type UnsafeTJLServiceServer interface {
	mustEmbedUnimplementedTJLServiceServer()
}

func RegisterTJLServiceServer(s grpc.ServiceRegistrar, srv TJLServiceServer) {
	s.RegisterService(&TJLService_ServiceDesc, srv)
}

// TJLService_ServiceDesc is the grpc.ServiceDesc for TJLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TJLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.TJLService",
	HandlerType: (*TJLServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "backend/restaupb/restau.proto",
}
