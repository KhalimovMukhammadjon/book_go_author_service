// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: author_service.proto

package author_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthorService_Create_FullMethodName      = "/author_service.AuthorService/Create"
	AuthorService_GetbyId_FullMethodName     = "/author_service.AuthorService/GetbyId"
	AuthorService_GetList_FullMethodName     = "/author_service.AuthorService/GetList"
	AuthorService_Update_FullMethodName      = "/author_service.AuthorService/Update"
	AuthorService_UpdatePatch_FullMethodName = "/author_service.AuthorService/UpdatePatch"
	AuthorService_Delete_FullMethodName      = "/author_service.AuthorService/Delete"
)

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	Create(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	GetbyId(ctx context.Context, in *AuthorPrimaryKey, opts ...grpc.CallOption) (*Author, error)
	GetList(ctx context.Context, in *GetAuthorListRequest, opts ...grpc.CallOption) (*GetAuthorListResponse, error)
	Update(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	UpdatePatch(ctx context.Context, in *PatchUpdateRequest, opts ...grpc.CallOption) (*Author, error)
	Delete(ctx context.Context, in *AuthorPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) Create(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, AuthorService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetbyId(ctx context.Context, in *AuthorPrimaryKey, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, AuthorService_GetbyId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetList(ctx context.Context, in *GetAuthorListRequest, opts ...grpc.CallOption) (*GetAuthorListResponse, error) {
	out := new(GetAuthorListResponse)
	err := c.cc.Invoke(ctx, AuthorService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) Update(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, AuthorService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) UpdatePatch(ctx context.Context, in *PatchUpdateRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, AuthorService_UpdatePatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) Delete(ctx context.Context, in *AuthorPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthorService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility
type AuthorServiceServer interface {
	Create(context.Context, *CreateAuthorRequest) (*Author, error)
	GetbyId(context.Context, *AuthorPrimaryKey) (*Author, error)
	GetList(context.Context, *GetAuthorListRequest) (*GetAuthorListResponse, error)
	Update(context.Context, *UpdateAuthorRequest) (*Author, error)
	UpdatePatch(context.Context, *PatchUpdateRequest) (*Author, error)
	Delete(context.Context, *AuthorPrimaryKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (UnimplementedAuthorServiceServer) Create(context.Context, *CreateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAuthorServiceServer) GetbyId(context.Context, *AuthorPrimaryKey) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetbyId not implemented")
}
func (UnimplementedAuthorServiceServer) GetList(context.Context, *GetAuthorListRequest) (*GetAuthorListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAuthorServiceServer) Update(context.Context, *UpdateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAuthorServiceServer) UpdatePatch(context.Context, *PatchUpdateRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatch not implemented")
}
func (UnimplementedAuthorServiceServer) Delete(context.Context, *AuthorPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAuthorServiceServer) mustEmbedUnimplementedAuthorServiceServer() {}

// UnsafeAuthorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServiceServer will
// result in compilation errors.
type UnsafeAuthorServiceServer interface {
	mustEmbedUnimplementedAuthorServiceServer()
}

func RegisterAuthorServiceServer(s grpc.ServiceRegistrar, srv AuthorServiceServer) {
	s.RegisterService(&AuthorService_ServiceDesc, srv)
}

func _AuthorService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Create(ctx, req.(*CreateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetbyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetbyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_GetbyId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetbyId(ctx, req.(*AuthorPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetList(ctx, req.(*GetAuthorListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Update(ctx, req.(*UpdateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_UpdatePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).UpdatePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_UpdatePatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).UpdatePatch(ctx, req.(*PatchUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Delete(ctx, req.(*AuthorPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorService_ServiceDesc is the grpc.ServiceDesc for AuthorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "author_service.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthorService_Create_Handler,
		},
		{
			MethodName: "GetbyId",
			Handler:    _AuthorService_GetbyId_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _AuthorService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AuthorService_Update_Handler,
		},
		{
			MethodName: "UpdatePatch",
			Handler:    _AuthorService_UpdatePatch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthorService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author_service.proto",
}
