package users

import (
	"errors"
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/user-api/mocks"
	"github.com/FelixAnna/web-service-dlw/user-api/users/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userModel entity.User
var address []entity.Address

var userModelInvalid entity.User
var addressInvalid []entity.Address

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

	userModelInvalid = entity.User{}
	addressInvalid = []entity.Address{{Country: ""}}
}

func setupService(t *testing.T) (*mocks.UserRepo, *UserApi) {
	mockRepo := mocks.NewUserRepo(t)
	service := &UserApi{mockRepo}

	return mockRepo, service
}

func TestUserApi(t *testing.T) {
	mockRepo, service := setupService(t)

	assert.NotNil(t, mockRepo)
	assert.NotNil(t, service)
}

func TestGetAllUsersFailed(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockRepo.EXPECT().GetAll().Return(nil, errors.New("any error"))

	service.GetAllUsers(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestGetAllUsers(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockRepo.EXPECT().GetAll().Return([]entity.User{userModel}, nil)

	service.GetAllUsers(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestGetUserByEmailFailed(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockRepo.EXPECT().GetByEmail(mock.Anything).Return(nil, errors.New("any error"))

	service.GetUserByEmail(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusNotFound)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestGetUserByEmail(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockRepo.EXPECT().GetByEmail(mock.Anything).Return(&userModel, nil)

	service.GetUserByEmail(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestGetUserByIdFailed(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockRepo.EXPECT().GetById(mock.Anything).Return(nil, errors.New("any error"))

	service.GetUserById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusNotFound)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestGetUserById(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockRepo.EXPECT().GetById(mock.Anything).Return(&userModel, nil)

	service.GetUserById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestUpdateUserBirthdayByIdFailed(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"userId": "any"}, Query: "birthday=any"})
	mockRepo.EXPECT().UpdateBirthday(mock.Anything, mock.Anything).Return(errors.New("any error"))

	service.UpdateUserBirthdayById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestUpdateUserBirthdayById(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"userId": "any"}, Query: "birthday=any"})
	mockRepo.EXPECT().UpdateBirthday(mock.Anything, mock.Anything).Return(nil)

	service.UpdateUserBirthdayById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestUpdateUserAddressByIdInvalid(t *testing.T) {
	_, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: addressInvalid})

	service.UpdateUserAddressById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	assert.NotEmpty(t, writer.Body.String())
}

func TestUpdateUserAddressByIdFailed(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: address})
	mockRepo.EXPECT().UpdateAddress(mock.Anything, mock.AnythingOfType("[]entity.Address")).Return(errors.New("any error"))

	service.UpdateUserAddressById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestUpdateUserAddressById(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: address})
	mockRepo.EXPECT().UpdateAddress(mock.Anything, mock.AnythingOfType("[]entity.Address")).Return(nil)

	service.UpdateUserAddressById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestAddUserInvalid(t *testing.T) {
	_, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: userModelInvalid})

	service.AddUser(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	assert.NotEmpty(t, writer.Body.String())
}

func TestAddUserFailed(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: userModel})
	mockRepo.EXPECT().Add(mock.AnythingOfType("*entity.User")).Return(nil, errors.New("any error"))

	service.AddUser(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestAddUser(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: userModel})
	mockRepo.EXPECT().Add(mock.AnythingOfType("*entity.User")).Return(&userModel.Id, nil)

	service.AddUser(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestRemoveUserFailed(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: userModel})
	mockRepo.EXPECT().Delete(mock.Anything).Return(errors.New("any error"))

	service.RemoveUser(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestRemoveUser(t *testing.T) {
	mockRepo, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: userModel})
	mockRepo.EXPECT().Delete(mock.Anything).Return(nil)

	service.RemoveUser(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}
