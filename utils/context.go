package utils

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ExtractUser(c *gin.Context) (models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return models.User{}, fmt.Errorf("user not found")
	}

	userStruct, ok := user.(models.User)
	if !ok {
		return models.User{}, fmt.Errorf("user type assertion failed")
	}

	return userStruct, nil
}
