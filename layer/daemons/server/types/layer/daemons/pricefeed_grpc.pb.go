// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

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

// PriceFeedServiceClient is the client API for PriceFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceFeedServiceClient interface {
	// Updates market prices.
	UpdateMarketPrices(ctx context.Context, in *UpdateMarketPricesRequest, opts ...grpc.CallOption) (*UpdateMarketPricesResponse, error)
}

type priceFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceFeedServiceClient(cc grpc.ClientConnInterface) PriceFeedServiceClient {
	return &priceFeedServiceClient{cc}
}

func (c *priceFeedServiceClient) UpdateMarketPrices(ctx context.Context, in *UpdateMarketPricesRequest, opts ...grpc.CallOption) (*UpdateMarketPricesResponse, error) {
	out := new(UpdateMarketPricesResponse)
	err := c.cc.Invoke(ctx, "/layer.daemons.PriceFeedService/UpdateMarketPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceFeedServiceServer is the server API for PriceFeedService service.
// All implementations must embed UnimplementedPriceFeedServiceServer
// for forward compatibility
type PriceFeedServiceServer interface {
	// Updates market prices.
	UpdateMarketPrices(context.Context, *UpdateMarketPricesRequest) (*UpdateMarketPricesResponse, error)
	mustEmbedUnimplementedPriceFeedServiceServer()
}

// UnimplementedPriceFeedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPriceFeedServiceServer struct {
}

func (UnimplementedPriceFeedServiceServer) UpdateMarketPrices(context.Context, *UpdateMarketPricesRequest) (*UpdateMarketPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarketPrices not implemented")
}
func (UnimplementedPriceFeedServiceServer) mustEmbedUnimplementedPriceFeedServiceServer() {}

// UnsafePriceFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceFeedServiceServer will
// result in compilation errors.
type UnsafePriceFeedServiceServer interface {
	mustEmbedUnimplementedPriceFeedServiceServer()
}

func RegisterPriceFeedServiceServer(s grpc.ServiceRegistrar, srv PriceFeedServiceServer) {
	s.RegisterService(&PriceFeedService_ServiceDesc, srv)
}

func _PriceFeedService_UpdateMarketPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMarketPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceFeedServiceServer).UpdateMarketPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layer.daemons.PriceFeedService/UpdateMarketPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceFeedServiceServer).UpdateMarketPrices(ctx, req.(*UpdateMarketPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceFeedService_ServiceDesc is the grpc.ServiceDesc for PriceFeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceFeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "layer.daemons.PriceFeedService",
	HandlerType: (*PriceFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateMarketPrices",
			Handler:    _PriceFeedService_UpdateMarketPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "layer/daemons/pricefeed.proto",
}
