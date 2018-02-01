package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/config"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/middleware"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/presenter"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/query"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/token"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/usecase"
)

func main() {

	//
	privateKey, err := config.InitPrivateKey()

	if err != nil {
		fmt.Println(err)
	}

	// init auth handler
	db := make(map[string]*model.Identity)
	var i model.Identity
	i.ID = "M1"
	i.Email = "wuriyanto48@yahoo.co.id"
	i.Password = "12345"

	db["M1"] = &i

	identityQuery := query.NewIdentityQueryInMemory(db)
	accessTokenGenerator := token.NewJwtGenerator(privateKey, time.Minute*10)

	authUseCase := usecase.NewAuthUseCase(identityQuery, accessTokenGenerator)

	authHttpHandler := presenter.NewHttpHandler(authUseCase)
	// end init auth handler

	http.Handle("/api/auth", middleware.LogRequest(authHttpHandler.Auth()))
	http.ListenAndServe(":3000", nil)
}
