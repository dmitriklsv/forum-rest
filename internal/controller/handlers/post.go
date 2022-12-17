package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"forum/internal/entity"
	"forum/internal/service"
	"forum/internal/tool/customErr"
	"forum/pkg/gayson"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	log.Println("| | post handler is done!")
	return &postHandler{
		service: service,
	}
}

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1)
	defer r.Body.Close()

	userID := r.Context().Value(userCtx)
	post := entity.Post{
		UserID: userID.(uint64),
	}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, customErr.InvalidData, http.StatusBadRequest)
		return
	}

	postID, err := p.service.CreatePost(r.Context(), post)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	fmt.Println(postID)
	gayson.SendJSON(w, postID)
}

// TODO: FILTER BY CREATED POSTS: NEWESET, OLDEST
func (p *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) == 1 {
		categories := r.URL.Query()["category"]
		if len(categories) == 0 {
			http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
			return
		}

		// var cat ctx = "categories"
		r = r.WithContext(context.WithValue(r.Context(), "categories", categories))
	}

	posts, err := p.service.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	gayson.SendJSON(w, posts)
}

func (p *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseUint(r.Context().Value("post_ID").(string), 10, 64)
	if err != nil {
		http.Error(w, customErr.InvalidData, http.StatusBadRequest)
		return
	}

	post, err := p.service.GetPostByID(r.Context(), postID)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	gayson.SendJSON(w, post)
}
