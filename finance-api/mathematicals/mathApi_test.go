package mathematicals

import (
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"

	"github.com/stretchr/testify/assert"
)

var service *MathApi
var criteria problem.Criteria
var criteria2 problem.Criteria
var criteria3 problem.Criteria

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

	mathService := problem.NewMathService(problem.NewTwoGenerationService())
	service = ProvideMathApi(mathService)
}

func TestProvideMathApi(t *testing.T) {
	assert.NotNil(t, service)
	assert.NotNil(t, service.mathService)
}
func TestGetQuestionsFailed(t *testing.T) {
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
	service.GetQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetQuestionsInvalid(t *testing.T) {
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
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: criteria})
	service.GetQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetAllQuestionsFailed(t *testing.T) {
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
	service.GetAllQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetAllQuestionsOk(t *testing.T) {
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: []problem.Criteria{criteria, criteria2, criteria3}})
	service.GetAllQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetAllQuestionFeedsFailed(t *testing.T) {
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
	service.GetAllQuestionFeeds(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetAllQuestionFeedsOk(t *testing.T) {
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: []problem.Criteria{criteria, criteria2, criteria3}})
	service.GetAllQuestionFeeds(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}
