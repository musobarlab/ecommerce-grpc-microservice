package usecase

import (
	"errors"

	grpcService "github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/services/grpc"
	serviceModel "github.com/wuriyanto48/ecommerce-grpc-microservice/order/src/services/model"
)

//orderUseCaseImpl
type orderUseCaseImpl struct {
	membershipGrpcService grpcService.MembershipGrpcClient
	productGrpcService    grpcService.ProductGrpcClient
	categoryGrpcService   grpcService.CategoryGrpcClient
}

//NewOrderUseCase
func NewOrderUseCase(membershipGrpcService grpcService.MembershipGrpcClient,
	productGrpcService grpcService.ProductGrpcClient, categoryGrpcService grpcService.CategoryGrpcClient) OrderUseCase {
	return &orderUseCaseImpl{
		membershipGrpcService: membershipGrpcService,
		productGrpcService:    productGrpcService,
		categoryGrpcService:   categoryGrpcService,
	}
}

//FindMemberByID
func (u *orderUseCaseImpl) FindMemberByID(id string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		memberResult := <-u.membershipGrpcService.FindByID(id)

		if memberResult.Error != nil {
			output <- UseCaseResult{Error: memberResult.Error}
			return
		}

		member, ok := memberResult.Result.(serviceModel.Member)

		if !ok {
			err := errors.New("Result is not member")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: member}
	}()
	return output
}

//FindProductByID
func (u *orderUseCaseImpl) FindProductByID(id int) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		productResult := <-u.productGrpcService.FindByID(id)

		if productResult.Error != nil {
			output <- UseCaseResult{Error: productResult.Error}
			return
		}

		product, ok := productResult.Result.(serviceModel.Product)

		if !ok {
			err := errors.New("Result is not Product")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: product}
	}()
	return output
}

//FindProductAll
func (u *orderUseCaseImpl) FindProductAll() <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		productResult := <-u.productGrpcService.FindAll()

		if productResult.Error != nil {
			output <- UseCaseResult{Error: productResult.Error}
			return
		}

		products, ok := productResult.Result.(serviceModel.Products)

		if !ok {
			err := errors.New("Result is not Products")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: products}
	}()
	return output
}

//FindCategoryByID
func (u *orderUseCaseImpl) FindCategoryByID(id int) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		categoryResult := <-u.categoryGrpcService.FindByID(id)

		if categoryResult.Error != nil {
			output <- UseCaseResult{Error: categoryResult.Error}
			return
		}

		category, ok := categoryResult.Result.(serviceModel.Category)

		if !ok {
			err := errors.New("Result is not Category")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: category}
	}()
	return output
}

//FindCategoryAll
func (u *orderUseCaseImpl) FindCategoryAll() <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		categoryResult := <-u.categoryGrpcService.FindAll()

		if categoryResult.Error != nil {
			output <- UseCaseResult{Error: categoryResult.Error}
			return
		}

		categories, ok := categoryResult.Result.(serviceModel.Categories)

		if !ok {
			err := errors.New("Result is not Categories")
			output <- UseCaseResult{Error: err}
			return
		}

		output <- UseCaseResult{Result: categories}
	}()
	return output
}
