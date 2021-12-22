package entity

type DLWDate struct {
	YMD   int
	Lunar string

	LeapMonth bool
	Animal    string
	Today     bool
	WeekDay   int
}

type Distance struct {
	StartYMD  int
	TargetYMD int
	Lunar     bool
	Before    int64
	After     int64
}
