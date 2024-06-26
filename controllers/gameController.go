package controllers

import (
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createGameRequest struct {
	User1ID uint `json:"user1Id"`
	User2ID uint `json:"user2Id"`
}

func CreateGame(c *gin.Context) {
	var req createGameRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	game, err := services.GameService.CreateNewGame(req.User1ID, req.User2ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"players": game.Players})
}

//
//func GetGameWithPlayers(gameID uint) (*models.Game, error) {
//	game := models.Game{}
//	// Preload the players association and fetch the game by ID
//	if err := initializers.DB.Preload("Players").First(&game, gameID).Error; err != nil {
//		return nil, err
//	}
//	return &game, nil
//}

func GetGames(c *gin.Context) {
	// Dummy implementation for extracting user
	// user, _ := utils.ExtractUser(c)

	c.JSON(http.StatusOK, gin.H{"message": "GetGames endpoint called"})
}
