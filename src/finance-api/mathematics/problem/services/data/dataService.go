package data

type DataService interface {
	GetData(criteria ...interface{}) int
}
