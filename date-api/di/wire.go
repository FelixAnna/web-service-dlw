//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/date-api/date"
	"github.com/FelixAnna/web-service-dlw/date-api/date/services"
	"github.com/google/wire"
)

func InitialDateApi() date.DateApi {
	wire.Build(date.ProvideDateApi, services.ProvideCarbonService)
	return date.DateApi{}
}
