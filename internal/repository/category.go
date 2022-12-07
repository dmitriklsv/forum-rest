package repository

import (
	"context"
	"database/sql"
	"log"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type categoryDB struct {
	storage *sql.DB
}

func NewCategoryRepo(database *sqlite3.DB) CategoryRepo {
	log.Println("| | category repository is done!")
	return &categoryDB{
		storage: database.Collection,
	}
}

func (c *categoryDB) CreateCategory(ctx context.Context, categories []entity.Category) ([]int64, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	createdCatIDs := make([]int64, len(categories))
	for i, category := range categories {
		query := `INSERT INTO categories (post_id, name) VALUES (?, ?)`
		st, err := c.storage.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}
		defer st.Close()

		res, err := st.ExecContext(ctx, category.PostID, category.Name)
		if err != nil {
			return nil, err
		}

		id, _ := res.LastInsertId()
		createdCatIDs[i] = id
	}

	return createdCatIDs, nil
}
