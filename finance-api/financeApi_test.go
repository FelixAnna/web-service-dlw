package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine
var validToken string

func init() {
	gin.SetMode(gin.TestMode)
	router = gin.New()
	initialMockDependency()
	defineRoutes(router)

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

	apiBoot.ZdjApi = &zdjApi
	apiBoot.Registry = di.InitialMockRegistry()
	apiBoot.AuthorizationHandler = di.InitialAuthorizationMiddleware()
	apiBoot.ErrorHandler = di.InitialErrorMiddleware()
}
