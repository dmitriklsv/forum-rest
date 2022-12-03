package repository

import (
	"context"
	"log"
	"time"

	"forum/internal/entity"
	"forum/pkg/sqlite3"
)

type sessionDB struct {
	storage *sqlite3.DB
}

func NewSessionRepo(database *sqlite3.DB) SessionRepo {
	log.Println("| | session repository is done!")
	return &sessionDB{
		storage: database,
	}
}

func (d *sessionDB) GetSession(ctx context.Context, sessionToken string) (entity.Session, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT * FROM sessions WHERE session_token = ?`
	row := d.storage.Collection.QueryRowContext(ctx, query, sessionToken)

	session := entity.Session{}
	if err := row.Scan(&session.ID, &session.UserID, &session.Token, &session.ExpireTime); err != nil {
		return entity.Session{}, err
	}

	return session, nil
}

func (d *sessionDB) CreateSession(ctx context.Context, session entity.Session) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO sessions (user_id, session_token, expire_time) VALUES (?, ?, ?)`
	st, err := d.storage.Collection.PrepareContext(ctx, query)
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// fmt.Println(session)
	query := `UPDATE sessions SET session_token = ?, expire_time = ? WHERE user_id = ?`
	st, err := d.storage.Collection.PrepareContext(ctx, query)
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `DELETE FROM sessions WHERE user_id = ?`
	st, err := d.storage.Collection.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err = st.ExecContext(ctx, id); err != nil {
		return err
	}

	return nil
}
