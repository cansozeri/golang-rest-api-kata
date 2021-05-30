package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-rest-api-kata/internal/records/entity"
)

type recordsRepo struct {
	db *mongo.Database
}

func NewRecordRepository(db *mongo.Database) *recordsRepo {
	return &recordsRepo{db: db}
}

func (r *recordsRepo) Create(e *entity.Record) error {
	panic("implement me")
}

func (r *recordsRepo) Update(e *entity.Record) error {
	panic("implement me")
}

func (r *recordsRepo) Search(query string) ([]*entity.Record, error) {
	panic("implement me")
}

func (r *recordsRepo) List() ([]*entity.Record, error) {
	panic("implement me")
}
