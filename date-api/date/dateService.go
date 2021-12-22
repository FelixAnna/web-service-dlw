package date

import (
	"net/http"
	"strconv"
	"time"

	"github.com/FelixAnna/web-service-dlw/date-api/date/entity"
	"github.com/FelixAnna/web-service-dlw/date-api/date/services"
	"github.com/gin-gonic/gin"
)

func GetMonthDate(c *gin.Context) {
	//ctx := context.Background()
	//generate state and return to client can stop CSRF
	date := c.Query("date")
	dateInt, err := strconv.Atoi(date)
	if err != nil {
		today := time.Now()
		dateInt = today.Year()*10000 + int(today.Month())*100 + today.Day()
	}
	dateList := services.GetMonthDate(dateInt)

	c.JSON(http.StatusOK, dateList)
}

func GetDateDistance(c *gin.Context) {
	//ctx := context.Background()
	//generate state and return to client can stop CSRF
	start := c.Query("start")
	end := c.Query("end")
	iStart, erStart := strconv.Atoi(start)
	iEnd, erEnd := strconv.Atoi(end)
	if erStart != nil || erEnd != nil {
		c.String(http.StatusBadRequest, "Not a number")
	}

	before, after := services.GetCarbonDistanceWithCacheAside(iStart, iEnd)

	distance := &entity.Distance{
		StartYMD:  iStart,
		TargetYMD: iEnd,
		Lunar:     false,
		Before:    before,
		After:     after,
	}
	c.JSON(http.StatusOK, distance)
}

func GetLunarDateDistance(c *gin.Context) {
	//ctx := context.Background()
	//generate state and return to client can stop CSRF
	start := c.Query("start")
	end := c.Query("end")
	iStart, erStart := strconv.Atoi(start)
	iEnd, erEnd := strconv.Atoi(end)
	if erStart != nil || erEnd != nil {
		c.String(http.StatusBadRequest, "Not a number")
	}

	before, after := services.GetLunarDistanceWithCacheAside(iStart, iEnd)

	distance := &entity.Distance{
		StartYMD:  iStart,
		TargetYMD: iEnd,
		Lunar:     true,
		Before:    before,
		After:     after,
	}

	c.JSON(http.StatusOK, distance)
}
