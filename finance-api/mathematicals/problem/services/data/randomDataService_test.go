package data

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDataInRange(t *testing.T) {
	service := RandomService{}

	result := service.GetData(100, 200)

	assert.True(t, result >= 100 && result <= 200)
}

func TestGetData(t *testing.T) {
	service := RandomService{}

	result := service.GetData()

	assert.True(t, result >= 0 && result <= math.MaxInt32)
}
