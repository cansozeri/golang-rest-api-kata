package logger

import (
	"github.com/labstack/gommon/log"
)

type Fields map[string]interface{}

type InfoLogger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoj(j log.JSON)
}

type Logger interface {
	InfoLogger

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugj(j log.JSON)

	Print(i ...interface{})
	Printf(format string, args ...interface{})
	Printj(j log.JSON)

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoj(j log.JSON)

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnj(j log.JSON)

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorj(j log.JSON)

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicj(j log.JSON)

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalj(j log.JSON)

	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}
