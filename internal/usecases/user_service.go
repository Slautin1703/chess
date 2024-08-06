package usecases

import (
	"awesomeProject/internal/domain/interfaces"
	"awesomeProject/internal/domain/models"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil || user.Password != password {
		return nil, err
	}
	return user, nil
}
