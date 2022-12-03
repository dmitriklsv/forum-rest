package repository

import (
	"context"
	"log"
	"time"

	"forum/internal/entity"
	"forum/pkg/sqlite3"
)

type commentDB struct {
	storage *sqlite3.DB
}

func NewCommentRepo(database *sqlite3.DB) CommentRepo {
	log.Println("| | comment repository is done!")
	return &commentDB{
		storage: database,
	}
}

func (c *commentDB) CreateComment(ctx context.Context, comment entity.Comment) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO comments (user_id, post_id, text) VALUES (?, ?, ?)`
	st, err := c.storage.Collection.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := st.ExecContext(ctx, comment.UserID, comment.PostID, comment.Text)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
