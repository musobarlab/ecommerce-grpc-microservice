package query

// QueryResult model
type QueryResult struct {
	Result interface{}
	Error  error
}

// IdentityQuery interface abstraction
type IdentityQuery interface {
	FindByEmail(email string) <-chan QueryResult
}
