package initializers

import (
	"awesomeProject/models"
)

func SyncDatabase() {
	// Create table for `User`
	err := DB.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
