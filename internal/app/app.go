package app

import (
	"log"
	"net"
	"net/http"
	"time"

	"forum/internal/controller"
)

// check rest enpoint naming standards!!!
const (
	welcome = "/"
	// auth
	signup = "/signup"
	signin = "/signin"
	// post
	createPost  = "/create_post"
	getAllPosts = "/get_all_posts"
	getPostByID = "/get_post_by_id"
	// category
	getAllCategories = "/get_all_categories"
	getCategoryByID  = "/get_category_by_id"
	// comment
	createComment       = "/create_comment"
	getCommentsByPostID = "/get_comments_by_post_id"
	getCommentByID      = "/get_comment_by_id"
	// reaction
	setPostReaction    = "/set_post_reaction"
	setCommentReaction = "/set_comment_reaction"
	// getReactionsByPostID
	// getReactionsByCommentID
)

func Run(handlers *controller.Handlers) error {
	log.Println("| creating router...")
	router := http.NewServeMux()

	// auth
	router.HandleFunc(signup, handlers.SignUp)
	router.HandleFunc(signin, handlers.SignIn)

	// post
	router.Handle(createPost, handlers.Middleware(handlers.CreatePost))
	router.HandleFunc(getAllPosts, handlers.GetAllPosts)
	router.HandleFunc(getPostByID, handlers.GetPostByID)

	// cat
	router.HandleFunc(getAllCategories, handlers.GetAllCategories)
	router.HandleFunc(getCategoryByID, handlers.GetCategoryByID)

	// comment
	router.Handle(createComment, handlers.Middleware(handlers.CreateComment))
	router.HandleFunc(getCommentByID, handlers.GetCommentByID)
	router.HandleFunc(getCommentsByPostID, handlers.GetCommentsByPostID)

	// reaction
	router.Handle(setPostReaction, handlers.Middleware(handlers.SetPostReaction))

	// home
	router.Handle(welcome, handlers.Middleware(handlers.WelcomePage))
	// router.Handle(home, handlers.Middleware(handlers.HomePage))

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
