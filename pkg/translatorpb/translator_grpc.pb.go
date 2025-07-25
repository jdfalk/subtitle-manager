// file: proto/translator.proto
// version: 1.2.0
// guid: 70eca88d-31fe-4044-8b76-a22bd3c9e0bd

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: proto/translator.proto

package translatorpb

import (
	context "context"
	configpb "github.com/jdfalk/subtitle-manager/pkg/configpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TranslatorService_Translate_FullMethodName = "/gcommon.v1.translator.TranslatorService/Translate"
	TranslatorService_GetConfig_FullMethodName = "/gcommon.v1.translator.TranslatorService/GetConfig"
	TranslatorService_SetConfig_FullMethodName = "/gcommon.v1.translator.TranslatorService/SetConfig"
)

// TranslatorServiceClient is the client API for TranslatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslatorServiceClient interface {
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*configpb.SubtitleManagerConfig, error)
	SetConfig(ctx context.Context, in *configpb.SubtitleManagerConfig, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type translatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslatorServiceClient(cc grpc.ClientConnInterface) TranslatorServiceClient {
	return &translatorServiceClient{cc}
}

func (c *translatorServiceClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranslateResponse)
	err := c.cc.Invoke(ctx, TranslatorService_Translate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorServiceClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*configpb.SubtitleManagerConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(configpb.SubtitleManagerConfig)
	err := c.cc.Invoke(ctx, TranslatorService_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translatorServiceClient) SetConfig(ctx context.Context, in *configpb.SubtitleManagerConfig, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TranslatorService_SetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServiceServer is the server API for TranslatorService service.
// All implementations must embed UnimplementedTranslatorServiceServer
// for forward compatibility.
type TranslatorServiceServer interface {
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
	GetConfig(context.Context, *emptypb.Empty) (*configpb.SubtitleManagerConfig, error)
	SetConfig(context.Context, *configpb.SubtitleManagerConfig) (*emptypb.Empty, error)
	mustEmbedUnimplementedTranslatorServiceServer()
}

// UnimplementedTranslatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTranslatorServiceServer struct{}

func (UnimplementedTranslatorServiceServer) Translate(context.Context, *TranslateRequest) (*TranslateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedTranslatorServiceServer) GetConfig(context.Context, *emptypb.Empty) (*configpb.SubtitleManagerConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedTranslatorServiceServer) SetConfig(context.Context, *configpb.SubtitleManagerConfig) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedTranslatorServiceServer) mustEmbedUnimplementedTranslatorServiceServer() {}
func (UnimplementedTranslatorServiceServer) testEmbeddedByValue()                           {}

// UnsafeTranslatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslatorServiceServer will
// result in compilation errors.
type UnsafeTranslatorServiceServer interface {
	mustEmbedUnimplementedTranslatorServiceServer()
}

func RegisterTranslatorServiceServer(s grpc.ServiceRegistrar, srv TranslatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedTranslatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TranslatorService_ServiceDesc, srv)
}

func _TranslatorService_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServiceServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslatorService_Translate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServiceServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslatorService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslatorService_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServiceServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslatorService_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(configpb.SubtitleManagerConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServiceServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslatorService_SetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServiceServer).SetConfig(ctx, req.(*configpb.SubtitleManagerConfig))
	}
	return interceptor(ctx, in, info, handler)
}

// TranslatorService_ServiceDesc is the grpc.ServiceDesc for TranslatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranslatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gcommon.v1.translator.TranslatorService",
	HandlerType: (*TranslatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _TranslatorService_Translate_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _TranslatorService_GetConfig_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _TranslatorService_SetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/translator.proto",
}
