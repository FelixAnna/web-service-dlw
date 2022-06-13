//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/memo-api/memo"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/repository"
	"github.com/FelixAnna/web-service-dlw/memo-api/memo/services"
	"github.com/google/wire"
)

func InitialMemoApi() memo.MemoApi {
	wire.Build(memo.MemoSet, repository.RepoSet, services.ProvideDateService)
	return memo.MemoApi{}
}
