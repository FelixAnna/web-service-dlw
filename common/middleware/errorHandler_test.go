package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var errorService *ErrorHandlingMiddleware

func init() {
	errorService = ProvideErrorHandlingMiddleware()
}

func TestProvideErrorHandlingMiddleware(t *testing.T) {
	assert.NotNil(t, errorService)
}

func TestErrorHandler(t *testing.T) {
	funx := errorService.ErrorHandler()

	assert.NotNil(t, funx)
}
