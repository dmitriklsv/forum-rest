package controller

import (
	"net/http"

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
}

type Handlers struct {
	Welcomer
	UserHandler
	PostHandler
	CategoryHandler
	CommentHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Welcomer:        NewWelcomeHandler(services.Authentication),
		UserHandler:     NewUserHandler(services.Authentication),
		PostHandler:     NewPostHandler(services.PostService),
		CategoryHandler: NewCategoryHandler(services.CategoryService),
		CommentHandler:  NewCommentHandler(services.CommentService),
	}
}
