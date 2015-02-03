package db

import (
	"fmt"
	"github.com/goibibo/mantle"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goibibo/t-settings"
)


func MgetPool(vertical string) interface{} {
        configs := settings.GetConfigsFor("memcache", vertical)
        return GetConnection(createMemcachePool, configs)
}


func GetMemcacheClientFor(vertical string) mantle.Mantle {
        return MgetPool(vertical).(*mantle.Orm).New()
}


func createMemcachePool(configs dbConfig) interface{} {
	connectionUrl := settings.ConstructMemcachePath(configs)
	fmt.Println(connectionUrl)
	options := map[string]string{"db": "1"}
	pool := mantle.Orm{
                Driver:       "memcache",
                HostAndPorts: []string{connectionUrl},
                Capacity:     10,
                Options:      options}
        return &pool
}
