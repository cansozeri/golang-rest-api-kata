package usecase_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang-rest-api-kata/config"
	"golang-rest-api-kata/internal/memory/entity"
	"golang-rest-api-kata/internal/memory/mock"
	"golang-rest-api-kata/internal/memory/request"
	"golang-rest-api-kata/internal/memory/usecase"
	"golang-rest-api-kata/pkg/logger"
	"testing"
)

func TestService_CreateInMemory(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	defer controller.Finish()

	cfg := &config.Config{
		Logger: config.Logger{
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: false,
			Encoding:          "json",
		},
	}

	apiLogger := logger.NewApiLogger(cfg)
	mockMemoryRepo := mock.NewMockRepository(controller)
	memoryUC := usecase.NewService(mockMemoryRepo, apiLogger)

	query := &entity.Memory{
		Key:   "test",
		Value: "test",
	}

	record := &entity.Memory{
		Key:   "test",
		Value: "test",
	}

	mockMemoryRepo.EXPECT().Create(query).Return(record, nil)

	record, err := memoryUC.CreateInMemory(query.Key, query.Value)

	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, record)

	require.Equal(t, "test", record.Key)
	require.Equal(t, "test", record.Value)

}

func TestService_GetInMemory(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	defer controller.Finish()

	cfg := &config.Config{
		Logger: config.Logger{
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: false,
			Encoding:          "json",
		},
	}

	apiLogger := logger.NewApiLogger(cfg)
	mockMemoryRepo := mock.NewMockRepository(controller)
	memoryUC := usecase.NewService(mockMemoryRepo, apiLogger)

	query := request.GetInMemoryRequest{
		Key: "test",
	}

	record := &entity.Memory{
		Key:   "test",
		Value: "test",
	}

	mockMemoryRepo.EXPECT().Get(query.Key).Return(record, nil)

	record, err := memoryUC.GetInMemory(query.Key)

	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, record)

	require.Equal(t, "test", record.Key)
	require.Equal(t, "test", record.Value)

}
