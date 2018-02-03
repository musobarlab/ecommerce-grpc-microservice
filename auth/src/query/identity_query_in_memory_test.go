package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/auth/src/model"
)

func TestIdentityQueryInMemory(t *testing.T) {

	db := make(map[string]*model.Identity)

	var exampleIdentity model.Identity
	exampleIdentity.ID = "M1"
	exampleIdentity.Email = "wuriyanto48@yahoo.co.id"
	exampleIdentity.Password = "123456"
	exampleIdentity.PasswordSalt = "salt"

	db["wuriyanto48@yahoo.co.id"] = &exampleIdentity

	q := NewIdentityQueryInMemory(db)

	t.Run("Find Identity By Email", func(t *testing.T) {

		identityResult := <-q.FindByEmail("wuriyanto48@yahoo.co.id")

		assert.NoError(t, identityResult.Error)

		wury, ok := identityResult.Result.(*model.Identity)

		assert.True(t, ok)
		assert.Equal(t, "wuriyanto48@yahoo.co.id", wury.Email)

	})

	t.Run("Find Identity By Email Not Found", func(t *testing.T) {

		identityResult := <-q.FindByEmail("invalid@email.com")

		assert.Error(t, identityResult.Error)

		_, ok := identityResult.Result.(*model.Identity)
		assert.False(t, ok)

	})

}
