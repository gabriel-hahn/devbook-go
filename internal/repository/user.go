package repository

import (
	"database/sql"
	"fmt"

	"github.com/gabriel-hahn/devbook/internal/model"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(user model.User) (uint64, error) {
	statement, err := u.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

func (u Users) FindAllByFilters(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := u.db.Query("select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u Users) FindByID(userID uint64) (model.User, error) {
	rows, err := u.db.Query("select id, name, nick, email, created_at from users where id = ?", userID)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

func (u Users) FindByEmail(userEmail string) (model.User, error) {
	rows, err := u.db.Query("select id, password from users where email = ?", userEmail)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		if err = rows.Scan(&user.ID, &user.Password); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

func (u Users) UpdateByID(ID uint64, user model.User) error {
	statement, err := u.db.Prepare("update users set name = ?, nick = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, ID); err != nil {
		return err
	}

	return nil
}

func (u Users) DeleteByID(ID uint64) error {
	statement, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (u Users) Follow(followerID, userID uint64) error {
	statement, err := u.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (u Users) Unfollow(unfollowID, userID uint64) error {
	statement, err := u.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(unfollowID, userID); err != nil {
		return err
	}

	return nil
}
