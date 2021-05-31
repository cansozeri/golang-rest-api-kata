package usecase

import (
	errors "golang-rest-api-kata/internal/errors/entity"
	"golang-rest-api-kata/internal/records/entity"
	"golang-rest-api-kata/internal/records/request"
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

func (rec *Service) SearchRecords(request request.SearchRecordRequest) ([]*entity.Record, error) {
	records, err := rec.repo.Search(request)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, errors.ErrNotFound
	}
	return records, nil
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
