package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FelixAnna/web-service-dlw/user-api/auth"
	"github.com/FelixAnna/web-service-dlw/user-api/di"
	"github.com/FelixAnna/web-service-dlw/user-api/users"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/server"
)

const SERVER_NAME = "user-api"

func main() {
	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8181"),
	)

	router := GetGinRouter()

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(apiBoot.Registry.GetRegistry()),
	)
	service.Init()
	service.Run()
}

type ApiBoot struct {
	UserApi              *users.UserApi
	AuthApi              *auth.GithubAuthApi
	ErrorHandler         *middleware.ErrorHandlingMiddleware
	AuthorizationHandler *middleware.AuthorizationMiddleware
	Registry             *mesh.Registry
}

var apiBoot *ApiBoot

func initialDependency() {
	apiBoot = &ApiBoot{}
	userApi := di.InitialUserApi()
	authApi := di.InitialGithubAuthApi()

	apiBoot.UserApi = &userApi
	apiBoot.AuthApi = &authApi
	apiBoot.AuthorizationHandler = di.InitialAuthorizationMiddleware()
	apiBoot.ErrorHandler = di.InitialErrorMiddleware()
	apiBoot.Registry = di.InitialRegistry()
}

func GetGinRouter() *gin.Engine {
	router := gin.New()
	initialDependency()

	//define middleware before apis
	initialLogger()
	router.Use(gin.Logger())
	router.Use(apiBoot.ErrorHandler.ErrorHandler())
	router.Use(gin.Recovery())

	defineRoutes(router)

	//router.Run(":8181")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	authGitHubRouter := router.Group("/oauth2/github")
	{
		authGitHubRouter.GET("/authorize", apiBoot.AuthApi.AuthorizeGithub)
		authGitHubRouter.GET("/authorize/url", apiBoot.AuthApi.AuthorizeGithubUrl)
		authGitHubRouter.GET("/redirect", apiBoot.AuthApi.GetGithubToken)
		authGitHubRouter.GET("/user", apiBoot.AuthApi.GetNativeToken)
		authGitHubRouter.GET("/checktoken", apiBoot.AuthApi.CheckNativeToken)
	}

	userGroupRouter := router.Group("/users", di.InitialAuthorizationMiddleware().AuthorizationHandler())
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

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
