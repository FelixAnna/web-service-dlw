package problem

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/stretchr/testify/assert"
)

var twoProblem *TwoProblem
var dataService data.DataService

func init() {
	dataService = data.CreateRandomService()
}

func TestCreateRandomService(t *testing.T) {
	twoProblem = NewTwoProblem(stratergy.NewTwoMinusStratergy(dataService))

	assert.NotNil(t, twoProblem)
	assert.NotNil(t, twoProblem.Stratergy)
}

func TestGenerateProblemPlus(t *testing.T) {
	twoProblem = NewTwoProblem(stratergy.NewTwoPlusStratergy(dataService))

	problem := twoProblem.GenerateProblem(100, 200)
	assert.NotNil(t, problem)
	assert.True(t, problem.A >= 100 && problem.A <= 200)
	assert.True(t, problem.B >= 100 && problem.B <= 200)
	assert.Equal(t, problem.Op, '+')
	assert.True(t, problem.A+problem.B == problem.C)
}

func TestGenerateProblemMinus(t *testing.T) {
	twoProblem = NewTwoProblem(stratergy.NewTwoMinusStratergy(dataService))

	problem := twoProblem.GenerateProblem(100, 200)
	assert.NotNil(t, problem)
	assert.True(t, problem.A >= 100 && problem.A <= 200)
	assert.True(t, problem.B >= 100 && problem.B <= 200)
	assert.Equal(t, problem.Op, '-')
	assert.True(t, problem.A-problem.B == problem.C)
}
