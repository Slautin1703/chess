package main

import (
	"awesomeProject/controllers"
	"awesomeProject/initializers"
	"awesomeProject/middlewares"
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	services.Initialize()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello! Take it, my friend, you need this")
	})

	r.POST("/signup", controllers.SignUp)

	r.POST("/login", controllers.Login)

	r.POST("/user", middlewares.RequireAuth, controllers.GetUser)

	r.POST("/gamesHistory", middlewares.RequireAuth, controllers.GetGames)

	r.POST("/createGame", middlewares.RequireAuth, controllers.CreateGame)

	r.Run()
}
