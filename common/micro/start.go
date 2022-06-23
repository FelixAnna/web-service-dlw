package micro

import (
	"log"

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
