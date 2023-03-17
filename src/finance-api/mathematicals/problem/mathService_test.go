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

var service *MathService[int]
var criteria Criteria[int]
var criteria2 Criteria[int]
var criteria3 Criteria[int]
var criteria4 Criteria[int]
var criteria5 Criteria[int]
var saveRquest SaveAnswersRequest
var questions entity.Questions

func init() {
	criteria = Criteria[int]{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQuestFirst,

		Category: CategoryPlus,
	}

	criteria2 = Criteria[int]{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQuestSecond,

		Category: CategoryPlus,

		Type: TypePlainExpression,
	}

	criteria3 = Criteria[int]{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQeustLast,

		Category: CategoryMinus,
		Type:     TypePlainApplication,
	}

	criteria4 = Criteria[int]{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQeustLast,

		Category: CategoryMinus,
		Type:     TypeAppleApplication,
	}

	criteria5 = Criteria[int]{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: KindQeustLast,

		Category: CategoryMinus,
		Type:     TypeTemplateApplication,
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
	service = NewMathService(NewTwoGenerationService[int](), mocks.NewQuestionRepo(t))

	assert.NotNil(t, service)
	assert.NotNil(t, service.genService)

	assert.NotNil(t, service.genService.TwoPlusService)
	assert.NotNil(t, service.genService.TwoMinusService)
}

func TestGenerateProblems(t *testing.T) {
	service = NewMathService(NewTwoGenerationService[int](), mocks.NewQuestionRepo(t))

	results := service.GenerateProblems(criteria)

	assert.NotNil(t, results)
	assert.NotEmpty(t, results.QuestionId)
	assert.Equal(t, len(results.Questions), criteria.Quantity)
}

func TestSaveResults_QueryError(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	service = NewMathService(NewTwoGenerationService[int](), mockRepo)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveQuestions(mock.Anything).Return(errors.New("any"))

	results := service.SaveResults(&saveRquest, "any")
	assert.NotNil(t, results)
}

func TestSaveResults_SaveError(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	service = NewMathService(NewTwoGenerationService[int](), mockRepo)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(&questions)
	mockRepo.EXPECT().SaveAnswers(mock.AnythingOfType("*entity.Answers")).Return(errors.New("any")).Times(1)

	results := service.SaveResults(&saveRquest, "any")
	assert.NotNil(t, results)
}

func TestSaveResults(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	service = NewMathService(NewTwoGenerationService[int](), mockRepo)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveQuestions(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveAnswers(mock.AnythingOfType("*entity.Answers")).Return(nil)

	results := service.SaveResults(&saveRquest, "any")
	assert.Nil(t, results)
}

func TestGenerateProblemsMulti(t *testing.T) {
	service = NewMathService(NewTwoGenerationService[int](), mocks.NewQuestionRepo(t))
	totalQuantity := criteria.Quantity + criteria2.Quantity + criteria3.Quantity + criteria4.Quantity + criteria5.Quantity

	results := service.GenerateProblems(criteria, criteria2, criteria3, criteria4, criteria5)

	assert.NotNil(t, results)
	assert.NotEmpty(t, results.QuestionId)
	assert.Equal(t, len(results.Questions), totalQuantity)
}

func TestGenerateFeeds(t *testing.T) {
	service = NewMathService(NewTwoGenerationService[int](), mocks.NewQuestionRepo(t))
	totalQuantity := criteria.Quantity + criteria2.Quantity + criteria3.Quantity + criteria4.Quantity + criteria5.Quantity

	results := service.GenerateFeeds(criteria, criteria2, criteria3, criteria4, criteria5)

	assert.NotNil(t, results)
	assert.Equal(t, len(results.Answers), totalQuantity)
	assert.Equal(t, len(results.FullText), totalQuantity)
	assert.Equal(t, len(results.Questions), totalQuantity)
}
