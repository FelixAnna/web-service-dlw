package stratergy

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/google/wire"
)

var TwoDivideStratergySet = wire.NewSet(NewTwoDivideStratergy[int], wire.Bind(new(Stratergy[int]), new(*TwoDivideStratergy[int])))
var TwoDivideStratergySetFloat = wire.NewSet(NewTwoDivideStratergy[float32], wire.Bind(new(Stratergy[float32]), new(*TwoDivideStratergy[float32])))

type TwoDivideStratergy[number entity.Number] struct {
	*TwoNumStratergy[number]
}

func NewTwoDivideStratergy[number entity.Number](service data.DataService[number]) *TwoDivideStratergy[number] {
	return &TwoDivideStratergy[number]{
		TwoNumStratergy: NewTwoNumStratergy(service),
	}
}

/*
	TwoDivideStratergy.Generate

criteria[0]: bottom num
criteria[1]: ceiling num
*/
func (tp *TwoDivideStratergy[number]) Generate(criteria ...interface{}) []number {
	nums := tp.TwoNumStratergy.Generate(criteria...)

	// avoid divide by 0
	for nums[0] == 0 {
		nums = tp.TwoNumStratergy.Generate(criteria...)
	}

	nums = append([]number{nums[0] * nums[1], nums[0]}, nums[1])
	return nums
}
