package date

import (
	"testing"

	commonmock "github.com/FelixAnna/web-service-dlw/common/mock"
	mock "github.com/FelixAnna/web-service-dlw/date-api/mock"
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

func SkipTestGetMonthDate(t *testing.T) {
	ctx := commonmock.GetGinContext("date=20200505", map[string][]string{})

	//need mock gin.Context.Writer
	service.GetMonthDate(ctx)
}
