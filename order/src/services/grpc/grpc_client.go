package grpc

//ServiceResult
type ServiceResult struct {
	Result interface{}
	Error  error
}

// MembershipGrpcClient interface
type MembershipGrpcClient interface {
	FindByID(id string) <-chan ServiceResult
	FindByEmail(email string) <-chan ServiceResult
}

//ProductGrpcClient interface
type ProductGrpcClient interface {
	FindByID(id int) <-chan ServiceResult
	FindByCategory(categoryID int) <-chan ServiceResult
	FindAll() <-chan ServiceResult
}

//CategoryGrpcClient interface
type CategoryGrpcClient interface {
	FindByID(id int) <-chan ServiceResult
	FindAll() <-chan ServiceResult
}
