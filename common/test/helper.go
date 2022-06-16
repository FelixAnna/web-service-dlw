package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type MockAwsHelper struct {
}

func ProvideMockAwsHelper() *MockAwsHelper {
	return &MockAwsHelper{}
}

func (service *MockAwsHelper) CreateSess() *session.Session {
	sess := func() *session.Session {
		// server is the mock server that simply writes a 200 status back to the client
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		return session.Must(session.NewSession(&aws.Config{
			DisableSSL: aws.Bool(true),
			Endpoint:   aws.String(server.URL),
		}))
	}()

	return sess
}

func (service *MockAwsHelper) LoadParameters(sess *session.Session) map[string]string {
	return map[string]string{
		"/dlf/dev/key1":            "value1",
		"/dlf/dev/key2":            "value2",
		"/dlf/dev/jwt/issuer":      "issuer",
		"/dlf/dev/jwt/signKey":     "signKey",
		"/dlf/dev/jwt/expiryAfter": "3600",
	}
}
