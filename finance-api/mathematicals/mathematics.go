package mathematicals

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
)

type MathService struct {
	TwoPlusService  problem.ProblemService
	TwoMinusService problem.ProblemService
}

const (
	KindQuestFirst  int = 1
	KindQuestSecond int = 2
	KindQeustLast   int = 3
)

const (
	CategoryPlus int = iota
	CategoryMinus
)

type Criteria struct {
	Min, Max int
	Quantity int

	PositiveOnly bool

	//+, -
	Category int

	//first, second, last
	Kind int
}

type QuestionModel struct {
	Question string
	Answer   string
}

func NewMathService() *MathService {
	return &MathService{
		TwoPlusService:  di.InitializeTwoPlusService(),
		TwoMinusService: di.InitializeTwoMinusService(),
	}
}

func (service *MathService) GenerateProblems(criteria *Criteria) []entity.Problem {
	postiveOnly := 0
	if criteria.PositiveOnly {
		postiveOnly = 1
	}

	if criteria.Quantity == 0 {
		criteria.Quantity = 10
	}

	var problems []entity.Problem = []entity.Problem{}

	switch criteria.Category {
	case CategoryMinus:
		service.GenerateMinusProblem(criteria, postiveOnly, &problems)
	case CategoryPlus:
		service.GeneratePlusProblem(criteria, &problems)
	}

	return problems
}

func (service *MathService) GenerateMinusProblem(criteria *Criteria, posOnly int, problems *[]entity.Problem) {
	for i := 0; i < criteria.Quantity; i++ {
		problem := service.TwoMinusService.GenerateProblem(criteria.Min, criteria.Max, posOnly)
		*problems = append(*problems, *problem)
	}
}

func (service *MathService) GeneratePlusProblem(criteria *Criteria, problems *[]entity.Problem) {
	for i := 0; i < criteria.Quantity; i++ {
		problem := service.TwoPlusService.GenerateProblem(criteria.Min, criteria.Max)
		*problems = append(*problems, *problem)
	}
}
