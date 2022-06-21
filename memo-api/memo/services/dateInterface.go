package services

type DateInterface interface {
	GetDistance(start, end int) (before, after int)
	GetLunarDistance(start, end int) (before, after int)
}
