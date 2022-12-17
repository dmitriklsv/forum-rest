package gayson

import (
	"encoding/json"
	"net/http"

	"forum/internal/tool/customErr"
)

func SendJSON(w http.ResponseWriter, v ...any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
	}
}
