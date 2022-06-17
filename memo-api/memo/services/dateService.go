package services

import (
	"context"
	"fmt"
	"log"

	"github.com/FelixAnna/web-service-dlw/common/mesh"
	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/selector"
)

const serviceName = "date-api"

/*
Date distance for the given month-day,
the previous and next startYMD around targetYMD
maps to before and after
*/
type Distance struct {
	StartYMD  int
	TargetYMD int
	Lunar     bool
	Before    int64
	After     int64
}

type DateService struct {
	DlwClient client.Client
}

//provide for wire
func ProvideDateService(registery *mesh.Registry) *DateService {
	reg := registery.GetRegistry()
	s := selector.NewSelector(selector.Registry(reg))
	dlwClient := httpClient.NewClient(client.Selector(s))

	return &DateService{DlwClient: dlwClient}
}

/*
Get distance from date-api for given date
Currently only support POST method
*/
func (service *DateService) GetDistance(start, end int) (before, after int) {
	return service.getDistance(start, end, false)
}

/*
Get distance from date-api for given date (Lunar)
Currently only support POST method
*/
func (service *DateService) GetLunarDistance(start, end int) (before, after int) {
	return service.getDistance(start, end, true)
}

/*
Get distance from date-api for given date (Lunar)
Currently only support POST method
*/
func (service *DateService) getDistance(start, end int, lunar bool) (before, after int) {
	category := "" //default calender
	if lunar {
		category = "/lunar"
	}

	path := fmt.Sprintf("/date/distance%v?start=%v&end=%v", category, start, end)

	// create request/response
	request := service.DlwClient.NewRequest(serviceName, path, "", client.WithContentType("application/json"))
	response := new(Distance)
	// call service
	err := service.DlwClient.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)

	return int(response.Before), int(response.After)
}

/*
func getDistance(url string) (*Distance, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var distance *Distance = &Distance{}
	json.Unmarshal(responseData, distance)

	return distance, nil
}
*/
