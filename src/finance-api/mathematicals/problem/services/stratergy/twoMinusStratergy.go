package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/google/wire"
)

var TwoMinusStratergySet = wire.NewSet(NewTwoMinusStratergy[int], wire.Bind(new(Stratergy[int]), new(*TwoMinusStratergy[int])))

type TwoMinusStratergy[number entity.Number] struct {
	*TwoNumStratergy[number]
}

func NewTwoMinusStratergy[number entity.Number](service data.DataService[number]) *TwoMinusStratergy[number] {
	return &TwoMinusStratergy[number]{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/*
	TwoMinusStratergy.Generate

criteria[0]: bottom num
criteria[1]: ceiling num
*/
func (tp *TwoMinusStratergy[number]) Generate(criteria ...interface{}) []number {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	nums = append(nums, nums[0]-nums[1])
	return nums
}
