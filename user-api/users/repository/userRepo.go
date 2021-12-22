package repository

import "github.com/FelixAnna/web-service-dlw/user-api/users/entity"

type UserRepo interface {
	GetAllTables()
	GetAll() ([]entity.User, error)

	Add(user *entity.User) (*string, error)
	GetById(userId string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)

	UpdateBirthday(userId, birthday string) error
	UpdateAddress(userId string, addresses []entity.Address) error

	Delete(userId string) error
}
