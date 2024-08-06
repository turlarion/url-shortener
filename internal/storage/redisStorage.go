package storage

import (
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	host   string
	port   int
	pass   string
	client redis.Client
}

func (rs *RedisStorage) OpenConnection() {
	rs.client = *redis.NewClient(&redis.Options{
		Addr:     rs.host + ":" + strconv.Itoa(rs.port),
		Password: rs.pass,
		DB:       0,
	})
}

func (rs RedisStorage) SaveNewUrl(key, url string, timeout int) {
	rs.client.Set(nil, key, url, time.Duration(timeout))
}
