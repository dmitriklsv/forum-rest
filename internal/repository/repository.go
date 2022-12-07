package repository

import (
	"context"

	"forum/internal/entity"
	"forum/pkg/sqlite3"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
	FindByID(ctx context.Context, id uint64) (entity.User, error)
	FindOne(ctx context.Context, user entity.User) (entity.User, error)
}

type SessionRepo interface {
	CreateSession(ctx context.Context, session entity.Session) error
	GetSession(ctx context.Context, sessionToken string) (entity.Session, error)
	UpdateSession(ctx context.Context, session entity.Session) (entity.Session, error)
	DeleteSession(ctx context.Context, id uint64) error
}

type PostRepo interface {
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID uint64) (entity.Post, error)
	CreatePost(ctx context.Context, post entity.Post) (int64, error)
}

type CategoryRepo interface {
	CreateCategory(ctx context.Context, categories []entity.Category) ([]int64, error)
}

type CommentRepo interface {
	CreateComment(ctx context.Context, comment entity.Comment) (int64, error)
}

type Repositories struct {
	UserRepo
	SessionRepo
	PostRepo
	CategoryRepo
	CommentRepo
}

func NewRepos(db *sqlite3.DB) *Repositories {
	return &Repositories{
		UserRepo:     NewUserRepo(db),
		SessionRepo:  NewSessionRepo(db),
		PostRepo:     NewPostRepo(db),
		CategoryRepo: NewCategoryRepo(db),
		CommentRepo:  NewCommentRepo(db),
	}
}
