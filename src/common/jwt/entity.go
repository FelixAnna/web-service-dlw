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

// TokenInfo struct
type GoogleTokenInfo struct {
	Iss string `json:"iss"`
	// userId
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	// clientId
	Aud string `json:"aud"`
	Iat int64  `json:"iat"`
	// expired time
	Exp int64 `json:"exp"`

	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Local         string `json:"locale"`
	jwt.StandardClaims
}
