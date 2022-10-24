package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	"github.com/FelixAnna/web-service-dlw/user-api/users/entity"
	"github.com/FelixAnna/web-service-dlw/user-api/users/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const googleUserUrl = "https://api.google.com/user"

type GoogleUser struct {
	Email     string `json:"email"`
	Login     string `json:"login"`
	Id        int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

//var GoogleAuthSet = wire.NewSet(provideGoogleAuth, repository.RepoSet)

type GoogleAuthApi struct {
	ConfGoogle *oauth2.Config
	Repo       repository.UserRepo

	jwtService *jwt.TokenService
}

func ProvideGoogleAuth(repo repository.UserRepo, awsService *aws.AWSService, jwtService *jwt.TokenService) *GoogleAuthApi {
	confGoogle := &oauth2.Config{
		ClientID:     awsService.GetParameterByKey("googleClientId"),
		ClientSecret: awsService.GetParameterByKey("googleClientSecret"),
		Scopes:       []string{"read:user", "user:email", "read:repo_hook"},
		Endpoint:     google.Endpoint,
	}

	return &GoogleAuthApi{ConfGoogle: confGoogle, Repo: repo, jwtService: jwtService}
}

//a combined API to get native token by auth code from google:
/*
	1. use google auth code  + state to get google token;
	2. use google token to get google user info;
	3. ensure google user email registered in our system;
	4. register new user if not already exists;
	5. generate token with our own signature.
*/
func (api *GoogleAuthApi) Login(c *gin.Context) {
	code := c.Query("code")
	//state := c.Query("state")

	if code == "" {
		c.JSON(http.StatusUnauthorized, "Token not found.")
		return
	}

	/* this part done by frontend
	if state != "state123" {
		c.JSON(http.StatusBadGateway, "Invalid state.")
		return
	}*/

	token, err := api.ConfGoogle.Exchange(c.Request.Context(), code)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err := api.getGoogleUser(googleUserUrl, token.AccessToken)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	nativeUser, err := api.Repo.GetByEmail(user.Email)
	if err != nil {
		api.Repo.Add(&entity.User{
			AvatarUrl: user.AvatarUrl,
			Email:     user.Email,
			Name:      user.Login,
			Birthday:  "2000-01-01",
			Address:   make([]entity.Address, 0),
		})

		nativeUser, err = api.Repo.GetByEmail(user.Email)

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	nativeToken, err := api.jwtService.NewToken(nativeUser.Id, nativeUser.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nativeToken)
}

/*
CheckNativeToken - verify native token
*/
func (api *GoogleAuthApi) CheckNativeToken(c *gin.Context) {
	token := api.jwtService.GetToken(c)

	claims, err := api.jwtService.ParseToken(token)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, claims)
}

func (api *GoogleAuthApi) getGoogleUser(url, token string) (*GoogleUser, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("token %v", token))
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var user *GoogleUser = &GoogleUser{}
	json.Unmarshal(responseData, &user)

	return user, nil
}
