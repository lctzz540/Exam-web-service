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
	router.Use(cors.Default())
	routes.AuthRoutes(router)
	routes.QuestionRoutes(router)
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "home")
	})

	router.Run("localhost:8080")
}
