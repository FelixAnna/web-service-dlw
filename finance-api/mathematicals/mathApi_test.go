package mathematicals

import (
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/stretchr/testify/assert"
)

var service *MathApi
var criteria Criteria

func init() {
	criteria = Criteria{
		Min:      100,
		Max:      200,
		Quantity: 10,

		Category: "+",
	}

	mathService := NewMathService()
	service = ProvideMathApi(mathService)
}

func TestProvideMathApi(t *testing.T) {
	assert.NotNil(t, service)
	assert.NotNil(t, service.mathService)
	assert.NotNil(t, service.mathService.TwoPlusService)
	assert.NotNil(t, service.mathService.TwoMinusService)
}
func TestGetQuestionsFailed(t *testing.T) {
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: "invalid"})
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
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: []Criteria{criteria, criteria}})
	service.GetAllQuestions(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}
