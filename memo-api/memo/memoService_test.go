package memo

import (
	"errors"
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/entity"
	"github.com/FelixAnna/web-service-dlw/memo-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func setupService(t *testing.T) (*mocks.MemoRepo, *mocks.DateInterface, *MemoApi) {
	mockRepo := mocks.NewMemoRepo(t)
	dateService := mocks.NewDateInterface(t)
	service := &MemoApi{mockRepo, dateService}

	return mockRepo, dateService, service
}

func TestProvideMemoApi(t *testing.T) {
	mockRepo, mockFileService, service := setupService(t)

	assert.NotNil(t, mockRepo)
	assert.NotNil(t, mockFileService)
	assert.NotNil(t, service)
}

func TestAddMemo(t *testing.T) {
	mockRepo, _, service := setupService(t)
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: &memoRequest})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().Add(mock.AnythingOfType("*entity.Memo")).Return("123", nil)

	service.AddMemo(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.Equal(t, writer.Body.String(), "\"Memo 123 created!\"")
}

func TestAddMemoInvalidArgs(t *testing.T) {
	_, _, service := setupService(t)
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: &memoRequestInvalid})
	ctx.Keys = map[string]any{"userId": "anyuser"}

	service.AddMemo(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	assert.NotNil(t, writer.Body.String())
}

func TestAddMemoFailedToAddToRepo(t *testing.T) {
	mockRepo, _, service := setupService(t)
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: &memoRequest})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().Add(mock.AnythingOfType("*entity.Memo")).Return("", errors.New("any error"))

	service.AddMemo(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.NotNil(t, writer.Body.String())
}

func TestGetMemosBy(t *testing.T) {
	mockRepo, mockDataService, service := setupService(t)
	params := map[string]string{"id": "123"}

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: params})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().GetById(mock.Anything).Return(&memoGregorian, nil)
	mockDataService.EXPECT().GetDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200, 20000101)

	service.GetMemoById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetMemosByLunar(t *testing.T) {
	mockRepo, mockDataService, service := setupService(t)
	params := map[string]string{"id": "123"}

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: params})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().GetById(mock.Anything).Return(&memoLunar, nil)
	mockDataService.EXPECT().GetLunarDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200, 20000101)

	service.GetMemoById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetMemosByIdNotFound(t *testing.T) {
	mockRepo, _, service := setupService(t)
	params := map[string]string{"id": "123"}

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: params})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().GetById(mock.Anything).Return(nil, errors.New("any error"))

	service.GetMemoById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusNotFound)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestGetMemosByUserId(t *testing.T) {
	mockRepo, mockDataService, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().GetByUserId(mock.Anything).Return([]entity.Memo{memoGregorian, memoLunar}, nil)
	mockDataService.EXPECT().GetDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200, 20000101)
	mockDataService.EXPECT().GetLunarDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200, 20000101)

	service.GetMemosByUserId(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestGetMemosByUserIdNotFound(t *testing.T) {
	mockRepo, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().GetByUserId(mock.Anything).Return(nil, errors.New("any error"))

	service.GetMemosByUserId(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestGetRecentMemos(t *testing.T) {
	mockRepo, mockDataService, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=123&end=456"})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().GetByDateRange(mock.Anything, mock.Anything, mock.Anything).Return([]entity.Memo{memoGregorian, memoLunar}, nil)
	mockDataService.EXPECT().GetDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200, 20000101)
	mockDataService.EXPECT().GetLunarDistance(mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(-100, 200, 20000101)

	service.GetRecentMemos(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.NotEmpty(t, writer.Body.String())
}

func TestGetRecentMemosNotFound(t *testing.T) {
	mockRepo, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=123&end=456"})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().GetByDateRange(mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("any error"))

	service.GetRecentMemos(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestUpdateMemoById(t *testing.T) {
	mockRepo, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"id": "123"}, Body: memoRequest})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().Update(mock.AnythingOfType("entity.Memo")).Return(nil)

	service.UpdateMemoById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.Equal(t, writer.Body.String(), "\"Memo 123 updated!\"")
}

func TestUpdateMemoByIdInvalid(t *testing.T) {
	_, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"id": "123"}, Body: memoRequestInvalid})
	ctx.Keys = map[string]any{"userId": "anyuser"}

	service.UpdateMemoById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	assert.NotEmpty(t, writer.Body.String())
}

func TestUpdateMemoByIdError(t *testing.T) {
	mockRepo, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"id": "123"}, Body: memoRequest})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().Update(mock.AnythingOfType("entity.Memo")).Return(errors.New("any error"))

	service.UpdateMemoById(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}

func TestRemoveMemo(t *testing.T) {
	mockRepo, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"id": "123"}})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().Delete(mock.Anything).Return(nil)

	service.RemoveMemo(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.Equal(t, writer.Body.String(), "\"Memo 123 deleted!\"")
}

func TestRemoveMemoError(t *testing.T) {
	mockRepo, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"id": "123"}})
	ctx.Keys = map[string]any{"userId": "anyuser"}
	mockRepo.EXPECT().Delete(mock.Anything).Return(errors.New("any error"))

	service.RemoveMemo(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusInternalServerError)
	assert.Equal(t, writer.Body.String(), "any error")
}
