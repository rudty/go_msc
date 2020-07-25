package main

import (
	"io"
	"log"
	"net"
	"sync"
)

type chatServer struct {
	lock        sync.RWMutex
	clients     map[uint32]*client
	OnConnected func(c *client)
	OnReceive   func(c request)
	OnClosed    func(c *client)
}

type request struct {
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

func (s *chatServer) callCallback(cb func(m request), m request) {
	defer defaultRecover()
	cb(m)
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

func (s *chatServer) receiveFromClient(c *client) {
	var buf [8094]byte
	for {
		l, err := c.ReadMessageInto(buf[:])
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}

		cb := s.OnReceive
		if cb != nil {
			s.callCallback(cb, request{
				Client:  c,
				Receive: buf[:l],
			})
		}
	}
}

func (s *chatServer) onAccept(c *client) {
	s.registerClient(c)
	onConnectedCallback := s.OnConnected
	if onConnectedCallback != nil {
		onConnectedCallback(c)
	}

	s.receiveFromClient(c)

	onClosedCallback := s.OnClosed
	if onClosedCallback != nil {
		onClosedCallback(c)
	}
	s.unRegisterClient(c)
	c.Conn.Close()
}

func (s *chatServer) FindClient(clientID uint32) (c *client, ok bool) {
	s.lock.Lock()
	c, ok = s.clients[clientID]
	s.lock.Unlock()
	return
}

func (s *chatServer) Range(run func(c *client)) {
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
			log.Println("client connect error: ", err)
			continue
		}

		c := newClient(clientSocket)
		go s.onAccept(c)
	}
}
