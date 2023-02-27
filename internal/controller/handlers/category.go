package handlers

import (
	"log"
	"net/http"
	"strconv"

	"forum/internal/service"
	"forum/internal/tool/config"
	"forum/internal/tool/customErr"
	"forum/pkg/gayson"
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
	categories, err := c.service.GetAllCategories(r.Context())
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	// if err := json.NewEncoder(w).Encode(categories); err != nil {
	// 	http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
	// 	return
	// }
	gayson.SendJSON(w, categories)
}

func (c *categoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.ParseUint(r.Context().Value(config.CategoryID).(string), 10, 64)
	if err != nil {
		http.Error(w, customErr.Bruhhh, http.StatusBadRequest)
		return
	}

	category, err := c.service.GetCategoryByID(r.Context(), categoryID)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	// if err := json.NewEncoder(w).Encode(category); err != nil {
	// 	http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
	// 	return
	// }
	gayson.SendJSON(w, category)
}
