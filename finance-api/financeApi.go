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
	zdj "github.com/FelixAnna/web-service-dlw/finance-api/zdj"
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

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(mesh.GetRegistry()),
	)

	service.Init()
	service.Run()
}

func GetGinRouter() *gin.Engine {
	router := gin.New()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//define middleware before apis
	initialLogger()
	router.Use(gin.Logger())
	router.Use(middleware.ErrorHandler())
	router.Use(gin.Recovery())

	defineRoutes(router)

	//router.Run(":8282")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	userGroupRouter := router.Group("/zdj", middleware.AuthorizationHandler())
	{
		userGroupRouter.GET("/", zdj.GetAll)
		userGroupRouter.POST("/search", zdj.Search)
		userGroupRouter.POST("/upload", zdj.Upload)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}