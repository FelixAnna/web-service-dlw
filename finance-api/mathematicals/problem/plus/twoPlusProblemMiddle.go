package plus

import (
	"fmt"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

type TwoPlusProblemMiddle struct {
	TwoPlusProblem
}

var TwoPlusProblemMiddleSet = wire.NewSet(NewTwoPlusProblemMiddle, wire.Bind(new(problem.ProblemService), new(*TwoPlusProblemMiddle)))

func NewTwoPlusProblemMiddle(stratergy stratergy.Stratergy) *TwoPlusProblemMiddle {
	return &TwoPlusProblemMiddle{
		TwoPlusProblem: *NewTwoPlusProblem(stratergy),
	}
}

func (tp *TwoPlusProblemMiddle) Print() string {
	return fmt.Sprintf("%v + ? =%v ", tp.Problem.A, tp.Problem.C)
}
