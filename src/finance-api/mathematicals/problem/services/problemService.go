package services

import "github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"

type ProblemService[number entity.Number] interface {
	GenerateProblem(criteria ...interface{}) *entity.Problem[number]
}
