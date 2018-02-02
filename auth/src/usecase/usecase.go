package usecase

// UseCaseResult model
type UseCaseResult struct {
	Result interface{}
	Error  error
}

// AuthUseCase interface abstraction
type AuthUseCase interface {
	GetAccessToken(email, password string) <-chan UseCaseResult
}
