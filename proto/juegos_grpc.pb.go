// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/juegos.proto

package juegos

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
	JuegosService_GetJuegosInfo_FullMethodName  = "/juegos.JuegosService/GetJuegosInfo"
	JuegosService_GetJuegosList_FullMethodName  = "/juegos.JuegosService/GetJuegosList"
	JuegosService_AddJuegos_FullMethodName      = "/juegos.JuegosService/AddJuegos"
	JuegosService_GetJuegosByWin_FullMethodName = "/juegos.JuegosService/GetJuegosByWin"
)

// JuegosServiceClient is the client API for JuegosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JuegosServiceClient interface {
	GetJuegosInfo(ctx context.Context, in *JuegosRequest, opts ...grpc.CallOption) (*JuegosResponse, error)
	GetJuegosList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JuegosResponse], error)
	AddJuegos(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[NewJuegosRequest, AddJuegosResponse], error)
	GetJuegosByWin(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[JuegosWinRequest, JuegosResponse], error)
}

type juegosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJuegosServiceClient(cc grpc.ClientConnInterface) JuegosServiceClient {
	return &juegosServiceClient{cc}
}

func (c *juegosServiceClient) GetJuegosInfo(ctx context.Context, in *JuegosRequest, opts ...grpc.CallOption) (*JuegosResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JuegosResponse)
	err := c.cc.Invoke(ctx, JuegosService_GetJuegosInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *juegosServiceClient) GetJuegosList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JuegosResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &JuegosService_ServiceDesc.Streams[0], JuegosService_GetJuegosList_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, JuegosResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JuegosService_GetJuegosListClient = grpc.ServerStreamingClient[JuegosResponse]

func (c *juegosServiceClient) AddJuegos(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[NewJuegosRequest, AddJuegosResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &JuegosService_ServiceDesc.Streams[1], JuegosService_AddJuegos_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NewJuegosRequest, AddJuegosResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JuegosService_AddJuegosClient = grpc.ClientStreamingClient[NewJuegosRequest, AddJuegosResponse]

func (c *juegosServiceClient) GetJuegosByWin(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[JuegosWinRequest, JuegosResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &JuegosService_ServiceDesc.Streams[2], JuegosService_GetJuegosByWin_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[JuegosWinRequest, JuegosResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JuegosService_GetJuegosByWinClient = grpc.BidiStreamingClient[JuegosWinRequest, JuegosResponse]

// JuegosServiceServer is the server API for JuegosService service.
// All implementations must embed UnimplementedJuegosServiceServer
// for forward compatibility.
type JuegosServiceServer interface {
	GetJuegosInfo(context.Context, *JuegosRequest) (*JuegosResponse, error)
	GetJuegosList(*Empty, grpc.ServerStreamingServer[JuegosResponse]) error
	AddJuegos(grpc.ClientStreamingServer[NewJuegosRequest, AddJuegosResponse]) error
	GetJuegosByWin(grpc.BidiStreamingServer[JuegosWinRequest, JuegosResponse]) error
	mustEmbedUnimplementedJuegosServiceServer()
}

// UnimplementedJuegosServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJuegosServiceServer struct{}

func (UnimplementedJuegosServiceServer) GetJuegosInfo(context.Context, *JuegosRequest) (*JuegosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJuegosInfo not implemented")
}
func (UnimplementedJuegosServiceServer) GetJuegosList(*Empty, grpc.ServerStreamingServer[JuegosResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetJuegosList not implemented")
}
func (UnimplementedJuegosServiceServer) AddJuegos(grpc.ClientStreamingServer[NewJuegosRequest, AddJuegosResponse]) error {
	return status.Errorf(codes.Unimplemented, "method AddJuegos not implemented")
}
func (UnimplementedJuegosServiceServer) GetJuegosByWin(grpc.BidiStreamingServer[JuegosWinRequest, JuegosResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetJuegosByWin not implemented")
}
func (UnimplementedJuegosServiceServer) mustEmbedUnimplementedJuegosServiceServer() {}
func (UnimplementedJuegosServiceServer) testEmbeddedByValue()                       {}

// UnsafeJuegosServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JuegosServiceServer will
// result in compilation errors.
type UnsafeJuegosServiceServer interface {
	mustEmbedUnimplementedJuegosServiceServer()
}

func RegisterJuegosServiceServer(s grpc.ServiceRegistrar, srv JuegosServiceServer) {
	// If the following call pancis, it indicates UnimplementedJuegosServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JuegosService_ServiceDesc, srv)
}

func _JuegosService_GetJuegosInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JuegosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JuegosServiceServer).GetJuegosInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JuegosService_GetJuegosInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JuegosServiceServer).GetJuegosInfo(ctx, req.(*JuegosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JuegosService_GetJuegosList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JuegosServiceServer).GetJuegosList(m, &grpc.GenericServerStream[Empty, JuegosResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JuegosService_GetJuegosListServer = grpc.ServerStreamingServer[JuegosResponse]

func _JuegosService_AddJuegos_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JuegosServiceServer).AddJuegos(&grpc.GenericServerStream[NewJuegosRequest, AddJuegosResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JuegosService_AddJuegosServer = grpc.ClientStreamingServer[NewJuegosRequest, AddJuegosResponse]

func _JuegosService_GetJuegosByWin_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JuegosServiceServer).GetJuegosByWin(&grpc.GenericServerStream[JuegosWinRequest, JuegosResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JuegosService_GetJuegosByWinServer = grpc.BidiStreamingServer[JuegosWinRequest, JuegosResponse]

// JuegosService_ServiceDesc is the grpc.ServiceDesc for JuegosService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JuegosService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "juegos.JuegosService",
	HandlerType: (*JuegosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJuegosInfo",
			Handler:    _JuegosService_GetJuegosInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetJuegosList",
			Handler:       _JuegosService_GetJuegosList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddJuegos",
			Handler:       _JuegosService_AddJuegos_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetJuegosByWin",
			Handler:       _JuegosService_GetJuegosByWin_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/juegos.proto",
}