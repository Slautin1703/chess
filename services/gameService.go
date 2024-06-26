package services

import (
	"awesomeProject/models"
	"fmt"
	"gorm.io/gorm"
)

const (
	errGameNotFound  = "game not found"
	errUserNotFound  = "user not found"
	errCreateGame    = "failed to create game"
	errAddGamePlayer = "failed to add game player"
	errCommitTx      = "failed to commit transaction"
)

type IGameService struct {
	DB *gorm.DB
}

func NewGameService(db *gorm.DB) *IGameService {
	return &IGameService{DB: db}
}

func (s *IGameService) GetGameWithPlayers(gameID uint) (*models.Game, error) {
	var game models.Game
	if err := s.DB.Preload("Players").First(&game, gameID).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", errGameNotFound, err)
	}
	return &game, nil
}

func (s *IGameService) CreateNewGame(user1ID, user2ID uint) (*models.Game, error) {
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	game := models.Game{}
	if err := tx.Create(&game).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("%s: %w", errCreateGame, err)
	}

	player1, err := s.AddGamePlayer(tx, &game, user1ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	player2, err := s.AddGamePlayer(tx, &game, user2ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	fmt.Println("Player 1 Email:", player1.Email)
	fmt.Println("Player 2 Email:", player2.Email)

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("%s: %w", errCommitTx, err)
	}

	return &game, nil
}

func (s *IGameService) FindGameById(tx *gorm.DB, gameID uint) (*models.Game, error) {
	var game models.Game
	if err := tx.First(&game, gameID).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", errGameNotFound, err)
	}

	return &game, nil
}

func (s *IGameService) AddGamePlayer(tx *gorm.DB, game *models.Game, userID uint) (*models.User, error) {
	var user models.User
	if err := tx.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", errUserNotFound, err)
	}

	if err := tx.Model(&game).Association("Players").Append(&user); err != nil {
		return nil, fmt.Errorf("%s: %w", errAddGamePlayer, err)
	}

	return &user, nil
}
