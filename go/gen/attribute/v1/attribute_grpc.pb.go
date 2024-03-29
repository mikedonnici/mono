// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package attributev1

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

// AttributeServiceClient is the client API for AttributeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttributeServiceClient interface {
	//  rpc AddAttribute(FetchAttributeRequest) returns (FetchAttributeResponse) {};
	//rpc ChangeAttribute(FetchAttributeRequest) returns (FetchAttributeResponse) {};
	FetchAttribute(ctx context.Context, in *FetchAttributeRequest, opts ...grpc.CallOption) (*FetchAttributeResponse, error)
}

type attributeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributeServiceClient(cc grpc.ClientConnInterface) AttributeServiceClient {
	return &attributeServiceClient{cc}
}

func (c *attributeServiceClient) FetchAttribute(ctx context.Context, in *FetchAttributeRequest, opts ...grpc.CallOption) (*FetchAttributeResponse, error) {
	out := new(FetchAttributeResponse)
	err := c.cc.Invoke(ctx, "/attribute.v1.AttributeService/FetchAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributeServiceServer is the server API for AttributeService service.
// All implementations should embed UnimplementedAttributeServiceServer
// for forward compatibility
type AttributeServiceServer interface {
	//  rpc AddAttribute(FetchAttributeRequest) returns (FetchAttributeResponse) {};
	//rpc ChangeAttribute(FetchAttributeRequest) returns (FetchAttributeResponse) {};
	FetchAttribute(context.Context, *FetchAttributeRequest) (*FetchAttributeResponse, error)
}

// UnimplementedAttributeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAttributeServiceServer struct {
}

func (UnimplementedAttributeServiceServer) FetchAttribute(context.Context, *FetchAttributeRequest) (*FetchAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAttribute not implemented")
}

// UnsafeAttributeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttributeServiceServer will
// result in compilation errors.
type UnsafeAttributeServiceServer interface {
	mustEmbedUnimplementedAttributeServiceServer()
}

func RegisterAttributeServiceServer(s grpc.ServiceRegistrar, srv AttributeServiceServer) {
	s.RegisterService(&AttributeService_ServiceDesc, srv)
}

func _AttributeService_FetchAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).FetchAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attribute.v1.AttributeService/FetchAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).FetchAttribute(ctx, req.(*FetchAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttributeService_ServiceDesc is the grpc.ServiceDesc for AttributeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttributeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attribute.v1.AttributeService",
	HandlerType: (*AttributeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAttribute",
			Handler:    _AttributeService_FetchAttribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attribute/v1/attribute.proto",
}
