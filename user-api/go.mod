module github.com/FelixAnna/web-service-dlw/user-api

go 1.14

require (
	github.com/FelixAnna/web-service-dlw/common v0.0.0-20211228152933-29d52f28e6b8
	//github.com/FelixAnna/web-service-dlw/common v0.0.0-00010101000000-000000000000
	github.com/asim/go-micro/plugins/server/http/v4 v4.0.0-20211210113221-37de747d195c
	github.com/aws/aws-sdk-go v1.42.21
	github.com/gin-gonic/gin v1.7.7
	github.com/go-oauth2/oauth2/v4 v4.4.2
	go-micro.dev/v4 v4.4.0
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
)

//replace github.com/FelixAnna/web-service-dlw/common => ../common
