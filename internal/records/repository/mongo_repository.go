package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-rest-api-kata/internal/records/entity"
	"golang-rest-api-kata/internal/records/request"
	"time"
)

type recordsRepo struct {
	db *mongo.Database
}

const (
	layoutISO = "2006-01-02"
)

func NewRecordRepository(db *mongo.Database) *recordsRepo {
	return &recordsRepo{db: db}
}

func (r *recordsRepo) Create(e *entity.Record) error {
	panic("implement me")
}

func (r *recordsRepo) Update(e *entity.Record) error {
	panic("implement me")
}

func (r *recordsRepo) Search(request request.SearchRecordRequest) (records []*entity.Record, err error) {
	collection := r.db.Collection("records")
	startDate, _ := time.Parse(layoutISO, request.StartDate)
	endDate, _ := time.Parse(layoutISO, request.EndDate)

	pipeline := mongo.Pipeline{
		{{"$match", bson.M{"createdAt": bson.M{"$gt": startDate, "$lt": endDate}}}},
		{{"$project", bson.M{"key": "$key", "createdAt": "$createdAt", "totalCount": bson.M{"$sum": "$counts"}}}},
		{{"$match", bson.M{"totalCount": bson.M{"$gt": request.MinCount, "$lt": request.MaxCount}}}},
		{{"$sort", bson.M{"totalCount": 1}}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &records)
	if err != nil {
		return nil, err
	}

	return records, nil

}

func (r *recordsRepo) List() ([]*entity.Record, error) {
	panic("implement me")
}
