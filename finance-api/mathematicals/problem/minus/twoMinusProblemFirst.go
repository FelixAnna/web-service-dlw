package minus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

type TwoMinusProblemFirst struct {
	TwoMinusProblem
}

var TwoMinusProblemFirstSet = wire.NewSet(NewTwoMinusProblemFirst, wire.Bind(new(problem.ProblemService), new(*TwoMinusProblemFirst)))

func NewTwoMinusProblemFirst(stratergy stratergy.Stratergy) *TwoMinusProblemFirst {
	return &TwoMinusProblemFirst{
		TwoMinusProblem: *NewTwoMinusProblem(stratergy),
	}
}

func (tp *TwoMinusProblemFirst) Print() string {
	return fmt.Sprintf("? - %v = %v", tp.Problem.B, tp.Problem.C)
}
