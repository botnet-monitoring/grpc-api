// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: storage.proto

package gen

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
	BMSStorageService_RegisterSession_FullMethodName          = "/grpc_api.BMSStorageService/RegisterSession"
	BMSStorageService_StoreDatedBotReplyBatch_FullMethodName  = "/grpc_api.BMSStorageService/StoreDatedBotReplyBatch"
	BMSStorageService_StoreDatedFailedTryBatch_FullMethodName = "/grpc_api.BMSStorageService/StoreDatedFailedTryBatch"
	BMSStorageService_StoreDatedEdgeBatch_FullMethodName      = "/grpc_api.BMSStorageService/StoreDatedEdgeBatch"
	BMSStorageService_Disconnect_FullMethodName               = "/grpc_api.BMSStorageService/Disconnect"
)

// BMSStorageServiceClient is the client API for BMSStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BMSStorageServiceClient interface {
	// The method to register a new monitoring session.
	//
	// This method gives you back a session token which is needed for all further requests corresponding to the session. Furthermore it (optionally) can tell the client the time of the last bot reply batch, failed try batch and edge batch for the last session of the same monitor (and thereby giving an indicator for the last successful insert).
	//
	// Note: This method needs authentication via providing the 256bit auth token as `auth-token-bin` and the monitor ID as `auth-monitor-id` in the metadata (last given value wins).
	RegisterSession(ctx context.Context, in *RegisterSessionRequest, opts ...grpc.CallOption) (*RegisterSessionResponse, error)
	// The method to send a stream of dated bot replies.
	//
	// The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`.
	// The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	//
	// @todo: Document errors which can occur
	StoreDatedBotReplyBatch(ctx context.Context, opts ...grpc.CallOption) (BMSStorageService_StoreDatedBotReplyBatchClient, error)
	// The method to send a stream of dated (failed) bot contact tries.
	//
	// The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`.
	// The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	//
	// @todo: Document errors which can occur
	StoreDatedFailedTryBatch(ctx context.Context, opts ...grpc.CallOption) (BMSStorageService_StoreDatedFailedTryBatchClient, error)
	// The method to send a stream of dated bot edges.
	//
	// The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`.
	// The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	// Note: The server will not accept batches with more than 10000 bot replies.
	//
	// @todo: Document errors which can occur
	StoreDatedEdgeBatch(ctx context.Context, opts ...grpc.CallOption) (BMSStorageService_StoreDatedEdgeBatchClient, error)
	// The method to disconnect the current monitoring session.
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	// Note: The server will not accept batches with more than 10000 edges.
	//
	// @todo: Document errors which can occur
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
}

type bMSStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBMSStorageServiceClient(cc grpc.ClientConnInterface) BMSStorageServiceClient {
	return &bMSStorageServiceClient{cc}
}

func (c *bMSStorageServiceClient) RegisterSession(ctx context.Context, in *RegisterSessionRequest, opts ...grpc.CallOption) (*RegisterSessionResponse, error) {
	out := new(RegisterSessionResponse)
	err := c.cc.Invoke(ctx, BMSStorageService_RegisterSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bMSStorageServiceClient) StoreDatedBotReplyBatch(ctx context.Context, opts ...grpc.CallOption) (BMSStorageService_StoreDatedBotReplyBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &BMSStorageService_ServiceDesc.Streams[0], BMSStorageService_StoreDatedBotReplyBatch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bMSStorageServiceStoreDatedBotReplyBatchClient{stream}
	return x, nil
}

type BMSStorageService_StoreDatedBotReplyBatchClient interface {
	Send(*StoreDatedBotReplyBatchRequest) error
	Recv() (*StoreDatedBotReplyBatchResponse, error)
	grpc.ClientStream
}

type bMSStorageServiceStoreDatedBotReplyBatchClient struct {
	grpc.ClientStream
}

func (x *bMSStorageServiceStoreDatedBotReplyBatchClient) Send(m *StoreDatedBotReplyBatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bMSStorageServiceStoreDatedBotReplyBatchClient) Recv() (*StoreDatedBotReplyBatchResponse, error) {
	m := new(StoreDatedBotReplyBatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bMSStorageServiceClient) StoreDatedFailedTryBatch(ctx context.Context, opts ...grpc.CallOption) (BMSStorageService_StoreDatedFailedTryBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &BMSStorageService_ServiceDesc.Streams[1], BMSStorageService_StoreDatedFailedTryBatch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bMSStorageServiceStoreDatedFailedTryBatchClient{stream}
	return x, nil
}

type BMSStorageService_StoreDatedFailedTryBatchClient interface {
	Send(*StoreDatedFailedTryBatchRequest) error
	Recv() (*StoreDatedFailedTryBatchResponse, error)
	grpc.ClientStream
}

type bMSStorageServiceStoreDatedFailedTryBatchClient struct {
	grpc.ClientStream
}

func (x *bMSStorageServiceStoreDatedFailedTryBatchClient) Send(m *StoreDatedFailedTryBatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bMSStorageServiceStoreDatedFailedTryBatchClient) Recv() (*StoreDatedFailedTryBatchResponse, error) {
	m := new(StoreDatedFailedTryBatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bMSStorageServiceClient) StoreDatedEdgeBatch(ctx context.Context, opts ...grpc.CallOption) (BMSStorageService_StoreDatedEdgeBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &BMSStorageService_ServiceDesc.Streams[2], BMSStorageService_StoreDatedEdgeBatch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bMSStorageServiceStoreDatedEdgeBatchClient{stream}
	return x, nil
}

type BMSStorageService_StoreDatedEdgeBatchClient interface {
	Send(*StoreDatedEdgeBatchRequest) error
	Recv() (*StoreDatedEdgeBatchResponse, error)
	grpc.ClientStream
}

type bMSStorageServiceStoreDatedEdgeBatchClient struct {
	grpc.ClientStream
}

func (x *bMSStorageServiceStoreDatedEdgeBatchClient) Send(m *StoreDatedEdgeBatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bMSStorageServiceStoreDatedEdgeBatchClient) Recv() (*StoreDatedEdgeBatchResponse, error) {
	m := new(StoreDatedEdgeBatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bMSStorageServiceClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, BMSStorageService_Disconnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BMSStorageServiceServer is the server API for BMSStorageService service.
// All implementations must embed UnimplementedBMSStorageServiceServer
// for forward compatibility
type BMSStorageServiceServer interface {
	// The method to register a new monitoring session.
	//
	// This method gives you back a session token which is needed for all further requests corresponding to the session. Furthermore it (optionally) can tell the client the time of the last bot reply batch, failed try batch and edge batch for the last session of the same monitor (and thereby giving an indicator for the last successful insert).
	//
	// Note: This method needs authentication via providing the 256bit auth token as `auth-token-bin` and the monitor ID as `auth-monitor-id` in the metadata (last given value wins).
	RegisterSession(context.Context, *RegisterSessionRequest) (*RegisterSessionResponse, error)
	// The method to send a stream of dated bot replies.
	//
	// The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`.
	// The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	//
	// @todo: Document errors which can occur
	StoreDatedBotReplyBatch(BMSStorageService_StoreDatedBotReplyBatchServer) error
	// The method to send a stream of dated (failed) bot contact tries.
	//
	// The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`.
	// The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	//
	// @todo: Document errors which can occur
	StoreDatedFailedTryBatch(BMSStorageService_StoreDatedFailedTryBatchServer) error
	// The method to send a stream of dated bot edges.
	//
	// The server acknowledges the successful handling (receiving, processing, storing) of the sent batch by sending a response with the respective `batchID`.
	// The client should cache the sent batch until the server acknowledges it explicitly (otherwise data might be lost).
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	// Note: The server will not accept batches with more than 10000 bot replies.
	//
	// @todo: Document errors which can occur
	StoreDatedEdgeBatch(BMSStorageService_StoreDatedEdgeBatchServer) error
	// The method to disconnect the current monitoring session.
	//
	// Note: This method needs (session) authentication via providing the 32bit session token as `session-token-bin` in the metadata (last given value wins).
	// Note: The server will not accept batches with more than 10000 edges.
	//
	// @todo: Document errors which can occur
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	mustEmbedUnimplementedBMSStorageServiceServer()
}

// UnimplementedBMSStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBMSStorageServiceServer struct {
}

func (UnimplementedBMSStorageServiceServer) RegisterSession(context.Context, *RegisterSessionRequest) (*RegisterSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSession not implemented")
}
func (UnimplementedBMSStorageServiceServer) StoreDatedBotReplyBatch(BMSStorageService_StoreDatedBotReplyBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreDatedBotReplyBatch not implemented")
}
func (UnimplementedBMSStorageServiceServer) StoreDatedFailedTryBatch(BMSStorageService_StoreDatedFailedTryBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreDatedFailedTryBatch not implemented")
}
func (UnimplementedBMSStorageServiceServer) StoreDatedEdgeBatch(BMSStorageService_StoreDatedEdgeBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreDatedEdgeBatch not implemented")
}
func (UnimplementedBMSStorageServiceServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedBMSStorageServiceServer) mustEmbedUnimplementedBMSStorageServiceServer() {}

// UnsafeBMSStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BMSStorageServiceServer will
// result in compilation errors.
type UnsafeBMSStorageServiceServer interface {
	mustEmbedUnimplementedBMSStorageServiceServer()
}

func RegisterBMSStorageServiceServer(s grpc.ServiceRegistrar, srv BMSStorageServiceServer) {
	s.RegisterService(&BMSStorageService_ServiceDesc, srv)
}

func _BMSStorageService_RegisterSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BMSStorageServiceServer).RegisterSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BMSStorageService_RegisterSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BMSStorageServiceServer).RegisterSession(ctx, req.(*RegisterSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BMSStorageService_StoreDatedBotReplyBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BMSStorageServiceServer).StoreDatedBotReplyBatch(&bMSStorageServiceStoreDatedBotReplyBatchServer{stream})
}

type BMSStorageService_StoreDatedBotReplyBatchServer interface {
	Send(*StoreDatedBotReplyBatchResponse) error
	Recv() (*StoreDatedBotReplyBatchRequest, error)
	grpc.ServerStream
}

type bMSStorageServiceStoreDatedBotReplyBatchServer struct {
	grpc.ServerStream
}

func (x *bMSStorageServiceStoreDatedBotReplyBatchServer) Send(m *StoreDatedBotReplyBatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bMSStorageServiceStoreDatedBotReplyBatchServer) Recv() (*StoreDatedBotReplyBatchRequest, error) {
	m := new(StoreDatedBotReplyBatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BMSStorageService_StoreDatedFailedTryBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BMSStorageServiceServer).StoreDatedFailedTryBatch(&bMSStorageServiceStoreDatedFailedTryBatchServer{stream})
}

type BMSStorageService_StoreDatedFailedTryBatchServer interface {
	Send(*StoreDatedFailedTryBatchResponse) error
	Recv() (*StoreDatedFailedTryBatchRequest, error)
	grpc.ServerStream
}

type bMSStorageServiceStoreDatedFailedTryBatchServer struct {
	grpc.ServerStream
}

func (x *bMSStorageServiceStoreDatedFailedTryBatchServer) Send(m *StoreDatedFailedTryBatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bMSStorageServiceStoreDatedFailedTryBatchServer) Recv() (*StoreDatedFailedTryBatchRequest, error) {
	m := new(StoreDatedFailedTryBatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BMSStorageService_StoreDatedEdgeBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BMSStorageServiceServer).StoreDatedEdgeBatch(&bMSStorageServiceStoreDatedEdgeBatchServer{stream})
}

type BMSStorageService_StoreDatedEdgeBatchServer interface {
	Send(*StoreDatedEdgeBatchResponse) error
	Recv() (*StoreDatedEdgeBatchRequest, error)
	grpc.ServerStream
}

type bMSStorageServiceStoreDatedEdgeBatchServer struct {
	grpc.ServerStream
}

func (x *bMSStorageServiceStoreDatedEdgeBatchServer) Send(m *StoreDatedEdgeBatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bMSStorageServiceStoreDatedEdgeBatchServer) Recv() (*StoreDatedEdgeBatchRequest, error) {
	m := new(StoreDatedEdgeBatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BMSStorageService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BMSStorageServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BMSStorageService_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BMSStorageServiceServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BMSStorageService_ServiceDesc is the grpc.ServiceDesc for BMSStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BMSStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_api.BMSStorageService",
	HandlerType: (*BMSStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSession",
			Handler:    _BMSStorageService_RegisterSession_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _BMSStorageService_Disconnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StoreDatedBotReplyBatch",
			Handler:       _BMSStorageService_StoreDatedBotReplyBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StoreDatedFailedTryBatch",
			Handler:       _BMSStorageService_StoreDatedFailedTryBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StoreDatedEdgeBatch",
			Handler:       _BMSStorageService_StoreDatedEdgeBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}
