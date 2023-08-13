package repositories

import "github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/entity"

type QuestionRepo interface {
	GetQuestion(id string) *entity.Questions
	SaveQuestions(questions *entity.Questions) error
	SaveAnswers(answers *entity.Answers) error
}
