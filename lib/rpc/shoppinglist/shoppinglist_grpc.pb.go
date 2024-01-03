// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: shoppinglist/shoppinglist.proto

package shoppinglist

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
	UserShoppingListService_CreateUserShoppingList_FullMethodName  = "/shoppinglist.UserShoppingListService/CreateUserShoppingList"
	UserShoppingListService_GetUserShoppingList_FullMethodName     = "/shoppinglist.UserShoppingListService/GetUserShoppingList"
	UserShoppingListService_GetAllUserShoppingLists_FullMethodName = "/shoppinglist.UserShoppingListService/GetAllUserShoppingLists"
	UserShoppingListService_UpdateUserShoppingList_FullMethodName  = "/shoppinglist.UserShoppingListService/UpdateUserShoppingList"
	UserShoppingListService_DeleteUserShoppingList_FullMethodName  = "/shoppinglist.UserShoppingListService/DeleteUserShoppingList"
)

// UserShoppingListServiceClient is the client API for UserShoppingListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserShoppingListServiceClient interface {
	// create method for adding a new shopping list
	CreateUserShoppingList(ctx context.Context, in *CreateUserShoppingListRequest, opts ...grpc.CallOption) (*CreateUserShoppingListResponse, error)
	// read method for getting a specific shopping list
	GetUserShoppingList(ctx context.Context, in *GetUserShoppingListRequest, opts ...grpc.CallOption) (*GetUserShoppingListResponse, error)
	// read method for getting all shopping lists for an user
	GetAllUserShoppingLists(ctx context.Context, in *GetAllUserShoppingListsRequest, opts ...grpc.CallOption) (*GetAllUserShoppingListsResponse, error)
	// update method for updating a shopping list
	UpdateUserShoppingList(ctx context.Context, in *UpdateUserShoppingListRequest, opts ...grpc.CallOption) (*UpdateUserShoppingListResponse, error)
	// delete method to get rid of a shopping list
	DeleteUserShoppingList(ctx context.Context, in *DeleteUserShoppingListRequest, opts ...grpc.CallOption) (*DeleteUserShoppingListResponse, error)
}

type userShoppingListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserShoppingListServiceClient(cc grpc.ClientConnInterface) UserShoppingListServiceClient {
	return &userShoppingListServiceClient{cc}
}

func (c *userShoppingListServiceClient) CreateUserShoppingList(ctx context.Context, in *CreateUserShoppingListRequest, opts ...grpc.CallOption) (*CreateUserShoppingListResponse, error) {
	out := new(CreateUserShoppingListResponse)
	err := c.cc.Invoke(ctx, UserShoppingListService_CreateUserShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListServiceClient) GetUserShoppingList(ctx context.Context, in *GetUserShoppingListRequest, opts ...grpc.CallOption) (*GetUserShoppingListResponse, error) {
	out := new(GetUserShoppingListResponse)
	err := c.cc.Invoke(ctx, UserShoppingListService_GetUserShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListServiceClient) GetAllUserShoppingLists(ctx context.Context, in *GetAllUserShoppingListsRequest, opts ...grpc.CallOption) (*GetAllUserShoppingListsResponse, error) {
	out := new(GetAllUserShoppingListsResponse)
	err := c.cc.Invoke(ctx, UserShoppingListService_GetAllUserShoppingLists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListServiceClient) UpdateUserShoppingList(ctx context.Context, in *UpdateUserShoppingListRequest, opts ...grpc.CallOption) (*UpdateUserShoppingListResponse, error) {
	out := new(UpdateUserShoppingListResponse)
	err := c.cc.Invoke(ctx, UserShoppingListService_UpdateUserShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListServiceClient) DeleteUserShoppingList(ctx context.Context, in *DeleteUserShoppingListRequest, opts ...grpc.CallOption) (*DeleteUserShoppingListResponse, error) {
	out := new(DeleteUserShoppingListResponse)
	err := c.cc.Invoke(ctx, UserShoppingListService_DeleteUserShoppingList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserShoppingListServiceServer is the server API for UserShoppingListService service.
// All implementations must embed UnimplementedUserShoppingListServiceServer
// for forward compatibility
type UserShoppingListServiceServer interface {
	// create method for adding a new shopping list
	CreateUserShoppingList(context.Context, *CreateUserShoppingListRequest) (*CreateUserShoppingListResponse, error)
	// read method for getting a specific shopping list
	GetUserShoppingList(context.Context, *GetUserShoppingListRequest) (*GetUserShoppingListResponse, error)
	// read method for getting all shopping lists for an user
	GetAllUserShoppingLists(context.Context, *GetAllUserShoppingListsRequest) (*GetAllUserShoppingListsResponse, error)
	// update method for updating a shopping list
	UpdateUserShoppingList(context.Context, *UpdateUserShoppingListRequest) (*UpdateUserShoppingListResponse, error)
	// delete method to get rid of a shopping list
	DeleteUserShoppingList(context.Context, *DeleteUserShoppingListRequest) (*DeleteUserShoppingListResponse, error)
	mustEmbedUnimplementedUserShoppingListServiceServer()
}

// UnimplementedUserShoppingListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserShoppingListServiceServer struct {
}

func (UnimplementedUserShoppingListServiceServer) CreateUserShoppingList(context.Context, *CreateUserShoppingListRequest) (*CreateUserShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserShoppingList not implemented")
}
func (UnimplementedUserShoppingListServiceServer) GetUserShoppingList(context.Context, *GetUserShoppingListRequest) (*GetUserShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserShoppingList not implemented")
}
func (UnimplementedUserShoppingListServiceServer) GetAllUserShoppingLists(context.Context, *GetAllUserShoppingListsRequest) (*GetAllUserShoppingListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserShoppingLists not implemented")
}
func (UnimplementedUserShoppingListServiceServer) UpdateUserShoppingList(context.Context, *UpdateUserShoppingListRequest) (*UpdateUserShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserShoppingList not implemented")
}
func (UnimplementedUserShoppingListServiceServer) DeleteUserShoppingList(context.Context, *DeleteUserShoppingListRequest) (*DeleteUserShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserShoppingList not implemented")
}
func (UnimplementedUserShoppingListServiceServer) mustEmbedUnimplementedUserShoppingListServiceServer() {
}

// UnsafeUserShoppingListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserShoppingListServiceServer will
// result in compilation errors.
type UnsafeUserShoppingListServiceServer interface {
	mustEmbedUnimplementedUserShoppingListServiceServer()
}

func RegisterUserShoppingListServiceServer(s grpc.ServiceRegistrar, srv UserShoppingListServiceServer) {
	s.RegisterService(&UserShoppingListService_ServiceDesc, srv)
}

func _UserShoppingListService_CreateUserShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListServiceServer).CreateUserShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListService_CreateUserShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListServiceServer).CreateUserShoppingList(ctx, req.(*CreateUserShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListService_GetUserShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListServiceServer).GetUserShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListService_GetUserShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListServiceServer).GetUserShoppingList(ctx, req.(*GetUserShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListService_GetAllUserShoppingLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserShoppingListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListServiceServer).GetAllUserShoppingLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListService_GetAllUserShoppingLists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListServiceServer).GetAllUserShoppingLists(ctx, req.(*GetAllUserShoppingListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListService_UpdateUserShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListServiceServer).UpdateUserShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListService_UpdateUserShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListServiceServer).UpdateUserShoppingList(ctx, req.(*UpdateUserShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListService_DeleteUserShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListServiceServer).DeleteUserShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListService_DeleteUserShoppingList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListServiceServer).DeleteUserShoppingList(ctx, req.(*DeleteUserShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserShoppingListService_ServiceDesc is the grpc.ServiceDesc for UserShoppingListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserShoppingListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shoppinglist.UserShoppingListService",
	HandlerType: (*UserShoppingListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserShoppingList",
			Handler:    _UserShoppingListService_CreateUserShoppingList_Handler,
		},
		{
			MethodName: "GetUserShoppingList",
			Handler:    _UserShoppingListService_GetUserShoppingList_Handler,
		},
		{
			MethodName: "GetAllUserShoppingLists",
			Handler:    _UserShoppingListService_GetAllUserShoppingLists_Handler,
		},
		{
			MethodName: "UpdateUserShoppingList",
			Handler:    _UserShoppingListService_UpdateUserShoppingList_Handler,
		},
		{
			MethodName: "DeleteUserShoppingList",
			Handler:    _UserShoppingListService_DeleteUserShoppingList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppinglist/shoppinglist.proto",
}

const (
	UserShoppingListEntryService_CreateUserShoppingListEntry_FullMethodName   = "/shoppinglist.UserShoppingListEntryService/CreateUserShoppingListEntry"
	UserShoppingListEntryService_GetUserShoppingListEntry_FullMethodName      = "/shoppinglist.UserShoppingListEntryService/GetUserShoppingListEntry"
	UserShoppingListEntryService_GetAllUserShoppingListEntries_FullMethodName = "/shoppinglist.UserShoppingListEntryService/GetAllUserShoppingListEntries"
	UserShoppingListEntryService_UpdateUserShoppingListEntry_FullMethodName   = "/shoppinglist.UserShoppingListEntryService/UpdateUserShoppingListEntry"
	UserShoppingListEntryService_DeleteUserShoppingListEntry_FullMethodName   = "/shoppinglist.UserShoppingListEntryService/DeleteUserShoppingListEntry"
)

// UserShoppingListEntryServiceClient is the client API for UserShoppingListEntryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserShoppingListEntryServiceClient interface {
	// create method for adding a new shopping list entry
	CreateUserShoppingListEntry(ctx context.Context, in *CreateUserShoppingListEntryRequest, opts ...grpc.CallOption) (*CreateUserShoppingListEntryResponse, error)
	// read method for getting a specific shopping list entry
	GetUserShoppingListEntry(ctx context.Context, in *GetUserShoppingListEntryRequest, opts ...grpc.CallOption) (*GetUserShoppingListEntryResponse, error)
	// read method for getting all shopping list entry for a specific shopping list
	GetAllUserShoppingListEntries(ctx context.Context, in *GetAllUserShoppingListEntriesRequest, opts ...grpc.CallOption) (*GetAllUserShoppingListEntriesResponse, error)
	// update method for updating a shopping list
	UpdateUserShoppingListEntry(ctx context.Context, in *UpdateUserShoppingListEntryRequest, opts ...grpc.CallOption) (*UpdateUserShoppingListEntryResponse, error)
	// delete method to get rid of a shopping list
	DeleteUserShoppingListEntry(ctx context.Context, in *DeleteUserShoppingListEntryRequest, opts ...grpc.CallOption) (*DeleteUserShoppingListEntryResponse, error)
}

type userShoppingListEntryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserShoppingListEntryServiceClient(cc grpc.ClientConnInterface) UserShoppingListEntryServiceClient {
	return &userShoppingListEntryServiceClient{cc}
}

func (c *userShoppingListEntryServiceClient) CreateUserShoppingListEntry(ctx context.Context, in *CreateUserShoppingListEntryRequest, opts ...grpc.CallOption) (*CreateUserShoppingListEntryResponse, error) {
	out := new(CreateUserShoppingListEntryResponse)
	err := c.cc.Invoke(ctx, UserShoppingListEntryService_CreateUserShoppingListEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListEntryServiceClient) GetUserShoppingListEntry(ctx context.Context, in *GetUserShoppingListEntryRequest, opts ...grpc.CallOption) (*GetUserShoppingListEntryResponse, error) {
	out := new(GetUserShoppingListEntryResponse)
	err := c.cc.Invoke(ctx, UserShoppingListEntryService_GetUserShoppingListEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListEntryServiceClient) GetAllUserShoppingListEntries(ctx context.Context, in *GetAllUserShoppingListEntriesRequest, opts ...grpc.CallOption) (*GetAllUserShoppingListEntriesResponse, error) {
	out := new(GetAllUserShoppingListEntriesResponse)
	err := c.cc.Invoke(ctx, UserShoppingListEntryService_GetAllUserShoppingListEntries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListEntryServiceClient) UpdateUserShoppingListEntry(ctx context.Context, in *UpdateUserShoppingListEntryRequest, opts ...grpc.CallOption) (*UpdateUserShoppingListEntryResponse, error) {
	out := new(UpdateUserShoppingListEntryResponse)
	err := c.cc.Invoke(ctx, UserShoppingListEntryService_UpdateUserShoppingListEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userShoppingListEntryServiceClient) DeleteUserShoppingListEntry(ctx context.Context, in *DeleteUserShoppingListEntryRequest, opts ...grpc.CallOption) (*DeleteUserShoppingListEntryResponse, error) {
	out := new(DeleteUserShoppingListEntryResponse)
	err := c.cc.Invoke(ctx, UserShoppingListEntryService_DeleteUserShoppingListEntry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserShoppingListEntryServiceServer is the server API for UserShoppingListEntryService service.
// All implementations must embed UnimplementedUserShoppingListEntryServiceServer
// for forward compatibility
type UserShoppingListEntryServiceServer interface {
	// create method for adding a new shopping list entry
	CreateUserShoppingListEntry(context.Context, *CreateUserShoppingListEntryRequest) (*CreateUserShoppingListEntryResponse, error)
	// read method for getting a specific shopping list entry
	GetUserShoppingListEntry(context.Context, *GetUserShoppingListEntryRequest) (*GetUserShoppingListEntryResponse, error)
	// read method for getting all shopping list entry for a specific shopping list
	GetAllUserShoppingListEntries(context.Context, *GetAllUserShoppingListEntriesRequest) (*GetAllUserShoppingListEntriesResponse, error)
	// update method for updating a shopping list
	UpdateUserShoppingListEntry(context.Context, *UpdateUserShoppingListEntryRequest) (*UpdateUserShoppingListEntryResponse, error)
	// delete method to get rid of a shopping list
	DeleteUserShoppingListEntry(context.Context, *DeleteUserShoppingListEntryRequest) (*DeleteUserShoppingListEntryResponse, error)
	mustEmbedUnimplementedUserShoppingListEntryServiceServer()
}

// UnimplementedUserShoppingListEntryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserShoppingListEntryServiceServer struct {
}

func (UnimplementedUserShoppingListEntryServiceServer) CreateUserShoppingListEntry(context.Context, *CreateUserShoppingListEntryRequest) (*CreateUserShoppingListEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserShoppingListEntry not implemented")
}
func (UnimplementedUserShoppingListEntryServiceServer) GetUserShoppingListEntry(context.Context, *GetUserShoppingListEntryRequest) (*GetUserShoppingListEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserShoppingListEntry not implemented")
}
func (UnimplementedUserShoppingListEntryServiceServer) GetAllUserShoppingListEntries(context.Context, *GetAllUserShoppingListEntriesRequest) (*GetAllUserShoppingListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserShoppingListEntries not implemented")
}
func (UnimplementedUserShoppingListEntryServiceServer) UpdateUserShoppingListEntry(context.Context, *UpdateUserShoppingListEntryRequest) (*UpdateUserShoppingListEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserShoppingListEntry not implemented")
}
func (UnimplementedUserShoppingListEntryServiceServer) DeleteUserShoppingListEntry(context.Context, *DeleteUserShoppingListEntryRequest) (*DeleteUserShoppingListEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserShoppingListEntry not implemented")
}
func (UnimplementedUserShoppingListEntryServiceServer) mustEmbedUnimplementedUserShoppingListEntryServiceServer() {
}

// UnsafeUserShoppingListEntryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserShoppingListEntryServiceServer will
// result in compilation errors.
type UnsafeUserShoppingListEntryServiceServer interface {
	mustEmbedUnimplementedUserShoppingListEntryServiceServer()
}

func RegisterUserShoppingListEntryServiceServer(s grpc.ServiceRegistrar, srv UserShoppingListEntryServiceServer) {
	s.RegisterService(&UserShoppingListEntryService_ServiceDesc, srv)
}

func _UserShoppingListEntryService_CreateUserShoppingListEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserShoppingListEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListEntryServiceServer).CreateUserShoppingListEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListEntryService_CreateUserShoppingListEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListEntryServiceServer).CreateUserShoppingListEntry(ctx, req.(*CreateUserShoppingListEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListEntryService_GetUserShoppingListEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserShoppingListEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListEntryServiceServer).GetUserShoppingListEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListEntryService_GetUserShoppingListEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListEntryServiceServer).GetUserShoppingListEntry(ctx, req.(*GetUserShoppingListEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListEntryService_GetAllUserShoppingListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserShoppingListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListEntryServiceServer).GetAllUserShoppingListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListEntryService_GetAllUserShoppingListEntries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListEntryServiceServer).GetAllUserShoppingListEntries(ctx, req.(*GetAllUserShoppingListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListEntryService_UpdateUserShoppingListEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserShoppingListEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListEntryServiceServer).UpdateUserShoppingListEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListEntryService_UpdateUserShoppingListEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListEntryServiceServer).UpdateUserShoppingListEntry(ctx, req.(*UpdateUserShoppingListEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserShoppingListEntryService_DeleteUserShoppingListEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserShoppingListEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserShoppingListEntryServiceServer).DeleteUserShoppingListEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserShoppingListEntryService_DeleteUserShoppingListEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserShoppingListEntryServiceServer).DeleteUserShoppingListEntry(ctx, req.(*DeleteUserShoppingListEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserShoppingListEntryService_ServiceDesc is the grpc.ServiceDesc for UserShoppingListEntryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserShoppingListEntryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shoppinglist.UserShoppingListEntryService",
	HandlerType: (*UserShoppingListEntryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserShoppingListEntry",
			Handler:    _UserShoppingListEntryService_CreateUserShoppingListEntry_Handler,
		},
		{
			MethodName: "GetUserShoppingListEntry",
			Handler:    _UserShoppingListEntryService_GetUserShoppingListEntry_Handler,
		},
		{
			MethodName: "GetAllUserShoppingListEntries",
			Handler:    _UserShoppingListEntryService_GetAllUserShoppingListEntries_Handler,
		},
		{
			MethodName: "UpdateUserShoppingListEntry",
			Handler:    _UserShoppingListEntryService_UpdateUserShoppingListEntry_Handler,
		},
		{
			MethodName: "DeleteUserShoppingListEntry",
			Handler:    _UserShoppingListEntryService_DeleteUserShoppingListEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppinglist/shoppinglist.proto",
}
