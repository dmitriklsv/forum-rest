package sqlite_repo

import (
	"context"
	"database/sql"
	"log"

	"forum/internal/entity"
	"forum/internal/tool/config"
	"forum/pkg/sqlite3"
)

type sessionDB struct {
	storage *sql.DB
}

func NewSessionRepo(database *sqlite3.DB) *sessionDB {
	log.Println("| | session repository is done!")
	return &sessionDB{
		storage: database.Collection,
	}
}

func (d *sessionDB) GetSession(ctx context.Context, sessionToken string) (entity.Session, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `SELECT * FROM sessions WHERE session_token = ?`
	row := d.storage.QueryRowContext(ctx, query, sessionToken)

	var session entity.Session
	if err := row.Scan(&session.ID, &session.UserID, &session.Token, &session.ExpireTime); err != nil {
		return entity.Session{}, err
	}

	return session, nil
}

func (d *sessionDB) CreateSession(ctx context.Context, session entity.Session) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `INSERT INTO sessions (user_id, session_token, expire_time) VALUES (?, ?, ?)`
	st, err := d.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, session.UserID, session.Token, session.ExpireTime); err != nil {
		return err
	}

	return nil
}

func (d *sessionDB) UpdateSession(ctx context.Context, session entity.Session) (entity.Session, error) {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	// fmt.Println(session)
	query := `UPDATE sessions SET session_token = ?, expire_time = ? WHERE user_id = ?`
	st, err := d.storage.PrepareContext(ctx, query)
	if err != nil {
		return session, err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, session.Token, session.ExpireTime, session.UserID); err != nil {
		return session, err
	}

	return session, nil
}

func (d *sessionDB) DeleteSession(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, config.DefaultTimeout)
	defer cancel()

	query := `DELETE FROM sessions WHERE user_id = ?`
	st, err := d.storage.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
