package services

import "github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/entity"

type ProblemService interface {
	GenerateProblem(criteria ...interface{}) *entity.Problem
}
