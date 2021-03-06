package problem

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/snowflake"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var service *MathService
var criteria Criteria
var criteria2 Criteria
var criteria3 Criteria
var saveRquest SaveAnswersRequest
var questions entity.Questions

func init() {
	criteria = Criteria{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQuestFirst,

		Category: CategoryPlus,
	}

	criteria2 = Criteria{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQuestSecond,

		Category: CategoryPlus,

		Type: TypePlainExpression,
	}

	criteria3 = Criteria{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQeustLast,

		Category: CategoryMinus,
		Type:     TypePlainApplication,
	}

	saveRquest = SaveAnswersRequest{
		Results: []QuestionAnswerItem{
			{
				Index: 1,

				Question: "any",
				Answer:   "any",

				Category: 1,
				Kind:     1,
				Type:     1,

				UserAnswer: "any",
			},
		},
		QuestionId: "anyId",
		Score:      100,
	}

	questions = entity.Questions{
		Id: "anyId",
		Questions: []entity.QuestionItem{
			{
				Index: 1,

				Question: "any",
				Answer:   "any",

				Category: 1,
				Kind:     1,
				Type:     1,
			},
		},
		CreatedTime: time.Now().UTC().Unix(),
	}

	os.Setenv("DLW_NODE_NO", "1023")
	snowflake.InitSnowflake()
}

func TestNewMathService(t *testing.T) {
	service = NewMathService(NewTwoGenerationService(), mocks.NewQuestionRepo(t))

	assert.NotNil(t, service)
	assert.NotNil(t, service.genService)

	assert.NotNil(t, service.genService.TwoPlusService)
	assert.NotNil(t, service.genService.TwoMinusService)
}

func TestGenerateProblems(t *testing.T) {
	service = NewMathService(NewTwoGenerationService(), mocks.NewQuestionRepo(t))

	results := service.GenerateProblems(criteria)

	assert.NotNil(t, results)
	assert.NotEmpty(t, results.QuestionId)
	assert.Equal(t, len(results.Questions), criteria.Quantity)
}

func TestSaveResults_QueryError(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	service = NewMathService(NewTwoGenerationService(), mockRepo)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveQuestions(mock.Anything).Return(errors.New("any"))

	results := service.SaveResults(&saveRquest, "any")
	assert.NotNil(t, results)
}

func TestSaveResults_SaveError(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	service = NewMathService(NewTwoGenerationService(), mockRepo)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(&questions)
	mockRepo.EXPECT().SaveAnswers(mock.AnythingOfType("*entity.Answers")).Return(errors.New("any")).Times(1)

	results := service.SaveResults(&saveRquest, "any")
	assert.NotNil(t, results)
}

func TestSaveResults(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	service = NewMathService(NewTwoGenerationService(), mockRepo)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveQuestions(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveAnswers(mock.AnythingOfType("*entity.Answers")).Return(nil)

	results := service.SaveResults(&saveRquest, "any")
	assert.Nil(t, results)
}

func TestGenerateProblemsMulti(t *testing.T) {
	service = NewMathService(NewTwoGenerationService(), mocks.NewQuestionRepo(t))

	results := service.GenerateProblems(criteria, criteria2, criteria3)

	assert.NotNil(t, results)
	assert.NotEmpty(t, results.QuestionId)
	assert.Equal(t, len(results.Questions), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
}

func TestGenerateFeeds(t *testing.T) {
	service = NewMathService(NewTwoGenerationService(), mocks.NewQuestionRepo(t))

	results := service.GenerateFeeds(criteria, criteria2, criteria3)

	assert.NotNil(t, results)
	assert.Equal(t, len(results.Answers), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
	assert.Equal(t, len(results.FullText), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
	assert.Equal(t, len(results.Questions), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
}
