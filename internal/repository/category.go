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
		switch {
		case err == nil:
		case err.Error() == "UNIQUE constraint failed: categories.name":
			createdCatIDs[i] = -1
			continue
		default:
			return nil, err
		}

		id, _ := res.LastInsertId()
		createdCatIDs[i] = id
	}

	return createdCatIDs, nil
}

func (c *categoryDB) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()
	// TODO: read join table
	query := `SELECT * FROM categories`
	rows, err := c.storage.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.PostID, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryDB) GetCategoryByID(ctx context.Context, categoryID uint64) (entity.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM categories WHERE id = ?`
	row := c.storage.QueryRowContext(ctx, query, categoryID)

	var category entity.Category
	if err := row.Scan(&category.ID, &category.ID, &category.Name); err != nil {
		return entity.Category{}, err
	}

	return category, nil
}
