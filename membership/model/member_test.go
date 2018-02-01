package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMember(t *testing.T) {

	t.Run("New Member", func(t *testing.T) {
		m := NewMember()
		m.FirstName = "Wuriyanto"
		m.LastName = "Musobar"
		m.Email = "wuriyanto48@yahoo.co.id"
		m.Password = "12345"
		m.BirthDate = time.Now()

		assert.NotNil(t, m)
		assert.Equal(t, "Wuriyanto", m.FirstName)
		assert.Equal(t, "Musobar", m.LastName)
		assert.Equal(t, 0, m.Version)
	})

}
