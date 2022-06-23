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
	"github.com/FelixAnna/web-service-dlw/date-api/date"
	"github.com/FelixAnna/web-service-dlw/date-api/di"
	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "date-api"

func main() {
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
	apiBoot = &ApiBoot{}
	dateApi := di.InitialDateApi()

	apiBoot.DateApi = dateApi
	//apiBoot.AuthorizationHandler = di.InitialAuthorizationMiddleware()
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
		userGroupRouter.GET("/distance", apiBoot.DateApi.GetDateDistance)
		userGroupRouter.GET("/distance/lunar", apiBoot.DateApi.GetLunarDateDistance)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
