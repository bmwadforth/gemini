// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: protocol_buffers/gemini_service/gemini_service.proto

package gemini

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

// GeminiClient is the client API for Gemini service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeminiClient interface {
	QueryGemini(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Gemini_QueryGeminiClient, error)
}

type geminiClient struct {
	cc grpc.ClientConnInterface
}

func NewGeminiClient(cc grpc.ClientConnInterface) GeminiClient {
	return &geminiClient{cc}
}

func (c *geminiClient) QueryGemini(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Gemini_QueryGeminiClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gemini_ServiceDesc.Streams[0], "/gemini_service.Gemini/QueryGemini", opts...)
	if err != nil {
		return nil, err
	}
	x := &geminiQueryGeminiClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gemini_QueryGeminiClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type geminiQueryGeminiClient struct {
	grpc.ClientStream
}

func (x *geminiQueryGeminiClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GeminiServer is the server API for Gemini service.
// All implementations must embed UnimplementedGeminiServer
// for forward compatibility
type GeminiServer interface {
	QueryGemini(*QueryRequest, Gemini_QueryGeminiServer) error
	mustEmbedUnimplementedGeminiServer()
}

// UnimplementedGeminiServer must be embedded to have forward compatible implementations.
type UnimplementedGeminiServer struct {
}

func (UnimplementedGeminiServer) QueryGemini(*QueryRequest, Gemini_QueryGeminiServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryGemini not implemented")
}
func (UnimplementedGeminiServer) mustEmbedUnimplementedGeminiServer() {}

// UnsafeGeminiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeminiServer will
// result in compilation errors.
type UnsafeGeminiServer interface {
	mustEmbedUnimplementedGeminiServer()
}

func RegisterGeminiServer(s grpc.ServiceRegistrar, srv GeminiServer) {
	s.RegisterService(&Gemini_ServiceDesc, srv)
}

func _Gemini_QueryGemini_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeminiServer).QueryGemini(m, &geminiQueryGeminiServer{stream})
}

type Gemini_QueryGeminiServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type geminiQueryGeminiServer struct {
	grpc.ServerStream
}

func (x *geminiQueryGeminiServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Gemini_ServiceDesc is the grpc.ServiceDesc for Gemini service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gemini_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gemini_service.Gemini",
	HandlerType: (*GeminiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryGemini",
			Handler:       _Gemini_QueryGemini_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol_buffers/gemini_service/gemini_service.proto",
}