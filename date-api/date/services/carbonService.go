package services

import (
	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/golang-module/carbon/v2"
)

type CarbonService struct {
	CarbonTimeMap map[int]int
	LunarTimeMap  map[int]int
}

//provide for wire
func ProvideCarbonService() *CarbonService {
	const start = 19500101
	const end = 20501231

	carbonService := CarbonService{}
	//init 1901-2050 carbon and lunar maps
	length := (start - end) / 10000 * 365
	carbonService.CarbonTimeMap = make(map[int]int, length)
	carbonService.LunarTimeMap = make(map[int]int, length)
	carbonService.initMap(start, end)

	return &carbonService
}

func (c *CarbonService) initMap(start, end int) {
	startCarbon := c.getCarbonDate(start)
	endCarbon := c.getCarbonDate(end)

	i, j := 0, 0
	for startCarbon.Lte(*endCarbon) {
		lunar := startCarbon.Lunar()

		carbonKey := startCarbon.Year()*10000 + startCarbon.Month()*100 + startCarbon.Day()
		lunarKey := lunar.Year()*100000 + lunar.Month()*1000 + lunar.Day()*10
		if lunar.IsLeapMonth() {
			lunarKey += 1
		}

		c.CarbonTimeMap[carbonKey] = i
		c.LunarTimeMap[lunarKey] = j

		i, j = i+1, j+1

		startCarbonNew := startCarbon.AddDay()
		startCarbon = &startCarbonNew
	}
}

/*
	Get distance of 2 datetime in days
	fist search in cache;
	if not found, calculate manually.
*/
func (c *CarbonService) GetCarbonDistanceWithCacheAside(startDate, targetDate int) (before, after int64) {
	targetValue, ok := c.CarbonTimeMap[targetDate]
	if ok {
		_, startMonthDay := startDate/10000, startDate%10000
		targetYear, targetMonthDay := targetDate/10000, targetDate%10000

		if startMonthDay < targetMonthDay {
			//targetYear + startMonthDay
			//targetYear+1 + startMonthDay
			preDate, nextDate := targetYear*10000+startMonthDay, (targetYear+1)*10000+startMonthDay
			preDateValue, okPre := c.CarbonTimeMap[preDate]
			nextDateValue, okNext := c.CarbonTimeMap[nextDate]
			if okPre && okNext {
				before = int64(preDateValue) - int64(targetValue)
				after = int64(nextDateValue) - int64(targetValue)
				return
			}

		} else if startMonthDay > targetMonthDay {
			//targetYear-1 + startMonthDay
			//targetYear + startMonthDay
			preDate, nextDate := (targetYear-1)*10000+startMonthDay, targetYear*10000+startMonthDay
			preDateValue, okPre := c.CarbonTimeMap[preDate]
			nextDateValue, okNext := c.CarbonTimeMap[nextDate]
			if okPre && okNext {
				before = int64(preDateValue) - int64(targetValue)
				after = int64(nextDateValue) - int64(targetValue)
				return
			}

		} else {
			return 0, 0
		}
	}

	return c.getCarbonDistance(startDate, targetDate)
}

/*
	Get distance of 2 datetime (will convert to lunar) in days
	fist search in cache;
	if not found, calculate manually.
*/
func (c *CarbonService) GetLunarDistanceWithCacheAside(startDate, targetDate int) (before, after int64) {
	startCarbon := c.getCarbonDate(startDate)
	targetCarbon := c.getCarbonDate(targetDate)

	startLunarDate := startCarbon.Lunar()
	targetLunarDate := targetCarbon.Lunar()

	startDate = startLunarDate.Year()*100000 + startLunarDate.Month()*1000 + startLunarDate.Day()*10
	targetDate = targetLunarDate.Year()*100000 + targetLunarDate.Month()*1000 + targetLunarDate.Day()*10

	if startLunarDate.IsLeapMonth() {
		startDate += 1
	}
	if targetLunarDate.IsLeapMonth() {
		targetDate += 1
	}

	targetValue, ok := c.LunarTimeMap[targetDate]
	if ok {
		_, startMonthDay, _ := startDate/100000, (startDate%100000)/10, startDate%10
		targetYear, targetMonthDay, _ := targetDate/100000, (targetDate%100000)/10, targetDate%10

		if startMonthDay < targetMonthDay {
			//targetYear + startMonthDay
			//targetYear+1 + startMonthDay
			preDate, nextDate := targetYear*100000+startMonthDay*10, (targetYear+1)*100000+startMonthDay*10

			preDateFinal, nextDateFinal := c.getLunarCacheValue(preDate, nextDate)
			if preDateFinal > 0 && nextDateFinal > 0 {
				before = int64(preDateFinal) - int64(targetValue)
				after = int64(nextDateFinal) - int64(targetValue)
				return
			}

		} else if startMonthDay > targetMonthDay {
			//targetYear-1 + startMonthDay
			//targetYear + startMonthDay
			preDate, nextDate := (targetYear-1)*10000+startMonthDay, targetYear*10000+startMonthDay
			preDateFinal, nextDateFinal := c.getLunarCacheValue(preDate, nextDate)
			if preDateFinal > 0 && nextDateFinal > 0 {
				before = int64(preDateFinal) - int64(targetValue)
				after = int64(nextDateFinal) - int64(targetValue)
				return
			}

		} else {
			return 0, 0
		}
	}

	return c.getLunarDistance(startDate, targetDate)
}

func (c *CarbonService) GetMonthDate(todayDate int) []entity.DLWDate {
	todayCarbon := c.getCarbonDate(todayDate)

	firstCarbon := c.getCarbonDate(todayDate/100*100 + 1)
	lastCarbon := c.getCarbonDate(todayDate/100*100 + todayCarbon.DaysInMonth())

	startCarbon := firstCarbon.AddDays(firstCarbon.DayOfWeek() * -1)
	endCarbon := lastCarbon.AddDays(6 - lastCarbon.DayOfWeek())

	result := make([]entity.DLWDate, 0)
	for startCarbon.Lte(endCarbon) {
		lunarCarbon := startCarbon.Lunar()

		ymd := startCarbon.Year()*10000 + startCarbon.Month()*100 + startCarbon.Day()
		lunar := lunarCarbon.ToDateString()
		lunarCarbon.Animal()
		item := entity.DLWDate{
			YMD:       ymd,
			Lunar:     lunar,
			Animal:    lunarCarbon.Animal(),
			LeapMonth: lunarCarbon.IsLeapMonth(),
			Today:     ymd == todayDate,
			WeekDay:   startCarbon.DayOfWeek(),
		}
		result = append(result, item)

		startCarbon = startCarbon.AddDay()
	}

	return result
}

/*
Get lunar cache value with consideration of Leap month
*/
func (c *CarbonService) getLunarCacheValue(preDate, nextDate int) (int, int) {
	preDateFinal := 0
	preDateLeapValue, okPreLeap := c.LunarTimeMap[preDate+1]
	if okPreLeap {
		preDateFinal = preDateLeapValue

	} else {
		preDateValue, okPre := c.LunarTimeMap[preDate]
		if okPre {
			preDateFinal = preDateValue
		}
	}

	nextDateFinal := 0
	nextDateValue, okNext := c.LunarTimeMap[nextDate]
	if okNext {
		nextDateFinal = nextDateValue
	} else {
		nextDateLeapValue, okNextLeap := c.LunarTimeMap[nextDate+1]
		if okNextLeap {
			nextDateFinal = nextDateLeapValue
		}
	}

	return preDateFinal, nextDateFinal
}

/*
getCarbonDistance - Get the distance between startDate and targetDate (ignore year)
Suppose target date is now,
return how many days before and how many days later if startDate (same month and day)
*/
func (c *CarbonService) getCarbonDistance(startDate, targetDate int) (before, after int64) {

	startCarbon := c.getCarbonDate(startDate)
	targetCarbon := c.getCarbonDate(targetDate)

	diffYear := startCarbon.DiffInYears(*targetCarbon)
	startCarbonThisYear := startCarbon.AddYears(int(diffYear))
	diffDays := targetCarbon.DiffInDays(startCarbonThisYear)

	if diffDays < 0 { //target after start - n days before were start, then find m days later when it will be start again
		before = diffDays

		startCarbonNextYear := startCarbonThisYear.AddYear()
		after = targetCarbon.DiffInDays(startCarbonNextYear)
	} else if diffDays > 0 { //target before start - n days later will be start, then find m days before when it was start
		after = diffDays

		startCarbonPreYear := startCarbonThisYear.SubYear()
		before = targetCarbon.DiffInDays(startCarbonPreYear)
	} else {
		return 0, 0
	}

	return
}

/*
getLunarDistance - Get the distance between startDate and targetDate (ignore year) after convert them to lunar
Suppose target date is now,
return how many days before and how many days later if startDate (same month and day)
*/
func (c *CarbonService) getLunarDistance(startDate, targetDate int) (before, after int64) {
	startCarbon := c.getCarbonDate(startDate)
	targetCarbon := c.getCarbonDate(targetDate)

	before = c.getLunarDistanceOneWay(startCarbon, targetCarbon, false)
	after = c.getLunarDistanceOneWay(startCarbon, targetCarbon, true)
	return
}

func (c *CarbonService) getLunarDistanceOneWay(startCarbon, targetCarbon *carbon.Carbon, forward bool) int64 {
	distance := 0
	startLunarDate := startCarbon.Lunar()
	targetLunarDate := targetCarbon.Lunar()

	startMMdd := startLunarDate.Month()*100 + startLunarDate.Day()
	targetMMdd := targetLunarDate.Month()*100 + targetLunarDate.Day()
	for startMMdd != targetMMdd {
		if forward {
			targetCarbonNew := targetCarbon.AddDays(1)
			targetCarbon = &targetCarbonNew

			targetLunarDate = targetCarbon.Lunar()
			targetMMdd = targetLunarDate.Month()*100 + targetLunarDate.Day()
			distance += 1
		} else {
			targetCarbonNew := targetCarbon.AddDays(-1)
			targetCarbon = &targetCarbonNew

			targetLunarDate = targetCarbon.Lunar()
			targetMMdd = targetLunarDate.Month()*100 + targetLunarDate.Day()
			distance -= 1
		}
	}

	return int64(distance)
}

func (c *CarbonService) getCarbonDate(date int) *carbon.Carbon {
	carbonDate := carbon.CreateFromDate(date/10000, (date%10000)/100, date%100)
	return &carbonDate
}
