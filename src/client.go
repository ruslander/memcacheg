package main

import (
	"flag"
	"fmt"
)

func main() {
	var messagesCount = flag.Int("messages", 100, "messages to send over")
	var messageSize = flag.Int("size", 100, "message size in bytes")

	flag.Parse()

	fmt.Print("workload of ", *messagesCount ," messages of size ", *messageSize)

	for i := 0; i < *messagesCount; i++ {
		fmt.Println("send ", i)
	}
}
