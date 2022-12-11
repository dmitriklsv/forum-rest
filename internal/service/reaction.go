package service

import (
	"context"
	"database/sql"
	"errors"

	"forum/internal/entity"
	"forum/internal/repository"
)

type reactionService struct {
	reactionRepo repository.ReactionRepo
}

func NewReactionService(rctRepo repository.ReactionRepo) ReactionService {
	return &reactionService{
		reactionRepo: rctRepo,
	}
}

func (r *reactionService) SetPostReaction(ctx context.Context, reaction entity.PostReaction) error {
	rct, err := r.reactionRepo.GetReactionByPost(ctx, reaction.UserID, reaction.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return r.reactionRepo.CreatePostReaction(ctx, reaction)
		}
		return err
	}

	switch rct.Reaction.Type {
	case "like":
		if rct.Reaction.State {
			return r.reactionRepo.UpdatePostReaction(ctx, reaction)
		}
		return r.reactionRepo.DeletePostReaction(ctx, rct)
	case "dislike":
		if rct.Reaction.State {
			return r.reactionRepo.UpdatePostReaction(ctx, reaction)
		}
		return r.reactionRepo.DeletePostReaction(ctx, rct)
	}

	return nil
}
