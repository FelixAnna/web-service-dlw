package services

import "github.com/FelixAnna/web-service-dlw/date-api/date/entity"

type CarbonService interface {
	GetCarbonDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64)
	GetLunarDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64)
	GetMonthDate(todayDate int) []entity.DLWDate
}
