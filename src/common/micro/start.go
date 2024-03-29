package micro

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
	httpServer "github.com/go-micro/plugins/v4/server/http"
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
	profile := os.Getenv("profile")
	corsSettings := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    true,
		AllowOrigins:     []string{"https://*.metadlw.com", "http://localhost:3000"},
	}

	//allow all origin for local debug/deployment
	if profile == "local" {
		corsSettings = cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
			AllowAllOrigins:  true,
		}
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
