package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	fmt.Println("hi")

	c := &serial.Config{Name: "COM5", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%q", buf[:n])
}
