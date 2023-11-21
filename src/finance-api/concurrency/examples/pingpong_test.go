package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	result := execute()
	assert.EqualValues(t, 10+1, result.hits)
	assert.EqualValues(t, 10+1, len(result.players))
}
