package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itismrx/jwt-auth/controllers"
	"github.com/itismrx/jwt-auth/initalizers"
	"github.com/itismrx/jwt-auth/middleware"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDb()
	initalizers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
