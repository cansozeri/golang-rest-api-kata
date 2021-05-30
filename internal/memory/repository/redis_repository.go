package repository

import (
	"github.com/go-redis/redis/v8"
	"golang-rest-api-kata/internal/memory/entity"
)

type RedisRepo struct {
	redisClient *redis.Client
}

func NewMemoryRedisRepository(redisClient *redis.Client) *RedisRepo {
	return &RedisRepo{redisClient: redisClient}
}

func (r *RedisRepo) Create(e *entity.Memory) (entity.Memory, error) {
	panic("implement me")
}

func (r *RedisRepo) Get(key string) (*entity.Memory, error) {
	panic("implement me")
}
