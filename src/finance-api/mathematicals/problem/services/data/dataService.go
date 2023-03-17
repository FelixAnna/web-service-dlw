package data

import "github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"

type DataService[number entity.Number] interface {
	GetData(criteria ...interface{}) number
}
