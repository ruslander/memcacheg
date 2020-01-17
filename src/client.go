package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net"
)

func main() {
	var messagesCount = flag.Int("messages", 100, "messages to send over")
	var messageSize = flag.Int("size", 100, "message size in bytes")

	flag.Parse()

	fmt.Print("workload of ", *messagesCount ," messages of size ", *messageSize)

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()

	msg := "msg#" + RandStringRunes(*messageSize)

	for i := 0; i < *messagesCount; i++ {

		fmt.Println("Send: ", msg)
		fmt.Fprintf(conn, msg + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Receive: "+message)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
