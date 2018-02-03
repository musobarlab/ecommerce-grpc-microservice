package main

import (
	"net/http"
	"os"
	"time"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/config"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/db"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/middleware"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/presenter"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/query"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/token"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/usecase"
)

func main() {

	//
	privateKey, err := config.InitPrivateKey()

	if err != nil {
		os.Exit(1)
	}

	// init auth handler

	identityQuery := query.NewIdentityQueryInMemory(db.GetInMemoryDb())
	accessTokenGenerator := token.NewJwtGenerator(privateKey, time.Minute*10)

	authUseCase := usecase.NewAuthUseCase(identityQuery, accessTokenGenerator)

	authHttpHandler := presenter.NewHttpHandler(authUseCase)
	// end init auth handler

	http.Handle("/api/auth", middleware.LogRequest(authHttpHandler.Auth()))
	http.ListenAndServe(":3000", nil)
}
