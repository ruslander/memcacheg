package transport

import (
	"bufio"
	"fmt"
	"net"
)

type TcpClient struct {
	address    string
	connection net.Conn
}

func (c TcpClient) Close() {
	c.connection.Close()
}

func (c TcpClient) Send(msg string) {
	fmt.Fprintf(c.connection, msg + "\n")
}

func (c TcpClient) Receive() string {
	message, _ := bufio.NewReader(c.connection).ReadString('\n')
	return message
}

func New(addr string) *TcpClient {
	conn, _ := net.Dial("tcp", addr)

	return &TcpClient{
		address: addr,
		connection: conn,
	}
}
