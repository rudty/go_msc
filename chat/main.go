package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

func main() {
	s := newChatServer()
	s.OnConnected = func(c *client) {
		fmt.Println(fmt.Sprintf("hello %d", c.ClientID))
	}
	s.OnClosed = func(c *client) {
		fmt.Println(fmt.Sprintf("bye %d", c.ClientID))
	}
	s.OnReceive = func(m request) {
		s.Range(func(c *client) {
			c.WriteMessage(m.Receive)
		})
	}
	go s.Serve()
	time.Sleep(500 * time.Millisecond)

	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		idx := i
		wg.Add(1)
		go func() {
			c, _ := net.Dial("tcp", "127.0.0.1:8080")
			time.Sleep(50 * time.Millisecond)
			msg := []byte("hello world" + strconv.Itoa(idx))
			var header [4]byte
			header[0] = byte(len(msg))
			c.Write(header[:])
			c.Write(msg)
			var buf [8094]byte
			for {
				_, err := c.Read(buf[:4])
				if err != nil {
					break
				}
				length := *(*int32)(unsafe.Pointer(&buf[0]))
				len, err := c.Read(buf[:length])
				if err != nil {
					break
				}
				s := string(buf[:len])
				log.Println(s)
			}
			// c.Close()
			wg.Done()
		}()
	}
	wg.Wait()
}
