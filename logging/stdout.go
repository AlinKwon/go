package main

import (
	"log"
	"os"
)

// Stdout logging sample

var logger = log.New(os.Stdout, "", 5)

func stdoutLoggingMain() {
	logger.Println("Hello, stdout logging")
	//or
	log.Println("Hello, stdout logging")
}
