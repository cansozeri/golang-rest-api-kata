package http_test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang-rest-api-kata/config"
	recordHttp "golang-rest-api-kata/internal/records/delivery/http"
	"golang-rest-api-kata/internal/records/delivery/presenter"
	"golang-rest-api-kata/internal/records/entity"
	"golang-rest-api-kata/internal/records/mock"
	"golang-rest-api-kata/internal/records/request"
	"golang-rest-api-kata/pkg/logger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRecordHandlers_SearchRecords(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)

	cfg := &config.Config{
		Logger: config.Logger{
			Development: true,
		},
	}

	apiLogger := logger.NewApiLogger(cfg)
	recordHandlers := recordHttp.NewRecordHandlers(cfg, service, apiLogger)

	e := &entity.Record{
		Key:        "test",
		CreatedAt:  time.Now(),
		TotalCount: 2812,
	}
	service.EXPECT().
		SearchRecords(request.SearchRecordRequest{
			StartDate: "2016-01-01",
			EndDate:   "2016-02-02",
			MinCount:  2700,
			MaxCount:  3000,
		}).
		Return([]*entity.Record{e}, nil)

	ts := httptest.NewServer(recordHandlers.SearchRecords())
	defer ts.Close()

	payload := fmt.Sprintf(`{
	   "startDate": "2016-01-01",
		"endDate": "2016-02-02",
		"minCount": 2700,
		"maxCount": 3000
	}`)

	resp, err := http.Post(ts.URL+"/records/search", "application/json", strings.NewReader(payload))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	r := &presenter.RecordPresenter{}
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 0, r.Code)
	assert.Equal(t, "Success", r.Msg)
}
