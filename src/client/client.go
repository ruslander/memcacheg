package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"transport"
)

var messagesCount = flag.Int("messages", 100, "messages to send over")
var messageSize = flag.Int("size", 100, "message size in bytes")

func main() {

	logger := log.New(os.Stdout, "cl ", log.Lmicroseconds)

	flag.Parse()
	logger.Print("workload of ", *messagesCount ," messages of size ", *messageSize)

	link := transport.New("127.0.0.1:8081")
	defer link.Close()

	msg := "msg#" + RandStringRunes(*messageSize)

	for i := 0; i < *messagesCount; i++ {
		logger.Println("Send: ", msg)

		start := time.Now()
		link.Send(msg)

		message := link.Receive()
		elapsed := time.Since(start)

		logger.Printf("Receive: " + strings.TrimSuffix(message, "\n") + " " + elapsed.String())
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
