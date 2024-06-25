package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("CONNECTION_STRING")
	DBInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = DBInstance
	if err != nil {
		panic("failed to connect database")
	}
}
