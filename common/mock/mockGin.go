package mock

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

//mock of gin.Context
func GetGinContext(query string, headers map[string][]string) *gin.Context {
	ctx := &gin.Context{}
	gin.CreateTestContext(httptest.NewRecorder())

	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: query,
		},
		Header: headers,
	}

	return ctx
}
