package usecase

import (
	"golang-rest-api-kata/internal/records/entity"
	"golang-rest-api-kata/internal/records/request"
)

//Reader interface
type Reader interface {
	Search(request request.SearchRecordRequest) ([]*entity.Record, error)
	List() ([]*entity.Record, error)
}

//Writer record writer
type Writer interface {
	Create(e *entity.Record) error
	Update(e *entity.Record) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	SearchRecords(request request.SearchRecordRequest) ([]*entity.Record, error)
	ListRecords() ([]*entity.Record, error)
	CreateRecord(key string, counts []entity.Count, value string) error
	UpdateRecord(e *entity.Record) error
}
