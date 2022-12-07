//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/common/mesh"
	"github.com/FelixAnna/web-service-dlw/common/middleware"
	"github.com/FelixAnna/web-service-dlw/date-api/date"
	"github.com/FelixAnna/web-service-dlw/date-api/date/services"
	"github.com/google/wire"
)

func InitialDateApi() *date.DateApi {
	wire.Build(date.ProvideDateApi, services.CarbonSet)
	return &date.DateApi{}
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
