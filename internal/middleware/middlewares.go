package middleware

import (
	"golang-rest-api-kata/config"
	"golang-rest-api-kata/pkg/logger"
)

type MiddlewareManager struct {
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

func NewMiddlewareManager(cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, origins: origins, logger: logger}
}
