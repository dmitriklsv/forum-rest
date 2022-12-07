package repository

import (
	"context"
	"database/sql"
	"log"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type postDB struct {
	storage *sql.DB
}

func NewPostRepo(database *sqlite3.DB) PostRepo {
	log.Println("| | post repository is done!")
	return &postDB{
		storage: database.Collection,
	}
}

func (p *postDB) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()
	// TODO: read join table
	query := `SELECT * FROM posts`
	rows, err := p.storage.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		post := &entity.Post{}
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Text); err != nil {
			return nil, err
		}

		query := `SELECT category_id FROM postcategory WHERE post_id = ?`
		rows, err := p.storage.QueryContext(ctx, query, post.ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var category entity.Category
		for rows.Next() {
			if err := rows.Scan(&category.ID); err != nil {
				return nil, err
			}

			query := `SELECT * FROM categories WHERE id = ?`
			row := p.storage.QueryRowContext(ctx, query, category.ID)

			if err := row.Scan(&category.ID, &category.Name); err != nil {
				return nil, err
			}

			post.Categories = append(post.Categories, category)
		}

		posts = append(posts, *post)
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

	query = `SELECT category_id FROM postcategory WHERE post_id = ?`
	rows, err := p.storage.QueryContext(ctx, query, postID)
	if err != nil {
		return entity.Post{}, err
	}
	defer rows.Close()

	var category entity.Category
	for rows.Next() {
		if err := rows.Scan(&category.ID); err != nil {
			return entity.Post{}, err
		}

		query := `SELECT * FROM categories WHERE id = ?`
		row := p.storage.QueryRowContext(ctx, query, category.ID)

		if err := row.Scan(&category.ID, &category.Name); err != nil {
			return entity.Post{}, err
		}

		post.Categories = append(post.Categories, category)
	}

	if err := rows.Err(); err != nil {
		return entity.Post{}, err
	}

	return post, nil
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
	// createdPostID, _ := res.LastInsertId()

	// var temp string
	// createdCatIDs := make([]int64, len(post.Categories))
	// for i := range post.Categories {
	// 	row := p.storage.QueryRow(`SELECT * FROM categories WHERE name = ?`, post.Categories[i].Name)
	// 	if err := row.Scan(&createdCatIDs[i], &temp); err == nil {
	// 		continue
	// 	}

	// 	query := `INSERT INTO categories (name) VALUES (?)`
	// 	st, err := p.storage.PrepareContext(ctx, query)
	// 	if err != nil {
	// 		return -1, err
	// 	}
	// 	defer st.Close()

	// 	res, err := st.ExecContext(ctx, post.Categories[i].Name)
	// 	if err != nil {
	// 		return -1, err
	// 	}

	// 	id, err := res.LastInsertId()
	// 	if err != nil {
	// 		return -1, err
	// 	}
	// 	createdCatIDs[i] = id
	// }

	// // fmt.Printf("this is list of categories id: %v\nthis is post id: %d\n", createdCatIDs, createdPostID)

	// for i := 0; i < len(createdCatIDs); i++ {
	// 	query := `INSERT INTO postcategory (post_id, category_id) VALUES (?, ?)`
	// 	st, err := p.storage.PrepareContext(ctx, query)
	// 	if err != nil {
	// 		return -1, err
	// 	}
	// 	defer st.Close()

	// 	if _, err := st.ExecContext(ctx, createdPostID, createdCatIDs[i]); err != nil {
	// 		return -1, err
	// 	}

	// 	// fmt.Println(res.LastInsertId())
	// }

	return res.LastInsertId()
}
