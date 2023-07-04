package repositories

import (
	"database/sql"

	"github.com/gabriel-hahn/devbook/models"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
