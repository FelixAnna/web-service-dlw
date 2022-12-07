package middleware

import (
	"net/http"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/stretchr/testify/assert"
)

var authService *AuthorizationMiddleware

func init() {
	helper := mocks.MockAwsHelper{}
	tokenService := jwt.ProvideTokenService(aws.ProvideAWSService(&helper))
	authService = ProvideAuthorizationMiddleware(tokenService)
}

func TestProvideAuthorizationMiddleware(t *testing.T) {
	assert.NotNil(t, authService)
	assert.NotEmpty(t, authService.TokenService)
}

func TestAuthorizationHandler(t *testing.T) {
	funx := authService.AuthorizationHandler()

	assert.NotNil(t, funx)
}

func TestAuthorizationHandlerFunc401(t *testing.T) {
	funx := authService.AuthorizationHandler()
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{Query: "access_code="})

	funx(ctx)

	assert.Equal(t, ctx.Writer.Status(), http.StatusUnauthorized)
}

func TestAuthorizationHandlerFunc403(t *testing.T) {
	funx := authService.AuthorizationHandler()
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{Query: "access_code=123"})

	funx(ctx)

	assert.Equal(t, ctx.Writer.Status(), http.StatusForbidden)
}

func TestAuthorizationHandlerFuncOK(t *testing.T) {
	funx := authService.AuthorizationHandler()
	token, _ := authService.TokenService.NewToken("testuser", "test@email.com")
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{Query: "access_code=" + token.Token})

	funx(ctx)

	assert.Equal(t, ctx.Writer.Status(), http.StatusOK)
}
