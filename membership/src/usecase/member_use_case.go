package usecase

import (
	"errors"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/query"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/repository"
)

// memberUseCaseImpl model
type memberUseCaseImpl struct {
	memberRepository repository.MembershipRepository
	memberQuery      query.MembershipQuery
}

// NewMemberUseCase for initialise memberUseCaseImpl model
func NewMemberUseCase(memberRepository repository.MembershipRepository, memberQuery query.MembershipQuery) MemberUseCase {
	return &memberUseCaseImpl{
		memberRepository: memberRepository,
		memberQuery:      memberQuery,
	}
}

// Save for saving Member model
func (mu *memberUseCaseImpl) Save(m *model.Member) <-chan error {
	output := make(chan error)

	go func() {

		defer close(output)

		memberExistResult := <-mu.memberQuery.FindByEmail(m.Email)

		if memberExistResult.Error != nil && memberExistResult.Error.Error() != "MEMBER_NOT_FOUND" {
			output <- memberExistResult.Error
			return
		} else if memberExistResult.Error != nil {
			output <- memberExistResult.Error
			return
		} else {
			err := <-mu.memberRepository.Save(m)

			if err != nil {
				output <- err
				return
			}
		}

		output <- nil

	}()

	return output
}

// FindByID for load Member by its ID
func (mu *memberUseCaseImpl) FindByID(id string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		memberResult := <-mu.memberRepository.Load(id)

		if memberResult.Error != nil {
			output <- UseCaseResult{Error: memberResult.Error}
			return
		}

		member, ok := memberResult.Result.(*model.Member)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not member")}
			return
		}

		output <- UseCaseResult{Result: member}

	}()

	return output
}

// FindByEmail for load Member by its Email
func (mu *memberUseCaseImpl) FindByEmail(email string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		memberResult := <-mu.memberQuery.FindByEmail(email)

		if memberResult.Error != nil {
			output <- UseCaseResult{Error: memberResult.Error}
			return
		}

		member, ok := memberResult.Result.(*model.Member)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not member")}
			return
		}

		output <- UseCaseResult{Result: member}

	}()

	return output
}
