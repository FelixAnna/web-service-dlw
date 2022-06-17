package mock

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

//mock of gin.Context
func GetGinContext(query string, header map[string][]string) *gin.Context {
	ctx := &gin.Context{}
	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: query,
		},
		Header: header,
	}

	return ctx
}
