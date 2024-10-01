package usecase

import "github.com/shiron-dev/rapi/internal/adapter/repository"

type LoggerUsecase interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	ErrorWithErr(err error)
}

type LoggerUsecaseImpl struct {
	logger repository.LoggerRepository
}

func NewLoggerUsecase(logger repository.LoggerRepository) LoggerUsecase {
	return &LoggerUsecaseImpl{logger: logger}
}

func (l *LoggerUsecaseImpl) Info(format string, args ...interface{}) {
	l.logger.Info(format, args...)
}

func (l *LoggerUsecaseImpl) Error(format string, args ...interface{}) {
	l.logger.Error(format, args...)
}

func (l *LoggerUsecaseImpl) ErrorWithErr(err error) {
	l.logger.ErrorWithErr(err)
}
