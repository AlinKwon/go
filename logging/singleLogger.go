package main

import (
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
var logger *log.Logger
var once sync.Once

func GetLogger() *log.Logger {
	once.Do(func() {
		if err := os.MkdirAll("logs", os.ModePerm); err != nil {
			panic(err)
		}
		f := &lumberjack.Logger{
			Filename:   "logs/log.log",
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		}

		mr := io.MultiWriter(f, os.Stdout)

		logger = log.New(mr, " ", 5)
		logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	})
	return logger
}
*/

var singleLogger *zap.SugaredLogger
var once sync.Once

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		if err := os.MkdirAll("logs", os.ModePerm); err != nil {
			panic(err)
		}
		f := &lumberjack.Logger{
			Filename:   "logs/log.log",
			MaxSize:    1,  // megabytes after which new file is created
			MaxBackups: 3,  // number of backups
			MaxAge:     28, //days
		}

		mr := io.MultiWriter(f, os.Stdout)

		w := zapcore.AddSync(mr)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.DebugLevel,
		)
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0)) // skip : 0(current caller), 1 (current method caller)
		singleLogger = logger.Sugar()
		// singleLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	})
	return singleLogger
}
