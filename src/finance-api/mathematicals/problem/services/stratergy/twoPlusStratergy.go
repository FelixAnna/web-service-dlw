package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/google/wire"
)

var TwoPlusStratergySet = wire.NewSet(NewTwoPlusStratergy, wire.Bind(new(Stratergy), new(*TwoPlusStratergy)))

type TwoPlusStratergy struct {
	*TwoNumStratergy
}

func NewTwoPlusStratergy(service data.DataService) *TwoPlusStratergy {
	return &TwoPlusStratergy{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/* TwoPlusStratergy.Generate
criteria[0]: bottom num
criteria[1]: ceiling num
*/
func (tp *TwoPlusStratergy) Generate(criteria ...interface{}) []int {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	nums = append(nums, nums[0]+nums[1])
	return nums
}
