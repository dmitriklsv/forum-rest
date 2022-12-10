package service

import (
	"context"

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
	return r.reactionRepo.SetPostReaction(ctx, reaction)
}
