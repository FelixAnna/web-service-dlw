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
	"github.com/FelixAnna/web-service-dlw/finance-api/di"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"

	"go-micro.dev/v4"
	"go-micro.dev/v4/server"

	"github.com/gin-gonic/gin"
)

const SERVER_NAME = "finance-api"

func main() {
	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8484"),
	)

	router := GetGinRouter()

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	registery := apiBoot.Registry
	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registery.GetRegistry()),
	)

	service.Init()
	service.Run()
}

type ApiBoot struct {
	ZdjApi               *zdj.ZdjApi
	ErrorHandler         *middleware.ErrorHandlingMiddleware
	AuthorizationHandler *middleware.AuthorizationMiddleware
	Registry             *mesh.Registry
}

var apiBoot *ApiBoot

func initialDependency() {
	apiBoot = &ApiBoot{}
	zdjApi, err := di.InitializeApi()
	if err != nil {
		log.Panic(err)
		return
	}

	apiBoot.ZdjApi = &zdjApi
	apiBoot.AuthorizationHandler = di.InitialAuthorizationMiddleware()
	apiBoot.ErrorHandler = di.InitialErrorMiddleware()
	apiBoot.Registry = di.InitialRegistry()
}

func GetGinRouter() *gin.Engine {
	router := gin.New()
	initialDependency()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//define middleware before apis
	initialLogger()
	router.Use(gin.Logger())
	router.Use(apiBoot.ErrorHandler.ErrorHandler())
	router.Use(gin.Recovery())

	defineRoutes(router)

	//router.Run(":8484")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	userGroupRouter := router.Group("/zdj", apiBoot.AuthorizationHandler.AuthorizationHandler())
	{
		userGroupRouter.GET("/", apiBoot.ZdjApi.GetAll)
		userGroupRouter.POST("/search", apiBoot.ZdjApi.Search)
		userGroupRouter.POST("/upload", apiBoot.ZdjApi.Upload)
		userGroupRouter.DELETE("/:id", apiBoot.ZdjApi.Delete)
		userGroupRouter.GET("/slow", apiBoot.ZdjApi.MemoryCosty)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
