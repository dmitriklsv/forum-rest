package repository

import (
	"context"
	"log"
	"time"

	"forum/internal/entity"
	"forum/pkg/sqlite3"
)

type userDB struct {
	storage *sqlite3.DB
}

func NewUserRepo(database *sqlite3.DB) UserRepo {
	log.Println("| | user repository is done!")
	return &userDB{
		storage: database,
	}
}

func (d *userDB) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `INSERT INTO users(email, username, password) VALUES (?, ?, ?)`
	st, err := d.storage.Collection.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer st.Close()

	res, err := st.ExecContext(ctx, user.Email, user.Username, user.Password)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (d *userDB) FindByID(ctx context.Context, id uint64) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id = ? `
	row := d.storage.Collection.QueryRowContext(ctx, query, id)

	user := entity.User{}
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (d *userDB) FindOne(ctx context.Context, user entity.User) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// fmt.Println(user)
	query := `SELECT * FROM users WHERE email = ? OR (username = ? AND password = ?)`
	row := d.storage.Collection.QueryRowContext(ctx, query, user.Email, user.Username, user.Password)

	user = entity.User{}
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		return entity.User{}, err
	}

	return user, nil
}
