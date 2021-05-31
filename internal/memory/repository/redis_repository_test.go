package repository_test

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"golang-rest-api-kata/internal/memory/entity"
	"golang-rest-api-kata/internal/memory/repository"
	"log"
	"testing"
)

func SetupRedis() *repository.RedisRepo {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	redisRepository := repository.NewMemoryRedisRepository(client)

	return redisRepository
}

func TestRedisRepo_Get(t *testing.T) {
	t.Parallel()

	redisRepository := SetupRedis()

	t.Run("Create", func(t *testing.T) {
		record := &entity.Memory{
			Key:   "test",
			Value: "test",
		}
		createdRecord, err := redisRepository.Create(record)
		require.NoError(t, err)
		require.NotEqual(t, createdRecord, nil)

		result, err := redisRepository.Get(record.Key)
		require.NoError(t, err)
		require.Equal(t, "test", result.Key)
		require.Equal(t, "test", result.Value)
	})
}

func TestRedisRepo_Create(t *testing.T) {
	t.Parallel()

	redisRepository := SetupRedis()

	t.Run("Create", func(t *testing.T) {
		record := &entity.Memory{
			Key:   "test",
			Value: "test",
		}
		result, err := redisRepository.Create(record)
		require.NoError(t, err)
		require.NotEqual(t, result, nil)
	})
}
