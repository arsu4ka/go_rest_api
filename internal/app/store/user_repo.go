package store

import "github.com/arsu4ka/go_rest_api/internal/app/model"

type UserRepo struct {
	store *Store
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreation(); err != nil {
		return nil, err
	}

	err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}
