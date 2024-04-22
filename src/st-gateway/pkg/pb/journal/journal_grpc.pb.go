// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: journal/journal.proto

package journal

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
	JournalService_CreateJournal_FullMethodName = "/journal.JournalService/CreateJournal"
	JournalService_RemoveJournal_FullMethodName = "/journal.JournalService/RemoveJournal"
	JournalService_UpdateJournal_FullMethodName = "/journal.JournalService/UpdateJournal"
	JournalService_ListJournals_FullMethodName  = "/journal.JournalService/ListJournals"
	JournalService_GetJournal_FullMethodName    = "/journal.JournalService/GetJournal"
)

// JournalServiceClient is the client API for JournalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JournalServiceClient interface {
	CreateJournal(ctx context.Context, in *CreateJournalRequest, opts ...grpc.CallOption) (*CreateJournalResponse, error)
	RemoveJournal(ctx context.Context, in *DeleteJournalRequest, opts ...grpc.CallOption) (*DeleteJournalResponse, error)
	UpdateJournal(ctx context.Context, in *UpdateJournalRequest, opts ...grpc.CallOption) (*UpdateJournalResponse, error)
	ListJournals(ctx context.Context, in *FindAllJournalsRequest, opts ...grpc.CallOption) (*FindAllJournalsResponse, error)
	GetJournal(ctx context.Context, in *FindOneJournalRequest, opts ...grpc.CallOption) (*FindOneJournalResponse, error)
}

type journalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJournalServiceClient(cc grpc.ClientConnInterface) JournalServiceClient {
	return &journalServiceClient{cc}
}

func (c *journalServiceClient) CreateJournal(ctx context.Context, in *CreateJournalRequest, opts ...grpc.CallOption) (*CreateJournalResponse, error) {
	out := new(CreateJournalResponse)
	err := c.cc.Invoke(ctx, JournalService_CreateJournal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) RemoveJournal(ctx context.Context, in *DeleteJournalRequest, opts ...grpc.CallOption) (*DeleteJournalResponse, error) {
	out := new(DeleteJournalResponse)
	err := c.cc.Invoke(ctx, JournalService_RemoveJournal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) UpdateJournal(ctx context.Context, in *UpdateJournalRequest, opts ...grpc.CallOption) (*UpdateJournalResponse, error) {
	out := new(UpdateJournalResponse)
	err := c.cc.Invoke(ctx, JournalService_UpdateJournal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) ListJournals(ctx context.Context, in *FindAllJournalsRequest, opts ...grpc.CallOption) (*FindAllJournalsResponse, error) {
	out := new(FindAllJournalsResponse)
	err := c.cc.Invoke(ctx, JournalService_ListJournals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *journalServiceClient) GetJournal(ctx context.Context, in *FindOneJournalRequest, opts ...grpc.CallOption) (*FindOneJournalResponse, error) {
	out := new(FindOneJournalResponse)
	err := c.cc.Invoke(ctx, JournalService_GetJournal_FullMethodName, in, out, opts...)
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
	RemoveJournal(context.Context, *DeleteJournalRequest) (*DeleteJournalResponse, error)
	UpdateJournal(context.Context, *UpdateJournalRequest) (*UpdateJournalResponse, error)
	ListJournals(context.Context, *FindAllJournalsRequest) (*FindAllJournalsResponse, error)
	GetJournal(context.Context, *FindOneJournalRequest) (*FindOneJournalResponse, error)
	mustEmbedUnimplementedJournalServiceServer()
}

// UnimplementedJournalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJournalServiceServer struct {
}

func (UnimplementedJournalServiceServer) CreateJournal(context.Context, *CreateJournalRequest) (*CreateJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJournal not implemented")
}
func (UnimplementedJournalServiceServer) RemoveJournal(context.Context, *DeleteJournalRequest) (*DeleteJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveJournal not implemented")
}
func (UnimplementedJournalServiceServer) UpdateJournal(context.Context, *UpdateJournalRequest) (*UpdateJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJournal not implemented")
}
func (UnimplementedJournalServiceServer) ListJournals(context.Context, *FindAllJournalsRequest) (*FindAllJournalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJournals not implemented")
}
func (UnimplementedJournalServiceServer) GetJournal(context.Context, *FindOneJournalRequest) (*FindOneJournalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJournal not implemented")
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
		FullMethod: JournalService_CreateJournal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).CreateJournal(ctx, req.(*CreateJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_RemoveJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).RemoveJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JournalService_RemoveJournal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).RemoveJournal(ctx, req.(*DeleteJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_UpdateJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).UpdateJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JournalService_UpdateJournal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).UpdateJournal(ctx, req.(*UpdateJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_ListJournals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllJournalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).ListJournals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JournalService_ListJournals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).ListJournals(ctx, req.(*FindAllJournalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JournalService_GetJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JournalServiceServer).GetJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JournalService_GetJournal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JournalServiceServer).GetJournal(ctx, req.(*FindOneJournalRequest))
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
			MethodName: "RemoveJournal",
			Handler:    _JournalService_RemoveJournal_Handler,
		},
		{
			MethodName: "UpdateJournal",
			Handler:    _JournalService_UpdateJournal_Handler,
		},
		{
			MethodName: "ListJournals",
			Handler:    _JournalService_ListJournals_Handler,
		},
		{
			MethodName: "GetJournal",
			Handler:    _JournalService_GetJournal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "journal/journal.proto",
}