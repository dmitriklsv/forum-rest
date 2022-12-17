package sqlite_repo

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

func NewCategoryRepo(database *sqlite3.DB) *categoryDB {
	log.Println("| | category repository is done!")
	return &categoryDB{
		storage: database.Collection,
	}
}

func (c *categoryDB) CreateCategory(ctx context.Context, postID uint64, categories []entity.Category) /* []int64, */ error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `INSERT INTO categories (post_id, name) VALUES (?, ?)`
	st, err := c.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	for _, category := range categories {
		if _, err := st.ExecContext(ctx, postID, category.Name); err != nil {
			return err
		}
	}

	return nil
}

func (c *categoryDB) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

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
	st, err := c.storage.PrepareContext(ctx, query)
	if err != nil {
		return entity.Category{}, err
	}
	defer st.Close()

	row := st.QueryRowContext(ctx, categoryID)

	var category entity.Category
	if err = row.Scan(&category.ID, &category.PostID, &category.Name); err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (c *categoryDB) GetCategoriesByPostID(ctx context.Context, postID uint64) ([]entity.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM categories WHERE post_id = ?`
	st, err := c.storage.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.QueryContext(ctx, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		if err = rows.Scan(&category.ID, &category.PostID, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return categories, nil
}
