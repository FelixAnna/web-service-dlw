package mathematicals

import (
	"log"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

const MaxGenerateTimes = 10000

type MathService struct {
	TwoPlusService  problem.ProblemService
	TwoMinusService problem.ProblemService
}

func NewMathService() *MathService {
	return &MathService{
		TwoPlusService:  di.InitializeTwoPlusService(),
		TwoMinusService: di.InitializeTwoMinusService(),
	}
}

func (service *MathService) GenerateProblems(criteria *Criteria) []entity.Problem {
	if criteria.Quantity == 0 {
		criteria.Quantity = 10
	}

	var problems []entity.Problem = []entity.Problem{}

	var problemService problem.ProblemService
	switch criteria.Category {
	case CategoryMinus:
		problemService = service.TwoMinusService
	case CategoryPlus:
		problemService = service.TwoPlusService
	default:
		log.Println("Invalid Category:", criteria.Category)
	}

	GenerateProblems(criteria, problemService, &problems)
	return problems
}

func GenerateProblems(criteria *Criteria, problemService problem.ProblemService, problems *[]entity.Problem) {
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
			problemTexts[problem.String()] {
			i--
			continue
		}

		problemTexts[problem.String()] = true
		*problems = append(*problems, *problem)
	}
}
