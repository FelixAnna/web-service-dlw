package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/micro"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/memo-api/di"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo"
	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "memo-api"

func main() {
	router := GetGinRouter()
	micro.StartApp(SERVER_NAME, ":8282", router, apiBoot.Registry.GetRegistry())
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
