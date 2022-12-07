package middleware

import (
	"errors"
	"net/http"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/gin-gonic/gin"
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

func TestErrorHandlerFuncError(t *testing.T) {
	funx := errorService.ErrorHandler()
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{})
	ctx.Errors = append(ctx.Errors, &gin.Error{
		Err:  errors.New("any error"),
		Type: gin.ErrorTypeAny,
		Meta: "any error",
	})

	funx(ctx)

	assert.NotNil(t, ctx.Writer.Status(), http.StatusInternalServerError)
}

func TestErrorHandlerFunc(t *testing.T) {
	funx := errorService.ErrorHandler()
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{})

	funx(ctx)

	assert.NotNil(t, ctx.Writer.Status(), http.StatusOK)
}
