package format

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/stretchr/testify/assert"
)

var plusProblem PlainExpression
var plusProblem2 PlainExpression
var minusProblem PlainExpression

func init() {
	plusProblem = PlainExpression{
		entity.Problem{
			A:  1,
			B:  2,
			C:  3,
			Op: '+',
		},
	}

	plusProblem2 = PlainExpression{
		entity.Problem{
			A:  2,
			B:  1,
			C:  3,
			Op: '+',
		},
	}

	minusProblem = PlainExpression{
		entity.Problem{
			A:  3,
			B:  2,
			C:  1,
			Op: '-',
		},
	}
}

func TestString(t *testing.T) {
	result := plusProblem.String()
	assert.EqualValues(t, "1 + 2 = 3", result)
}

func TestIdenticalString(t *testing.T) {
	result := plusProblem.IndenticalString()
	result2 := plusProblem2.IndenticalString()
	assert.EqualValues(t, result2, result)
}

func TestNotIdenticalString(t *testing.T) {
	result := plusProblem.IndenticalString()
	result2 := minusProblem.IndenticalString()
	assert.NotEqualValues(t, result2, result)
}

func TestPrintFirst(t *testing.T) {
	result := plusProblem.QuestFirst()
	assert.EqualValues(t, "(  ) + 2 = 3", result)
}

func TestPrintSecond(t *testing.T) {
	result := plusProblem.QuestSecond()
	assert.EqualValues(t, "1 + (  ) = 3", result)
}

func TestPrintLast(t *testing.T) {
	result := plusProblem.QuestResult()
	assert.EqualValues(t, "1 + 2 = (  )", result)
}

func TestPrintMinusLast(t *testing.T) {
	result := minusProblem.QuestResult()
	assert.EqualValues(t, "3 - 2 = (  )", result)
}
