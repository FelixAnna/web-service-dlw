package stratergy

import "github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"

type Stratergy interface {
	Generate(criteria ...interface{}) []int
}

type TwoNumStratergy struct {
	DataService data.DataService
}

func NewTwoNumStratergy(service data.DataService) *TwoNumStratergy {
	return &TwoNumStratergy{
		DataService: service,
	}
}

func (ts *TwoNumStratergy) Generate(criteria ...interface{}) []int {
	a := ts.DataService.GetData(criteria)
	b := ts.DataService.GetData(criteria)

	return []int{a, b}
}
