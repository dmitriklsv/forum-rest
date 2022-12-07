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

func (c *categoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	categories, err := c.service.GetAllCategories(r.Context())
	if err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (c *categoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var category entity.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, errors.Bruhhh, http.StatusBadRequest)
		return
	}

	category, err := c.service.GetCategoryByID(r.Context(), category.ID)
	if err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, errors.InvalidContract, http.StatusInternalServerError)
		return
	}
}
