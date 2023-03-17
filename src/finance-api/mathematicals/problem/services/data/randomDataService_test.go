package data

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var service DataService[int]

func init() {
	service = CreateRandomService()
}
func TestCreateRandomService(t *testing.T) {
	assert.NotNil(t, service)
}

func TestGetDataInRange(t *testing.T) {
	result := service.GetData(100, 200)

	assert.True(t, result >= 100 && result <= 200)
}

func TestGetData(t *testing.T) {
	result := service.GetData()

	assert.True(t, result >= 0 && result <= math.MaxInt32)
}
