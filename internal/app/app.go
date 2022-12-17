package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"

	"forum/internal/controller"
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
	getAllCategories = "/get_all_categories"
	getCategoryByID  = "/get_category_by_id"
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
	router.HandleFunc(posts, func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.URL.Path)
		switch r.Method {
		case http.MethodPost:
			handlers.Middleware(handlers.CreatePost)(w, r)
		case http.MethodGet:
			reg := regexp.MustCompile(`^/posts/(\d+)$`)
			if reg.MatchString(r.URL.String()) {
				// var post_ID ctx = "post_ID"
				handlers.GetPostByID(w, r.WithContext(context.WithValue(r.Context(), "post_ID", reg.FindStringSubmatch(r.URL.Path)[1])))
			} else if len(r.URL.Query()) == 1 && r.URL.Path == "/posts/" {
				handlers.GetAllPosts(w, r)
			} else {
				http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
				return
			}
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	})

	// cat
	router.HandleFunc(getAllCategories, handlers.GetAllCategories)
	router.HandleFunc(getCategoryByID, handlers.GetCategoryByID)

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
