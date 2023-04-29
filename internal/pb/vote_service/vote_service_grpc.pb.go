// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: api/vote_service.proto

package vote_service

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
	VoteService_Vote_FullMethodName = "/vote_service.VoteService/Vote"
)

// VoteServiceClient is the client API for VoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoteServiceClient interface {
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
}

type voteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteServiceClient(cc grpc.ClientConnInterface) VoteServiceClient {
	return &voteServiceClient{cc}
}

func (c *voteServiceClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, VoteService_Vote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoteServiceServer is the server API for VoteService service.
// All implementations must embed UnimplementedVoteServiceServer
// for forward compatibility
type VoteServiceServer interface {
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
	mustEmbedUnimplementedVoteServiceServer()
}

// UnimplementedVoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVoteServiceServer struct {
}

func (UnimplementedVoteServiceServer) Vote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedVoteServiceServer) mustEmbedUnimplementedVoteServiceServer() {}

// UnsafeVoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoteServiceServer will
// result in compilation errors.
type UnsafeVoteServiceServer interface {
	mustEmbedUnimplementedVoteServiceServer()
}

func RegisterVoteServiceServer(s grpc.ServiceRegistrar, srv VoteServiceServer) {
	s.RegisterService(&VoteService_ServiceDesc, srv)
}

func _VoteService_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_Vote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VoteService_ServiceDesc is the grpc.ServiceDesc for VoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vote_service.VoteService",
	HandlerType: (*VoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Vote",
			Handler:    _VoteService_Vote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/vote_service.proto",
}
