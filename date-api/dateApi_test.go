package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/FelixAnna/web-service-dlw/date-api/di"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = gin.New()
	initialMockDependency()
	defineRoutes(router)
}

func TestRunning(t *testing.T) {

	w := performRequest(router, "GET", "/status")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "running", w.Body.String())
}

func TestGetMonthDateNil(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/date/current/month")

	var response []entity.DLWDate
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, len(response) >= 28)
}

func TestGetMonthDate(t *testing.T) {

	//Act
	w := performRequest(router, "GET", "/date/current/month?date=20200505")

	var response []entity.DLWDate
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, response[0].YMD, 20200426)
	assert.True(t, len(response) > 31)
}

func TestGetDateDistance(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/date/distance?start=20200101&end=20200505")

	var response entity.Distance
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, response.After, 241)
	assert.EqualValues(t, response.Before, -125)
}

func TestGetLunarDateDistance(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/date/distance/lunar?start=20200101&end=20200505")

	var response entity.Distance
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, response.After, 229)
	assert.EqualValues(t, response.Before, -155)
}

func TestGetDateDistanceInvalid(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/date/distance?start=&end=20200505")

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.EqualValues(t, w.Body.String(), "Not a number")
}

func TestGetLunarDateDistanceInvalid(t *testing.T) {
	//Act
	w := performRequest(router, "GET", "/date/distance/lunar?start=&end=20200505")

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.EqualValues(t, w.Body.String(), "Not a number")
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func initialMockDependency() {
	apiBoot = &ApiBoot{}
	dateApi := di.InitialDateApi()

	apiBoot.DateApi = dateApi
	//apiBoot.AuthorizationHandler = di.InitialAuthorizationMiddleware()
	apiBoot.ErrorHandler = di.InitialErrorMiddleware()
	apiBoot.Registry = di.InitialMockRegistry()
}
