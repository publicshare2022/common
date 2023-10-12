package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	LogLevel logger.LogLevel
}

func NewGormLogger() logger.Interface {
	return new(GormLogger)
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Debugf(msg, data)
}
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(msg, data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	msg := fmt.Sprintf("Trace sql: %v row: %v", sql, rows)
	if err != nil {
		msg += fmt.Sprintf(" error: %v", err)
	}
	logx.WithContext(ctx).WithDuration(elapsed).Slow(msg)
}
