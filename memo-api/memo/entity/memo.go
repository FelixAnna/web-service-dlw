package entity

import (
	"time"
)

type Memo struct {
	Id               string `json:"Id" binding:"" bson:"_id"`
	Subject          string `json:"Subject" binding:"required" bson:"subject"`
	Description      string `json:"Description" binding:"" bson:"description"`
	UserId           string `json:"UserId" binding:"required" bson:"userid"`
	MonthDay         int    `json:"MonthDay" binding:"required" bson:"monthday"`
	StartYear        int    `json:"StartYear" binding:"" bson:"startyear"`
	Lunar            bool   `json:"Lunar" binding:"" bson:"lunar"` //user care about chinese Lunar only if checked
	CreateTime       string `json:"CreateTime,omitempty" bson:"createtime"`
	LastModifiedTime string `json:"LastModifiedTime,omitempty" bson:"lastmodifiedtime"`
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
	LunarYMD         int    `json:"LunarYMD" binding:""`
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
		//Distance:         memo.getDistance(now),
	}

	return resp
}
