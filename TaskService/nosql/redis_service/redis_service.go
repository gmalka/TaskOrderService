package redisservice

import (
	"fmt"
	"log"
	"taskServer/nosql"
	"time"

	"github.com/go-redis/redis"
)

const TTL = 5

type redisService struct {
	cli *redis.Client
}

func NewRedisService (addr string) (nosql.NoSqlService, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return redisService{cli: client}, nil
}

func (r redisService) Get(username string, page int) ([]byte, error) {
	key := fmt.Sprintf("%s:%d", username, page)
	result := r.cli.Get(key)

	if result == nil {
		return nil, fmt.Errorf("cant find refresh token for key %s", key)
	}

	return result.Bytes()
}

func (r redisService) Set(username string, page int, val []byte) error {
	key := fmt.Sprintf("%s:%d", username, page)
	status := r.cli.Set(key, val, time.Duration(TTL) * time.Minute)
	
	if status.Err() != nil {
		log.Println("Error while seting redis token for: ", key)
		return status.Err()
	}

	return nil
}

func (r redisService) Delete(username string) {
	r.cli.Del(fmt.Sprintf("%s:*", username))
}