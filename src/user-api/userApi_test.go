package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixAnna/web-service-dlw/user-api/di"
	"github.com/FelixAnna/web-service-dlw/user-api/mocks"
	"github.com/FelixAnna/web-service-dlw/user-api/users"
	"github.com/FelixAnna/web-service-dlw/user-api/users/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userModel entity.User
var address []entity.Address

func init() {
	address = []entity.Address{
		{
			Country: "any",
			State:   "any",
			City:    "any",
			Details: "any",
		},
	}

	userModel = entity.User{
		Id:         "any",
		Name:       "any",
		AvatarUrl:  "any",
		Email:      "any@any.com",
		Birthday:   "any",
		Address:    address,
		CreateTime: "any",
	}
}

var router *gin.Engine
var validToken string

func TestRunning(t *testing.T) {
	//Arrange
	setupService(t)

	w := performRequest(router, "GET", "/status", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "running", w.Body.String())
}

func TestGetAllUsersUnAuthorized(t *testing.T) {
	//Arrange
	setupService(t)

	//Act
	w := performRequest(router, "GET", "/users/", nil)

	var response []entity.User
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotNil(t, err)
}

func TestGetAllUsersForbidden(t *testing.T) {
	//Arrange
	setupService(t)

	//Act
	w := performRequest(router, "GET", "/users/?access_code=123", nil)

	var response []entity.User
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.NotNil(t, err)
}

func TestGetAllUsersAuthorized(t *testing.T) {
	//Arrange
	mockRepo := setupService(t)
	mockRepo.EXPECT().GetAll().Return([]entity.User{userModel}, nil)

	//Act
	w := performRequest(router, "GET", "/users/?access_code="+validToken, nil)

	var response []entity.User
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetUserByEmailAuthorized(t *testing.T) {
	//Arrange
	mockRepo := setupService(t)
	mockRepo.EXPECT().GetByEmail(mock.Anything).Return(&userModel, nil)

	//Act
	w := performRequest(router, "GET", "/users/email/any@any.com?access_code="+validToken, nil)

	var response entity.User
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetUserByIdAuthorized(t *testing.T) {
	//Arrange
	mockRepo := setupService(t)
	mockRepo.EXPECT().GetById(mock.Anything).Return(&userModel, nil)

	//Act
	w := performRequest(router, "GET", "/users/123?access_code="+validToken, nil)

	var response entity.User
	err := json.Unmarshal(w.Body.Bytes(), &response)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestUpdateUserBirthdayByIdAuthorized(t *testing.T) {
	//Arrange
	mockRepo := setupService(t)
	mockRepo.EXPECT().UpdateBirthday(mock.Anything, mock.Anything).Return(nil)

	//Act
	w := performRequest(router, "POST", "/users/123?birthday=any&access_code="+validToken, nil)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}

func TestUpdateUserAddressByIdAuthorized(t *testing.T) {
	//Arrange
	mockRepo := setupService(t)
	mockRepo.EXPECT().UpdateAddress(mock.Anything, mock.AnythingOfType("[]entity.Address")).Return(nil)

	//Act
	w := performRequest(router, "POST", "/users/123/address?access_code="+validToken, address)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}

func TestAddUserAuthorized(t *testing.T) {
	//Arrange
	mockRepo := setupService(t)
	mockRepo.EXPECT().Add(mock.AnythingOfType("*entity.User")).Return(&userModel.Id, nil)

	//Act
	w := performRequest(router, "PUT", "/users/?access_code="+validToken, userModel)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}

func TestDeleteUserAuthorized(t *testing.T) {
	//Arrange
	mockRepo := setupService(t)
	mockRepo.EXPECT().Delete(mock.Anything).Return(nil)

	//Act
	w := performRequest(router, "DELETE", "/users/123?access_code="+validToken, address)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
}

func TestGetGinRouter(t *testing.T) {

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

func setupService(t *testing.T) *mocks.UserRepo {
	gin.SetMode(gin.TestMode)

	mockRepo := mocks.NewUserRepo(t)
	service := &users.UserApi{Repo: mockRepo}
	apiBoot = &ApiBoot{
		UserApi: service,
		//AuthGithubApi:              di.InitialGithubAuthApi(),
		//AuthGoogleApi:              di.InitialGoogleAuthApi(),
		AuthorizationHandler: di.InitialMockAuthorizationMiddleware(),
		ErrorHandler:         di.InitialErrorMiddleware(),
		Registry:             di.InitialMockRegistry(),
	}

	router = GetGinRouter()

	token, _ := apiBoot.AuthorizationHandler.TokenService.NewToken("testuser", "test@email.com")
	validToken = token.Token

	return mockRepo
}
