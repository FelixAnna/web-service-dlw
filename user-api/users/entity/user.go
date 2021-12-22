package entity

import "fmt"

type Date struct {
	Year  int `json:"year" binding:"required"`
	Month int `json:"month" binding:"required"`
	Day   int `json:"day" binding:"required"`
}

type Address struct {
	Country string `json:"Country" binding:"required"`
	State   string `json:"State" binding:"required"`
	City    string `json:"City" binding:"required"`
	Details string `json:"Details" binding:"required"`
}

type User struct {
	Id         string    `json:"Id" binding:""`
	Name       string    `json:"Name" binding:"required"`
	AvatarUrl  string    `json:"AvatarUrl" binding:""`
	Email      string    `json:"Email" binding:"required,email"`
	Phone      string    `json:"Phone" binding:"-"`
	Birthday   string    `json:"Birthday" binding:"required"`
	Address    []Address `json:"Address,omitempty" binding:"required,dive,required"`
	CreateTime string    `json:"CreateTime,omitempty"`
}

func (d *Date) String() string {
	return fmt.Sprintf("%04d%02d%02d", d.Year, d.Month, d.Day)
}
