package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	t.Run("Test Valid Password", func(t *testing.T) {
		var i Identity
		i.ID = "M1"
		i.Email = "wuriyanto48@yahoo.co.id"
		i.Password = "12345"

		err := i.IsValidPassword("12345")
		assert.NoError(t, err)
	})

	t.Run("Test Invalid Password", func(t *testing.T) {
		var i Identity
		i.ID = "M1"
		i.Email = "wuriyanto48@yahoo.co.id"
		i.Password = "12345"

		err := i.IsValidPassword("123456")
		assert.Error(t, err)
	})
}
