package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"forum/internal/entity"
	"forum/internal/repository"
)

type reactionService struct {
	reactionRepo repository.ReactionRepo
}

func NewReactionService(curReactionRepo repository.ReactionRepo) *reactionService {
	return &reactionService{
		reactionRepo: curReactionRepo,
	}
}

func (r *reactionService) SetPostReaction(ctx context.Context, sentReaction entity.PostReaction) error {
	if sentReaction.Reaction < 0 || sentReaction.Reaction > 1 {
		return fmt.Errorf("invalid reaction type")
	}

	curReaction, err := r.reactionRepo.PostReactionRepo.GetReactionByPost(ctx, sentReaction.UserID, sentReaction.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) && sentReaction.PostID != 0 {
			return r.reactionRepo.PostReactionRepo.CreatePostReaction(ctx, sentReaction)
		}
		return err
	}

	if sentReaction.Reaction == curReaction.Reaction {
		return r.reactionRepo.PostReactionRepo.DeletePostReaction(ctx, curReaction)
	}

	return r.reactionRepo.PostReactionRepo.UpdatePostReaction(ctx, sentReaction)
}

func (r *reactionService) SetCommentReaction(ctx context.Context, sentReaction entity.CommentReaction) error {
	if sentReaction.Reaction < 0 || sentReaction.Reaction > 1 {
		return fmt.Errorf("invalid reaction type")
	}

	curReaction, err := r.reactionRepo.CommentReactionRepo.GetReactionByComment(ctx, sentReaction.UserID, sentReaction.CommentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) && sentReaction.CommentID != 0 {
			return r.reactionRepo.CommentReactionRepo.CreateCommentReaction(ctx, sentReaction)
		}
		return err
	}

	if sentReaction.Reaction == curReaction.Reaction {
		return r.reactionRepo.CommentReactionRepo.DeleteCommentReaction(ctx, curReaction)
	}

	return r.reactionRepo.CommentReactionRepo.UpdateCommentReaction(ctx, sentReaction)
}
