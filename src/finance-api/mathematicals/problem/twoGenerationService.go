package problem

import (
	"log"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services"
)

const MaxGenerateTimes = 10000

type TwoGenerationService[number entity.Number] struct {
	TwoPlusService     services.ProblemService[number]
	TwoMinusService    services.ProblemService[number]
	TwoMultiplyService services.ProblemService[number]
	TwoDivideService   services.ProblemService[number]
}

func NewTwoGenerationService[number entity.Number]() *TwoGenerationService[number] {
	return &TwoGenerationService[number]{
		TwoPlusService:     di.InitializeTwoPlusService[number](),
		TwoMinusService:    di.InitializeTwoMinusService[number](),
		TwoMultiplyService: di.InitializeTwoMultiplyService[number](),
		TwoDivideService:   di.InitializeTwoDivideService[number](),
	}
}

func (service *TwoGenerationService[number]) GenerateProblems(criteria *Criteria[number]) []entity.Problem[number] {
	if criteria.Quantity == 0 {
		criteria.Quantity = 10
	}

	var problems []entity.Problem[number] = []entity.Problem[number]{}

	var problemService services.ProblemService[number]
	switch criteria.Category {
	case CategoryMinus:
		problemService = service.TwoMinusService
	case CategoryPlus:
		problemService = service.TwoPlusService
	case CategoryMultiply:
		problemService = service.TwoMultiplyService
	case CategoryDivide:
		problemService = service.TwoDivideService
	default:
		log.Println("Invalid Category:", criteria.Category)
	}

	GenerateProblems(criteria, problemService, &problems)
	return problems
}

func GenerateProblems[number entity.Number](criteria *Criteria[number], problemService services.ProblemService[number], problems *[]entity.Problem[number]) {
	round := 0
	problemTexts := map[string]bool{}
	for i := 0; i < criteria.Quantity; i++ {
		round++
		if round > MaxGenerateTimes {
			log.Println("Too many attampts: ", MaxGenerateTimes)
			break
		}

		problem := problemService.GenerateProblem(criteria.Min, criteria.Max)

		minResult, maxResult := criteria.GetRange()
		if problem.C > maxResult ||
			problem.C < minResult ||
			problemTexts[problem.IndenticalString()] {
			i--
			continue
		}

		problemTexts[problem.IndenticalString()] = true
		*problems = append(*problems, *problem)
	}
}
