package query

import (
	"errors"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/model"
)

// identityQueryInMemory model, this private model implement IdentityQuery
type identityQueryInMemory struct {
	db map[string]*model.Identity
}

// NewIdentityQueryInMemory function for initialise identityQueryInMemory model
func NewIdentityQueryInMemory(db map[string]*model.Identity) IdentityQuery {
	return &identityQueryInMemory{db}
}

// FindByEmail function return Identity by its Email
func (q *identityQueryInMemory) FindByEmail(email string) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		var identity *model.Identity
		for _, v := range q.db {
			if v.Email == email {
				identity = v
				break
			} else {
				output <- QueryResult{Error: errors.New("IDENTITY_NOT_FOUND")}
				return
			}
		}

		output <- QueryResult{Result: identity}
	}()
	return output
}
