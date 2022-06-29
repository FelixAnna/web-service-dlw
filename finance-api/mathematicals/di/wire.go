//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

func InitializeTwoPlusService() problem.ProblemService {
	wire.Build(problem.ProblemServiceSet, stratergy.TwoPlusStratergySet, data.RandomServiceSet)
	return nil
}

func InitializeTwoMinusService() problem.ProblemService {
	wire.Build(problem.ProblemServiceSet, stratergy.TwoMinusStratergySet, data.RandomServiceSet)
	return nil
}
