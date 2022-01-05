module github.com/FelixAnna/web-service-dlw/memo-api

go 1.17

require (
	//github.com/FelixAnna/web-service-dlw/common v0.0.0-00010101000000-000000000000
	github.com/FelixAnna/web-service-dlw/common v0.0.0-20211231152311-9d083a2c0544
	github.com/asim/go-micro/plugins/client/http/v4 v4.0.0-20211210113221-37de747d195c
	github.com/asim/go-micro/plugins/server/http/v4 v4.0.0-20211210113221-37de747d195c
	github.com/aws/aws-sdk-go v1.42.21
	github.com/gin-gonic/gin v1.7.7
	go-micro.dev/v4 v4.4.0
)

//replace github.com/FelixAnna/web-service-dlw/common => ../common
