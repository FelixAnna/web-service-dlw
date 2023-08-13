package stratergy

import (
	"fmt"
	"testing"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematics/problem/services/data"
	"github.com/stretchr/testify/assert"
)

var twoNumStratergy *TwoNumStratergy

func init() {
	twoNumStratergy = NewTwoNumStratergy(data.CreateRandomService())
}

func TestNewTwoNumStratergy(t *testing.T) {
	assert.NotNil(t, twoNumStratergy)
	assert.NotNil(t, twoNumStratergy.DataService)
}

func TestGenerate(t *testing.T) {
	nums := twoNumStratergy.Generate(100, 200)

	fmt.Println(nums)
	assert.Equal(t, len(nums), 2)
	assert.True(t, nums[0] >= 100 && nums[0] <= 200)
	assert.True(t, nums[1] >= 100 && nums[1] <= 200)
}
