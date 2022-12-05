package service

import (
	"context"

	"forum/internal/entity"
	"forum/internal/repository"
)

type Authentication interface {
	SetUser(ctx context.Context, user entity.User) (int64, error)
	GetUser(ctx context.Context, id uint64) (entity.User, error)
	SetSession(ctx context.Context, user entity.User) (entity.Session, error)
	GetSession(ctx context.Context, sessionToken string) (entity.Session, error)
	UpdateSession(ctx context.Context, session entity.Session) (entity.Session, error)
}

type PostService interface {
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID uint64) (entity.Post, error)
	CreatePost(ctx context.Context, post entity.Post) (int64, error)
}

type CommentService interface {
	CreateComment(ctx context.Context, comment entity.Comment) (int64, error)
}

type Services struct {
	Authentication
	PostService
	CommentService
}

func NewServices(repository *repository.Repositories) *Services {
	return &Services{
		Authentication: NewAuthService(repository.UserRepo, repository.SessionRepo),
		PostService:    NewPostService(repository.PostRepo),
		CommentService: NewCommentService(repository.CommentRepo),
	}
}
