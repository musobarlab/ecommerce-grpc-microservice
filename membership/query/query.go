package query

// QueryResult model
type QueryResult struct {
	Result interface{}
	Error  error
}

// MembershipQuery interface abstraction
type MembershipQuery interface {
	FindByEmail(email string) <-chan QueryResult
}
