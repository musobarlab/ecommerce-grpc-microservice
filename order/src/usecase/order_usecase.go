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
	return nil
}

//FindProductAll
func (u *orderUseCaseImpl) FindProductAll() <-chan UseCaseResult {
	return nil
}

//FindCategoryByID
func (u *orderUseCaseImpl) FindCategoryByID(id int) <-chan UseCaseResult {
	return nil
}

//FindCategoryAll
func (u *orderUseCaseImpl) FindCategoryAll() <-chan UseCaseResult {
	return nil
}
