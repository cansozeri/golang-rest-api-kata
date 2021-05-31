package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	errors "golang-rest-api-kata/internal/errors/entity"
	"golang-rest-api-kata/internal/memory/entity"
)

type RedisRepo struct {
	redisClient *redis.Client
}

func NewMemoryRedisRepository(redisClient *redis.Client) *RedisRepo {
	return &RedisRepo{redisClient: redisClient}
}

func (r *RedisRepo) Create(e *entity.Memory) (*entity.Memory, error) {

	if err := r.redisClient.Set(context.TODO(), e.Key, e.Value, 0).Err(); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *RedisRepo) Get(key string) (*entity.Memory, error) {
	value, err := r.redisClient.Get(context.TODO(), key).Result()

	if err == redis.Nil {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	record := &entity.Memory{
		Key:   key,
		Value: value,
	}

	return record, nil
}
