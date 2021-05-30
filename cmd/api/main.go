package main

import (
	"golang-rest-api-kata/config"
	"golang-rest-api-kata/internal/server"
	"golang-rest-api-kata/pkg/db/mongo"
	"golang-rest-api-kata/pkg/db/redis"
	"golang-rest-api-kata/pkg/logger"
	"golang-rest-api-kata/pkg/utils"
	"log"
	"os"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	mongoDB, err := mongo.NewMongoDb(cfg)
	if err != nil {
		appLogger.Fatalf("Mongo init: %s", err)
	} else {
		appLogger.Infof("Mongo connected, Status: %#v", mongoDB.Name())
	}
	defer mongo.Disconnect(mongoDB.Client())

	redisClient := redis.NewRedisClient(cfg)
	defer redis.Disconnect(redisClient)

	appLogger.Info("Redis connected")

	s := server.NewServer(cfg, mongoDB, redisClient, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
