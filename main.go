package main

import (
	"jwtgen/handlers"
	"jwtgen/initializers"
	"jwtgen/middleware"
	"jwtgen/router"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
	initializers.LoadKeys()
}

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.LogIn)
	r.GET("/validate", middleware.RequestAuth, handlers.Validate)
	r.POST("/getUserData", middleware.RequestAuth, handlers.GetUserData)

	r.POST("/model", middleware.RequestAuth, router.GraphQLHandler())
	r.GET("/", router.PlaygroundHandler())
	r.Run(":" + port)
}
