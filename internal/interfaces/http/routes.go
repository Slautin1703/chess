package http

import (
	"awesomeProject/internal/interfaces/http/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/logout", handlers.LogoutHandler)
	mux.HandleFunc("/signup", handlers.SignUpHandler)
	mux.HandleFunc("/user", AuthMiddleware(handlers.GetUserHandler))
	mux.HandleFunc("/createGame", AuthMiddleware(handlers.CreateGameHandler))

	return mux
}
