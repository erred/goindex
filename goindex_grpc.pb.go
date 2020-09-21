// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package goindex

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GoindexClient is the client API for Goindex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoindexClient interface {
	Versions(ctx context.Context, in *VersionsRequest, opts ...grpc.CallOption) (*ProjectVersions, error)
}

type goindexClient struct {
	cc grpc.ClientConnInterface
}

func NewGoindexClient(cc grpc.ClientConnInterface) GoindexClient {
	return &goindexClient{cc}
}

var goindexVersionsStreamDesc = &grpc.StreamDesc{
	StreamName: "Versions",
}

func (c *goindexClient) Versions(ctx context.Context, in *VersionsRequest, opts ...grpc.CallOption) (*ProjectVersions, error) {
	out := new(ProjectVersions)
	err := c.cc.Invoke(ctx, "/stream.Goindex/Versions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoindexService is the service API for Goindex service.
// Fields should be assigned to their respective handler implementations only before
// RegisterGoindexService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type GoindexService struct {
	Versions func(context.Context, *VersionsRequest) (*ProjectVersions, error)
}

func (s *GoindexService) versions(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Versions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/stream.Goindex/Versions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Versions(ctx, req.(*VersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterGoindexService registers a service implementation with a gRPC server.
func RegisterGoindexService(s grpc.ServiceRegistrar, srv *GoindexService) {
	srvCopy := *srv
	if srvCopy.Versions == nil {
		srvCopy.Versions = func(context.Context, *VersionsRequest) (*ProjectVersions, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Versions not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "stream.Goindex",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Versions",
				Handler:    srvCopy.versions,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "goindex.proto",
	}

	s.RegisterService(&sd, nil)
}