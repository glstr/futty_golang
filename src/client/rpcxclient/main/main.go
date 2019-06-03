package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "address")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func clientStart() {
	log.Printf("enter")
	client := client.NewClient(client.DefaultOption)
	err := client.Connect("tcp", *addr)
	if err != nil {
		log.Printf("error_msg:%s", err.Error())
		return
	}
	defer client.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}
	err = client.Call(context.Background(), "Arith", "Mul", args, reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}

	if reply.C != 200 {
		log.Printf("expect 200 but got %d", reply.C)
	}
	log.Printf("res:%d", reply.C)
	return
}

func main() {
	flag.Parse()
	clientStart()
	return
}
