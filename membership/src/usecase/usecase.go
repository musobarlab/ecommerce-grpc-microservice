package usecase

import (
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
)

// UseCaseResult model
type UseCaseResult struct {
	Result interface{}
	Error  error
}

// MemberUseCase interface abstraction
type MemberUseCase interface {
	Save(m *model.Member) <-chan error
	FindByID(id string) <-chan UseCaseResult
	FindByEmail(email string) <-chan UseCaseResult
}
