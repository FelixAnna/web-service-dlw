package plus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
)

type TwoPlusProblem struct {
	*problem.TwoProblem
}

func NewTwoPlusProblem(stratergy stratergy.Stratergy) *TwoPlusProblem {
	return &TwoPlusProblem{
		TwoProblem: problem.NewTwoProblem(stratergy),
	}
}

func (tp *TwoPlusProblem) PrintAll() string {
	return fmt.Sprintf("%v + %v = %v", tp.Problem.A, tp.Problem.B, tp.Problem.C)
}
