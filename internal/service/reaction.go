package service

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

func NewReactionService(curReactionRepo repository.ReactionRepo) ReactionService {
	return &reactionService{
		reactionRepo: curReactionRepo,
	}
}

func (r *reactionService) SetPostReaction(ctx context.Context, sentReaction entity.PostReaction) error {
	if sentReaction.Reaction < 0 || sentReaction.Reaction > 1 {
		return fmt.Errorf("invalid reaction type")
	}

	curReaction, err := r.reactionRepo.GetReactionByPost(ctx, sentReaction.UserID, sentReaction.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return r.reactionRepo.CreatePostReaction(ctx, sentReaction)
		}
		return err
	}

	if sentReaction.Reaction == curReaction.Reaction {
		return r.reactionRepo.DeletePostReaction(ctx, curReaction)
	}

	return r.reactionRepo.UpdatePostReaction(ctx, sentReaction)
}
