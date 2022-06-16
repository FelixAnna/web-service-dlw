package middleware

import (
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/common/jwt"

	"github.com/gin-gonic/gin"
)

type AuthorizationMiddleware struct {
	tokenService *jwt.TokenService
}

func ProvideAuthorizationMiddleware(service *jwt.TokenService) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{tokenService: service}
}

func (service *AuthorizationMiddleware) AuthorizationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		token := service.tokenService.GetToken(c)

		if token == "" {
			c.String(http.StatusForbidden, "token not found!")
			c.Abort()
			return
		}

		claims, err := service.tokenService.ParseToken(token)
		if err != nil {
			log.Println(err.Error())
			c.String(http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

		//set UserId in the request context
		c.Set("userId", claims.UserId)

		// before request
		log.Printf("User with email %v, Id %v send this request", claims.Email, claims.UserId)
		c.Next()

	}
}
