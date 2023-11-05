// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: admin.proto

package proto

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	GetSubsGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubscriptionsResponse, error)
	GetCouponsGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCouponsResponse, error)
	GetCategoriesGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCategoryResponse, error)
	GetRegionsGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegionResponse, error)
	GetCouponsByRegionGRPC(ctx context.Context, in *Region, opts ...grpc.CallOption) (*GetCouponsResponse, error)
	GetCouponsByCategoryGRPC(ctx context.Context, in *Category, opts ...grpc.CallOption) (*GetCouponsResponse, error)
	GetOrganizationInfo(ctx context.Context, in *InfoOrganizationRequest, opts ...grpc.CallOption) (*InfoOrganizationResponse, error)
	UpdateOrganizationInfo(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error)
	UpdateMembersInfo(ctx context.Context, in *UpdateMembersRequest, opts ...grpc.CallOption) (*UpdateMembersResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) GetSubsGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubscriptionsResponse, error) {
	out := new(SubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetSubsGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetCouponsGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCouponsResponse, error) {
	out := new(GetCouponsResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetCouponsGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetCategoriesGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCategoryResponse, error) {
	out := new(GetCategoryResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetCategoriesGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetRegionsGRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RegionResponse, error) {
	out := new(RegionResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetRegionsGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetCouponsByRegionGRPC(ctx context.Context, in *Region, opts ...grpc.CallOption) (*GetCouponsResponse, error) {
	out := new(GetCouponsResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetCouponsByRegionGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetCouponsByCategoryGRPC(ctx context.Context, in *Category, opts ...grpc.CallOption) (*GetCouponsResponse, error) {
	out := new(GetCouponsResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetCouponsByCategoryGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetOrganizationInfo(ctx context.Context, in *InfoOrganizationRequest, opts ...grpc.CallOption) (*InfoOrganizationResponse, error) {
	out := new(InfoOrganizationResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/GetOrganizationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateOrganizationInfo(ctx context.Context, in *UpdateOrganizationRequest, opts ...grpc.CallOption) (*UpdateOrganizationResponse, error) {
	out := new(UpdateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/UpdateOrganizationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateMembersInfo(ctx context.Context, in *UpdateMembersRequest, opts ...grpc.CallOption) (*UpdateMembersResponse, error) {
	out := new(UpdateMembersResponse)
	err := c.cc.Invoke(ctx, "/api.AdminService/UpdateMembersInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	GetSubsGRPC(context.Context, *Empty) (*SubscriptionsResponse, error)
	GetCouponsGRPC(context.Context, *Empty) (*GetCouponsResponse, error)
	GetCategoriesGRPC(context.Context, *Empty) (*GetCategoryResponse, error)
	GetRegionsGRPC(context.Context, *Empty) (*RegionResponse, error)
	GetCouponsByRegionGRPC(context.Context, *Region) (*GetCouponsResponse, error)
	GetCouponsByCategoryGRPC(context.Context, *Category) (*GetCouponsResponse, error)
	GetOrganizationInfo(context.Context, *InfoOrganizationRequest) (*InfoOrganizationResponse, error)
	UpdateOrganizationInfo(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error)
	UpdateMembersInfo(context.Context, *UpdateMembersRequest) (*UpdateMembersResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) GetSubsGRPC(context.Context, *Empty) (*SubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubsGRPC not implemented")
}
func (UnimplementedAdminServiceServer) GetCouponsGRPC(context.Context, *Empty) (*GetCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponsGRPC not implemented")
}
func (UnimplementedAdminServiceServer) GetCategoriesGRPC(context.Context, *Empty) (*GetCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoriesGRPC not implemented")
}
func (UnimplementedAdminServiceServer) GetRegionsGRPC(context.Context, *Empty) (*RegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegionsGRPC not implemented")
}
func (UnimplementedAdminServiceServer) GetCouponsByRegionGRPC(context.Context, *Region) (*GetCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponsByRegionGRPC not implemented")
}
func (UnimplementedAdminServiceServer) GetCouponsByCategoryGRPC(context.Context, *Category) (*GetCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponsByCategoryGRPC not implemented")
}
func (UnimplementedAdminServiceServer) GetOrganizationInfo(context.Context, *InfoOrganizationRequest) (*InfoOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationInfo not implemented")
}
func (UnimplementedAdminServiceServer) UpdateOrganizationInfo(context.Context, *UpdateOrganizationRequest) (*UpdateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganizationInfo not implemented")
}
func (UnimplementedAdminServiceServer) UpdateMembersInfo(context.Context, *UpdateMembersRequest) (*UpdateMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMembersInfo not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_GetSubsGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetSubsGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetSubsGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetSubsGRPC(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetCouponsGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetCouponsGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetCouponsGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetCouponsGRPC(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetCategoriesGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetCategoriesGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetCategoriesGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetCategoriesGRPC(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetRegionsGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetRegionsGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetRegionsGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetRegionsGRPC(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetCouponsByRegionGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Region)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetCouponsByRegionGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetCouponsByRegionGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetCouponsByRegionGRPC(ctx, req.(*Region))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetCouponsByCategoryGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetCouponsByCategoryGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetCouponsByCategoryGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetCouponsByCategoryGRPC(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetOrganizationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetOrganizationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/GetOrganizationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetOrganizationInfo(ctx, req.(*InfoOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateOrganizationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateOrganizationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/UpdateOrganizationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateOrganizationInfo(ctx, req.(*UpdateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateMembersInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateMembersInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AdminService/UpdateMembersInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateMembersInfo(ctx, req.(*UpdateMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubsGRPC",
			Handler:    _AdminService_GetSubsGRPC_Handler,
		},
		{
			MethodName: "GetCouponsGRPC",
			Handler:    _AdminService_GetCouponsGRPC_Handler,
		},
		{
			MethodName: "GetCategoriesGRPC",
			Handler:    _AdminService_GetCategoriesGRPC_Handler,
		},
		{
			MethodName: "GetRegionsGRPC",
			Handler:    _AdminService_GetRegionsGRPC_Handler,
		},
		{
			MethodName: "GetCouponsByRegionGRPC",
			Handler:    _AdminService_GetCouponsByRegionGRPC_Handler,
		},
		{
			MethodName: "GetCouponsByCategoryGRPC",
			Handler:    _AdminService_GetCouponsByCategoryGRPC_Handler,
		},
		{
			MethodName: "GetOrganizationInfo",
			Handler:    _AdminService_GetOrganizationInfo_Handler,
		},
		{
			MethodName: "UpdateOrganizationInfo",
			Handler:    _AdminService_UpdateOrganizationInfo_Handler,
		},
		{
			MethodName: "UpdateMembersInfo",
			Handler:    _AdminService_UpdateMembersInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
