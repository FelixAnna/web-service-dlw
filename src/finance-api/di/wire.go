//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/filesystem"
	"github.com/FelixAnna/web-service-dlw/common/jwt"
	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals"
	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem"

	"github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/repositories"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/repository"
	"github.com/google/wire"
)

func InitializeZdjApi() (*zdj.ZdjApi, error) {
	wire.Build(zdj.ProvideZdjApi, repository.SqlRepoSet, filesystem.FileSet, aws.ProvideAWSService, aws.AwsSet) //sql
	//wire.Build(zdj.ProvideZdjApi, repository.MemoryRepoSet) //InMemory
	return &zdj.ZdjApi{}, nil
}

func InitializeMathApi() *mathematicals.MathApi[int] {
	wire.Build(mathematicals.ProvideMathApi, problem.NewMathService, problem.ProvideTwoGenerationService, repositories.MongoRepoSet, aws.ProvideAWSService, aws.AwsSet) //sql
	//wire.Build(zdj.ProvideZdjApi, repository.MemoryRepoSet) //InMemory
	return &mathematicals.MathApi{}
}

func InitializeMathApi2() *mathematicals.MathApi[float32] {
	wire.Build(mathematicals.ProvideMathApi2, problem.NewMathService, problem.ProvideTwoGenerationService, repositories.MongoRepoSet, aws.ProvideAWSService, aws.AwsSet) //sql
	//wire.Build(zdj.ProvideZdjApi, repository.MemoryRepoSet) //InMemory
	return &mathematicals.MathApi{}
}

func InitializeMockApi() (*zdj.ZdjApi, error) {
	wire.Build(zdj.ProvideZdjApi, repository.MemoryRepoSet, filesystem.FileSet) //inmemory
	//wire.Build(zdj.ProvideZdjApi, repository.MemoryRepoSet) //InMemory
	return &zdj.ZdjApi{}, nil
}

func InitializeMockMathApi() *mathematicals.MathApi {
	wire.Build(mathematicals.ProvideMathApi, problem.NewMathService, problem.ProvideTwoGenerationService2, repositories.MongoRepoSet, aws.ProvideAWSService, aws.AwsMockSet) //sql
	//wire.Build(zdj.ProvideZdjApi, repository.MemoryRepoSet) //InMemory
	return &mathematicals.MathApi{}
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

func InitialMockAuthorizationMiddleware() *middleware.AuthorizationMiddleware {
	wire.Build(middleware.ProvideAuthorizationMiddleware, jwt.JwtMockSet)
	return &middleware.AuthorizationMiddleware{}
}
