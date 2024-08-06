package http

import (
	"awesomeProject/internal/domain/models"
	"awesomeProject/internal/infrastructure/db"
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type UserClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Authorization")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println(err)
			return
		}

		claims, ok := token.Claims.(*UserClaim)
		if !ok || !token.Valid || claims.RegisteredClaims.ExpiresAt.Unix() < time.Now().Unix() {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println("Invalid or expired token")
			return
		}

		var user models.User
		err = db.DB.QueryRow("SELECT id, email FROM users WHERE email = $1", claims.Email).Scan(&user.ID, &user.Email)
		if err == sql.ErrNoRows {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the user in the request context
		r = r.WithContext(context.WithValue(r.Context(), "user", user))

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

// CORSMiddleware handles CORS settings
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
