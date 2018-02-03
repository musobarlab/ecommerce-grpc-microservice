package presenter

import (
	"errors"
	"strconv"

	pb "github.com/wuriyanto48/ecommerce-grpc-microservice/product/protogo/product"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/usecase"

	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GrpcProductHandler
type GrpcProductHandler struct {
	productUseCase usecase.ProductUseCase
}

// NewGrpcProductHandler
func NewGrpcProductHandler(productUseCase usecase.ProductUseCase) *GrpcProductHandler {
	return &GrpcProductHandler{productUseCase}
}

// FindByID
func (h *GrpcProductHandler) FindByID(ctx context.Context, arg *pb.ProductQueryRequest) (*pb.ProductResponse, error) {

	id := arg.ID

	productResult := <-h.productUseCase.FindByID(int(id))

	if productResult.Error != nil {
		return nil, status.Error(codes.InvalidArgument, productResult.Error.Error())
	}

	product, ok := productResult.Result.(*model.Product)

	if !ok {
		err := errors.New("Result is Not Product")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	stock, _ := strconv.Atoi(product.Stock.String())
	price, _ := strconv.ParseFloat(product.Price.String(), 32)
	productResponse := &pb.ProductResponse{
		ID:          int32(product.ID),
		CategoryID:  int32(product.CategoryID),
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Stock:       int32(stock),
		Price:       price,
	}

	return productResponse, nil
}

// FindByCategory
func (h *GrpcProductHandler) FindByCategory(arg *pb.ProductQueryRequest, stream pb.ProductService_FindByCategoryServer) error {

	categoryID := arg.CategoryID

	productResult := <-h.productUseCase.FindByCategory(int(categoryID))

	if productResult.Error != nil {
		return status.Error(codes.InvalidArgument, productResult.Error.Error())
	}

	products, ok := productResult.Result.(model.Products)

	if !ok {
		err := errors.New("Result is Not Products")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	for _, product := range products {

		stock, _ := strconv.Atoi(product.Stock.String())
		price, _ := strconv.ParseFloat(product.Price.String(), 32)
		productResponse := &pb.ProductResponse{
			ID:          int32(product.ID),
			CategoryID:  int32(product.CategoryID),
			Name:        product.Name,
			Description: product.Description,
			Image:       product.Image,
			Stock:       int32(stock),
			Price:       price,
		}

		if err := stream.Send(productResponse); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

//FindAll
func (h *GrpcProductHandler) FindAll(arg *pb.ProductQueryRequest, stream pb.ProductService_FindAllServer) error {

	productResult := <-h.productUseCase.FindAll()

	if productResult.Error != nil {
		return status.Error(codes.InvalidArgument, productResult.Error.Error())
	}

	products, ok := productResult.Result.(model.Products)

	if !ok {
		err := errors.New("Result is Not Products")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	for _, product := range products {

		stock, _ := strconv.Atoi(product.Stock.String())
		price, _ := strconv.ParseFloat(product.Price.String(), 32)
		productResponse := &pb.ProductResponse{
			ID:          int32(product.ID),
			CategoryID:  int32(product.CategoryID),
			Name:        product.Name,
			Description: product.Description,
			Image:       product.Image,
			Stock:       int32(stock),
			Price:       price,
		}

		if err := stream.Send(productResponse); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
