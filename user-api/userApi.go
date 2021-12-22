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
	"github.com/FelixAnna/web-service-dlw/user-api/auth"
	userService "github.com/FelixAnna/web-service-dlw/user-api/users"

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
		micro.Registry(mesh.GetConsulRegistry()),
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

	authGitHubRouter := router.Group("/oauth2/github")
	{
		authGitHubRouter.GET("/authorize", auth.AuthorizeGithub)
		authGitHubRouter.GET("/authorize/url", auth.AuthorizeGithubUrl)
		authGitHubRouter.GET("/redirect", auth.GetGithubToken)
		authGitHubRouter.GET("/user", auth.GetNativeToken)
		authGitHubRouter.GET("/checktoken", auth.CheckNativeToken)
	}

	userGroupRouter := router.Group("/users", middleware.AuthorizationHandler())
	{
		userGroupRouter.GET("/", userService.GetAllUsers)
		userGroupRouter.GET("/:userId", userService.GetUserById)
		userGroupRouter.GET("/email/:email", userService.GetUserByEmail)

		userGroupRouter.POST("/:userId", userService.UpdateUserBirthdayById)
		userGroupRouter.POST("/:userId/address", userService.UpdateUserAddressById)

		userGroupRouter.PUT("/", userService.AddUser)

		userGroupRouter.DELETE("/:userId", userService.RemoveUser)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

/*

update name:
curl -H "Content-Type:application/json" -X PUT -d '{"name":"felix","email":"felix@example.com","phone":"+8612345678901","birthday": "1989-07-11","address":[{"country":"China","state":"Guangdong","city":"Shenzhen","details":"futian"}]}' http://localhost:8181/users/

curl -H "Content-Type:application/json" -X POST  http://localhost:8181/users/1637418999081?birthday=1989-07-12

curl -H "Content-Type:application/json" -X POST -d '[{"country":"China","state":"Guangdong","city":"Shenzhen","details":"futian2"}]' http://localhost:8181/users/1637418999081/address


curl -H "Content-Type:application/json" -X DELETE  http://localhost:8181/users/1637418999081

http://localhost:8181/users/email/felix@example.com
*/
