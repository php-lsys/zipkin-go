// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testing

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BaggageServiceClient is the client API for BaggageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaggageServiceClient interface {
	Handler1(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Handler2(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type baggageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaggageServiceClient(cc grpc.ClientConnInterface) BaggageServiceClient {
	return &baggageServiceClient{cc}
}

func (c *baggageServiceClient) Handler1(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zipkin.testing.BaggageService/Handler1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baggageServiceClient) Handler2(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/zipkin.testing.BaggageService/Handler2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaggageServiceServer is the server API for BaggageService service.
// All implementations must embed UnimplementedBaggageServiceServer
// for forward compatibility
type BaggageServiceServer interface {
	Handler1(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Handler2(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedBaggageServiceServer()
}

// UnimplementedBaggageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBaggageServiceServer struct {
}

func (*UnimplementedBaggageServiceServer) Handler1(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handler1 not implemented")
}
func (*UnimplementedBaggageServiceServer) Handler2(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handler2 not implemented")
}
func (*UnimplementedBaggageServiceServer) mustEmbedUnimplementedBaggageServiceServer() {}

func RegisterBaggageServiceServer(s *grpc.Server, srv BaggageServiceServer) {
	s.RegisterService(&_BaggageService_serviceDesc, srv)
}

func _BaggageService_Handler1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaggageServiceServer).Handler1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipkin.testing.BaggageService/Handler1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaggageServiceServer).Handler1(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaggageService_Handler2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaggageServiceServer).Handler2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zipkin.testing.BaggageService/Handler2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaggageServiceServer).Handler2(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BaggageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zipkin.testing.BaggageService",
	HandlerType: (*BaggageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handler1",
			Handler:    _BaggageService_Handler1_Handler,
		},
		{
			MethodName: "Handler2",
			Handler:    _BaggageService_Handler2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/testing/baggage.proto",
}