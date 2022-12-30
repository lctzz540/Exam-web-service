package models

type AnswerOption struct {
	AnswerText *string `json:"answerText" validate:"required"`
	IsCorrect  bool    `json:"isCorrect" validate:"required"`
}
