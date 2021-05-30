package usecase

import (
	"golang-rest-api-kata/internal/records/entity"
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

func (rec *Service) SearchRecords(query string) ([]*entity.Record, error) {
	panic("implement me")
}

func (rec *Service) ListRecords() ([]*entity.Record, error) {
	panic("implement me")
}

func (rec *Service) CreateRecord(key string, counts []entity.Count, value string) error {
	panic("implement me")
}

func (rec *Service) UpdateRecord(e *entity.Record) error {
	panic("implement me")
}
