// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/user-api/auth"
	"github.com/FelixAnna/web-service-dlw/user-api/users"
	"github.com/FelixAnna/web-service-dlw/user-api/users/repository"
)

// Injectors from wire.go:

func InitialUserApi() *users.UserApi {
	awsHelper := aws.ProvideAwsHelper()
	awsService := aws.ProvideAWSService(awsHelper)
	userRepoMongoDB := repository.ProvideUserRepoMongoDB(awsService)
	userApi := &users.UserApi{
		Repo: userRepoMongoDB,
	}
	return userApi
}

func InitialGithubAuthApi() *auth.GithubAuthApi {
	awsHelper := aws.ProvideAwsHelper()
	awsService := aws.ProvideAWSService(awsHelper)
	userRepoMongoDB := repository.ProvideUserRepoMongoDB(awsService)
	tokenService := jwt.ProvideTokenService(awsService)
	githubAuthApi := auth.ProvideGithubAuth(userRepoMongoDB, awsService, tokenService)
	return githubAuthApi
}

func InitialRegistry() *mesh.Registry {
	awsHelper := aws.ProvideAwsHelper()
	awsService := aws.ProvideAWSService(awsHelper)
	registry := mesh.ProvideRegistry(awsService)
	return registry
}

func InitialMockRegistry() *mesh.Registry {
	mockAwsHelper := mocks.ProvideMockAwsHelper()
	awsService := aws.ProvideAWSService(mockAwsHelper)
	registry := mesh.ProvideRegistry(awsService)
	return registry
}

func InitialErrorMiddleware() *middleware.ErrorHandlingMiddleware {
	errorHandlingMiddleware := middleware.ProvideErrorHandlingMiddleware()
	return errorHandlingMiddleware
}

func InitialAuthorizationMiddleware() *middleware.AuthorizationMiddleware {
	awsHelper := aws.ProvideAwsHelper()
	awsService := aws.ProvideAWSService(awsHelper)
	tokenService := jwt.ProvideTokenService(awsService)
	authorizationMiddleware := middleware.ProvideAuthorizationMiddleware(tokenService)
	return authorizationMiddleware
}

func InitialMockAuthorizationMiddleware() *middleware.AuthorizationMiddleware {
	mockAwsHelper := mocks.ProvideMockAwsHelper()
	awsService := aws.ProvideAWSService(mockAwsHelper)
	tokenService := jwt.ProvideTokenService(awsService)
	authorizationMiddleware := middleware.ProvideAuthorizationMiddleware(tokenService)
	return authorizationMiddleware
}
