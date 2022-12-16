package sqlite_repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
		// fmt.Printf("%#v\n", categories)
		query += ` WHERE id IN (SELECT post_id FROM categories WHERE name IN (?))`
		rows, err = p.storage.QueryContext(ctx, query, categories.(string))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	defer rows.Close()

	// fmt.Println(rows)

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Text); err != nil {
			return nil, err
		}
		fmt.Println(post)
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(posts)
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
