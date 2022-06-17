//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/repository"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/services"
	"github.com/google/wire"
)

func InitialMemoApi() memo.MemoApi {
	wire.Build(memo.MemoSet, repository.RepoSet, services.ProvideDateService, aws.ProvideAWSService, aws.AwsSet, mesh.ProvideRegistry)
	return memo.MemoApi{}
}

func InitialRegistry() *mesh.Registry {
	wire.Build(mesh.ProvideRegistry,
		aws.ProvideAWSService,
		aws.AwsSet)
	return &mesh.Registry{}
}

func InitialMockRegistry() *mesh.Registry {
	wire.Build(mesh.ProvideRegistry,
		aws.ProvideAWSService,
		aws.AwsMockSet)
	return &mesh.Registry{}
}

func InitialErrorMiddleware() *middleware.ErrorHandlingMiddleware {
	wire.Build(middleware.ProvideErrorHandlingMiddleware)
	return &middleware.ErrorHandlingMiddleware{}
}

func InitialAuthorizationMiddleware() *middleware.AuthorizationMiddleware {
	wire.Build(middleware.ProvideAuthorizationMiddleware, jwt.JwtSet)
	return &middleware.AuthorizationMiddleware{}
}
