package log

import (
	"context"
	"go.uber.org/zap"
	"time"

	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type GormLogger struct {
	logger *zapgorm2.Logger
}

func newGormLogger(logger *zap.Logger, levelName string) logger.Interface {
	zap2Logger := zapgorm2.New(logger)
	zap2Logger.LogLevel = resolveGormLogLevel(levelName)
	zap2Logger.SetAsDefault()
	return &GormLogger{&zap2Logger}
}

func resolveGormLogLevel(levelName string) logger.LogLevel {
	switch levelName {
	case "debug":
		return logger.Info
	default:
		return logger.Warn
	}
}

// gorm logger.Interface conformance

func (gl *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return gl.logger.LogMode(level)
}

func (gl *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	gl.logger.Info(ctx, s, i)
}

func (gl *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	gl.logger.Warn(ctx, s, i)
}

func (gl *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	gl.logger.Error(ctx, s, i)
}

func (gl *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	gl.logger.Trace(ctx, begin, fc, err)
}
