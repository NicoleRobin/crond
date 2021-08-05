package log

import (
	"fmt"

	"go.uber.org/zap"
)

var DefaultLogger *zap.Logger

func init() {
	var err error
	DefaultLogger, err = zap.NewDevelopment(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
	}
}

func Debug(msg string) {
	DefaultLogger.Debug(msg)
}

func Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	DefaultLogger.Debug(msg)
}

func Info(msg string) {
	DefaultLogger.Info(msg)
}

func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	DefaultLogger.Info(msg)
}
func Warn(msg string) {
	DefaultLogger.Warn(msg)
}

func Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	DefaultLogger.Warn(msg)
}

func Error(msg string) {
	DefaultLogger.Error(msg)
}

func Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	DefaultLogger.Error(msg)
}
