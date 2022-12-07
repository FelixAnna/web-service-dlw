package mathematicals

import (
	"fmt"
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
func (api *MathApi) SaveResults(c *gin.Context) {
	userId, _ := c.Get("userId")
	var request problem.SaveAnswersRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := api.mathService.SaveResults(&request, userId.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Saved for user: %v!", userId))
}
