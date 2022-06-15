package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var service *CarbonService

func init() {
	service = ProvideCarbonService()
}

func TestProvideCarbonService(t *testing.T) {
	//Act
	//service := ProvideCarbonService()

	//Assert
	assert.NotNil(t, service)
	assert.True(t, len(service.CarbonTimeMap) > 0)
	assert.True(t, len(service.LunarTimeMap) > 0)
}

func TestGetCarbonDistanceWithCacheAsideInCacheBefore(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20200101, 20200505

	//Act
	before, after := service.GetCarbonDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -125, before)
	assert.EqualValues(t, 241, after)
}

func TestGetCarbonDistanceWithCacheAsideInCacheAfter(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20190701, 20200505

	//Act
	before, after := service.GetCarbonDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -309, before)
	assert.EqualValues(t, 57, after)
}

func TestGetCarbonDistanceWithCacheAsideRealTimeCalcBefore(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20790101, 20800505

	//Act
	before, after := service.GetCarbonDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -125, before)
	assert.EqualValues(t, 241, after)
}

func TestGetCarbonDistanceWithCacheAsideRealTimeCalcAfter(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20790701, 20800505

	//Act
	before, after := service.GetCarbonDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -309, before)
	assert.EqualValues(t, 57, after)
}

func TestGetLunarDistanceWithCacheAsideBefore(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20200101, 20200505

	//Act
	before, after := service.GetLunarDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -155, before)
	assert.EqualValues(t, 229, after)
}

func TestGetLunarDistanceWithCacheAsideAfter(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20200701, 20200505

	//Act
	before, after := service.GetLunarDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -357, before)
	assert.EqualValues(t, 27, after)
}

func TestGetLunarDistanceWithCacheAsideRealTimeCalcBefore(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20790101, 20800505

	//Act
	before, after := service.GetLunarDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -136, before)
	assert.EqualValues(t, 248, after)
}

func TestGetLunarDistanceWithCacheAsideRealTimeCalcAfter(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20790701, 20800505

	//Act
	before, after := service.GetLunarDistanceWithCacheAside(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -309, before)
	assert.EqualValues(t, 75, after)
}

func TestGetMonthDate(t *testing.T) {
	//Arrange
	todayDate := 20800505

	//Act
	dateList := service.GetMonthDate(todayDate)

	//Assert
	assert.GreaterOrEqual(t, len(dateList), 31)
}

func SkipTestGetCarbonDate(t *testing.T) {
	//Arrange
	date := 20200101

	//Act
	result := service.getCarbonDate(date)

	//Assert
	assert.NotNil(t, result)
	assert.Equal(t, result.Year(), 2020)
	assert.Equal(t, result.Month(), 1)
	assert.Equal(t, result.Day(), 1)

	assert.Equal(t, result.Lunar().Year(), 2019)
	assert.Equal(t, result.Lunar().Month(), 12)
	assert.Equal(t, result.Lunar().Day(), 7)
}

func SkipTestGetLunarDistanceOneWayForward(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20200101, 20200505
	alignToCarbon := service.getCarbonDate(alignToDate)
	targetCarbon := service.getCarbonDate(targetDate)

	//Act
	result := service.getLunarDistanceOneWay(alignToCarbon, targetCarbon, true)

	//Assert
	assert.EqualValues(t, 259, result)
}

func SkipTestGetLunarDistanceOneWayAfterword(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20200101, 20200505
	alignToCarbon := service.getCarbonDate(alignToDate)
	targetCarbon := service.getCarbonDate(targetDate)

	//Act
	result := service.getLunarDistanceOneWay(alignToCarbon, targetCarbon, false)

	//Assert
	assert.EqualValues(t, -125, result)
}

func SkipTestGetLunarDistance(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20200101, 20200505

	//Act
	before, after := service.getLunarDistance(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -125, before)
	assert.EqualValues(t, 259, after)
}

func SkipTestGetCarbonDistanceBefore(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20200101, 20200505

	//Act
	before, after := service.getCarbonDistance(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -125, before)
	assert.EqualValues(t, 241, after)
}

func SkipTestGetCarbonDistanceAfter(t *testing.T) {
	//Arrange
	alignToDate, targetDate := 20190701, 20200505

	//Act
	before, after := service.getCarbonDistance(alignToDate, targetDate)

	//Assert
	assert.EqualValues(t, -309, before)
	assert.EqualValues(t, 57, after)
}
