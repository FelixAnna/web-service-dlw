//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/user-api/auth"
	"github.com/FelixAnna/web-service-dlw/user-api/users"
	"github.com/FelixAnna/web-service-dlw/user-api/users/repository"
	"github.com/google/wire"
)

func InitialUserApi() users.UserApi {
	wire.Build(users.UserSet, repository.RepoSet, aws.ProvideAWSService, aws.AwsSet)
	return users.UserApi{}
}

func InitialGithubAuthApi() auth.GithubAuthApi {
	wire.Build(auth.ProvideGithubAuth, repository.RepoSet, jwt.JwtSet)
	return auth.GithubAuthApi{}
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
