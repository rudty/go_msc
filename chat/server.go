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
	OnReceive func(c clientMessage)
}

type clientMessage struct {
	Client  *client
	Receive []byte
}

func newChatServer() *chatServer {
	s := &chatServer{}
	s.clients = make(map[uint32]*client, 64)
	return s
}

// 아무것도 안함
func defaultRecover() {
	recover()
}

var uniqueID uint32 = 0

func (s *chatServer) callCallback(cb func(c clientMessage), c *client, data []byte) {
	defer defaultRecover()
	cb(clientMessage{
		Client:  c,
		Receive: data,
	})
}

func (s *chatServer) registerClient(c *client) {
	s.lock.Lock()
	s.clients[c.ClientID] = c
	s.lock.Unlock()
}

func (s *chatServer) unRegisterClient(c *client) {
	s.lock.Lock()
	delete(s.clients, c.ClientID)
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
			log.Println(err)
			break
		}

		cb := s.OnReceive
		if cb != nil {
			s.callCallback(cb, &c, buf[:len])
		}
	}

	s.unRegisterClient(&c)

	clientSocket.Close()
}

func (s *chatServer) Do(run func(c *client)) {
	s.lock.Lock()
	w := sync.WaitGroup{}
	for _, e := range s.clients {
		w.Add(1)
		client := e
		go func() {
			defer defaultRecover()
			defer w.Done()
			run(client)
		}()
	}
	w.Wait()
	s.lock.Unlock()
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
