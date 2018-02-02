package repository

import (
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
)

// RepositoryResult model
type RepositoryResult struct {
	Result interface{}
	Error  error
}

// MembershipRepository interface abstraction
type MembershipRepository interface {
	Save(m *model.Member) <-chan error
	Load(id string) <-chan RepositoryResult
}
