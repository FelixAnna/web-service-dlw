package repository

import (
	"log"
	"strings"

	"github.com/FelixAnna/web-service-dlw/finance-api/zdj/entity"
	linq "github.com/ahmetb/go-linq/v3"
	"github.com/google/wire"
)

var zdjList []entity.Zhidaojia

func init() {
	zdjList = make([]entity.Zhidaojia, 0)
}

var MemoryRepoSet = wire.NewSet(ProvideZdjInMemoryRepo, wire.Bind(new(ZdjRepo), new(*ZdjInMemoryRepo)))

type ZdjInMemoryRepo struct {
}

//provide for wire
func ProvideZdjInMemoryRepo() *ZdjInMemoryRepo {
	return &ZdjInMemoryRepo{}
}

func (repo *ZdjInMemoryRepo) Append(zdj *[]entity.Zhidaojia) error {
	count := 0
	var current entity.Zhidaojia
	for i := 0; i < len(*zdj); i++ {
		current = (*zdj)[i]
		if !linq.From(zdjList).AnyWith(func(i interface{}) bool {
			return i.(entity.Zhidaojia).Id == current.Id && i.(entity.Zhidaojia).Version == current.Version
		}) {
			zdjList = append(zdjList, current)
			count++
		}
	}

	log.Printf("Inserted %v of %v items.", count, len(*zdj))
	return nil
}

func (repo *ZdjInMemoryRepo) Search(criteria *entity.Criteria) ([]entity.Zhidaojia, error) {

	var query linq.Query = linq.From(zdjList)
	if len(criteria.Distrct) > 0 {
		query = query.Where(func(i interface{}) bool {
			return i.(entity.Zhidaojia).Distrct == criteria.Distrct
		})
	}

	if len(criteria.Street) > 0 {
		query = query.Where(func(i interface{}) bool {
			return i.(entity.Zhidaojia).Street == criteria.Street
		})
	}

	if len(criteria.KeyWords) > 0 {
		query = query.Where(func(i interface{}) bool {
			return strings.Contains(i.(entity.Zhidaojia).Community, criteria.KeyWords)
		})
	}

	if criteria.MinPrice > 0 {
		query = query.Where(func(i interface{}) bool {
			return i.(entity.Zhidaojia).Price >= criteria.MinPrice
		})
	}

	if criteria.MaxPrice > 0 {
		query = query.Where(func(i interface{}) bool {
			return i.(entity.Zhidaojia).Price <= criteria.MaxPrice
		})
	}

	if criteria.Version > 0 {
		query = query.Where(func(i interface{}) bool {
			return i.(entity.Zhidaojia).Version == criteria.Version
		})
	}

	if len(criteria.SortKey) > 0 {
		switch strings.ToLower(criteria.SortKey) {
		case "id":
			query = query.Sort(func(i, j interface{}) bool { return i.(entity.Zhidaojia).Id < j.(entity.Zhidaojia).Id })
		case "district":
			query = query.Sort(func(i, j interface{}) bool { return i.(entity.Zhidaojia).Distrct < j.(entity.Zhidaojia).Distrct })
		case "street":
			query = query.Sort(func(i, j interface{}) bool { return i.(entity.Zhidaojia).Street < j.(entity.Zhidaojia).Street })
		case "community":
			query = query.Sort(func(i, j interface{}) bool { return i.(entity.Zhidaojia).Community < j.(entity.Zhidaojia).Community })
		default: //price
			query = query.Sort(func(i, j interface{}) bool { return i.(entity.Zhidaojia).Price < j.(entity.Zhidaojia).Price })
		}
	}

	var items []entity.Zhidaojia = []entity.Zhidaojia{}
	query.Skip((criteria.Page - 1) * criteria.Size).Take(criteria.Size).ToSlice(&items)

	return items, nil
}

func (repo *ZdjInMemoryRepo) Delete(id int, version int) error {
	for i := 0; i < len(zdjList); i++ {
		if id == zdjList[i].Id && version == zdjList[i].Version {
			zdjList = append(zdjList[:i], zdjList[i+1:]...)
			return nil
		}
	}

	return nil
}
