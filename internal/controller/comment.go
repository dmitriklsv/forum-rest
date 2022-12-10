package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"forum/internal/entity"
	"forum/internal/service"
	"forum/internal/tool/errors"
)

type commentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) CommentHandler {
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
		http.Error(w, errors.InvalidData, http.StatusBadRequest)
		return
	}

	commentID, err := c.service.CreateComment(r.Context(), comment)
	if err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(commentID); err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (c *commentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var comment entity.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, errors.Bruhhh, http.StatusBadRequest)
		return
	}

	comment, err := c.service.GetCommentByID(r.Context(), comment.ID)
	if err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (c *commentHandler) GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var comment entity.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, errors.Bruhhh, http.StatusBadRequest)
		return
	}

	comments, err := c.service.GetCommentsByPostID(r.Context(), comment.PostID)
	if err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}
}
