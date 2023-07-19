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

func (p Posts) FindAllByUserID(userID uint64) ([]model.Post, error) {
	rows, err := p.db.Query("select distinct p.*, u.nick from posts p join users u on p.author_id = u.id left join followers f on u.id = f.user_id where u.id = ? or f.follower_id = ? order by 1 desc", userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
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

func (p Posts) UpdateByID(ID uint64, post model.Post) error {
	statement, err := p.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, ID); err != nil {
		return err
	}

	return nil
}
