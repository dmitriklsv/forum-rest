package repository

import (
	"context"
	"log"
	"time"

	"forum/internal/entity"
	"forum/pkg/sqlite3"
)

type postDB struct {
	storage *sqlite3.DB
}

func NewPostRepo(database *sqlite3.DB) PostRepo {
	log.Println("| | post repository is done!")
	return &postDB{
		storage: database,
	}
}

func (p *postDB) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT * FROM posts`
	rows, err := p.storage.Collection.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	posts := []entity.Post{}
	for rows.Next() {
		post := &entity.Post{}
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Text); err != nil {
			return nil, err
		}
		posts = append(posts, *post)
	}

	return posts, nil
}

func (p *postDB) GetPostByID(ctx context.Context, postID uint64) (entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT * FROM posts WHERE id = ?`
	row := p.storage.Collection.QueryRowContext(ctx, query, postID)

	post := entity.Post{}
	if err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Text); err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

func (p *postDB) CreatePost(ctx context.Context, post entity.Post) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO posts (user_id, title, text) VALUES (?, ?, ?)`
	st, err := p.storage.Collection.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer st.Close()

	res, err := st.ExecContext(ctx, post.UserID, post.Title, post.Text)
	if err != nil {
		return -1, err
	}
	createdPostID, _ := res.LastInsertId()

	createdCatIDs := make([]int64, len(post.Categories))
	for i := 0; i < len(post.Categories); i++ {
		var temp string // anyways garbage collector will delete it MUHAHAHAHAHA!!!
		row := p.storage.Collection.QueryRow(`SELECT * FROM categories WHERE name = ?`, post.Categories[i])
		if err := row.Scan(&createdCatIDs[i], &temp); err == nil {
			continue
		}

		query := `INSERT INTO categories (name) VALUES (?)`
		st, err := p.storage.Collection.PrepareContext(ctx, query)
		if err != nil {
			return -1, err
		}
		defer st.Close()

		res, err := st.ExecContext(ctx, post.Categories[i])
		if err != nil {
			return -1, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return -1, err
		}
		createdCatIDs[i] = id
	}

	// fmt.Printf("this is list of categories id: %v\nthis is post id: %d\n", createdCatIDs, createdPostID)

	for i := 0; i < len(createdCatIDs); i++ {
		query := `INSERT INTO postcategory (post_id, category_id) VALUES (?, ?)`
		st, err := p.storage.Collection.PrepareContext(ctx, query)
		if err != nil {
			return -1, err
		}
		defer st.Close()

		if _, err := st.ExecContext(ctx, createdPostID, createdCatIDs[i]); err != nil {
			return -1, err
		}

		// fmt.Println(res.LastInsertId())
	}

	return res.LastInsertId()
}
