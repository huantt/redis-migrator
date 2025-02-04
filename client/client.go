package client

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/huantt/redis-migrator/config"
	"time"
)

func generateClient(connectionURL string, redisPassword string, redisDatabase int) (redis.Conn, error) {
	rdb, err := redis.Dial("tcp", connectionURL,
		redis.DialPassword(redisPassword),
		redis.DialDatabase(redisDatabase),
	)
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

// OldRedisClient generates connection for old redis
func OldRedisClient(redisConfig config.Configuration, redisDatabase int) (redis.Conn, error) {
	connectionURL := fmt.Sprintf("%s:%d", redisConfig.OldRedis.Host, redisConfig.OldRedis.Port)
	redisClient, err := generateClient(connectionURL, redisConfig.OldRedis.Password, redisDatabase)
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

// NewRedisClient generates connection for new redis
func NewRedisClient(redisConfig config.Configuration, redisDatabase int) (redis.Conn, error) {
	connectionURL := fmt.Sprintf("%s:%d", redisConfig.NewRedis.Host, redisConfig.NewRedis.Port)
	redisClient, err := generateClient(connectionURL, redisConfig.NewRedis.Password, redisDatabase)
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

func NewPool(redisConfig config.Redis, database int, maxIdle int, idleTimeout time.Duration) *redis.Pool {
	connectionURL := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				connectionURL,
				redis.DialDatabase(database),
				redis.DialPassword(redisConfig.Password),
			)
		},
	}
}
