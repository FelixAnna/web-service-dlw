package data

import (
	"math"
	"math/rand"
	"time"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	"github.com/google/wire"
)

var RandomServiceSet = wire.NewSet(CreateRandomService[int], wire.Bind(new(DataService[int]), new(*RandomService[int])))
var RandomServiceSetFloat = wire.NewSet(CreateRandomService[float32], wire.Bind(new(DataService[float32]), new(*RandomService[float32])))

var r1 *rand.Rand

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
}

type RandomService[number entity.Number] struct{}

func CreateRandomService[number entity.Number]() *RandomService[number] {
	return &RandomService[number]{}
}

func (rd *RandomService[number]) GetData(criteria ...interface{}) number {
	var bottom, up number

	_, ok := interface{}(&bottom).(float32)
	if ok {
		bottom = 0.0
		up = interface{}(math.MaxFloat32).(number)
	} else {
		bottom = 0.0
		up = interface{}(math.MaxInt32).(number)
	}

	if len(criteria) >= 2 {
		bottom, up = criteria[0].(number), criteria[1].(number)
		return rd.GetRand(up+1, bottom, ok) + bottom //+1 to include {up} itself
	}

	return rd.GetRand(up, bottom, ok) + bottom
}

func (rd *RandomService[number]) GetRand(up, bottom number, ok bool) number {
	if ok {
		var data = (interface{}(r1.Float32())).(number)
		for data > up-bottom {
			data -= up - bottom
		}
		return data
	} else {
		var start = (interface{}(up - bottom)).(int)
		return (interface{}(r1.Intn(start))).(number)
	}
}
