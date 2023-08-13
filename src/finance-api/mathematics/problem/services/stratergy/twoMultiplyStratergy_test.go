package stratergy

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/services/data"
	"github.com/stretchr/testify/assert"
)

var twoMultiplyStratergy *TwoMultiplyStratergy

func init() {
	twoMultiplyStratergy = NewTwoMultiplyStratergy(data.CreateRandomService())
}

func TestNewTwoMultiplyStratergy(t *testing.T) {
	assert.NotNil(t, twoMultiplyStratergy)
	assert.NotNil(t, twoMultiplyStratergy.DataService)
}

func TestTwoMultiplyGenerate(t *testing.T) {
	nums := twoMultiplyStratergy.Generate(100, 200)

	assert.Equal(t, len(nums), 3)
	assert.True(t, nums[0] >= 100 && nums[0] <= 200)
	assert.True(t, nums[1] >= 100 && nums[1] <= 200)
	assert.True(t, nums[0]*nums[1] == nums[2])
}
