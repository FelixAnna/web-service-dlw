package repository

import (
	"database/sql"
	"fmt"
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
	if len(criteria.Districts) > 0 {
		query = query.Where("Distrct IN ?", criteria.Districts)
	}

	if len(criteria.Streets) > 0 {
		query = query.Where("Street IN ?", criteria.Streets)
	}

	if len(criteria.KeyWords) > 0 {
		keywords := fmt.Sprintf("%%%s%%", criteria.KeyWords)
		query = query.Where("(Community LIKE ? OR Street LIKE ?)", keywords, keywords)
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
		case "price_asc":
			query.Order("Price ASC")
		case "price_desc":
			query.Order("Price DESC")
		case "district_asc":
			query.Order("Distrct ASC")
		case "district_desc":
			query.Order("Distrct DESC")
		case "street_asc":
			query.Order("Street ASC")
		case "street_desc":
			query.Order("Street DESC")
		case "community_asc":
			query.Order("Community ASC")
		case "community_desc":
			query.Order("Community DESC")
		default: //price
			query.Order("Id ASC")
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
