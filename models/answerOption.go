package models

type AnswerOption struct {
	AnswerText *string `bson:"answerText" json:"answerText" validate:"required"`
	IsCorrect  bool    `bson:"isCorrect" json:"isCorrect" validate:"required"`
}
