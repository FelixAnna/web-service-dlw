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
	"github.com/FelixAnna/web-service-dlw/memo-api/memo"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/repository"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/services"
)

// Injectors from wire.go:

func InitialMemoApi() memo.MemoApi {
	awsHelper := aws.ProvideAwsHelper()
	awsService := aws.ProvideAWSService(awsHelper)
	memoRepoDynamoDB := repository.ProvideMemoRepoDynamoDB(awsService)
	registry := mesh.ProvideRegistry(awsService)
	dateService := services.ProvideDateService(registry)
	memoApi := memo.MemoApi{
		Repo:        memoRepoDynamoDB,
		DateService: dateService,
	}
	return memoApi
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
