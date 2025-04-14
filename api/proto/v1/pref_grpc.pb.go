// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: v1/pref.proto

package routiq

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PrefService_ListPrefs_FullMethodName = "/asaycle.routiq.v1.PrefService/ListPrefs"
)

// PrefServiceClient is the client API for PrefService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrefServiceClient interface {
	ListPrefs(ctx context.Context, in *ListPrefsRequest, opts ...grpc.CallOption) (*ListPrefsResponse, error)
}

type prefServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrefServiceClient(cc grpc.ClientConnInterface) PrefServiceClient {
	return &prefServiceClient{cc}
}

func (c *prefServiceClient) ListPrefs(ctx context.Context, in *ListPrefsRequest, opts ...grpc.CallOption) (*ListPrefsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPrefsResponse)
	err := c.cc.Invoke(ctx, PrefService_ListPrefs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrefServiceServer is the server API for PrefService service.
// All implementations must embed UnimplementedPrefServiceServer
// for forward compatibility.
type PrefServiceServer interface {
	ListPrefs(context.Context, *ListPrefsRequest) (*ListPrefsResponse, error)
	mustEmbedUnimplementedPrefServiceServer()
}

// UnimplementedPrefServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPrefServiceServer struct{}

func (UnimplementedPrefServiceServer) ListPrefs(context.Context, *ListPrefsRequest) (*ListPrefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrefs not implemented")
}
func (UnimplementedPrefServiceServer) mustEmbedUnimplementedPrefServiceServer() {}
func (UnimplementedPrefServiceServer) testEmbeddedByValue()                     {}

// UnsafePrefServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrefServiceServer will
// result in compilation errors.
type UnsafePrefServiceServer interface {
	mustEmbedUnimplementedPrefServiceServer()
}

func RegisterPrefServiceServer(s grpc.ServiceRegistrar, srv PrefServiceServer) {
	// If the following call pancis, it indicates UnimplementedPrefServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PrefService_ServiceDesc, srv)
}

func _PrefService_ListPrefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPrefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefServiceServer).ListPrefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrefService_ListPrefs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefServiceServer).ListPrefs(ctx, req.(*ListPrefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrefService_ServiceDesc is the grpc.ServiceDesc for PrefService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrefService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "asaycle.routiq.v1.PrefService",
	HandlerType: (*PrefServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPrefs",
			Handler:    _PrefService_ListPrefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/pref.proto",
}
