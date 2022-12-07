package repository

import (
	"context"
	"database/sql"
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

	query := `INSERT INTO comments (user_id, post_id, text) VALUES (?, ?, ?)`
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
