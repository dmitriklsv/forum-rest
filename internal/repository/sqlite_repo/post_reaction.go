package sqlite_repo

import (
	"context"
	"database/sql"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type postReactionRepo struct {
	storage *sql.DB
}

func NewPostReactionRepo(database *sqlite3.DB) *postReactionRepo {
	return &postReactionRepo{
		storage: database.Collection,
	}
}

func (rct *postReactionRepo) CreatePostReaction(ctx context.Context, reaction entity.PostReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `INSERT INTO post_reactions (post_id, user_id, reaction) VALUES (?, ?, ?)`
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

func (rct *postReactionRepo) GetReactionsByPostID(ctx context.Context, postID uint64) ([]entity.PostReaction, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM post_reactions WHERE post_id = ?`
	rows, err := rct.storage.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reactions []entity.PostReaction
	for rows.Next() {
		var reaction entity.PostReaction
		if rows.Scan(&reaction.ID, &reaction.PostID, &reaction.UserID, &reaction.Reaction); err != nil {
			return nil, err
		}

		reactions = append(reactions, reaction)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return reactions, nil
}

func (rct *postReactionRepo) GetReactionByPost(ctx context.Context, userID, postID uint64) (entity.PostReaction, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM post_reactions WHERE post_id = ? AND user_id = ?`
	row := rct.storage.QueryRowContext(ctx, query, postID, userID)

	var reaction entity.PostReaction
	if err := row.Scan(&reaction.ID, &reaction.PostID, &reaction.UserID, &reaction.Reaction); err != nil {
		return entity.PostReaction{}, err
	}

	return reaction, nil
}

func (rct *postReactionRepo) UpdatePostReaction(ctx context.Context, reaction entity.PostReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `UPDATE post_reactions SET reaction = ? WHERE post_id = ? AND user_id = ?`
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

func (rct *postReactionRepo) DeletePostReaction(ctx context.Context, reaction entity.PostReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `DELETE FROM post_reactions WHERE id = ?`
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
