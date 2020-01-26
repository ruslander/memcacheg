package main

import (
	"log"
	"net"
	"os"
)
import "fmt"
import "bufio"

func main() {

	logger := log.New(os.Stdout, "sr ", log.Lmicroseconds)

	logger.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn, logger)
	}
}

func handleConnection(conn net.Conn, logger *log.Logger) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			conn.Close()
			logger.Println("Error, connection closed")
			return
		}

		logger.Print("Message Received: ", message)
		fmt.Fprintf(conn, message)
	}
}
