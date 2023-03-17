package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var plusProblem Problem[int]
var plusProblem2 Problem[int]
var minusProblem Problem[int]
var multipleProblem Problem[int]
var multipleProblem2 Problem[int]

func init() {
	plusProblem = Problem[int]{
		A:  1,
		B:  2,
		C:  3,
		Op: '+',
	}

	plusProblem2 = Problem[int]{
		A:  2,
		B:  1,
		C:  3,
		Op: '+',
	}

	minusProblem = Problem[int]{
		A:  3,
		B:  2,
		C:  1,
		Op: '-',
	}

	multipleProblem = Problem[int]{
		A:  3,
		B:  2,
		C:  6,
		Op: '*',
	}

	multipleProblem2 = Problem[int]{
		A:  2,
		B:  3,
		C:  6,
		Op: '*',
	}
}

func TestIdenticalStringPlus(t *testing.T) {
	result := plusProblem.IndenticalString()
	result2 := plusProblem2.IndenticalString()
	assert.EqualValues(t, result2, result)
}

func TestIdenticalStringMultiple(t *testing.T) {
	result := multipleProblem.IndenticalString()
	result2 := multipleProblem2.IndenticalString()
	assert.EqualValues(t, result2, result)
}

func TestNotIdenticalString(t *testing.T) {
	result := plusProblem.IndenticalString()
	result2 := minusProblem.IndenticalString()
	assert.NotEqualValues(t, result2, result)
}
