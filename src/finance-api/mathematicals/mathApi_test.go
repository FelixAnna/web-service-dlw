package mathematicals

import (
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/common/snowflake"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"

	"github.com/FelixAnna/web-service-dlw/finance-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var service *MathApi
var criteria problem.Criteria
var criteria2 problem.Criteria
var criteria3 problem.Criteria
var criteria4 problem.Criteria
var saveRquest problem.SaveAnswersRequest
var questions entity.Questions

func init() {
	criteria = problem.Criteria{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: problem.KindQuestFirst,

		Category: problem.CategoryPlus,
	}

	criteria2 = problem.Criteria{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: problem.KindQuestSecond,

		Category: problem.CategoryPlus,

		Type: problem.TypePlainExpression,
	}

	criteria3 = problem.Criteria{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Kind: problem.KindQeustLast,

		Category: problem.CategoryMinus,
		Type:     problem.TypePlainApplication,
	}

	criteria4 = problem.Criteria{
		Min:      1,
		Max:      100,
		Quantity: 10,

		Kind: problem.KindQeustLast,

		Category: problem.CategoryMultiply,
		Type:     problem.TypePlainApplication,
	}

	saveRquest = problem.SaveAnswersRequest{
		Results: []problem.QuestionAnswerItem{
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

func TestProvideMathApi(t *testing.T) {
	initialService(t)

	assert.NotNil(t, service)
	assert.NotNil(t, service.mathService)
}
func TestGetQuestionsFailed(t *testing.T) {
	initialService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
	service.GetQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetQuestionsInvalid(t *testing.T) {
	initialService(t)

	criteriaInvalid := problem.Criteria{
		Kind: 255,
	}
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: criteriaInvalid})
	service.GetQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetQuestionsOk(t *testing.T) {
	initialService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: criteria})
	service.GetQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetAllQuestionsFailed(t *testing.T) {
	initialService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
	service.GetAllQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetAllQuestionsOk(t *testing.T) {
	initialService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: []problem.Criteria{criteria, criteria2, criteria3, criteria4}})
	service.GetAllQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetAllQuestionFeedsFailed(t *testing.T) {
	initialService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
	service.GetAllQuestionFeeds(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetAllQuestionFeedsOk(t *testing.T) {
	initialService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: []problem.Criteria{criteria, criteria2, criteria3, criteria4}})
	service.GetAllQuestionFeeds(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestSaveResultsFailed(t *testing.T) {
	initialService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
	service.SaveResults(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestSaveResultsInvalid(t *testing.T) {
	initialService(t)

	request := problem.SaveAnswersRequest{
		Score:   100,
		Results: []problem.QuestionAnswerItem{},
	}
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: request})
	service.SaveResults(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestSaveResultsError(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	mathService := problem.NewMathService(problem.NewTwoGenerationService(), mockRepo)
	service = ProvideMathApi(mathService)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveQuestions(mock.Anything).Return(errors.New("any"))

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: saveRquest})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	service.SaveResults(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
}

func TestSaveResultsOk(t *testing.T) {
	mockRepo := mocks.NewQuestionRepo(t)
	mathService := problem.NewMathService(problem.NewTwoGenerationService(), mockRepo)
	service = ProvideMathApi(mathService)

	mockRepo.EXPECT().GetQuestion(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveQuestions(mock.Anything).Return(nil)
	mockRepo.EXPECT().SaveAnswers(mock.AnythingOfType("*entity.Answers")).Return(nil)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: saveRquest})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	service.SaveResults(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func initialService(t *testing.T) {
	mathService := problem.NewMathService(problem.NewTwoGenerationService(), mocks.NewQuestionRepo(t))
	service = ProvideMathApi(mathService)
}
