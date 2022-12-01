package data

import (
	"math"
	"math/rand"
	"time"

	"github.com/google/wire"
)

var RandomServiceSet = wire.NewSet(CreateRandomService, wire.Bind(new(DataService), new(*RandomService)))

var r1 *rand.Rand

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
}

type RandomService struct{}

func CreateRandomService() *RandomService {
	return &RandomService{}
}

func (rd *RandomService) GetData(criteria ...interface{}) int {
	bottom, up := 0, math.MaxInt32
	if len(criteria) >= 2 {
		bottom, up = criteria[0].(int), criteria[1].(int)
		return r1.Intn(up-bottom+1) + bottom //+1 to include {up} itself
	}
	return r1.Intn(up-bottom) + bottom
}
