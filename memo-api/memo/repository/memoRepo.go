package repository

import "github.com/FelixAnna/web-service-dlw/memo-api/memo/entity"

type MemoRepo interface {
	Add(memo *entity.Memo) (*string, error)

	GetById(id string) (*entity.Memo, error)
	GetByUserId(userId string) ([]entity.Memo, error)
	GetByDateRange(start, end, userId string) ([]entity.Memo, error)

	Update(memo entity.Memo) error

	Delete(id string) error
}
