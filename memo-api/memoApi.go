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
	"github.com/FelixAnna/web-service-dlw/memo-api/di"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/server"
)

const SERVER_NAME = "memo-api"

func main() {
	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8282"),
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
	MemoApi              *memo.MemoApi
	ErrorHandler         *middleware.ErrorHandlingMiddleware
	AuthorizationHandler *middleware.AuthorizationMiddleware
	Registry             *mesh.Registry
}

var apiBoot *ApiBoot

func initialDependency() {
	apiBoot = &ApiBoot{}
	memoApi := di.InitialMemoApi()

	apiBoot.MemoApi = &memoApi
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

	//router.Run(":8282")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	userGroupRouter := router.Group("/memos", apiBoot.AuthorizationHandler.AuthorizationHandler())
	{
		userGroupRouter.GET("/", apiBoot.MemoApi.GetMemosByUserId)
		userGroupRouter.GET("/:id", apiBoot.MemoApi.GetMemoById)
		userGroupRouter.GET("/recent", apiBoot.MemoApi.GetRecentMemos)

		userGroupRouter.PUT("/", apiBoot.MemoApi.AddMemo)

		userGroupRouter.POST("/:id", apiBoot.MemoApi.UpdateMemoById)
		userGroupRouter.DELETE("/:id", apiBoot.MemoApi.RemoveMemo)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
