package interfaces

import "awesomeProject/internal/domain/models"

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Save(user *models.User) error
}
