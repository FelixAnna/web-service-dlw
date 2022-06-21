package memo

import (
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/memo-api/mocks"
	"github.com/stretchr/testify/assert"
)

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

func SkipTestAddMemo(t *testing.T) {
	mockRepo, _, service := setupService(t)

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	//mockRepo.
	//mockRepo(&entity.Memo{}).

	//need mock gin.Context.Writer
	service.AddMemo(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockRepo.AssertExpectations(t)
}
