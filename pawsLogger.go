package log

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

func resolveLogLevel(levelName string) zapcore.Level {
	switch levelName {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "warn":
		return zap.WarnLevel
	case "panic":
		return zap.PanicLevel
	case "off":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}

func CreateLogger(logDir string, loggerName string, levelName string, maxDaysAge int) (*zap.Logger, error) {
	logFilePath := logDir + "/%Y-%m-%d/" + loggerName + ".json"
	rotator, err := rotatelogs.New(
		logFilePath,
		rotatelogs.WithMaxAge(time.Duration(maxDaysAge)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour))
	if err != nil {
		return nil, err
	}

	// add the encoder config and rotator to create a new zap logger
	w := zapcore.AddSync(rotator)

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	writer := w
	defaultLogLevel := resolveLogLevel(levelName)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger, nil
}

func CreateGormLogger(logDir string, loggerName string, levelName string, maxDaysAge int) (gormlogger.Interface, error) {
	logger, err := CreateLogger(logDir, loggerName, levelName, maxDaysAge)
	if err != nil {
		return nil, err
	}
	gormLogger := newGormLogger(logger, levelName)
	return gormLogger, nil
}
