package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	configEnv "github.com/joho/godotenv"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/config"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/middleware"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/presenter"
	grpcService "github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/services/grpc"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/usecase"
)

func main() {

	//load environtment variables
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Printf(".env is not loaded: %s", err.Error())
		os.Exit(2)
	}

	publicKey, err := config.InitPublicKey()

	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	// membership GRPC Host
	grpcMembershipHost, ok := os.LookupEnv("MEMBERSHIP_GRPC_HOST")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	// membership Auth Key
	grpcMembershipAuthKey, ok := os.LookupEnv("MEMBERSHIP_GRPC_AUTH_KEY")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	// product GRPC Host
	grpcProductHost, ok := os.LookupEnv("PRODUCT_GRPC_HOST")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	// product Auth Key
	grpcProductAuthKey, ok := os.LookupEnv("PRODUCT_GRPC_AUTH_KEY")
	if !ok {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	//init order handler

	//init grpc service
	membershipGrpcService, err := grpcService.NewMembershipGrpcClient(grpcMembershipHost, grpcMembershipAuthKey)
	if err != nil {
		fmt.Printf("Error %s", err.Error())
		os.Exit(1)
	}

	productGrpcService, err := grpcService.NewProductGrpcClient(grpcProductHost, grpcProductAuthKey)
	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	categoryGrpcService, err := grpcService.NewCategoryGrpcClient(grpcProductHost, grpcProductAuthKey)
	if err != nil {
		fmt.Printf("Error %e", err.Error())
		os.Exit(1)
	}

	orderUseCase := usecase.NewOrderUseCase(membershipGrpcService, productGrpcService, categoryGrpcService)

	orderHttpHandler := presenter.NewHttpOrderHandler(orderUseCase)

	//routing

	r := mux.NewRouter()

	r.Handle("/api/me", middleware.LogRequest(middleware.Bearer(publicKey, orderHttpHandler.Me()))).Methods("GET")

	r.Handle("/api/products", middleware.LogRequest(orderHttpHandler.GetProducts())).Methods("GET")
	r.Handle("/api/products/{id}", middleware.LogRequest(orderHttpHandler.GetProduct())).Methods("GET")

	r.Handle("/api/categories", middleware.LogRequest(orderHttpHandler.GetCategories())).Methods("GET")
	r.Handle("/api/categories/{id}", middleware.LogRequest(orderHttpHandler.GetCategory())).Methods("GET")

	http.ListenAndServe(":3004", r)
}
