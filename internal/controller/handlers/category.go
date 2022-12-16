package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"forum/internal/service"
	"forum/internal/tool/customErr"
)

type categoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *categoryHandler {
	log.Println("| | category handler is done!")
	return &categoryHandler{
		service: service,
	}
}

func (c *categoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	categories, err := c.service.GetAllCategories(r.Context())
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (c *categoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	categoryID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
		return
	}

	category, err := c.service.GetCategoryByID(r.Context(), categoryID)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}
