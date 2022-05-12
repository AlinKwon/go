package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func rotateExpireMain() {
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic(err)
	}

	f := &lumberjack.Logger{
		Filename:   "logs/log.log",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	}
	mv := io.MultiWriter(f, os.Stdout)
	l := log.New(mv, "", 5)
	l.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	l.Println("hi file")
	l.Printf("%d file", 4)
}
