package mathematicals

import (
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
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

	results := api.mathService.GenerateProblems(criteria)
	c.JSON(http.StatusOK, results)
}

func (api *MathApi) GetAllQuestions(c *gin.Context) {
	var criterias []problem.Criteria
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	results := api.mathService.GenerateProblems(criterias...)

	c.JSON(http.StatusOK, results)
}

func (api *MathApi) GetAllQuestionFeeds(c *gin.Context) {
	var criterias []problem.Criteria
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	results := api.mathService.GenerateFeeds(criterias...)

	c.JSON(http.StatusOK, results)
}
