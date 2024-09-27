package repository

import "github.com/shiron-dev/rapi/internal/infrastructure/infra"

type LoggerRepositoryImpl struct {
	Logger *infra.LoggerInterfaceImpl
}

func NewLoggerRepositoryImpl(logger *infra.LoggerInterfaceImpl) *LoggerRepositoryImpl {
	return &LoggerRepositoryImpl{Logger: logger}
}

func (l *LoggerRepositoryImpl) Info(format string, args ...interface{}) {
	l.Logger.Info(format, args...)
}

func (l *LoggerRepositoryImpl) Error(format string, args ...interface{}) {
	l.Logger.Error(format, args...)
}

func (l *LoggerRepositoryImpl) ErrorWithErr(err error) {
	l.Logger.ErrorWithErr(err)
}
