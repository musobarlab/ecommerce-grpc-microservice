package query

// QueryResult model
type QueryResult struct {
	Result interface{}
	Error  error
}

// ProductQuery interface abstraction
type ProductQuery interface {
	FindByID(id int) <-chan QueryResult
	FindAll() <-chan QueryResult
	FindByCategory(categoryID int) <-chan QueryResult
}

//CategoryQuery interface abstraction
type CategoryQuery interface {
	FindByID(id int) <-chan QueryResult
	FindAll() <-chan QueryResult
}
