package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixAnna/web-service-dlw/memo-api/di"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/entity"
	"github.com/FelixAnna/web-service-dlw/memo-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var router *gin.Engine
var validToken string

var memoGregorian entity.Memo
var memoLunar entity.Memo

var memoRequest entity.MemoRequest
var memoRequestInvalid entity.MemoRequest

func init() {
	memoGregorian = entity.Memo{
		Id:          "any",
		Subject:     "any",
		Description: "any",
		UserId:      "any",
		MonthDay:    1228,
		StartYear:   1990,
		Lunar:       false,
	}

	memoLunar = entity.Memo{
		Id:          "any",
		Subject:     "any",
		Description: "any",
		UserId:      "any",
		MonthDay:    1228,
		StartYear:   1990,
		Lunar:       true,
	}

	memoRequest = entity.MemoRequest{
		Subject:  "any",
		MonthDay: 1228,
	}

	memoRequestInvalid = entity.MemoRequest{}
}

func TestRunning(t *testing.T) {
	//Arrange
	setupService(t)

	w := performRequest(router, "GET", "/status", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "running", w.Body.String())
}

func TestGetMemosUnAuthorized(t *testing.T) {
	//Arrange
	setupService(t)

	//Act
	w := performRequest(router, "GET", "/memos/", nil)

	var response []entity.MemoResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotNil(t, err)
}

func TestGetMemosForbidden(t *testing.T) {
	//Arrange
	setupService(t)

	//Act
	w := performRequest(router, "GET", "/memos/?access_code=123", nil)

	var response []entity.MemoResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.NotNil(t, err)
}

func TestGetMemosAuthorized(t *testing.T) {
	//Arrange
	mockRepo, mockDataService := setupService(t)
	mockRepo.EXPECT().GetByUserId(mock.Anything).Return([]entity.Memo{memoGregorian, memoLunar}, nil)
	mockDataService.EXPECT().GetDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200)
	mockDataService.EXPECT().GetLunarDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200)

	//Act
	w := performRequest(router, "GET", "/memos/?access_code="+validToken, nil)

	var response []entity.MemoResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetByIdAuthorized(t *testing.T) {
	//Arrange
	mockRepo, mockDataService := setupService(t)
	mockRepo.EXPECT().GetById(mock.Anything).Return(&memoGregorian, nil)
	mockDataService.EXPECT().GetDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200)

	//Act
	w := performRequest(router, "GET", "/memos/123?access_code="+validToken, nil)

	var response entity.MemoResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetRecentMemosAuthorized(t *testing.T) {
	//Arrange
	mockRepo, mockDataService := setupService(t)
	mockRepo.EXPECT().GetByDateRange(mock.Anything, mock.Anything, mock.Anything).Return([]entity.Memo{memoGregorian, memoLunar}, nil)
	mockDataService.EXPECT().GetDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200)
	mockDataService.EXPECT().GetLunarDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200)

	//Act
	w := performRequest(router, "GET", "/memos/recent?start=123&end=234&access_code="+validToken, nil)

	var response []entity.MemoResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestAddMemoAuthorized(t *testing.T) {
	//Arrange
	mockRepo, _ := setupService(t)
	mockRepo.EXPECT().Add(mock.AnythingOfType("*entity.Memo")).Return("123", nil)

	//Act
	w := performRequest(router, "PUT", "/memos/?access_code="+validToken, memoRequest)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}

func TestUpdateMemoByIdAuthorized(t *testing.T) {
	//Arrange
	mockRepo, _ := setupService(t)
	mockRepo.EXPECT().Update(mock.AnythingOfType("entity.Memo")).Return(nil)

	//Act
	w := performRequest(router, "POST", "/memos/123?access_code="+validToken, memoRequest)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}

func TestRemoveMemoAuthorized(t *testing.T) {
	//Arrange
	mockRepo, _ := setupService(t)
	mockRepo.EXPECT().Delete(mock.Anything).Return(nil)

	//Act
	w := performRequest(router, "DELETE", "/memos/123?access_code="+validToken, memoRequest)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
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

func setupService(t *testing.T) (*mocks.MemoRepo, *mocks.DateInterface) {
	gin.SetMode(gin.TestMode)

	mockRepo := mocks.NewMemoRepo(t)
	dateService := mocks.NewDateInterface(t)
	service := &memo.MemoApi{Repo: mockRepo, DateService: dateService}

	apiBoot = &ApiBoot{}
	apiBoot.MemoApi = service

	apiBoot.Registry = di.InitialMockRegistry()
	apiBoot.AuthorizationHandler = di.InitialAuthorizationMiddleware()
	apiBoot.ErrorHandler = di.InitialErrorMiddleware()

	router = gin.New()
	defineRoutes(router)

	token, _ := apiBoot.AuthorizationHandler.TokenService.NewToken("testuser", "test@email.com")
	validToken = token.Token

	return mockRepo, dateService
}
