package usecase

import "golang-rest-api-kata/internal/memory/entity"

//Reader interface
type Reader interface {
	Get(key string) (*entity.Memory, error)
}

//Writer book writer
type Writer interface {
	Create(e *entity.Memory) (entity.Memory, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetInMemory(key string) (*entity.Memory, error)
	CreateInMemory(key string, value string) (entity.Memory, error)
}
