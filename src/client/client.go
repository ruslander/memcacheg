package main

import (
	"flag"
	"github.com/codahale/hdrhistogram"
	"log"
	"math/rand"
	"os"
	"time"
	"transport"
)

var messagesCount = flag.Int("messages", 100, "messages to send over")
var messageSize = flag.Int("size", 100, "message size in bytes")
var serverIp = flag.String("server", "127.0.0.1", "server Ip address")
var serverPort = flag.String("port", "8081", "server port")

func main() {

	logger := log.New(os.Stdout, "cl ", log.Lmicroseconds)

	flag.Parse()
	logger.Print("workload of ", *messagesCount ," messages of size ", *messageSize)

	link := transport.New(*serverIp + ":" + *serverPort)
	defer link.Close()

	msg := "msg#" + RandStringRunes(*messageSize)

	hist := hdrhistogram.New(1, 1000, 2)


	for i := 0; i < *messagesCount; i++ {
		start := time.Now()
		link.Send(msg)

		link.Receive()
		duration := time.Since(start)

		hist.RecordValue(duration.Microseconds())
	}

	printSummary(hist, logger)
}

func printSummary(hist *hdrhistogram.Histogram, logger *log.Logger) {
	logger.Printf("Latency distribution (usec)")
	logger.Printf("%5s, %10s, %5s", "Value", "Percentile", "TotalCount")

	for _, s := range hist.CumulativeDistribution() {
		logger.Printf("%5d, %10.2f, %9d", s.ValueAt, s.Quantile, s.Count)
	}

	logger.Printf("Min = %d, Max = %d , Mean = %f, StdDeviation = %f",
		hist.Min(), hist.Max(), hist.Mean(), hist.StdDev())

	logger.Printf("Throughput %5.1f op/sec,", (1000000)/hist.Mean())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")



func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
