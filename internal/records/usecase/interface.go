package usecase

import "golang-rest-api-kata/internal/records/entity"

//Reader interface
type Reader interface {
	Search(query string) ([]*entity.Record, error)
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
	SearchRecords(query string) ([]*entity.Record, error)
	ListRecords() ([]*entity.Record, error)
	CreateRecord(key string, counts []entity.Count, value string) error
	UpdateRecord(e *entity.Record) error
}
