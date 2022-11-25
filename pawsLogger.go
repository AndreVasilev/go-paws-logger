package log

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func CreateLogger(logDir string, loggerName string, levelName string) (*zap.Logger, error) {
	logFile := logDir + "/%Y-%m-%d/" + loggerName + ".json"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(60*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour))
	if err != nil {
		return nil, err
	}

	// add the encoder config and rotator to create a new zap logger
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		resolveLogLevel(levelName))
	logger := zap.New(core)
	return logger, nil
}
