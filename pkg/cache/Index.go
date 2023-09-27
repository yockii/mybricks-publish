package cache

import (
	"fmt"
	"manatee-publish/pkg/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Prefix string
var Redis *redis.Pool
var enabled bool

func init() {
	config.DefaultInstance.SetDefault("redis.app", "ruomu")
	config.DefaultInstance.SetDefault("redis.host", "localhost")
	config.DefaultInstance.SetDefault("redis.port", "6379")

	InitRedis(
		config.GetString("redis.app"),
		config.GetString("redis.host"),
		config.GetString("redis.password"),
		config.GetInt("redis.port"),
		config.GetInt("redis.maxIdle"),
		config.GetInt("redis.maxActive"),
	)
}

func InitRedis(redisPrefix, host, password string, port, maxIdle, maxActive int, options ...redis.DialOption) {
	Prefix = redisPrefix
	if password != "" {
		options = append(options, redis.DialPassword(password))
	}
	Redis = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", host, port),
				options...)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
	}
	enabled = true
}

func Get() redis.Conn {
	return Redis.Get()
}

func Close() {
	_ = Redis.Close()
}

func Enabled() bool {
	return enabled
}
