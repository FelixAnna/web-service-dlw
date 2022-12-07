package mocks

import (
	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/stretchr/testify/mock"
)

type MockCarbonService struct {
	mock.Mock
}

func (service *MockCarbonService) GetCarbonDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64) {
	args := service.Called(alignToDate, targetDate)
	return args.Get(0).(int64), args.Get(1).(int64)
}

func (service *MockCarbonService) GetLunarDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64) {
	args := service.Called(alignToDate, targetDate)
	return args.Get(0).(int64), args.Get(1).(int64)
}

func (service *MockCarbonService) GetMonthDate(todayDate int) []entity.DLWDate {
	args := service.Called(todayDate)
	return args.Get(0).([]entity.DLWDate)
}

func (service *MockCarbonService) ToCarbonDate(todayDate int) *entity.DLWDate {
	args := service.Called(todayDate)
	return args.Get(0).(*entity.DLWDate)
}
