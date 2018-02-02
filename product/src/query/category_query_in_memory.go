package query

import (
	"errors"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/model"
)

// categoryQueryInMemory model
type categoryQueryInMemory struct {
	db map[int]*model.Category
}

// NewCategoryQueryInMemory for initialise categoryQueryInMemory model
func NewCategoryQueryInMemory(db map[int]*model.Category) CategoryQuery {
	return &categoryQueryInMemory{db}
}

// FindByID will return Category by its id
func (q *categoryQueryInMemory) FindByID(id int) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		category, ok := q.db[id]
		if !ok {
			output <- QueryResult{Error: errors.New("category not found")}
			return
		}

		output <- QueryResult{Result: category}
	}()
	return output
}

// FindAll will return all categories
func (q *categoryQueryInMemory) FindAll() <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		var categories model.Categories
		for _, v := range q.db {
			categories = append(categories, *v)
		}

		output <- QueryResult{Result: categories}
	}()
	return output
}
