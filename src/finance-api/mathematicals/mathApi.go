package mathematicals

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/gin-gonic/gin"
)

type MathApi[number entity.Number] struct {
	mathService *problem.MathService[number]
}

// provide for wire
func ProvideMathApi(mathService *problem.MathService[int]) *MathApi[int] {
	return &MathApi[int]{mathService: mathService}
}

func ProvideMathApi2(mathService *problem.MathService[float32]) *MathApi[float32] {
	return &MathApi[float32]{mathService: mathService}
}

func (api *MathApi[number]) GetQuestions(c *gin.Context) {
	var criteria problem.Criteria[number]
	if err := c.BindJSON(&criteria); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	results := api.mathService.GenerateProblems(criteria)
	c.JSON(http.StatusOK, results)
}

func (api *MathApi[number]) GetAllQuestions(c *gin.Context) {
	var criterias []problem.Criteria[number]
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	results := api.mathService.GenerateProblems(criterias...)

	c.JSON(http.StatusOK, results)
}

func (api *MathApi[number]) GetAllQuestionFeeds(c *gin.Context) {
	var criterias []problem.Criteria[number]
	if err := c.BindJSON(&criterias); err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	results := api.mathService.GenerateFeeds(criterias...)

	c.JSON(http.StatusOK, results)
}
func (api *MathApi[number]) SaveResults(c *gin.Context) {
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
