package usecase

import (
	"errors"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/model"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/query"
)

// categoryUseCaseImpl model
type categoryUseCaseImpl struct {
	categoryQuery query.CategoryQuery
}

// NewCategoryUseCase for initialise categoryUseCaseImpl model
func NewCategoryUseCase(categoryQuery query.CategoryQuery) CategoryUseCase {
	return &categoryUseCaseImpl{
		categoryQuery: categoryQuery,
	}
}

// FindByID
func (u *categoryUseCaseImpl) FindByID(id int) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		categoryResult := <-u.categoryQuery.FindByID(id)

		if categoryResult.Error != nil {
			output <- UseCaseResult{Error: categoryResult.Error}
			return
		}

		category, ok := categoryResult.Result.(model.Category)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not Category")}
			return
		}

		output <- UseCaseResult{Result: category}
	}()

	return output
}

// FindAll
func (u *categoryUseCaseImpl) FindAll() <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		categoryResult := <-u.categoryQuery.FindAll()

		if categoryResult.Error != nil {
			output <- UseCaseResult{Error: categoryResult.Error}
			return
		}

		categories, ok := categoryResult.Result.(model.Categories)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not Categories")}
			return
		}

		output <- UseCaseResult{Result: categories}
	}()

	return output
}
