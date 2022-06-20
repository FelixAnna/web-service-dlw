package mock

import (
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	"github.com/stretchr/testify/mock"
)

type ZdjMockRepo struct {
	mock.Mock
}

func (repo *ZdjMockRepo) Append(zdj *[]entity.Zhidaojia) error {
	args := repo.Called(zdj)
	return args.Get(0).(error)
}

func (repo *ZdjMockRepo) Search(criteria *entity.Criteria) ([]entity.Zhidaojia, error) {
	args := repo.Called(criteria)
	return args.Get(0).([]entity.Zhidaojia), args.Error(1)
}

func (repo *ZdjMockRepo) Delete(id int, version int) error {
	args := repo.Called(id, version)
	return args.Error(0)
}
