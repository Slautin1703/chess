package initializers

import (
	"awesomeProject/models"
)

func SyncDatabase() {
	// Create table for `User`
	err := DB.Migrator().AutoMigrate(&models.User{}, &models.Game{}, &models.Move{}, &models.Figure{})
	if err != nil {
		return
	}
}
