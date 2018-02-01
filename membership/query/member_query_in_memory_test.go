package query

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/model"
)

func TestMemberQueryInMemory(t *testing.T) {

	db := make(map[string]*model.Member)

	exampleMember := model.NewMember()
	exampleMember.ID = "M1"
	exampleMember.FirstName = "Wuriyanto"
	exampleMember.LastName = "Musobar"
	exampleMember.Email = "wuriyanto48@yahoo.co.id"
	exampleMember.Password = "123456"
	exampleMember.PasswordSalt = "salt"
	exampleMember.BirthDate = time.Now()

	db["M1"] = exampleMember

	q := NewMemberQueryInMemory(db)

	t.Run("Find Member By Email", func(t *testing.T) {

		memberResult := <-q.FindByEmail("wuriyanto48@yahoo.co.id")

		assert.NoError(t, memberResult.Error)

		wury, ok := memberResult.Result.(*model.Member)

		assert.True(t, ok)
		assert.Equal(t, "Wuriyanto", wury.FirstName)
		assert.Equal(t, "Musobar", wury.LastName)
		assert.Equal(t, 0, wury.Version)

	})

	t.Run("Find Member By Email Not Found", func(t *testing.T) {

		memberResult := <-q.FindByEmail("invalid@email.com")

		assert.Error(t, memberResult.Error)

		_, ok := memberResult.Result.(*model.Member)
		assert.False(t, ok)

	})

}
