package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	QuestionID    primitive.ObjectID `bson:"_id"`
	QuestionText  *string            `bson:"questionText" json:"questionText" validate:"required,min=2,max=100"`
	AnswerOptions []AnswerOption     `bson:"answerOptions" json:"answerOptions" validate:"required"`
	Describle     *string            `json:"describle"`
	Subject       *string            `json:"subject" validate:"required"`
	Owner         *string            `json:"owner"`
	Create_at     time.Time
	Lasted_update time.Time
}
