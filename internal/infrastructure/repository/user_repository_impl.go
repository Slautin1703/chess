package repository

import (
	"database/sql"
	"github.com/Slautin1703/games/internal/domain/interfaces"
	"github.com/Slautin1703/games/internal/domain/models"
	"github.com/Slautin1703/games/internal/infrastructure/db"
)

type UserRepositoryImpl struct{}

func NewUserRepository() interfaces.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return nil, nil
		}
		// Other errors
		return nil, err
	}
	return &user, nil
}

// Save inserts a new user into the database.
func (r *UserRepositoryImpl) Save(user *models.User) error {
	_, err := db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}
