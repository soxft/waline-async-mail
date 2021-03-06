package redisutil

import (
	"github.com/gomodule/redigo/redis"
	"github.com/soxft/waline-async-mail/config"
	"log"
)

var R *redis.Pool

func Init() {
	if !config.Redis.Enable {
		return
	}

	R = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 50,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis.Addr,
				redis.DialPassword(config.Redis.Pwd),
				redis.DialDatabase(config.Redis.Db),
			)
			if err != nil {
				log.Fatalf(err.Error())
			}
			return c, err
		},
	}
	if _, err := R.Get().Do("PING"); err != nil {
		log.Fatalf("redis connect error: %s", err.Error())
	}
	log.Println("[INFO] Redis init success")
}
