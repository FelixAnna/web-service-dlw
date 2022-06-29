package minus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

type TwoMinusProblemLast struct {
	TwoMinusProblem
}

var TwoMinusProblemLastSet = wire.NewSet(NewTwoMinusProblemLast, wire.Bind(new(problem.ProblemService), new(*TwoMinusProblemLast)))

func NewTwoMinusProblemLast(stratergy stratergy.Stratergy) *TwoMinusProblemLast {
	return &TwoMinusProblemLast{
		TwoMinusProblem: *NewTwoMinusProblem(stratergy),
	}
}

func (tp *TwoMinusProblemLast) Print() string {
	return fmt.Sprintf("%v - %v = ?", tp.Problem.A, tp.Problem.B)
}
