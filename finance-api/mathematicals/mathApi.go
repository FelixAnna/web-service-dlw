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
	c.JSON(http.StatusOK, results)
}

func (api *MathApi) GetAllQuestions(c *gin.Context) {
	var criterias []Criteria
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var results []entity.Problem
	for _, criteria := range criterias {
		results = append(results, api.mathService.GenerateProblems(&criteria)...)
	}

	c.JSON(http.StatusOK, results)
}
