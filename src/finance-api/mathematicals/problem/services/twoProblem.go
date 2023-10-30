package services

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

var ProblemServiceSet = wire.NewSet(NewTwoProblem[int], wire.Bind(new(ProblemService[int]), new(*TwoProblem[int])))
var ProblemServiceSetFloat = wire.NewSet(NewTwoProblem[float32], wire.Bind(new(ProblemService[float32]), new(*TwoProblem[float32])))

type TwoProblem[number entity.Number] struct {
	Stratergy stratergy.Stratergy[number]
}

func NewTwoProblem[number entity.Number](stratergy stratergy.Stratergy[number]) *TwoProblem[number] {
	return &TwoProblem[number]{
		Stratergy: stratergy,
	}
}

func (tp *TwoProblem[number]) GenerateProblem(criteria ...interface{}) *entity.Problem[number] {
	nums := tp.Stratergy.Generate(criteria...)

	var op rune
	switch tp.Stratergy.(type) {
	case *stratergy.TwoPlusStratergy[number]:
		op = '+'
	case *stratergy.TwoMinusStratergy[number]:
		op = '-'
	case *stratergy.TwoMultiplyStratergy[number]:
		op = '*'
	case *stratergy.TwoDivideStratergy[number]:
		op = '/'
	}

	problem := &entity.Problem[number]{
		A: nums[0],
		B: nums[1],
		C: nums[2],

		Op: op,
	}

	return problem
}
