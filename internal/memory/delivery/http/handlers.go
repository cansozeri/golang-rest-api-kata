package http

import (
	"golang-rest-api-kata/config"
	memoryService "golang-rest-api-kata/internal/memory/usecase"
	"golang-rest-api-kata/pkg/logger"
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
		mHandler.logger.Info("redis alınıyor")
	})
}

func (mHandler *MemoryHandlers) CreateInMemory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mHandler.logger.Info("redis yaratılıyor")
	})
}
