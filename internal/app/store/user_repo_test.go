package store_test

import (
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/app/model"
	"github.com/arsu4ka/go_rest_api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u_test := model.TestUser(t)
	_, err := s.User().FindByEmail(u_test.Email)
	assert.Error(t, err)

	s.User().Create(u_test)

	u, err := s.User().FindByEmail(u_test.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u_test.Email, u.Email)
}
