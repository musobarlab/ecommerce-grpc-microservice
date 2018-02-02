package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/query"
	queryMocks "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/query/mocks"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/repository"
	repoMocks "github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/repository/mocks"
)

func generateFindByIdSuccessMemberResult() <-chan repository.RepositoryResult {
	output := make(chan repository.RepositoryResult)
	go func() {
		exampleMember := model.NewMember()
		exampleMember.ID = "M1"
		exampleMember.FirstName = "Wuriyanto"
		exampleMember.LastName = "Musobar"
		exampleMember.Email = "wuriyanto48@yahoo.co.id"
		exampleMember.Password = "123456"
		exampleMember.PasswordSalt = "salt"
		exampleMember.BirthDate = time.Now()

		output <- repository.RepositoryResult{Result: exampleMember}
	}()
	return output
}

func generateFindByEmailSuccessMemberResult() <-chan query.QueryResult {
	output := make(chan query.QueryResult)
	go func() {
		exampleMember := model.NewMember()
		exampleMember.ID = "M1"
		exampleMember.FirstName = "Wuriyanto"
		exampleMember.LastName = "Musobar"
		exampleMember.Email = "wuriyanto48@yahoo.co.id"
		exampleMember.Password = "123456"
		exampleMember.PasswordSalt = "salt"
		exampleMember.BirthDate = time.Now()

		output <- query.QueryResult{Result: exampleMember}
	}()
	return output
}

func generateFindByEmailFailMemberResult() <-chan query.QueryResult {
	output := make(chan query.QueryResult)
	go func() {
		output <- query.QueryResult{Error: nil}
	}()
	return output
}

func generateSaveSuccessMemberResult() <-chan error {
	output := make(chan error)
	go func() {
		output <- nil
	}()
	return output
}

func generateSaveFailMemberResult() <-chan error {
	output := make(chan error)
	go func() {
		output <- errors.New("Error occured")
	}()
	return output
}

func TestMembershipUseCase(t *testing.T) {

	t.Run("Test Save", func(t *testing.T) {
		mockMemberRepo := new(repoMocks.MembershipRepository)
		mockMemberQuery := new(queryMocks.MembershipQuery)

		exampleMember := model.NewMember()
		exampleMember.ID = "M1"
		exampleMember.FirstName = "Wuriyanto"
		exampleMember.LastName = "Musobar"
		exampleMember.Email = "wuriyanto48@yahoo.co.id"
		exampleMember.Password = "123456"
		exampleMember.PasswordSalt = "salt"
		exampleMember.BirthDate = time.Now()

		mockMemberQuery.On("FindByEmail", mock.AnythingOfType("string")).Return(generateFindByEmailFailMemberResult())
		mockMemberRepo.On("Save", mock.AnythingOfType("*model.Member")).Return(generateSaveSuccessMemberResult())

		defer mockMemberQuery.AssertCalled(t, "FindByEmail", mock.AnythingOfType("string"))
		defer mockMemberRepo.AssertCalled(t, "Save", mock.AnythingOfType("*model.Member"))

		uc := NewMemberUseCase(mockMemberRepo, mockMemberQuery)

		err := <-uc.Save(exampleMember)

		assert.NoError(t, err)

	})

	t.Run("Find By Id", func(t *testing.T) {
		mockMemberRepo := new(repoMocks.MembershipRepository)
		mockMemberQuery := new(queryMocks.MembershipQuery)

		mockMemberRepo.On("Load", mock.AnythingOfType("string")).Return(generateFindByIdSuccessMemberResult())

		defer mockMemberRepo.AssertCalled(t, "Load", mock.AnythingOfType("string"))

		uc := NewMemberUseCase(mockMemberRepo, mockMemberQuery)

		result := <-uc.FindByID("M1")

		assert.NoError(t, result.Error)

	})

	t.Run("Find By Email", func(t *testing.T) {
		mockMemberRepo := new(repoMocks.MembershipRepository)
		mockMemberQuery := new(queryMocks.MembershipQuery)

		mockMemberQuery.On("FindByEmail", mock.AnythingOfType("string")).Return(generateFindByEmailSuccessMemberResult())

		defer mockMemberQuery.AssertCalled(t, "FindByEmail", mock.AnythingOfType("string"))

		uc := NewMemberUseCase(mockMemberRepo, mockMemberQuery)

		result := <-uc.FindByEmail("wuriyanto48@yahoo.co.id")

		assert.NoError(t, result.Error)

	})
}
