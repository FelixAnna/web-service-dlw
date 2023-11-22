package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	seconds := 2
	result := execute(seconds)
	assert.EqualValues(t, result.hits, len(result.players))
	assert.GreaterOrEqual(t, seconds*10+1, len(result.players))
}
