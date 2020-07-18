package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	s := newChatServer()
	s.OnReceive = func(c clientMessage) {
		fmt.Println("receive:" + string(c.Receive))
		c.Client.WriteByteArray([]byte("OK"))
	}
	go s.Serve()
	time.Sleep(1 * time.Second)

	c, _ := net.Dial("tcp", "127.0.0.1:8080")
	c.Write([]byte("hello world"))
	time.Sleep(1 * time.Second)
	var buf [8094]byte
	len, _ := c.Read(buf[:])
	fmt.Println(string(buf[:len]))
	c.Close()
}
