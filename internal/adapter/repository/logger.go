package repository

import "github.com/shiron-dev/rapi/internal/infrastructure/infra"

type LoggerRepositoryImpl struct {
	logger *infra.LoggerInterfaceImpl
}

func NewLoggerRepositoryImpl(logger *infra.LoggerInterfaceImpl) *LoggerRepositoryImpl {
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
