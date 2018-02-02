package usecase

// UseCaseResult model
type UseCaseResult struct {
	Result interface{}
	Error  error
}

// ProductUseCase interface abstraction
type ProductUseCase interface {
	FindByID(id int) <-chan UseCaseResult
	FindAll() <-chan UseCaseResult
	FindByCategory(categoryID int) <-chan UseCaseResult
}

type CategoryUseCase interface {
	FindByID(id int) <-chan UseCaseResult
	FindAll() <-chan UseCaseResult
}
