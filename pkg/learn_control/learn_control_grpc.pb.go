// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: learn_control.proto

package learn_control

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
	LearnControl_GetEmployeesByName_FullMethodName                   = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/GetEmployeesByName"
	LearnControl_GetEmployeePersonalCard_FullMethodName              = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/GetEmployeePersonalCard"
	LearnControl_UpdateEmployeeTrainingDate_FullMethodName           = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/UpdateEmployeeTrainingDate"
	LearnControl_GetDepartments_FullMethodName                       = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/GetDepartments"
	LearnControl_GetPositions_FullMethodName                         = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/GetPositions"
	LearnControl_GetTrainings_FullMethodName                         = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/GetTrainings"
	LearnControl_GetEmployeesByFilters_FullMethodName                = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/GetEmployeesByFilters"
	LearnControl_AddEmployee_FullMethodName                          = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/AddEmployee"
	LearnControl_GetTasksByFilters_FullMethodName                    = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/GetTasksByFilters"
	LearnControl_CloseAssignTask_FullMethodName                      = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/CloseAssignTask"
	LearnControl_CloseTaskWithTrainingDateSet_FullMethodName         = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/CloseTaskWithTrainingDateSet"
	LearnControl_CloseTaskWithPositionTrainingsSet_FullMethodName    = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/CloseTaskWithPositionTrainingsSet"
	LearnControl_CloseTaskWithTrainingProtocolConfirm_FullMethodName = "/github.com.DarYur13.learn_control.api.learn_control.LearnControl/CloseTaskWithTrainingProtocolConfirm"
)

// LearnControlClient is the client API for LearnControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearnControlClient interface {
	GetEmployeesByName(ctx context.Context, in *GetEmployeesByNameRequest, opts ...grpc.CallOption) (*GetEmployeesByNameResponse, error)
	GetEmployeePersonalCard(ctx context.Context, in *GetEmployeePersonalCardRequest, opts ...grpc.CallOption) (*GetEmployeePersonalCardResponse, error)
	UpdateEmployeeTrainingDate(ctx context.Context, in *UpdateEmployeeTrainingDateRequest, opts ...grpc.CallOption) (*UpdateEmployeeTrainingDateResponse, error)
	GetDepartments(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDepartmentsResponse, error)
	GetPositions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPositionsResponse, error)
	GetTrainings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTrainingsResponse, error)
	GetEmployeesByFilters(ctx context.Context, in *GetEmployeesByFiltersRequest, opts ...grpc.CallOption) (*GetEmployeesByFiltersResponse, error)
	AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTasksByFilters(ctx context.Context, in *GetTasksByFiltersRequest, opts ...grpc.CallOption) (*GetTasksByFiltersResponse, error)
	CloseAssignTask(ctx context.Context, in *CloseAssignTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CloseTaskWithTrainingDateSet(ctx context.Context, in *CloseTaskWithTrainingDateSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CloseTaskWithPositionTrainingsSet(ctx context.Context, in *CloseTaskWithPositionTrainingsSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CloseTaskWithTrainingProtocolConfirm(ctx context.Context, in *CloseTaskWithTrainingProtocolConfirmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type learnControlClient struct {
	cc grpc.ClientConnInterface
}

func NewLearnControlClient(cc grpc.ClientConnInterface) LearnControlClient {
	return &learnControlClient{cc}
}

func (c *learnControlClient) GetEmployeesByName(ctx context.Context, in *GetEmployeesByNameRequest, opts ...grpc.CallOption) (*GetEmployeesByNameResponse, error) {
	out := new(GetEmployeesByNameResponse)
	err := c.cc.Invoke(ctx, LearnControl_GetEmployeesByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) GetEmployeePersonalCard(ctx context.Context, in *GetEmployeePersonalCardRequest, opts ...grpc.CallOption) (*GetEmployeePersonalCardResponse, error) {
	out := new(GetEmployeePersonalCardResponse)
	err := c.cc.Invoke(ctx, LearnControl_GetEmployeePersonalCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) UpdateEmployeeTrainingDate(ctx context.Context, in *UpdateEmployeeTrainingDateRequest, opts ...grpc.CallOption) (*UpdateEmployeeTrainingDateResponse, error) {
	out := new(UpdateEmployeeTrainingDateResponse)
	err := c.cc.Invoke(ctx, LearnControl_UpdateEmployeeTrainingDate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) GetDepartments(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDepartmentsResponse, error) {
	out := new(GetDepartmentsResponse)
	err := c.cc.Invoke(ctx, LearnControl_GetDepartments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) GetPositions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPositionsResponse, error) {
	out := new(GetPositionsResponse)
	err := c.cc.Invoke(ctx, LearnControl_GetPositions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) GetTrainings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTrainingsResponse, error) {
	out := new(GetTrainingsResponse)
	err := c.cc.Invoke(ctx, LearnControl_GetTrainings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) GetEmployeesByFilters(ctx context.Context, in *GetEmployeesByFiltersRequest, opts ...grpc.CallOption) (*GetEmployeesByFiltersResponse, error) {
	out := new(GetEmployeesByFiltersResponse)
	err := c.cc.Invoke(ctx, LearnControl_GetEmployeesByFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) AddEmployee(ctx context.Context, in *AddEmployeeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LearnControl_AddEmployee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) GetTasksByFilters(ctx context.Context, in *GetTasksByFiltersRequest, opts ...grpc.CallOption) (*GetTasksByFiltersResponse, error) {
	out := new(GetTasksByFiltersResponse)
	err := c.cc.Invoke(ctx, LearnControl_GetTasksByFilters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) CloseAssignTask(ctx context.Context, in *CloseAssignTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LearnControl_CloseAssignTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) CloseTaskWithTrainingDateSet(ctx context.Context, in *CloseTaskWithTrainingDateSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LearnControl_CloseTaskWithTrainingDateSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) CloseTaskWithPositionTrainingsSet(ctx context.Context, in *CloseTaskWithPositionTrainingsSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LearnControl_CloseTaskWithPositionTrainingsSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnControlClient) CloseTaskWithTrainingProtocolConfirm(ctx context.Context, in *CloseTaskWithTrainingProtocolConfirmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LearnControl_CloseTaskWithTrainingProtocolConfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LearnControlServer is the server API for LearnControl service.
// All implementations must embed UnimplementedLearnControlServer
// for forward compatibility
type LearnControlServer interface {
	GetEmployeesByName(context.Context, *GetEmployeesByNameRequest) (*GetEmployeesByNameResponse, error)
	GetEmployeePersonalCard(context.Context, *GetEmployeePersonalCardRequest) (*GetEmployeePersonalCardResponse, error)
	UpdateEmployeeTrainingDate(context.Context, *UpdateEmployeeTrainingDateRequest) (*UpdateEmployeeTrainingDateResponse, error)
	GetDepartments(context.Context, *emptypb.Empty) (*GetDepartmentsResponse, error)
	GetPositions(context.Context, *emptypb.Empty) (*GetPositionsResponse, error)
	GetTrainings(context.Context, *emptypb.Empty) (*GetTrainingsResponse, error)
	GetEmployeesByFilters(context.Context, *GetEmployeesByFiltersRequest) (*GetEmployeesByFiltersResponse, error)
	AddEmployee(context.Context, *AddEmployeeRequest) (*emptypb.Empty, error)
	GetTasksByFilters(context.Context, *GetTasksByFiltersRequest) (*GetTasksByFiltersResponse, error)
	CloseAssignTask(context.Context, *CloseAssignTaskRequest) (*emptypb.Empty, error)
	CloseTaskWithTrainingDateSet(context.Context, *CloseTaskWithTrainingDateSetRequest) (*emptypb.Empty, error)
	CloseTaskWithPositionTrainingsSet(context.Context, *CloseTaskWithPositionTrainingsSetRequest) (*emptypb.Empty, error)
	CloseTaskWithTrainingProtocolConfirm(context.Context, *CloseTaskWithTrainingProtocolConfirmRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLearnControlServer()
}

// UnimplementedLearnControlServer must be embedded to have forward compatible implementations.
type UnimplementedLearnControlServer struct {
}

func (UnimplementedLearnControlServer) GetEmployeesByName(context.Context, *GetEmployeesByNameRequest) (*GetEmployeesByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeesByName not implemented")
}
func (UnimplementedLearnControlServer) GetEmployeePersonalCard(context.Context, *GetEmployeePersonalCardRequest) (*GetEmployeePersonalCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeePersonalCard not implemented")
}
func (UnimplementedLearnControlServer) UpdateEmployeeTrainingDate(context.Context, *UpdateEmployeeTrainingDateRequest) (*UpdateEmployeeTrainingDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployeeTrainingDate not implemented")
}
func (UnimplementedLearnControlServer) GetDepartments(context.Context, *emptypb.Empty) (*GetDepartmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartments not implemented")
}
func (UnimplementedLearnControlServer) GetPositions(context.Context, *emptypb.Empty) (*GetPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPositions not implemented")
}
func (UnimplementedLearnControlServer) GetTrainings(context.Context, *emptypb.Empty) (*GetTrainingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainings not implemented")
}
func (UnimplementedLearnControlServer) GetEmployeesByFilters(context.Context, *GetEmployeesByFiltersRequest) (*GetEmployeesByFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeesByFilters not implemented")
}
func (UnimplementedLearnControlServer) AddEmployee(context.Context, *AddEmployeeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedLearnControlServer) GetTasksByFilters(context.Context, *GetTasksByFiltersRequest) (*GetTasksByFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasksByFilters not implemented")
}
func (UnimplementedLearnControlServer) CloseAssignTask(context.Context, *CloseAssignTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseAssignTask not implemented")
}
func (UnimplementedLearnControlServer) CloseTaskWithTrainingDateSet(context.Context, *CloseTaskWithTrainingDateSetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTaskWithTrainingDateSet not implemented")
}
func (UnimplementedLearnControlServer) CloseTaskWithPositionTrainingsSet(context.Context, *CloseTaskWithPositionTrainingsSetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTaskWithPositionTrainingsSet not implemented")
}
func (UnimplementedLearnControlServer) CloseTaskWithTrainingProtocolConfirm(context.Context, *CloseTaskWithTrainingProtocolConfirmRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTaskWithTrainingProtocolConfirm not implemented")
}
func (UnimplementedLearnControlServer) mustEmbedUnimplementedLearnControlServer() {}

// UnsafeLearnControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearnControlServer will
// result in compilation errors.
type UnsafeLearnControlServer interface {
	mustEmbedUnimplementedLearnControlServer()
}

func RegisterLearnControlServer(s grpc.ServiceRegistrar, srv LearnControlServer) {
	s.RegisterService(&LearnControl_ServiceDesc, srv)
}

func _LearnControl_GetEmployeesByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeesByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).GetEmployeesByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_GetEmployeesByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).GetEmployeesByName(ctx, req.(*GetEmployeesByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_GetEmployeePersonalCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeePersonalCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).GetEmployeePersonalCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_GetEmployeePersonalCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).GetEmployeePersonalCard(ctx, req.(*GetEmployeePersonalCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_UpdateEmployeeTrainingDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmployeeTrainingDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).UpdateEmployeeTrainingDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_UpdateEmployeeTrainingDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).UpdateEmployeeTrainingDate(ctx, req.(*UpdateEmployeeTrainingDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_GetDepartments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).GetDepartments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_GetDepartments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).GetDepartments(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_GetPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).GetPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_GetPositions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).GetPositions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_GetTrainings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).GetTrainings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_GetTrainings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).GetTrainings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_GetEmployeesByFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmployeesByFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).GetEmployeesByFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_GetEmployeesByFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).GetEmployeesByFilters(ctx, req.(*GetEmployeesByFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_AddEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).AddEmployee(ctx, req.(*AddEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_GetTasksByFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTasksByFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).GetTasksByFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_GetTasksByFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).GetTasksByFilters(ctx, req.(*GetTasksByFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_CloseAssignTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseAssignTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).CloseAssignTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_CloseAssignTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).CloseAssignTask(ctx, req.(*CloseAssignTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_CloseTaskWithTrainingDateSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseTaskWithTrainingDateSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).CloseTaskWithTrainingDateSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_CloseTaskWithTrainingDateSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).CloseTaskWithTrainingDateSet(ctx, req.(*CloseTaskWithTrainingDateSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_CloseTaskWithPositionTrainingsSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseTaskWithPositionTrainingsSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).CloseTaskWithPositionTrainingsSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_CloseTaskWithPositionTrainingsSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).CloseTaskWithPositionTrainingsSet(ctx, req.(*CloseTaskWithPositionTrainingsSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearnControl_CloseTaskWithTrainingProtocolConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseTaskWithTrainingProtocolConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnControlServer).CloseTaskWithTrainingProtocolConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearnControl_CloseTaskWithTrainingProtocolConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnControlServer).CloseTaskWithTrainingProtocolConfirm(ctx, req.(*CloseTaskWithTrainingProtocolConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LearnControl_ServiceDesc is the grpc.ServiceDesc for LearnControl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LearnControl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.DarYur13.learn_control.api.learn_control.LearnControl",
	HandlerType: (*LearnControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployeesByName",
			Handler:    _LearnControl_GetEmployeesByName_Handler,
		},
		{
			MethodName: "GetEmployeePersonalCard",
			Handler:    _LearnControl_GetEmployeePersonalCard_Handler,
		},
		{
			MethodName: "UpdateEmployeeTrainingDate",
			Handler:    _LearnControl_UpdateEmployeeTrainingDate_Handler,
		},
		{
			MethodName: "GetDepartments",
			Handler:    _LearnControl_GetDepartments_Handler,
		},
		{
			MethodName: "GetPositions",
			Handler:    _LearnControl_GetPositions_Handler,
		},
		{
			MethodName: "GetTrainings",
			Handler:    _LearnControl_GetTrainings_Handler,
		},
		{
			MethodName: "GetEmployeesByFilters",
			Handler:    _LearnControl_GetEmployeesByFilters_Handler,
		},
		{
			MethodName: "AddEmployee",
			Handler:    _LearnControl_AddEmployee_Handler,
		},
		{
			MethodName: "GetTasksByFilters",
			Handler:    _LearnControl_GetTasksByFilters_Handler,
		},
		{
			MethodName: "CloseAssignTask",
			Handler:    _LearnControl_CloseAssignTask_Handler,
		},
		{
			MethodName: "CloseTaskWithTrainingDateSet",
			Handler:    _LearnControl_CloseTaskWithTrainingDateSet_Handler,
		},
		{
			MethodName: "CloseTaskWithPositionTrainingsSet",
			Handler:    _LearnControl_CloseTaskWithPositionTrainingsSet_Handler,
		},
		{
			MethodName: "CloseTaskWithTrainingProtocolConfirm",
			Handler:    _LearnControl_CloseTaskWithTrainingProtocolConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "learn_control.proto",
}
