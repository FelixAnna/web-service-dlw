//go:build wireinject
// +build wireinject

package dependencyInjection

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/repository"
	"github.com/google/wire"
)

func InitializeApi() (zdj.ZdjApi, error) {
	wire.Build(zdj.ProvideZdjApi, repository.SqlRepoSet) //sql
	//wire.Build(zdj.ProvideZdjApi, repository.MemoryRepoSet) //InMemory
	return zdj.ZdjApi{}, nil
}
