package db

import (
	"github.com/goibibo/mantle"
	mRedis "github.com/goibibo/mantle/backends"
	"github.com/goibibo/t-settings"
	"strconv"
)

const (
	pool_size  = "10"
	default_db = "0"
)

func getPool(vertical string) interface{} {
	configs := settings.GetConfigsFor("redis", vertical)
	return GetConnection(createRedisPool, configs)
}

func PureRedisClientFor(vertical string) (*mRedis.RedisConn, error) {
	return getPool(vertical).(*mantle.Orm).GetRedisConn()
}

func GetRedisClientFor(vertical string) mantle.Mantle {
	return getPool(vertical).(*mantle.Orm).New()
}

func foundOrSetDefault(configs dbConfig, key string, fallback string) string {
	value, ok := configs[key]
	if !ok {
		value = fallback
	}
	return value
}

func createRedisPool(configs dbConfig) interface{} {
	connectionUrl := settings.ConstructRedisPath(configs)
	db := foundOrSetDefault(configs, "db", default_db)
	capacity, _ := strconv.Atoi(foundOrSetDefault(configs, "pool_size", pool_size))
	options := map[string]string{"db": db}
	pool := mantle.Orm{
		Driver:       "redis",
		HostAndPorts: []string{connectionUrl},
		Capacity:     capacity,
		Options:      options}
	return &pool
}
