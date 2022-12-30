package routes

import (
	"lctzz540/controllers"

	"github.com/gin-gonic/gin"
)

func QuestionRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/question/getownquestion", controllers.GetOwnQuestions())
	incomingRoutes.POST("/question/addownquestion", controllers.AddOwnQuestions())
}
