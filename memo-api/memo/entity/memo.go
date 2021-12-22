package entity

import (
	"time"

	"github.com/FelixAnna/web-service-dlw/memo-api/memo/services"
)

type Memo struct {
	Id               string `json:"Id" binding:""`
	Subject          string `json:"Subject" binding:"required"`
	Description      string `json:"Description" binding:""`
	UserId           string `json:"UserId" binding:"required"`
	MonthDay         int    `json:"MonthDay" binding:"required"`
	StartYear        int    `json:"StartYear" binding:""`
	Lunar            bool   `json:"Lunar" binding:""` //user care about chinese Lunar only if checked
	CreateTime       string `json:"CreateTime,omitempty"`
	LastModifiedTime string `json:"LastModifiedTime,omitempty"`
}

type MemoRequest struct {
	Subject     string `json:"Subject" binding:"required"`
	Description string `json:"Description" binding:""`
	MonthDay    int    `json:"MonthDay" binding:"required"`
	StartYear   int    `json:"StartYear" binding:""`
	Lunar       bool   `json:"Lunar" binding:""`
}

type MemoResponse struct {
	Id               string `json:"Id" binding:""`
	UserId           string `json:"UserId" binding:"required"`
	Subject          string `json:"Subject" binding:"required"`
	Description      string `json:"Description" binding:""`
	MonthDay         int    `json:"MonthDay" binding:"required"`
	StartYear        int    `json:"StartYear" binding:""`
	Lunar            bool   `json:"Lunar" binding:""`
	Distance         []int  `json:"Distance" binding:"required"`
	CreateTime       string `json:"CreateTime,omitempty"`       //  - TODO convert to formated datetime
	LastModifiedTime string `json:"LastModifiedTime,omitempty"` //  - TODO convert to formated datetime
}

func (memo *Memo) ToResponse(now *time.Time) *MemoResponse {
	resp := &MemoResponse{
		Id:               memo.Id,
		UserId:           memo.UserId,
		Subject:          memo.Subject,
		Description:      memo.Description,
		MonthDay:         memo.MonthDay,
		StartYear:        memo.StartYear,
		Lunar:            memo.Lunar,
		CreateTime:       memo.CreateTime,
		LastModifiedTime: memo.LastModifiedTime,
		Distance:         memo.getDistance(now),
	}

	return resp
}

func (memo *Memo) getDistance(target *time.Time) []int {
	year := memo.StartYear
	if year <= 1900 {
		year = time.Now().Year()
	}

	startDate := year*10000 + memo.MonthDay
	targetDate := target.Year()*10000 + int(target.Month())*100 + target.Day()

	if memo.Lunar {
		before, after := services.GetLunarDistance(startDate, targetDate)
		return []int{before, after}
	} else {
		before, after := services.GetDistance(startDate, targetDate)
		return []int{before, after}
	}
}
