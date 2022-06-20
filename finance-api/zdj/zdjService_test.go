package zdj

import (
	"net/http"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/filesystem"
	commonmock "github.com/FelixAnna/web-service-dlw/common/mock"
	"github.com/FelixAnna/web-service-dlw/finance-api/mock"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	"github.com/stretchr/testify/assert"
)

func setupService() (*mock.ZdjMockRepo, filesystem.FileInterface, *ZdjApi) {
	mockRepo := &mock.ZdjMockRepo{}
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

	ctx, writer := commonmock.GetGinContext("", map[string][]string{})
	criteria := &entity.Criteria{Page: 1, Size: 20}
	mockRepo.On("Search", criteria).Return([]entity.Zhidaojia{}, nil)

	//need mock gin.Context.Writer
	service.GetAll(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockRepo.AssertExpectations(t)
}

func SkipTestSerach(t *testing.T) {
	mockRepo, _, service := setupService()

	ctx, writer := commonmock.GetGinContext("", map[string][]string{})
	criteria := &entity.Criteria{Page: 1, Size: 20}
	mockRepo.On("Search", criteria).Return([]entity.Zhidaojia{}, nil)

	//need mock gin.Context.Writer
	service.Search(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockRepo.AssertExpectations(t)
}
