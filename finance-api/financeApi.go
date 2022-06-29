package main

import (
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/micro"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/finance-api/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj"

	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "finance-api"

func main() {
	initialDependency()
	router := GetGinRouter()
	micro.StartApp(SERVER_NAME, ":8484", router, apiBoot.Registry.GetRegistry())
}

type ApiBoot struct {
	ZdjApi               *zdj.ZdjApi
	MathApi              *mathematicals.MathApi
	ErrorHandler         *middleware.ErrorHandlingMiddleware
	AuthorizationHandler *middleware.AuthorizationMiddleware
	Registry             *mesh.Registry
}

var apiBoot *ApiBoot

func initialDependency() {
	zdjApi, err := di.InitializeZdjApi()
	if err != nil {
		log.Panic(err)
		return
	}

	apiBoot = &ApiBoot{
		ZdjApi:               zdjApi,
		MathApi:              di.InitializeMathApi(),
		AuthorizationHandler: di.InitialAuthorizationMiddleware(),
		ErrorHandler:         di.InitialErrorMiddleware(),
		Registry:             di.InitialRegistry(),
	}
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//define middleware before apis
	micro.RegisterMiddlewares(router, apiBoot.ErrorHandler.ErrorHandler())
	defineRoutes(router)

	//router.Run(":8484")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	zdjGroupRouter := router.Group("/zdj", apiBoot.AuthorizationHandler.AuthorizationHandler())
	{
		zdjGroupRouter.GET("/", apiBoot.ZdjApi.GetAll)
		zdjGroupRouter.POST("/search", apiBoot.ZdjApi.Search)
		zdjGroupRouter.POST("/upload", apiBoot.ZdjApi.Upload)
		zdjGroupRouter.DELETE("/:id", apiBoot.ZdjApi.Delete)
		zdjGroupRouter.GET("/slow", apiBoot.ZdjApi.MemoryCosty)
	}

	mathGroupRouter := router.Group("/homework/math")
	{
		mathGroupRouter.POST("/", apiBoot.MathApi.GetQuestions)
		mathGroupRouter.POST("/multiple", apiBoot.MathApi.GetAllQuestions)
	}
}
