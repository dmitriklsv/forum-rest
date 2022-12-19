package handlers

import (
	"encoding/json"
	"net/http"

	"forum/internal/entity"
	"forum/internal/service"
	"forum/internal/tool/config"
	"forum/internal/tool/customErr"
)

type reactionHandler struct {
	service service.ReactionService
}

func NewReactionHandler(service service.ReactionService) *reactionHandler {
	return &reactionHandler{
		service: service,
	}
}

func (rct *reactionHandler) SetPostReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	userID := r.Context().Value(config.UserID)
	reaction := entity.PostReaction{
		UserID: userID.(uint64),
	}

	if err := json.NewDecoder(r.Body).Decode(&reaction); err != nil {
		http.Error(w, customErr.InvalidData, http.StatusBadRequest)
		return
	}

	if err := rct.service.SetPostReaction(r.Context(), reaction); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (rct *reactionHandler) SetCommentReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	userID := r.Context().Value(config.UserID)
	reaction := entity.CommentReaction{
		UserID: userID.(uint64),
	}

	if err := json.NewDecoder(r.Body).Decode(&reaction); err != nil {
		http.Error(w, customErr.InvalidData, http.StatusBadRequest)
		return
	}

	if err := rct.service.SetCommentReaction(r.Context(), reaction); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}
