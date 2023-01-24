package teststore_test

import (
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/app/model"
	"github.com/arsu4ka/go_rest_api/internal/app/store"
	"github.com/arsu4ka/go_rest_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	s := teststore.New()
	userTest := model.TestUser(t)

	_, err := s.User().FindByEmail(userTest.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(userTest)
	u, err := s.User().FindByEmail(userTest.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, userTest.Email, u.Email)
}
