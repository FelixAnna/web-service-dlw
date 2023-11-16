package stratergy

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/stretchr/testify/assert"
)

var twoMinusStratergy *TwoMinusStratergy[int]

func init() {
	twoMinusStratergy = NewTwoMinusStratergy[int](data.CreateRandomService[int]())
}

func TestNewTwoMinusStratergy(t *testing.T) {
	assert.NotNil(t, twoMinusStratergy)
	assert.NotNil(t, twoMinusStratergy.DataService)
}

func TestTwoMinusGenerate(t *testing.T) {
	nums := twoMinusStratergy.Generate(100, 200)

	assert.Equal(t, len(nums), 3)
	assert.True(t, nums[0] >= 100 && nums[0] <= 200)
	assert.True(t, nums[1] >= 100 && nums[1] <= 200)
	assert.True(t, nums[0]-nums[1] == nums[2])
}
