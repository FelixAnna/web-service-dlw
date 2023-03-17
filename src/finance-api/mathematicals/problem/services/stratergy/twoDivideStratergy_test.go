package stratergy

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/stretchr/testify/assert"
)

var twoDivideStratergy *TwoDivideStratergy[int]

func init() {
	twoDivideStratergy = NewTwoDivideStratergy[int](data.CreateRandomService())
}

func TestNewTwoDivideStratergy(t *testing.T) {
	assert.NotNil(t, twoDivideStratergy)
	assert.NotNil(t, twoDivideStratergy.DataService)
}

func TestTwoDivideGenerate(t *testing.T) {
	nums := twoDivideStratergy.Generate(100, 200)

	assert.Equal(t, len(nums), 3)
	assert.True(t, nums[1] >= 100 && nums[1] <= 200)
	assert.True(t, nums[2] >= 100 && nums[2] <= 200)
	assert.True(t, nums[0]/nums[1] == nums[2])
}
