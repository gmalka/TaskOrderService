package redisservice

import (
	"fmt"
	"log"
	"time"
	"userService/internal/nosql"

	"github.com/go-redis/redis"
)

const (
	// key should be like `username:key`
	ACCESS = "access"
	REFRESH = "refresh"
)

type redisService struct {
	cli *redis.Client
}

func NewRedisService (cli *redis.Client) nosql.NoSqlService {
	return redisService{cli: cli}
}

func (r redisService) Get(key string, ttl int) (string, error) {
	result := r.cli.Get(key)

	if result == nil {
		return "", fmt.Errorf("cant find refresh token for key %s", key)
	}

	return result.String(), nil
}

func (r redisService) Set(key, value string, ttl int) error {
	status := r.cli.Set(key, value, time.Duration(ttl) * time.Minute)
	
	if status.Err() != nil {
		log.Println("Error while seting redis token for: ", key)
		return status.Err()
	}

	return nil
}

func (r redisService) Delete(key string) {
	r.cli.Del(key)
}