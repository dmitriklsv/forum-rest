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

	// fmt.Println(comment)

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
