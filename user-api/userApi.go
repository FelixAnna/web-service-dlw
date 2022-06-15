package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/user-api/di"

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
		micro.Registry(mesh.GetRegistry()),
	)
	service.Init()
	service.Run()
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	//define middleware before apis
	initialLogger()
	router.Use(gin.Logger())
	router.Use(middleware.ErrorHandler())
	router.Use(gin.Recovery())

	defineRoutes(router)

	//router.Run(":8181")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	var authApi = di.InitialGithubAuthApi()
	var userApi = di.InitialUserApi()

	authGitHubRouter := router.Group("/oauth2/github")
	{
		authGitHubRouter.GET("/authorize", authApi.AuthorizeGithub)
		authGitHubRouter.GET("/authorize/url", authApi.AuthorizeGithubUrl)
		authGitHubRouter.GET("/redirect", authApi.GetGithubToken)
		authGitHubRouter.GET("/user", authApi.GetNativeToken)
		authGitHubRouter.GET("/checktoken", authApi.CheckNativeToken)
	}

	userGroupRouter := router.Group("/users", middleware.AuthorizationHandler())
	{
		userGroupRouter.GET("/", userApi.GetAllUsers)
		userGroupRouter.GET("/:userId", userApi.GetUserById)
		userGroupRouter.GET("/email/:email", userApi.GetUserByEmail)

		userGroupRouter.POST("/:userId", userApi.UpdateUserBirthdayById)
		userGroupRouter.POST("/:userId/address", userApi.UpdateUserAddressById)

		userGroupRouter.PUT("/", userApi.AddUser)

		userGroupRouter.DELETE("/:userId", userApi.RemoveUser)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
