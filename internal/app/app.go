package app

import (
	"log"
	"net"
	"net/http"
	"time"

	"forum/internal/controller"
)

const (
	welcome = "/"
	// home       = "/home"
	signup        = "/signup"
	signin        = "/signin"
	createPost    = "/create_post"
	createComment = "/create_comment"
	getAllPosts   = "/get_all_posts"
	getPostByID   = "/get_post_by_id"
)

func Run(handlers *controller.Handlers) error {
	log.Println("| creating router...")
	router := http.NewServeMux()

	// auth
	router.HandleFunc(signup, handlers.SignUp)
	router.HandleFunc(signin, handlers.SignIn)

	// home
	router.HandleFunc(getAllPosts, handlers.GetAllPosts)
	router.HandleFunc(getPostByID, handlers.GetPostByID)
	router.Handle(welcome, handlers.Middleware(handlers.WelcomePage))
	// router.Handle(home, handlers.Middleware(handlers.HomePage))
	router.Handle(createPost, handlers.Middleware(handlers.CreatePost))
	router.Handle(createComment, handlers.Middleware(handlers.CreateComment))

	return ListenAndServe(router)
}

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
