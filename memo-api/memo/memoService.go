package memo

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FelixAnna/web-service-dlw/memo-api/memo/entity"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/repository"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/services"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var MemoSet = wire.NewSet(wire.Struct(new(MemoApi), "*"))

type MemoApi struct {
	Repo        repository.MemoRepo
	DateService services.DateInterface
}

func (api *MemoApi) AddMemo(c *gin.Context) {
	userId := api.getUserIdFromContext(c)
	var request entity.MemoRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var new_memo entity.Memo = entity.Memo{
		Subject:     request.Subject,
		Description: request.Description,
		UserId:      userId,
		MonthDay:    request.MonthDay,
		StartYear:   request.StartYear,
		Lunar:       request.Lunar,
	}

	id, err := api.Repo.Add(&new_memo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Memo %v created!", *id))
}

func (api *MemoApi) GetMemoById(c *gin.Context) {
	id := c.Param("id")
	memo, err := api.Repo.GetById(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	now := time.Now()

	resp := memo.ToResponse(&now)
	resp.Distance = api.getDistance(&now, memo)
	c.JSON(http.StatusOK, resp)
}

func (api *MemoApi) GetMemosByUserId(c *gin.Context) {
	userId := api.getUserIdFromContext(c)
	memos, err := api.Repo.GetByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	respMemos := make([]*entity.MemoResponse, len(memos))
	now := time.Now()
	for i, val := range memos {
		resp := val.ToResponse(&now)
		resp.Distance = api.getDistance(&now, &val)
		respMemos[i] = resp
	}

	c.JSON(http.StatusOK, respMemos)
}

func (api *MemoApi) GetRecentMemos(c *gin.Context) {
	userId := api.getUserIdFromContext(c)
	start := c.Query("start")
	end := c.Query("end")
	//TODO - get by month+ day range
	//TODO - calculate distance
	memos, err := api.Repo.GetByDateRange(start, end, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	respMemos := make([]*entity.MemoResponse, len(memos))
	now := time.Now()
	for i, val := range memos {
		resp := val.ToResponse(&now)
		resp.Distance = api.getDistance(&now, &val)
		respMemos[i] = resp
	}

	c.JSON(http.StatusOK, respMemos)
}

func (api *MemoApi) UpdateMemoById(c *gin.Context) {
	id := c.Param("id")
	userId := api.getUserIdFromContext(c)
	var request entity.MemoRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var new_memo entity.Memo = entity.Memo{
		Id:          id,
		Subject:     request.Subject,
		Description: request.Description,
		UserId:      userId, //keep current userId for compare before update db
		MonthDay:    request.MonthDay,
		StartYear:   request.StartYear,
		Lunar:       request.Lunar,
	}

	err := api.Repo.Update(new_memo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Memo %v updated!", id))
}

func (api *MemoApi) RemoveMemo(c *gin.Context) {
	id := c.Param("id")
	err := api.Repo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("Memo %v deleted!", id))
	}
}

func (api *MemoApi) getUserIdFromContext(c *gin.Context) (userId string) {
	val, _ := c.Get("userId")
	userId = val.(string)
	return
}

func (api *MemoApi) getDistance(target *time.Time, memo *entity.Memo) []int {
	year := memo.StartYear
	if year <= 1900 {
		year = time.Now().Year()
	}

	startDate := year*10000 + memo.MonthDay
	targetDate := target.Year()*10000 + int(target.Month())*100 + target.Day()

	if memo.Lunar {
		before, after := api.DateService.GetLunarDistance(startDate, targetDate)
		return []int{before, after}
	} else {
		before, after := api.DateService.GetDistance(startDate, targetDate)
		return []int{before, after}
	}
}
