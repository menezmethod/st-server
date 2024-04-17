// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: pkg/pb/record.proto

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

const (
	RecordService_CreateRecord_FullMethodName = "/record.RecordService/CreateRecord"
	RecordService_UpdateRecord_FullMethodName = "/record.RecordService/UpdateRecord"
	RecordService_GetRecord_FullMethodName    = "/record.RecordService/GetRecord"
	RecordService_ListRecords_FullMethodName  = "/record.RecordService/ListRecords"
	RecordService_RemoveRecord_FullMethodName = "/record.RecordService/RemoveRecord"
)

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordServiceClient interface {
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error)
	GetRecord(ctx context.Context, in *FindOneRecordRequest, opts ...grpc.CallOption) (*FindOneRecordResponse, error)
	ListRecords(ctx context.Context, in *FindAllRecordsRequest, opts ...grpc.CallOption) (*FindAllRecordsResponse, error)
	RemoveRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error)
}

type recordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordServiceClient(cc grpc.ClientConnInterface) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, RecordService_CreateRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error) {
	out := new(UpdateRecordResponse)
	err := c.cc.Invoke(ctx, RecordService_UpdateRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) GetRecord(ctx context.Context, in *FindOneRecordRequest, opts ...grpc.CallOption) (*FindOneRecordResponse, error) {
	out := new(FindOneRecordResponse)
	err := c.cc.Invoke(ctx, RecordService_GetRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) ListRecords(ctx context.Context, in *FindAllRecordsRequest, opts ...grpc.CallOption) (*FindAllRecordsResponse, error) {
	out := new(FindAllRecordsResponse)
	err := c.cc.Invoke(ctx, RecordService_ListRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordServiceClient) RemoveRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error) {
	out := new(DeleteRecordResponse)
	err := c.cc.Invoke(ctx, RecordService_RemoveRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServiceServer is the server API for RecordService service.
// All implementations must embed UnimplementedRecordServiceServer
// for forward compatibility
type RecordServiceServer interface {
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	UpdateRecord(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error)
	GetRecord(context.Context, *FindOneRecordRequest) (*FindOneRecordResponse, error)
	ListRecords(context.Context, *FindAllRecordsRequest) (*FindAllRecordsResponse, error)
	RemoveRecord(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error)
	mustEmbedUnimplementedRecordServiceServer()
}

// UnimplementedRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServiceServer struct {
}

func (UnimplementedRecordServiceServer) CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedRecordServiceServer) UpdateRecord(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}
func (UnimplementedRecordServiceServer) GetRecord(context.Context, *FindOneRecordRequest) (*FindOneRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (UnimplementedRecordServiceServer) ListRecords(context.Context, *FindAllRecordsRequest) (*FindAllRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecords not implemented")
}
func (UnimplementedRecordServiceServer) RemoveRecord(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRecord not implemented")
}
func (UnimplementedRecordServiceServer) mustEmbedUnimplementedRecordServiceServer() {}

// UnsafeRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServiceServer will
// result in compilation errors.
type UnsafeRecordServiceServer interface {
	mustEmbedUnimplementedRecordServiceServer()
}

func RegisterRecordServiceServer(s grpc.ServiceRegistrar, srv RecordServiceServer) {
	s.RegisterService(&RecordService_ServiceDesc, srv)
}

func _RecordService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_CreateRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_UpdateRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).UpdateRecord(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_GetRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).GetRecord(ctx, req.(*FindOneRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_ListRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).ListRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_ListRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).ListRecords(ctx, req.(*FindAllRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordService_RemoveRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).RemoveRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecordService_RemoveRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).RemoveRecord(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordService_ServiceDesc is the grpc.ServiceDesc for RecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "record.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecord",
			Handler:    _RecordService_CreateRecord_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _RecordService_UpdateRecord_Handler,
		},
		{
			MethodName: "GetRecord",
			Handler:    _RecordService_GetRecord_Handler,
		},
		{
			MethodName: "ListRecords",
			Handler:    _RecordService_ListRecords_Handler,
		},
		{
			MethodName: "RemoveRecord",
			Handler:    _RecordService_RemoveRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/record.proto",
}
