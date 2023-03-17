package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/google/wire"
)

var TwoMultiplyStratergySet = wire.NewSet(NewTwoMultiplyStratergy[int], wire.Bind(new(Stratergy[int]), new(*TwoMultiplyStratergy[int])))

type TwoMultiplyStratergy[number entity.Number] struct {
	*TwoNumStratergy[number]
}

func NewTwoMultiplyStratergy[number entity.Number](service data.DataService[number]) *TwoMultiplyStratergy[number] {
	return &TwoMultiplyStratergy[number]{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/*
	TwoMultiplyStratergy.Generate

criteria[0]: bottom num
criteria[1]: ceiling num
*/
func (tp *TwoMultiplyStratergy[number]) Generate(criteria ...interface{}) []number {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	nums = append(nums, nums[0]*nums[1])
	return nums
}
