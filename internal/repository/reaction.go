package repository

import (
	"context"
	"database/sql"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type reactionRepo struct {
	storage *sql.DB
}

func NewReactionRepo(database *sqlite3.DB) ReactionRepo {
	return &reactionRepo{
		storage: database.Collection,
	}
}

func (rct *reactionRepo) SetPostReaction(ctx context.Context, reaction entity.PostReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `INSERT INTO reactions (post_id, user_id, reaction) VALUES (?, ?, ?)`
	st, err := rct.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, reaction.PostID, reaction.UserID, reaction.Reaction); err != nil {
		return err
	}

	return nil
}
