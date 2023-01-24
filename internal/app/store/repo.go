package store

import "github.com/arsu4ka/go_rest_api/internal/app/model"

type UserRepo interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
