package usecase

import (
	"golang-rest-api-kata/internal/memory/entity"
	"golang-rest-api-kata/pkg/logger"
)

type Service struct {
	repo   Repository
	logger logger.Logger
}

func NewService(r Repository, logger logger.Logger) *Service {
	return &Service{
		repo:   r,
		logger: logger,
	}
}

func (mem *Service) GetInMemory(key string) (*entity.Memory, error) {
	panic("implement me")
}

func (mem *Service) CreateInMemory(key string, value string) (entity.Memory, error) {
	panic("implement me")
}
