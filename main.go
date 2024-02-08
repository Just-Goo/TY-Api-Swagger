package main

import (
	"log"

	_ "github.com/TY-Api-Swagger/docs"
	"github.com/TY-Api-Swagger/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Documentation
// @version 1
// @Description Test description

// @contact.name Just Goo
// @contact.url https://github.com/Just-Goo
// @contact.email sample@gmail.com

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	user := v1.Group("/users")
	{
		user.GET("/", handlers.GetUsers)
		user.POST("/", handlers.CreateUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

