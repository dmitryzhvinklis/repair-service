package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	PhoneNumber       string `json:"phonenumber"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
	Password          string `json:"password,omiempty"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(
			&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
