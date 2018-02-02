package config

import (
	"time"

	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
)

// GetInMemoryDb return *model.Member map, this fake database just for testing purposes only
func GetInMemoryDb() map[string]*model.Member {
	db := make(map[string]*model.Member)

	exampleMember := model.NewMember()
	exampleMember.ID = "M1"
	exampleMember.FirstName = "Wuriyanto"
	exampleMember.LastName = "Musobar"
	exampleMember.Email = "wuriyanto48@yahoo.co.id"
	exampleMember.Password = "123456"
	exampleMember.PasswordSalt = "salt"
	exampleMember.BirthDate = time.Now()

	db[exampleMember.ID] = exampleMember

	return db
}
