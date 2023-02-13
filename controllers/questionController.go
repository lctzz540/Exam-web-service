package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lctzz540/Exam-web-service/database"
	"github.com/lctzz540/Exam-web-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var questionCollection *mongo.Collection = database.OpenCollection(database.Client, "questions")

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func GetOwnQuestions() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var foundQuestion []models.Question

		defer cancel()
		userEmail, exited := c.Get("contextEmail")
		if !exited {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "error while get context"})
			return
		}

		cur, err := questionCollection.Find(ctx, bson.M{"owner": userEmail})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No question is found"})
			return
		}
		for cur.Next(context.TODO()) {
			var question models.Question
			err := cur.Decode(&question)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when get questions"})
				return
			}
			foundQuestion = append(foundQuestion, question)
		}

		c.JSON(http.StatusOK, foundQuestion)
	}
}

func AddOwnQuestions() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var question models.Question

		defer cancel()
		if err := c.ShouldBindJSON(&question); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer cancel()
		validationErr := validate.Struct(question)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		var filter = bson.M{"questionText": question.QuestionText}
		var foundQuestion models.Question
		existedErr := questionCollection.FindOne(ctx, filter).Decode(&foundQuestion)
		if existedErr == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Question already exists"})
			return
		}

		question.QuestionID = primitive.NewObjectID()
		question.Create_at = time.Now()
		question.Lasted_update = time.Now()

		value, ok := c.Value("contextEmail").(string)
		if !ok {

		}
		question.Owner = &value

		defer cancel()
		_, err := questionCollection.InsertOne(ctx, question)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "success"})
	}
}

func AddManyOwnQuestions() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var questions []models.Question
		if err := c.ShouldBindJSON(&questions); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var successCount int
		var failCount int
		for _, question := range questions {
			validationErr := validate.Struct(question)
			if validationErr != nil {
				failCount++
				continue
			}

			var filter = bson.M{"questionText": question.QuestionText}
			var foundQuestion models.Question
			existedErr := questionCollection.FindOne(ctx, filter).Decode(&foundQuestion)
			if existedErr == nil {
				failCount++
				continue
			}

			question.QuestionID = primitive.NewObjectID()
			question.Create_at = time.Now()
			question.Lasted_update = time.Now()

			value, ok := c.Value("contextEmail").(string)
			if !ok {
				failCount++
				continue
			}
			question.Owner = &value

			_, err := questionCollection.InsertOne(ctx, question)
			if err != nil {
				failCount++
				continue
			}

			successCount++
		}

		c.JSON(http.StatusCreated, gin.H{"status": "success", "successCount": successCount, "failCount": failCount})
	}
}
func GetQuestionByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			QuestionID string `json:"questionID"`
		}

		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		questionID, err := primitive.ObjectIDFromHex(requestBody.QuestionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
			return
		}

		var question models.Question
		err = questionCollection.FindOne(context.TODO(), bson.M{"_id": questionID}).Decode(&question)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
			return
		}

		c.JSON(http.StatusOK, question)
	}
}
