package main

import (
	"gihub.com/mthsnts/go-JWT-API/Middleware"
	"gihub.com/mthsnts/go-JWT-API/controllers"
	"gihub.com/mthsnts/go-JWT-API/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", Middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
