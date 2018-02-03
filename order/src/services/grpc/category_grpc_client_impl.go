package grpc

import (
	"io"

	"golang.org/x/net/context"

	pb "github.com/wuriyanto48/ecommerce-grpc-microservice/order/protogo/category"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/services/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//categoryGrpcClientImpl struct
type categoryGrpcClientImpl struct {
	grpcAuthKey string
	client      pb.CategoryServiceClient
}

//NewCategoryGrpcClient
func NewCategoryGrpcClient(host, grpcAuthKey string) (*categoryGrpcClientImpl, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewCategoryServiceClient(conn)

	return &categoryGrpcClientImpl{
		grpcAuthKey: grpcAuthKey,
		client:      client,
	}, nil
}

//FindByID function, for find Category By Id using GRPC service client
func (c *categoryGrpcClientImpl) FindByID(id int) <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		defer close(output)

		md := metadata.Pairs("authorization", c.grpcAuthKey)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		arg := &pb.CategoryQueryRequest{ID: int32(id)}
		res, err := c.client.FindByID(ctx, arg)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		category := model.Category{
			ID:          int(res.ID),
			Name:        res.Name,
			Description: res.Description,
			Image:       res.Image,
		}

		output <- ServiceResult{Result: category}
	}()

	return output
}

//FindAll function, for find all Category Using GRPC service client
func (c *categoryGrpcClientImpl) FindAll() <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		defer close(output)

		md := metadata.Pairs("authorization", c.grpcAuthKey)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		arg := &pb.CategoryQueryRequest{}
		resStream, err := c.client.FindAll(ctx, arg)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		var categories model.Categories

		for {
			res, err := resStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				output <- ServiceResult{Error: err}
				return
			}

			category := model.Category{
				ID:          int(res.ID),
				Name:        res.Name,
				Description: res.Description,
				Image:       res.Image,
			}

			categories = append(categories, category)
		}

		output <- ServiceResult{Result: categories}

	}()
	return output
}
