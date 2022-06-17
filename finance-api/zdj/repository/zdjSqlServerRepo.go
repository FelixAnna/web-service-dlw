package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/FelixAnna/web-service-dlw/common/aws"
	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	"github.com/google/wire"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var SqlRepoSet = wire.NewSet(ProvideZdjSqlServerRepo, wire.Bind(new(ZdjRepo), new(*ZdjSqlServerRepo)))

type ZdjSqlServerRepo struct {
	Db *gorm.DB
}

//provide for wire
func ProvideZdjSqlServerRepo(awsService *aws.AWSService) (*ZdjSqlServerRepo, error) {
	dsn := awsService.GetParameterByKey("sqldsn")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return &ZdjSqlServerRepo{}, err
	}

	return &ZdjSqlServerRepo{Db: db}, nil
}

func (s *ZdjSqlServerRepo) Append(zdj *[]entity.Zhidaojia) error {
	count := 0
	if s.Db.Model(&entity.Zhidaojia{}).Where("Version=?", (*zdj)[0].Version).FirstOrInit(&entity.Zhidaojia{}).RowsAffected == 0 {
		result := s.Db.CreateInBatches(*zdj, 100)
		if result.Error != nil {
			return result.Error
		}

		count = len(*zdj)
	}

	log.Printf("Inserted %v of %v items.", count, len(*zdj))
	return nil
}

func (s *ZdjSqlServerRepo) Search(criteria *entity.Criteria) ([]entity.Zhidaojia, error) {
	query := s.Db.Model(&entity.Zhidaojia{})
	if len(criteria.Distrct) > 0 {
		query = query.Where(&entity.Zhidaojia{Distrct: criteria.Distrct})
	}

	if len(criteria.Street) > 0 {
		query = query.Where(&entity.Zhidaojia{Street: criteria.Street})
	}

	if len(criteria.Community) > 0 {
		query = query.Where(&entity.Zhidaojia{Community: criteria.Community})
	}

	if criteria.MinPrice > 0 {
		query = query.Where("Price >= ?", criteria.MinPrice)
	}

	if criteria.MaxPrice > 0 {
		query = query.Where("Price <= ?", criteria.MaxPrice)
	}

	if criteria.Version > 0 {
		query = query.Where("Version = @version", sql.Named("version", criteria.Version))
	}

	if len(criteria.SortKey) > 0 {
		switch strings.ToLower(criteria.SortKey) {
		case "id":
			query.Order("Id DESC")
		case "district":
			query.Order("District DESC")
		case "street":
			query.Order("Street DESC")
		case "community":
			query.Order("Community DESC")
		default: //price
			query.Order("Price DESC")
		}
	}

	var items []entity.Zhidaojia
	result := query.Offset((criteria.Page - 1) * criteria.Size).Limit(criteria.Size).Find(&items)

	return items, result.Error
}

func (s *ZdjSqlServerRepo) Delete(id int, version int) error {
	result := s.Db.Model(&entity.Zhidaojia{}).Where(map[string]interface{}{"Id": id, "Version": version}).Delete(&entity.Zhidaojia{})
	if result.RowsAffected > 0 {
		log.Printf("deleted: id=%v, version=%v.", id, version)
	}
	return result.Error
}
