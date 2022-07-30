package repository

import "github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"

type ZdjRepo interface {
	Append(zdj *[]entity.Zhidaojia) error

	Search(criteria *entity.Criteria) ([]entity.Zhidaojia, error)

	Delete(id int, version int) error
}
