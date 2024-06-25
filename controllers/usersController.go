package controllers

import (
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	user, _ := utils.ExtractUser(c)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
