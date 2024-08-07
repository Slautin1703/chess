package interfaces

import "github.com/Slautin1703/games/internal/domain/models"

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Save(user *models.User) error
}
