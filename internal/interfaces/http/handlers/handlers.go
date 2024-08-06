package handlers

import (
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"id":    1,
		"email": "test@gmail.com",
	}

	json.NewEncoder(w).Encode(response)
}
