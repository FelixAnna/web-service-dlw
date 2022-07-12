package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/micro"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/common/snowflake"
	"github.com/FelixAnna/web-service-dlw/finance-api/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj"

	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "finance-api"

func main() {
	os.Setenv("DLW_NODE_NO", "1023") //Debug Only

	initialDependency()
	router := GetGinRouter()

	router.Run(":8484") //Debug Only
	//micro.StartApp(SERVER_NAME, ":8484", router, apiBoot.Registry.GetRegistry())
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

	snowflake.InitSnowflake()
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//define middleware before apis
	micro.RegisterMiddlewares(router, apiBoot.ErrorHandler.ErrorHandler())
	defineRoutes(router)

	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	authorizationHandler := apiBoot.AuthorizationHandler.AuthorizationHandler()
	zdjGroupRouter := router.Group("/zdj", authorizationHandler)
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
		mathGroupRouter.POST("/save", apiBoot.MathApi.SaveResults, authorizationHandler)
		mathGroupRouter.POST("/multiple/feeds", apiBoot.MathApi.GetAllQuestionFeeds)
	}
}
