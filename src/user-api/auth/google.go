package auth

import (
	"fmt"
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
		RedirectURL:  awsService.GetParameterByKey("googleRedirectURL"),
		Scopes:       []string{"profile", "email", "openid"},
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

	//google support openid connect, will return id_token which have user profile
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		fmt.Println("No id_token")
	}

	user, err := jwt.ParseUserFromGoogleIDToken(idToken)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	nativeUser, err := api.Repo.GetByEmail(user.Email)
	if err != nil {
		api.Repo.Add(&entity.User{
			AvatarUrl: user.Picture,
			Email:     user.Email,
			Name:      user.Name,
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
