package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/google/wire"
)

var TwoPlusStratergySet = wire.NewSet(NewTwoPlusStratergy[int], wire.Bind(new(Stratergy[int]), new(*TwoPlusStratergy[int])))

type TwoPlusStratergy[number entity.Number] struct {
	*TwoNumStratergy[number]
}

func NewTwoPlusStratergy[number entity.Number](service data.DataService[number]) *TwoPlusStratergy[number] {
	return &TwoPlusStratergy[number]{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/*
	TwoPlusStratergy.Generate

criteria[0]: bottom num
criteria[1]: ceiling num
*/
func (tp *TwoPlusStratergy[number]) Generate(criteria ...interface{}) []number {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	nums = append(nums, nums[0]+nums[1])
	return nums
}
