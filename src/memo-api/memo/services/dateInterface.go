package services

type DateInterface interface {
	GetDistance(start, end int) (before, after, lunarYMD int)
	GetLunarDistance(start, end int) (before, after, lunarYMD int)
}
