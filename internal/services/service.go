package services

import (
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"turlarion.ru/url-shortener/internal/base64"
)

type ShortenerService struct {
	rc redis.Client
}

func New(host string, port int, pwd string) ShortenerService {

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: pwd,
		DB:       0,
	})

	return ShortenerService{*client}
}

func (s ShortenerService) SaveNewUrl(url string, timeout int) (string, error) {
	key := base64.Generate(20)

	err := s.rc.Set(nil, key, url, time.Duration(timeout)).Err()

	if err != nil {
		return "", err
	}

	return key, nil

}

func (s ShortenerService) GetFullLink(id string) (string, error) {

	val, err := s.rc.Get(nil, id).Result()
	if err != nil {
		return "", err
	}
	return val, nil

}
