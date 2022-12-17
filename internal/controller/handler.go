package controller

import (
	"net/http"

	"forum/internal/controller/handlers"
	"forum/internal/service"
)

type Welcomer interface {
	WelcomePage(w http.ResponseWriter, r *http.Request)
	HomePage(w http.ResponseWriter, r *http.Request)
	Middleware(next http.HandlerFunc) http.HandlerFunc
}

type UserHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
}

type PostHandler interface {
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
}

type CategoryHandler interface {
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	GetCategoryByID(w http.ResponseWriter, r *http.Request)
}

type CommentHandler interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
	GetCommentByID(w http.ResponseWriter, r *http.Request)
	GetCommentsByPostID(w http.ResponseWriter, r *http.Request)
}

type ReactionHandler interface {
	SetPostReaction(w http.ResponseWriter, r *http.Request)
	SetCommentReaction(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Welcomer
	UserHandler
	PostHandler
	CategoryHandler
	CommentHandler
	ReactionHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Welcomer:        handlers.NewWelcomeHandler(services.Authentication),
		UserHandler:     handlers.NewUserHandler(services.Authentication),
		PostHandler:     handlers.NewPostHandler(services.PostService),
		CategoryHandler: handlers.NewCategoryHandler(services.CategoryService),
		CommentHandler:  handlers.NewCommentHandler(services.CommentService),
		ReactionHandler: handlers.NewReactionHandler(services.ReactionService),
	}
}
