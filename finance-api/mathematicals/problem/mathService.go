package problem

import (
	"fmt"
	"sync"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/format"
)

type MathService struct {
	genService *TwoGenerationService
}

func NewMathService(genService *TwoGenerationService) *MathService {
	return &MathService{
		genService: genService,
	}
}

func (service *MathService) GenerateProblems(criterias ...Criteria) []QuestionModel {
	var results []QuestionModel
	for _, criteria := range criterias {
		cr := criteria
		problems := GetResponse(service.genService.GenerateProblems(&cr), &cr)
		results = append(results, problems...)
	}

	return results
}

func (service *MathService) GenerateFeeds(criterias ...Criteria) *QuestionFeedModel {
	ch := make(chan []string)
	wg := &sync.WaitGroup{}
	wg.Add(len(criterias))
	for _, criteria := range criterias {
		cr := criteria
		go GetResponseFeed(service.genService.GenerateProblems(&cr), &cr, wg, ch)
	}

	result, wg2 := processingResults(ch)

	//wait for generating
	wg.Wait()
	close(ch)

	//wait for result processing
	wg2.Wait()

	return result
}

func GetResponse(results []entity.Problem, criteria *Criteria) []QuestionModel {
	questions := []QuestionModel{}

	for _, pb := range results {
		expression := getFormatInterface(&pb, criteria)
		question, answer := getDisplayQandA(criteria, expression, pb)

		model := QuestionModel{
			FullText: expression.String(),
			Kind:     criteria.Kind,
			Question: question,
			Answer:   answer,
		}

		questions = append(questions, model)
	}

	return questions
}

func GetResponseFeed(results []entity.Problem, criteria *Criteria, wg *sync.WaitGroup, ch chan<- []string) {
	defer wg.Done()

	for _, pb := range results {
		expression := getFormatInterface(&pb, criteria)
		question, answer := getDisplayQandA(criteria, expression, pb)
		ch <- []string{question, fmt.Sprintf("%v", answer), expression.String()}
	}
}

func getDisplayQandA(criteria *Criteria, expression format.FormatInterface, pb entity.Problem) (string, int) {
	var question string
	var answer int
	switch criteria.Kind {
	case KindQuestFirst:
		question = expression.QuestFirst()
		answer = pb.A
	case KindQuestSecond:
		question = expression.QuestSecond()
		answer = pb.B
	case KindQeustLast:
		question = expression.QuestResult()
		answer = pb.C
	}
	return question, answer
}

func formart(input interface{}, idx int) string {
	return fmt.Sprintf("%v. %v", idx, input)
}

func processingResults(ch chan []string) (*QuestionFeedModel, *sync.WaitGroup) {
	var result = &QuestionFeedModel{}
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg2.Done()
		idx := 0
		for vals := range ch {
			idx++
			result.Questions = append(result.Questions, formart(vals[0], idx))
			result.Answers = append(result.Answers, formart(vals[1], idx))
			result.FullText = append(result.FullText, formart(vals[2], idx))
		}
	}(wg2)
	return result, wg2
}

func getFormatInterface(pb *entity.Problem, criteria *Criteria) format.FormatInterface {
	var expression format.FormatInterface
	switch criteria.Type {
	case TypePlainExpression:
		expression = &format.PlainExpression{
			Problem: pb,
		}
	case TypePlainApplication:
		expression = &format.PlainApplication{
			Problem: pb,
		}
	}

	return expression
}
