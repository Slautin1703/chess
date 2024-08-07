package http

import (
	"net/http"
	"user-service/internal/interfaces/http/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/logout", handlers.LogoutHandler)
	mux.HandleFunc("/signup", handlers.SignUpHandler)
	mux.HandleFunc("/user", AuthMiddleware(handlers.GetUserHandler))

	return mux
}
