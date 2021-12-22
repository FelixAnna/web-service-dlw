package services

import (
	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/golang-module/carbon/v2"
)

var CarbonTimeMap map[int]int
var LunarTimeMap map[int]int

const start = 19500101
const end = 20501231

func init() {
	//init 1901-2050 carbon and lunar maps
	length := (start - end) / 10000 * 365
	CarbonTimeMap = make(map[int]int, length)
	LunarTimeMap = make(map[int]int, length)
	initMap()
}

func initMap() {
	startCarbon := getCarbonDate(start)
	endCarbon := getCarbonDate(end)

	i, j := 0, 0
	for startCarbon.Lte(*endCarbon) {
		lunar := startCarbon.Lunar()

		carbonKey := startCarbon.Year()*10000 + startCarbon.Month()*100 + startCarbon.Day()
		lunarKey := lunar.Year()*100000 + lunar.Month()*1000 + lunar.Day()*10
		if lunar.IsLeapMonth() {
			lunarKey += 1
		}

		CarbonTimeMap[carbonKey] = i
		LunarTimeMap[lunarKey] = j

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
func GetCarbonDistanceWithCacheAside(startDate, targetDate int) (before, after int64) {
	targetValue, ok := CarbonTimeMap[targetDate]
	if ok {
		_, startMonthDay := startDate/10000, startDate%10000
		targetYear, targetMonthDay := targetDate/10000, targetDate%10000

		if startMonthDay < targetMonthDay {
			//targetYear + startMonthDay
			//targetYear+1 + startMonthDay
			preDate, nextDate := targetYear*10000+startMonthDay, (targetYear+1)*10000+startMonthDay
			preDateValue, okPre := CarbonTimeMap[preDate]
			nextDateValue, okNext := CarbonTimeMap[nextDate]
			if okPre && okNext {
				before = int64(preDateValue) - int64(targetValue)
				after = int64(nextDateValue) - int64(targetValue)
				return
			}

		} else if startMonthDay > targetMonthDay {
			//targetYear-1 + startMonthDay
			//targetYear + startMonthDay
			preDate, nextDate := (targetYear-1)*10000+startMonthDay, targetYear*10000+startMonthDay
			preDateValue, okPre := CarbonTimeMap[preDate]
			nextDateValue, okNext := CarbonTimeMap[nextDate]
			if okPre && okNext {
				before = int64(preDateValue) - int64(targetValue)
				after = int64(nextDateValue) - int64(targetValue)
				return
			}

		} else {
			return 0, 0
		}
	}

	return getCarbonDistance(startDate, targetDate)
}

/*
	Get distance of 2 datetime (will convert to lunar) in days
	fist search in cache;
	if not found, calculate manually.
*/
func GetLunarDistanceWithCacheAside(startDate, targetDate int) (before, after int64) {
	startCarbon := getCarbonDate(startDate)
	targetCarbon := getCarbonDate(targetDate)

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

	targetValue, ok := LunarTimeMap[targetDate]
	if ok {
		_, startMonthDay, _ := startDate/100000, (startDate%100000)/10, startDate%10
		targetYear, targetMonthDay, _ := targetDate/100000, (targetDate%100000)/10, targetDate%10

		if startMonthDay < targetMonthDay {
			//targetYear + startMonthDay
			//targetYear+1 + startMonthDay
			preDate, nextDate := targetYear*100000+startMonthDay*10, (targetYear+1)*100000+startMonthDay*10

			preDateFinal, nextDateFinal := getLunarCacheValue(preDate, nextDate)
			if preDateFinal > 0 && nextDateFinal > 0 {
				before = int64(preDateFinal) - int64(targetValue)
				after = int64(nextDateFinal) - int64(targetValue)
				return
			}

		} else if startMonthDay > targetMonthDay {
			//targetYear-1 + startMonthDay
			//targetYear + startMonthDay
			preDate, nextDate := (targetYear-1)*10000+startMonthDay, targetYear*10000+startMonthDay
			preDateFinal, nextDateFinal := getLunarCacheValue(preDate, nextDate)
			if preDateFinal > 0 && nextDateFinal > 0 {
				before = int64(preDateFinal) - int64(targetValue)
				after = int64(nextDateFinal) - int64(targetValue)
				return
			}

		} else {
			return 0, 0
		}
	}

	return getLunarDistance(startDate, targetDate)
}

func GetMonthDate(todayDate int) []entity.DLWDate {
	todayCarbon := getCarbonDate(todayDate)

	firstCarbon := getCarbonDate(todayDate/100*100 + 1)
	lastCarbon := getCarbonDate(todayDate/100*100 + todayCarbon.DaysInMonth())

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
func getLunarCacheValue(preDate, nextDate int) (int, int) {
	preDateFinal := 0
	preDateLeapValue, okPreLeap := LunarTimeMap[preDate+1]
	if okPreLeap {
		preDateFinal = preDateLeapValue

	} else {
		preDateValue, okPre := LunarTimeMap[preDate]
		if okPre {
			preDateFinal = preDateValue
		}
	}

	nextDateFinal := 0
	nextDateValue, okNext := LunarTimeMap[nextDate]
	if okNext {
		nextDateFinal = nextDateValue
	} else {
		nextDateLeapValue, okNextLeap := LunarTimeMap[nextDate+1]
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
func getCarbonDistance(startDate, targetDate int) (before, after int64) {

	startCarbon := getCarbonDate(startDate)
	targetCarbon := getCarbonDate(targetDate)

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
func getLunarDistance(startDate, targetDate int) (before, after int64) {
	startCarbon := getCarbonDate(startDate)
	targetCarbon := getCarbonDate(targetDate)

	before = getLunarDistanceOneWay(startCarbon, targetCarbon, false)
	after = getLunarDistanceOneWay(startCarbon, targetCarbon, true)
	return
}

func getLunarDistanceOneWay(startCarbon, targetCarbon *carbon.Carbon, forward bool) int64 {
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

func getCarbonDate(date int) *carbon.Carbon {
	carbonDate := carbon.CreateFromDate(date/10000, (date%10000)/100, date%100)
	return &carbonDate
}
