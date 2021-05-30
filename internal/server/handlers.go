package server

import (
	"github.com/urfave/negroni"
	memoryHttp "golang-rest-api-kata/internal/memory/delivery/http"
	memoryRepository "golang-rest-api-kata/internal/memory/repository"
	memoryService "golang-rest-api-kata/internal/memory/usecase"
	apiMiddlewares "golang-rest-api-kata/internal/middleware"
	recordHttp "golang-rest-api-kata/internal/records/delivery/http"
	recordRepository "golang-rest-api-kata/internal/records/repository"
	recordService "golang-rest-api-kata/internal/records/usecase"
	"golang-rest-api-kata/pkg/utils"
	"net/http"
)

func (s *Server) MapHandlers() error {
	// Init repositories
	recordRepo := recordRepository.NewRecordRepository(s.db)
	memoryRepo := memoryRepository.NewMemoryRedisRepository(s.redisClient)

	// Init useCases
	recordUseCase := recordService.NewService(recordRepo, s.logger)
	memoryUseCase := memoryService.NewService(memoryRepo, s.logger)

	//Init Handlers
	recordHandlers := recordHttp.NewRecordHandlers(s.cfg, recordUseCase, s.logger)
	memoryHandlers := memoryHttp.NewMemoryHandlers(s.cfg, memoryUseCase, s.logger)

	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)

	s.httpRouter.PREFIX("/api/v1")

	n := negroni.New()
	n.Use(utils.JSONRecovery(true))
	n.UseFunc(mw.Cors)

	recordHttp.MapRecordRoutes(s.httpRouter, recordHandlers, n)
	memoryHttp.MapMemoryRoutes(s.httpRouter, memoryHandlers, n)

	s.httpRouter.GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("Health check")
		w.WriteHeader(http.StatusOK)
	}))

	return nil
}
