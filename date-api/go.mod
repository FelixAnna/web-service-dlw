module github.com/FelixAnna/web-service-dlw/date-api

go 1.17

require (
	github.com/FelixAnna/web-service-dlw/common v0.0.0-20211231152311-9d083a2c0544
	github.com/gin-gonic/gin v1.7.7
	github.com/golang-module/carbon/v2 v2.0.1
	go-micro.dev/v4 v4.4.0
)

require github.com/asim/go-micro/plugins/server/http/v4 v4.0.0-20211210113221-37de747d195c

// replace github.com/FelixAnna/web-service-dlw/common => ../common
