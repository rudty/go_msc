package main

import (
	"log"
	"net"
	"sync"
	"sync/atomic"
)

type chatServer struct {
	lock      sync.RWMutex
	clients   map[uint32]*client
	OnReceive func(c *client)
}

func newChatServer() *chatServer {
	return &chatServer{}
}

var uniqueID uint32 = 0

func (s *chatServer) callCallback(cb func(c *client), c *client) {
	defer func() {
		recover()
	}()
	cb(c)
}

func (s *chatServer) registerClient(c *client) {
	s.lock.Lock()
	s.clients[c.ClientID] = c
	s.lock.Unlock()
}

func (s *chatServer) onAccept(clientSocket net.Conn) {
	var buf [8094]byte
	clientID := atomic.AddUint32(&uniqueID, 1)

	c := client{
		Conn:     clientSocket,
		ClientID: clientID,
	}

	s.registerClient(&c)

	for {
		len, err := c.Conn.Read(buf[:])
		if err != nil {
			return
		}

		c.Receive = buf[:len]

		cb := s.OnReceive
		if cb != nil {
			s.callCallback(cb, &c)
		}
	}

	clientSocket.Close()
}

func (s *chatServer) Serve() {
	serverSocket, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic("tcp open error")
	}

	for {
		clientSocket, err := serverSocket.Accept()
		if err != nil {
			log.Panic("tcp open error")
		}
		go s.onAccept(clientSocket)
	}
}
