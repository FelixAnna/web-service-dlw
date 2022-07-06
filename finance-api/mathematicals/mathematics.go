package mathematicals

import (
	"log"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services"
)

const MaxGenerateTimes = 10000

type MathService struct {
	TwoPlusService  services.ProblemService
	TwoMinusService services.ProblemService
}

func NewMathService() *MathService {
	return &MathService{
		TwoPlusService:  di.InitializeTwoPlusService(),
		TwoMinusService: di.InitializeTwoMinusService(),
	}
}

func (service *MathService) GenerateProblems(criteria *problem.Criteria) []entity.Problem {
	if criteria.Quantity == 0 {
		criteria.Quantity = 10
	}

	var problems []entity.Problem = []entity.Problem{}

	var problemService services.ProblemService
	switch criteria.Category {
	case problem.CategoryMinus:
		problemService = service.TwoMinusService
	case problem.CategoryPlus:
		problemService = service.TwoPlusService
	default:
		log.Println("Invalid Category:", criteria.Category)
	}

	GenerateProblems(criteria, problemService, &problems)
	return problems
}

func GenerateProblems(criteria *problem.Criteria, problemService services.ProblemService, problems *[]entity.Problem) {
	round := 0
	problemTexts := map[string]bool{}
	for i := 0; i < criteria.Quantity; i++ {
		round++
		if round > MaxGenerateTimes {
			log.Println("Too many attampts")
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
