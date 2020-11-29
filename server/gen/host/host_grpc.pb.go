// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package host

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostServiceClient interface {
	Status(ctx context.Context, in *HostStatusRequest, opts ...grpc.CallOption) (*HostStatusResponse, error)
}

type hostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostServiceClient(cc grpc.ClientConnInterface) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) Status(ctx context.Context, in *HostStatusRequest, opts ...grpc.CallOption) (*HostStatusResponse, error) {
	out := new(HostStatusResponse)
	err := c.cc.Invoke(ctx, "/doclocker.host.HostService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
// All implementations must embed UnimplementedHostServiceServer
// for forward compatibility
type HostServiceServer interface {
	Status(context.Context, *HostStatusRequest) (*HostStatusResponse, error)
	mustEmbedUnimplementedHostServiceServer()
}

// UnimplementedHostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHostServiceServer struct {
}

func (UnimplementedHostServiceServer) Status(context.Context, *HostStatusRequest) (*HostStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedHostServiceServer) mustEmbedUnimplementedHostServiceServer() {}

// UnsafeHostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostServiceServer will
// result in compilation errors.
type UnsafeHostServiceServer interface {
	mustEmbedUnimplementedHostServiceServer()
}

func RegisterHostServiceServer(s grpc.ServiceRegistrar, srv HostServiceServer) {
	s.RegisterService(&_HostService_serviceDesc, srv)
}

func _HostService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doclocker.host.HostService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).Status(ctx, req.(*HostStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "doclocker.host.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _HostService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "host.proto",
}
