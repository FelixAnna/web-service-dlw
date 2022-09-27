package entity

import "fmt"

type Date struct {
	Year  int `json:"year" binding:"required"`
	Month int `json:"month" binding:"required"`
	Day   int `json:"day" binding:"required"`
}

type Address struct {
	Country string `json:"Country" binding:"required" bson:"country"`
	State   string `json:"State" binding:"required" bson:"state"`
	City    string `json:"City" binding:"required" bson:"city"`
	Details string `json:"Details" binding:"required" bson:"details"`
}

type User struct {
	Id         string    `json:"Id" binding:"" bson:"_id"`
	Name       string    `json:"Name" binding:"required" bson:"name"`
	AvatarUrl  string    `json:"AvatarUrl" binding:"" bson:"avatarurl"`
	Email      string    `json:"Email" binding:"required,email" bson:"email"`
	Phone      string    `json:"Phone" binding:"-" bson:"phone"`
	Birthday   string    `json:"Birthday" binding:"required" bson:"birthday"`
	Address    []Address `json:"Address,omitempty" binding:"required,dive,required" bson:"address"`
	CreateTime string    `json:"CreateTime,omitempty" bson:"createtime"`
}

func (d *Date) String() string {
	return fmt.Sprintf("%04d%02d%02d", d.Year, d.Month, d.Day)
}
