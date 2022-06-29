package plus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

type TwoPlusProblemLast struct {
	TwoPlusProblem
}

var TwoPlusProblemLastSet = wire.NewSet(NewTwoPlusProblemLast, wire.Bind(new(problem.ProblemService), new(*TwoPlusProblemLast)))

func NewTwoPlusProblemLast(stratergy stratergy.Stratergy) *TwoPlusProblemLast {
	return &TwoPlusProblemLast{
		TwoPlusProblem: *NewTwoPlusProblem(stratergy),
	}
}

func (tp *TwoPlusProblemLast) Print() string {
	return fmt.Sprintf("%v + %v = ?", tp.Problem.A, tp.Problem.B)
}
