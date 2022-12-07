package date

import (
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/FelixAnna/web-service-dlw/date-api/mocks"
	"github.com/stretchr/testify/assert"
	mockit "github.com/stretchr/testify/mock"
)

func setupService() (*mocks.MockCarbonService, *DateApi) {
	mockService := &mocks.MockCarbonService{}
	service := ProvideDateApi(mockService)

	return mockService, service
}

func TestProvideDateApi(t *testing.T) {
	_, service := setupService()

	assert.NotNil(t, service)
	assert.NotNil(t, service.CarbonService)
}

func TestGetMonthDateDefault(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockService.On("GetMonthDate", mockit.Anything).Return([]entity.DLWDate{})

	//need mock gin.Context.Writer
	service.GetMonthDate(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}

func TestGetMonthDate(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "date=20200505"})
	mockService.On("GetMonthDate", mockit.Anything).Return([]entity.DLWDate{})

	//need mock gin.Context.Writer
	service.GetMonthDate(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}

func TestToCarbonDateDefault(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{})
	mockService.On("ToCarbonDate", mockit.Anything).Return(&entity.DLWDate{})

	//need mock gin.Context.Writer
	service.ToCarbonDate(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}

func TestToCarbonDate(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "date=20200505"})
	mockService.On("ToCarbonDate", mockit.Anything).Return(&entity.DLWDate{})

	//need mock gin.Context.Writer
	service.ToCarbonDate(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}

func TestGetDateDistanceInvalidStart(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=&end=20200505"})
	mockService.On("GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetDateDistanceInvalidEnd(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=20200505&end="})
	mockService.On("GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)

	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetDateDistance(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=20200505&end=20200101"})
	mockService.On("GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(int64(1), int64(2))

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}

func TestGetLunarDateDistanceInvalidStart(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=&end=20200505"})
	mockService.On("GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetLunarDateDistanceInvalidEnd(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=20200505&end="})
	mockService.On("GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)

	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetLunarDateDistance(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext(&commonmock.Parameter{Query: "start=20200505&end=20200101"})
	mockService.On("GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(int64(1), int64(2))
	mockService.On("ToCarbonDate", mockit.Anything).Return(&entity.DLWDate{})

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}
