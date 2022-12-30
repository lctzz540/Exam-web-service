package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	QuestionID    primitive.ObjectID `bson:"_id"`
	QuestionText  *string            `json:"questionText" validate:"required,min=2,max=100"`
	AnswerOptions []AnswerOption     `json:"answerOptions" validate:"required"`
	Describle     *string            `json:"describle"`
	Subject       *string            `json:"Subject" validate:"required"`
	Owner         *string
	Create_at     time.Time
	Lasted_update time.Time
}
