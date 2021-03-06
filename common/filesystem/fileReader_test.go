package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var service *FileService

func init() {
	service = ProvideFileService()
}

func TestProvideFileService(t *testing.T) {
	assert.NotNil(t, service)
}

func TestReadLinesFailed(t *testing.T) {
	fileName := "./notexists.txt"

	results := service.ReadLines(fileName)

	assert.NotNil(t, results)
	assert.Equal(t, len(results), 0)
}

func TestReadLines(t *testing.T) {
	fileName := "./fileReader.go"

	results := service.ReadLines(fileName)

	assert.NotNil(t, results)
	assert.Greater(t, len(results), 0)
}
