package mathematicals

import (
	"log"
	"math"

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

type Range struct {
	Min, Max int
}
type Criteria struct {
	Min, Max int
	Quantity int

	Range *Range

	//+, -
	Category int

	//first, second, last
	Kind int
}

func (s *Criteria) GetRange() (min, max int) {
	min, max = math.MinInt32, math.MaxInt32
	if s.Range == nil {
		return
	}

	if s.Range.Min > s.Range.Max {
		return s.Range.Max, s.Range.Min
	}

	return s.Range.Min, s.Range.Max
}

type QuestionModel struct {
	Question string
	Answer   int

	Kind     int
	FullText string
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
	}

	GenerateProblems(criteria, problemService, &problems)
	return problems
}

func GenerateProblems(criteria *Criteria, problemService problem.ProblemService, problems *[]entity.Problem) {
	round := 0
	for i := 0; i < criteria.Quantity; i++ {
		round++
		if round > 10000 {
			log.Println("Too many attampts")
			break
		}

		problem := problemService.GenerateProblem(criteria.Min, criteria.Max)

		minResult, maxResult := criteria.GetRange()
		if problem.C > maxResult {
			i--
			continue
		}

		if problem.C < minResult {
			i--
			continue
		}

		*problems = append(*problems, *problem)
	}
}
