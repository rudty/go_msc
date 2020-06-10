package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

func newRPCServer() {
	auctionService := NewAuctionService()
	auctionService.Start()
	rpc.Register(auctionService)
	// rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error", err)
	}
	// http.Serve(l, nil)
	rpc.Accept(l)
}

func newRPCClient() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dial error", err)
	}
	item := AuctionRegisterItemRequest{
		ItemID:   1,
		BidPrice: 5,
	}
	var reply int64
	if err := client.Call("AuctionServer.RegisterItem", item, &reply); err != nil {
		log.Fatal("query error", err)
	}
	fmt.Println("response", reply)
}

func main() {
	go newRPCServer()
	time.Sleep(1 * time.Second)
	newRPCClient()
}
