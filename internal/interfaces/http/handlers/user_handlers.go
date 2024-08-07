package handlers

import (
	"encoding/json"
	"github.com/Slautin1703/games/internal/domain/models"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
