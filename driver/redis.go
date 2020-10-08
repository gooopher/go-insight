package driver

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/config"
	"reflect"
	"sync"
)

var redisMap sync.Map
var redisClusterMap sync.Map

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db,string"`
}

func GetRedis(key string) (client *redis.Client, err error) {
	clientIfc, ok := redisMap.Load(key)
	if !ok {
		client, err = getRedis(key)
		if err != nil {
			return
		}
		redisMap.Store(key, client)
	} else {
		client = clientIfc.(*redis.Client)
	}
	return
}

func getRedis(key string) (*redis.Client, error) {
	var conf RedisConfig
	if err := config.Get("redis", key).Scan(&conf); err != nil {
		return nil, err
	}
	if reflect.DeepEqual(conf, RedisConfig{}) {
		return nil, fmt.Errorf("not found redis config: %s", key)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

type RedisClusterConfig struct {
	Addrs    []string `json:"addrs"`
	Password string   `json:"password"`
}

func GetRedisCluster(key string) (client *redis.ClusterClient, err error) {
	clientIfc, ok := redisClusterMap.Load(key)
	if !ok {
		client, err = getRedisCluster(key)
		if err != nil {
			return
		}
		redisClusterMap.Store(key, client)
	} else {
		client = clientIfc.(*redis.ClusterClient)
	}
	return
}

func getRedisCluster(key string) (*redis.ClusterClient, error) {
	var conf RedisClusterConfig
	if err := config.Get("redis_cluster", key).Scan(&conf); err != nil {
		return nil, err
	}
	if reflect.DeepEqual(conf, RedisClusterConfig{}) {
		return nil, fmt.Errorf("not found redis_cluster config: %s", key)
	}
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    conf.Addrs,
		Password: conf.Password,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
