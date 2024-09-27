package infra

import "fmt"

type LoggerInterface interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	ErrorWithErr(err error)
}

type LoggerInterfaceImpl struct{}

func NewLoggerInterface() LoggerInterface {
	return &LoggerInterfaceImpl{}
}

func (l *LoggerInterfaceImpl) Info(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *LoggerInterfaceImpl) Error(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *LoggerInterfaceImpl) ErrorWithErr(err error) {
	panic(err)
}
