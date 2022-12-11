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

func (rct *reactionRepo) CreatePostReaction(ctx context.Context, reaction entity.PostReaction) error {
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

func (rct *reactionRepo) GetReactionByPost(ctx context.Context, userID, postID uint64) (entity.PostReaction, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM reactions WHERE post_id = ? AND user_id = ?`
	row := rct.storage.QueryRowContext(ctx, query, postID, userID)

	var reaction entity.PostReaction
	if err := row.Scan(&reaction.ID, &reaction.PostID, &reaction.UserID, &reaction.Reaction); err != nil {
		return entity.PostReaction{}, err
	}

	return reaction, nil
}

func (rct *reactionRepo) UpdatePostReaction(ctx context.Context, reaction entity.PostReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `UPDATE reactions SET reaction = ? WHERE post_id = ? AND user_id = ?`
	st, err := rct.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, reaction.Reaction, reaction.PostID, reaction.UserID); err != nil {
		return err
	}

	return nil
}

func (rct *reactionRepo) DeletePostReaction(ctx context.Context, reaction entity.PostReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `DELETE FROM reactions WHERE id = ?`
	st, err := rct.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, reaction.ID); err != nil {
		return err
	}

	return nil
}
