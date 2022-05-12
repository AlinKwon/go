package main

import (
	"log"
	"os"
)

type Logger interface {
	Dubug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type BuiltinLogger struct {
	logger *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *BuiltinLogger) Dubug(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func interfaceLoggerMain() {
	var logger Logger = NewBuiltinLogger()
	logger.Debugf("Hello with %s", "formatting")
}
