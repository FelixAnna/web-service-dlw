package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var plusProblem Problem
var plusProblem2 Problem
var minusProblem Problem

func init() {
	plusProblem = Problem{
		A:  1,
		B:  2,
		C:  3,
		Op: '+',
	}

	plusProblem2 = Problem{
		A:  2,
		B:  1,
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
