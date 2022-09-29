package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/google/wire"
)

var TwoMultiplyStratergySet = wire.NewSet(NewTwoMultiplyStratergy, wire.Bind(new(Stratergy), new(*TwoMultiplyStratergy)))

type TwoMultiplyStratergy struct {
	*TwoNumStratergy
}

func NewTwoMultiplyStratergy(service data.DataService) *TwoMultiplyStratergy {
	return &TwoMultiplyStratergy{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/* TwoMultiplyStratergy.Generate
criteria[0]: bottom num
criteria[1]: ceiling num
*/
func (tp *TwoMultiplyStratergy) Generate(criteria ...interface{}) []int {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	nums = append(nums, nums[0]*nums[1])
	return nums
}
