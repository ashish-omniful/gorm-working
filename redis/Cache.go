package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm/models"
	"time"
)

func getClient() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SetCache(ctx context.Context, cacheKey string, data interface{}) error {

	client := getClient()

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = client.Set(ctx, cacheKey, string(jsonData), time.Second*1000).Err()
	if errors.Is(err, redis.Nil) {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func GetCache(ctx context.Context, cacheKey string) (*models.User, error) {

	client := getClient()

	val, err := client.Get(ctx, cacheKey).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	user := models.User{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
