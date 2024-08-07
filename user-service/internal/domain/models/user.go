package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserClaim struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
