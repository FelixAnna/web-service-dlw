//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

func InitializeTwoPlusService() services.ProblemService {
	wire.Build(services.ProblemServiceSet, stratergy.TwoPlusStratergySet, data.RandomServiceSet)
	return nil
}

func InitializeTwoMinusService() services.ProblemService {
	wire.Build(services.ProblemServiceSet, stratergy.TwoMinusStratergySet, data.RandomServiceSet)
	return nil
}
