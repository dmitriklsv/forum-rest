package services

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/repository"
)

type categoryService struct {
	catRepo repository.CategoryRepo
}

func NewCategoryService(catRepo repository.CategoryRepo) *categoryService {
	log.Println("| | category service is done!")
	return &categoryService{
		catRepo: catRepo,
	}
}

func (c *categoryService) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return c.catRepo.GetAllCategories(ctx)
}

func (c *categoryService) GetCategoryByID(ctx context.Context, categoryID uint64) (entity.Category, error) {
	return c.catRepo.GetCategoryByID(ctx, categoryID)
}
