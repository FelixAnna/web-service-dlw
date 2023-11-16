// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/filesystem"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/common/mocks"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/repositories"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/repository"
)

// Injectors from wire.go:

func InitializeZdjApi() (*zdj.ZdjApi, error) {
	awsHelper := aws.ProvideAwsHelper()
	awsService := aws.ProvideAWSService(awsHelper)
	zdjSqlServerRepo, err := repository.ProvideZdjSqlServerRepo(awsService)
	if err != nil {
		return nil, err
	}
	fileService := filesystem.ProvideFileService()
	zdjApi := zdj.ProvideZdjApi(zdjSqlServerRepo, fileService)
	return zdjApi, nil
}

func InitializeMathApi() *mathematicals.MathApi[int] {
	twoGenerationService := problem.NewTwoGenerationService[int]()
	awsHelper := aws.ProvideAwsHelper()
	awsService := aws.ProvideAWSService(awsHelper)
	mongoQuestionRepo := repositories.ProvideMongoQuestionRepo(awsService)
	mathService := problem.NewMathService(twoGenerationService, mongoQuestionRepo)
	mathApi := mathematicals.ProvideMathApi(mathService)
	return mathApi
}

func InitializeMockApi() (*zdj.ZdjApi, error) {
	zdjInMemoryRepo := repository.ProvideZdjInMemoryRepo()
	fileService := filesystem.ProvideFileService()
	zdjApi := zdj.ProvideZdjApi(zdjInMemoryRepo, fileService)
	return zdjApi, nil
}

func InitializeMockMathApi() *mathematicals.MathApi[int] {
	twoGenerationService := problem.NewTwoGenerationService[int]()
	mockAwsHelper := mocks.ProvideMockAwsHelper()
	awsService := aws.ProvideAWSService(mockAwsHelper)
	mongoQuestionRepo := repositories.ProvideMongoQuestionRepo(awsService)
	mathService := problem.NewMathService(twoGenerationService, mongoQuestionRepo)
	mathApi := mathematicals.ProvideMathApi(mathService)
	return mathApi
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
