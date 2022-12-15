package sqlite_repo

import (
	"context"
	"database/sql"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type commentReactionRepo struct {
	storage *sql.DB
}

func NewCommentReactionRepo(database *sqlite3.DB) *commentReactionRepo {
	return &commentReactionRepo{
		storage: database.Collection,
	}
}

func (rct *commentReactionRepo) CreateCommentReaction(ctx context.Context, reaction entity.CommentReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `INSERT INTO comment_reactions (comment_id, user_id, reaction) VALUES (?, ?, ?)`
	st, err := rct.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, reaction.CommentID, reaction.UserID, reaction.Reaction); err != nil {
		return err
	}

	return nil
}

func (rct *commentReactionRepo) GetReactionsByCommentID(ctx context.Context, commentID uint64) ([]entity.CommentReaction, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM comment_reactions WHERE comment_id = ?`
	rows, err := rct.storage.QueryContext(ctx, query, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reactions []entity.CommentReaction
	for rows.Next() {
		var reaction entity.CommentReaction
		if err := rows.Scan(&reaction.ID, &reaction.CommentID, &reaction.UserID, &reaction.Reaction); err != nil {
			return nil, err
		}

		reactions = append(reactions, reaction)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return reactions, nil
}

func (rct *commentReactionRepo) GetReactionByComment(ctx context.Context, userID, commentID uint64) (entity.CommentReaction, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM comment_reactions WHERE comment_id = ? AND user_id = ?`
	row := rct.storage.QueryRowContext(ctx, query, commentID, userID)

	var reaction entity.CommentReaction
	if err := row.Scan(&reaction.ID, &reaction.CommentID, &reaction.UserID, &reaction.Reaction); err != nil {
		return entity.CommentReaction{}, err
	}

	return reaction, nil
}

func (rct *commentReactionRepo) UpdateCommentReaction(ctx context.Context, reaction entity.CommentReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `UPDATE comment_reactions SET reaction = ? WHERE comment_id = ? AND user_id = ?`
	st, err := rct.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, reaction.Reaction, reaction.CommentID, reaction.UserID); err != nil {
		return err
	}

	return nil
}

func (rct *commentReactionRepo) DeleteCommentReaction(ctx context.Context, reaction entity.CommentReaction) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `DELETE FROM comment_reactions WHERE id = ?`
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
