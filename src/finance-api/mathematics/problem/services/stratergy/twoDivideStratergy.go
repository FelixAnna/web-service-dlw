package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/services/data"
	"github.com/google/wire"
)

var TwoDivideStratergySet = wire.NewSet(NewTwoDivideStratergy, wire.Bind(new(Stratergy), new(*TwoDivideStratergy)))

type TwoDivideStratergy struct {
	*TwoNumStratergy
}

func NewTwoDivideStratergy(service data.DataService) *TwoDivideStratergy {
	return &TwoDivideStratergy{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/*
	TwoDivideStratergy.Generate

criteria[0]: bottom num
criteria[1]: ceiling num
*/
func (tp *TwoDivideStratergy) Generate(criteria ...interface{}) []int {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	// avoid divide by 0
	for nums[0] == 0 {
		nums = tp.TwoNumStratergy.Generate(criteria...)
	}

	nums = append([]int{nums[0] * nums[1], nums[0]}, nums[1])
	return nums
}
