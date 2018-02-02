package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuriyanto48/ecommerce-grpc-microservice/membership/src/model"
)

func TestMemberRepositoryInMemory(t *testing.T) {

	db := make(map[string]*model.Member)

	r := NewMemberRepositoryInMemory(db)

	t.Run("Save Member", func(t *testing.T) {
		wury := model.NewMember()
		wury.ID = "M1"
		wury.FirstName = "Wuriyanto"
		wury.LastName = "Musobar"

		err := <-r.Save(wury)
		assert.NoError(t, err)

	})

	t.Run("Load Member", func(t *testing.T) {

		memberResult := <-r.Load("M1")

		assert.NoError(t, memberResult.Error)

		wury, ok := memberResult.Result.(*model.Member)

		assert.True(t, ok)
		assert.Equal(t, "Wuriyanto", wury.FirstName)
		assert.Equal(t, "Musobar", wury.LastName)
		assert.Equal(t, 1, wury.Version)

	})

	t.Run("Update Member", func(t *testing.T) {

		memberResult := <-r.Load("M1")

		assert.NoError(t, memberResult.Error)

		wury, ok := memberResult.Result.(*model.Member)
		assert.True(t, ok)

		wury.LastName = "Lone Wolf"

		err := <-r.Save(wury)
		assert.NoError(t, err)

		assert.Equal(t, "Wuriyanto", wury.FirstName)
		assert.Equal(t, "Lone Wolf", wury.LastName)
		assert.Equal(t, 2, wury.Version)

	})

	t.Run("Load Member Not Found", func(t *testing.T) {

		memberResult := <-r.Load("M2")

		assert.Error(t, memberResult.Error)

		_, ok := memberResult.Result.(*model.Member)
		assert.False(t, ok)

	})

}
