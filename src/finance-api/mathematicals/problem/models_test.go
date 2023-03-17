package problem

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRange(t *testing.T) {
	criteria := &Criteria[int]{
		Min: 100,
		Max: 200,

		Quantity: 5,

		Category: CategoryPlus,
	}

	min, max := criteria.GetRange()

	assert.Equal(t, min, math.MinInt32)
	assert.Equal(t, max, math.MaxInt32)
}

func TestGetRangeReverse(t *testing.T) {
	criteria := &Criteria[int]{
		Range: &Range{
			Min: 200,
			Max: 100,
		},
	}

	min, max := criteria.GetRange()

	assert.Equal(t, min, criteria.Range.Max)
	assert.Equal(t, max, criteria.Range.Min)
}

func TestGetRangeSorted(t *testing.T) {
	criteria := &Criteria[int]{
		Range: &Range{
			Min: 100,
			Max: 200,
		},
	}

	min, max := criteria.GetRange()

	assert.Equal(t, min, criteria.Range.Min)
	assert.Equal(t, max, criteria.Range.Max)
}
