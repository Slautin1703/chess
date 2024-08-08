package http

import (
	"game-service/internal/delivery/http/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	//mux.HandleFunc("/createGame", AuthMiddleware(handlers.CreateGameHandler))
	mux.HandleFunc("/createGame", handlers.CreateGameHandler)

	return mux
}
