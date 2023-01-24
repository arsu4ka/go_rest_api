package sqlstore_test

import (
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/app/model"
	"github.com/arsu4ka/go_rest_api/internal/app/store"
	"github.com/arsu4ka/go_rest_api/internal/app/store/sqlstore"
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

	u_test := model.TestUser(t)
	_, err := s.User().FindByEmail(u_test.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u_test)

	u, err := s.User().FindByEmail(u_test.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u_test.Email, u.Email)
}
