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
	"golang.org/x/oauth2/github"
)

const githubUserUrl = "https://api.github.com/user"

type GitHubUser struct {
	Email     string `json:"email"`
	Login     string `json:"login"`
	Id        int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

//var GithubAuthSet = wire.NewSet(provideGithubAuth, repository.RepoSet)

type GithubAuthApi struct {
	ConfGitHub *oauth2.Config
	Repo       repository.UserRepo

	jwtService *jwt.TokenService
}

func ProvideGithubAuth(repo repository.UserRepo, awsService *aws.AWSService, jwtService *jwt.TokenService) *GithubAuthApi {
	confGitHub := &oauth2.Config{
		ClientID:     awsService.GetParameterByKey("githubClientId"),
		ClientSecret: awsService.GetParameterByKey("githubClientSecret"),
		Scopes:       []string{"read:user", "user:email", "read:repo_hook"},
		Endpoint:     github.Endpoint,
	}

	return &GithubAuthApi{ConfGitHub: confGitHub, Repo: repo, jwtService: jwtService}
}

/* AuthorizeGithub
generate github authorize url and redirect directly
*/
func (api *GithubAuthApi) AuthorizeGithub(c *gin.Context) {
	//ctx := context.Background()
	//generate state and return to client can stop CSRF
	url := api.ConfGitHub.AuthCodeURL("state123", oauth2.AccessTypeOffline)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

/*
AuthorizeGithubUrl
generate github authorize url and return url
*/
func (api *GithubAuthApi) AuthorizeGithubUrl(c *gin.Context) {
	//ctx := context.Background()
	//generate state and return to client can stop CSRF
	url := api.ConfGitHub.AuthCodeURL("state123", oauth2.AccessTypeOffline)

	c.String(http.StatusOK, url)
}

/*
GetGithubToken
 redirect from github with code, call github api again in backend (for security reason), get access token generated by github
**/
func (api *GithubAuthApi) GetGithubToken(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if code == "" {
		c.JSON(http.StatusUnauthorized, "Token not found.")
	}

	if state != "state123" {
		c.JSON(http.StatusBadGateway, "Invalid state.")
	}

	//TODO: how to verify dynamic csrf token
	token, err := api.ConfGitHub.Exchange(c.Request.Context(), code)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, token)
}

/*
GetNativeToken - Get github user info and register into to our database, finally generate native jwt token
 1. get github user
 2. check native user by email
   2.1 update native user  -- email exists
   2.2 add native user     -- email not exists
 3 generate jwt token
 4 return token
*/
func (api *GithubAuthApi) GetNativeToken(c *gin.Context) {
	token := api.jwtService.GetToken(c)

	user, err := api.getGithubUser(githubUserUrl, token)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
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
		}
	}

	nativeToken, err := api.jwtService.NewToken(nativeUser.Id, nativeUser.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, nativeToken)
}

/*
CheckNativeToken - verify native token
*/
func (api *GithubAuthApi) CheckNativeToken(c *gin.Context) {
	token := api.jwtService.GetToken(c)

	claims, err := api.jwtService.ParseToken(token)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, claims)
}

func (api *GithubAuthApi) getGithubUser(url, token string) (*GitHubUser, error) {
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

	var user *GitHubUser = &GitHubUser{}
	json.Unmarshal(responseData, &user)

	return user, nil
}
