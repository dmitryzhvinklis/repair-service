package sqlstore

import (
	"database/sql"

	"github.com/dmitryzhvinklis/carfix_dev/internal/app/model"
	"github.com/dmitryzhvinklis/carfix_dev/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (id,first_name,last_name,phone_number,email,encrypted_password)VALUES ($1,$2,$3,$4,$5,$6) RETURNING id",
		u.ID,
		u.FirstName,
		u.LastName,
		u.PhoneNumber,
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, first_name, last_name, phone_number, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
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
