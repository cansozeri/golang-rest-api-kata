package logger

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"golang-rest-api-kata/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type apiLogger struct {
	Logger
	level       int
	verbosity   int
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

func NewZapApiLogger(cfg *config.Config) *apiLogger {
	return &apiLogger{cfg: cfg}
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *apiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	_ = l.sugarLogger.Sync()
	// TODO: fix error
	// -> https://github.com/uber-go/zap/issues/772
	// -> https://github.com/uber-go/zap/issues/328
}

func (l *apiLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *apiLogger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

func (l *apiLogger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.sugarLogger.Debugw(string(b))
}

func (l *apiLogger) Print(i ...interface{}) {
	l.Info(i...)
}

func (l *apiLogger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func (l *apiLogger) Printj(j log.JSON) {
	l.Infoj(j)
}

func (l *apiLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *apiLogger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *apiLogger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.sugarLogger.Infow(string(b))
}

func (l *apiLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *apiLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *apiLogger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.sugarLogger.Warnw(string(b))
}

func (l *apiLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *apiLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

func (l *apiLogger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.sugarLogger.Errorw(string(b))
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *apiLogger) Panicf(format string, args ...interface{}) {
	l.sugarLogger.Panicf(format, args...)
}

func (l *apiLogger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.sugarLogger.Panicw(string(b))
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *apiLogger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

func (l *apiLogger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.sugarLogger.Fatalw(string(b))
}

func (l *apiLogger) V(level int) InfoLogger {
	return &apiLogger{
		level:       level,
		verbosity:   l.verbosity,
		sugarLogger: l.sugarLogger,
	}
}

func (l *apiLogger) WithField(key string, value interface{}) Logger {

	fields := Fields{
		key: value,
	}

	f := prepareForWith(fields)

	return &apiLogger{
		level:       l.level,
		verbosity:   l.verbosity,
		sugarLogger: l.sugarLogger.With(f...),
	}
}

func (l *apiLogger) WithFields(fields map[string]interface{}) Logger {

	f := prepareForWith(fields)

	return &apiLogger{
		level:       l.level,
		verbosity:   l.verbosity,
		sugarLogger: l.sugarLogger.With(f...),
	}
}

func (l *apiLogger) Enabled() bool {
	return l.level <= l.verbosity
}

func prepareForWith(fields Fields) []interface{} {

	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}

	return f
}
