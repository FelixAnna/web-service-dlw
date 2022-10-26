package main

import (
	"net/http"

	"github.com/FelixAnna/web-service-dlw/user-api/auth"
	"github.com/FelixAnna/web-service-dlw/user-api/di"
	"github.com/FelixAnna/web-service-dlw/user-api/users"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/micro"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "user-api"

func main() {
	initialDependency()
	router := GetGinRouter()

	//router.Run(":8181")
	micro.StartApp(SERVER_NAME, ":8181", router, apiBoot.Registry.GetRegistry())
}

type ApiBoot struct {
	UserApi              *users.UserApi
	AuthGithubApi        *auth.GithubAuthApi
	AuthGoogleApi        *auth.GoogleAuthApi
	ErrorHandler         *middleware.ErrorHandlingMiddleware
	AuthorizationHandler *middleware.AuthorizationMiddleware
	Registry             *mesh.Registry
}

var apiBoot *ApiBoot

func initialDependency() {
	apiBoot = &ApiBoot{
		UserApi:              di.InitialUserApi(),
		AuthGithubApi:        di.InitialGithubAuthApi(),
		AuthGoogleApi:        di.InitialGoogleAuthApi(),
		AuthorizationHandler: di.InitialAuthorizationMiddleware(),
		ErrorHandler:         di.InitialErrorMiddleware(),
		Registry:             di.InitialRegistry(),
	}
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	micro.RegisterMiddlewares(router, apiBoot.ErrorHandler.ErrorHandler())
	defineRoutes(router)

	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	authGitHubRouter := router.Group("/oauth2/github")
	{
		authGitHubRouter.GET("/authorize", apiBoot.AuthGithubApi.AuthorizeGithub)
		authGitHubRouter.GET("/authorize/url", apiBoot.AuthGithubApi.AuthorizeGithubUrl)
		authGitHubRouter.GET("/redirect", apiBoot.AuthGithubApi.GetGithubToken)
		authGitHubRouter.GET("/login", apiBoot.AuthGithubApi.Login)
		authGitHubRouter.GET("/user", apiBoot.AuthGithubApi.GetNativeToken)
		authGitHubRouter.GET("/checktoken", apiBoot.AuthGithubApi.CheckNativeToken)
	}

	authGoogleRouter := router.Group("/oauth2/google")
	{
		authGoogleRouter.GET("/login", apiBoot.AuthGoogleApi.Login)
		authGoogleRouter.GET("/checktoken", apiBoot.AuthGoogleApi.CheckNativeToken)
	}

	userGroupRouter := router.Group("/users", apiBoot.AuthorizationHandler.AuthorizationHandler())
	{
		userGroupRouter.GET("/", apiBoot.UserApi.GetAllUsers)
		userGroupRouter.GET("/:userId", apiBoot.UserApi.GetUserById)
		userGroupRouter.GET("/email/:email", apiBoot.UserApi.GetUserByEmail)

		userGroupRouter.POST("/:userId", apiBoot.UserApi.UpdateUserBirthdayById)
		userGroupRouter.POST("/:userId/address", apiBoot.UserApi.UpdateUserAddressById)

		userGroupRouter.PUT("/", apiBoot.UserApi.AddUser)

		userGroupRouter.DELETE("/:userId", apiBoot.UserApi.RemoveUser)
	}
}
