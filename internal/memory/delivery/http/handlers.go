package http

import (
	"golang-rest-api-kata/config"
	"golang-rest-api-kata/internal/memory/entity"
	"golang-rest-api-kata/internal/memory/request"
	memoryService "golang-rest-api-kata/internal/memory/usecase"
	"golang-rest-api-kata/pkg/logger"
	"golang-rest-api-kata/pkg/utils"
	"golang-rest-api-kata/pkg/validator"
	"net/http"
)

type MemoryHandlers struct {
	cfg      *config.Config
	memoryUC memoryService.UseCase
	logger   logger.Logger
}

func NewMemoryHandlers(cfg *config.Config, memoryUC memoryService.UseCase, log logger.Logger) *MemoryHandlers {
	return &MemoryHandlers{cfg: cfg, memoryUC: memoryUC, logger: log}
}

func (mHandler *MemoryHandlers) GetInMemory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var parameters request.GetInMemoryRequest
		validate, err := validator.NewValidate()

		err = utils.ReadRequest(r, &parameters, validate)

		if err != nil {
			_ = utils.Render(w, http.StatusBadRequest, validate.FormErrorMessage(err))
			return
		}

		result, err := mHandler.memoryUC.GetInMemory(parameters.Key)

		if err != nil {
			mHandler.logger.Errorf("Error: %s", err.Error())
			_ = utils.Render(w, http.StatusBadRequest, validate.FormErrorMessage(err))
			return
		}

		_ = utils.Render(w, http.StatusOK, result)

	})
}

func (mHandler *MemoryHandlers) CreateInMemory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post entity.Memory
		validate, err := validator.NewValidate()

		err = utils.ReadRequestBody(r.Body, &post, validate)
		if err != nil {
			mHandler.logger.Errorf("Error: %s", err.Error())
			_ = utils.Render(w, http.StatusBadRequest, validate.FormErrorMessage(err))
			return
		}

		result, err := mHandler.memoryUC.CreateInMemory(post.Key, post.Value)

		if err != nil {
			mHandler.logger.Errorf("Error: %s", err.Error())
			_ = utils.Render(w, http.StatusBadRequest, validate.FormErrorMessage(err))
			return
		}

		_ = utils.Render(w, http.StatusOK, result)
	})
}
