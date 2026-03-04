package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const greeterServiceName = "api.Greeter"

// GreeterServer defines the gRPC server API for Greeter.
// It intentionally uses well-known types to avoid custom protobuf generation.
type GreeterServer interface {
	SayHello(context.Context, *wrapperspb.StringValue) (*wrapperspb.StringValue, error)
	SayHelloStream(*wrapperspb.StringValue, Greeter_SayHelloStreamServer) error
	mustEmbedUnimplementedGreeterServer()
}

type UnimplementedGreeterServer struct{}

func (UnimplementedGreeterServer) SayHello(context.Context, *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func (UnimplementedGreeterServer) SayHelloStream(*wrapperspb.StringValue, Greeter_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}

func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// RegisterGreeterServer registers the Greeter service with the gRPC server.
func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&greeterServiceDesc, srv)
}

type Greeter_SayHelloStreamServer interface {
	Send(*wrapperspb.StringValue) error
	grpc.ServerStream
}

type greeterSayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloStreamServer) Send(m *wrapperspb.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/" + greeterServiceName + "/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	in := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(in); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloStream(in, &greeterSayHelloStreamServer{stream})
}

var greeterServiceDesc = grpc.ServiceDesc{
	ServiceName: greeterServiceName,
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _Greeter_SayHelloStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greeter.proto",
}

