package mathematicals

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/format"
	"github.com/gin-gonic/gin"
)

type MathApi struct {
	mathService *MathService
}

//provide for wire
func ProvideMathApi(mathService *MathService) *MathApi {
	return &MathApi{mathService: mathService}
}

func (api *MathApi) GetQuestions(c *gin.Context) {
	var criteria Criteria
	if err := c.BindJSON(&criteria); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	results := api.mathService.GenerateProblems(&criteria)
	c.JSON(http.StatusOK, GetResponse(results, criteria.Kind))
}

func (api *MathApi) GetAllQuestions(c *gin.Context) {
	var criterias []Criteria
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var results []QuestionModel
	for _, criteria := range criterias {
		problems := GetResponse(api.mathService.GenerateProblems(&criteria), criteria.Kind)
		results = append(results, problems...)
	}

	c.JSON(http.StatusOK, results)
}

func (api *MathApi) GetAllQuestionFeeds(c *gin.Context) {
	var criterias []Criteria
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ch := make(chan []string)
	wg := &sync.WaitGroup{}
	wg.Add(len(criterias))
	for _, criteria := range criterias {
		go GetResponseFeed(api.mathService.GenerateProblems(&criteria), criteria.Kind, wg, ch)
	}

	result, wg2 := processingResults(ch)

	//wait for generating
	wg.Wait()
	close(ch)

	//wait for result processing
	wg2.Wait()

	c.JSON(http.StatusOK, result)
}

func GetResponse(results []entity.Problem, kind int) []QuestionModel {
	questions := []QuestionModel{}

	for _, problem := range results {
		expression := &format.PlainExpression{
			Problem: problem,
		}
		model := QuestionModel{
			FullText: expression.String(),
			Kind:     kind,
		}

		switch kind {
		case KindQuestFirst:
			model.Question = expression.QuestFirst()
			model.Answer = problem.A
		case KindQuestSecond:
			model.Question = expression.QuestSecond()
			model.Answer = problem.B
		case KindQeustLast:
			model.Question = expression.QuestResult()
			model.Answer = problem.C
		}

		questions = append(questions, model)
	}

	return questions
}

func GetResponseFeed(results []entity.Problem, kind int, wg *sync.WaitGroup, ch chan<- []string) {
	defer wg.Done()

	for _, problem := range results {
		expression := &format.PlainApplication{
			Problem: problem,
		}
		var question string
		var answer int
		switch kind {
		case KindQuestFirst:
			question = expression.QuestFirst()
			answer = problem.A
		case KindQuestSecond:
			question = expression.QuestSecond()
			answer = problem.B
		case KindQeustLast:
			question = expression.QuestResult()
			answer = problem.C
		}

		ch <- []string{question, fmt.Sprintf("%v", answer), expression.String()}
	}
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
