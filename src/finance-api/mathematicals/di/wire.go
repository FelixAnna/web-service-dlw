//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
	"github.com/google/wire"
)

func InitializeTwoPlusService[number entity.Number]() services.ProblemService[number] {
	var zero number
	switch zero.(type) {
	case int:
		wire.Build(services.ProblemServiceSet, stratergy.TwoPlusStratergySet, data.RandomServiceSet)
	case float32:
		wire.Build(services.ProblemServiceSetFloat, stratergy.TwoPlusStratergySetFloat, data.RandomServiceSetFloat)
	default:
		panic("Invalid type")
	}
	return nil
}

func InitializeTwoMinusService[number entity.Number]() services.ProblemService[number] {
	switch (number); {
	case int:
		wire.Build(services.ProblemServiceSet, stratergy.TwoMinusStratergySet, data.RandomServiceSet)
	case float32:
		wire.Build(services.ProblemServiceSetFloat, stratergy.TwoMinusStratergySetFloat, data.RandomServiceSetFloat)
	default:
		panic("Invalid type")
	}
	return nil
}

func InitializeTwoMultiplyService[number entity.Number]() services.ProblemService[number] {
	switch (number); {
	case int:
		wire.Build(services.ProblemServiceSet, stratergy.TwoMultiplyStratergySet, data.RandomServiceSet)
	case float32:
		wire.Build(services.ProblemServiceSetFloat, stratergy.TwoMultiplyStratergySetFloat, data.RandomServiceSetFloat)
	default:
		panic("Invalid type")
	}
	return nil
}

func InitializeTwoDivideService[number entity.Number]() services.ProblemService[number] {
	switch (number); {
	case int:
		wire.Build(services.ProblemServiceSet, stratergy.TwoDivideStratergySet, data.RandomServiceSet)
	case float32:
		wire.Build(services.ProblemServiceSetFloat, stratergy.TwoDivideStratergySetFloat, data.RandomServiceSetFloat)
	default:
		panic("Invalid type")
	}
	return nil
}
