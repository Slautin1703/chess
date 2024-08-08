package handlers

import (
	"encoding/json"
	"game-service/internal/infrastructure/db"
	"net/http"
	"time"
)

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	//_, ok := r.Context().Value("user").(models.User)
	//if !ok {
	//	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	//	return
	//}

	var body struct {
		GameType string `json:"type"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	if body.GameType != "fool" && body.GameType != "chess" {
		http.Error(w, "There are just two games available: chess and fool", http.StatusBadRequest)
		return
	}
	db.DB.QueryRow(
		"INSERT INTO games (game_type, start_time) VALUES ($1, $2) RETURNING id",
		body.GameType,
		time.Now(),
	)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Game created successfully"))
}
