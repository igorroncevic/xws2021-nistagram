package grpc_common

import (
	"context"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	otgo "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func CreateGrpcConnection(address string) (*grpc.ClientConn, error) {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	return grpc.DialContext(
		context.Background(),
		address,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
		grpc.WithStreamInterceptor(
			grpc_opentracing.StreamClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
	)
}