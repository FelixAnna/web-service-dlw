package services

import (
	"context"
	"fmt"
	"log"

	httpClient "github.com/asim/go-micro/plugins/client/http/v4"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
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

var dateClient client.Client

func init() {
	consulReg := consul.NewRegistry(registry.Addrs("localhost:8500"))
	//r := registry.NewRegistry()
	s := selector.NewSelector(selector.Registry(consulReg))
	// new client
	dateClient = httpClient.NewClient(client.Selector(s))
}

/*
Get distance from date-api for given date
Currently only support POST method
*/
func GetDistance(start, end int) (before, after int) {
	return getDistance(start, end, false)
}

/*
Get distance from date-api for given date (Lunar)
Currently only support POST method
*/
func GetLunarDistance(start, end int) (before, after int) {
	return getDistance(start, end, true)
}

/*
Get distance from date-api for given date (Lunar)
Currently only support POST method
*/
func getDistance(start, end int, lunar bool) (before, after int) {
	lunarPath := ""
	if lunar {
		lunarPath = "/lunar"
	}

	path := fmt.Sprintf("/date/distance%v?start=%v&end=%v", lunarPath, start, end)

	// create request/response
	request := dateClient.NewRequest(serviceName, path, "", client.WithContentType("application/json"))
	response := new(Distance)
	// call service
	err := dateClient.Call(context.Background(), request, response)
	log.Printf("err:%v response:%#v\n", err, response)

	return int(response.Before), int(response.After)
}

/*
func getDistance(url string) (*Distance, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var distance *Distance = &Distance{}
	json.Unmarshal(responseData, distance)

	return distance, nil
}
*/
