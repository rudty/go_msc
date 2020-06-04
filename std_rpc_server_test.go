package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
	"time"
)

type Args struct {
	A int
	B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func newRpcServer() {
	arith := new(Arith)
	fmt.Println(arith)
	rpc.Register(arith)
	// rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error", err)
	}
	// http.Serve(l, nil)
	rpc.Accept(l)
}

func newRpcClient() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dial error", err)
	}
	args := Args{7, 8}
	var reply int
	if err := client.Call("Arith.Multiply", args, &reply); err != nil {
		log.Fatal("arith error", err)
	}
	fmt.Println("response", reply)
}

func TestRpcServer(t *testing.T) {
	go newRpcServer()
	time.Sleep(1 * time.Second)
	newRpcClient()
}
