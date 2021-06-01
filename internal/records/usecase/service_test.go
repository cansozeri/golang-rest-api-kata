package usecase_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang-rest-api-kata/config"
	"golang-rest-api-kata/internal/records/delivery/presenter"
	"golang-rest-api-kata/internal/records/entity"
	"golang-rest-api-kata/internal/records/mock"
	"golang-rest-api-kata/internal/records/request"
	"golang-rest-api-kata/internal/records/usecase"
	"golang-rest-api-kata/pkg/logger"
	"testing"
	"time"
)

func TestService_SearchRecords(t *testing.T) {
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
	mockRecordRepo := mock.NewMockRepository(controller)
	recordUC := usecase.NewService(mockRecordRepo, apiLogger)

	query := request.SearchRecordRequest{
		StartDate: "2016-01-01",
		EndDate:   "2016-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}

	recordList := []*entity.Record{{
		Key:        "test",
		CreatedAt:  time.Now(),
		TotalCount: 2812,
	}}

	mockRecordRepo.EXPECT().Search(query).Return(recordList, nil)

	recordList, err := recordUC.SearchRecords(query)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, recordList)

	toJ := presenter.RecordPresenter{Msg: "Success"}

	for _, d := range recordList {
		toJ.Records = append(toJ.Records, &presenter.Record{
			Key:        d.Key,
			CreatedAt:  d.CreatedAt,
			TotalCount: d.TotalCount,
		})
	}

	require.Equal(t, 0, toJ.Code)
	require.Equal(t, "Success", toJ.Msg)
	require.Equal(t, 1, len(toJ.Records))
}
