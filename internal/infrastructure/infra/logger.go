package infra

import "fmt"

type LoggerInterfaceImpl struct{}

func NewLoggerInterfaceImpl() *LoggerInterfaceImpl {
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
