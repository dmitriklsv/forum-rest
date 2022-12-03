package service

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/repository"
)

type postService struct {
	postRepo repository.PostRepo
}

func NewPostService(pRepo repository.PostRepo) PostService {
	log.Println("| | post service is done!")
	return &postService{
		postRepo: pRepo,
	}
}

func (p *postService) CreatePost(ctx context.Context, post entity.Post) (int64, error) {
	return p.postRepo.CreatePost(ctx, post)
}
