package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/google/wire"
)

var TwoMinusStratergySet = wire.NewSet(NewTwoMinusStratergy, wire.Bind(new(Stratergy), new(*TwoMinusStratergy)))

type TwoMinusStratergy struct {
	*TwoNumStratergy
}

func NewTwoMinusStratergy(service data.DataService) *TwoMinusStratergy {
	return &TwoMinusStratergy{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/* TwoMinusStratergy.Generate
criteria[0]: bottom num
criteria[1]: ceiling num
criteria[2]: positive Only (1: pos only)
*/
func (tp *TwoMinusStratergy) Generate(criteria ...interface{}) []int {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	if len(criteria) > 2 && criteria[2] == 1 {
		if nums[0] < nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
	}

	nums = append(nums, nums[0]-nums[1])
	return nums
}
