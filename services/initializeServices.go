package services

import "awesomeProject/initializers"

var GameService *IGameService

func Initialize() {
	GameService = NewGameService(initializers.DB)
}
