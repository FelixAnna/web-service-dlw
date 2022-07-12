package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/snowflake"
	"github.com/FelixAnna/web-service-dlw/finance-api/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	mathEntity "github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var validToken string

func init() {
	gin.SetMode(gin.TestMode)
	initialMockDependency()
	router = GetGinRouter()

	token, _ := apiBoot.AuthorizationHandler.TokenService.NewToken("testuser", "test@email.com")
	validToken = token.Token
}

func TestRunning(t *testing.T) {

	w := performRequest(router, "GET", "/status", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "running", w.Body.String())
}

func TestGetZdjUnAuthorized(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/zdj/", nil)

	var response []entity.Zhidaojia
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotNil(t, err)
}

func TestGetZdjForbidden(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/zdj/?access_code=123", nil)

	var response []entity.Zhidaojia
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.NotNil(t, err)
}

func TestGetZdjAuthorized(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/zdj/?access_code="+validToken, nil)

	var response []entity.Zhidaojia
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestSearchAuthorized(t *testing.T) {
	//Act
	w := performRequest(router, "POST", "/zdj/search?access_code="+validToken, entity.Criteria{Page: 1, Size: 20})

	var response []entity.Zhidaojia
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetQuestionsInvalid(t *testing.T) {
	//Act
	w := performRequest(router, "POST", "/homework/math/", problem.Criteria{Min: 10, Max: 20, Category: 1000})

	var response []mathEntity.Problem
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotNil(t, err)
}

func TestGetQuestions(t *testing.T) {
	//Act
	w := performRequest(router, "POST", "/homework/math/", problem.Criteria{
		Min:      10,
		Max:      20,
		Quantity: 100,
		Category: problem.CategoryPlus,
		Kind:     problem.KindQeustLast,
		Type:     problem.TypePlainExpression,
	})

	var response problem.QuestionResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetQuestionsMultiple(t *testing.T) {
	//Act
	w := performRequest(router, "POST", "/homework/math/multiple", []problem.Criteria{
		{
			Min:      10,
			Max:      20,
			Quantity: 100,
			Category: problem.CategoryPlus,
			Kind:     problem.KindQeustLast,
			Type:     problem.TypePlainExpression,
		},
	})

	var response problem.QuestionResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	for _, val := range response.Questions {
		assert.True(t, val.Answer > 0)
	}
}

func TestGetQuestionFeedsMultiple(t *testing.T) {
	//Act
	w := performRequest(router, "POST", "/homework/math/multiple/feeds", []problem.Criteria{
		{
			Min:      10,
			Max:      20,
			Quantity: 100,
			Category: problem.CategoryPlus,
			Kind:     problem.KindQeustLast,
		},
	})

	var response problem.QuestionFeedModel
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	for _, val := range response.Answers {
		vals := strings.Split(val, ". ")
		//fmt.Println(val, vals)
		assert.True(t, len(vals) == 2)
		answer, err := strconv.ParseInt(vals[1], 10, 32)
		assert.Nil(t, err)
		assert.Greater(t, answer, int64(0))
	}
}

func TestDeleteAuthorized(t *testing.T) {
	//Act
	w := performRequest(router, "DELETE", "/zdj/123?access_code="+validToken, nil)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.EqualValues(t, w.Body.String(), "\"Data deleted.\"")
}

func TestSlowAuthorized(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/zdj/slow?access_code="+validToken, nil)

	var response map[int]int
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func performRequest(r http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	var readerOfBody io.Reader = nil
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			log.Fatal(err)
		}

		readerOfBody = bytes.NewReader(data)
	}

	req, _ := http.NewRequest(method, path, readerOfBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func initialMockDependency() {
	apiBoot = &ApiBoot{}
	zdjApi, err := di.InitializeMockApi()
	if err != nil {
		log.Panic(err)
		return
	}

	apiBoot = &ApiBoot{
		ZdjApi:               zdjApi,
		MathApi:              di.InitializeMathApi(),
		AuthorizationHandler: di.InitialMockAuthorizationMiddleware(),
		ErrorHandler:         di.InitialErrorMiddleware(),
		Registry:             di.InitialMockRegistry(),
	}

	os.Setenv("DLW_NODE_NO", "1023")
	snowflake.InitSnowflake()
}
