// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
)

// Injectors from wire.go:

func InitializeTwoPlusService() services.ProblemService {
	randomService := data.CreateRandomService()
	twoPlusStratergy := stratergy.NewTwoPlusStratergy(randomService)
	twoProblem := services.NewTwoProblem(twoPlusStratergy)
	return twoProblem
}

func InitializeTwoMinusService() services.ProblemService {
	randomService := data.CreateRandomService()
	twoMinusStratergy := stratergy.NewTwoMinusStratergy(randomService)
	twoProblem := services.NewTwoProblem(twoMinusStratergy)
	return twoProblem
}

func InitializeTwoMultiplyService() services.ProblemService {
	randomService := data.CreateRandomService()
	twoMultiplyStratergy := stratergy.NewTwoMultiplyStratergy(randomService)
	twoProblem := services.NewTwoProblem(twoMultiplyStratergy)
	return twoProblem
}

func InitializeTwoDivideService() services.ProblemService {
	randomService := data.CreateRandomService()
	twoDivideStratergy := stratergy.NewTwoDivideStratergy(randomService)
	twoProblem := services.NewTwoProblem(twoDivideStratergy)
	return twoProblem
}
