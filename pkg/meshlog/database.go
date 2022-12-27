package meshlog

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
)

type Database struct {
	enabled bool
	base    *log.Logger
}

func (d Database) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return d
}

func (d Database) Info(ctx context.Context, msg string, data ...interface{}) {
	d.base.Log(log.InfoLevel, "msg", data)
}

func (d Database) Warn(ctx context.Context, msg string, data ...interface{}) {
	d.base.Log(log.WarnLevel, "msg", data)
}

func (d Database) Error(ctx context.Context, msg string, data ...interface{}) {
	d.base.Log(log.ErrorLevel, "msg", data)
}

func (d Database) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}

func (l *Logger) DatabaseLogger() gormlogger.Interface {
	return &Database{
		enabled: true,
		// base:    l,
	}
}
