package services

import (
	"time"

	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/golang-module/carbon/v2"
	"github.com/google/wire"
)

var start, end int

func init() {
	start = (time.Now().Year()-100)*10000 + 101
	end = (time.Now().Year()+100)*10000 + 1231
}

var CarbonSet = wire.NewSet(ProvideCarbonInMemoryService, wire.Bind(new(CarbonService), new(*CarbonInMemoryService)))

type CarbonInMemoryService struct {
	CarbonTimeMap map[int]int
	LunarTimeMap  map[int]int
}

//provide for wire
func ProvideCarbonInMemoryService() *CarbonInMemoryService {
	carbonService := CarbonInMemoryService{}
	//init 1901-2050 carbon and lunar maps
	length := (start - end) / 10000 * 365
	carbonService.CarbonTimeMap = make(map[int]int, length)
	carbonService.LunarTimeMap = make(map[int]int, length)
	carbonService.initMap(start, end)

	return &carbonService
}

func (c *CarbonInMemoryService) initMap(start, end int) {
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
func (c *CarbonInMemoryService) GetCarbonDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64) {
	targetValue, ok := c.CarbonTimeMap[targetDate]
	if ok {
		_, alignToMonthDay := alignToDate/10000, alignToDate%10000
		targetYear, targetMonthDay := targetDate/10000, targetDate%10000

		if alignToMonthDay < targetMonthDay {
			//targetYear + alignToMonthDay
			//targetYear+1 + alignToMonthDay
			preDate, nextDate := targetYear*10000+alignToMonthDay, (targetYear+1)*10000+alignToMonthDay
			preDateValue, okPre := c.CarbonTimeMap[preDate]
			nextDateValue, okNext := c.CarbonTimeMap[nextDate]
			if okPre && okNext {
				before = int64(preDateValue) - int64(targetValue)
				after = int64(nextDateValue) - int64(targetValue)
				return
			}

		} else if alignToMonthDay > targetMonthDay {
			//targetYear-1 + alignToMonthDay
			//targetYear + alignToMonthDay
			preDate, nextDate := (targetYear-1)*10000+alignToMonthDay, targetYear*10000+alignToMonthDay
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

	return c.getCarbonDistance(alignToDate, targetDate)
}

/*
GetLunarDistanceWithCacheAside - Get distance of 2 datetime (will convert to lunar) in days
	fist search in cache;
	if not found, calculate manually.
*/
func (c *CarbonInMemoryService) GetLunarDistanceWithCacheAside(alignToDate, targetDate int) (before, after int64) {
	alignToCarbon := c.getCarbonDate(alignToDate)
	targetCarbon := c.getCarbonDate(targetDate)

	alignToLunarDate := alignToCarbon.Lunar()
	targetLunarDate := targetCarbon.Lunar()

	alignToDate = alignToLunarDate.Year()*100000 + alignToLunarDate.Month()*1000 + alignToLunarDate.Day()*10
	targetDate = targetLunarDate.Year()*100000 + targetLunarDate.Month()*1000 + targetLunarDate.Day()*10

	if alignToLunarDate.IsLeapMonth() {
		alignToDate += 1
	}
	if targetLunarDate.IsLeapMonth() {
		targetDate += 1
	}

	targetValue, ok := c.LunarTimeMap[targetDate]
	if ok {
		_, alignToMonthDay, _ := alignToDate/100000, (alignToDate%100000)/10, alignToDate%10
		targetYear, targetMonthDay, _ := targetDate/100000, (targetDate%100000)/10, targetDate%10

		if alignToMonthDay < targetMonthDay {
			//targetYear + alignToMonthDay
			//targetYear+1 + alignToMonthDay
			preDate, nextDate := targetYear*100000+alignToMonthDay*10, (targetYear+1)*100000+alignToMonthDay*10

			preDateFinal, nextDateFinal := c.getLunarCacheValue(preDate, nextDate)
			if preDateFinal > 0 && nextDateFinal > 0 {
				before = int64(preDateFinal) - int64(targetValue)
				after = int64(nextDateFinal) - int64(targetValue)
				return
			}

		} else if alignToMonthDay > targetMonthDay {
			//targetYear-1 + alignToMonthDay
			//targetYear + alignToMonthDay
			preDate, nextDate := (targetYear-1)*100000+alignToMonthDay*10, targetYear*100000+alignToMonthDay*10
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

	return c.getLunarDistance(alignToDate, targetDate)
}

func (c *CarbonInMemoryService) GetMonthDate(todayDate int) []entity.DLWDate {
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
func (c *CarbonInMemoryService) getLunarCacheValue(preDate, nextDate int) (int, int) {
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
getCarbonDistance - Get the distance between alignToDate and targetDate (ignore year)
	Suppose target date is now,
	alignToDate is the day we want to align,

	return how many days before and how many days later if alignToDate in MMdd (same month and day)
*/
func (c *CarbonInMemoryService) getCarbonDistance(alignToDate, targetDate int) (before, after int64) {

	alignToCarbon := c.getCarbonDate(alignToDate)
	targetCarbon := c.getCarbonDate(targetDate)

	diffYear := alignToCarbon.DiffInYears(*targetCarbon)
	alignToCarbonThisYear := alignToCarbon.AddYears(int(diffYear))
	diffDays := targetCarbon.DiffInDays(alignToCarbonThisYear)

	if diffDays < 0 { //target after alignToDate - n days before were alignToDate in MMdd, then find m days later when it will be alignToDate in MMdd again
		before = diffDays

		alignToCarbonNextYear := alignToCarbonThisYear.AddYear()
		after = targetCarbon.DiffInDays(alignToCarbonNextYear)
	} else if diffDays > 0 { //target before alignToDate - n days later will be alignToDate in MMdd, then find m days before when it was alignToDate in MMdd
		after = diffDays

		alignToCarbonPreYear := alignToCarbonThisYear.SubYear()
		before = targetCarbon.DiffInDays(alignToCarbonPreYear)
	} else {
		return 0, 0
	}

	return
}

/*
getLunarDistance - Get the distance between alignToDate and targetDate (ignore year) after convert them to lunar
	Suppose targetDate is now,
	alignToDate is the day we want to align,

	return how many days before and how many days later if same as alignToDate in MMdd (same month and day)
*/
func (c *CarbonInMemoryService) getLunarDistance(alignToDate, targetDate int) (before, after int64) {
	alignToCarbon := c.getCarbonDate(alignToDate)
	targetCarbon := c.getCarbonDate(targetDate)

	before = c.getLunarDistanceOneWay(alignToCarbon, targetCarbon, false)
	after = c.getLunarDistanceOneWay(alignToCarbon, targetCarbon, true)
	return
}

func (c *CarbonInMemoryService) getLunarDistanceOneWay(alignToCarbon, targetCarbon *carbon.Carbon, forward bool) int64 {
	distance := 0
	alignToLunarDate := alignToCarbon.Lunar()
	targetLunarDate := targetCarbon.Lunar()

	alignToMMdd := alignToLunarDate.Month()*100 + alignToLunarDate.Day()
	targetMMdd := targetLunarDate.Month()*100 + targetLunarDate.Day()
	for alignToMMdd != targetMMdd {
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

func (c *CarbonInMemoryService) getCarbonDate(date int) *carbon.Carbon {
	carbonDate := carbon.CreateFromDate(date/10000, (date%10000)/100, date%100)
	return &carbonDate
}
