package service

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/repository"
)

type commentService struct {
	commentRepo repository.CommentRepo
}

func NewCommentService(cRepo repository.CommentRepo) CommentService {
	log.Println("| | comment service is done!")
	return &commentService{
		commentRepo: cRepo,
	}
}

func (c *commentService) CreateComment(ctx context.Context, comment entity.Comment) (int64, error) {
	return c.commentRepo.CreateComment(ctx, comment)
}
