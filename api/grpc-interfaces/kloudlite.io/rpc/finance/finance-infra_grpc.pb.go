// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package finance

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

// FinanceInfraClient is the client API for FinanceInfra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinanceInfraClient interface {
	ComputeStart(ctx context.Context, in *ComputeStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	ComputeEnd(ctx context.Context, in *ComputeEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	LambdaStart(ctx context.Context, in *LambdaStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	LambdaEnd(ctx context.Context, in *LambdaEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	BlockStorageStart(ctx context.Context, in *BlockStorageStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	BlockStorageEnd(ctx context.Context, in *BlockStorageEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	ObjectStorageStart(ctx context.Context, in *ObjectStorageStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	ObjectStorageEnd(ctx context.Context, in *ObjectStorageEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	CIStart(ctx context.Context, in *CIStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
	CIEnd(ctx context.Context, in *CIEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error)
}

type financeInfraClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceInfraClient(cc grpc.ClientConnInterface) FinanceInfraClient {
	return &financeInfraClient{cc}
}

func (c *financeInfraClient) ComputeStart(ctx context.Context, in *ComputeStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/ComputeStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) ComputeEnd(ctx context.Context, in *ComputeEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/ComputeEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) LambdaStart(ctx context.Context, in *LambdaStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/LambdaStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) LambdaEnd(ctx context.Context, in *LambdaEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/LambdaEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) BlockStorageStart(ctx context.Context, in *BlockStorageStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/BlockStorageStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) BlockStorageEnd(ctx context.Context, in *BlockStorageEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/BlockStorageEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) ObjectStorageStart(ctx context.Context, in *ObjectStorageStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/ObjectStorageStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) ObjectStorageEnd(ctx context.Context, in *ObjectStorageEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/ObjectStorageEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) CIStart(ctx context.Context, in *CIStartIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/CIStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeInfraClient) CIEnd(ctx context.Context, in *CIEndIn, opts ...grpc.CallOption) (*FinanceInfraVoid, error) {
	out := new(FinanceInfraVoid)
	err := c.cc.Invoke(ctx, "/FinanceInfra/CIEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceInfraServer is the server API for FinanceInfra service.
// All implementations must embed UnimplementedFinanceInfraServer
// for forward compatibility
type FinanceInfraServer interface {
	ComputeStart(context.Context, *ComputeStartIn) (*FinanceInfraVoid, error)
	ComputeEnd(context.Context, *ComputeEndIn) (*FinanceInfraVoid, error)
	LambdaStart(context.Context, *LambdaStartIn) (*FinanceInfraVoid, error)
	LambdaEnd(context.Context, *LambdaEndIn) (*FinanceInfraVoid, error)
	BlockStorageStart(context.Context, *BlockStorageStartIn) (*FinanceInfraVoid, error)
	BlockStorageEnd(context.Context, *BlockStorageEndIn) (*FinanceInfraVoid, error)
	ObjectStorageStart(context.Context, *ObjectStorageStartIn) (*FinanceInfraVoid, error)
	ObjectStorageEnd(context.Context, *ObjectStorageEndIn) (*FinanceInfraVoid, error)
	CIStart(context.Context, *CIStartIn) (*FinanceInfraVoid, error)
	CIEnd(context.Context, *CIEndIn) (*FinanceInfraVoid, error)
	mustEmbedUnimplementedFinanceInfraServer()
}

// UnimplementedFinanceInfraServer must be embedded to have forward compatible implementations.
type UnimplementedFinanceInfraServer struct {
}

func (UnimplementedFinanceInfraServer) ComputeStart(context.Context, *ComputeStartIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeStart not implemented")
}
func (UnimplementedFinanceInfraServer) ComputeEnd(context.Context, *ComputeEndIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeEnd not implemented")
}
func (UnimplementedFinanceInfraServer) LambdaStart(context.Context, *LambdaStartIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LambdaStart not implemented")
}
func (UnimplementedFinanceInfraServer) LambdaEnd(context.Context, *LambdaEndIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LambdaEnd not implemented")
}
func (UnimplementedFinanceInfraServer) BlockStorageStart(context.Context, *BlockStorageStartIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockStorageStart not implemented")
}
func (UnimplementedFinanceInfraServer) BlockStorageEnd(context.Context, *BlockStorageEndIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockStorageEnd not implemented")
}
func (UnimplementedFinanceInfraServer) ObjectStorageStart(context.Context, *ObjectStorageStartIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObjectStorageStart not implemented")
}
func (UnimplementedFinanceInfraServer) ObjectStorageEnd(context.Context, *ObjectStorageEndIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObjectStorageEnd not implemented")
}
func (UnimplementedFinanceInfraServer) CIStart(context.Context, *CIStartIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CIStart not implemented")
}
func (UnimplementedFinanceInfraServer) CIEnd(context.Context, *CIEndIn) (*FinanceInfraVoid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CIEnd not implemented")
}
func (UnimplementedFinanceInfraServer) mustEmbedUnimplementedFinanceInfraServer() {}

// UnsafeFinanceInfraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinanceInfraServer will
// result in compilation errors.
type UnsafeFinanceInfraServer interface {
	mustEmbedUnimplementedFinanceInfraServer()
}

func RegisterFinanceInfraServer(s grpc.ServiceRegistrar, srv FinanceInfraServer) {
	s.RegisterService(&FinanceInfra_ServiceDesc, srv)
}

func _FinanceInfra_ComputeStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeStartIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).ComputeStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/ComputeStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).ComputeStart(ctx, req.(*ComputeStartIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_ComputeEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeEndIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).ComputeEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/ComputeEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).ComputeEnd(ctx, req.(*ComputeEndIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_LambdaStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LambdaStartIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).LambdaStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/LambdaStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).LambdaStart(ctx, req.(*LambdaStartIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_LambdaEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LambdaEndIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).LambdaEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/LambdaEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).LambdaEnd(ctx, req.(*LambdaEndIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_BlockStorageStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockStorageStartIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).BlockStorageStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/BlockStorageStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).BlockStorageStart(ctx, req.(*BlockStorageStartIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_BlockStorageEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockStorageEndIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).BlockStorageEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/BlockStorageEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).BlockStorageEnd(ctx, req.(*BlockStorageEndIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_ObjectStorageStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectStorageStartIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).ObjectStorageStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/ObjectStorageStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).ObjectStorageStart(ctx, req.(*ObjectStorageStartIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_ObjectStorageEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectStorageEndIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).ObjectStorageEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/ObjectStorageEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).ObjectStorageEnd(ctx, req.(*ObjectStorageEndIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_CIStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIStartIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).CIStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/CIStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).CIStart(ctx, req.(*CIStartIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceInfra_CIEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CIEndIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceInfraServer).CIEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FinanceInfra/CIEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceInfraServer).CIEnd(ctx, req.(*CIEndIn))
	}
	return interceptor(ctx, in, info, handler)
}

// FinanceInfra_ServiceDesc is the grpc.ServiceDesc for FinanceInfra service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinanceInfra_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FinanceInfra",
	HandlerType: (*FinanceInfraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeStart",
			Handler:    _FinanceInfra_ComputeStart_Handler,
		},
		{
			MethodName: "ComputeEnd",
			Handler:    _FinanceInfra_ComputeEnd_Handler,
		},
		{
			MethodName: "LambdaStart",
			Handler:    _FinanceInfra_LambdaStart_Handler,
		},
		{
			MethodName: "LambdaEnd",
			Handler:    _FinanceInfra_LambdaEnd_Handler,
		},
		{
			MethodName: "BlockStorageStart",
			Handler:    _FinanceInfra_BlockStorageStart_Handler,
		},
		{
			MethodName: "BlockStorageEnd",
			Handler:    _FinanceInfra_BlockStorageEnd_Handler,
		},
		{
			MethodName: "ObjectStorageStart",
			Handler:    _FinanceInfra_ObjectStorageStart_Handler,
		},
		{
			MethodName: "ObjectStorageEnd",
			Handler:    _FinanceInfra_ObjectStorageEnd_Handler,
		},
		{
			MethodName: "CIStart",
			Handler:    _FinanceInfra_CIStart_Handler,
		},
		{
			MethodName: "CIEnd",
			Handler:    _FinanceInfra_CIEnd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance-infra.proto",
}
