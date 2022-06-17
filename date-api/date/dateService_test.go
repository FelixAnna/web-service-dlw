package date

import (
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mock"
	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/FelixAnna/web-service-dlw/date-api/mock"
	"github.com/stretchr/testify/assert"
	mockit "github.com/stretchr/testify/mock"
)

func setupService() (*mock.MockCarbonService, *DateApi) {
	mockService := &mock.MockCarbonService{}
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

	ctx, writer := commonmock.GetGinContext("", map[string][]string{})
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

	ctx, writer := commonmock.GetGinContext("date=20200505", map[string][]string{})
	mockService.On("GetMonthDate", mockit.Anything).Return([]entity.DLWDate{})

	//need mock gin.Context.Writer
	service.GetMonthDate(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}

func TestGetDateDistanceInvalidStart(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext("start=&end=20200505", map[string][]string{})
	mockService.On("GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetDateDistanceInvalidEnd(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext("start=20200505&end=", map[string][]string{})
	mockService.On("GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)

	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetDateDistance(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext("start=20200505&end=20200101", map[string][]string{})
	mockService.On("GetCarbonDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(int64(1), int64(2))

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}

func TestGetLunarDateDistanceInvalidStart(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext("start=&end=20200505", map[string][]string{})
	mockService.On("GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetLunarDateDistanceInvalidEnd(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext("start=20200505&end=", map[string][]string{})
	mockService.On("GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(mockit.Anything, mockit.Anything)

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)

	assert.Equal(t, writer.Code, http.StatusBadRequest)
	mockService.AssertNotCalled(t, "GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything)
}

func TestGetLunarDateDistance(t *testing.T) {
	mockService, service := setupService()

	ctx, writer := commonmock.GetGinContext("start=20200505&end=20200101", map[string][]string{})
	mockService.On("GetLunarDistanceWithCacheAside", mockit.Anything, mockit.Anything).Return(int64(1), int64(2))

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
	mockService.AssertExpectations(t)
}
