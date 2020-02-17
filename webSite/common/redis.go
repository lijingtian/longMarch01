package common

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisConn redis.Conn

var redisHost = "127.0.0.1:6379"

func init() {
	if redisConn == nil {
		var err error
		redisConn, err = redis.Dial("tcp", redisHost)
		if err != nil {
			fmt.Println("redis.Dial err", err.Error())
			return
		}
	}
}

func GetRedisConn() redis.Conn {
	if redisConn == nil {
		dialRedis()
	}
	return redisConn
}

func dialRedis() {
	if redisConn == nil {
		var err error
		redisConn, err = redis.Dial("tcp", redisHost)
		if err != nil {
			fmt.Println("redis.Dial err", err.Error())
		}
	}
}
