package controllers

import (
	"context"
	"lctzz540/database"
	"lctzz540/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
		var question models.Question
		var foundQuestion []models.Question

		defer cancel()
		if err := c.BindJSON(&question); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if question.Owner == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "required owner"})
			return
		}

		cur, err := questionCollection.Find(ctx, bson.M{"owner": question.Owner})
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

		question.QuestionID = primitive.NewObjectID()
		question.Create_at = time.Now()
		question.Lasted_update = time.Now()

		defer cancel()
		_, err := questionCollection.InsertOne(ctx, question)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "success"})
	}
}
