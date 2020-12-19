// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package filepc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ExplorerClient is the client API for Explorer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExplorerClient interface {
	GetFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error)
}

type explorerClient struct {
	cc grpc.ClientConnInterface
}

func NewExplorerClient(cc grpc.ClientConnInterface) ExplorerClient {
	return &explorerClient{cc}
}

func (c *explorerClient) GetFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/main.Explorer/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExplorerServer is the server API for Explorer service.
// All implementations must embed UnimplementedExplorerServer
// for forward compatibility
type ExplorerServer interface {
	GetFile(context.Context, *FileRequest) (*FileResponse, error)
	mustEmbedUnimplementedExplorerServer()
}

// UnimplementedExplorerServer must be embedded to have forward compatible implementations.
type UnimplementedExplorerServer struct {
}

func (UnimplementedExplorerServer) GetFile(context.Context, *FileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedExplorerServer) mustEmbedUnimplementedExplorerServer() {}

// UnsafeExplorerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExplorerServer will
// result in compilation errors.
type UnsafeExplorerServer interface {
	mustEmbedUnimplementedExplorerServer()
}

func RegisterExplorerServer(s grpc.ServiceRegistrar, srv ExplorerServer) {
	s.RegisterService(&_Explorer_serviceDesc, srv)
}

func _Explorer_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Explorer/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServer).GetFile(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Explorer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Explorer",
	HandlerType: (*ExplorerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFile",
			Handler:    _Explorer_GetFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filepc.proto",
}
