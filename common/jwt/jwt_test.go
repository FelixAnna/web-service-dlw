package jwt

import (
	"fmt"
	"testing"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/stretchr/testify/assert"
)

var service *TokenService

func init() {
	helper := mocks.MockAwsHelper{}
	service = ProvideTokenService(aws.ProvideAWSService(&helper))
}

func TestProvideTokenService(t *testing.T) {
	assert.NotNil(t, service)
	assert.NotEmpty(t, service.myExpireAt)
	assert.NotEmpty(t, service.myIssuer)
	assert.NotNil(t, service.mySigningKey)
}

func TestNewTokenInvalidExpire(t *testing.T) {
	service.myExpireAt = "not a number"
	id, email := "123", "test@example.com"

	result, err := service.NewToken(id, email)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.NotNil(t, result.Token)
}

func TestNewTokenNormal(t *testing.T) {
	service.myExpireAt = "100"
	id, email := "123", "test@example.com"

	result, err := service.NewToken(id, email)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.NotNil(t, result.Token)
}

func TestParseTokenInvalid(t *testing.T) {
	token := "invalid token"

	claims, err := service.ParseToken(token)

	assert.NotNil(t, err)
	assert.Nil(t, claims)
}

func TestParseTokenExpired(t *testing.T) {
	service.myExpireAt = "-1800"
	id, email := "123", "test@example.com"
	token, _ := service.NewToken(id, email)

	claims, err := service.ParseToken(token.Token)

	fmt.Println(claims, err)
	assert.NotNil(t, err)
	assert.Nil(t, claims)
}

func TestParseToken(t *testing.T) {
	service.myExpireAt = "1800"
	id, email := "123", "test@example.com"
	token, _ := service.NewToken(id, email)

	claims, err := service.ParseToken(token.Token)

	assert.Nil(t, err)
	assert.NotNil(t, claims)
}

func TestGetTokenByHeader(t *testing.T) {
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{Headers: map[string][]string{"Authorization": {"Bearer abc"}}})

	token := service.GetToken(ctx)

	assert.NotEmpty(t, token)
}

func TestGetTokenByCode(t *testing.T) {
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{Query: "access_code=abc"})

	token := service.GetToken(ctx)

	assert.NotEmpty(t, token)
}

func TestGetTokenEmptyCodeAndHeader(t *testing.T) {
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{Headers: map[string][]string{}})

	token := service.GetToken(ctx)

	assert.Empty(t, token)
}

func TestGetTokenInvalid(t *testing.T) {
	ctx, _ := mocks.GetGinContext(&mocks.Parameter{Query: "access_code=", Headers: map[string][]string{"Authorization": {"invalid"}}})

	token := service.GetToken(ctx)

	assert.Empty(t, token)
}
