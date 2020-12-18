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

// FileGetterClient is the client API for FileGetter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileGetterClient interface {
	GetFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error)
}

type fileGetterClient struct {
	cc grpc.ClientConnInterface
}

func NewFileGetterClient(cc grpc.ClientConnInterface) FileGetterClient {
	return &fileGetterClient{cc}
}

func (c *fileGetterClient) GetFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/FileGetter/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileGetterServer is the server API for FileGetter service.
// All implementations must embed UnimplementedFileGetterServer
// for forward compatibility
type FileGetterServer interface {
	GetFile(context.Context, *FileRequest) (*FileResponse, error)
	mustEmbedUnimplementedFileGetterServer()
}

// UnimplementedFileGetterServer must be embedded to have forward compatible implementations.
type UnimplementedFileGetterServer struct {
}

func (UnimplementedFileGetterServer) GetFile(context.Context, *FileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileGetterServer) mustEmbedUnimplementedFileGetterServer() {}

// UnsafeFileGetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileGetterServer will
// result in compilation errors.
type UnsafeFileGetterServer interface {
	mustEmbedUnimplementedFileGetterServer()
}

func RegisterFileGetterServer(s grpc.ServiceRegistrar, srv FileGetterServer) {
	s.RegisterService(&_FileGetter_serviceDesc, srv)
}

func _FileGetter_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileGetterServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileGetter/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileGetterServer).GetFile(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileGetter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FileGetter",
	HandlerType: (*FileGetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFile",
			Handler:    _FileGetter_GetFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filepc.proto",
}