// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// MovieSvcClient is the client API for MovieSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieSvcClient interface {
	PutMovie(ctx context.Context, in *PutMovieRequest, opts ...grpc.CallOption) (*PutMovieResponse, error)
	GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error)
	GetMovies(ctx context.Context, in *GetMoviesRequest, opts ...grpc.CallOption) (*GetMoviesResponse, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error)
}

type movieSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieSvcClient(cc grpc.ClientConnInterface) MovieSvcClient {
	return &movieSvcClient{cc}
}

func (c *movieSvcClient) PutMovie(ctx context.Context, in *PutMovieRequest, opts ...grpc.CallOption) (*PutMovieResponse, error) {
	out := new(PutMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieSvc/PutMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieSvcClient) GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error) {
	out := new(GetMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieSvc/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieSvcClient) GetMovies(ctx context.Context, in *GetMoviesRequest, opts ...grpc.CallOption) (*GetMoviesResponse, error) {
	out := new(GetMoviesResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieSvc/GetMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieSvcClient) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error) {
	out := new(DeleteMovieResponse)
	err := c.cc.Invoke(ctx, "/pb.MovieSvc/DeleteMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieSvcServer is the server API for MovieSvc service.
// All implementations must embed UnimplementedMovieSvcServer
// for forward compatibility
type MovieSvcServer interface {
	PutMovie(context.Context, *PutMovieRequest) (*PutMovieResponse, error)
	GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error)
	GetMovies(context.Context, *GetMoviesRequest) (*GetMoviesResponse, error)
	DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error)
	mustEmbedUnimplementedMovieSvcServer()
}

// UnimplementedMovieSvcServer must be embedded to have forward compatible implementations.
type UnimplementedMovieSvcServer struct {
}

func (UnimplementedMovieSvcServer) PutMovie(context.Context, *PutMovieRequest) (*PutMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutMovie not implemented")
}
func (UnimplementedMovieSvcServer) GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieSvcServer) GetMovies(context.Context, *GetMoviesRequest) (*GetMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMovieSvcServer) DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieSvcServer) mustEmbedUnimplementedMovieSvcServer() {}

// UnsafeMovieSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieSvcServer will
// result in compilation errors.
type UnsafeMovieSvcServer interface {
	mustEmbedUnimplementedMovieSvcServer()
}

func RegisterMovieSvcServer(s grpc.ServiceRegistrar, srv MovieSvcServer) {
	s.RegisterService(&MovieSvc_ServiceDesc, srv)
}

func _MovieSvc_PutMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieSvcServer).PutMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieSvc/PutMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieSvcServer).PutMovie(ctx, req.(*PutMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieSvc_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieSvcServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieSvc/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieSvcServer).GetMovie(ctx, req.(*GetMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieSvc_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieSvcServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieSvc/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieSvcServer).GetMovies(ctx, req.(*GetMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieSvc_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieSvcServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MovieSvc/DeleteMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieSvcServer).DeleteMovie(ctx, req.(*DeleteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieSvc_ServiceDesc is the grpc.ServiceDesc for MovieSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MovieSvc",
	HandlerType: (*MovieSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutMovie",
			Handler:    _MovieSvc_PutMovie_Handler,
		},
		{
			MethodName: "GetMovie",
			Handler:    _MovieSvc_GetMovie_Handler,
		},
		{
			MethodName: "GetMovies",
			Handler:    _MovieSvc_GetMovies_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieSvc_DeleteMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcapi.proto",
}
