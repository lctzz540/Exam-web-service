package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lctzz540/Exam-web-service/controllers"
	"github.com/lctzz540/Exam-web-service/middlewares"
)

func QuestionRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/question/getownquestion", middlewares.JWTMiddleware(), controllers.GetOwnQuestions())
	incomingRoutes.POST("/question/addownquestion", middlewares.JWTMiddleware(), controllers.AddOwnQuestions())
	incomingRoutes.POST("/question/addmanyownquestion", middlewares.JWTMiddleware(), controllers.AddManyOwnQuestions())
	incomingRoutes.POST("/question/getquestionbyid", middlewares.JWTMiddleware(), controllers.GetQuestionByID())
}
