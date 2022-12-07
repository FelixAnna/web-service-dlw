package format

import (
	"fmt"
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/stretchr/testify/assert"
)

var plusProblem PlainExpression
var plusProblem2 PlainExpression
var minusProblem PlainExpression
var multiplyProblem PlainExpression
var divideProblem PlainExpression
var UnSupportedProblem PlainExpression

func init() {
	plusProblem = PlainExpression{
		&entity.Problem{
			A:  1,
			B:  2,
			C:  3,
			Op: '+',
		},
	}

	plusProblem2 = PlainExpression{
		&entity.Problem{
			A:  2,
			B:  1,
			C:  3,
			Op: '+',
		},
	}

	minusProblem = PlainExpression{
		&entity.Problem{
			A:  3,
			B:  2,
			C:  1,
			Op: '-',
		},
	}

	multiplyProblem = PlainExpression{
		&entity.Problem{
			A:  3,
			B:  2,
			C:  6,
			Op: '*',
		},
	}

	divideProblem = PlainExpression{
		&entity.Problem{
			A:  3,
			B:  2,
			C:  6,
			Op: '/',
		},
	}

	UnSupportedProblem = PlainExpression{
		&entity.Problem{
			A:  3,
			B:  2,
			C:  6,
			Op: '&',
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
	assert.EqualValues(t, fmt.Sprintf("%v + 2 = 3", placeHolder), result)
}

func TestPrintSecond(t *testing.T) {
	result := plusProblem.QuestSecond()
	assert.EqualValues(t, fmt.Sprintf("1 + %v = 3", placeHolder), result)
}

func TestPrintLast(t *testing.T) {
	result := plusProblem.QuestResult()
	assert.EqualValues(t, fmt.Sprintf("1 + 2 = %v", placeHolder), result)
}

func TestPrintMinusLast(t *testing.T) {
	result := minusProblem.QuestResult()
	assert.EqualValues(t, fmt.Sprintf("3 - 2 = %v", placeHolder), result)
}

func TestPrintMultiplyLast(t *testing.T) {
	result := multiplyProblem.QuestResult()
	assert.EqualValues(t, fmt.Sprintf("3 * 2 = %v", placeHolder), result)
}

func TestPrintDivide(t *testing.T) {
	result := divideProblem.QuestResult()
	assert.EqualValues(t, fmt.Sprintf("3 / 2 = %v", placeHolder), result)
}

func TestPrintNotSupport(t *testing.T) {
	result := UnSupportedProblem.QuestResult()
	assert.EqualValues(t, fmt.Sprintf("3 ? 2 = %v", placeHolder), result)
}
