package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
)

type Stratergy[number entity.Number] interface {
	Generate(criteria ...interface{}) []number
}

type TwoNumStratergy[number entity.Number] struct {
	DataService data.DataService[number]
}

func NewTwoNumStratergy[number entity.Number](service data.DataService[number]) *TwoNumStratergy[number] {
	return &TwoNumStratergy[number]{
		DataService: service,
	}
}

func (ts *TwoNumStratergy[number]) Generate(criteria ...interface{}) []number {
	a := ts.DataService.GetData(criteria...)
	b := ts.DataService.GetData(criteria...)

	return []number{a, b}
}
