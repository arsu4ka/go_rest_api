package sqlstore_test

import (
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/model"
	"github.com/arsu4ka/go_rest_api/internal/store"
	"github.com/arsu4ka/go_rest_api/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	userTest := model.TestUser(t)
	_, err := s.User().FindByEmail(userTest.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(userTest)

	u, err := s.User().FindByEmail(userTest.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, userTest.Email, u.Email)
}

func TestUserRepo_FindByID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	userTest := model.TestUser(t)
	s.User().Create(userTest)
	u, err := s.User().FindByID(userTest.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, userTest.ID, u.ID)
}
