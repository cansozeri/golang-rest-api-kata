package entity

import (
	errors "golang-rest-api-kata/internal/errors/entity"
	"time"
)

type Record struct {
	Key        string    `bson:"key"`
	Counts     []Count   `bson:"counts"`
	Value      string    `bson:"value"`
	TotalCount int       `bson:"totalCount,omitempty"`
	CreatedAt  time.Time `bson:"createdAt"`
}

type Count struct {
	Count int
}

func NewRecord(key string, counts []Count, value string) (*Record, error) {
	r := &Record{
		Key:       key,
		Counts:    counts,
		Value:     value,
		CreatedAt: time.Now(),
	}
	err := r.Validate()
	if err != nil {
		return nil, errors.ErrInvalidEntity
	}
	return r, nil
}

func (r *Record) Validate() error {
	if r.Key == "" || r.Value == "" || len(r.Counts) == 0 {
		return errors.ErrInvalidEntity
	}
	return nil
}
