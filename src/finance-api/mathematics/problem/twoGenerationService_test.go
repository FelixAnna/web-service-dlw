package problem

import (
	"fmt"
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mocks"
	"github.com/stretchr/testify/assert"
)

var twoGenService *TwoGenerationService

func init() {
	twoGenService = NewTwoGenerationService()
}

func TestTwoGenerationService(t *testing.T) {
	assert.NotNil(t, twoGenService)
	assert.NotNil(t, twoGenService.TwoPlusService)
	assert.NotNil(t, twoGenService.TwoMinusService)
}

func TestGetQuestions(t *testing.T) {
	criteria := &Criteria{
		Min: 100,
		Max: 200,

		Quantity: 5,

		Category: CategoryPlus,
	}

	problems := twoGenService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), criteria.Quantity)
	for _, problem := range problems {
		assert.True(t, problem.A >= criteria.Min && problem.A <= criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B <= criteria.Max)
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

	problems := twoGenService.GenerateProblems(criteria)

	fmt.Println(problems)
	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), criteria.Quantity)
	for _, problem := range problems {
		assert.True(t, problem.A >= criteria.Min && problem.A <= criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B <= criteria.Max)
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

	problems := twoGenService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), criteria.Quantity)
	for _, problem := range problems {
		fmt.Println(problem)
		assert.True(t, problem.A >= criteria.Min && problem.A <= criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B <= criteria.Max)
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

	problems := twoGenService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), 10)
	for _, problem := range problems {
		assert.True(t, problem.A >= criteria.Min && problem.A <= criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B <= criteria.Max)
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

	problems := twoGenService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), 0)
}

func TestGenerateProblemsNoDuplocates(t *testing.T) {
	criteria := &Criteria{
		Min: 10,
		Max: 100,

		Quantity: 10,

		Range: &Range{
			Min: 0,
			Max: 30,
		},
		Category: CategoryMinus,
	}

	problemService := mocks.NewProblemService(t)

	problems := []entity.Problem{}

	problemService.EXPECT().GenerateProblem(criteria.Min, criteria.Max).Return(
		&entity.Problem{
			A:  15,
			B:  12,
			C:  3,
			Op: '-',
		}).Times(MaxGenerateTimes)

	GenerateProblems(criteria, problemService, &problems)

	assert.Equal(t, len(problems), 1)
}

func TestGenerateProblemsMultiply(t *testing.T) {
	criteria := &Criteria{
		Min: 1,
		Max: 10,

		Range: &Range{
			Min: 1,
			Max: 50,
		},
		Category: CategoryMultiply,
	}

	problems := twoGenService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), 10)
	for _, problem := range problems {
		assert.True(t, problem.A >= criteria.Min && problem.A <= criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B <= criteria.Max)
		assert.Equal(t, problem.Op, '*')
		assert.True(t, problem.A*problem.B == problem.C)
	}
}

func TestGenerateProblemsDivide(t *testing.T) {
	criteria := &Criteria{
		Min: 1,
		Max: 100,

		Range: &Range{
			Min: 1,
			Max: 10,
		},
		Category: CategoryDivide,
	}

	problems := twoGenService.GenerateProblems(criteria)

	assert.NotNil(t, problems)
	assert.Equal(t, len(problems), 10)
	for _, problem := range problems {
		assert.True(t, problem.C >= criteria.Min && problem.C <= criteria.Max)
		assert.True(t, problem.B >= criteria.Min && problem.B <= criteria.Max)
		assert.Equal(t, problem.Op, '/')
		assert.True(t, problem.A/problem.B == problem.C)
	}
}
