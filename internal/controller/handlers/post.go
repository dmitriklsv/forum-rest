package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"forum/internal/entity"
	"forum/internal/service"
	"forum/internal/tool/config"
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
	defer r.Body.Close()

	userID := r.Context().Value(config.UserID)
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

	gayson.SendJSON(w, postID)
}

func (p *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) == 1 {
		if r.URL.Query().Has("category") {
			categories := r.URL.Query()["category"]
			if len(categories) == 0 {
				http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), config.Categories, categories))
		} else if r.URL.Query().Has("own") {
			r = r.WithContext(context.WithValue(r.Context(), config.Filter, "own"))
		} else if r.URL.Query().Has("liked") {
			r = r.WithContext(context.WithValue(r.Context(), config.Filter, "liked"))
		} else if r.URL.Query().Has("disliked") {
			r = r.WithContext(context.WithValue(r.Context(), config.Filter, "disliked"))
		}
	}

	posts, err := p.service.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	gayson.SendJSON(w, posts)
}

func (p *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseUint(r.Context().Value(config.PostID).(string), 10, 64)
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
