package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gabriel-hahn/devbook/internal/crypto"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type UserResponse struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ValidationType string

const (
	Signup ValidationType = "signup"
	Update ValidationType = "update"
)

func (u *User) Prepare(step ValidationType) error {
	if step == Signup {
		if err := u.validate(); err != nil {
			return err
		}

		if err := u.generatePasswordHash(); err != nil {
			return err
		}
	}

	u.trimValues()

	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Nick == "" {
		return errors.New("nick is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email format")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) trimValues() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}

func (u *User) generatePasswordHash() error {
	hashedPassword, err := crypto.GenerateHash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}
