package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type recvWrapper struct {
	ctx context.Context
	grpc.ServerStream
}

func (r *recvWrapper) Context() context.Context {
	return r.ctx
}

func TestAuthInterceptor(t *testing.T) {

	grpcMiddleware := NewInterceptor("123456")

	t.Run("Test Auth Positive Test", func(t *testing.T) {
		md := metadata.Pairs("authorization", "123456")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		unaryInfo := &grpc.UnaryServerInfo{
			FullMethod: "TestService.UnaryMethod",
		}

		unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "fake-response", nil
		}

		_, err := grpcMiddleware.Auth(ctx, "Interceptor", unaryInfo, unaryHandler)

		assert.NoError(t, err)
	})

	t.Run("Test Auth Stream Positive Test", func(t *testing.T) {
		md := metadata.Pairs("authorization", "123456")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		stream := &recvWrapper{
			ctx: ctx,
		}

		streamInfo := &grpc.StreamServerInfo{
			FullMethod: "TestService.StreamUnaryMethod",
		}

		streamHandler := func(srv interface{}, req grpc.ServerStream) error {
			return nil
		}

		err := grpcMiddleware.AuthStream("test", stream, streamInfo, streamHandler)

		assert.NoError(t, err)
	})

	t.Run("Test Auth Stream Missing Authorization", func(t *testing.T) {
		md := metadata.Pairs("authorization", "")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		stream := &recvWrapper{
			ctx: ctx,
		}

		streamInfo := &grpc.StreamServerInfo{
			FullMethod: "TestService.StreamUnaryMethod",
		}

		streamHandler := func(srv interface{}, req grpc.ServerStream) error {
			return nil
		}

		err := grpcMiddleware.AuthStream("test", stream, streamInfo, streamHandler)

		assert.Error(t, err)
	})

	t.Run("Test Auth Stream Negative Test", func(t *testing.T) {
		md := metadata.Pairs("authorization", "12345")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		stream := &recvWrapper{
			ctx: ctx,
		}

		streamInfo := &grpc.StreamServerInfo{
			FullMethod: "TestService.StreamUnaryMethod",
		}

		streamHandler := func(srv interface{}, req grpc.ServerStream) error {
			return nil
		}

		err := grpcMiddleware.AuthStream("test", stream, streamInfo, streamHandler)

		assert.Error(t, err)
	})

	t.Run("Test Auth Negative Test", func(t *testing.T) {
		md := metadata.Pairs("authorization", "invalid-auth")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		unaryInfo := &grpc.UnaryServerInfo{
			FullMethod: "TestService.UnaryMethod",
		}

		unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "fake-response", nil
		}

		_, err := grpcMiddleware.Auth(ctx, "Interceptor", unaryInfo, unaryHandler)

		assert.Error(t, err)
	})

	t.Run("Test Auth Missing Authorization", func(t *testing.T) {
		md := metadata.Pairs("authorization", "")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		unaryInfo := &grpc.UnaryServerInfo{
			FullMethod: "TestService.UnaryMethod",
		}

		unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return "fake-response", nil
		}

		_, err := grpcMiddleware.Auth(ctx, "Interceptor", unaryInfo, unaryHandler)

		assert.Error(t, err)
	})
}
