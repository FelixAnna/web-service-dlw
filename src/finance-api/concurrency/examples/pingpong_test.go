package examples

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	seconds := 2
	result := execute(seconds)
	assert.EqualValues(t, result.hits, len(result.players))
	assert.GreaterOrEqual(t, seconds*10+1, len(result.players))
}

func TestFactory(t *testing.T) {
	seconds := 2
	workers := []string{"Sam", "Alice", "Jack", "Helm"}
	result := mine(seconds, workers)

	fmt.Printf("%f", result)
	assert.GreaterOrEqual(t, result, float32(0))
}
