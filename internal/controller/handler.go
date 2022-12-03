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
	CreatePost(w http.ResponseWriter, r *http.Request)
	// GetPostByID()
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	// GetAllPostsByCategoryID()
}

type CommentHandler interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Welcomer
	UserHandler
	PostHandler
	CommentHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Welcomer:       NewWelcomeHandler(services.Authentication),
		UserHandler:    NewUserHandler(services.Authentication),
		PostHandler:    NewPostHandler(services.PostService),
		CommentHandler: NewCommentHandler(services.CommentService),
	}
}
