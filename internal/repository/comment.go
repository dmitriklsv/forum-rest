package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type commentDB struct {
	storage *sql.DB
}

func NewCommentRepo(database *sqlite3.DB) CommentRepo {
	log.Println("| | comment repository is done!")
	return &commentDB{
		storage: database.Collection,
	}
}

func (c *commentDB) CreateComment(ctx context.Context, comment entity.Comment) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT EXISTS (SELECT 1 FROM posts WHERE id = ?)`
	row := c.storage.QueryRowContext(ctx, query, comment.PostID)
	var t int
	if err := row.Scan(&t); err != nil {
		return 0, err
	}
	if t == 0 {
		return 0, fmt.Errorf("invalid post")
	}

	query = `INSERT INTO comments (user_id, post_id, text) VALUES (?, ?, ?)`
	st, err := c.storage.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer st.Close()

	res, err := st.ExecContext(ctx, comment.UserID, comment.PostID, comment.Text)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (c *commentDB) GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM comments WHERE id = ?`
	row := c.storage.QueryRowContext(ctx, query, commentID)

	var comment entity.Comment
	if err := row.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Text); err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (c *commentDB) GetCommentsByPostID(ctx context.Context, postID uint64) ([]entity.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM comments WHERE post_id = ?`
	rows, err := c.storage.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		var comment entity.Comment
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Text); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}
