package sqlstore

import (
	"database/sql"
	"github.com/alex-dev-master/golang-rest-example/internal/app/model"
	"github.com/alex-dev-master/golang-rest-example/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	Store *Store
}


// Create ...
func (r *UserRepository) Create(u *model.User) error {

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.Store.Db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES (?, ?)",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.Store.Db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = ?",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.Store.Db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = ?",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}