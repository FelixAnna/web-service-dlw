package jwt

import (
	"log"
	"strconv"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type MyToken struct {
	Token string `json:"token"`
}

var (
	mySigningKey = []byte(aws.GetParameterByKey("jwt/signKey"))
	myIssuer     = aws.GetParameterByKey("jwt/issuer")
	myExpireAt   = aws.GetParameterByKey("jwt/expiryAfter")
)

func NewToken(id, email string) (*MyToken, error) {
	iExpiryAfter, err := strconv.ParseInt(myExpireAt, 10, 64)
	if err != nil {
		iExpiryAfter = 86400
	}

	// Create the Claims
	claims := MyCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + iExpiryAfter,
			Issuer:    myIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return &MyToken{Token: ss}, err
}

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		log.Println("Valid token and claims")
		return claims, nil
	} else {
		log.Println("invalid token and claims")
		return nil, err
	}
}

func GetToken(c *gin.Context) string {
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
