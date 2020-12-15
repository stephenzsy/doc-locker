// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package configurations

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ConfigurationsServiceClient is the client API for ConfigurationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigurationsServiceClient interface {
	SiteConfigurations(ctx context.Context, in *SiteConfigurationsRequest, opts ...grpc.CallOption) (*SiteConfigurationsResponse, error)
}

type configurationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationsServiceClient(cc grpc.ClientConnInterface) ConfigurationsServiceClient {
	return &configurationsServiceClient{cc}
}

func (c *configurationsServiceClient) SiteConfigurations(ctx context.Context, in *SiteConfigurationsRequest, opts ...grpc.CallOption) (*SiteConfigurationsResponse, error) {
	out := new(SiteConfigurationsResponse)
	err := c.cc.Invoke(ctx, "/doclocker.configurations.ConfigurationsService/SiteConfigurations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationsServiceServer is the server API for ConfigurationsService service.
// All implementations must embed UnimplementedConfigurationsServiceServer
// for forward compatibility
type ConfigurationsServiceServer interface {
	SiteConfigurations(context.Context, *SiteConfigurationsRequest) (*SiteConfigurationsResponse, error)
	mustEmbedUnimplementedConfigurationsServiceServer()
}

// UnimplementedConfigurationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigurationsServiceServer struct {
}

func (UnimplementedConfigurationsServiceServer) SiteConfigurations(context.Context, *SiteConfigurationsRequest) (*SiteConfigurationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SiteConfigurations not implemented")
}
func (UnimplementedConfigurationsServiceServer) mustEmbedUnimplementedConfigurationsServiceServer() {}

// UnsafeConfigurationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigurationsServiceServer will
// result in compilation errors.
type UnsafeConfigurationsServiceServer interface {
	mustEmbedUnimplementedConfigurationsServiceServer()
}

func RegisterConfigurationsServiceServer(s grpc.ServiceRegistrar, srv ConfigurationsServiceServer) {
	s.RegisterService(&_ConfigurationsService_serviceDesc, srv)
}

func _ConfigurationsService_SiteConfigurations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteConfigurationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationsServiceServer).SiteConfigurations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/doclocker.configurations.ConfigurationsService/SiteConfigurations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationsServiceServer).SiteConfigurations(ctx, req.(*SiteConfigurationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigurationsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "doclocker.configurations.ConfigurationsService",
	HandlerType: (*ConfigurationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SiteConfigurations",
			Handler:    _ConfigurationsService_SiteConfigurations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configurations.proto",
}
