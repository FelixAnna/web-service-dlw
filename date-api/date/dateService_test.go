package date

import (
	"net/http"
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mock"
	"github.com/FelixAnna/web-service-dlw/date-api/mock"
	"github.com/stretchr/testify/assert"
)

var service *DateApi

func init() {
	carbonService := mock.MockCarbonService{}
	service = ProvideDateApi(&carbonService)
}

func TestProvideDateApi(t *testing.T) {
	assert.NotNil(t, service)
	assert.NotNil(t, service.CarbonService)
}

func TestGetMonthDateDefault(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("", map[string][]string{})

	//need mock gin.Context.Writer
	service.GetMonthDate(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetMonthDate(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("date=20200505", map[string][]string{})

	//need mock gin.Context.Writer
	service.GetMonthDate(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetDateDistanceInvalidStart(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("start=&end=20200505", map[string][]string{})

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetDateDistanceInvalidEnd(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("start=20200505&end=", map[string][]string{})

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)

	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetDateDistance(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("start=20200505&end=20200101", map[string][]string{})

	service.GetDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}

func TestGetLunarDateDistanceInvalidStart(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("start=&end=20200505", map[string][]string{})

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetLunarDateDistanceInvalidEnd(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("start=20200505&end=", map[string][]string{})

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)

	assert.Equal(t, writer.Code, http.StatusBadRequest)
}

func TestGetLunarDateDistance(t *testing.T) {
	ctx, writer := commonmock.GetGinContext("start=20200505&end=20200101", map[string][]string{})

	service.GetLunarDateDistance(ctx)

	assert.NotNil(t, ctx)
	assert.NotNil(t, writer)
	assert.Equal(t, writer.Code, http.StatusOK)
}
