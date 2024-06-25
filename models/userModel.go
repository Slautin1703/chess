package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type UserClaim struct {
	Sub string `json:"sub"`
	jwt.RegisteredClaims
}
