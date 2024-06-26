package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}

type UserClaim struct {
	Sub string `json:"sub"`
	jwt.RegisteredClaims
}
