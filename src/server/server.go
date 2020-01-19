package main

import "net"
import "fmt"
import "bufio"

func main() {

	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			conn.Close()
			fmt.Println("Error, connection closed")
			return
		}

		fmt.Print("Message Received: ", message)
		fmt.Fprintf(conn, message, "\n")
	}
}
