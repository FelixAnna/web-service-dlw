package jwt

import (
	"log"
	"strconv"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
)

var JwtSet = wire.NewSet(ProvideTokenService, aws.ProvideAWSService, aws.AwsSet)
var JwtMockSet = wire.NewSet(ProvideTokenService, aws.ProvideAWSService, aws.AwsMockSet)

type TokenService struct {
	mySigningKey []byte
	myIssuer     string
	myExpireAt   string
}

func ProvideTokenService(awsService *aws.AWSService) *TokenService {
	mySigningKey := []byte(awsService.GetParameterByKey("jwt/signKey"))
	myIssuer := awsService.GetParameterByKey("jwt/issuer")
	myExpireAt := awsService.GetParameterByKey("jwt/expiryAfter")

	return &TokenService{mySigningKey: mySigningKey, myIssuer: myIssuer, myExpireAt: myExpireAt}
}

func (service *TokenService) NewToken(id, email string) (*MyToken, error) {
	iExpiryAfter, err := strconv.ParseInt(service.myExpireAt, 10, 64)
	if err != nil {
		iExpiryAfter = 86400
	}

	// Create the Claims
	claims := MyCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + iExpiryAfter,
			Issuer:    service.myIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(service.mySigningKey)

	return &MyToken{Token: ss}, err
}

func (service *TokenService) ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return service.mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		log.Println("Valid token and claims")
		return claims, nil
	} else {
		log.Println("invalid token and claims")
		return nil, err
	}
}

func (service *TokenService) GetToken(c *gin.Context) string {
	token := c.Query("access_code")
	if token == "" {
		token = c.GetHeader("Authorization")
		if len(token) <= 7 {
			return ""
		}

		token = token[7:]
	}
	return token
}

func ParseUserFromGoogleIDToken(id_token string) (*GoogleTokenInfo, error) {
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(id_token, &GoogleTokenInfo{})
	if err == nil {
		if tokenInfo, ok := token.Claims.(*GoogleTokenInfo); ok {
			return tokenInfo, nil
		}
	}

	// parse token.payload failed
	return nil, err
}
