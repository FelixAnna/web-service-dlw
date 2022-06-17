package middleware

import (
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	mock "github.com/FelixAnna/web-service-dlw/common/mock"
	"github.com/stretchr/testify/assert"
)

var authService *AuthorizationMiddleware

func init() {
	helper := mock.MockAwsHelper{}
	tokenService := jwt.ProvideTokenService(aws.ProvideAWSService(&helper))
	authService = ProvideAuthorizationMiddleware(tokenService)
}

func TestProvideAuthorizationMiddleware(t *testing.T) {
	assert.NotNil(t, authService)
	assert.NotEmpty(t, authService.tokenService)
}

func TestAuthorizationHandler(t *testing.T) {
	funx := authService.AuthorizationHandler()

	assert.NotNil(t, funx)
}
