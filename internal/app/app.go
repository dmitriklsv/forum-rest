package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"

	"forum/internal/controller"
	"forum/internal/tool/config"
	"forum/internal/tool/customErr"
)

// TODO: CHECK REST ENDPOINT STANDARDS!!!
const (
	welcome = "/"
	// auth
	signup = "/signup"
	signin = "/signin"
	// post
	posts = "/posts/"
	// category
	categories = "/categories/"
	// comment
	createComment       = "/create_comment"
	getCommentsByPostID = "/get_comments_by_post_id"
	getCommentByID      = "/get_comment_by_id"
	// reaction
	setPostReaction    = "/set_post_reaction"
	setCommentReaction = "/set_comment_reaction"
)

func Run(handlers *controller.Handlers) error {
	log.Println("| creating router...")
	router := http.NewServeMux()

	// auth
	router.HandleFunc(signup, handlers.SignUp)
	router.HandleFunc(signin, handlers.SignIn)

	// post
	router.HandleFunc(posts, Posts(handlers))

	// cat
	router.HandleFunc(categories, Categories(handlers))

	// comment
	router.Handle(createComment, handlers.Middleware(handlers.CreateComment))
	router.HandleFunc(getCommentByID, handlers.GetCommentByID)
	router.HandleFunc(getCommentsByPostID, handlers.GetCommentsByPostID)

	// reaction
	router.Handle(setPostReaction, handlers.Middleware(handlers.SetPostReaction))
	router.Handle(setCommentReaction, handlers.Middleware(handlers.SetCommentReaction))

	// home
	router.HandleFunc(welcome, handlers.WelcomePage)
	// router.Handle(home, handlers.Middleware(handlers.HomePage))

	return ListenAndServe(router)
}

// TODO: MAKE CONFIG
func ListenAndServe(router *http.ServeMux) error {
	log.Println("| starting application...")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	server := &http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("| listening http://localhost:8080")
	return server.Serve(listener)
}

func Posts(handlers *controller.Handlers) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.Welcomer.Middleware(handlers.PostHandler.CreatePost)(w, r)
		case http.MethodGet:
			reg := regexp.MustCompile(`^/posts/(\d+)$`)
			if reg.MatchString(r.URL.String()) {
				handlers.PostHandler.GetPostByID(w, r.WithContext(context.WithValue(r.Context(), config.PostID, reg.FindStringSubmatch(r.URL.Path)[1])))
			} else if len(r.URL.Query()) == 1 && r.URL.Path == "/posts/" {
				if r.URL.Query().Has("own") || r.URL.Query().Has("liked") || r.URL.Query().Has("disliked") {
					handlers.Welcomer.Middleware(handlers.PostHandler.GetAllPosts)(w, r)
				} else {
					handlers.PostHandler.GetAllPosts(w, r)
				}
			} else if r.URL.Path == "/posts/" {
				handlers.PostHandler.GetAllPosts(w, r)
			} else {
				http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}

func Categories(handlers *controller.Handlers) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			reg := regexp.MustCompile(`^/categories/(\d+)$`)
			if reg.MatchString(r.URL.String()) {
				handlers.CategoryHandler.GetCategoryByID(w, r.WithContext(context.WithValue(r.Context(), config.CategoryID, reg.FindStringSubmatch(r.URL.Path)[1])))
			} else if r.URL.Path == "/categories/" {
				handlers.CategoryHandler.GetAllCategories(w, r)
			} else {
				http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}
}
