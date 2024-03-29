package main

import (
	"net/http"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/micro"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/memo-api/di"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo"
	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "memo-api"

func main() {
	initialDependency()
	router := GetGinRouter()

	//router.Run(":8282")
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
	apiBoot = &ApiBoot{
		MemoApi:              di.InitialMemoApi(),
		AuthorizationHandler: di.InitialAuthorizationMiddleware(),
		ErrorHandler:         di.InitialErrorMiddleware(),
		Registry:             di.InitialRegistry(),
	}
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	//define middleware before apis
	micro.RegisterMiddlewares(router, apiBoot.ErrorHandler.ErrorHandler())
	defineRoutes(router)

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
