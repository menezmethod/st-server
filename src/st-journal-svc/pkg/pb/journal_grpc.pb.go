// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: pkg/pb/journal.proto

package pb

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

// JournalServiceClient is the client API for JournalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JournalServiceClient interface {
	CreateJournal(ctx context.Context, in *CreateJournalRequest, opts ...grpc.CallOption) (*CreateJournalResponse, error)
	DeleteJournal(ctx context.Context, in *DeleteJournalRequest, opts ...grpc.CallOption) (*DeleteJournalResponse, error)
	EditJournal(ctx context.Context, in *EditJournalRequest, opts ...grpc.CallOption) (*EditJournalResponse, error)
	FindOneJournal(ctx context.Context, in *FindOneJournalRequest, opts ...grpc.CallOption) (*FindOneJournalResponse, error)
	CreateTrade(ctx context.Context, in *CreateTradeRequest, opts ...grpc.CallOption) (*CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *DeleteTradeRequest, opts ...grpc.CallOption) (*DeleteTradeResponse, error)
	EditTrade(ctx context.Context, in *EditTradeRequest, opts ...grpc.CallOption) (*EditTradeResponse, error)
	FindOneTrade(ctx context.Context, in *FindOneTradeRequest, opts ...grpc.CallOption) (*FindOneTradeResponse, error)
}

type journalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJournalServiceClient(cc grpc.ClientConnInterface) JournalServiceClient {
	return &journalServiceClient{cc}
}

func (c *journalServiceClient) CreateJournal(ctx context.Context, in *CreateJournalRequest, opts ...grpc.CallOption) (*CreateJournalResponse, error) {
	out := new(CreateJournalResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/CreateJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) DeleteJournal(ctx context.Context, in *DeleteJournalRequest, opts ...grpc.CallOption) (*DeleteJournalResponse, error) {
	out := new(DeleteJournalResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/DeleteJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) EditJournal(ctx context.Context, in *EditJournalRequest, opts ...grpc.CallOption) (*EditJournalResponse, error) {
	out := new(EditJournalResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/EditJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) FindOneJournal(ctx context.Context, in *FindOneJournalRequest, opts ...grpc.CallOption) (*FindOneJournalResponse, error) {
	out := new(FindOneJournalResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/FindOneJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) CreateTrade(ctx context.Context, in *CreateTradeRequest, opts ...grpc.CallOption) (*CreateTradeResponse, error) {
	out := new(CreateTradeResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/CreateTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) DeleteTrade(ctx context.Context, in *DeleteTradeRequest, opts ...grpc.CallOption) (*DeleteTradeResponse, error) {
	out := new(DeleteTradeResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/DeleteTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) EditTrade(ctx context.Context, in *EditTradeRequest, opts ...grpc.CallOption) (*EditTradeResponse, error) {
	out := new(EditTradeResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/EditTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) FindOneTrade(ctx context.Context, in *FindOneTradeRequest, opts ...grpc.CallOption) (*FindOneTradeResponse, error) {
	out := new(FindOneTradeResponse)
	err := c.cc.Invoke(ctx, "/journal.JournalService/FindOneTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JournalServiceServer is the server API for JournalService service.
// All implementations must embed UnimplementedJournalServiceServer
// for forward compatibility
type JournalServiceServer interface {
	CreateJournal(context.Context, *CreateJournalRequest) (*CreateJournalResponse, error)
	DeleteJournal(context.Context, *DeleteJournalRequest) (*DeleteJournalResponse, error)
	EditJournal(context.Context, *EditJournalRequest) (*EditJournalResponse, error)
	FindOneJournal(context.Context, *FindOneJournalRequest) (*FindOneJournalResponse, error)
	CreateTrade(context.Context, *CreateTradeRequest) (*CreateTradeResponse, error)
	DeleteTrade(context.Context, *DeleteTradeRequest) (*DeleteTradeResponse, error)
	EditTrade(context.Context, *EditTradeRequest) (*EditTradeResponse, error)
	FindOneTrade(context.Context, *FindOneTradeRequest) (*FindOneTradeResponse, error)
	mustEmbedUnimplementedJournalServiceServer()
}

// UnimplementedJournalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJournalServiceServer struct {
}

func (UnimplementedJournalServiceServer) CreateJournal(context.Context, *CreateJournalRequest) (*CreateJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJournal not implemented")
}
func (UnimplementedJournalServiceServer) DeleteJournal(context.Context, *DeleteJournalRequest) (*DeleteJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJournal not implemented")
}
func (UnimplementedJournalServiceServer) EditJournal(context.Context, *EditJournalRequest) (*EditJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditJournal not implemented")
}
func (UnimplementedJournalServiceServer) FindOneJournal(context.Context, *FindOneJournalRequest) (*FindOneJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneJournal not implemented")
}
func (UnimplementedJournalServiceServer) CreateTrade(context.Context, *CreateTradeRequest) (*CreateTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrade not implemented")
}
func (UnimplementedJournalServiceServer) DeleteTrade(context.Context, *DeleteTradeRequest) (*DeleteTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrade not implemented")
}
func (UnimplementedJournalServiceServer) EditTrade(context.Context, *EditTradeRequest) (*EditTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTrade not implemented")
}
func (UnimplementedJournalServiceServer) FindOneTrade(context.Context, *FindOneTradeRequest) (*FindOneTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneTrade not implemented")
}
func (UnimplementedJournalServiceServer) mustEmbedUnimplementedJournalServiceServer() {}

// UnsafeJournalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JournalServiceServer will
// result in compilation errors.
type UnsafeJournalServiceServer interface {
	mustEmbedUnimplementedJournalServiceServer()
}

func RegisterJournalServiceServer(s grpc.ServiceRegistrar, srv JournalServiceServer) {
	s.RegisterService(&JournalService_ServiceDesc, srv)
}

func _JournalService_CreateJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).CreateJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/CreateJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).CreateJournal(ctx, req.(*CreateJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_DeleteJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).DeleteJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/DeleteJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).DeleteJournal(ctx, req.(*DeleteJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_EditJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).EditJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/EditJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).EditJournal(ctx, req.(*EditJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_FindOneJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).FindOneJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/FindOneJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).FindOneJournal(ctx, req.(*FindOneJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_CreateTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).CreateTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/CreateTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).CreateTrade(ctx, req.(*CreateTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_DeleteTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).DeleteTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/DeleteTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).DeleteTrade(ctx, req.(*DeleteTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_EditTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).EditTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/EditTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).EditTrade(ctx, req.(*EditTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_FindOneTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).FindOneTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/journal.JournalService/FindOneTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).FindOneTrade(ctx, req.(*FindOneTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JournalService_ServiceDesc is the grpc.ServiceDesc for JournalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JournalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "journal.JournalService",
	HandlerType: (*JournalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJournal",
			Handler:    _JournalService_CreateJournal_Handler,
		},
		{
			MethodName: "DeleteJournal",
			Handler:    _JournalService_DeleteJournal_Handler,
		},
		{
			MethodName: "EditJournal",
			Handler:    _JournalService_EditJournal_Handler,
		},
		{
			MethodName: "FindOneJournal",
			Handler:    _JournalService_FindOneJournal_Handler,
		},
		{
			MethodName: "CreateTrade",
			Handler:    _JournalService_CreateTrade_Handler,
		},
		{
			MethodName: "DeleteTrade",
			Handler:    _JournalService_DeleteTrade_Handler,
		},
		{
			MethodName: "EditTrade",
			Handler:    _JournalService_EditTrade_Handler,
		},
		{
			MethodName: "FindOneTrade",
			Handler:    _JournalService_FindOneTrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/journal.proto",
}
