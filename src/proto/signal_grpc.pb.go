// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SignalServiceClient is the client API for SignalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignalServiceClient interface {
	About(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AboutResponse, error)
	GetConfiguration(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	SetConfiguration(ctx context.Context, in *SetConfigurationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	RegisterNumber(ctx context.Context, in *RegisterNumberRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	VerifyRegisteredNumber(ctx context.Context, in *VerifyRegisteredNumberRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error)
	GetGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	DeleteGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	JoinGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	BlockGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetQrCodeLink(ctx context.Context, in *GetQrCodeLinkRequest, opts ...grpc.CallOption) (*GetQrCodeLinkResponse, error)
	GetAttachments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAttachmentsResponse, error)
	RemoveAttachment(ctx context.Context, in *RemoveAttachmentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ServeAttachment(ctx context.Context, in *ServeAttachmentRequest, opts ...grpc.CallOption) (*ServeAttachmentResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListIdentities(ctx context.Context, in *ListIdentitiesRequest, opts ...grpc.CallOption) (*ListIdentitiesResponse, error)
	TrustIdentity(ctx context.Context, in *TrustIdentityRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SendV2(ctx context.Context, in *SendV2Request, opts ...grpc.CallOption) (*SendResponse, error)
}

type signalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignalServiceClient(cc grpc.ClientConnInterface) SignalServiceClient {
	return &signalServiceClient{cc}
}

func (c *signalServiceClient) About(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AboutResponse, error) {
	out := new(AboutResponse)
	err := c.cc.Invoke(ctx, "/SignalService/About", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) GetConfiguration(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, "/SignalService/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) SetConfiguration(ctx context.Context, in *SetConfigurationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/SetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) RegisterNumber(ctx context.Context, in *RegisterNumberRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/RegisterNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) VerifyRegisteredNumber(ctx context.Context, in *VerifyRegisteredNumberRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/VerifyRegisteredNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/SignalService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error) {
	out := new(ReceiveResponse)
	err := c.cc.Invoke(ctx, "/SignalService/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/SignalService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	out := new(GetGroupsResponse)
	err := c.cc.Invoke(ctx, "/SignalService/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) GetGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, "/SignalService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) DeleteGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) JoinGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/JoinGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) BlockGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/BlockGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) GetQrCodeLink(ctx context.Context, in *GetQrCodeLinkRequest, opts ...grpc.CallOption) (*GetQrCodeLinkResponse, error) {
	out := new(GetQrCodeLinkResponse)
	err := c.cc.Invoke(ctx, "/SignalService/GetQrCodeLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) GetAttachments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAttachmentsResponse, error) {
	out := new(GetAttachmentsResponse)
	err := c.cc.Invoke(ctx, "/SignalService/GetAttachments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) RemoveAttachment(ctx context.Context, in *RemoveAttachmentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/RemoveAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) ServeAttachment(ctx context.Context, in *ServeAttachmentRequest, opts ...grpc.CallOption) (*ServeAttachmentResponse, error) {
	out := new(ServeAttachmentResponse)
	err := c.cc.Invoke(ctx, "/SignalService/ServeAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) ListIdentities(ctx context.Context, in *ListIdentitiesRequest, opts ...grpc.CallOption) (*ListIdentitiesResponse, error) {
	out := new(ListIdentitiesResponse)
	err := c.cc.Invoke(ctx, "/SignalService/ListIdentities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) TrustIdentity(ctx context.Context, in *TrustIdentityRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/SignalService/TrustIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) SendV2(ctx context.Context, in *SendV2Request, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/SignalService/SendV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalServiceServer is the server API for SignalService service.
// All implementations must embed UnimplementedSignalServiceServer
// for forward compatibility
type SignalServiceServer interface {
	About(context.Context, *empty.Empty) (*AboutResponse, error)
	GetConfiguration(context.Context, *empty.Empty) (*GetConfigurationResponse, error)
	SetConfiguration(context.Context, *SetConfigurationRequest) (*empty.Empty, error)
	Health(context.Context, *empty.Empty) (*empty.Empty, error)
	RegisterNumber(context.Context, *RegisterNumberRequest) (*empty.Empty, error)
	VerifyRegisteredNumber(context.Context, *VerifyRegisteredNumberRequest) (*empty.Empty, error)
	Send(context.Context, *SendRequest) (*SendResponse, error)
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error)
	GetGroup(context.Context, *GroupRequest) (*GetGroupResponse, error)
	DeleteGroup(context.Context, *GroupRequest) (*empty.Empty, error)
	JoinGroup(context.Context, *GroupRequest) (*empty.Empty, error)
	BlockGroup(context.Context, *GroupRequest) (*empty.Empty, error)
	GetQrCodeLink(context.Context, *GetQrCodeLinkRequest) (*GetQrCodeLinkResponse, error)
	GetAttachments(context.Context, *empty.Empty) (*GetAttachmentsResponse, error)
	RemoveAttachment(context.Context, *RemoveAttachmentRequest) (*empty.Empty, error)
	ServeAttachment(context.Context, *ServeAttachmentRequest) (*ServeAttachmentResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*empty.Empty, error)
	ListIdentities(context.Context, *ListIdentitiesRequest) (*ListIdentitiesResponse, error)
	TrustIdentity(context.Context, *TrustIdentityRequest) (*empty.Empty, error)
	SendV2(context.Context, *SendV2Request) (*SendResponse, error)
	mustEmbedUnimplementedSignalServiceServer()
}

// UnimplementedSignalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSignalServiceServer struct {
}

func (UnimplementedSignalServiceServer) About(context.Context, *empty.Empty) (*AboutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method About not implemented")
}
func (UnimplementedSignalServiceServer) GetConfiguration(context.Context, *empty.Empty) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedSignalServiceServer) SetConfiguration(context.Context, *SetConfigurationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfiguration not implemented")
}
func (UnimplementedSignalServiceServer) Health(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedSignalServiceServer) RegisterNumber(context.Context, *RegisterNumberRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNumber not implemented")
}
func (UnimplementedSignalServiceServer) VerifyRegisteredNumber(context.Context, *VerifyRegisteredNumberRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRegisteredNumber not implemented")
}
func (UnimplementedSignalServiceServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSignalServiceServer) Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedSignalServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedSignalServiceServer) GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedSignalServiceServer) GetGroup(context.Context, *GroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedSignalServiceServer) DeleteGroup(context.Context, *GroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedSignalServiceServer) JoinGroup(context.Context, *GroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (UnimplementedSignalServiceServer) BlockGroup(context.Context, *GroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockGroup not implemented")
}
func (UnimplementedSignalServiceServer) GetQrCodeLink(context.Context, *GetQrCodeLinkRequest) (*GetQrCodeLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQrCodeLink not implemented")
}
func (UnimplementedSignalServiceServer) GetAttachments(context.Context, *empty.Empty) (*GetAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachments not implemented")
}
func (UnimplementedSignalServiceServer) RemoveAttachment(context.Context, *RemoveAttachmentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAttachment not implemented")
}
func (UnimplementedSignalServiceServer) ServeAttachment(context.Context, *ServeAttachmentRequest) (*ServeAttachmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeAttachment not implemented")
}
func (UnimplementedSignalServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedSignalServiceServer) ListIdentities(context.Context, *ListIdentitiesRequest) (*ListIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIdentities not implemented")
}
func (UnimplementedSignalServiceServer) TrustIdentity(context.Context, *TrustIdentityRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrustIdentity not implemented")
}
func (UnimplementedSignalServiceServer) SendV2(context.Context, *SendV2Request) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendV2 not implemented")
}
func (UnimplementedSignalServiceServer) mustEmbedUnimplementedSignalServiceServer() {}

// UnsafeSignalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignalServiceServer will
// result in compilation errors.
type UnsafeSignalServiceServer interface {
	mustEmbedUnimplementedSignalServiceServer()
}

func RegisterSignalServiceServer(s grpc.ServiceRegistrar, srv SignalServiceServer) {
	s.RegisterService(&SignalService_ServiceDesc, srv)
}

func _SignalService_About_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).About(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/About",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).About(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetConfiguration(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_SetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).SetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/SetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).SetConfiguration(ctx, req.(*SetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_RegisterNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).RegisterNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/RegisterNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).RegisterNumber(ctx, req.(*RegisterNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_VerifyRegisteredNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRegisteredNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).VerifyRegisteredNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/VerifyRegisteredNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).VerifyRegisteredNumber(ctx, req.(*VerifyRegisteredNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetGroups(ctx, req.(*GetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).DeleteGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/JoinGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).JoinGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_BlockGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).BlockGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/BlockGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).BlockGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_GetQrCodeLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQrCodeLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetQrCodeLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/GetQrCodeLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetQrCodeLink(ctx, req.(*GetQrCodeLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_GetAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/GetAttachments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetAttachments(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_RemoveAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).RemoveAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/RemoveAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).RemoveAttachment(ctx, req.(*RemoveAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_ServeAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServeAttachmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).ServeAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/ServeAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).ServeAttachment(ctx, req.(*ServeAttachmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_ListIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).ListIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/ListIdentities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).ListIdentities(ctx, req.(*ListIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_TrustIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).TrustIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/TrustIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).TrustIdentity(ctx, req.(*TrustIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_SendV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).SendV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignalService/SendV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).SendV2(ctx, req.(*SendV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

// SignalService_ServiceDesc is the grpc.ServiceDesc for SignalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SignalService",
	HandlerType: (*SignalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "About",
			Handler:    _SignalService_About_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _SignalService_GetConfiguration_Handler,
		},
		{
			MethodName: "SetConfiguration",
			Handler:    _SignalService_SetConfiguration_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _SignalService_Health_Handler,
		},
		{
			MethodName: "RegisterNumber",
			Handler:    _SignalService_RegisterNumber_Handler,
		},
		{
			MethodName: "VerifyRegisteredNumber",
			Handler:    _SignalService_VerifyRegisteredNumber_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _SignalService_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _SignalService_Receive_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _SignalService_CreateGroup_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _SignalService_GetGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _SignalService_GetGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _SignalService_DeleteGroup_Handler,
		},
		{
			MethodName: "JoinGroup",
			Handler:    _SignalService_JoinGroup_Handler,
		},
		{
			MethodName: "BlockGroup",
			Handler:    _SignalService_BlockGroup_Handler,
		},
		{
			MethodName: "GetQrCodeLink",
			Handler:    _SignalService_GetQrCodeLink_Handler,
		},
		{
			MethodName: "GetAttachments",
			Handler:    _SignalService_GetAttachments_Handler,
		},
		{
			MethodName: "RemoveAttachment",
			Handler:    _SignalService_RemoveAttachment_Handler,
		},
		{
			MethodName: "ServeAttachment",
			Handler:    _SignalService_ServeAttachment_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _SignalService_UpdateProfile_Handler,
		},
		{
			MethodName: "ListIdentities",
			Handler:    _SignalService_ListIdentities_Handler,
		},
		{
			MethodName: "TrustIdentity",
			Handler:    _SignalService_TrustIdentity_Handler,
		},
		{
			MethodName: "SendV2",
			Handler:    _SignalService_SendV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/signal.proto",
}
