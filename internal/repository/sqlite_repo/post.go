package sqlite_repo

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type postDB struct {
	storage *sql.DB
}

func NewPostRepo(database *sqlite3.DB) *postDB {
	log.Println("| | post repository is done!")
	return &postDB{
		storage: database.Collection,
	}
}

func (p *postDB) CreatePost(ctx context.Context, post entity.Post) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `INSERT INTO posts (user_id, title, text) VALUES (?, ?, ?)`
	st, err := p.storage.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer st.Close()

	res, err := st.ExecContext(ctx, post.UserID, post.Title, post.Text)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (p *postDB) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM posts`
	rows, err := p.storage.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	categories := ctx.Value("categories")
	if categories != nil {
		categories_len := len(categories.([]string))
		args := make([]any, categories_len)
		for i, category := range categories.([]string) {
			args[i] = category
		}

		marks := strings.Repeat(" AND post_id IN (SELECT post_id FROM categories WHERE name = ?)", categories_len)
		query += ` WHERE id IN (SELECT DISTINCT post_id FROM categories WHERE ` + marks[5:] + `)`
		rows, err = p.storage.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Text); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *postDB) GetPostByID(ctx context.Context, postID uint64) (entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM posts WHERE id = ?`
	row := p.storage.QueryRowContext(ctx, query, postID)

	var post entity.Post
	if err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Text); err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
