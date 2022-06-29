package plus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

type TwoPlusProblemFirst struct {
	TwoPlusProblem
}

var TwoPlusProblemFirstSet = wire.NewSet(NewTwoPlusProblemFirst, wire.Bind(new(problem.ProblemService), new(*TwoPlusProblemFirst)))

func NewTwoPlusProblemFirst(stratergy stratergy.Stratergy) *TwoPlusProblemFirst {
	return &TwoPlusProblemFirst{
		TwoPlusProblem: *NewTwoPlusProblem(stratergy),
	}
}

func (tp *TwoPlusProblemFirst) Print() string {
	return fmt.Sprintf("? + %v = %v", tp.Problem.B, tp.Problem.C)
}
