package db

import (
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"time"
)

type gormLogger struct {
}

// createGormLogger returns a gorm logger that logs to zerolog
func createGormLogger() logger.Interface {
	return &gormLogger{}
}

func (g gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	//TODO implement me
	panic("implement me")
}

func (g gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	log.Info().Msg(s)
}

func (g gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	log.Warn().Msg(s)
}

func (g gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	log.Error().Msg(s)
}

func (g gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	if err != nil {
		log.Error().Str("sql", sql).Int64("rows", rowsAffected).Err(err).Msg("failed to execute sql statement")
	} else {
		log.Trace().Str("sql", sql).Int64("rows", rowsAffected).Err(err).Msg("trace sql statement")
	}
}
