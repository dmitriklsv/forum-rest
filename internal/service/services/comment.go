package services

import (
	"context"
	"log"

	"forum/internal/entity"
	"forum/internal/repository"
)

type commentService struct {
	commentRepo  repository.CommentRepo
	reactionRepo repository.CommentReactionRepo // make set reaction for comments
}

func NewCommentService(cRepo repository.CommentRepo, reactionRepo repository.CommentReactionRepo) *commentService {
	log.Println("| | comment service is done!")
	return &commentService{
		commentRepo:  cRepo,
		reactionRepo: reactionRepo,
	}
}

func (c *commentService) CreateComment(ctx context.Context, comment entity.Comment) (int64, error) {
	return c.commentRepo.CreateComment(ctx, comment)
}

func (c *commentService) GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error) {
	reactions, err := c.reactionRepo.GetReactionsByCommentID(ctx, commentID)
	if err != nil {
		return entity.Comment{}, err
	}

	comment, err := c.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return entity.Comment{}, err
	}
	comment.Rating = c.setRating(reactions)

	return comment, nil
}

func (c *commentService) GetCommentsByPostID(ctx context.Context, postID uint64) ([]entity.Comment, error) {
	comments, err := c.commentRepo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	comments_len := len(comments)
	for i := 0; i < comments_len; i++ {
		reactions, err := c.reactionRepo.GetReactionsByCommentID(ctx, comments[i].ID)
		if err != nil {
			return nil, err
		}

		comments[i].Rating = c.setRating(reactions)
	}

	return comments, nil
}

func (c *commentService) setRating(reactions []entity.CommentReaction) int64 {
	var rating int64

	for _, reaction := range reactions {
		if reaction.Reaction == 1 {
			rating++
		} else {
			rating--
		}
	}

	return rating
}
