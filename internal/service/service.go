package service

import (
	"context"

	"forum/internal/entity"
	"forum/internal/repository"
	"forum/internal/service/services"
)

type Authentication interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
	GetUser(ctx context.Context, id uint64) (entity.User, error)
	SetSession(ctx context.Context, user entity.User) (entity.Session, error)
	GetSession(ctx context.Context, sessionToken string) (entity.Session, error)
	UpdateSession(ctx context.Context, session entity.Session) (entity.Session, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post entity.Post) (int64, error)
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID uint64) (entity.Post, error)
}

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	GetCategoryByID(ctx context.Context, categoryID uint64) (entity.Category, error)
}

type CommentService interface {
	CreateComment(ctx context.Context, comment entity.Comment) (int64, error)
	GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID uint64) ([]entity.Comment, error)
}

type ReactionService interface {
	SetPostReaction(ctx context.Context, reaction entity.PostReaction) error
	SetCommentReaction(ctx context.Context, reaction entity.CommentReaction) error
}

type Services struct {
	Authentication
	PostService
	CategoryService
	CommentService
	ReactionService
}

func NewServices(repository *repository.Repositories) *Services {
	return &Services{
		Authentication:  services.NewAuthService(repository.UserRepo, repository.SessionRepo),
		PostService:     services.NewPostService(repository.PostRepo, repository.CategoryRepo, repository.ReactionRepo),
		CategoryService: services.NewCategoryService(repository.CategoryRepo),
		CommentService:  services.NewCommentService(repository.CommentRepo, repository.CommentReactionRepo),
		ReactionService: services.NewReactionService(repository.ReactionRepo),
	}
}
