package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var plusProblem Problem
var minusProblem Problem

func init() {
	plusProblem = Problem{
		A:  1,
		B:  2,
		C:  3,
		Op: '+',
	}

	minusProblem = Problem{
		A:  3,
		B:  2,
		C:  1,
		Op: '-',
	}
}

func TestPrintAll(t *testing.T) {
	result := plusProblem.String()
	assert.EqualValues(t, "1 + 2 = 3", result)
}

func TestPrintFirst(t *testing.T) {
	result := plusProblem.QuestFirst()
	assert.EqualValues(t, "? + 2 = 3", result)
}

func TestPrintSecond(t *testing.T) {
	result := plusProblem.QuestSecond()
	assert.EqualValues(t, "1 + ? = 3", result)
}

func TestPrintLast(t *testing.T) {
	result := plusProblem.QuestResult()
	assert.EqualValues(t, "1 + 2 = ?", result)
}
