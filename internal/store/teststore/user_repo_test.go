package teststore_test

import (
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/model"
	"github.com/arsu4ka/go_rest_api/internal/store"
	"github.com/arsu4ka/go_rest_api/internal/store/teststore"
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

	err = s.User().Create(userTest)
	assert.NoError(t, err)
	u, err := s.User().FindByEmail(userTest.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, userTest.Email, u.Email)
}

func TestUserRepo_FindByID(t *testing.T) {
	s := teststore.New()
	userTest := model.TestUser(t)
	err := s.User().Create(userTest)
	assert.NoError(t, err)
	u, err := s.User().FindByID(userTest.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, userTest.ID, u.ID)
}
