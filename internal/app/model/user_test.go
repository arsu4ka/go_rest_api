package model_test

import (
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreation())
	assert.NotEmpty(t, u.EncryptedPassword)
}
