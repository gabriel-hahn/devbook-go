package repository

import (
	"database/sql"

	"github.com/gabriel-hahn/devbook/internal/model"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (p Posts) Create(post model.Post) (uint64, error) {
	statement, err := p.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

func (p Posts) FindByID(postID uint64) (model.Post, error) {
	row, err := p.db.Query("select p.*, u.nick from posts p inner join users u on u.id = p.author_id where p.id = ?", postID)
	if err != nil {
		return model.Post{}, err
	}
	defer row.Close()

	var post model.Post
	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return model.Post{}, err
		}
	}

	return post, nil
}
