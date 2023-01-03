package routes

import (
	"lctzz540/controllers"
	"lctzz540/middlewares"

	"github.com/gin-gonic/gin"
)

func QuestionRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/question/getownquestion", middlewares.JWTMiddleware(), controllers.GetOwnQuestions())
	incomingRoutes.POST("/question/addownquestion", middlewares.JWTMiddleware(), controllers.AddOwnQuestions())
}
