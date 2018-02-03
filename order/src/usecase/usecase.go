package usecase

// UseCaseResult model
type UseCaseResult struct {
	Result interface{}
	Error  error
}

// OrderUseCase interface abstraction
type OrderUseCase interface {
	FindMemberByID(id string) <-chan UseCaseResult
	FindProductByID(id int) <-chan UseCaseResult
	FindProductAll() <-chan UseCaseResult
	FindCategoryByID(id int) <-chan UseCaseResult
	FindCategoryAll() <-chan UseCaseResult
}
