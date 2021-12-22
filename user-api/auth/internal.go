package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	authServerURL = aws.GetParameterByKey("oauth2/serverDomain")

	config = oauth2.Config{
		ClientID:     aws.GetParameterByKey("oauth2/clientId"),
		ClientSecret: aws.GetParameterByKey("oauth2/clientSecret"),
		Scopes:       []string{"all"},
		RedirectURL:  fmt.Sprintf("%v/oauth2/token", aws.GetParameterByKey("oauth2/clientDomain")),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}

	globalToken *oauth2.Token // Non-concurrent security
)

var NativeServer *server.Server

func init() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:9096",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	NativeServer = srv
}

/*
func FireNativeAuthorize(c *gin.Context) {
	err := NativeServer.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	//c.JSON(http.StatusAccepted, NativeServer)
}

func GetNativeToken(c *gin.Context) {
	err := NativeServer.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
}*/

func GetRedirectUrl(c *gin.Context) {
	u := config.AuthCodeURL("xyz",
		oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256("s256example")),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"))
	c.String(http.StatusFound, u)
}

func GetToken(c *gin.Context) {
	state := c.Query("state")
	if state != "xyz" {
		c.String(http.StatusBadRequest, "State invalid")
		return
	}
	code := c.Query("code")
	if code == "" {
		c.String(http.StatusBadRequest, "Code not found")
		return
	}
	token, err := config.Exchange(context.Background(), code, oauth2.SetAuthURLParam("code_verifier", "s256example"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	globalToken = token

	c.JSON(http.StatusOK, token)
}

func RefreshToken(c *gin.Context) {
	if globalToken == nil {
		GetToken(c)
		return
	}

	globalToken.Expiry = time.Now()
	token, err := config.TokenSource(context.Background(), globalToken).Token()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	globalToken = token
	c.JSON(http.StatusOK, token)
}

func TestAccess(c *gin.Context) {
	if globalToken == nil {
		GetToken(c)
		return
	}

	resp, err := http.Get(fmt.Sprintf("%s/test?access_token=%s", authServerURL, globalToken.AccessToken))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer resp.Body.Close()

	io.Copy(c.Writer, resp.Body)
}

func PassordLogin(c *gin.Context) {
	token, err := config.PasswordCredentialsToken(context.Background(), "test", "test")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	globalToken = token
	c.JSON(http.StatusOK, token)
}

func ClientSecretLogin(c *gin.Context) {
	cfg := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.Endpoint.TokenURL,
	}

	token, err := cfg.Token(context.Background())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	globalToken = token
	c.JSON(http.StatusOK, token)
}

func genCodeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))
	return base64.URLEncoding.EncodeToString(s256[:])
}

/* native oauth2 server
authNativeRouter := router.Group("/oauth2")
{
	authNativeRouter.GET("/redirect", auth.GetRedirectUrl)
	authNativeRouter.GET("/token", auth.GetToken)
	authNativeRouter.GET("/refresh", auth.RefreshToken)
	authNativeRouter.GET("/test", auth.TestAccess)
	authNativeRouter.GET("/pwd", auth.PassordLogin)
	authNativeRouter.GET("/client", auth.ClientSecretLogin)
}*/
