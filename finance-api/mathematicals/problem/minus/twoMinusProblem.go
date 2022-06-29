package minus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
)

type TwoMinusProblem struct {
	*problem.TwoProblem
}

func NewTwoMinusProblem(stratergy stratergy.Stratergy) *TwoMinusProblem {
	return &TwoMinusProblem{
		TwoProblem: problem.NewTwoProblem(stratergy),
	}
}

func (tp *TwoMinusProblem) PrintAll() string {
	return fmt.Sprintf("%v - %v = %v", tp.Problem.A, tp.Problem.B, tp.Problem.C)
}
