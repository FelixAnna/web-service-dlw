package jwt

import "github.com/golang-jwt/jwt"

type MyCustomClaims struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type MyToken struct {
	Token string `json:"token"`
}
