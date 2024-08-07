package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"turlarion.ru/url-shortener/internal/base64"
)

type ShortenerService struct {
	rc  redis.Client
	ctx context.Context
}

func New(host string, port int, pwd string) ShortenerService {

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: pwd,
		DB:       0,
	})

	return ShortenerService{rc: *client, ctx: context.TODO()}
}

func (s ShortenerService) SaveNewUrl(url string, timeout int) (string, error) {
	key := base64.Generate(20)

	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		url = "http://" + url
	}

	fmt.Println(url)

	err := s.rc.Set(s.ctx, key, url, time.Duration(timeout)*time.Second).Err()

	if err != nil {
		return "", err
	}

	return key, nil

}

func (s ShortenerService) GetFullLink(id string) (string, error) {

	val, err := s.rc.Get(s.ctx, id).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (s ShortenerService) DeleteShortLink(id string) error {
	_, err := s.rc.Del(s.ctx, id).Result()
	return err
}
