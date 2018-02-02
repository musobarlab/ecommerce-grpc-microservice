package servers

import (
	"errors"
	"fmt"
	"net"
	"os"

	middleware "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/grpc/middleware"

	pbMembership "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/protogo/membership"

	membershipPresenterPackage "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/presenter"

	"google.golang.org/grpc"
)

//Server data structure, grpc server model
type Server struct {
	membershipGrpcHandler *membershipPresenterPackage.GrpcHandler
	grpcMiddleware        *middleware.Interceptor
}

// Serve insecure server/ no server side encryption
func (s *Server) Serve(port uint) error {
	address := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer(
		//Unary interceptor
		grpc.UnaryInterceptor(s.grpcMiddleware.Auth),
		//Stream interceptor
		grpc.StreamInterceptor(s.grpcMiddleware.AuthStream),
	)

	//Register all sub server here
	pbMembership.RegisterMembershipServiceServer(server, s.membershipGrpcHandler)
	//end register server

	err = server.Serve(l)

	if err != nil {
		return err
	}

	fmt.Sprintf("Product GRPC Server running on PORT %d", port)

	return nil
}

//NewGrpcServer function, return: Pointer GRPC Server, or error otherwise
func NewGrpcServer(membershipGrpcHandler *membershipPresenterPackage.GrpcHandler) (*Server, error) {
	//init Auth Key

	grpcAuthKey, ok := os.LookupEnv("GRPC_AUTH_KEY")
	if !ok {
		err := errors.New("you need to specify GRPC_AUTH_KEY in the environment variable")
		return nil, err
	}

	grpcMiddleware := middleware.NewInterceptor(grpcAuthKey)

	return &Server{
		membershipGrpcHandler: membershipGrpcHandler,
		grpcMiddleware:        grpcMiddleware,
	}, nil

}
