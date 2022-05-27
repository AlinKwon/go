package main

import "fmt"

func main() {
	fmt.Println("start logging sample")

	stdoutLoggingMain()

	interfaceLoggerMain()

	rotateExpireMain()

	logger := GetLogger()
	logger.Info("hi")
}
