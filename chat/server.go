package main

import (
	"fmt"
	"log"
	"net"
)

type ChatServer struct {
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}

func sendMessage(clientSocket net.Conn, buf []byte) error {
	var sendIndex int = 0
	for {
		sendLength, err := clientSocket.Write(buf[sendIndex:])
		if err != nil {
			return err
		}

		sendIndex += sendLength

		if sendLength >= len(buf) {
			return nil
		}
	}
}

func onAccept(clientSocket net.Conn) {
	var buf [128]byte
	for {
		len, err := clientSocket.Read(buf[:])
		if err != nil {
			fmt.Println("conn close1")
			return
		}

		fmt.Println(string(buf[0:len]))

		err = sendMessage(clientSocket, buf[0:len])
		if err != nil {
			fmt.Println("conn close2")
			return
		}
	}
}

func (c *ChatServer) Serve() {
	serverSocket, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Panic("tcp open error")
	}

	for {
		clientSocket, err := serverSocket.Accept()
		if err != nil {
			log.Panic("tcp open error")
		}
		go onAccept(clientSocket)
	}
}
