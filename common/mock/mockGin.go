package mock

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

//mock of gin.Context
func GetGinContext(query string, headers map[string][]string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)

	ctx.Request = &http.Request{
		URL: &url.URL{
			RawQuery: query,
		},
		Header: headers,
	}

	return ctx, writer
}
