package mathematicals

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/format"
	"github.com/gin-gonic/gin"
)

type MathApi struct {
	mathService *problem.MathService
}

//provide for wire
func ProvideMathApi(mathService *problem.MathService) *MathApi {
	return &MathApi{mathService: mathService}
}

func (api *MathApi) GetQuestions(c *gin.Context) {
	var criteria problem.Criteria
	if err := c.BindJSON(&criteria); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	results := api.mathService.GenerateProblems(&criteria)
	c.JSON(http.StatusOK, GetResponse(results, &criteria))
}

func (api *MathApi) GetAllQuestions(c *gin.Context) {
	var criterias []problem.Criteria
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var results []problem.QuestionModel
	for _, criteria := range criterias {
		cr := criteria
		problems := GetResponse(api.mathService.GenerateProblems(&cr), &cr)
		results = append(results, problems...)
	}

	c.JSON(http.StatusOK, results)
}

func (api *MathApi) GetAllQuestionFeeds(c *gin.Context) {
	var criterias []problem.Criteria
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ch := make(chan []string)
	wg := &sync.WaitGroup{}
	wg.Add(len(criterias))
	for _, criteria := range criterias {
		cr := criteria
		go GetResponseFeed(api.mathService.GenerateProblems(&cr), &cr, wg, ch)
	}

	result, wg2 := processingResults(ch)

	//wait for generating
	wg.Wait()
	close(ch)

	//wait for result processing
	wg2.Wait()

	c.JSON(http.StatusOK, result)
}

func GetResponse(results []entity.Problem, criteria *problem.Criteria) []problem.QuestionModel {
	questions := []problem.QuestionModel{}

	for _, pb := range results {
		expression := getFormatInterface(&pb, criteria)
		question, answer := getDisplayQandA(criteria, expression, pb)

		model := problem.QuestionModel{
			FullText: expression.String(),
			Kind:     criteria.Kind,
			Question: question,
			Answer:   answer,
		}

		questions = append(questions, model)
	}

	return questions
}

func GetResponseFeed(results []entity.Problem, criteria *problem.Criteria, wg *sync.WaitGroup, ch chan<- []string) {
	defer wg.Done()

	for _, pb := range results {
		expression := getFormatInterface(&pb, criteria)
		question, answer := getDisplayQandA(criteria, expression, pb)
		ch <- []string{question, fmt.Sprintf("%v", answer), expression.String()}
	}
}

func getDisplayQandA(criteria *problem.Criteria, expression format.FormatInterface, pb entity.Problem) (string, int) {
	var question string
	var answer int
	switch criteria.Kind {
	case problem.KindQuestFirst:
		question = expression.QuestFirst()
		answer = pb.A
	case problem.KindQuestSecond:
		question = expression.QuestSecond()
		answer = pb.B
	case problem.KindQeustLast:
		question = expression.QuestResult()
		answer = pb.C
	}
	return question, answer
}

func formart(input interface{}, idx int) string {
	return fmt.Sprintf("%v. %v", idx, input)
}

func processingResults(ch chan []string) (*problem.QuestionFeedModel, *sync.WaitGroup) {
	var result = &problem.QuestionFeedModel{}
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

func getFormatInterface(pb *entity.Problem, criteria *problem.Criteria) format.FormatInterface {
	var expression format.FormatInterface
	switch criteria.Type {
	case problem.TypePlainExpression:
		expression = &format.PlainExpression{
			Problem: pb,
		}
	case problem.TypePlainApplication:
		expression = &format.PlainApplication{
			Problem: pb,
		}
	}

	return expression
}
