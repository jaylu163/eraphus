package config

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisDict = make(map[string]*redisProxy)
)

type redisProxy struct {
	host            string
	port            int
	password        string
	db              int
	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}
type RedisConn struct {
	client *redis.Client
}

func redisMap(list []RedisConf) {
	for _, item := range list {
		redisDict[item.Name] = &redisProxy{
			host:     item.Host,
			port:     item.Port,
			password: item.Password,
			db:       item.DB,
		}
	}
}

func InitRedis(name string) *RedisConn {
	client := new(RedisConn)
	if r, ok := redisDict[name]; ok {
		client = r.new()
	}
	return client
}

func (r *redisProxy) setAddr() string {
	return fmt.Sprintf("%d:%d", r.host, r.port)
}

func (r *redisProxy) new() *RedisConn {
	conn := redis.NewClient(&redis.Options{
		Addr:     r.setAddr(),
		Password: r.password,
		DB:       r.db,
	})
	client := &RedisConn{
		client: conn,
	}
	return client
}
