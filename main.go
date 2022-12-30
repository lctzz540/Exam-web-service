package main

import (
	"lctzz540/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.QuestionRoutes(router)
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "home")
	})

	router.Run("localhost:8080")
}
