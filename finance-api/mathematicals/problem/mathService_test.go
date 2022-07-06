package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var service *MathService
var criteria Criteria
var criteria2 Criteria
var criteria3 Criteria

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

	service = NewMathService(NewTwoGenerationService())
}

func TestNewMathService(t *testing.T) {
	assert.NotNil(t, service)
	assert.NotNil(t, service.genService)

	assert.NotNil(t, service.genService.TwoPlusService)
	assert.NotNil(t, service.genService.TwoMinusService)
}

func TestGenerateProblems(t *testing.T) {
	results := service.GenerateProblems(criteria)

	assert.NotNil(t, results)
	assert.Equal(t, len(results), criteria.Quantity)
}

func TestGenerateProblemsMulti(t *testing.T) {
	results := service.GenerateProblems(criteria, criteria2, criteria3)

	assert.NotNil(t, results)
	assert.Equal(t, len(results), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
}

func TestGenerateFeeds(t *testing.T) {
	results := service.GenerateFeeds(criteria, criteria2, criteria3)

	assert.NotNil(t, results)
	assert.Equal(t, len(results.Answers), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
	assert.Equal(t, len(results.FullText), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
	assert.Equal(t, len(results.Questions), criteria.Quantity+criteria2.Quantity+criteria3.Quantity)
}
