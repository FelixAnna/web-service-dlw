package mock

import "github.com/FelixAnna/web-service-dlw/date-api/date/entity"

type MockCarbonService struct{}

func (service *MockCarbonService) GetCarbonDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64) {
	return 100, 100
}

func (service *MockCarbonService) GetLunarDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64) {
	return 100, 100
}

func (service *MockCarbonService) GetMonthDate(todayDate int) []entity.DLWDate {
	return []entity.DLWDate{}
}
