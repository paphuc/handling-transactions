// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package employee

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

// EmployeeClient is the client API for Employee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeClient interface {
	InsertEmployee(ctx context.Context, in *InsertEmployeeRequest, opts ...grpc.CallOption) (*InsertEmployeeResponse, error)
	InsertEmployeeDetail(ctx context.Context, in *InsertEmployeeDetailRequest, opts ...grpc.CallOption) (*InsertEmployeeDetailResponse, error)
}

type employeeClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeClient(cc grpc.ClientConnInterface) EmployeeClient {
	return &employeeClient{cc}
}

func (c *employeeClient) InsertEmployee(ctx context.Context, in *InsertEmployeeRequest, opts ...grpc.CallOption) (*InsertEmployeeResponse, error) {
	out := new(InsertEmployeeResponse)
	err := c.cc.Invoke(ctx, "/employee.Employee/InsertEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) InsertEmployeeDetail(ctx context.Context, in *InsertEmployeeDetailRequest, opts ...grpc.CallOption) (*InsertEmployeeDetailResponse, error) {
	out := new(InsertEmployeeDetailResponse)
	err := c.cc.Invoke(ctx, "/employee.Employee/InsertEmployeeDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServer is the server API for Employee service.
// All implementations should embed UnimplementedEmployeeServer
// for forward compatibility
type EmployeeServer interface {
	InsertEmployee(context.Context, *InsertEmployeeRequest) (*InsertEmployeeResponse, error)
	InsertEmployeeDetail(context.Context, *InsertEmployeeDetailRequest) (*InsertEmployeeDetailResponse, error)
}

// UnimplementedEmployeeServer should be embedded to have forward compatible implementations.
type UnimplementedEmployeeServer struct {
}

func (UnimplementedEmployeeServer) InsertEmployee(context.Context, *InsertEmployeeRequest) (*InsertEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEmployee not implemented")
}
func (UnimplementedEmployeeServer) InsertEmployeeDetail(context.Context, *InsertEmployeeDetailRequest) (*InsertEmployeeDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEmployeeDetail not implemented")
}

// UnsafeEmployeeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServer will
// result in compilation errors.
type UnsafeEmployeeServer interface {
	mustEmbedUnimplementedEmployeeServer()
}

func RegisterEmployeeServer(s grpc.ServiceRegistrar, srv EmployeeServer) {
	s.RegisterService(&Employee_ServiceDesc, srv)
}

func _Employee_InsertEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).InsertEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.Employee/InsertEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).InsertEmployee(ctx, req.(*InsertEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_InsertEmployeeDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertEmployeeDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).InsertEmployeeDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.Employee/InsertEmployeeDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).InsertEmployeeDetail(ctx, req.(*InsertEmployeeDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Employee_ServiceDesc is the grpc.ServiceDesc for Employee service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Employee_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee.Employee",
	HandlerType: (*EmployeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertEmployee",
			Handler:    _Employee_InsertEmployee_Handler,
		},
		{
			MethodName: "InsertEmployeeDetail",
			Handler:    _Employee_InsertEmployeeDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee/employee.proto",
}
