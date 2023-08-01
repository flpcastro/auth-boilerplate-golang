package main

import (
	"v1/src/controllers"
	"v1/src/initializers"
	"v1/src/middlewares"

	docs "v1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Template Golang Kukac"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"HTTP"}

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
