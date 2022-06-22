package zdj

import (
	"net/http"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/filesystem"
	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/finance-api/mocks"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	"github.com/stretchr/testify/assert"
	mockit "github.com/stretchr/testify/mock"
)

func setupService() (*mocks.ZdjMockRepo, filesystem.FileInterface, *ZdjApi) {
	mockRepo := &mocks.ZdjMockRepo{}
	mockFileService := &commonmock.MockFileService{}
	service := ProvideZdjApi(mockRepo, mockFileService)

	return mockRepo, mockFileService, &service
}

func TestProvideZdjApi(t *testing.T) {
	mockRepo, mockFileService, service := setupService()

	assert.NotNil(t, mockRepo)
	assert.NotNil(t, mockFileService)
	assert.NotNil(t, service)
}

func TestGetAll(t *testing.T) {
	mockRepo, _, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	criteria := &entity.Criteria{Page: 1, Size: 20}
	mockRepo.On("Search", criteria).Return([]entity.Zhidaojia{}, nil)

	//need mock gin.Context.Writer
	service.GetAll(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockRepo.AssertExpectations(t)
}

func TestSearchNil(t *testing.T) {
	mockRepo, _, service := setupService()

	criteria := &entity.Criteria{Page: 1, Size: 20}
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockRepo.On("Search", criteria).Return([]entity.Zhidaojia{}, nil)

	//need mock gin.Context.Writer
	service.Search(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockRepo.AssertNotCalled(t, "Search", criteria)
}

func TestSearch(t *testing.T) {
	mockRepo, _, service := setupService()

	criteria := &entity.Criteria{Page: 1, Size: 20}
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Body: criteria})
	mockRepo.On("Search", criteria).Return([]entity.Zhidaojia{}, nil)

	//need mock gin.Context.Writer
	service.Search(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockRepo.AssertExpectations(t)
}

func TestMemoryCostyDefault(t *testing.T) {
	_, _, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})

	//need mock gin.Context.Writer
	service.MemoryCosty(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestMemoryCosty(t *testing.T) {
	_, _, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "times=10000"})

	//need mock gin.Context.Writer
	service.MemoryCosty(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestDeleteInvalidId(t *testing.T) {
	mockRepo, _, service := setupService()

	query := "version=2021"
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: query, Params: map[string]string{"id": "123a"}})
	mockRepo.On("Delete", mockit.AnythingOfType("int"), mockit.AnythingOfType("int")).Return(nil)

	//need mock gin.Context.Writer
	service.Delete(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockRepo.AssertNotCalled(t, "Delete", mockit.AnythingOfType("int"), mockit.AnythingOfType("int"))
}

func TestDeleteInvalidVersion(t *testing.T) {
	mockRepo, _, service := setupService()

	query := "version=abc"
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: query, Params: map[string]string{"id": "123"}})
	mockRepo.On("Delete", mockit.AnythingOfType("int"), mockit.AnythingOfType("int")).Return(nil)

	//need mock gin.Context.Writer
	service.Delete(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockRepo.AssertNotCalled(t, "Delete", mockit.AnythingOfType("int"), mockit.AnythingOfType("int"))
}

func TestDeleteInvalidVersionNil(t *testing.T) {
	mockRepo, _, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Params: map[string]string{"id": "123"}})
	mockRepo.On("Delete", mockit.AnythingOfType("int"), mockit.AnythingOfType("int")).Return(nil)

	//need mock gin.Context.Writer
	service.Delete(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo, _, service := setupService()

	query := "version=2021"
	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: query, Params: map[string]string{"id": "123"}})
	mockRepo.On("Delete", mockit.AnythingOfType("int"), mockit.AnythingOfType("int")).Return(nil)

	//need mock gin.Context.Writer
	service.Delete(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockRepo.AssertExpectations(t)
}