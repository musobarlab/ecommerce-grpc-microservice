package repository

import (
	"errors"
	"time"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/model"
)

// memberRepositoryInMemory model
type memberRepositoryInMemory struct {
	db map[string]*model.Member
}

// NewMemberRepositoryInMemory for initialise memberRepositoryInMemory model
func NewMemberRepositoryInMemory(db map[string]*model.Member) MembershipRepository {
	return &memberRepositoryInMemory{db}
}

// Save function for saving Member model
func (r *memberRepositoryInMemory) Save(m *model.Member) <-chan error {
	output := make(chan error)
	go func() {
		member, ok := r.db[m.ID]
		if !ok {
			m.Version++
			r.db[m.ID] = m
			output <- nil
			return
		} else {
			member.Version++
			member.UpdatedAt = time.Now()
			r.db[m.ID] = member
			output <- nil
			return
		}
	}()
	return output
}

// Load for load Member by its ID
func (r *memberRepositoryInMemory) Load(id string) <-chan RepositoryResult {
	output := make(chan RepositoryResult)
	go func() {
		member, ok := r.db[id]
		if !ok {
			output <- RepositoryResult{Error: errors.New("member not found")}
			return
		}

		output <- RepositoryResult{Result: member}
	}()
	return output
}
