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
	"github.com/FelixAnna/web-service-dlw/date-api/di"

	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/server"
)

const SERVER_NAME = "date-api"

func main() {
	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8383"),
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

	//define middleware before apis
	initialLogger()
	router.Use(gin.Logger())
	router.Use(middleware.ErrorHandler())
	router.Use(gin.Recovery())

	defineRoutes(router)

	//router.Run(":8383")
	return router
}

func defineRoutes(router *gin.Engine) {
	router.Any("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "running")
	})

	var dateApi = di.InitialDateApi()
	userGroupRouter := router.Group("/date")
	{
		userGroupRouter.GET("/current/month", dateApi.GetMonthDate)
		userGroupRouter.Any("/distance", dateApi.GetDateDistance)
		userGroupRouter.Any("/distance/lunar", dateApi.GetLunarDateDistance)
	}
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
