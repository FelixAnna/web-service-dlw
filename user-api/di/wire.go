//go:build wireinject
// +build wireinject

package di

import (
	"github.com/FelixAnna/web-service-dlw/user-api/auth"
	"github.com/FelixAnna/web-service-dlw/user-api/users"
	"github.com/FelixAnna/web-service-dlw/user-api/users/repository"
	"github.com/google/wire"
)

func InitialUserApi() users.UserApi {
	wire.Build(users.UserSet, repository.RepoSet)
	return users.UserApi{}
}

func InitialGithubAuthApi() auth.GithubAuthApi {
	wire.Build(auth.ProvideGithubAuth, repository.RepoSet)
	return auth.GithubAuthApi{}
}
