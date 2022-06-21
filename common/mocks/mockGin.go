package mocks

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Parameter struct {
	Query   string
	Headers map[string][]string
	Body    interface{}
	Params  map[string]string
}

func GetGinContext(param *Parameter) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)

	ctx.Request = &http.Request{}

	if param != nil {
		ctx.Request.URL = &url.URL{RawQuery: param.Query}

		if param.Headers != nil {
			ctx.Request.Header = param.Headers
		}

		if param.Body != nil {
			jsonValue, _ := json.Marshal(param.Body)
			ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(jsonValue))
		}

		if param.Params != nil {
			ctx.Params = gin.Params{}
			for k, v := range param.Params {
				ctx.Params = append(ctx.Params, gin.Param{Key: k, Value: v})
			}
		}
	}

	return ctx, writer
}
