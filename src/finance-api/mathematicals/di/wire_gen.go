// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/stratergy"
)

// Injectors from wire.go:

func InitializeTwoPlusService[number entity.Number]() services.ProblemService[number] {
	randomService := data.CreateRandomService[number]()
	twoPlusStratergy := stratergy.NewTwoPlusStratergy[number](randomService)
	twoProblem := services.NewTwoProblem[number](twoPlusStratergy)
	return twoProblem
}

func InitializeTwoMinusService[number entity.Number]() services.ProblemService[number] {
	randomService := data.CreateRandomService[number]()
	twoMinusStratergy := stratergy.NewTwoMinusStratergy[number](randomService)
	twoProblem := services.NewTwoProblem[number](twoMinusStratergy)
	return twoProblem
}

func InitializeTwoMultiplyService[number entity.Number]() services.ProblemService[number] {
	randomService := data.CreateRandomService[number]()
	twoMultiplyStratergy := stratergy.NewTwoMultiplyStratergy[number](randomService)
	twoProblem := services.NewTwoProblem[number](twoMultiplyStratergy)
	return twoProblem
}

func InitializeTwoDivideService[number entity.Number]() services.ProblemService[number] {
	randomService := data.CreateRandomService[number]()
	twoDivideStratergy := stratergy.NewTwoDivideStratergy[number](randomService)
	twoProblem := services.NewTwoProblem[number](twoDivideStratergy)
	return twoProblem
}
