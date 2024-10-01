package repository

import "github.com/shiron-dev/rapi/internal/infrastructure/infra"

type LoggerRepository interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	ErrorWithErr(err error)
}

type LoggerRepositoryImpl struct {
	logger infra.LoggerInterface
}

func NewLoggerRepository(logger infra.LoggerInterface) LoggerRepository {
	return &LoggerRepositoryImpl{logger: logger}
}

func (l *LoggerRepositoryImpl) Info(format string, args ...interface{}) {
	l.logger.Info(format, args...)
}

func (l *LoggerRepositoryImpl) Error(format string, args ...interface{}) {
	l.logger.Error(format, args...)
}

func (l *LoggerRepositoryImpl) ErrorWithErr(err error) {
	l.logger.ErrorWithErr(err)
}
