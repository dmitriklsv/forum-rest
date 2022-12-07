package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"forum/internal/entity"
	"forum/internal/service"
	"forum/internal/tool/errors"
)

type categoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	log.Println("| | category handler is done!")
	return &categoryHandler{
		service: service,
	}
}

func (c *categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var categories []entity.Category
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
		http.Error(w, errors.InvalidData, http.StatusBadRequest)
		return
	}

	catIDs, err := c.service.CreateCategory(r.Context(), categories)
	if err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(catIDs)
}
