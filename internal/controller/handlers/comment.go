package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"forum/internal/entity"
	"forum/internal/service"
	"forum/internal/tool/customErr"
)

type commentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) *commentHandler {
	log.Println("| | comment handler is done!")
	return &commentHandler{
		service: service,
	}
}

func (c *commentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	userID := r.Context().Value(userCtx)
	comment := entity.Comment{
		UserID: userID.(uint64),
	}
	// TODO: create customer
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, customErr.InvalidData, http.StatusBadRequest)
		return
	}

	commentID, err := c.service.CreateComment(r.Context(), comment)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(commentID); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (c *commentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	commentID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
		return
	}

	comment, err := c.service.GetCommentByID(r.Context(), commentID)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	// fmt.Println(comment)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (c *commentHandler) GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	postID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
		return
	}

	comments, err := c.service.GetCommentsByPostID(r.Context(), postID)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}
