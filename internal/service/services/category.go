package services

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/service"
)

type categoryService struct {
	catRepo service.CategoryRepo
}

func NewCategoryService(catRepo service.CategoryRepo) *categoryService {
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
