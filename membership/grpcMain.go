package main

import (
	"fmt"
	"os"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/config"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/presenter"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/query"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/repository"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/usecase"

	membershipGrpc "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/grpc/servers"
)

//GrpcPortDefault default port for GRPC Server
const GrpcPortDefault = 3001

func grpcMain() {

	//init member handler
	memberQuery := query.NewMemberQueryInMemory(config.GetInMemoryDb())
	memberRepository := repository.NewMemberRepositoryInMemory(config.GetInMemoryDb())

	memberUseCase := usecase.NewMemberUseCase(memberRepository, memberQuery)

	memberGrpcHandler := presenter.NewGrpcHandler(memberUseCase)
	//end init member handler

	grpcServer, err := membershipGrpc.NewGrpcServer(memberGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}

	err = grpcServer.Serve(GrpcPortDefault)

	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}
}
