package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	pb "github.com/wuriyanto48/ecommerce-grpc-microservice/product/protogo/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	c, err := clientProduct("localhost:3002")

	if err != nil {
		log.Fatalf("error 1 %v", err)
	}

	GetProduct(c)
}

func clientProduct(serverHost string) (pb.ProductServiceClient, error) {

	conn, err := grpc.Dial(serverHost, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	//defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	return client, nil
}

func GetProduct(client pb.ProductServiceClient) {
	md := metadata.Pairs("authorization", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	param := &pb.ProductQueryRequest{ID: 2}
	res, err := client.FindByID(ctx, param)

	if err != nil {
		log.Fatalf("error 1 %v", err)
	}

	fmt.Println(res)
}
