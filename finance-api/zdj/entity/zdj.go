package entity

type Zhidaojia struct {
	Id int

	Distrct   string
	Street    string
	Community string
	Price     int

	Version int
}

type Criteria struct {
	Districts []string
	Streets   []string

	KeyWords string

	MinPrice int
	MaxPrice int

	Version int

	SortKey string

	Page int
	Size int
}
