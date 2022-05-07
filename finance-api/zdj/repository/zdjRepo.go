package repository

import "github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"

type ZdjRepo interface {
	//Add(zdj *entity.Zhidaojia) (int, error)

	Append(zdj *[]entity.Zhidaojia) (bool, error)

	Search(criteria *entity.Criteria) ([]entity.Zhidaojia, error)

	Delete(id int, version int) error
}
