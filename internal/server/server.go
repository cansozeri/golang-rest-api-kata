package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-rest-api-kata/config"
	router "golang-rest-api-kata/internal/http"
	"golang-rest-api-kata/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	httpRouter  router.Router
	cfg         *config.Config
	db          *mongo.Database
	redisClient *redis.Client
	logger      logger.Logger
}

func NewServer(cfg *config.Config, db *mongo.Database, redisClient *redis.Client, logger logger.Logger) *Server {
	var port string

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = cfg.Server.Port
	}
	s := &http.Server{
		Addr:           port,
		ReadTimeout:    time.Second * cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	return &Server{httpRouter: router.NewMuxRouter(s), cfg: cfg, db: db, redisClient: redisClient, logger: logger}
}

func (s *Server) Run() (err error) {
	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err = s.httpRouter.SERVE(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err = s.MapHandlers(); err != nil {
		return err
	}

	s.handleSignals()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()

	err = s.httpRouter.SHUTDOWN(ctxShutDown)

	s.logger.Info("Server exited properly")

	return err
}

func (s *Server) handleSignals() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	s.logger.Info("Shutting down")
	go time.AfterFunc(15*time.Second, func() {
		os.Exit(1)
	})
}
