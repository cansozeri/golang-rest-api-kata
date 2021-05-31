package http

import (
	"golang-rest-api-kata/config"
	"golang-rest-api-kata/internal/records/delivery/presenter"
	"golang-rest-api-kata/internal/records/request"
	recordService "golang-rest-api-kata/internal/records/usecase"
	"golang-rest-api-kata/pkg/logger"
	"golang-rest-api-kata/pkg/utils"
	"golang-rest-api-kata/pkg/validator"
	"net/http"
)

type RecordHandlers struct {
	cfg      *config.Config
	recordUC recordService.UseCase
	logger   logger.Logger
}

func NewRecordHandlers(cfg *config.Config, recordUC recordService.UseCase, log logger.Logger) *RecordHandlers {
	return &RecordHandlers{cfg: cfg, recordUC: recordUC, logger: log}
}

func (rHandler *RecordHandlers) SearchRecords() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post request.SearchRecordRequest
		validate, err := validator.NewValidate()

		err = utils.ReadRequestBody(r.Body, &post, validate)
		if err != nil {
			rHandler.logger.Errorf("Error: %s", err.Error())
			_ = utils.Render(w, http.StatusBadRequest, validate.FormErrorMessage(err))
			return
		}

		records, err := rHandler.recordUC.SearchRecords(post)
		if err != nil {
			rHandler.logger.Errorf("Error: %s", err.Error())
			_ = utils.Render(w, http.StatusBadRequest, validate.FormErrorMessage(err))
			return
		}

		toJ := presenter.RecordPresenter{Msg: "Success"}

		for _, d := range records {
			toJ.Records = append(toJ.Records, &presenter.Record{
				Key:        d.Key,
				CreatedAt:  d.CreatedAt,
				TotalCount: d.TotalCount,
			})
		}

		_ = utils.Render(w, http.StatusOK, toJ)
	})
}
