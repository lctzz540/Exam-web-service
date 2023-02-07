package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lctzz540/Exam-web-service/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))
	routes.AuthRoutes(router)
	routes.QuestionRoutes(router)
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "home")
	})

	router.Run("0.0.0.0:8080")
}
