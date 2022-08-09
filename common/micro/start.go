package micro

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"go-micro.dev/v4"

	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

func StartApp(serverName, port string, router *gin.Engine, reg registry.Registry) {
	srv := httpServer.NewServer(
		server.Name(serverName),
		server.Address(port),
	)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(reg),
	)
	service.Init()
	service.Run()
}

func RegisterMiddlewares(router *gin.Engine, errorHandler gin.HandlerFunc) {
	//define middleware before apis
	corsSettings := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{"https://dlw-mi.azureedge.net", "http://localhost:3000"},
	}

	initialLogger()
	router.Use(gin.Logger())
	router.Use(errorHandler)
	router.Use(cors.New(corsSettings))
	router.Use(gin.Recovery())
}

func initialLogger() {
	year, month, day := time.Now().UTC().Date()
	date := fmt.Sprintf("%v%v%v", year, int(month), day)
	f, _ := os.Create("../logs/" + date + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
