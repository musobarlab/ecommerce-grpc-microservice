package main

import (
	"net/http"
	"os"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/config"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/middleware"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/presenter"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/query"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/repository"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/usecase"
)

func main() {

	publicKey, err := config.InitPublicKey()

	if err != nil {
		os.Exit(1)
	}

	//init member handler
	memberQuery := query.NewMemberQueryInMemory(config.GetInMemoryDb())
	memberRepository := repository.NewMemberRepositoryInMemory(config.GetInMemoryDb())

	memberUseCase := usecase.NewMemberUseCase(memberRepository, memberQuery)

	memberHttpHandler := presenter.NewHttpMemberHandler(memberUseCase)
	//end init member handler

	http.Handle("/api/me", middleware.LogRequest(middleware.Bearer(publicKey, memberHttpHandler.Me())))
	http.ListenAndServe(":3001", nil)
}
