package main

import (
	"net/http"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/micro"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/date-api/date"
	"github.com/FelixAnna/web-service-dlw/date-api/di"
	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "date-api"

func main() {
	initialDependency()
	router := GetGinRouter()
	micro.StartApp(SERVER_NAME, ":8383", router, apiBoot.Registry.GetRegistry())
}

type ApiBoot struct {
	DateApi      *date.DateApi
	ErrorHandler *middleware.ErrorHandlingMiddleware
	//AuthorizationHandler *middleware.AuthorizationMiddleware
	Registry *mesh.Registry
}

var apiBoot *ApiBoot

func initialDependency() {
	apiBoot = &ApiBoot{
		DateApi:      di.InitialDateApi(),
		ErrorHandler: di.InitialErrorMiddleware(),
		Registry:     di.InitialRegistry(),
	}
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	//define middleware before apis
	micro.RegisterMiddlewares(router, apiBoot.ErrorHandler.ErrorHandler())
	defineRoutes(router)

	//router.Run(":8383")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	userGroupRouter := router.Group("/date")
	{
		userGroupRouter.GET("/current/month", apiBoot.DateApi.GetMonthDate)
		userGroupRouter.POST("/distance", apiBoot.DateApi.GetDateDistance)
		userGroupRouter.POST("/distance/lunar", apiBoot.DateApi.GetLunarDateDistance)
	}
}
