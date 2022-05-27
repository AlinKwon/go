package main

import (
	"io"
	"log"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

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
