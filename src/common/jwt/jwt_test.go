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

func TestParseUserFromGoogleIDTokenInvalid(t *testing.T) {
	id_token := "invalid"
	tokenInfo, err := ParseUserFromGoogleIDToken(id_token)

	assert.NotNil(t, err)
	assert.Empty(t, tokenInfo)
}

func TestParseUserFromGoogleIDToken(t *testing.T) {
	id_token := "eyJraWQiOiIxZTlnZGs3IiwiYWxnIjoiUlMyNTYifQ.ewogImlzcyI6ICJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwKICJzdWIiOiAiMjQ4Mjg5NzYxMDAxIiwKICJhdWQiOiAiczZCaGRSa3F0MyIsCiAibm9uY2UiOiAibi0wUzZfV3pBMk1qIiwKICJleHAiOiAxMzExMjgxOTcwLAogImlhdCI6IDEzMTEyODA5NzAsCiAibmFtZSI6ICJKYW5lIERvZSIsCiAiZ2l2ZW5fbmFtZSI6ICJKYW5lIiwKICJmYW1pbHlfbmFtZSI6ICJEb2UiLAogImdlbmRlciI6ICJmZW1hbGUiLAogImJpcnRoZGF0ZSI6ICIwMDAwLTEwLTMxIiwKICJlbWFpbCI6ICJqYW5lZG9lQGV4YW1wbGUuY29tIiwKICJwaWN0dXJlIjogImh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyIKfQ.rHQjEmBqn9Jre0OLykYNnspA10Qql2rvx4FsD00jwlB0Sym4NzpgvPKsDjn_wMkHxcp6CilPcoKrWHcipR2iAjzLvDNAReF97zoJqq880ZD1bwY82JDauCXELVR9O6_B0w3K-E7yM2macAAgNCUwtik6SjoSUZRcf-O5lygIyLENx882p6MtmwaL1hd6qn5RZOQ0TLrOYu0532g9Exxcm-ChymrB4xLykpDj3lUivJt63eEGGN6DH5K6o33TcxkIjNrCD4XB1CKKumZvCedgHHF3IAK4dVEDSUoGlH9z4pP_eWYNXvqQOjGs-rDaQzUHl6cQQWNiDpWOl_lxXjQEvQ"
	tokenInfo, err := ParseUserFromGoogleIDToken(id_token)

	assert.Nil(t, err)
	assert.NotNil(t, tokenInfo)
}
