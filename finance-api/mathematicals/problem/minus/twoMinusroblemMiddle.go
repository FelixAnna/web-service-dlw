package minus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

type TwoMinusProblemMiddle struct {
	TwoMinusProblem
}

var TwoMinusProblemMiddleSet = wire.NewSet(NewTwoMinusProblemMiddle, wire.Bind(new(problem.ProblemService), new(*TwoMinusProblemMiddle)))

func NewTwoMinusProblemMiddle(stratergy stratergy.Stratergy) *TwoMinusProblemMiddle {
	return &TwoMinusProblemMiddle{
		TwoMinusProblem: *NewTwoMinusProblem(stratergy),
	}
}

func (tp *TwoMinusProblemMiddle) Print() string {
	return fmt.Sprintf("%v - ? =%v ", tp.Problem.A, tp.Problem.C)
}
