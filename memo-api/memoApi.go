package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FelixAnna/web-service-dlw/memo-api/di"
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
		micro.Registry(di.InitialRegistry().GetRegistry()),
	)

	service.Init()
	service.Run()
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	//define middleware before apis
	initialLogger()
	router.Use(gin.Logger())
	router.Use(di.InitialErrorMiddleware().ErrorHandler())
	router.Use(gin.Recovery())

	defineRoutes(router)

	//router.Run(":8282")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	var memoApi = di.InitialMemoApi()
	userGroupRouter := router.Group("/memos", di.InitialAuthorizationMiddleware().AuthorizationHandler())
	{
		userGroupRouter.PUT("/", memoApi.AddMemo)

		userGroupRouter.GET("/:id", memoApi.GetMemoById)
		userGroupRouter.GET("/", memoApi.GetMemosByUserId)
		userGroupRouter.GET("/recent", memoApi.GetRecentMemos)

		userGroupRouter.POST("/:id", memoApi.UpdateMemoById)
		userGroupRouter.DELETE("/:id", memoApi.RemoveMemo)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
