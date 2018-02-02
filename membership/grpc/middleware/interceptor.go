package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Interceptor data structure
type Interceptor struct {
	authKey string
}

// NewInterceptor function, for init Interceptor object
func NewInterceptor(authKey string) *Interceptor {
	return &Interceptor{authKey}
}

//Auth function,
//or Unary interceptor
//additional security for our GRPC server
func (i *Interceptor) Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(meta["authorization"]) != 1 {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	authorization := meta["authorization"][0]

	if authorization != i.authKey {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	return handler(ctx, req)
}

// AuthStream
func (i *Interceptor) AuthStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	meta, ok := metadata.FromIncomingContext(stream.Context())

	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(meta["authorization"]) != 1 {
		return grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	authorization := meta["authorization"][0]

	if authorization != i.authKey {
		return grpc.Errorf(codes.Unauthenticated, "invalid authorization")
	}

	return handler(srv, stream)
}
