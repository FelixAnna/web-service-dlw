package mathematicals

import (
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
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

func GetResponse(results []entity.Problem, kind int) []QuestionModel {
	questions := []QuestionModel{}

	for _, problem := range results {
		model := QuestionModel{
			FullText: problem.String(),
			Kind:     kind,
		}

		switch kind {
		case KindQuestFirst:
			model.Question = problem.QuestFirst()
			model.Answer = problem.A
		case KindQuestSecond:
			model.Question = problem.QuestSecond()
			model.Answer = problem.B
		case KindQeustLast:
			model.Question = problem.QuestResult()
			model.Answer = problem.C
		}

		questions = append(questions, model)
	}

	return questions
}
