package stratergy

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/services/data"
	"github.com/stretchr/testify/assert"
)

var twoPlusStratergy *TwoPlusStratergy

func init() {
	twoPlusStratergy = NewTwoPlusStratergy(data.CreateRandomService())
}

func TestNewTwoPlusStratergy(t *testing.T) {
	assert.NotNil(t, twoPlusStratergy)
	assert.NotNil(t, twoPlusStratergy.DataService)
}

func TestTwoPlusGenerate(t *testing.T) {
	nums := twoPlusStratergy.Generate(100, 200)

	assert.Equal(t, len(nums), 3)
	assert.True(t, nums[0] >= 100 && nums[0] <= 200)
	assert.True(t, nums[1] >= 100 && nums[1] <= 200)
	assert.True(t, nums[0]+nums[1] == nums[2])
}
