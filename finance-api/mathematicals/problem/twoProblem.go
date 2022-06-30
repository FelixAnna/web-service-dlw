package problem

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

var ProblemServiceSet = wire.NewSet(NewTwoProblem, wire.Bind(new(ProblemService), new(*TwoProblem)))

type TwoProblem struct {
	Stratergy stratergy.Stratergy
}

func NewTwoProblem(stratergy stratergy.Stratergy) *TwoProblem {
	return &TwoProblem{
		Stratergy: stratergy,
	}
}

func (tp *TwoProblem) GenerateProblem(criteria ...interface{}) *entity.Problem {
	nums := tp.Stratergy.Generate(criteria...)

	var op rune
	switch tp.Stratergy.(type) {
	case *stratergy.TwoPlusStratergy:
		op = '+'
	case *stratergy.TwoMinusStratergy:
		op = '-'
	}

	problem := &entity.Problem{
		A: nums[0],
		B: nums[1],
		C: nums[2],

		Op: op,
	}

	return problem
}
