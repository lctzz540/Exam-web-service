package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	QuestionID    primitive.ObjectID `bson:"_id"`
	QuestionText  *string            `bson:"questionText" json:"questionText" validate:"required"`
	AnswerOptions []AnswerOption     `bson:"answerOptions" json:"answerOptions" validate:"required"`
	Describle     *string            `json:"describle"`
	Image         *string            `bson:"image" json:"image"`
	Subject       *string            `json:"subject" validate:"required"`
	Owner         *string            `json:"owner"`
	Create_at     time.Time
	Lasted_update time.Time
}
