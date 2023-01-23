package store_test

import (
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/app/model"
	"github.com/arsu4ka/go_rest_api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "example@gmail.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "example@gmail.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: email,
	})

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Email, email)
}
