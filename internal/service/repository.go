package service

import (
	"context"

	"forum/internal/entity"
	"forum/internal/repository/sqlite_repo"
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
	CreatePost(ctx context.Context, post entity.Post) (int64, error)
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	// GetPostsByCategory(ctx context.Context, category string) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID uint64) (entity.Post, error)
}

type CategoryRepo interface {
	CreateCategory(ctx context.Context, postID uint64, categories []entity.Category) /* []int64,  */ error
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	GetCategoryByID(ctx context.Context, categoryID uint64) (entity.Category, error)
	GetCategoriesByPostID(ctx context.Context, postID uint64) ([]entity.Category, error)
}

type CommentRepo interface {
	CreateComment(ctx context.Context, comment entity.Comment) (int64, error)
	GetCommentByID(ctx context.Context, commentID uint64) (entity.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID uint64) ([]entity.Comment, error)
}

type PostReactionRepo interface {
	CreatePostReaction(ctx context.Context, reaction entity.PostReaction) error
	GetReactionsByPostID(ctx context.Context, postID uint64) ([]entity.PostReaction, error)
	GetReactionByPost(ctx context.Context, userID, postID uint64) (entity.PostReaction, error)
	UpdatePostReaction(ctx context.Context, reaction entity.PostReaction) error
	DeletePostReaction(ctx context.Context, reaction entity.PostReaction) error
}

type CommentReactionRepo interface {
	CreateCommentReaction(ctx context.Context, reaction entity.CommentReaction) error
	GetReactionsByCommentID(ctx context.Context, commentID uint64) ([]entity.CommentReaction, error)
	GetReactionByComment(ctx context.Context, userID, commentID uint64) (entity.CommentReaction, error)
	UpdateCommentReaction(ctx context.Context, reaction entity.CommentReaction) error
	DeleteCommentReaction(ctx context.Context, reaction entity.CommentReaction) error
}

type ReactionRepo struct {
	PostReactionRepo
	CommentReactionRepo
}

type Repositories struct {
	UserRepo
	SessionRepo
	PostRepo
	CategoryRepo
	CommentRepo
	ReactionRepo
}

func NewRepos(db *sqlite3.DB) *Repositories {
	return &Repositories{
		UserRepo:     sqlite_repo.NewUserRepo(db),
		SessionRepo:  sqlite_repo.NewSessionRepo(db),
		PostRepo:     sqlite_repo.NewPostRepo(db),
		CategoryRepo: sqlite_repo.NewCategoryRepo(db),
		CommentRepo:  sqlite_repo.NewCommentRepo(db),
		ReactionRepo: ReactionRepo{
			sqlite_repo.NewPostReactionRepo(db),
			sqlite_repo.NewCommentReactionRepo(db),
		},
	}
}
