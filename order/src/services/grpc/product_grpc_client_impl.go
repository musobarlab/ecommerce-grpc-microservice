package grpc

import (
	"io"
	"strconv"
	"time"

	"golang.org/x/net/context"

	pb "github.com/wuriyanto48/ecommerce-grpc-microservice/order/protogo/product"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/services/model"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//productGrpcClientImpl struct
type productGrpcClientImpl struct {
	grpcAuthKey string
	client      pb.ProductServiceClient
}

//NewProductGrpcClient
func NewProductGrpcClient(host, grpcAuthKey string) (*productGrpcClientImpl, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewProductServiceClient(conn)

	return &productGrpcClientImpl{
		grpcAuthKey: grpcAuthKey,
		client:      client,
	}, nil
}

//FindByID function, for find Product By Id using GRPC service client
func (c *productGrpcClientImpl) FindByID(id int) <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		defer close(output)

		md := metadata.Pairs("authorization", c.grpcAuthKey)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		arg := &pb.ProductQueryRequest{ID: int32(id)}
		res, err := c.client.FindByID(ctx, arg)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		//stock from int32 to decimal
		stock, err := decimal.NewFromString(strconv.Itoa(int(res.Stock)))

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		//price from float64 to decimal
		price := decimal.NewFromFloat(res.Price)

		product := model.Product{
			ID:          int(res.ID),
			CategoryID:  int(res.CategoryID),
			Name:        res.Name,
			Description: res.Description,
			Image:       res.Image,
			Stock:       stock,
			Price:       price,
			Version:     1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		output <- ServiceResult{Result: product}
	}()

	return output
}

//FindByCategory function, for find Product By Category Id using GRPC service client
func (c *productGrpcClientImpl) FindByCategory(categoryID int) <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		md := metadata.Pairs("authorization", c.grpcAuthKey)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		arg := &pb.ProductQueryRequest{CategoryID: int32(categoryID)}
		resStream, err := c.client.FindByCategory(ctx, arg)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		var products model.Products

		for {
			res, err := resStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				output <- ServiceResult{Error: err}
				return
			}

			//stock from int32 to decimal
			stock, err := decimal.NewFromString(strconv.Itoa(int(res.Stock)))

			if err != nil {
				output <- ServiceResult{Error: err}
				return
			}

			//price from float64 to decimal
			price := decimal.NewFromFloat(res.Price)

			product := model.Product{
				ID:          int(res.ID),
				CategoryID:  int(res.CategoryID),
				Name:        res.Name,
				Description: res.Description,
				Image:       res.Image,
				Stock:       stock,
				Price:       price,
				Version:     1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			products = append(products, product)
		}

		output <- ServiceResult{Result: products}

	}()
	return output
}

//FindAll function, for find all Product Using GRPC service client
func (c *productGrpcClientImpl) FindAll() <-chan ServiceResult {
	output := make(chan ServiceResult)

	go func() {
		md := metadata.Pairs("authorization", c.grpcAuthKey)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		arg := &pb.ProductQueryRequest{}
		resStream, err := c.client.FindAll(ctx, arg)

		if err != nil {
			output <- ServiceResult{Error: err}
			return
		}

		var products model.Products

		for {
			res, err := resStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				output <- ServiceResult{Error: err}
				return
			}

			//stock from int32 to decimal
			stock, err := decimal.NewFromString(strconv.Itoa(int(res.Stock)))

			if err != nil {
				output <- ServiceResult{Error: err}
				return
			}

			//price from float64 to decimal
			price := decimal.NewFromFloat(res.Price)

			product := model.Product{
				ID:          int(res.ID),
				CategoryID:  int(res.CategoryID),
				Name:        res.Name,
				Description: res.Description,
				Image:       res.Image,
				Stock:       stock,
				Price:       price,
				Version:     1,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			products = append(products, product)
		}

		output <- ServiceResult{Result: products}

	}()
	return output
}
