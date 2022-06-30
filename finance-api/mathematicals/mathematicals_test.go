package mathematicals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mathService *MathService

func init() {
	mathService = NewMathService()
}

func TestNewMathService(t *testing.T) {
	assert.NotNil(t, mathService)
	assert.NotNil(t, mathService.TwoPlusService)
	assert.NotNil(t, mathService.TwoMinusService)
}

func TestGetQuestions(t *testing.T) {
	criteria := &Criteria{
		Min: 100,
		Max: 200,

		Quantity: 5,

		Category: CategoryPlus,
	}

	problems := mathService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), criteria.Quantity)
	for _, problem := range problems {
		assert.True(t, problem.A >= criteria.Min && problem.A < criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B < criteria.Max)
		assert.Equal(t, problem.Op, '+')
		assert.True(t, problem.A+problem.B == problem.C)
	}
}

func TestGetQuestionsWithinRange(t *testing.T) {
	criteria := &Criteria{
		Min: 100,
		Max: 200,

		Range: &Range{
			Max: 230,
		},

		Quantity: 5,

		Category: CategoryPlus,
	}

	problems := mathService.GenerateProblems(criteria)

	fmt.Println(problems)
	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), criteria.Quantity)
	for _, problem := range problems {
		assert.True(t, problem.A >= criteria.Min && problem.A < criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B < criteria.Max)
		assert.Equal(t, problem.Op, '+')
		assert.True(t, problem.C <= criteria.Range.Max)
		assert.True(t, problem.A+problem.B == problem.C)
	}
}

func TestGenerateProblemsMinus(t *testing.T) {
	criteria := &Criteria{
		Min: 10,
		Max: 100,

		Range: &Range{
			Min: 0,
			Max: 20,
		},

		Quantity: 5,

		Category: CategoryMinus,
	}

	problems := mathService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), criteria.Quantity)
	for _, problem := range problems {
		fmt.Println(problem)
		assert.True(t, problem.A >= criteria.Min && problem.A < criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B < criteria.Max)
		assert.Equal(t, problem.Op, '-')
		assert.True(t, problem.A-problem.B == problem.C)
		assert.True(t, problem.A >= problem.B)
	}
}
func TestGenerateProblemsMinusPos(t *testing.T) {
	criteria := &Criteria{
		Min: 10,
		Max: 100,

		Range: &Range{
			Min: 0,
			Max: 30,
		},
		Category: CategoryMinus,
	}

	problems := mathService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), 10)
	for _, problem := range problems {
		assert.True(t, problem.A >= criteria.Min && problem.A < criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B < criteria.Max)
		assert.Equal(t, problem.Op, '-')
		assert.True(t, problem.A >= problem.B)
		assert.True(t, problem.A-problem.B == problem.C)
	}
}

func TestGenerateProblemsMinusImpossible(t *testing.T) {
	criteria := &Criteria{
		Min: 10,
		Max: 100,

		Range: &Range{
			Min: 100,
			Max: 300,
		},
		Category: CategoryMinus,
	}

	problems := mathService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), 0)
}
