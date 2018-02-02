package query

import (
	"errors"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/product/src/model"
)

// productQueryInMemory model
type productQueryInMemory struct {
	db map[int]*model.Product
}

// NewProductQueryInMemory for initialise productQueryInMemory model
func NewProductQueryInMemory(db map[int]*model.Product) ProductQuery {
	return &productQueryInMemory{db}
}

// FindByID will return Product by its id
func (q *productQueryInMemory) FindByID(id int) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		product, ok := q.db[id]
		if !ok {
			output <- QueryResult{Error: errors.New("product not found")}
			return
		}

		output <- QueryResult{Result: product}
	}()
	return output
}

// FindAll will return all products
func (q *productQueryInMemory) FindAll() <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		var products model.Products
		for _, v := range q.db {
			products = append(products, *v)
		}

		output <- QueryResult{Result: products}
	}()
	return output
}

// FindByCategory will return all product by categoryID
func (q *productQueryInMemory) FindByCategory(categoryID int) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		var products model.Products
		for _, v := range q.db {
			if v.CategoryID == categoryID {
				products = append(products, *v)
			}
		}

		output <- QueryResult{Result: products}
	}()
	return output
}
